// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_kv.proto

package telemetry_kv

/*
Package with obsolete definition of Model Driven Telemetry GPB K/V format message.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TelemetryOLD struct {
	CollectionId           uint64               `protobuf:"varint,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	BasePath               string               `protobuf:"bytes,2,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	SubscriptionIdentifier string               `protobuf:"bytes,3,opt,name=subscription_identifier,json=subscriptionIdentifier,proto3" json:"subscription_identifier,omitempty"`
	ModelVersion           string               `protobuf:"bytes,4,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	CollectionStartTime    uint64               `protobuf:"varint,5,opt,name=collection_start_time,json=collectionStartTime,proto3" json:"collection_start_time,omitempty"`
	MsgTimestamp           uint64               `protobuf:"varint,6,opt,name=msg_timestamp,json=msgTimestamp,proto3" json:"msg_timestamp,omitempty"`
	Fields                 []*TelemetryFieldOLD `protobuf:"bytes,14,rep,name=fields,proto3" json:"fields,omitempty"`
	CollectionEndTime      uint64               `protobuf:"varint,15,opt,name=collection_end_time,json=collectionEndTime,proto3" json:"collection_end_time,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *TelemetryOLD) Reset()         { *m = TelemetryOLD{} }
func (m *TelemetryOLD) String() string { return proto.CompactTextString(m) }
func (*TelemetryOLD) ProtoMessage()    {}
func (*TelemetryOLD) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_kv_b443f9e7e0cc95f1, []int{0}
}
func (m *TelemetryOLD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryOLD.Unmarshal(m, b)
}
func (m *TelemetryOLD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryOLD.Marshal(b, m, deterministic)
}
func (dst *TelemetryOLD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryOLD.Merge(dst, src)
}
func (m *TelemetryOLD) XXX_Size() int {
	return xxx_messageInfo_TelemetryOLD.Size(m)
}
func (m *TelemetryOLD) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryOLD.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryOLD proto.InternalMessageInfo

func (m *TelemetryOLD) GetCollectionId() uint64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *TelemetryOLD) GetBasePath() string {
	if m != nil {
		return m.BasePath
	}
	return ""
}

func (m *TelemetryOLD) GetSubscriptionIdentifier() string {
	if m != nil {
		return m.SubscriptionIdentifier
	}
	return ""
}

func (m *TelemetryOLD) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func (m *TelemetryOLD) GetCollectionStartTime() uint64 {
	if m != nil {
		return m.CollectionStartTime
	}
	return 0
}

func (m *TelemetryOLD) GetMsgTimestamp() uint64 {
	if m != nil {
		return m.MsgTimestamp
	}
	return 0
}

func (m *TelemetryOLD) GetFields() []*TelemetryFieldOLD {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TelemetryOLD) GetCollectionEndTime() uint64 {
	if m != nil {
		return m.CollectionEndTime
	}
	return 0
}

type TelemetryFieldOLD struct {
	Timestamp   uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AugmentData bool   `protobuf:"varint,3,opt,name=augment_data,json=augmentData,proto3" json:"augment_data,omitempty"`
	// Types that are valid to be assigned to ValueByType:
	//	*TelemetryFieldOLD_BytesValue
	//	*TelemetryFieldOLD_StringValue
	//	*TelemetryFieldOLD_BoolValue
	//	*TelemetryFieldOLD_Uint32Value
	//	*TelemetryFieldOLD_Uint64Value
	//	*TelemetryFieldOLD_Sint32Value
	//	*TelemetryFieldOLD_Sint64Value
	//	*TelemetryFieldOLD_DoubleValue
	//	*TelemetryFieldOLD_FloatValue
	ValueByType          isTelemetryFieldOLD_ValueByType `protobuf_oneof:"value_by_type"`
	Fields               []*TelemetryFieldOLD            `protobuf:"bytes,15,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TelemetryFieldOLD) Reset()         { *m = TelemetryFieldOLD{} }
func (m *TelemetryFieldOLD) String() string { return proto.CompactTextString(m) }
func (*TelemetryFieldOLD) ProtoMessage()    {}
func (*TelemetryFieldOLD) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_kv_b443f9e7e0cc95f1, []int{1}
}
func (m *TelemetryFieldOLD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryFieldOLD.Unmarshal(m, b)
}
func (m *TelemetryFieldOLD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryFieldOLD.Marshal(b, m, deterministic)
}
func (dst *TelemetryFieldOLD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryFieldOLD.Merge(dst, src)
}
func (m *TelemetryFieldOLD) XXX_Size() int {
	return xxx_messageInfo_TelemetryFieldOLD.Size(m)
}
func (m *TelemetryFieldOLD) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryFieldOLD.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryFieldOLD proto.InternalMessageInfo

func (m *TelemetryFieldOLD) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TelemetryFieldOLD) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TelemetryFieldOLD) GetAugmentData() bool {
	if m != nil {
		return m.AugmentData
	}
	return false
}

type isTelemetryFieldOLD_ValueByType interface {
	isTelemetryFieldOLD_ValueByType()
}

type TelemetryFieldOLD_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,4,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type TelemetryFieldOLD_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type TelemetryFieldOLD_BoolValue struct {
	BoolValue bool `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type TelemetryFieldOLD_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,7,opt,name=uint32_value,json=uint32Value,proto3,oneof"`
}

type TelemetryFieldOLD_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,8,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type TelemetryFieldOLD_Sint32Value struct {
	Sint32Value int32 `protobuf:"zigzag32,9,opt,name=sint32_value,json=sint32Value,proto3,oneof"`
}

type TelemetryFieldOLD_Sint64Value struct {
	Sint64Value int64 `protobuf:"zigzag64,10,opt,name=sint64_value,json=sint64Value,proto3,oneof"`
}

type TelemetryFieldOLD_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,11,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type TelemetryFieldOLD_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,12,opt,name=float_value,json=floatValue,proto3,oneof"`
}

func (*TelemetryFieldOLD_BytesValue) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_StringValue) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_BoolValue) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_Uint32Value) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_Uint64Value) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_Sint32Value) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_Sint64Value) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_DoubleValue) isTelemetryFieldOLD_ValueByType() {}

func (*TelemetryFieldOLD_FloatValue) isTelemetryFieldOLD_ValueByType() {}

func (m *TelemetryFieldOLD) GetValueByType() isTelemetryFieldOLD_ValueByType {
	if m != nil {
		return m.ValueByType
	}
	return nil
}

func (m *TelemetryFieldOLD) GetBytesValue() []byte {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *TelemetryFieldOLD) GetStringValue() string {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TelemetryFieldOLD) GetBoolValue() bool {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TelemetryFieldOLD) GetUint32Value() uint32 {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (m *TelemetryFieldOLD) GetUint64Value() uint64 {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *TelemetryFieldOLD) GetSint32Value() int32 {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_Sint32Value); ok {
		return x.Sint32Value
	}
	return 0
}

func (m *TelemetryFieldOLD) GetSint64Value() int64 {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_Sint64Value); ok {
		return x.Sint64Value
	}
	return 0
}

func (m *TelemetryFieldOLD) GetDoubleValue() float64 {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TelemetryFieldOLD) GetFloatValue() float32 {
	if x, ok := m.GetValueByType().(*TelemetryFieldOLD_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *TelemetryFieldOLD) GetFields() []*TelemetryFieldOLD {
	if m != nil {
		return m.Fields
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TelemetryFieldOLD) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TelemetryFieldOLD_OneofMarshaler, _TelemetryFieldOLD_OneofUnmarshaler, _TelemetryFieldOLD_OneofSizer, []interface{}{
		(*TelemetryFieldOLD_BytesValue)(nil),
		(*TelemetryFieldOLD_StringValue)(nil),
		(*TelemetryFieldOLD_BoolValue)(nil),
		(*TelemetryFieldOLD_Uint32Value)(nil),
		(*TelemetryFieldOLD_Uint64Value)(nil),
		(*TelemetryFieldOLD_Sint32Value)(nil),
		(*TelemetryFieldOLD_Sint64Value)(nil),
		(*TelemetryFieldOLD_DoubleValue)(nil),
		(*TelemetryFieldOLD_FloatValue)(nil),
	}
}

func _TelemetryFieldOLD_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TelemetryFieldOLD)
	// value_by_type
	switch x := m.ValueByType.(type) {
	case *TelemetryFieldOLD_BytesValue:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesValue)
	case *TelemetryFieldOLD_StringValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *TelemetryFieldOLD_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *TelemetryFieldOLD_Uint32Value:
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint32Value))
	case *TelemetryFieldOLD_Uint64Value:
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *TelemetryFieldOLD_Sint32Value:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeZigzag32(uint64(x.Sint32Value))
	case *TelemetryFieldOLD_Sint64Value:
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeZigzag64(uint64(x.Sint64Value))
	case *TelemetryFieldOLD_DoubleValue:
		b.EncodeVarint(11<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *TelemetryFieldOLD_FloatValue:
		b.EncodeVarint(12<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case nil:
	default:
		return fmt.Errorf("TelemetryFieldOLD.ValueByType has unexpected type %T", x)
	}
	return nil
}

func _TelemetryFieldOLD_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TelemetryFieldOLD)
	switch tag {
	case 4: // value_by_type.bytes_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ValueByType = &TelemetryFieldOLD_BytesValue{x}
		return true, err
	case 5: // value_by_type.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueByType = &TelemetryFieldOLD_StringValue{x}
		return true, err
	case 6: // value_by_type.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueByType = &TelemetryFieldOLD_BoolValue{x != 0}
		return true, err
	case 7: // value_by_type.uint32_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueByType = &TelemetryFieldOLD_Uint32Value{uint32(x)}
		return true, err
	case 8: // value_by_type.uint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueByType = &TelemetryFieldOLD_Uint64Value{x}
		return true, err
	case 9: // value_by_type.sint32_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.ValueByType = &TelemetryFieldOLD_Sint32Value{int32(x)}
		return true, err
	case 10: // value_by_type.sint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag64()
		m.ValueByType = &TelemetryFieldOLD_Sint64Value{int64(x)}
		return true, err
	case 11: // value_by_type.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ValueByType = &TelemetryFieldOLD_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 12: // value_by_type.float_value
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.ValueByType = &TelemetryFieldOLD_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _TelemetryFieldOLD_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TelemetryFieldOLD)
	// value_by_type
	switch x := m.ValueByType.(type) {
	case *TelemetryFieldOLD_BytesValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BytesValue)))
		n += len(x.BytesValue)
	case *TelemetryFieldOLD_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *TelemetryFieldOLD_BoolValue:
		n += 1 // tag and wire
		n += 1
	case *TelemetryFieldOLD_Uint32Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Uint32Value))
	case *TelemetryFieldOLD_Uint64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *TelemetryFieldOLD_Sint32Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64((uint32(x.Sint32Value) << 1) ^ uint32((int32(x.Sint32Value) >> 31))))
	case *TelemetryFieldOLD_Sint64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(uint64(x.Sint64Value<<1) ^ uint64((int64(x.Sint64Value) >> 63))))
	case *TelemetryFieldOLD_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *TelemetryFieldOLD_FloatValue:
		n += 1 // tag and wire
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*TelemetryOLD)(nil), "telemetry_kv.TelemetryOLD")
	proto.RegisterType((*TelemetryFieldOLD)(nil), "telemetry_kv.TelemetryFieldOLD")
}

func init() { proto.RegisterFile("telemetry_kv.proto", fileDescriptor_telemetry_kv_b443f9e7e0cc95f1) }

var fileDescriptor_telemetry_kv_b443f9e7e0cc95f1 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x8d, 0x5a, 0x6b, 0x5e, 0x62, 0xc5, 0x29, 0x6d, 0x03, 0x2d, 0x6c, 0xdc, 0xbd, 0xe4,
	0xe4, 0xc1, 0x5d, 0x76, 0xef, 0xc5, 0x16, 0x17, 0x16, 0xb6, 0xa4, 0xb2, 0xd7, 0x61, 0x62, 0x46,
	0x77, 0x68, 0x92, 0x91, 0xcc, 0x53, 0xf0, 0xbf, 0xee, 0xad, 0xd7, 0x32, 0x3f, 0x34, 0x91, 0x9e,
	0x7a, 0x93, 0xcf, 0xfb, 0xf0, 0x7d, 0x2f, 0x5f, 0x46, 0x20, 0xc8, 0x0b, 0x5e, 0x72, 0xac, 0x8f,
	0xf4, 0xd7, 0x61, 0xb6, 0xab, 0x25, 0x4a, 0x12, 0xb6, 0xd9, 0xf5, 0xef, 0x2e, 0x84, 0xab, 0x13,
	0x78, 0x7e, 0x5a, 0x90, 0x1b, 0x18, 0xad, 0x65, 0x51, 0xf0, 0x35, 0x0a, 0x59, 0x51, 0x91, 0x47,
	0x5e, 0xec, 0x25, 0xfd, 0x34, 0x6c, 0xe0, 0x63, 0x4e, 0x3e, 0x83, 0x9f, 0x31, 0xc5, 0xe9, 0x8e,
	0xe1, 0x6b, 0xd4, 0x8d, 0xbd, 0xc4, 0x4f, 0x87, 0x1a, 0xfc, 0x60, 0xf8, 0x4a, 0x1e, 0xe0, 0x93,
	0xda, 0x67, 0x6a, 0x5d, 0x8b, 0x9d, 0xcb, 0xe0, 0x15, 0x8a, 0x8d, 0xe0, 0x75, 0xd4, 0x33, 0xea,
	0xc7, 0xf6, 0xf8, 0xf1, 0x3c, 0xd5, 0xab, 0x4b, 0x99, 0xf3, 0x82, 0x1e, 0x78, 0xad, 0x84, 0xac,
	0xa2, 0xbe, 0xd1, 0x43, 0x03, 0x5f, 0x2c, 0x23, 0x73, 0xf8, 0xd0, 0xba, 0x4f, 0x21, 0xab, 0x91,
	0xa2, 0x28, 0x79, 0xf4, 0xc6, 0xdc, 0xf9, 0xbe, 0x19, 0xfe, 0xd4, 0xb3, 0x95, 0x28, 0xb9, 0x09,
	0x56, 0x5b, 0xa3, 0x29, 0x64, 0xe5, 0x2e, 0x1a, 0xd8, 0x6f, 0x2a, 0xd5, 0x76, 0x75, 0x62, 0xe4,
	0x01, 0x06, 0x1b, 0xc1, 0x8b, 0x5c, 0x45, 0xef, 0xe2, 0x5e, 0x12, 0xcc, 0xaf, 0x66, 0x17, 0xe5,
	0x9d, 0x4b, 0xfa, 0xae, 0xa5, 0xe7, 0xa7, 0x45, 0xea, 0x74, 0x32, 0x83, 0xd6, 0x52, 0xca, 0xab,
	0xdc, 0xde, 0x33, 0x36, 0x3b, 0x26, 0xcd, 0xe8, 0x5b, 0x95, 0xeb, 0x6d, 0xd7, 0x7f, 0x7a, 0x30,
	0xf9, 0x27, 0x8d, 0x7c, 0x01, 0xbf, 0xb9, 0xcf, 0x76, 0xde, 0x00, 0x42, 0xa0, 0x5f, 0xb1, 0x92,
	0xbb, 0xae, 0xcd, 0x6f, 0x32, 0x85, 0x90, 0xed, 0xb7, 0x25, 0xaf, 0x90, 0xe6, 0x0c, 0x99, 0x29,
	0x77, 0x98, 0x06, 0x8e, 0x2d, 0x18, 0x32, 0x32, 0x85, 0x20, 0x3b, 0x22, 0x57, 0xf4, 0xc0, 0x8a,
	0x3d, 0x37, 0x7d, 0x86, 0xcb, 0x4e, 0x0a, 0x06, 0xbe, 0x68, 0x46, 0x6e, 0x20, 0x54, 0x58, 0x8b,
	0x6a, 0xeb, 0x1c, 0x5d, 0xa3, 0xbf, 0xec, 0xa4, 0x81, 0xa5, 0x56, 0xba, 0x02, 0xc8, 0xa4, 0x2c,
	0x9c, 0xa2, 0xdb, 0x1b, 0x2e, 0x3b, 0xa9, 0xaf, 0xd9, 0x39, 0x65, 0x2f, 0x2a, 0xbc, 0x9d, 0x3b,
	0xe5, 0x6d, 0xec, 0x25, 0x23, 0x9d, 0x62, 0xe9, 0x85, 0x74, 0x7f, 0xe7, 0xa4, 0xa1, 0xfe, 0xca,
	0x93, 0x74, 0x7f, 0xd7, 0xdc, 0xd3, 0x4e, 0xf2, 0x63, 0x2f, 0x99, 0x98, 0x7b, 0x2e, 0x93, 0x54,
	0x3b, 0x09, 0x62, 0x2f, 0x21, 0x27, 0xa9, 0x95, 0x94, 0xcb, 0x7d, 0x56, 0x70, 0x27, 0x05, 0xb1,
	0x97, 0x78, 0x5a, 0xb2, 0xd4, 0x4a, 0x53, 0x08, 0x36, 0x85, 0x64, 0xe8, 0x9c, 0x30, 0xf6, 0x92,
	0xae, 0x6e, 0xc8, 0x40, 0xab, 0x34, 0x0f, 0x63, 0xfc, 0x5f, 0x0f, 0xe3, 0xeb, 0x18, 0x46, 0x26,
	0x95, 0x66, 0x47, 0x8a, 0xc7, 0x1d, 0xcf, 0x06, 0xe6, 0x1f, 0x78, 0xfb, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0x05, 0x1f, 0x56, 0x97, 0x03, 0x00, 0x00,
}
