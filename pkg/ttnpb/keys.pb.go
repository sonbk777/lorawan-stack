// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/keys.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type KeyEnvelope struct {
	// The unencrypted AES key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The label of the RFC 3394 key-encryption-key (KEK) that was used to encrypt the key.
	KekLabel             string   `protobuf:"bytes,2,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	EncryptedKey         []byte   `protobuf:"bytes,3,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{0}
}
func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

// Root keys for a LoRaWAN device.
// These are stored on the Join Server.
type RootKeys struct {
	// Join Server issued identifier for the root keys.
	RootKeyId string `protobuf:"bytes,1,opt,name=root_key_id,json=rootKeyId,proto3" json:"root_key_id,omitempty"`
	// The (encrypted) Application Key.
	AppKey *KeyEnvelope `protobuf:"bytes,2,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	// The (encrypted) Network Key.
	NwkKey               *KeyEnvelope `protobuf:"bytes,3,opt,name=nwk_key,json=nwkKey,proto3" json:"nwk_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RootKeys) Reset()         { *m = RootKeys{} }
func (m *RootKeys) String() string { return proto.CompactTextString(m) }
func (*RootKeys) ProtoMessage()    {}
func (*RootKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{1}
}
func (m *RootKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootKeys.Unmarshal(m, b)
}
func (m *RootKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootKeys.Marshal(b, m, deterministic)
}
func (m *RootKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootKeys.Merge(m, src)
}
func (m *RootKeys) XXX_Size() int {
	return xxx_messageInfo_RootKeys.Size(m)
}
func (m *RootKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_RootKeys.DiscardUnknown(m)
}

var xxx_messageInfo_RootKeys proto.InternalMessageInfo

func (m *RootKeys) GetRootKeyId() string {
	if m != nil {
		return m.RootKeyId
	}
	return ""
}

func (m *RootKeys) GetAppKey() *KeyEnvelope {
	if m != nil {
		return m.AppKey
	}
	return nil
}

func (m *RootKeys) GetNwkKey() *KeyEnvelope {
	if m != nil {
		return m.NwkKey
	}
	return nil
}

// Session keys for a LoRaWAN session.
// Only the components for which the keys were meant, will have the key-encryption-key (KEK) to decrypt the individual keys.
type SessionKeys struct {
	// Join Server issued identifier for the session keys.
	// This ID can be used to request the keys from the Join Server in case the are lost.
	SessionKeyId []byte `protobuf:"bytes,1,opt,name=session_key_id,json=sessionKeyId,proto3" json:"session_key_id,omitempty"`
	// The (encrypted) Forwarding Network Session Integrity Key (or Network Session Key in 1.0 compatibility mode).
	// This key is stored by the (forwarding) Network Server.
	FNwkSIntKey *KeyEnvelope `protobuf:"bytes,2,opt,name=f_nwk_s_int_key,json=fNwkSIntKey,proto3" json:"f_nwk_s_int_key,omitempty"`
	// The (encrypted) Serving Network Session Integrity Key.
	// This key is stored by the (serving) Network Server.
	SNwkSIntKey *KeyEnvelope `protobuf:"bytes,3,opt,name=s_nwk_s_int_key,json=sNwkSIntKey,proto3" json:"s_nwk_s_int_key,omitempty"`
	// The (encrypted) Network Session Encryption Key.
	// This key is stored by the (serving) Network Server.
	NwkSEncKey *KeyEnvelope `protobuf:"bytes,4,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	// The (encrypted) Application Session Key.
	// This key is stored by the Application Server.
	AppSKey              *KeyEnvelope `protobuf:"bytes,5,opt,name=app_s_key,json=appSKey,proto3" json:"app_s_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SessionKeys) Reset()         { *m = SessionKeys{} }
func (m *SessionKeys) String() string { return proto.CompactTextString(m) }
func (*SessionKeys) ProtoMessage()    {}
func (*SessionKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{2}
}
func (m *SessionKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionKeys.Unmarshal(m, b)
}
func (m *SessionKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionKeys.Marshal(b, m, deterministic)
}
func (m *SessionKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionKeys.Merge(m, src)
}
func (m *SessionKeys) XXX_Size() int {
	return xxx_messageInfo_SessionKeys.Size(m)
}
func (m *SessionKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionKeys.DiscardUnknown(m)
}

var xxx_messageInfo_SessionKeys proto.InternalMessageInfo

func (m *SessionKeys) GetSessionKeyId() []byte {
	if m != nil {
		return m.SessionKeyId
	}
	return nil
}

func (m *SessionKeys) GetFNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.FNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetSNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.SNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetNwkSEncKey() *KeyEnvelope {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func (m *SessionKeys) GetAppSKey() *KeyEnvelope {
	if m != nil {
		return m.AppSKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyEnvelope)(nil), "ttn.lorawan.v3.KeyEnvelope")
	golang_proto.RegisterType((*KeyEnvelope)(nil), "ttn.lorawan.v3.KeyEnvelope")
	proto.RegisterType((*RootKeys)(nil), "ttn.lorawan.v3.RootKeys")
	golang_proto.RegisterType((*RootKeys)(nil), "ttn.lorawan.v3.RootKeys")
	proto.RegisterType((*SessionKeys)(nil), "ttn.lorawan.v3.SessionKeys")
	golang_proto.RegisterType((*SessionKeys)(nil), "ttn.lorawan.v3.SessionKeys")
}

func init() { proto.RegisterFile("lorawan-stack/api/keys.proto", fileDescriptor_ee170ee4ccd55993) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/keys.proto", fileDescriptor_ee170ee4ccd55993)
}

var fileDescriptor_ee170ee4ccd55993 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xbf, 0x6f, 0x1a, 0x49,
	0x14, 0xc7, 0xb5, 0x60, 0x30, 0x0c, 0xd8, 0x67, 0x6d, 0x85, 0xb9, 0xd3, 0xc9, 0x42, 0x3a, 0xc9,
	0x0d, 0xbb, 0x67, 0xe3, 0x1f, 0x77, 0x57, 0xdc, 0x1d, 0x6b, 0x63, 0x9b, 0xe3, 0x42, 0x24, 0x70,
	0xa4, 0x28, 0x0d, 0x1a, 0x96, 0xc7, 0xb2, 0xd9, 0x65, 0x66, 0xb4, 0x33, 0xb0, 0x5e, 0x57, 0x96,
	0xd2, 0xa7, 0x48, 0x97, 0x74, 0x49, 0x91, 0x82, 0xbf, 0x22, 0x55, 0xaa, 0xfc, 0x17, 0xe9, 0x92,
	0xce, 0xa5, 0x8b, 0x28, 0x9a, 0x01, 0x39, 0x8b, 0x1d, 0xc9, 0x76, 0xdc, 0xa0, 0x7d, 0xe2, 0xfb,
	0x3e, 0x7c, 0xbf, 0xef, 0x3d, 0x16, 0xfd, 0xe2, 0xd3, 0x00, 0x87, 0x98, 0x94, 0xb9, 0xc0, 0xb6,
	0x67, 0x62, 0xe6, 0x9a, 0x1e, 0x44, 0xdc, 0x60, 0x01, 0x15, 0x54, 0x5f, 0x16, 0x82, 0x18, 0x33,
	0x85, 0x31, 0xae, 0x14, 0xab, 0x8e, 0x2b, 0x06, 0xa3, 0xae, 0x61, 0xd3, 0xa1, 0x09, 0x64, 0x4c,
	0x23, 0x16, 0xd0, 0x93, 0xc8, 0x54, 0x62, 0xbb, 0xec, 0x00, 0x29, 0x8f, 0xb1, 0xef, 0xf6, 0xb0,
	0x00, 0xf3, 0xda, 0xc3, 0x14, 0x59, 0x2c, 0xc7, 0x10, 0x0e, 0x75, 0xe8, 0xb4, 0xb9, 0x3b, 0xea,
	0xab, 0x4a, 0x15, 0xea, 0x69, 0x26, 0xdf, 0x8f, 0xc9, 0x8f, 0x07, 0x70, 0x3c, 0x70, 0x89, 0xc3,
	0xeb, 0xa4, 0x37, 0xe2, 0x22, 0x70, 0x81, 0xc7, 0x7f, 0xda, 0xa1, 0xe5, 0xbe, 0x8f, 0x1d, 0x6e,
	0x62, 0x42, 0xa8, 0xc0, 0xc2, 0xa5, 0x64, 0x96, 0xa3, 0xb8, 0x77, 0x27, 0xca, 0x53, 0x4e, 0xc9,
	0x77, 0x20, 0x46, 0x4c, 0x45, 0x19, 0x10, 0xcc, 0xdc, 0xf1, 0xa6, 0x49, 0x99, 0xd2, 0x5c, 0xd7,
	0x97, 0x9e, 0xa7, 0x50, 0xae, 0x01, 0x51, 0x8d, 0x8c, 0xc1, 0xa7, 0x0c, 0xf4, 0x2f, 0x09, 0x94,
	0xf4, 0x20, 0x2a, 0x68, 0x6b, 0xda, 0x7a, 0xde, 0xfa, 0x9c, 0x78, 0x51, 0xdd, 0xf8, 0xaf, 0x54,
	0xfa, 0x7d, 0x63, 0xb3, 0xb2, 0xb5, 0xbd, 0xb3, 0xfb, 0xc7, 0x9f, 0x55, 0x6b, 0x6f, 0xbf, 0x76,
	0x70, 0xb5, 0x2e, 0xbd, 0x4a, 0x68, 0x8b, 0x6f, 0x12, 0x69, 0xe9, 0x94, 0x38, 0x17, 0x56, 0xfa,
	0x74, 0x61, 0xb0, 0xc2, 0xb4, 0x4f, 0x93, 0xd5, 0x67, 0x1a, 0xfa, 0xc7, 0xa1, 0x86, 0x18, 0x80,
	0x50, 0x79, 0x0c, 0x02, 0x22, 0xa4, 0x81, 0x67, 0xce, 0xaf, 0x74, 0x5c, 0x31, 0x99, 0xe7, 0x98,
	0x22, 0x62, 0xc0, 0x8d, 0x07, 0x38, 0xe0, 0x03, 0xec, 0x1f, 0xd5, 0x1e, 0x5b, 0x91, 0x00, 0xae,
	0xff, 0x7b, 0x57, 0xc0, 0x23, 0x32, 0x9c, 0x22, 0x36, 0x76, 0x14, 0xe1, 0x7c, 0xb2, 0xfa, 0x5a,
	0x2b, 0x3e, 0xbc, 0x1d, 0xc6, 0x1e, 0xf6, 0x4c, 0x21, 0x48, 0xd9, 0x0f, 0xcb, 0xb6, 0xef, 0x9a,
	0xf6, 0x88, 0x0b, 0x3a, 0x54, 0x3b, 0x33, 0x9a, 0x10, 0xce, 0x88, 0x07, 0x3e, 0x76, 0x4a, 0xcd,
	0x7b, 0x03, 0x0f, 0x41, 0xd4, 0x4e, 0xb0, 0x2d, 0x14, 0xb3, 0x25, 0x07, 0xaf, 0xff, 0x86, 0xb2,
	0x1e, 0x78, 0x1d, 0x1f, 0x77, 0xc1, 0x2f, 0x24, 0xd6, 0xb4, 0xf5, 0xac, 0x95, 0xb9, 0xb0, 0x52,
	0x41, 0xb2, 0x70, 0xb6, 0xd2, 0xca, 0x78, 0xe0, 0xfd, 0x2f, 0xbf, 0xd1, 0x3f, 0x68, 0x68, 0x09,
	0x88, 0x1d, 0x44, 0x4c, 0x40, 0xaf, 0x23, 0x37, 0x96, 0x54, 0x1b, 0x7b, 0xab, 0x5d, 0x58, 0xa9,
	0xd3, 0x64, 0xe1, 0x2c, 0x73, 0x3e, 0x59, 0x7d, 0xa9, 0x15, 0x9b, 0x3f, 0x70, 0x9a, 0xea, 0x93,
	0xf9, 0x23, 0xc7, 0x25, 0x32, 0xf2, 0x11, 0x9c, 0x7c, 0xcb, 0x7c, 0x78, 0x4f, 0xde, 0x21, 0xcc,
	0xc2, 0xe6, 0x2f, 0xdd, 0x37, 0x20, 0xfa, 0x4b, 0xba, 0x5d, 0xc8, 0x68, 0x2b, 0x5a, 0x69, 0xa2,
	0xa1, 0x4c, 0x8b, 0x52, 0xd1, 0x80, 0x88, 0xeb, 0xeb, 0x28, 0x17, 0x50, 0x2a, 0x64, 0xbe, 0x8e,
	0xdb, 0x53, 0x47, 0x19, 0x1f, 0x47, 0x36, 0x98, 0x0a, 0xeb, 0x3d, 0x7d, 0x0b, 0x2d, 0x62, 0xc6,
	0xd4, 0x20, 0xe4, 0xd0, 0x72, 0x9b, 0x3f, 0x1b, 0xf3, 0xaf, 0x05, 0x23, 0x76, 0xe5, 0xad, 0x34,
	0x66, 0xac, 0x01, 0x91, 0xec, 0x22, 0xa1, 0x77, 0x39, 0xbe, 0x9b, 0xba, 0x48, 0xe8, 0xcd, 0x9b,
	0x7d, 0x9f, 0x40, 0xb9, 0x36, 0x70, 0xee, 0x52, 0xa2, 0xfc, 0x1a, 0x68, 0x99, 0x4f, 0xcb, 0xb8,
	0xe5, 0xbc, 0xb2, 0x7c, 0xaa, 0x2c, 0xe7, 0xf9, 0xa5, 0xbc, 0xde, 0xd3, 0xab, 0xe8, 0xa7, 0x7e,
	0x47, 0x3a, 0xe0, 0x1d, 0x97, 0x88, 0xdb, 0xba, 0xcf, 0xf5, 0x9b, 0xa1, 0xd7, 0xae, 0x13, 0x19,
	0x5d, 0x22, 0xf8, 0x15, 0xc4, 0x2d, 0xa2, 0xe4, 0x78, 0x0c, 0xf1, 0x37, 0x5a, 0x9a, 0x02, 0x80,
	0xd8, 0x0a, 0xb0, 0x70, 0x33, 0x00, 0x91, 0xd0, 0x6b, 0xd7, 0x88, 0x2d, 0xfb, 0x77, 0x51, 0x56,
	0xce, 0x9e, 0xab, 0xde, 0xd4, 0xcd, 0xbd, 0x72, 0x53, 0xed, 0xb9, 0x41, 0x5a, 0xdb, 0xef, 0x3e,
	0xfe, 0xaa, 0x3d, 0x31, 0xef, 0xf0, 0x47, 0x17, 0x84, 0x75, 0xbb, 0x69, 0x75, 0x6e, 0x95, 0xaf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0xb5, 0xcd, 0x58, 0x21, 0x06, 0x00, 0x00,
}
