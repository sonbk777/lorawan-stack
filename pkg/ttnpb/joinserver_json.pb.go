// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.1
// - protoc             v3.9.1
// source: lorawan-stack/api/joinserver.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the SessionKeyRequest message to JSON.
func (x *SessionKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.SessionKeyId) > 0 || s.HasField("session_key_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("session_key_id")
		s.WriteBytes(x.SessionKeyId)
	}
	if len(x.DevEui) > 0 || s.HasField("dev_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_eui")
		types.MarshalHEXBytes(s.WithField("dev_eui"), x.DevEui)
	}
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SessionKeyRequest to JSON.
func (x SessionKeyRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the SessionKeyRequest message from JSON.
func (x *SessionKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "session_key_id", "sessionKeyId":
			s.AddField("session_key_id")
			x.SessionKeyId = s.ReadBytes()
		case "dev_eui", "devEui":
			s.AddField("dev_eui")
			x.DevEui = types.Unmarshal8Bytes(s.WithField("dev_eui", false))
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the SessionKeyRequest from JSON.
func (x *SessionKeyRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the CryptoServicePayloadRequest message to JSON.
func (x *CryptoServicePayloadRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: EndDeviceIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if x.LorawanVersion != 0 || s.HasField("lorawan_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_version")
		x.LorawanVersion.MarshalProtoJSON(s)
	}
	if len(x.Payload) > 0 || s.HasField("payload") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("payload")
		s.WriteBytes(x.Payload)
	}
	if x.ProvisionerId != "" || s.HasField("provisioner_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioner_id")
		s.WriteString(x.ProvisionerId)
	}
	if x.ProvisioningData != nil || s.HasField("provisioning_data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioning_data")
		if x.ProvisioningData == nil {
			s.WriteNil()
		} else {
			gogo.MarshalStruct(s, x.ProvisioningData)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the CryptoServicePayloadRequest to JSON.
func (x CryptoServicePayloadRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the CryptoServicePayloadRequest message from JSON.
func (x *CryptoServicePayloadRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: EndDeviceIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v EndDeviceIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "lorawan_version", "lorawanVersion":
			s.AddField("lorawan_version")
			x.LorawanVersion.UnmarshalProtoJSON(s)
		case "payload":
			s.AddField("payload")
			x.Payload = s.ReadBytes()
		case "provisioner_id", "provisionerId":
			s.AddField("provisioner_id")
			x.ProvisionerId = s.ReadString()
		case "provisioning_data", "provisioningData":
			s.AddField("provisioning_data")
			if s.ReadNil() {
				x.ProvisioningData = nil
				return
			}
			v := gogo.UnmarshalStruct(s)
			if s.Err() != nil {
				return
			}
			x.ProvisioningData = v
		}
	})
}

// UnmarshalJSON unmarshals the CryptoServicePayloadRequest from JSON.
func (x *CryptoServicePayloadRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the JoinAcceptMICRequest message to JSON.
func (x *JoinAcceptMICRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.PayloadRequest != nil || s.HasField("payload_request") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("payload_request")
		x.PayloadRequest.MarshalProtoJSON(s.WithField("payload_request"))
	}
	if x.JoinRequestType != 0 || s.HasField("join_request_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_request_type")
		x.JoinRequestType.MarshalProtoJSON(s)
	}
	if len(x.DevNonce) > 0 || s.HasField("dev_nonce") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_nonce")
		types.MarshalHEXBytes(s.WithField("dev_nonce"), x.DevNonce)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the JoinAcceptMICRequest to JSON.
func (x JoinAcceptMICRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the JoinAcceptMICRequest message from JSON.
func (x *JoinAcceptMICRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "payload_request", "payloadRequest":
			if s.ReadNil() {
				x.PayloadRequest = nil
				return
			}
			x.PayloadRequest = &CryptoServicePayloadRequest{}
			x.PayloadRequest.UnmarshalProtoJSON(s.WithField("payload_request", true))
		case "join_request_type", "joinRequestType":
			s.AddField("join_request_type")
			x.JoinRequestType.UnmarshalProtoJSON(s)
		case "dev_nonce", "devNonce":
			s.AddField("dev_nonce")
			x.DevNonce = types.Unmarshal2Bytes(s.WithField("dev_nonce", false))
		}
	})
}

// UnmarshalJSON unmarshals the JoinAcceptMICRequest from JSON.
func (x *JoinAcceptMICRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the DeriveSessionKeysRequest message to JSON.
func (x *DeriveSessionKeysRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: EndDeviceIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if x.LorawanVersion != 0 || s.HasField("lorawan_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_version")
		x.LorawanVersion.MarshalProtoJSON(s)
	}
	if len(x.JoinNonce) > 0 || s.HasField("join_nonce") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_nonce")
		types.MarshalHEXBytes(s.WithField("join_nonce"), x.JoinNonce)
	}
	if len(x.DevNonce) > 0 || s.HasField("dev_nonce") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_nonce")
		types.MarshalHEXBytes(s.WithField("dev_nonce"), x.DevNonce)
	}
	if len(x.NetId) > 0 || s.HasField("net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("net_id")
		types.MarshalHEXBytes(s.WithField("net_id"), x.NetId)
	}
	if x.ProvisionerId != "" || s.HasField("provisioner_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioner_id")
		s.WriteString(x.ProvisionerId)
	}
	if x.ProvisioningData != nil || s.HasField("provisioning_data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioning_data")
		if x.ProvisioningData == nil {
			s.WriteNil()
		} else {
			gogo.MarshalStruct(s, x.ProvisioningData)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the DeriveSessionKeysRequest to JSON.
func (x DeriveSessionKeysRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the DeriveSessionKeysRequest message from JSON.
func (x *DeriveSessionKeysRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: EndDeviceIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v EndDeviceIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "lorawan_version", "lorawanVersion":
			s.AddField("lorawan_version")
			x.LorawanVersion.UnmarshalProtoJSON(s)
		case "join_nonce", "joinNonce":
			s.AddField("join_nonce")
			x.JoinNonce = types.Unmarshal3Bytes(s.WithField("join_nonce", false))
		case "dev_nonce", "devNonce":
			s.AddField("dev_nonce")
			x.DevNonce = types.Unmarshal2Bytes(s.WithField("dev_nonce", false))
		case "net_id", "netId":
			s.AddField("net_id")
			x.NetId = types.Unmarshal3Bytes(s.WithField("net_id", false))
		case "provisioner_id", "provisionerId":
			s.AddField("provisioner_id")
			x.ProvisionerId = s.ReadString()
		case "provisioning_data", "provisioningData":
			s.AddField("provisioning_data")
			if s.ReadNil() {
				x.ProvisioningData = nil
				return
			}
			v := gogo.UnmarshalStruct(s)
			if s.Err() != nil {
				return
			}
			x.ProvisioningData = v
		}
	})
}

// UnmarshalJSON unmarshals the DeriveSessionKeysRequest from JSON.
func (x *DeriveSessionKeysRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ProvisionEndDevicesRequest_IdentifiersList message to JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersList) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	if len(x.EndDeviceIds) > 0 || s.HasField("end_device_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("end_device_ids")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.EndDeviceIds {
			s.WriteMoreIf(&wroteElement)
			// NOTE: EndDeviceIdentifiers does not seem to implement MarshalProtoJSON.
			gogo.MarshalMessage(s, element)
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ProvisionEndDevicesRequest_IdentifiersList to JSON.
func (x ProvisionEndDevicesRequest_IdentifiersList) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the ProvisionEndDevicesRequest_IdentifiersList message from JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersList) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		case "end_device_ids", "endDeviceIds":
			s.AddField("end_device_ids")
			if s.ReadNil() {
				x.EndDeviceIds = nil
				return
			}
			s.ReadArray(func() {
				// NOTE: EndDeviceIdentifiers does not seem to implement UnmarshalProtoJSON.
				var v EndDeviceIdentifiers
				gogo.UnmarshalMessage(s, &v)
				x.EndDeviceIds = append(x.EndDeviceIds, &v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the ProvisionEndDevicesRequest_IdentifiersList from JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersList) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ProvisionEndDevicesRequest_IdentifiersRange message to JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersRange) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	if len(x.StartDevEui) > 0 || s.HasField("start_dev_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("start_dev_eui")
		types.MarshalHEXBytes(s.WithField("start_dev_eui"), x.StartDevEui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ProvisionEndDevicesRequest_IdentifiersRange to JSON.
func (x ProvisionEndDevicesRequest_IdentifiersRange) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the ProvisionEndDevicesRequest_IdentifiersRange message from JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersRange) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		case "start_dev_eui", "startDevEui":
			s.AddField("start_dev_eui")
			x.StartDevEui = types.Unmarshal8Bytes(s.WithField("start_dev_eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the ProvisionEndDevicesRequest_IdentifiersRange from JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersRange) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ProvisionEndDevicesRequest_IdentifiersFromData message to JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersFromData) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ProvisionEndDevicesRequest_IdentifiersFromData to JSON.
func (x ProvisionEndDevicesRequest_IdentifiersFromData) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the ProvisionEndDevicesRequest_IdentifiersFromData message from JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersFromData) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the ProvisionEndDevicesRequest_IdentifiersFromData from JSON.
func (x *ProvisionEndDevicesRequest_IdentifiersFromData) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ProvisionEndDevicesRequest message to JSON.
func (x *ProvisionEndDevicesRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ApplicationIds)
	}
	if x.ProvisionerId != "" || s.HasField("provisioner_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioner_id")
		s.WriteString(x.ProvisionerId)
	}
	if len(x.ProvisioningData) > 0 || s.HasField("provisioning_data") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("provisioning_data")
		s.WriteBytes(x.ProvisioningData)
	}
	if x.EndDevices != nil {
		switch ov := x.EndDevices.(type) {
		case *ProvisionEndDevicesRequest_List:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("list")
			ov.List.MarshalProtoJSON(s.WithField("list"))
		case *ProvisionEndDevicesRequest_Range:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("range")
			ov.Range.MarshalProtoJSON(s.WithField("range"))
		case *ProvisionEndDevicesRequest_FromData:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("from_data")
			ov.FromData.MarshalProtoJSON(s.WithField("from_data"))
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ProvisionEndDevicesRequest to JSON.
func (x ProvisionEndDevicesRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the ProvisionEndDevicesRequest message from JSON.
func (x *ProvisionEndDevicesRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "provisioner_id", "provisionerId":
			s.AddField("provisioner_id")
			x.ProvisionerId = s.ReadString()
		case "provisioning_data", "provisioningData":
			s.AddField("provisioning_data")
			x.ProvisioningData = s.ReadBytes()
		case "list":
			ov := &ProvisionEndDevicesRequest_List{}
			x.EndDevices = ov
			if s.ReadNil() {
				ov.List = nil
				return
			}
			ov.List = &ProvisionEndDevicesRequest_IdentifiersList{}
			ov.List.UnmarshalProtoJSON(s.WithField("list", true))
		case "range":
			ov := &ProvisionEndDevicesRequest_Range{}
			x.EndDevices = ov
			if s.ReadNil() {
				ov.Range = nil
				return
			}
			ov.Range = &ProvisionEndDevicesRequest_IdentifiersRange{}
			ov.Range.UnmarshalProtoJSON(s.WithField("range", true))
		case "from_data", "fromData":
			ov := &ProvisionEndDevicesRequest_FromData{}
			x.EndDevices = ov
			if s.ReadNil() {
				ov.FromData = nil
				return
			}
			ov.FromData = &ProvisionEndDevicesRequest_IdentifiersFromData{}
			ov.FromData.UnmarshalProtoJSON(s.WithField("from_data", true))
		}
	})
}

// UnmarshalJSON unmarshals the ProvisionEndDevicesRequest from JSON.
func (x *ProvisionEndDevicesRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ApplicationActivationSettings message to JSON.
func (x *ApplicationActivationSettings) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.KekLabel != "" || s.HasField("kek_label") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("kek_label")
		s.WriteString(x.KekLabel)
	}
	if x.Kek != nil || s.HasField("kek") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("kek")
		// NOTE: KeyEnvelope does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Kek)
	}
	if len(x.HomeNetId) > 0 || s.HasField("home_net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("home_net_id")
		types.MarshalHEXBytes(s.WithField("home_net_id"), x.HomeNetId)
	}
	if x.ApplicationServerId != "" || s.HasField("application_server_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_server_id")
		s.WriteString(x.ApplicationServerId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ApplicationActivationSettings to JSON.
func (x ApplicationActivationSettings) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the ApplicationActivationSettings message from JSON.
func (x *ApplicationActivationSettings) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "kek_label", "kekLabel":
			s.AddField("kek_label")
			x.KekLabel = s.ReadString()
		case "kek":
			s.AddField("kek")
			if s.ReadNil() {
				x.Kek = nil
				return
			}
			// NOTE: KeyEnvelope does not seem to implement UnmarshalProtoJSON.
			var v KeyEnvelope
			gogo.UnmarshalMessage(s, &v)
			x.Kek = &v
		case "home_net_id", "homeNetId":
			s.AddField("home_net_id")
			x.HomeNetId = types.Unmarshal3Bytes(s.WithField("home_net_id", false))
		case "application_server_id", "applicationServerId":
			s.AddField("application_server_id")
			x.ApplicationServerId = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the ApplicationActivationSettings from JSON.
func (x *ApplicationActivationSettings) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SetApplicationActivationSettingsRequest message to JSON.
func (x *SetApplicationActivationSettingsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ApplicationIds)
	}
	if x.Settings != nil || s.HasField("settings") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("settings")
		x.Settings.MarshalProtoJSON(s.WithField("settings"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SetApplicationActivationSettingsRequest to JSON.
func (x SetApplicationActivationSettingsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the SetApplicationActivationSettingsRequest message from JSON.
func (x *SetApplicationActivationSettingsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "settings":
			if s.ReadNil() {
				x.Settings = nil
				return
			}
			x.Settings = &ApplicationActivationSettings{}
			x.Settings.UnmarshalProtoJSON(s.WithField("settings", true))
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// UnmarshalJSON unmarshals the SetApplicationActivationSettingsRequest from JSON.
func (x *SetApplicationActivationSettingsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the JoinEUIPrefix message to JSON.
func (x *JoinEUIPrefix) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	if x.Length != 0 || s.HasField("length") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("length")
		s.WriteUint32(x.Length)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the JoinEUIPrefix to JSON.
func (x JoinEUIPrefix) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the JoinEUIPrefix message from JSON.
func (x *JoinEUIPrefix) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		case "length":
			s.AddField("length")
			x.Length = s.ReadUint32()
		}
	})
}

// UnmarshalJSON unmarshals the JoinEUIPrefix from JSON.
func (x *JoinEUIPrefix) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the JoinEUIPrefixes message to JSON.
func (x *JoinEUIPrefixes) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Prefixes) > 0 || s.HasField("prefixes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("prefixes")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Prefixes {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("prefixes"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the JoinEUIPrefixes to JSON.
func (x JoinEUIPrefixes) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the JoinEUIPrefixes message from JSON.
func (x *JoinEUIPrefixes) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "prefixes":
			s.AddField("prefixes")
			if s.ReadNil() {
				x.Prefixes = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Prefixes = append(x.Prefixes, nil)
					return
				}
				v := &JoinEUIPrefix{}
				v.UnmarshalProtoJSON(s.WithField("prefixes", false))
				if s.Err() != nil {
					return
				}
				x.Prefixes = append(x.Prefixes, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the JoinEUIPrefixes from JSON.
func (x *JoinEUIPrefixes) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetDefaultJoinEUIResponse message to JSON.
func (x *GetDefaultJoinEUIResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetDefaultJoinEUIResponse to JSON.
func (x GetDefaultJoinEUIResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the GetDefaultJoinEUIResponse message from JSON.
func (x *GetDefaultJoinEUIResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the GetDefaultJoinEUIResponse from JSON.
func (x *GetDefaultJoinEUIResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
