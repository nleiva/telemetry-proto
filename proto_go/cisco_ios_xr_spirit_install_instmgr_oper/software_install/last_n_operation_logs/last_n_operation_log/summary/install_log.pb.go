// Code generated by protoc-gen-go. DO NOT EDIT.
// source: install_log.proto

package cisco_ios_xr_spirit_install_instmgr_oper_software_install_last_n_operation_logs_last_n_operation_log_summary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Install Log
type InstallLog_KEYS struct {
	LastNLogs            uint32   `protobuf:"varint,1,opt,name=last_n_logs,json=lastNLogs,proto3" json:"last_n_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallLog_KEYS) Reset()         { *m = InstallLog_KEYS{} }
func (m *InstallLog_KEYS) String() string { return proto.CompactTextString(m) }
func (*InstallLog_KEYS) ProtoMessage()    {}
func (*InstallLog_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0e50cc9d3e98eb, []int{0}
}

func (m *InstallLog_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallLog_KEYS.Unmarshal(m, b)
}
func (m *InstallLog_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallLog_KEYS.Marshal(b, m, deterministic)
}
func (m *InstallLog_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallLog_KEYS.Merge(m, src)
}
func (m *InstallLog_KEYS) XXX_Size() int {
	return xxx_messageInfo_InstallLog_KEYS.Size(m)
}
func (m *InstallLog_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallLog_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InstallLog_KEYS proto.InternalMessageInfo

func (m *InstallLog_KEYS) GetLastNLogs() uint32 {
	if m != nil {
		return m.LastNLogs
	}
	return 0
}

type InstallLog struct {
	// log
	Log                  string   `protobuf:"bytes,50,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallLog) Reset()         { *m = InstallLog{} }
func (m *InstallLog) String() string { return proto.CompactTextString(m) }
func (*InstallLog) ProtoMessage()    {}
func (*InstallLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa0e50cc9d3e98eb, []int{1}
}

func (m *InstallLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallLog.Unmarshal(m, b)
}
func (m *InstallLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallLog.Marshal(b, m, deterministic)
}
func (m *InstallLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallLog.Merge(m, src)
}
func (m *InstallLog) XXX_Size() int {
	return xxx_messageInfo_InstallLog.Size(m)
}
func (m *InstallLog) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallLog.DiscardUnknown(m)
}

var xxx_messageInfo_InstallLog proto.InternalMessageInfo

func (m *InstallLog) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterType((*InstallLog_KEYS)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.last_n_operation_logs.last_n_operation_log.summary.install_log_KEYS")
	proto.RegisterType((*InstallLog)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.last_n_operation_logs.last_n_operation_log.summary.install_log")
}

func init() { proto.RegisterFile("install_log.proto", fileDescriptor_fa0e50cc9d3e98eb) }

var fileDescriptor_fa0e50cc9d3e98eb = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x0a, 0xc2, 0x40,
	0x0c, 0x86, 0x29, 0x82, 0xd0, 0x2b, 0x42, 0xed, 0xd4, 0x49, 0x4b, 0xa7, 0x4e, 0x37, 0xd4, 0x67,
	0x70, 0x52, 0x1c, 0xea, 0xe4, 0x14, 0xce, 0x52, 0x8f, 0x83, 0x6b, 0x53, 0x92, 0x88, 0xfa, 0xf6,
	0x72, 0xc5, 0x42, 0x07, 0xa7, 0x84, 0xff, 0xcb, 0x17, 0x12, 0xb5, 0x75, 0x03, 0x8b, 0xf1, 0x1e,
	0x3c, 0x5a, 0x3d, 0x12, 0x0a, 0x66, 0xbe, 0x75, 0xdc, 0x22, 0x38, 0x64, 0x78, 0x13, 0xf0, 0xe8,
	0xc8, 0x09, 0xcc, 0x63, 0xa1, 0xf6, 0x96, 0x00, 0xc7, 0x8e, 0x34, 0xe3, 0x43, 0x5e, 0x86, 0xba,
	0x99, 0x6a, 0x6f, 0x58, 0x60, 0x98, 0xa0, 0x11, 0x87, 0x43, 0xd8, 0xca, 0x7f, 0x53, 0xcd, 0xcf,
	0xbe, 0x37, 0xf4, 0x29, 0x6b, 0x95, 0x2e, 0x4e, 0x80, 0xd3, 0xf1, 0x76, 0xcd, 0x76, 0x2a, 0xf9,
	0x39, 0xc1, 0xcf, 0xa3, 0x22, 0xaa, 0x36, 0x4d, 0x1c, 0xa2, 0xcb, 0x19, 0x2d, 0x97, 0x7b, 0x95,
	0x2c, 0x9c, 0x2c, 0x55, 0x2b, 0x8f, 0x36, 0xaf, 0x8b, 0xa8, 0x8a, 0x9b, 0xd0, 0xde, 0xd7, 0xd3,
	0x27, 0x87, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x41, 0xac, 0xa0, 0xde, 0x00, 0x00, 0x00,
}
