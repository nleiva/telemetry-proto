// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_err_log_struc.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_error_log

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

// Error log structure
type IsisShErrLogStruc_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShErrLogStruc_KEYS) Reset()         { *m = IsisShErrLogStruc_KEYS{} }
func (m *IsisShErrLogStruc_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShErrLogStruc_KEYS) ProtoMessage()    {}
func (*IsisShErrLogStruc_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee, []int{0}
}
func (m *IsisShErrLogStruc_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShErrLogStruc_KEYS.Unmarshal(m, b)
}
func (m *IsisShErrLogStruc_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShErrLogStruc_KEYS.Marshal(b, m, deterministic)
}
func (dst *IsisShErrLogStruc_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShErrLogStruc_KEYS.Merge(dst, src)
}
func (m *IsisShErrLogStruc_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShErrLogStruc_KEYS.Size(m)
}
func (m *IsisShErrLogStruc_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShErrLogStruc_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShErrLogStruc_KEYS proto.InternalMessageInfo

func (m *IsisShErrLogStruc_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type IsisShErrLogStruc struct {
	// Error Log entries
	LogEntry             []*IsisShErrLogEnt `protobuf:"bytes,50,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IsisShErrLogStruc) Reset()         { *m = IsisShErrLogStruc{} }
func (m *IsisShErrLogStruc) String() string { return proto.CompactTextString(m) }
func (*IsisShErrLogStruc) ProtoMessage()    {}
func (*IsisShErrLogStruc) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee, []int{1}
}
func (m *IsisShErrLogStruc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShErrLogStruc.Unmarshal(m, b)
}
func (m *IsisShErrLogStruc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShErrLogStruc.Marshal(b, m, deterministic)
}
func (dst *IsisShErrLogStruc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShErrLogStruc.Merge(dst, src)
}
func (m *IsisShErrLogStruc) XXX_Size() int {
	return xxx_messageInfo_IsisShErrLogStruc.Size(m)
}
func (m *IsisShErrLogStruc) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShErrLogStruc.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShErrLogStruc proto.InternalMessageInfo

func (m *IsisShErrLogStruc) GetLogEntry() []*IsisShErrLogEnt {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

// Timestamp for an event
type IsisShTimestampType struct {
	// Timestamp value (seconds)
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Timestamp value (nanoseconds)
	NanoSeconds          uint32   `protobuf:"varint,2,opt,name=nano_seconds,json=nanoSeconds,proto3" json:"nano_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShTimestampType) Reset()         { *m = IsisShTimestampType{} }
func (m *IsisShTimestampType) String() string { return proto.CompactTextString(m) }
func (*IsisShTimestampType) ProtoMessage()    {}
func (*IsisShTimestampType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee, []int{2}
}
func (m *IsisShTimestampType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTimestampType.Unmarshal(m, b)
}
func (m *IsisShTimestampType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTimestampType.Marshal(b, m, deterministic)
}
func (dst *IsisShTimestampType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTimestampType.Merge(dst, src)
}
func (m *IsisShTimestampType) XXX_Size() int {
	return xxx_messageInfo_IsisShTimestampType.Size(m)
}
func (m *IsisShTimestampType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTimestampType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTimestampType proto.InternalMessageInfo

func (m *IsisShTimestampType) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *IsisShTimestampType) GetNanoSeconds() uint32 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

// Generic portion of a log entry
type IsisShGenericLogEnt struct {
	// Time in UTC relative to Jan 1st, 1970
	Timestamp            *IsisShTimestampType `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IsisShGenericLogEnt) Reset()         { *m = IsisShGenericLogEnt{} }
func (m *IsisShGenericLogEnt) String() string { return proto.CompactTextString(m) }
func (*IsisShGenericLogEnt) ProtoMessage()    {}
func (*IsisShGenericLogEnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee, []int{3}
}
func (m *IsisShGenericLogEnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShGenericLogEnt.Unmarshal(m, b)
}
func (m *IsisShGenericLogEnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShGenericLogEnt.Marshal(b, m, deterministic)
}
func (dst *IsisShGenericLogEnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShGenericLogEnt.Merge(dst, src)
}
func (m *IsisShGenericLogEnt) XXX_Size() int {
	return xxx_messageInfo_IsisShGenericLogEnt.Size(m)
}
func (m *IsisShGenericLogEnt) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShGenericLogEnt.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShGenericLogEnt proto.InternalMessageInfo

func (m *IsisShGenericLogEnt) GetTimestamp() *IsisShTimestampType {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// An error log entry
type IsisShErrLogEnt struct {
	// Generic entry data
	GenericData *IsisShGenericLogEnt `protobuf:"bytes,1,opt,name=generic_data,json=genericData,proto3" json:"generic_data,omitempty"`
	// Logging level
	ErrorLogLevel string `protobuf:"bytes,2,opt,name=error_log_level,json=errorLogLevel,proto3" json:"error_log_level,omitempty"`
	// Error code
	ErrorCode uint32 `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// Error details
	ErrorReason          string   `protobuf:"bytes,4,opt,name=error_reason,json=errorReason,proto3" json:"error_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShErrLogEnt) Reset()         { *m = IsisShErrLogEnt{} }
func (m *IsisShErrLogEnt) String() string { return proto.CompactTextString(m) }
func (*IsisShErrLogEnt) ProtoMessage()    {}
func (*IsisShErrLogEnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee, []int{4}
}
func (m *IsisShErrLogEnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShErrLogEnt.Unmarshal(m, b)
}
func (m *IsisShErrLogEnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShErrLogEnt.Marshal(b, m, deterministic)
}
func (dst *IsisShErrLogEnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShErrLogEnt.Merge(dst, src)
}
func (m *IsisShErrLogEnt) XXX_Size() int {
	return xxx_messageInfo_IsisShErrLogEnt.Size(m)
}
func (m *IsisShErrLogEnt) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShErrLogEnt.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShErrLogEnt proto.InternalMessageInfo

func (m *IsisShErrLogEnt) GetGenericData() *IsisShGenericLogEnt {
	if m != nil {
		return m.GenericData
	}
	return nil
}

func (m *IsisShErrLogEnt) GetErrorLogLevel() string {
	if m != nil {
		return m.ErrorLogLevel
	}
	return ""
}

func (m *IsisShErrLogEnt) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *IsisShErrLogEnt) GetErrorReason() string {
	if m != nil {
		return m.ErrorReason
	}
	return ""
}

func init() {
	proto.RegisterType((*IsisShErrLogStruc_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.error_log.isis_sh_err_log_struc_KEYS")
	proto.RegisterType((*IsisShErrLogStruc)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.error_log.isis_sh_err_log_struc")
	proto.RegisterType((*IsisShTimestampType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.error_log.isis_sh_timestamp_type")
	proto.RegisterType((*IsisShGenericLogEnt)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.error_log.isis_sh_generic_log_ent")
	proto.RegisterType((*IsisShErrLogEnt)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.error_log.isis_sh_err_log_ent")
}

func init() {
	proto.RegisterFile("isis_sh_err_log_struc.proto", fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee)
}

var fileDescriptor_isis_sh_err_log_struc_9d5c623707242eee = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4b, 0xe3, 0x40,
	0x14, 0xc6, 0xc9, 0x76, 0xd9, 0xdd, 0xbc, 0xb4, 0x2c, 0xcc, 0xb2, 0x1a, 0x14, 0xa1, 0x46, 0x90,
	0x9e, 0x72, 0xa8, 0x67, 0x0f, 0xa2, 0x3d, 0x59, 0x3c, 0x4c, 0xa9, 0xe0, 0x69, 0x18, 0x27, 0x8f,
	0x18, 0x48, 0xe6, 0x85, 0x99, 0x51, 0xec, 0x7f, 0xe0, 0xc9, 0x3f, 0xd8, 0x93, 0x64, 0xda, 0x69,
	0xb1, 0xf4, 0x66, 0x6f, 0x33, 0xbf, 0xf7, 0xe6, 0xfb, 0xde, 0xf7, 0x12, 0x38, 0xae, 0x6c, 0x65,
	0x85, 0x7d, 0x12, 0x68, 0x8c, 0xa8, 0xa9, 0x14, 0xd6, 0x99, 0x67, 0x95, 0xb7, 0x86, 0x1c, 0xb1,
	0x4b, 0x55, 0x59, 0x45, 0xa2, 0x22, 0x2b, 0x5e, 0x8d, 0x50, 0xb5, 0xb6, 0xc2, 0xb7, 0x53, 0x8b,
	0x26, 0xef, 0x4e, 0x79, 0xa5, 0xad, 0x93, 0x5a, 0xe1, 0xe6, 0x94, 0xa3, 0x31, 0xe4, 0xa5, 0xb2,
	0x2b, 0x38, 0xda, 0xa9, 0x2e, 0x6e, 0x27, 0x0f, 0x33, 0x76, 0x06, 0x83, 0xf0, 0x46, 0x68, 0xd9,
	0x60, 0x1a, 0x0d, 0xa3, 0x51, 0xcc, 0xfb, 0x01, 0xde, 0xc9, 0x06, 0xb3, 0xb7, 0x08, 0xfe, 0xef,
	0xd4, 0x60, 0x04, 0x71, 0x77, 0x41, 0xed, 0xcc, 0x22, 0x1d, 0x0f, 0x7b, 0xa3, 0x64, 0xcc, 0xf3,
	0x6f, 0xcd, 0x9b, 0x6f, 0x1b, 0xa1, 0x76, 0xfc, 0x4f, 0x4d, 0xe5, 0xa4, 0xf3, 0xc8, 0xe6, 0x70,
	0x10, 0x1a, 0x5c, 0xd5, 0xa0, 0x75, 0xb2, 0x69, 0x85, 0x5b, 0xb4, 0xc8, 0x52, 0xf8, 0x6d, 0x51,
	0x91, 0x2e, 0xac, 0xcf, 0x30, 0xe0, 0xe1, 0xca, 0x4e, 0xa1, 0xaf, 0xa5, 0x26, 0x11, 0xca, 0x3f,
	0x7c, 0x39, 0xe9, 0xd8, 0x6c, 0x89, 0xb2, 0xf7, 0x08, 0x0e, 0x83, 0x6e, 0x89, 0x1a, 0x4d, 0xa5,
	0x82, 0x39, 0xb3, 0x10, 0xaf, 0xad, 0xbc, 0x74, 0x32, 0x9e, 0xef, 0x29, 0xe3, 0xd7, 0x08, 0x7c,
	0xe3, 0x93, 0x7d, 0x44, 0xf0, 0x6f, 0xc7, 0x26, 0xd8, 0x02, 0xfa, 0x61, 0xbe, 0x42, 0x3a, 0xb9,
	0x9a, 0xe7, 0x7e, 0x4f, 0xf3, 0x6c, 0x45, 0xe7, 0xc9, 0x0a, 0xdc, 0x48, 0x27, 0xd9, 0x39, 0xfc,
	0x5d, 0xbf, 0x10, 0x35, 0xbe, 0x60, 0xed, 0x37, 0x19, 0xf3, 0x81, 0xc7, 0x53, 0x2a, 0xa7, 0x1d,
	0x64, 0x27, 0x00, 0xcb, 0x3e, 0x45, 0x05, 0xa6, 0x3d, 0xbf, 0xec, 0xd8, 0x93, 0x6b, 0x2a, 0xb0,
	0xfb, 0x1a, 0xcb, 0xb2, 0x41, 0x69, 0x49, 0xa7, 0x3f, 0xbd, 0x46, 0xe2, 0x19, 0xf7, 0xe8, 0xf1,
	0x97, 0xff, 0xf1, 0x2f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xa1, 0x88, 0x52, 0x17, 0x03,
	0x00, 0x00,
}
