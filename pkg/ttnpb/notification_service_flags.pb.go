// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.6
// - protoc              v3.9.1
// source: lorawan-stack/api/notification_service.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
)

// AddSetFlagsForListNotificationsRequest adds flags to select fields in ListNotificationsRequest.
func AddSetFlagsForListNotificationsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForUserIdentifiers(flags, flagsplugin.Prefix("receiver-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("status", prefix), flagsplugin.EnumValueDesc(NotificationStatus_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ListNotificationsRequest message from flags.
func (m *ListNotificationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("receiver_ids", prefix)); changed {
		if m.ReceiverIds == nil {
			m.ReceiverIds = &UserIdentifiers{}
		}
		if setPaths, err := m.ReceiverIds.SetFromFlags(flags, flagsplugin.Prefix("receiver_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("status", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Status = make([]NotificationStatus, len(val))
		for i, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, NotificationStatus_value)
			if err != nil {
				return nil, err
			}
			m.Status[i] = NotificationStatus(enumValue)
		}
		paths = append(paths, flagsplugin.Prefix("status", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	return paths, nil
}

// AddSetFlagsForUpdateNotificationStatusRequest adds flags to select fields in UpdateNotificationStatusRequest.
func AddSetFlagsForUpdateNotificationStatusRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForUserIdentifiers(flags, flagsplugin.Prefix("receiver-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("ids", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("status", prefix), flagsplugin.EnumValueDesc(NotificationStatus_value), flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the UpdateNotificationStatusRequest message from flags.
func (m *UpdateNotificationStatusRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("receiver_ids", prefix)); changed {
		if m.ReceiverIds == nil {
			m.ReceiverIds = &UserIdentifiers{}
		}
		if setPaths, err := m.ReceiverIds.SetFromFlags(flags, flagsplugin.Prefix("receiver_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Ids = val
		paths = append(paths, flagsplugin.Prefix("ids", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("status", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, NotificationStatus_value)
		if err != nil {
			return nil, err
		}
		m.Status = NotificationStatus(enumValue)
		paths = append(paths, flagsplugin.Prefix("status", prefix))
	}
	return paths, nil
}
