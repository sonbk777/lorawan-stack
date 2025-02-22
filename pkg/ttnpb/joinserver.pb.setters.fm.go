// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *SessionKeyRequest) SetFields(src *SessionKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "session_key_id":
			if len(subs) > 0 {
				return fmt.Errorf("'session_key_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SessionKeyId = src.SessionKeyId
			} else {
				dst.SessionKeyId = nil
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEui = src.DevEui
			} else {
				dst.DevEui = nil
			}
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *NwkSKeysResponse) SetFields(src *NwkSKeysResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "f_nwk_s_int_key":
			if len(subs) > 0 {
				var newDst, newSrc *KeyEnvelope
				if (src == nil || src.FNwkSIntKey == nil) && dst.FNwkSIntKey == nil {
					continue
				}
				if src != nil {
					newSrc = src.FNwkSIntKey
				}
				if dst.FNwkSIntKey != nil {
					newDst = dst.FNwkSIntKey
				} else {
					newDst = &KeyEnvelope{}
					dst.FNwkSIntKey = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.FNwkSIntKey = src.FNwkSIntKey
				} else {
					dst.FNwkSIntKey = nil
				}
			}
		case "s_nwk_s_int_key":
			if len(subs) > 0 {
				var newDst, newSrc *KeyEnvelope
				if (src == nil || src.SNwkSIntKey == nil) && dst.SNwkSIntKey == nil {
					continue
				}
				if src != nil {
					newSrc = src.SNwkSIntKey
				}
				if dst.SNwkSIntKey != nil {
					newDst = dst.SNwkSIntKey
				} else {
					newDst = &KeyEnvelope{}
					dst.SNwkSIntKey = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SNwkSIntKey = src.SNwkSIntKey
				} else {
					dst.SNwkSIntKey = nil
				}
			}
		case "nwk_s_enc_key":
			if len(subs) > 0 {
				var newDst, newSrc *KeyEnvelope
				if (src == nil || src.NwkSEncKey == nil) && dst.NwkSEncKey == nil {
					continue
				}
				if src != nil {
					newSrc = src.NwkSEncKey
				}
				if dst.NwkSEncKey != nil {
					newDst = dst.NwkSEncKey
				} else {
					newDst = &KeyEnvelope{}
					dst.NwkSEncKey = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.NwkSEncKey = src.NwkSEncKey
				} else {
					dst.NwkSEncKey = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AppSKeyResponse) SetFields(src *AppSKeyResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "app_s_key":
			if len(subs) > 0 {
				var newDst, newSrc *KeyEnvelope
				if (src == nil || src.AppSKey == nil) && dst.AppSKey == nil {
					continue
				}
				if src != nil {
					newSrc = src.AppSKey
				}
				if dst.AppSKey != nil {
					newDst = dst.AppSKey
				} else {
					newDst = &KeyEnvelope{}
					dst.AppSKey = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AppSKey = src.AppSKey
				} else {
					dst.AppSKey = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CryptoServicePayloadRequest) SetFields(src *CryptoServicePayloadRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "lorawan_version":
			if len(subs) > 0 {
				return fmt.Errorf("'lorawan_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LorawanVersion = src.LorawanVersion
			} else {
				var zero MACVersion
				dst.LorawanVersion = zero
			}
		case "payload":
			if len(subs) > 0 {
				return fmt.Errorf("'payload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Payload = src.Payload
			} else {
				dst.Payload = nil
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerId = src.ProvisionerId
			} else {
				var zero string
				dst.ProvisionerId = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CryptoServicePayloadResponse) SetFields(src *CryptoServicePayloadResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "payload":
			if len(subs) > 0 {
				return fmt.Errorf("'payload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Payload = src.Payload
			} else {
				dst.Payload = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *JoinAcceptMICRequest) SetFields(src *JoinAcceptMICRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "payload_request":
			if len(subs) > 0 {
				var newDst, newSrc *CryptoServicePayloadRequest
				if (src == nil || src.PayloadRequest == nil) && dst.PayloadRequest == nil {
					continue
				}
				if src != nil {
					newSrc = src.PayloadRequest
				}
				if dst.PayloadRequest != nil {
					newDst = dst.PayloadRequest
				} else {
					newDst = &CryptoServicePayloadRequest{}
					dst.PayloadRequest = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.PayloadRequest = src.PayloadRequest
				} else {
					dst.PayloadRequest = nil
				}
			}
		case "join_request_type":
			if len(subs) > 0 {
				return fmt.Errorf("'join_request_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinRequestType = src.JoinRequestType
			} else {
				var zero JoinRequestType
				dst.JoinRequestType = zero
			}
		case "dev_nonce":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_nonce' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevNonce = src.DevNonce
			} else {
				dst.DevNonce = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeriveSessionKeysRequest) SetFields(src *DeriveSessionKeysRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "lorawan_version":
			if len(subs) > 0 {
				return fmt.Errorf("'lorawan_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LorawanVersion = src.LorawanVersion
			} else {
				var zero MACVersion
				dst.LorawanVersion = zero
			}
		case "join_nonce":
			if len(subs) > 0 {
				return fmt.Errorf("'join_nonce' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinNonce = src.JoinNonce
			} else {
				dst.JoinNonce = nil
			}
		case "dev_nonce":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_nonce' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevNonce = src.DevNonce
			} else {
				dst.DevNonce = nil
			}
		case "net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NetId = src.NetId
			} else {
				dst.NetId = nil
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerId = src.ProvisionerId
			} else {
				var zero string
				dst.ProvisionerId = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRootKeysRequest) SetFields(src *GetRootKeysRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerId = src.ProvisionerId
			} else {
				var zero string
				dst.ProvisionerId = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest) SetFields(src *ProvisionEndDevicesRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "provisioner_id":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioner_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisionerId = src.ProvisionerId
			} else {
				var zero string
				dst.ProvisionerId = zero
			}
		case "provisioning_data":
			if len(subs) > 0 {
				return fmt.Errorf("'provisioning_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ProvisioningData = src.ProvisioningData
			} else {
				dst.ProvisioningData = nil
			}

		case "end_devices":
			if len(subs) == 0 && src == nil {
				dst.EndDevices = nil
				continue
			} else if len(subs) == 0 {
				dst.EndDevices = src.EndDevices
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "list":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.EndDevices.(*ProvisionEndDevicesRequest_List)
					}
					if srcValid := srcTypeOk || src == nil || src.EndDevices == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'list', while different oneof is set in source")
					}
					_, dstTypeOk := dst.EndDevices.(*ProvisionEndDevicesRequest_List)
					if dstValid := dstTypeOk || dst.EndDevices == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'list', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ProvisionEndDevicesRequest_IdentifiersList
						if srcTypeOk {
							newSrc = src.EndDevices.(*ProvisionEndDevicesRequest_List).List
						}
						if dstTypeOk {
							newDst = dst.EndDevices.(*ProvisionEndDevicesRequest_List).List
						} else if srcTypeOk {
							newDst = &ProvisionEndDevicesRequest_IdentifiersList{}
							dst.EndDevices = &ProvisionEndDevicesRequest_List{List: newDst}
						} else {
							dst.EndDevices = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.EndDevices = src.EndDevices
						} else {
							dst.EndDevices = nil
						}
					}
				case "range":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.EndDevices.(*ProvisionEndDevicesRequest_Range)
					}
					if srcValid := srcTypeOk || src == nil || src.EndDevices == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'range', while different oneof is set in source")
					}
					_, dstTypeOk := dst.EndDevices.(*ProvisionEndDevicesRequest_Range)
					if dstValid := dstTypeOk || dst.EndDevices == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'range', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ProvisionEndDevicesRequest_IdentifiersRange
						if srcTypeOk {
							newSrc = src.EndDevices.(*ProvisionEndDevicesRequest_Range).Range
						}
						if dstTypeOk {
							newDst = dst.EndDevices.(*ProvisionEndDevicesRequest_Range).Range
						} else if srcTypeOk {
							newDst = &ProvisionEndDevicesRequest_IdentifiersRange{}
							dst.EndDevices = &ProvisionEndDevicesRequest_Range{Range: newDst}
						} else {
							dst.EndDevices = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.EndDevices = src.EndDevices
						} else {
							dst.EndDevices = nil
						}
					}
				case "from_data":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.EndDevices.(*ProvisionEndDevicesRequest_FromData)
					}
					if srcValid := srcTypeOk || src == nil || src.EndDevices == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'from_data', while different oneof is set in source")
					}
					_, dstTypeOk := dst.EndDevices.(*ProvisionEndDevicesRequest_FromData)
					if dstValid := dstTypeOk || dst.EndDevices == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'from_data', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ProvisionEndDevicesRequest_IdentifiersFromData
						if srcTypeOk {
							newSrc = src.EndDevices.(*ProvisionEndDevicesRequest_FromData).FromData
						}
						if dstTypeOk {
							newDst = dst.EndDevices.(*ProvisionEndDevicesRequest_FromData).FromData
						} else if srcTypeOk {
							newDst = &ProvisionEndDevicesRequest_IdentifiersFromData{}
							dst.EndDevices = &ProvisionEndDevicesRequest_FromData{FromData: newDst}
						} else {
							dst.EndDevices = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.EndDevices = src.EndDevices
						} else {
							dst.EndDevices = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationActivationSettings) SetFields(src *ApplicationActivationSettings, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "kek_label":
			if len(subs) > 0 {
				return fmt.Errorf("'kek_label' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.KekLabel = src.KekLabel
			} else {
				var zero string
				dst.KekLabel = zero
			}
		case "kek":
			if len(subs) > 0 {
				var newDst, newSrc *KeyEnvelope
				if (src == nil || src.Kek == nil) && dst.Kek == nil {
					continue
				}
				if src != nil {
					newSrc = src.Kek
				}
				if dst.Kek != nil {
					newDst = dst.Kek
				} else {
					newDst = &KeyEnvelope{}
					dst.Kek = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Kek = src.Kek
				} else {
					dst.Kek = nil
				}
			}
		case "home_net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'home_net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HomeNetId = src.HomeNetId
			} else {
				dst.HomeNetId = nil
			}
		case "application_server_id":
			if len(subs) > 0 {
				return fmt.Errorf("'application_server_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationServerId = src.ApplicationServerId
			} else {
				var zero string
				dst.ApplicationServerId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationActivationSettingsRequest) SetFields(src *GetApplicationActivationSettingsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetApplicationActivationSettingsRequest) SetFields(src *SetApplicationActivationSettingsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "settings":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationActivationSettings
				if (src == nil || src.Settings == nil) && dst.Settings == nil {
					continue
				}
				if src != nil {
					newSrc = src.Settings
				}
				if dst.Settings != nil {
					newDst = dst.Settings
				} else {
					newDst = &ApplicationActivationSettings{}
					dst.Settings = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Settings = src.Settings
				} else {
					dst.Settings = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteApplicationActivationSettingsRequest) SetFields(src *DeleteApplicationActivationSettingsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *JoinEUIPrefix) SetFields(src *JoinEUIPrefix, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}
		case "length":
			if len(subs) > 0 {
				return fmt.Errorf("'length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Length = src.Length
			} else {
				var zero uint32
				dst.Length = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *JoinEUIPrefixes) SetFields(src *JoinEUIPrefixes, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "prefixes":
			if len(subs) > 0 {
				return fmt.Errorf("'prefixes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Prefixes = src.Prefixes
			} else {
				dst.Prefixes = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetDefaultJoinEUIResponse) SetFields(src *GetDefaultJoinEUIResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest_IdentifiersList) SetFields(src *ProvisionEndDevicesRequest_IdentifiersList, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}
		case "end_device_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'end_device_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.EndDeviceIds = src.EndDeviceIds
			} else {
				dst.EndDeviceIds = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest_IdentifiersRange) SetFields(src *ProvisionEndDevicesRequest_IdentifiersRange, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}
		case "start_dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'start_dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.StartDevEui = src.StartDevEui
			} else {
				dst.StartDevEui = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ProvisionEndDevicesRequest_IdentifiersFromData) SetFields(src *ProvisionEndDevicesRequest_IdentifiersFromData, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
