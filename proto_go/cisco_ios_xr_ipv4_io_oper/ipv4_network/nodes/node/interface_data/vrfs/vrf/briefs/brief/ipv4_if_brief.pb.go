// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_if_brief.proto

package cisco_ios_xr_ipv4_io_oper_ipv4_network_nodes_node_interface_data_vrfs_vrf_briefs_brief

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

// Brief Summary of IP Interface
type Ipv4IfBrief_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfBrief_KEYS) Reset()         { *m = Ipv4IfBrief_KEYS{} }
func (m *Ipv4IfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfBrief_KEYS) ProtoMessage()    {}
func (*Ipv4IfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_if_brief_a48dfae5a8c14480, []int{0}
}
func (m *Ipv4IfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Unmarshal(m, b)
}
func (m *Ipv4IfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv4IfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfBrief_KEYS.Merge(dst, src)
}
func (m *Ipv4IfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Size(m)
}
func (m *Ipv4IfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfBrief_KEYS proto.InternalMessageInfo

func (m *Ipv4IfBrief_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv4IfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4IfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv4IfBrief struct {
	// Primary address
	PrimaryAddress string `protobuf:"bytes,50,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	// VRF ID of the interface
	VrfId uint32 `protobuf:"varint,51,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// Line state of the interface
	LineState            string   `protobuf:"bytes,52,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfBrief) Reset()         { *m = Ipv4IfBrief{} }
func (m *Ipv4IfBrief) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfBrief) ProtoMessage()    {}
func (*Ipv4IfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_if_brief_a48dfae5a8c14480, []int{1}
}
func (m *Ipv4IfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfBrief.Unmarshal(m, b)
}
func (m *Ipv4IfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfBrief.Marshal(b, m, deterministic)
}
func (dst *Ipv4IfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfBrief.Merge(dst, src)
}
func (m *Ipv4IfBrief) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfBrief.Size(m)
}
func (m *Ipv4IfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfBrief proto.InternalMessageInfo

func (m *Ipv4IfBrief) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *Ipv4IfBrief) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Ipv4IfBrief) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv4IfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv4_if_brief_KEYS")
	proto.RegisterType((*Ipv4IfBrief)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv4_if_brief")
}

func init() { proto.RegisterFile("ipv4_if_brief.proto", fileDescriptor_ipv4_if_brief_a48dfae5a8c14480) }

var fileDescriptor_ipv4_if_brief_a48dfae5a8c14480 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0xff, 0x1f, 0x6b, 0x77, 0x60, 0x2b, 0x44, 0x84, 0x15, 0x11, 0x4a, 0x41, 0xec,
	0x29, 0x07, 0xdb, 0x2f, 0xe0, 0xc1, 0x83, 0x08, 0x1e, 0x5a, 0x10, 0x3c, 0x0d, 0xe9, 0x66, 0x02,
	0x41, 0x37, 0x59, 0x26, 0x61, 0xd5, 0x6f, 0x2f, 0x99, 0x15, 0xa5, 0x97, 0x09, 0xf9, 0xbd, 0x37,
	0xef, 0xc1, 0xc0, 0xb9, 0x1f, 0xc6, 0x2d, 0x7a, 0x87, 0x07, 0xf6, 0xe4, 0xf4, 0xc0, 0x31, 0x47,
	0xf5, 0xd2, 0xf9, 0xd4, 0x45, 0xf4, 0x31, 0xe1, 0x27, 0xe3, 0xe4, 0x88, 0x18, 0x07, 0x62, 0x2d,
	0x9f, 0x40, 0xf9, 0x23, 0xf2, 0x9b, 0x0e, 0xd1, 0x52, 0x92, 0xa9, 0x7d, 0xc8, 0xc4, 0xce, 0x74,
	0x84, 0xd6, 0x64, 0xa3, 0x47, 0x76, 0xa9, 0x0c, 0x2d, 0xb1, 0x69, 0x7a, 0x56, 0x09, 0xd4, 0x51,
	0x1d, 0x3e, 0x3d, 0xbc, 0xee, 0xd5, 0x15, 0xd4, 0x25, 0x02, 0x83, 0xe9, 0xa9, 0xad, 0x96, 0xd5,
	0xba, 0xde, 0xcd, 0x0b, 0x78, 0x36, 0x3d, 0xa9, 0x4b, 0x98, 0x8f, 0xec, 0x26, 0xed, 0x9f, 0x68,
	0xa7, 0x23, 0x3b, 0x91, 0x6e, 0x60, 0xf1, 0xd7, 0x2a, 0x86, 0xff, 0x62, 0x68, 0x7e, 0x69, 0xb1,
	0xad, 0x02, 0x34, 0x47, 0xa5, 0xea, 0x16, 0xce, 0x06, 0xf6, 0xbd, 0xe1, 0x2f, 0x34, 0xd6, 0x32,
	0xa5, 0xd4, 0xde, 0xc9, 0xe2, 0xe2, 0x07, 0xdf, 0x4f, 0x54, 0x5d, 0xc0, 0xac, 0x74, 0x7b, 0xdb,
	0x6e, 0x96, 0xd5, 0xba, 0xd9, 0x9d, 0x8c, 0xec, 0x1e, 0xad, 0xba, 0x06, 0x78, 0xf7, 0x81, 0x30,
	0x65, 0x93, 0xa9, 0xdd, 0xca, 0x6a, 0x5d, 0xc8, 0xbe, 0x80, 0xc3, 0x4c, 0x6e, 0xb8, 0xf9, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x4a, 0xa4, 0x4a, 0x5a, 0x01, 0x00, 0x00,
}
