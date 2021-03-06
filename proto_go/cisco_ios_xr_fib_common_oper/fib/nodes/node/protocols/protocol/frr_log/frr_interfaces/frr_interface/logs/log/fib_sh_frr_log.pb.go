// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fib_sh_frr_log.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_frr_log_frr_interfaces_frr_interface_logs_log

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

// FIB frr log information
type FibShFrrLog_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	FrrInterfaceName     string   `protobuf:"bytes,3,opt,name=frr_interface_name,json=frrInterfaceName,proto3" json:"frr_interface_name,omitempty"`
	LogIndex             uint32   `protobuf:"varint,4,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShFrrLog_KEYS) Reset()         { *m = FibShFrrLog_KEYS{} }
func (m *FibShFrrLog_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibShFrrLog_KEYS) ProtoMessage()    {}
func (*FibShFrrLog_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_frr_log_3f71f8fba0051a94, []int{0}
}
func (m *FibShFrrLog_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShFrrLog_KEYS.Unmarshal(m, b)
}
func (m *FibShFrrLog_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShFrrLog_KEYS.Marshal(b, m, deterministic)
}
func (dst *FibShFrrLog_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShFrrLog_KEYS.Merge(dst, src)
}
func (m *FibShFrrLog_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibShFrrLog_KEYS.Size(m)
}
func (m *FibShFrrLog_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShFrrLog_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibShFrrLog_KEYS proto.InternalMessageInfo

func (m *FibShFrrLog_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibShFrrLog_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *FibShFrrLog_KEYS) GetFrrInterfaceName() string {
	if m != nil {
		return m.FrrInterfaceName
	}
	return ""
}

func (m *FibShFrrLog_KEYS) GetLogIndex() uint32 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

type FibShFrrLog struct {
	// FIB Protocol Type
	FrrProtocolType string `protobuf:"bytes,50,opt,name=frr_protocol_type,json=frrProtocolType,proto3" json:"frr_protocol_type,omitempty"`
	// Interface assoc w frr nh
	FrrInterfaceName string `protobuf:"bytes,51,opt,name=frr_interface_name,json=frrInterfaceName,proto3" json:"frr_interface_name,omitempty"`
	// nh prefix
	FrrPrefix string `protobuf:"bytes,52,opt,name=frr_prefix,json=frrPrefix,proto3" json:"frr_prefix,omitempty"`
	// frr timestamp
	FrrTimestamp *FibShTimespec `protobuf:"bytes,53,opt,name=frr_timestamp,json=frrTimestamp,proto3" json:"frr_timestamp,omitempty"`
	// frr switching time
	FrrSwitchingTime uint32 `protobuf:"varint,54,opt,name=frr_switching_time,json=frrSwitchingTime,proto3" json:"frr_switching_time,omitempty"`
	// bundle member
	BundleMemberInterfaceName string   `protobuf:"bytes,55,opt,name=bundle_member_interface_name,json=bundleMemberInterfaceName,proto3" json:"bundle_member_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *FibShFrrLog) Reset()         { *m = FibShFrrLog{} }
func (m *FibShFrrLog) String() string { return proto.CompactTextString(m) }
func (*FibShFrrLog) ProtoMessage()    {}
func (*FibShFrrLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_frr_log_3f71f8fba0051a94, []int{1}
}
func (m *FibShFrrLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShFrrLog.Unmarshal(m, b)
}
func (m *FibShFrrLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShFrrLog.Marshal(b, m, deterministic)
}
func (dst *FibShFrrLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShFrrLog.Merge(dst, src)
}
func (m *FibShFrrLog) XXX_Size() int {
	return xxx_messageInfo_FibShFrrLog.Size(m)
}
func (m *FibShFrrLog) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShFrrLog.DiscardUnknown(m)
}

var xxx_messageInfo_FibShFrrLog proto.InternalMessageInfo

func (m *FibShFrrLog) GetFrrProtocolType() string {
	if m != nil {
		return m.FrrProtocolType
	}
	return ""
}

func (m *FibShFrrLog) GetFrrInterfaceName() string {
	if m != nil {
		return m.FrrInterfaceName
	}
	return ""
}

func (m *FibShFrrLog) GetFrrPrefix() string {
	if m != nil {
		return m.FrrPrefix
	}
	return ""
}

func (m *FibShFrrLog) GetFrrTimestamp() *FibShTimespec {
	if m != nil {
		return m.FrrTimestamp
	}
	return nil
}

func (m *FibShFrrLog) GetFrrSwitchingTime() uint32 {
	if m != nil {
		return m.FrrSwitchingTime
	}
	return 0
}

func (m *FibShFrrLog) GetBundleMemberInterfaceName() string {
	if m != nil {
		return m.BundleMemberInterfaceName
	}
	return ""
}

type FibShTimespec struct {
	Seconds              int32    `protobuf:"zigzag32,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	NanoSeconds          int32    `protobuf:"zigzag32,2,opt,name=nano_seconds,json=nanoSeconds,proto3" json:"nano_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShTimespec) Reset()         { *m = FibShTimespec{} }
func (m *FibShTimespec) String() string { return proto.CompactTextString(m) }
func (*FibShTimespec) ProtoMessage()    {}
func (*FibShTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_frr_log_3f71f8fba0051a94, []int{2}
}
func (m *FibShTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShTimespec.Unmarshal(m, b)
}
func (m *FibShTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShTimespec.Marshal(b, m, deterministic)
}
func (dst *FibShTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShTimespec.Merge(dst, src)
}
func (m *FibShTimespec) XXX_Size() int {
	return xxx_messageInfo_FibShTimespec.Size(m)
}
func (m *FibShTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_FibShTimespec proto.InternalMessageInfo

func (m *FibShTimespec) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *FibShTimespec) GetNanoSeconds() int32 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*FibShFrrLog_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.frr_log.frr_interfaces.frr_interface.logs.log.fib_sh_frr_log_KEYS")
	proto.RegisterType((*FibShFrrLog)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.frr_log.frr_interfaces.frr_interface.logs.log.fib_sh_frr_log")
	proto.RegisterType((*FibShTimespec)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.frr_log.frr_interfaces.frr_interface.logs.log.fib_sh_timespec")
}

func init() {
	proto.RegisterFile("fib_sh_frr_log.proto", fileDescriptor_fib_sh_frr_log_3f71f8fba0051a94)
}

var fileDescriptor_fib_sh_frr_log_3f71f8fba0051a94 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x3f, 0x8f, 0xda, 0x30,
	0x18, 0xc6, 0x15, 0xa8, 0x5a, 0x62, 0xa0, 0x14, 0xb7, 0x43, 0x2a, 0x5a, 0x89, 0xd2, 0x05, 0x55,
	0x55, 0x06, 0xe8, 0x9f, 0xb1, 0x53, 0x07, 0x74, 0x3a, 0x74, 0x0a, 0x2c, 0x37, 0x59, 0x89, 0x63,
	0x07, 0x4b, 0xb1, 0x1d, 0xd9, 0x39, 0x1d, 0x0c, 0xf7, 0x51, 0x6e, 0xe1, 0x93, 0x9e, 0xfc, 0x06,
	0x23, 0xe5, 0x4e, 0xac, 0xb7, 0x18, 0xfb, 0x79, 0x7f, 0x7e, 0xdf, 0xc7, 0x0f, 0x41, 0x9f, 0xb8,
	0xc8, 0x88, 0xdd, 0x11, 0x6e, 0x0c, 0x29, 0x75, 0x11, 0x57, 0x46, 0xd7, 0x1a, 0x97, 0x54, 0x58,
	0xaa, 0x89, 0xd0, 0x96, 0xec, 0x0d, 0x71, 0x08, 0xd5, 0x52, 0x6a, 0x45, 0x74, 0xc5, 0x4c, 0xcc,
	0x45, 0x16, 0x2b, 0x9d, 0x33, 0x0b, 0x6b, 0x73, 0x85, 0xea, 0xd2, 0x9e, 0x77, 0xb1, 0xef, 0xe6,
	0x7e, 0x85, 0xaa, 0x99, 0xe1, 0x29, 0x65, 0xb6, 0x7d, 0x8c, 0x4b, 0x5d, 0x58, 0xb7, 0xcc, 0x8e,
	0x01, 0xfa, 0xd8, 0xb6, 0x41, 0xae, 0xfe, 0xdf, 0x6e, 0xf0, 0x04, 0x85, 0x6e, 0x00, 0x51, 0xa9,
	0x64, 0x51, 0x30, 0x0d, 0xe6, 0x61, 0xd2, 0x73, 0xc2, 0x3a, 0x95, 0x0c, 0x7f, 0x47, 0x43, 0x3f,
	0xae, 0x01, 0x3a, 0x00, 0x0c, 0xbc, 0x08, 0xd0, 0x4f, 0x84, 0x5b, 0x33, 0x1b, 0xb2, 0x0b, 0xe4,
	0x07, 0x6e, 0xcc, 0xca, 0x17, 0x80, 0x9e, 0xa0, 0xd0, 0xcd, 0x16, 0x2a, 0x67, 0xfb, 0xe8, 0xcd,
	0x34, 0x98, 0x0f, 0x93, 0x5e, 0xa9, 0x8b, 0x95, 0x3b, 0xcf, 0x1e, 0xbb, 0xe8, 0x7d, 0xdb, 0x24,
	0xfe, 0x81, 0xc6, 0x6e, 0x7b, 0xb6, 0x51, 0x1f, 0x2a, 0x16, 0x2d, 0xa0, 0xf9, 0x88, 0x1b, 0x73,
	0x73, 0xd2, 0xb7, 0x87, 0xea, 0x92, 0x93, 0xe5, 0x05, 0x27, 0x5f, 0x11, 0x6a, 0x3a, 0x33, 0x2e,
	0xf6, 0xd1, 0x2f, 0xa0, 0x42, 0x68, 0xe9, 0x04, 0x7c, 0x0c, 0xd0, 0xd0, 0xd5, 0x6b, 0x21, 0x99,
	0xad, 0x53, 0x59, 0x45, 0xbf, 0xa7, 0xc1, 0xbc, 0xbf, 0x78, 0x88, 0x5f, 0xf3, 0x7f, 0x8b, 0x4f,
	0x71, 0x80, 0x8b, 0x8a, 0xd1, 0x64, 0xc0, 0x8d, 0xd9, 0x7a, 0x4b, 0xfe, 0xc5, 0xf6, 0x5e, 0xd4,
	0x74, 0x27, 0x54, 0x01, 0x5c, 0xf4, 0x07, 0x62, 0x75, 0x2f, 0xde, 0xf8, 0x82, 0xbb, 0x82, 0xff,
	0xa1, 0x2f, 0xd9, 0x9d, 0xca, 0x4b, 0x46, 0x24, 0x93, 0x19, 0x7b, 0x91, 0xd4, 0x5f, 0xc8, 0xe0,
	0x73, 0xc3, 0x5c, 0x03, 0xd2, 0x8a, 0x6c, 0xb6, 0x46, 0xa3, 0x67, 0x7e, 0x70, 0x84, 0xde, 0x59,
	0x46, 0xb5, 0xca, 0x2d, 0x7c, 0x3d, 0xe3, 0xc4, 0x1f, 0xf1, 0x37, 0x34, 0x50, 0xa9, 0xd2, 0xc4,
	0x97, 0x3b, 0x50, 0xee, 0x3b, 0x6d, 0xd3, 0x48, 0xd9, 0x5b, 0x08, 0x65, 0xf9, 0x14, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0x49, 0xeb, 0xe3, 0x21, 0x03, 0x00, 0x00,
}
