// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package ttnmage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

const (
	protocOut             = "/out"
	gogoProtoImage        = "ghcr.io/thethingsindustries/protoc:gen-gogo-1.3.1"
	jsonProtoImage        = "ghcr.io/thethingsindustries/protoc:3.9.1-gen-go-json-1.4.0"
	fieldMaskProtoImage   = "ghcr.io/thethingsindustries/protoc:3.9.1-gen-fieldmask-0.6.2"
	grpcGatewayProtoImage = "ghcr.io/thethingsindustries/protoc:gen-grpc-gateway-1.16.0"
	openAPIv2ProtoImage   = "ghcr.io/thethingsindustries/protoc:gen-grpc-gateway-2.10.3"
	docProtoImage         = "ghcr.io/thethingsindustries/protoc:gen-doc-1.4.1"
	flagProtoImage        = "ghcr.io/thethingsindustries/protoc:3.9.1-gen-go-flags-1.0.6"
)

// Proto namespace.
type Proto mg.Namespace

type protocContext struct {
	WorkingDirectory string
	UID, GID         string
}

func makeProtoc(image string) (func(...string) error, *protocContext, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get working directory: %w", err)
	}
	usr, err := user.Current()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	mountWD := filepath.ToSlash(filepath.Join(filepath.Dir(wd), "lorawan-stack"))
	return sh.RunCmd("docker", "run",
			"--rm",
			"--user", fmt.Sprintf("%s:%s", usr.Uid, usr.Gid),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/api", filepath.Join(wd, "api"), mountWD),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/go.thethings.network/lorawan-stack/v3/pkg/ttnpb", filepath.Join(wd, "pkg", "ttnpb"), protocOut),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/v3/sdk/js", filepath.Join(wd, "sdk", "js"), mountWD),
			"-w", mountWD,
			image,
			fmt.Sprintf("-I%s/api/third_party", mountWD),
			fmt.Sprintf("-I%s", filepath.Dir(wd)),
		), &protocContext{
			WorkingDirectory: mountWD,
			UID:              usr.Uid,
			GID:              usr.Gid,
		}, nil
}

func withProtoc(image string, f func(pCtx *protocContext, protoc func(...string) error) error) error {
	protoc, pCtx, err := makeProtoc(image)
	if err != nil {
		return errors.New("failed to construct protoc command")
	}
	return f(pCtx, protoc)
}

var gogoConvs = func() []string {
	var convs []string
	for _, t := range []string{"any", "duration", "empty", "field_mask", "struct", "timestamp", "wrappers"} {
		convs = append(convs, fmt.Sprintf("Mgoogle/protobuf/%s.proto=github.com/gogo/protobuf/types", t))
	}
	return convs
}()

func (p Proto) gogo(context.Context) error {
	return withProtoc(gogoProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			fmt.Sprintf("--gogo_out=plugins=grpc,%s:%s", strings.Join(gogoConvs, ","), protocOut),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

var jsonConvs = func() []string {
	var convs []string
	for _, t := range gogoConvs {
		convs = append(convs, t+";types")
	}
	return convs
}()

func (p Proto) json(context.Context) error {
	return withProtoc(jsonProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			"--go-json_opt=lang=gogo",
			"--go-json_opt=std=true",
			fmt.Sprintf("--go-json_out=%s:%s", strings.Join(jsonConvs, ","), protocOut),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

func (p Proto) flags(context.Context) error {
	return withProtoc(flagProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			"--go-flags_opt=lang=gogo",
			"--go-flags_opt=customtype.getter-suffix=FromFlag",
			fmt.Sprintf("--go-flags_out=%s:%s", strings.Join(jsonConvs, ","), protocOut),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

func (p Proto) fieldMask(context.Context) error {
	return withProtoc(fieldMaskProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			fmt.Sprintf("--fieldmask_out=lang=gogo,%s:%s", strings.Join(gogoConvs, ","), protocOut),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

func (p Proto) grpcGateway(context.Context) error {
	if err := withProtoc(grpcGatewayProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			fmt.Sprintf("--grpc-gateway_out=allow_delete_body=true,%s:%s", strings.Join(gogoConvs, ","), protocOut),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	}); err != nil {
		return err
	}

	if err := sh.RunV(filepath.Join("tools", "mage", "scripts", "fix-grpc-gateway-names.sh"), "api"); err != nil {
		return fmt.Errorf("failed to fix gRPC-gateway names: %w", err)
	}

	return nil
}

// Go generates Go protos.
func (p Proto) Go(context.Context) error {
	mg.Deps(p.gogo, p.fieldMask, p.grpcGateway, p.json, p.flags)

	ttnpb, err := filepath.Abs(filepath.Join("pkg", "ttnpb"))
	if err != nil {
		return fmt.Errorf("failed to construct absolute path to pkg/ttnpb: %w", err)
	}
	return runGoTool(gofumpt, "-w", ttnpb)
}

// GoClean removes generated Go protos.
func (p Proto) GoClean(context.Context) error {
	return filepath.Walk(filepath.Join("pkg", "ttnpb"), func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ext := range []string{".pb.go", ".pb.gw.go", ".pb.paths.fm.go", ".pb.setters.fm.go", ".pb.validate.go", ".pb.util.fm.go"} {
			if strings.HasSuffix(path, ext) {
				if err := sh.Rm(path); err != nil {
					return err
				}
				return nil
			}
		}
		return nil
	})
}

// Swagger generates Swagger protos.
func (p Proto) Swagger(context.Context) error {
	ok, err := target.Glob(
		filepath.Join("api", "api.swagger.json"),
		filepath.Join("api", "*.proto"),
	)
	if err != nil {
		return targetError(err)
	}
	if !ok {
		return nil
	}
	return withProtoc(openAPIv2ProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			"--openapiv2_opt allow_delete_body=true",
			"--openapiv2_opt=json_names_for_fields=false",
			fmt.Sprintf("--openapiv2_out=allow_merge,merge_file_name=api:%s/api", pCtx.WorkingDirectory),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

// SwaggerClean removes generated Swagger protos.
func (p Proto) SwaggerClean(context.Context) error {
	return sh.Rm(filepath.Join("api", "api.swagger.json"))
}

// Markdown generates Markdown protos.
func (p Proto) Markdown(context.Context) error {
	ok, err := target.Glob(
		filepath.Join("api", "api.md"),
		filepath.Join("api", "*.proto"),
	)
	if err != nil {
		return targetError(err)
	}
	if !ok {
		return nil
	}
	return withProtoc(docProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			fmt.Sprintf("--doc_opt=%s/api/api.md.tmpl,api.md --doc_out=%s/api", pCtx.WorkingDirectory, pCtx.WorkingDirectory),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

// MarkdownClean removes generated Markdown protos.
func (p Proto) MarkdownClean(context.Context) error {
	return sh.Rm(filepath.Join("api", "api.md"))
}

// JsSDK generates javascript SDK protos.
func (p Proto) JsSDK(context.Context) error {
	ok, err := target.Glob(
		filepath.Join("sdk", "js", "generated", "api.json"),
		filepath.Join("api", "*.proto"),
	)
	if err != nil {
		return targetError(err)
	}
	if !ok {
		return nil
	}
	return withProtoc(docProtoImage, func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			fmt.Sprintf("--doc_opt=json,api.json --doc_out=%s/v3/sdk/js/generated", pCtx.WorkingDirectory),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

// JsSDKClean removes generated javascript SDK protos.
func (p Proto) JsSDKClean(context.Context) error {
	return sh.Rm(filepath.Join("sdk", "js", "generated", "api.json"))
}

// All generates protos.
func (p Proto) All(ctx context.Context) {
	mg.CtxDeps(ctx, Proto.Go, Proto.Swagger, Proto.Markdown, Proto.JsSDK)
}

// Clean removes generated protos.
func (p Proto) Clean(ctx context.Context) {
	mg.CtxDeps(ctx, Proto.GoClean, Proto.SwaggerClean, Proto.MarkdownClean, Proto.JsSDKClean)
}
