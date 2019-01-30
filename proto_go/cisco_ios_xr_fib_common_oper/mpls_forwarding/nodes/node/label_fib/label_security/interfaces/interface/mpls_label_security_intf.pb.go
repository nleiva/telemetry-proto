// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_label_security_intf.proto

package cisco_ios_xr_fib_common_oper_mpls_forwarding_nodes_node_label_fib_label_security_interfaces_interface

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

type MplsLabelSecurityIntf_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLabelSecurityIntf_KEYS) Reset()         { *m = MplsLabelSecurityIntf_KEYS{} }
func (m *MplsLabelSecurityIntf_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLabelSecurityIntf_KEYS) ProtoMessage()    {}
func (*MplsLabelSecurityIntf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_label_security_intf_38e39010b9ee43fb, []int{0}
}
func (m *MplsLabelSecurityIntf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLabelSecurityIntf_KEYS.Unmarshal(m, b)
}
func (m *MplsLabelSecurityIntf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLabelSecurityIntf_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsLabelSecurityIntf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLabelSecurityIntf_KEYS.Merge(dst, src)
}
func (m *MplsLabelSecurityIntf_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLabelSecurityIntf_KEYS.Size(m)
}
func (m *MplsLabelSecurityIntf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLabelSecurityIntf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLabelSecurityIntf_KEYS proto.InternalMessageInfo

func (m *MplsLabelSecurityIntf_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsLabelSecurityIntf_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsLabelSecurityIntf struct {
	// RPF is enabled on interface
	RpfEnabled bool `protobuf:"varint,50,opt,name=rpf_enabled,json=rpfEnabled,proto3" json:"rpf_enabled,omitempty"`
	// RPF stats supported per interface
	RpfSupported bool `protobuf:"varint,51,opt,name=rpf_supported,json=rpfSupported,proto3" json:"rpf_supported,omitempty"`
	// Multi-label drop is enabled on interface
	MldEnabled bool `protobuf:"varint,52,opt,name=mld_enabled,json=mldEnabled,proto3" json:"mld_enabled,omitempty"`
	// Multi-label drop counters are supported per interface
	MldSupported bool `protobuf:"varint,53,opt,name=mld_supported,json=mldSupported,proto3" json:"mld_supported,omitempty"`
	// RPF drops
	RpfDrops uint64 `protobuf:"varint,54,opt,name=rpf_drops,json=rpfDrops,proto3" json:"rpf_drops,omitempty"`
	// Multi-label drops
	MultiLabelDrops uint64 `protobuf:"varint,55,opt,name=multi_label_drops,json=multiLabelDrops,proto3" json:"multi_label_drops,omitempty"`
	// RPF interface handle
	Rpfifh               string   `protobuf:"bytes,56,opt,name=rpfifh,proto3" json:"rpfifh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLabelSecurityIntf) Reset()         { *m = MplsLabelSecurityIntf{} }
func (m *MplsLabelSecurityIntf) String() string { return proto.CompactTextString(m) }
func (*MplsLabelSecurityIntf) ProtoMessage()    {}
func (*MplsLabelSecurityIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_label_security_intf_38e39010b9ee43fb, []int{1}
}
func (m *MplsLabelSecurityIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLabelSecurityIntf.Unmarshal(m, b)
}
func (m *MplsLabelSecurityIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLabelSecurityIntf.Marshal(b, m, deterministic)
}
func (dst *MplsLabelSecurityIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLabelSecurityIntf.Merge(dst, src)
}
func (m *MplsLabelSecurityIntf) XXX_Size() int {
	return xxx_messageInfo_MplsLabelSecurityIntf.Size(m)
}
func (m *MplsLabelSecurityIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLabelSecurityIntf.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLabelSecurityIntf proto.InternalMessageInfo

func (m *MplsLabelSecurityIntf) GetRpfEnabled() bool {
	if m != nil {
		return m.RpfEnabled
	}
	return false
}

func (m *MplsLabelSecurityIntf) GetRpfSupported() bool {
	if m != nil {
		return m.RpfSupported
	}
	return false
}

func (m *MplsLabelSecurityIntf) GetMldEnabled() bool {
	if m != nil {
		return m.MldEnabled
	}
	return false
}

func (m *MplsLabelSecurityIntf) GetMldSupported() bool {
	if m != nil {
		return m.MldSupported
	}
	return false
}

func (m *MplsLabelSecurityIntf) GetRpfDrops() uint64 {
	if m != nil {
		return m.RpfDrops
	}
	return 0
}

func (m *MplsLabelSecurityIntf) GetMultiLabelDrops() uint64 {
	if m != nil {
		return m.MultiLabelDrops
	}
	return 0
}

func (m *MplsLabelSecurityIntf) GetRpfifh() string {
	if m != nil {
		return m.Rpfifh
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsLabelSecurityIntf_KEYS)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.label_security.interfaces.interface.mpls_label_security_intf_KEYS")
	proto.RegisterType((*MplsLabelSecurityIntf)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.label_security.interfaces.interface.mpls_label_security_intf")
}

func init() {
	proto.RegisterFile("mpls_label_security_intf.proto", fileDescriptor_mpls_label_security_intf_38e39010b9ee43fb)
}

var fileDescriptor_mpls_label_security_intf_38e39010b9ee43fb = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcb, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x99, 0x22, 0xa5, 0x8d, 0x56, 0x71, 0x16, 0x32, 0x50, 0xd4, 0x52, 0x11, 0x8a, 0x8b,
	0x59, 0x58, 0x6f, 0x0f, 0x60, 0x57, 0x8a, 0x8b, 0x76, 0xe5, 0x2a, 0xcc, 0x24, 0x27, 0x1a, 0xc8,
	0x8d, 0x93, 0x14, 0xf5, 0x09, 0x7c, 0x6d, 0x49, 0x52, 0x3b, 0xba, 0xe8, 0x26, 0xe4, 0xfc, 0xf9,
	0xf2, 0x9d, 0x5c, 0xc8, 0x99, 0x76, 0xca, 0x53, 0xd5, 0xb4, 0xa0, 0xa8, 0x07, 0xb6, 0x46, 0x19,
	0xbe, 0xa8, 0x34, 0x41, 0xd4, 0x0e, 0x6d, 0xb0, 0x25, 0x30, 0xe9, 0x99, 0xa5, 0xd2, 0x7a, 0xfa,
	0x89, 0x54, 0xc8, 0x96, 0x32, 0xab, 0xb5, 0x35, 0xd4, 0x3a, 0xc0, 0x3a, 0x6d, 0x16, 0x16, 0x3f,
	0x1a, 0xe4, 0xd2, 0xbc, 0xd5, 0xc6, 0x72, 0xf0, 0x69, 0xac, 0xb3, 0x52, 0xc8, 0xb6, 0xfe, 0x2f,
	0xaf, 0xa5, 0x09, 0x80, 0xa2, 0x61, 0xe0, 0xbb, 0xe9, 0x94, 0x91, 0xd3, 0x5d, 0x07, 0xa1, 0x4f,
	0x8b, 0xd7, 0x55, 0x39, 0x26, 0xc3, 0xa8, 0xa5, 0xa6, 0xd1, 0x50, 0x15, 0x93, 0x62, 0x36, 0x5c,
	0x0e, 0x62, 0xf0, 0xd2, 0x68, 0x28, 0x2f, 0xc9, 0xe1, 0x56, 0x95, 0x89, 0x5e, 0x22, 0x46, 0xdb,
	0x34, 0x62, 0xd3, 0xef, 0x1e, 0xa9, 0x76, 0x75, 0x29, 0xcf, 0xc9, 0x3e, 0x3a, 0x41, 0xc1, 0x34,
	0xad, 0x02, 0x5e, 0x5d, 0x4f, 0x8a, 0xd9, 0x60, 0x49, 0xd0, 0x89, 0x45, 0x4e, 0xca, 0x0b, 0x32,
	0x8a, 0x80, 0x5f, 0x3b, 0x67, 0x31, 0x00, 0xaf, 0xe6, 0x09, 0x39, 0x40, 0x27, 0x56, 0xbf, 0x59,
	0xb4, 0x68, 0xc5, 0xb7, 0x96, 0x9b, 0x6c, 0xd1, 0x8a, 0xff, 0xb1, 0x44, 0xa0, 0xb3, 0xdc, 0x66,
	0x8b, 0x56, 0xbc, 0xb3, 0x8c, 0xc9, 0x30, 0xb6, 0xe2, 0x68, 0x9d, 0xaf, 0xee, 0x26, 0xc5, 0x6c,
	0x6f, 0x39, 0x40, 0x27, 0x1e, 0x63, 0x5d, 0x5e, 0x91, 0x63, 0xbd, 0x56, 0x41, 0x6e, 0x6e, 0x91,
	0xa1, 0xfb, 0x04, 0x1d, 0xa5, 0x85, 0xe7, 0x98, 0x67, 0xf6, 0x84, 0xf4, 0xd1, 0x09, 0x29, 0xde,
	0xab, 0x87, 0xf4, 0x20, 0x9b, 0xaa, 0xed, 0xa7, 0xcf, 0x9d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0xe8, 0x23, 0x20, 0xfe, 0x01, 0x00, 0x00,
}
