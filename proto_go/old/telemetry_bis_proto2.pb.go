// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_bis_proto2.proto

package telemetry_proto2

/*
Telemetry message (proto2). See telemetry.proto for documentation
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

//
// Common message used as a header to both compact and self-describing
// telemetry messages.
type Telemetry struct {
	// oneof node_id_str, node_id_uuid
	NodeIdStr *string `protobuf:"bytes,1,opt,name=node_id_str,json=nodeIdStr" json:"node_id_str,omitempty"`
	// oneof subscription_id_str, subscription_id
	SubscriptionIdStr *string `protobuf:"bytes,3,opt,name=subscription_id_str,json=subscriptionIdStr" json:"subscription_id_str,omitempty"`
	// optional string         sensor_path               = 5;    // not produced
	EncodingPath *string `protobuf:"bytes,6,opt,name=encoding_path,json=encodingPath" json:"encoding_path,omitempty"`
	// optional string         model_version             = 7;    // not produced
	CollectionId        *uint64 `protobuf:"varint,8,opt,name=collection_id,json=collectionId" json:"collection_id,omitempty"`
	CollectionStartTime *uint64 `protobuf:"varint,9,opt,name=collection_start_time,json=collectionStartTime" json:"collection_start_time,omitempty"`
	MsgTimestamp        *uint64 `protobuf:"varint,10,opt,name=msg_timestamp,json=msgTimestamp" json:"msg_timestamp,omitempty"`
	// Either data_gpbkv or data_gpb is encoded
	DataGpbkv            []*TelemetryField  `protobuf:"bytes,11,rep,name=data_gpbkv,json=dataGpbkv" json:"data_gpbkv,omitempty"`
	DataGpb              *TelemetryGPBTable `protobuf:"bytes,12,opt,name=data_gpb,json=dataGpb" json:"data_gpb,omitempty"`
	CollectionEndTime    *uint64            `protobuf:"varint,13,opt,name=collection_end_time,json=collectionEndTime" json:"collection_end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Telemetry) Reset()         { *m = Telemetry{} }
func (m *Telemetry) String() string { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()    {}
func (*Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_bis_proto2_be78366a61b59189, []int{0}
}
func (m *Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry.Unmarshal(m, b)
}
func (m *Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry.Marshal(b, m, deterministic)
}
func (dst *Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry.Merge(dst, src)
}
func (m *Telemetry) XXX_Size() int {
	return xxx_messageInfo_Telemetry.Size(m)
}
func (m *Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry proto.InternalMessageInfo

func (m *Telemetry) GetNodeIdStr() string {
	if m != nil && m.NodeIdStr != nil {
		return *m.NodeIdStr
	}
	return ""
}

func (m *Telemetry) GetSubscriptionIdStr() string {
	if m != nil && m.SubscriptionIdStr != nil {
		return *m.SubscriptionIdStr
	}
	return ""
}

func (m *Telemetry) GetEncodingPath() string {
	if m != nil && m.EncodingPath != nil {
		return *m.EncodingPath
	}
	return ""
}

func (m *Telemetry) GetCollectionId() uint64 {
	if m != nil && m.CollectionId != nil {
		return *m.CollectionId
	}
	return 0
}

func (m *Telemetry) GetCollectionStartTime() uint64 {
	if m != nil && m.CollectionStartTime != nil {
		return *m.CollectionStartTime
	}
	return 0
}

func (m *Telemetry) GetMsgTimestamp() uint64 {
	if m != nil && m.MsgTimestamp != nil {
		return *m.MsgTimestamp
	}
	return 0
}

func (m *Telemetry) GetDataGpbkv() []*TelemetryField {
	if m != nil {
		return m.DataGpbkv
	}
	return nil
}

func (m *Telemetry) GetDataGpb() *TelemetryGPBTable {
	if m != nil {
		return m.DataGpb
	}
	return nil
}

func (m *Telemetry) GetCollectionEndTime() uint64 {
	if m != nil && m.CollectionEndTime != nil {
		return *m.CollectionEndTime
	}
	return 0
}

//
// Messages used to export content in GPB K/V form.
//
// The set of messages in this .proto are sufficient to decode all
// telemetry messages.
type TelemetryField struct {
	Timestamp *uint64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Name      *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	// oneof *_value below, encoded based on the native type of each value
	BytesValue           []byte            `protobuf:"bytes,4,opt,name=bytes_value,json=bytesValue" json:"bytes_value,omitempty"`
	StringValue          *string           `protobuf:"bytes,5,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	BoolValue            *bool             `protobuf:"varint,6,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	Uint32Value          *uint32           `protobuf:"varint,7,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	Uint64Value          *uint64           `protobuf:"varint,8,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	Sint32Value          *int32            `protobuf:"zigzag32,9,opt,name=sint32_value,json=sint32Value" json:"sint32_value,omitempty"`
	Sint64Value          *int64            `protobuf:"zigzag64,10,opt,name=sint64_value,json=sint64Value" json:"sint64_value,omitempty"`
	DoubleValue          *float64          `protobuf:"fixed64,11,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	FloatValue           *float32          `protobuf:"fixed32,12,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	Fields               []*TelemetryField `protobuf:"bytes,15,rep,name=fields" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TelemetryField) Reset()         { *m = TelemetryField{} }
func (m *TelemetryField) String() string { return proto.CompactTextString(m) }
func (*TelemetryField) ProtoMessage()    {}
func (*TelemetryField) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_bis_proto2_be78366a61b59189, []int{1}
}
func (m *TelemetryField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryField.Unmarshal(m, b)
}
func (m *TelemetryField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryField.Marshal(b, m, deterministic)
}
func (dst *TelemetryField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryField.Merge(dst, src)
}
func (m *TelemetryField) XXX_Size() int {
	return xxx_messageInfo_TelemetryField.Size(m)
}
func (m *TelemetryField) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryField.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryField proto.InternalMessageInfo

func (m *TelemetryField) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TelemetryField) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TelemetryField) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func (m *TelemetryField) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *TelemetryField) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *TelemetryField) GetUint32Value() uint32 {
	if m != nil && m.Uint32Value != nil {
		return *m.Uint32Value
	}
	return 0
}

func (m *TelemetryField) GetUint64Value() uint64 {
	if m != nil && m.Uint64Value != nil {
		return *m.Uint64Value
	}
	return 0
}

func (m *TelemetryField) GetSint32Value() int32 {
	if m != nil && m.Sint32Value != nil {
		return *m.Sint32Value
	}
	return 0
}

func (m *TelemetryField) GetSint64Value() int64 {
	if m != nil && m.Sint64Value != nil {
		return *m.Sint64Value
	}
	return 0
}

func (m *TelemetryField) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *TelemetryField) GetFloatValue() float32 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *TelemetryField) GetFields() []*TelemetryField {
	if m != nil {
		return m.Fields
	}
	return nil
}

//
// Messages used to export content in compact GPB form
//
// Per encoding-path .proto files are required to decode keys/content
// pairs below.
type TelemetryGPBTable struct {
	Row                  []*TelemetryRowGPB `protobuf:"bytes,1,rep,name=row" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TelemetryGPBTable) Reset()         { *m = TelemetryGPBTable{} }
func (m *TelemetryGPBTable) String() string { return proto.CompactTextString(m) }
func (*TelemetryGPBTable) ProtoMessage()    {}
func (*TelemetryGPBTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_bis_proto2_be78366a61b59189, []int{2}
}
func (m *TelemetryGPBTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryGPBTable.Unmarshal(m, b)
}
func (m *TelemetryGPBTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryGPBTable.Marshal(b, m, deterministic)
}
func (dst *TelemetryGPBTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryGPBTable.Merge(dst, src)
}
func (m *TelemetryGPBTable) XXX_Size() int {
	return xxx_messageInfo_TelemetryGPBTable.Size(m)
}
func (m *TelemetryGPBTable) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryGPBTable.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryGPBTable proto.InternalMessageInfo

func (m *TelemetryGPBTable) GetRow() []*TelemetryRowGPB {
	if m != nil {
		return m.Row
	}
	return nil
}

type TelemetryRowGPB struct {
	Timestamp            *uint64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Keys                 []byte   `protobuf:"bytes,10,req,name=keys" json:"keys,omitempty"`
	Content              []byte   `protobuf:"bytes,11,req,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryRowGPB) Reset()         { *m = TelemetryRowGPB{} }
func (m *TelemetryRowGPB) String() string { return proto.CompactTextString(m) }
func (*TelemetryRowGPB) ProtoMessage()    {}
func (*TelemetryRowGPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_telemetry_bis_proto2_be78366a61b59189, []int{3}
}
func (m *TelemetryRowGPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryRowGPB.Unmarshal(m, b)
}
func (m *TelemetryRowGPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryRowGPB.Marshal(b, m, deterministic)
}
func (dst *TelemetryRowGPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryRowGPB.Merge(dst, src)
}
func (m *TelemetryRowGPB) XXX_Size() int {
	return xxx_messageInfo_TelemetryRowGPB.Size(m)
}
func (m *TelemetryRowGPB) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryRowGPB.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryRowGPB proto.InternalMessageInfo

func (m *TelemetryRowGPB) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TelemetryRowGPB) GetKeys() []byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TelemetryRowGPB) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Telemetry)(nil), "telemetry_proto2.Telemetry")
	proto.RegisterType((*TelemetryField)(nil), "telemetry_proto2.TelemetryField")
	proto.RegisterType((*TelemetryGPBTable)(nil), "telemetry_proto2.TelemetryGPBTable")
	proto.RegisterType((*TelemetryRowGPB)(nil), "telemetry_proto2.TelemetryRowGPB")
}

func init() {
	proto.RegisterFile("telemetry_bis_proto2.proto", fileDescriptor_telemetry_bis_proto2_be78366a61b59189)
}

var fileDescriptor_telemetry_bis_proto2_be78366a61b59189 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0x86, 0x71, 0x92, 0x35, 0xf1, 0x27, 0x67, 0x5d, 0x54, 0x06, 0x62, 0x6c, 0xab, 0x9b, 0x5e,
	0x7c, 0xca, 0x21, 0x1d, 0x63, 0xa7, 0x0d, 0x0a, 0x5b, 0xd6, 0x5b, 0x50, 0xc3, 0x6e, 0xc3, 0xc8,
	0x91, 0x9a, 0x9a, 0xda, 0x52, 0xb0, 0x94, 0x96, 0xfc, 0xaa, 0xfd, 0xc5, 0x21, 0xc9, 0xae, 0x9d,
	0x0e, 0xba, 0x9e, 0xa2, 0x3c, 0xdf, 0xa3, 0x17, 0x49, 0xaf, 0xe1, 0x9d, 0x11, 0x85, 0x28, 0x85,
	0xa9, 0xf6, 0x69, 0x96, 0xeb, 0x74, 0x5b, 0x29, 0xa3, 0xe6, 0x33, 0xf7, 0x83, 0xdf, 0xb4, 0x33,
	0xcf, 0xa7, 0x7f, 0xfa, 0x10, 0xae, 0x1a, 0x88, 0x3f, 0x02, 0x92, 0x8a, 0x8b, 0x34, 0xe7, 0xa9,
	0x36, 0x15, 0x09, 0xe2, 0x20, 0x09, 0x69, 0x68, 0xd1, 0x15, 0xbf, 0x36, 0x15, 0x9e, 0xc1, 0x89,
	0xde, 0x65, 0x7a, 0x5d, 0xe5, 0x5b, 0x93, 0x2b, 0xd9, 0x78, 0x7d, 0xe7, 0x4d, 0xba, 0x23, 0xef,
	0x9f, 0xc3, 0x58, 0xc8, 0xb5, 0xe2, 0xb9, 0xdc, 0xa4, 0x5b, 0x66, 0x6e, 0xc9, 0x91, 0x33, 0xa3,
	0x06, 0x2e, 0x99, 0xb9, 0xb5, 0xd2, 0x5a, 0x15, 0x85, 0x58, 0xd7, 0x91, 0x64, 0x14, 0x07, 0xc9,
	0x80, 0x46, 0x2d, 0xbc, 0xe2, 0x78, 0x0e, 0x6f, 0x3b, 0x92, 0x36, 0xac, 0x32, 0xa9, 0xc9, 0x4b,
	0x41, 0x42, 0x27, 0x9f, 0xb4, 0xc3, 0x6b, 0x3b, 0x5b, 0xe5, 0xa5, 0xb0, 0xc1, 0xa5, 0xde, 0x38,
	0x4d, 0x1b, 0x56, 0x6e, 0x09, 0xf8, 0xe0, 0x52, 0x6f, 0x56, 0x0d, 0xc3, 0xdf, 0x00, 0x38, 0x33,
	0x2c, 0xdd, 0x6c, 0xb3, 0xbb, 0x7b, 0x82, 0xe2, 0x7e, 0x82, 0xe6, 0xf1, 0xec, 0xe9, 0x3b, 0xcd,
	0x1e, 0xdf, 0xe8, 0x47, 0x2e, 0x0a, 0x4e, 0x43, 0xbb, 0x67, 0x61, 0xb7, 0xe0, 0xaf, 0x30, 0x6a,
	0x02, 0x48, 0x14, 0x07, 0x09, 0x9a, 0x9f, 0x3f, 0xb3, 0x7d, 0xb1, 0xbc, 0x5c, 0xb1, 0xac, 0x10,
	0x74, 0x58, 0x27, 0xd8, 0x37, 0xed, 0xdc, 0x4c, 0x48, 0xee, 0xef, 0x35, 0x76, 0x67, 0x9d, 0xb4,
	0xa3, 0xef, 0x92, 0xdb, 0x53, 0xdb, 0xc6, 0x5e, 0x1f, 0x9e, 0x06, 0xbf, 0x87, 0xb0, 0xbd, 0x64,
	0xe0, 0x36, 0xb6, 0x00, 0x63, 0x18, 0x48, 0x56, 0x0a, 0xd2, 0x8b, 0x7b, 0x49, 0x48, 0xdd, 0x1a,
	0x9f, 0x02, 0xca, 0xf6, 0x46, 0xe8, 0xf4, 0x9e, 0x15, 0x3b, 0x41, 0x06, 0x71, 0x90, 0x44, 0x14,
	0x1c, 0xfa, 0x65, 0x09, 0x3e, 0x83, 0x48, 0x9b, 0xca, 0xf6, 0xe6, 0x8d, 0x57, 0xae, 0x38, 0xe4,
	0x99, 0x57, 0x3e, 0x00, 0x64, 0x4a, 0x15, 0xb5, 0x60, 0x9b, 0x1d, 0xd1, 0xd0, 0x92, 0xc7, 0x84,
	0x5d, 0x2e, 0xcd, 0xc5, 0xbc, 0x16, 0x86, 0x71, 0x90, 0x8c, 0x29, 0xf2, 0xec, 0x40, 0xf9, 0xfc,
	0xa9, 0x56, 0x7c, 0xf1, 0xc8, 0xb3, 0xf6, 0x1c, 0xdd, 0x14, 0x5b, 0xf7, 0x84, 0x22, 0x7d, 0x98,
	0xa2, 0xbb, 0x29, 0xb6, 0x65, 0xec, 0x95, 0x4e, 0x0a, 0x57, 0xbb, 0xac, 0x10, 0xb5, 0x82, 0xe2,
	0x20, 0x09, 0x28, 0xf2, 0xcc, 0x2b, 0xa7, 0x80, 0x6e, 0x0a, 0xc5, 0x4c, 0x6d, 0xd8, 0x26, 0x7b,
	0x14, 0x1c, 0xf2, 0xc2, 0x17, 0x38, 0xba, 0xb1, 0xaf, 0xad, 0xc9, 0xf1, 0x0b, 0x3f, 0x92, 0xda,
	0x9f, 0xfe, 0x84, 0xc9, 0x3f, 0xfd, 0xe3, 0x0b, 0xe8, 0x57, 0xea, 0x81, 0x04, 0x2e, 0xeb, 0xec,
	0x99, 0x2c, 0xaa, 0x1e, 0x16, 0xcb, 0x4b, 0x6a, 0xed, 0xe9, 0x6f, 0x38, 0x7e, 0xc2, 0xff, 0xdf,
	0xfd, 0x9d, 0xd8, 0x6b, 0x02, 0x71, 0x2f, 0x89, 0xa8, 0x5b, 0x63, 0x02, 0xc3, 0xb5, 0x92, 0x46,
	0x48, 0x43, 0x90, 0xc3, 0xcd, 0xdf, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x88, 0x1e, 0x41,
	0x3b, 0x04, 0x00, 0x00,
}
