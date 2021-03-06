// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_if_brief.proto

package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_vrfs_vrf_global_briefs_global_brief

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

// Brief Summary of IPv6 Interface
type Ipv6IfBrief_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6IfBrief_KEYS) Reset()         { *m = Ipv6IfBrief_KEYS{} }
func (m *Ipv6IfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6IfBrief_KEYS) ProtoMessage()    {}
func (*Ipv6IfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_brief_5bb3f7d8b3180cb3, []int{0}
}
func (m *Ipv6IfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6IfBrief_KEYS.Unmarshal(m, b)
}
func (m *Ipv6IfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6IfBrief_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv6IfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6IfBrief_KEYS.Merge(dst, src)
}
func (m *Ipv6IfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6IfBrief_KEYS.Size(m)
}
func (m *Ipv6IfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6IfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6IfBrief_KEYS proto.InternalMessageInfo

func (m *Ipv6IfBrief_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6IfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6IfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6IfBrief struct {
	// State of Interface Line
	LineState string `protobuf:"bytes,50,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	// Address List
	AddressList []*Ipv6AddrNode `protobuf:"bytes,51,rep,name=address_list,json=addressList,proto3" json:"address_list,omitempty"`
	// Link Local Address
	LinkLocalAddress     *Ipv6AddrNode `protobuf:"bytes,52,opt,name=link_local_address,json=linkLocalAddress,proto3" json:"link_local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ipv6IfBrief) Reset()         { *m = Ipv6IfBrief{} }
func (m *Ipv6IfBrief) String() string { return proto.CompactTextString(m) }
func (*Ipv6IfBrief) ProtoMessage()    {}
func (*Ipv6IfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_brief_5bb3f7d8b3180cb3, []int{1}
}
func (m *Ipv6IfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6IfBrief.Unmarshal(m, b)
}
func (m *Ipv6IfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6IfBrief.Marshal(b, m, deterministic)
}
func (dst *Ipv6IfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6IfBrief.Merge(dst, src)
}
func (m *Ipv6IfBrief) XXX_Size() int {
	return xxx_messageInfo_Ipv6IfBrief.Size(m)
}
func (m *Ipv6IfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6IfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6IfBrief proto.InternalMessageInfo

func (m *Ipv6IfBrief) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *Ipv6IfBrief) GetAddressList() []*Ipv6AddrNode {
	if m != nil {
		return m.AddressList
	}
	return nil
}

func (m *Ipv6IfBrief) GetLinkLocalAddress() *Ipv6AddrNode {
	if m != nil {
		return m.LinkLocalAddress
	}
	return nil
}

// List of IPv6 Addresses
type Ipv6AddrNode struct {
	// IPv6 Address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Prefix Length of IPv6 Address
	PrefixLength uint32 `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	// State of Address
	AddressState string `protobuf:"bytes,3,opt,name=address_state,json=addressState,proto3" json:"address_state,omitempty"`
	// Anycast address
	IsAnycast bool `protobuf:"varint,4,opt,name=is_anycast,json=isAnycast,proto3" json:"is_anycast,omitempty"`
	// Route-tag of the Address
	RouteTag             uint32   `protobuf:"varint,5,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6AddrNode) Reset()         { *m = Ipv6AddrNode{} }
func (m *Ipv6AddrNode) String() string { return proto.CompactTextString(m) }
func (*Ipv6AddrNode) ProtoMessage()    {}
func (*Ipv6AddrNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_brief_5bb3f7d8b3180cb3, []int{2}
}
func (m *Ipv6AddrNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6AddrNode.Unmarshal(m, b)
}
func (m *Ipv6AddrNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6AddrNode.Marshal(b, m, deterministic)
}
func (dst *Ipv6AddrNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6AddrNode.Merge(dst, src)
}
func (m *Ipv6AddrNode) XXX_Size() int {
	return xxx_messageInfo_Ipv6AddrNode.Size(m)
}
func (m *Ipv6AddrNode) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6AddrNode.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6AddrNode proto.InternalMessageInfo

func (m *Ipv6AddrNode) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6AddrNode) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6AddrNode) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *Ipv6AddrNode) GetIsAnycast() bool {
	if m != nil {
		return m.IsAnycast
	}
	return false
}

func (m *Ipv6AddrNode) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6IfBrief_KEYS)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_briefs.global_brief.ipv6_if_brief_KEYS")
	proto.RegisterType((*Ipv6IfBrief)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_briefs.global_brief.ipv6_if_brief")
	proto.RegisterType((*Ipv6AddrNode)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_briefs.global_brief.ipv6_addr_node")
}

func init() { proto.RegisterFile("ipv6_if_brief.proto", fileDescriptor_ipv6_if_brief_5bb3f7d8b3180cb3) }

var fileDescriptor_ipv6_if_brief_5bb3f7d8b3180cb3 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0x86, 0x95, 0x76, 0x77, 0x9b, 0xb8, 0x4d, 0xb5, 0xf2, 0x5e, 0xb2, 0x42, 0x95, 0xa2, 0x22,
	0xa4, 0x9e, 0x72, 0x68, 0x11, 0xf7, 0x1e, 0x38, 0x51, 0x71, 0x48, 0xb9, 0x70, 0x1a, 0xb9, 0x89,
	0x13, 0xac, 0xba, 0x76, 0x64, 0x9b, 0x50, 0x9e, 0x84, 0x3b, 0x4f, 0xc0, 0x8b, 0xf1, 0x0e, 0xc8,
	0x76, 0x03, 0xca, 0x0b, 0xc0, 0x65, 0xa4, 0xf9, 0xe6, 0xcf, 0xcc, 0xaf, 0xfc, 0x09, 0xfa, 0xc7,
	0x9a, 0xf6, 0x0a, 0x58, 0x05, 0x3b, 0xc5, 0x68, 0x95, 0x35, 0x4a, 0x1a, 0x89, 0xcb, 0x82, 0xe9,
	0x42, 0x02, 0x93, 0x1a, 0x8e, 0x0a, 0x9c, 0xe2, 0x40, 0x40, 0x36, 0x54, 0x65, 0xae, 0x11, 0xd4,
	0x3c, 0x49, 0xb5, 0xcf, 0x84, 0x2c, 0xa9, 0x76, 0x35, 0x63, 0xc2, 0x50, 0x55, 0x91, 0x82, 0x42,
	0x49, 0x0c, 0xc9, 0x5a, 0x55, 0x69, 0x5b, 0xb2, 0x9a, 0xcb, 0x1d, 0xe1, 0x7e, 0xbb, 0xee, 0x75,
	0x73, 0x8d, 0x70, 0xef, 0x38, 0xdc, 0x5c, 0xdf, 0x6f, 0xf1, 0x19, 0x8a, 0xec, 0x42, 0x10, 0xe4,
	0x40, 0x93, 0x20, 0x0d, 0x16, 0x51, 0x1e, 0x5a, 0x70, 0x4b, 0x0e, 0x14, 0xff, 0x47, 0x61, 0xab,
	0x2a, 0x3f, 0x1b, 0xb8, 0xd9, 0xa8, 0x55, 0x95, 0x1b, 0x5d, 0xa0, 0xe9, 0x97, 0x07, 0x27, 0x18,
	0x3a, 0x41, 0xfc, 0x49, 0xad, 0x6c, 0xfe, 0x3e, 0x40, 0x71, 0xef, 0x2a, 0x9e, 0x21, 0xc4, 0x99,
	0xa0, 0xa0, 0x0d, 0x31, 0x34, 0x59, 0xba, 0x87, 0x22, 0x4b, 0xb6, 0x16, 0xe0, 0x97, 0x00, 0x4d,
	0x48, 0x59, 0x2a, 0xaa, 0x35, 0x70, 0xa6, 0x4d, 0xb2, 0x4a, 0x87, 0x8b, 0xf1, 0xd2, 0x64, 0xdf,
	0xf1, 0x8e, 0xfc, 0x2a, 0x7b, 0x1e, 0xec, 0x86, 0x7c, 0x7c, 0x72, 0xb2, 0x61, 0xda, 0xe0, 0xd7,
	0x00, 0x61, 0xce, 0xc4, 0x1e, 0xb8, 0x2c, 0x08, 0x87, 0xd3, 0x28, 0xb9, 0x4c, 0x83, 0x1f, 0xf3,
	0xf7, 0xd7, 0xfa, 0xd9, 0x58, 0x3b, 0x6b, 0xef, 0x66, 0xfe, 0x16, 0xa0, 0x69, 0x5f, 0x84, 0x13,
	0x34, 0xea, 0xbc, 0xfa, 0x7c, 0xbb, 0x16, 0x9f, 0xa3, 0xb8, 0x51, 0xb4, 0x62, 0x47, 0xe0, 0x54,
	0xd4, 0xe6, 0xc1, 0x65, 0x1c, 0xe7, 0x13, 0x0f, 0x37, 0x8e, 0x59, 0x51, 0x97, 0x87, 0x8f, 0xcc,
	0xe7, 0xdc, 0x85, 0xe4, 0x53, 0x9b, 0x21, 0xc4, 0x34, 0x10, 0xf1, 0x5c, 0x10, 0x6d, 0x92, 0x5f,
	0x69, 0xb0, 0x08, 0xf3, 0x88, 0xe9, 0xb5, 0x07, 0xf6, 0x23, 0x53, 0xf2, 0xd1, 0x50, 0x30, 0xa4,
	0x4e, 0x7e, 0xbb, 0x23, 0xa1, 0x03, 0x77, 0xa4, 0xde, 0xfd, 0x71, 0x3f, 0xc1, 0xea, 0x23, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x34, 0xee, 0xf4, 0x1b, 0x03, 0x00, 0x00,
}
