// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/join.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
import time "time"

import bytes "bytes"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type JoinRequest struct {
	RawPayload           []byte   `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	Payload              *Message `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	EndDeviceIdentifiers `protobuf:"bytes,3,opt,name=end_device_ids,json=endDeviceIds,proto3,embedded=end_device_ids" json:"end_device_ids"`
	SelectedMACVersion   MACVersion                                         `protobuf:"varint,4,opt,name=selected_mac_version,json=selectedMacVersion,proto3,enum=ttn.lorawan.v3.MACVersion" json:"selected_mac_version,omitempty"`
	NetID                go_thethings_network_lorawan_stack_pkg_types.NetID `protobuf:"bytes,5,opt,name=net_id,json=netId,proto3,customtype=go.thethings.network/lorawan-stack/pkg/types.NetID" json:"net_id"`
	DownlinkSettings     DLSettings                                         `protobuf:"bytes,6,opt,name=downlink_settings,json=downlinkSettings,proto3" json:"downlink_settings"`
	RxDelay              RxDelay                                            `protobuf:"varint,7,opt,name=rx_delay,json=rxDelay,proto3,enum=ttn.lorawan.v3.RxDelay" json:"rx_delay,omitempty"`
	// Optional CFList.
	CFList               *CFList  `protobuf:"bytes,8,opt,name=cf_list,json=cfList,proto3" json:"cf_list,omitempty"`
	CorrelationIDs       []string `protobuf:"bytes,10,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()      { *m = JoinRequest{} }
func (*JoinRequest) ProtoMessage() {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_join_0f271ae05d562c67, []int{0}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(dst, src)
}
func (m *JoinRequest) XXX_Size() int {
	return m.Size()
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinRequest) GetPayload() *Message {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *JoinRequest) GetSelectedMACVersion() MACVersion {
	if m != nil {
		return m.SelectedMACVersion
	}
	return MAC_UNKNOWN
}

func (m *JoinRequest) GetDownlinkSettings() DLSettings {
	if m != nil {
		return m.DownlinkSettings
	}
	return DLSettings{}
}

func (m *JoinRequest) GetRxDelay() RxDelay {
	if m != nil {
		return m.RxDelay
	}
	return RX_DELAY_0
}

func (m *JoinRequest) GetCFList() *CFList {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *JoinRequest) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

type JoinResponse struct {
	RawPayload           []byte `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	SessionKeys          `protobuf:"bytes,2,opt,name=session_keys,json=sessionKeys,proto3,embedded=session_keys" json:"session_keys"`
	Lifetime             time.Duration `protobuf:"bytes,3,opt,name=lifetime,proto3,stdduration" json:"lifetime"`
	CorrelationIDs       []string      `protobuf:"bytes,4,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JoinResponse) Reset()      { *m = JoinResponse{} }
func (*JoinResponse) ProtoMessage() {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_join_0f271ae05d562c67, []int{1}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(dst, src)
}
func (m *JoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinResponse) GetLifetime() time.Duration {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *JoinResponse) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	golang_proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
	golang_proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
}
func (this *JoinRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinRequest)
	if !ok {
		that2, ok := that.(JoinRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	if !this.EndDeviceIdentifiers.Equal(&that1.EndDeviceIdentifiers) {
		return false
	}
	if this.SelectedMACVersion != that1.SelectedMACVersion {
		return false
	}
	if !this.NetID.Equal(that1.NetID) {
		return false
	}
	if !this.DownlinkSettings.Equal(&that1.DownlinkSettings) {
		return false
	}
	if this.RxDelay != that1.RxDelay {
		return false
	}
	if !this.CFList.Equal(that1.CFList) {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	return true
}
func (this *JoinResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinResponse)
	if !ok {
		that2, ok := that.(JoinResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.SessionKeys.Equal(&that1.SessionKeys) {
		return false
	}
	if this.Lifetime != that1.Lifetime {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	return true
}
func (m *JoinRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RawPayload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i += copy(dAtA[i:], m.RawPayload)
	}
	if m.Payload != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.Payload.Size()))
		n1, err := m.Payload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.EndDeviceIdentifiers.Size()))
	n2, err := m.EndDeviceIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.SelectedMACVersion != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.SelectedMACVersion))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.NetID.Size()))
	n3, err := m.NetID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.DownlinkSettings.Size()))
	n4, err := m.DownlinkSettings.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.RxDelay != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.RxDelay))
	}
	if m.CFList != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.CFList.Size()))
		n5, err := m.CFList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			dAtA[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *JoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RawPayload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i += copy(dAtA[i:], m.RawPayload)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.SessionKeys.Size()))
	n6, err := m.SessionKeys.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x1a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime)))
	n7, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Lifetime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintJoin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedJoinResponse(r randyJoin, easy bool) *JoinResponse {
	this := &JoinResponse{}
	v1 := r.Intn(100)
	this.RawPayload = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.RawPayload[i] = byte(r.Intn(256))
	}
	v2 := NewPopulatedSessionKeys(r, easy)
	this.SessionKeys = *v2
	v3 := github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.Lifetime = *v3
	v4 := r.Intn(10)
	this.CorrelationIDs = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.CorrelationIDs[i] = randStringJoin(r)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyJoin interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneJoin(r randyJoin) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringJoin(r randyJoin) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneJoin(r)
	}
	return string(tmps)
}
func randUnrecognizedJoin(r randyJoin, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldJoin(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldJoin(dAtA []byte, r randyJoin, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateJoin(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *JoinRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.EndDeviceIdentifiers.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.SelectedMACVersion != 0 {
		n += 1 + sovJoin(uint64(m.SelectedMACVersion))
	}
	l = m.NetID.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = m.DownlinkSettings.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.RxDelay != 0 {
		n += 1 + sovJoin(uint64(m.RxDelay))
	}
	if m.CFList != nil {
		l = m.CFList.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	return n
}

func (m *JoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.SessionKeys.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime)
	n += 1 + l + sovJoin(uint64(l))
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	return n
}

func sovJoin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJoin(x uint64) (n int) {
	return sovJoin((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *JoinRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinRequest{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "Message", "Message", 1) + `,`,
		`EndDeviceIdentifiers:` + strings.Replace(strings.Replace(this.EndDeviceIdentifiers.String(), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1), `&`, ``, 1) + `,`,
		`SelectedMACVersion:` + fmt.Sprintf("%v", this.SelectedMACVersion) + `,`,
		`NetID:` + fmt.Sprintf("%v", this.NetID) + `,`,
		`DownlinkSettings:` + strings.Replace(strings.Replace(this.DownlinkSettings.String(), "DLSettings", "DLSettings", 1), `&`, ``, 1) + `,`,
		`RxDelay:` + fmt.Sprintf("%v", this.RxDelay) + `,`,
		`CFList:` + strings.Replace(fmt.Sprintf("%v", this.CFList), "CFList", "CFList", 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JoinResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinResponse{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`SessionKeys:` + strings.Replace(strings.Replace(this.SessionKeys.String(), "SessionKeys", "SessionKeys", 1), `&`, ``, 1) + `,`,
		`Lifetime:` + strings.Replace(strings.Replace(this.Lifetime.String(), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJoin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *JoinRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &Message{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedMACVersion", wireType)
			}
			m.SelectedMACVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SelectedMACVersion |= (MACVersion(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownlinkSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DownlinkSettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxDelay", wireType)
			}
			m.RxDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RxDelay |= (RxDelay(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CFList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CFList == nil {
				m.CFList = &CFList{}
			}
			if err := m.CFList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationIDs = append(m.CorrelationIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SessionKeys.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lifetime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Lifetime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationIDs = append(m.CorrelationIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipJoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJoin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthJoin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJoin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipJoin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthJoin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJoin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_join_0f271ae05d562c67) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_join_0f271ae05d562c67)
}

var fileDescriptor_join_0f271ae05d562c67 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3d, 0x6c, 0x23, 0x45,
	0x18, 0x9d, 0xe1, 0xfc, 0x77, 0x63, 0xcb, 0x1c, 0x2b, 0x74, 0x18, 0x83, 0x66, 0xad, 0x83, 0xc2,
	0x05, 0xb7, 0x16, 0xbe, 0x8e, 0x14, 0x08, 0xdb, 0x20, 0x1c, 0x12, 0x84, 0x36, 0x88, 0x22, 0xcd,
	0x6a, 0xbd, 0x3b, 0xde, 0x0c, 0xde, 0xcc, 0x98, 0x9d, 0x71, 0x1c, 0x77, 0x11, 0x55, 0x4a, 0xca,
	0x94, 0x11, 0x55, 0xca, 0x94, 0x29, 0x53, 0xa6, 0x4c, 0x19, 0x51, 0x98, 0x78, 0xb6, 0x49, 0x99,
	0x8e, 0x94, 0x68, 0xff, 0x70, 0x12, 0x47, 0x22, 0x95, 0xbf, 0xfd, 0xbe, 0xf7, 0xbe, 0x79, 0x7e,
	0xf3, 0x06, 0x7d, 0xea, 0xf3, 0xc0, 0x9e, 0xda, 0xec, 0xad, 0x90, 0xb6, 0x33, 0x6a, 0xd9, 0x63,
	0xda, 0xfa, 0x95, 0x53, 0x66, 0x8c, 0x03, 0x2e, 0xb9, 0x56, 0x95, 0x92, 0x19, 0x29, 0xc2, 0xd8,
	0x7b, 0x57, 0x7f, 0xeb, 0x51, 0xb9, 0x33, 0x19, 0x18, 0x0e, 0xdf, 0x6d, 0x79, 0xdc, 0xe3, 0xad,
	0x18, 0x36, 0x98, 0x0c, 0xe3, 0xaf, 0xf8, 0x23, 0xae, 0x12, 0x7a, 0x1d, 0x7b, 0x9c, 0x7b, 0x3e,
	0x59, 0xa2, 0xdc, 0x49, 0x60, 0x4b, 0xca, 0xd3, 0xf5, 0xf5, 0xcf, 0x56, 0x0f, 0xa7, 0x2e, 0x61,
	0x92, 0x0e, 0x29, 0x09, 0x44, 0x0a, 0x7a, 0x42, 0xe1, 0x88, 0xcc, 0xb2, 0xa9, 0xbe, 0x3a, 0xcd,
	0xf4, 0xc6, 0x80, 0x37, 0xbf, 0xe7, 0x51, 0x79, 0x9d, 0x53, 0x66, 0x92, 0xdf, 0x26, 0x44, 0x48,
	0x4d, 0x47, 0xe5, 0xc0, 0x9e, 0x5a, 0x63, 0x7b, 0xe6, 0x73, 0xdb, 0xad, 0xc1, 0x06, 0x6c, 0x56,
	0x4c, 0x14, 0xd8, 0xd3, 0x9f, 0x92, 0x8e, 0xf6, 0x25, 0x2a, 0x66, 0xc3, 0xf7, 0x1a, 0xb0, 0x59,
	0x6e, 0x7f, 0x64, 0x3c, 0x74, 0xc1, 0xd8, 0x24, 0x42, 0xd8, 0x1e, 0x31, 0x33, 0x9c, 0xf6, 0x33,
	0xaa, 0x12, 0xe6, 0x5a, 0x2e, 0xd9, 0xa3, 0x0e, 0xb1, 0xa8, 0x2b, 0x6a, 0x2f, 0x62, 0xe6, 0xe7,
	0x8f, 0x99, 0xdf, 0x32, 0xb7, 0x17, 0x83, 0xfa, 0xcb, 0xbf, 0xd9, 0x29, 0x5d, 0xcc, 0x75, 0x70,
	0x39, 0xd7, 0xa1, 0x59, 0x21, 0xcb, 0xb9, 0xd0, 0x5c, 0xf4, 0xa1, 0x20, 0x3e, 0x71, 0x24, 0x71,
	0xad, 0x5d, 0xdb, 0xb1, 0xf6, 0x48, 0x20, 0x28, 0x67, 0xb5, 0x5c, 0x03, 0x36, 0xab, 0xed, 0xfa,
	0x8a, 0xaa, 0x6f, 0xba, 0xbf, 0x24, 0x88, 0xce, 0x6b, 0x35, 0xd7, 0xb5, 0xad, 0x94, 0xbb, 0xec,
	0x9b, 0x5a, 0xb6, 0x6f, 0xd3, 0x76, 0xd2, 0x9e, 0xb6, 0x8d, 0x0a, 0x8c, 0x48, 0x8b, 0xba, 0xb5,
	0x7c, 0x64, 0x45, 0xa7, 0x1b, 0xa9, 0xf9, 0x6b, 0xae, 0xb7, 0x3d, 0x6e, 0xc8, 0x1d, 0x22, 0x77,
	0x28, 0xf3, 0x84, 0xc1, 0x88, 0x9c, 0xf2, 0x60, 0xd4, 0x7a, 0xe8, 0xf6, 0x78, 0xe4, 0xb5, 0xe4,
	0x6c, 0x4c, 0x84, 0xf1, 0x23, 0x91, 0xfd, 0x9e, 0x9a, 0xeb, 0xf9, 0xb8, 0x30, 0xf3, 0x8c, 0xc8,
	0xbe, 0xab, 0x6d, 0xa2, 0x0f, 0x5c, 0x3e, 0x65, 0x3e, 0x65, 0x23, 0x4b, 0x10, 0x29, 0xa3, 0x55,
	0xb5, 0x42, 0x6c, 0xcd, 0x8a, 0xfc, 0xde, 0xc6, 0x56, 0x8a, 0xe8, 0xe4, 0x22, 0x09, 0xe6, 0xab,
	0x8c, 0x9a, 0xf5, 0xb5, 0x36, 0x2a, 0x05, 0xfb, 0x96, 0x4b, 0x7c, 0x7b, 0x56, 0x2b, 0xc6, 0x26,
	0xac, 0x5c, 0x8d, 0xb9, 0xdf, 0x8b, 0xc6, 0x66, 0x31, 0x48, 0x0a, 0x6d, 0x0d, 0x15, 0x9d, 0xa1,
	0xe5, 0x53, 0x21, 0x6b, 0xa5, 0xf8, 0xe0, 0xd7, 0x8f, 0x29, 0xdd, 0xef, 0x36, 0xa8, 0x90, 0x1d,
	0xa4, 0xe6, 0x7a, 0x21, 0xa9, 0xcd, 0x82, 0x33, 0x8c, 0x7e, 0xb5, 0x35, 0xf4, 0xbe, 0xc3, 0x83,
	0x80, 0xf8, 0x71, 0x68, 0xe3, 0x8b, 0x45, 0x8d, 0x17, 0xcd, 0x97, 0x1d, 0x4d, 0xcd, 0xf5, 0x6a,
	0x77, 0x39, 0xea, 0xf7, 0x84, 0x59, 0xbd, 0x07, 0xed, 0xbb, 0xe2, 0xab, 0xdc, 0xd9, 0xb1, 0x0e,
	0xd6, 0x73, 0xa5, 0x97, 0xaf, 0xd0, 0x9b, 0x7f, 0x20, 0xaa, 0x24, 0x21, 0x14, 0x63, 0xce, 0x04,
	0xf9, 0xff, 0x14, 0x7e, 0x8f, 0x2a, 0x82, 0x88, 0xe8, 0x86, 0xac, 0x28, 0xed, 0x69, 0x14, 0x3f,
	0x79, 0x2c, 0x7e, 0x2b, 0xc1, 0xfc, 0x40, 0x66, 0xf7, 0x73, 0x54, 0x16, 0xcb, 0xb6, 0xf6, 0x35,
	0x2a, 0xf9, 0x74, 0x48, 0x24, 0xdd, 0x25, 0x69, 0x2c, 0x3f, 0x36, 0x92, 0x77, 0x69, 0x64, 0xef,
	0xd2, 0xe8, 0xa5, 0xef, 0x32, 0xd9, 0x71, 0xf4, 0xb7, 0x0e, 0xcd, 0xff, 0x48, 0x4f, 0xb9, 0x90,
	0x7b, 0xae, 0x0b, 0x9d, 0x3f, 0xe1, 0xc5, 0x02, 0xc3, 0xcb, 0x05, 0x86, 0x57, 0x0b, 0x0c, 0xae,
	0x17, 0x18, 0xdc, 0x2c, 0x30, 0xb8, 0x5d, 0x60, 0x70, 0xb7, 0xc0, 0xf0, 0x40, 0x61, 0x78, 0xa8,
	0x30, 0x38, 0x51, 0x18, 0x9e, 0x2a, 0x0c, 0xce, 0x14, 0x06, 0xe7, 0x0a, 0x83, 0x0b, 0x85, 0xe1,
	0xa5, 0xc2, 0xf0, 0x4a, 0x61, 0x70, 0xad, 0x30, 0xbc, 0x51, 0x18, 0xdc, 0x2a, 0x0c, 0xef, 0x14,
	0x06, 0x07, 0x21, 0x06, 0x87, 0x21, 0x86, 0x7f, 0x84, 0x18, 0x1c, 0x85, 0x18, 0x1e, 0x87, 0x18,
	0x9c, 0x84, 0x18, 0x9c, 0x86, 0x18, 0x9e, 0x85, 0x18, 0x9e, 0x87, 0x18, 0x6e, 0x7f, 0xf1, 0xdc,
	0xf4, 0x4a, 0x36, 0x1e, 0x0c, 0x0a, 0xb1, 0x11, 0xef, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfa,
	0x84, 0xd1, 0xa5, 0x0d, 0x05, 0x00, 0x00,
}
