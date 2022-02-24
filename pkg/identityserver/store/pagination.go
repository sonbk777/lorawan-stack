// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"
	"fmt"
	"strings"
)

type paginationOptionsKeyType struct{}

var paginationOptionsKey paginationOptionsKeyType

type PaginationOptions struct {
	limit  uint32
	offset uint32
	total  *uint64
}

// WithPagination instructs the store to paginate the results, and set the total
// number of results into total.
func WithPagination(ctx context.Context, limit, page uint32, total *uint64) context.Context {
	if page == 0 {
		page = 1
	}
	return context.WithValue(ctx, paginationOptionsKey, PaginationOptions{
		limit:  limit,
		offset: (page - 1) * limit,
		total:  total,
	})
}

// SetTotal sets the total number of results into the destination set by
// SetTotalCount if not already set.
func SetTotal(ctx context.Context, total uint64) {
	if opts, ok := ctx.Value(paginationOptionsKey).(PaginationOptions); ok && opts.total != nil && *opts.total == 0 {
		*opts.total = total
	}
}

// LimitAndOffsetFromContext gets the limit and offset from the context.
func LimitAndOffsetFromContext(ctx context.Context) (limit, offset uint32) {
	if opts, ok := ctx.Value(paginationOptionsKey).(PaginationOptions); ok {
		return opts.limit, opts.offset
	}
	return 0, 0
}

// WithOrder instructs the store to sort the results by the given field.
// If the field is prefixed with a minus, the order is reversed.
func WithOrder(ctx context.Context, spec string) context.Context {
	if spec == "" {
		return ctx
	}
	field := spec
	order := "ASC"
	if strings.HasPrefix(spec, "-") {
		field = strings.TrimPrefix(spec, "-")
		order = "DESC"
	}
	return context.WithValue(ctx, orderOptionsKey, OrderOptions{
		field: field,
		order: order,
	})
}

type orderOptionsKeyType struct{}

var orderOptionsKey orderOptionsKeyType

type OrderOptions struct {
	field string
	order string
}

func OrderFromContext(ctx context.Context, table, defaultTableField, defaultOrder string) string {
	if opts, ok := ctx.Value(orderOptionsKey).(OrderOptions); ok && opts.field != "" {
		order := opts.order
		if order == "" {
			order = "ASC"
		}
		if (table == "organizations" && opts.field == "organization_id") || (table == "users" && opts.field == "user_id") {
			table = "accounts"
			opts.field = "uid"
		}
		tableField := fmt.Sprintf(`"%s"."%s"`, table, opts.field)
		if tableField != defaultTableField && opts.field != defaultTableField {
			return fmt.Sprintf(`%s %s, %s %s`, tableField, order, defaultTableField, order)
		}
		return fmt.Sprintf(`%s %s`, tableField, order)
	}
	return fmt.Sprintf("%s %s", defaultTableField, defaultOrder)
}
