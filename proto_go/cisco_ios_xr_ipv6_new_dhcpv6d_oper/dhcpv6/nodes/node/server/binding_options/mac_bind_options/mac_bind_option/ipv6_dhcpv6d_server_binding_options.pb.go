// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_dhcpv6d_server_binding_options.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option

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

// DHCPv6 server binding inserted option values
type Ipv6Dhcpv6DServerBindingOptions_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	MacAddress           string   `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) Reset()         { *m = Ipv6Dhcpv6DServerBindingOptions_KEYS{} }
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBindingOptions_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBindingOptions_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0286bae149069f3, []int{0}
}

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

type Ipv6Dhcpv6DServerBindingOptions struct {
	// Client MAC address
	MacAddress string `protobuf:"bytes,50,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// Client DUID
	Duid string `protobuf:"bytes,51,opt,name=duid,proto3" json:"duid,omitempty"`
	// DNS address count
	DnsCount uint32 `protobuf:"varint,52,opt,name=dns_count,json=dnsCount,proto3" json:"dns_count,omitempty"`
	// DNS addresses
	DnsAddress []*IPV6AddressType `protobuf:"bytes,53,rep,name=dns_address,json=dnsAddress,proto3" json:"dns_address,omitempty"`
	// Client Option 17 value
	Opt17                string   `protobuf:"bytes,54,opt,name=opt17,proto3" json:"opt17,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBindingOptions) Reset()         { *m = Ipv6Dhcpv6DServerBindingOptions{} }
func (m *Ipv6Dhcpv6DServerBindingOptions) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBindingOptions) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBindingOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0286bae149069f3, []int{1}
}

func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Size(m)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBindingOptions) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetDnsCount() uint32 {
	if m != nil {
		return m.DnsCount
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetDnsAddress() []*IPV6AddressType {
	if m != nil {
		return m.DnsAddress
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetOpt17() string {
	if m != nil {
		return m.Opt17
	}
	return ""
}

// IPV6 Address type
type IPV6AddressType struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPV6AddressType) Reset()         { *m = IPV6AddressType{} }
func (m *IPV6AddressType) String() string { return proto.CompactTextString(m) }
func (*IPV6AddressType) ProtoMessage()    {}
func (*IPV6AddressType) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0286bae149069f3, []int{2}
}

func (m *IPV6AddressType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPV6AddressType.Unmarshal(m, b)
}
func (m *IPV6AddressType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPV6AddressType.Marshal(b, m, deterministic)
}
func (m *IPV6AddressType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPV6AddressType.Merge(m, src)
}
func (m *IPV6AddressType) XXX_Size() int {
	return xxx_messageInfo_IPV6AddressType.Size(m)
}
func (m *IPV6AddressType) XXX_DiscardUnknown() {
	xxx_messageInfo_IPV6AddressType.DiscardUnknown(m)
}

var xxx_messageInfo_IPV6AddressType proto.InternalMessageInfo

func (m *IPV6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DServerBindingOptions_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DServerBindingOptions)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options")
	proto.RegisterType((*IPV6AddressType)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.IPV6AddressType")
}

func init() {
	proto.RegisterFile("ipv6_dhcpv6d_server_binding_options.proto", fileDescriptor_c0286bae149069f3)
}

var fileDescriptor_c0286bae149069f3 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0xf4, 0x30,
	0x10, 0xc7, 0xe9, 0x7e, 0x9f, 0xb2, 0x3b, 0x8b, 0x08, 0xc1, 0x43, 0xc1, 0x83, 0x4b, 0x3d, 0x58,
	0x2f, 0x01, 0x77, 0xb5, 0x9e, 0x45, 0x3c, 0x88, 0x20, 0x52, 0x45, 0xf0, 0x14, 0xb2, 0x99, 0xe0,
	0x06, 0xb6, 0x49, 0x48, 0xda, 0xaa, 0x07, 0x5f, 0xc5, 0x47, 0xf0, 0x19, 0x25, 0xa9, 0xf5, 0x50,
	0x10, 0xf6, 0xe4, 0x25, 0xe4, 0x3f, 0x93, 0xf9, 0xcd, 0xfc, 0x87, 0xc0, 0xb1, 0xb2, 0x6d, 0xc1,
	0x70, 0x25, 0x6c, 0x5b, 0x20, 0xf3, 0xd2, 0xb5, 0xd2, 0xb1, 0xa5, 0xd2, 0xa8, 0xf4, 0x33, 0x33,
	0xb6, 0x56, 0x46, 0x7b, 0x6a, 0x9d, 0xa9, 0x0d, 0x59, 0x0b, 0xe5, 0x85, 0x61, 0xca, 0x78, 0xf6,
	0xea, 0x58, 0xac, 0xd3, 0xf2, 0xe5, 0xa7, 0xd6, 0x58, 0xe9, 0x68, 0x27, 0xa8, 0x36, 0x28, 0x7d,
	0x3c, 0x69, 0x87, 0xa4, 0x43, 0x64, 0xc5, 0x45, 0x6c, 0xf3, 0x5b, 0x20, 0x5b, 0x41, 0xbe, 0xc1,
	0x68, 0xec, 0xe6, 0xea, 0xe9, 0x9e, 0xec, 0xc3, 0x24, 0xf4, 0x62, 0x9a, 0x57, 0x32, 0x4d, 0x66,
	0x49, 0x3e, 0x29, 0xc7, 0x21, 0x70, 0xcb, 0x2b, 0x49, 0x0e, 0x60, 0x1a, 0xd8, 0x1c, 0xd1, 0x49,
	0xef, 0xd3, 0x51, 0x4c, 0x43, 0xc5, 0xc5, 0x45, 0x17, 0xc9, 0x3e, 0x47, 0x70, 0xb8, 0x41, 0xab,
	0x21, 0x68, 0x3e, 0x04, 0x11, 0x02, 0xff, 0xb1, 0x51, 0x98, 0x2e, 0x62, 0x26, 0xde, 0xc3, 0x68,
	0xa8, 0x3d, 0x13, 0xa6, 0xd1, 0x75, 0x7a, 0x3a, 0x4b, 0xf2, 0x9d, 0x72, 0x8c, 0xda, 0x5f, 0x06,
	0x4d, 0x3e, 0x12, 0x98, 0x86, 0x6c, 0x8f, 0x3c, 0x9b, 0xfd, 0xcb, 0xa7, 0xf3, 0x77, 0xfa, 0x97,
	0x8b, 0xa6, 0xd7, 0x77, 0x8f, 0xc5, 0xb7, 0x83, 0x87, 0x37, 0x2b, 0x4b, 0x40, 0xed, 0x7b, 0x47,
	0x7b, 0xb0, 0x65, 0x6c, 0x7d, 0x72, 0x9e, 0x16, 0xd1, 0x52, 0x27, 0xb2, 0x23, 0xd8, 0x1d, 0x14,
	0x85, 0x87, 0x2d, 0x5f, 0x37, 0xfd, 0xf6, 0x3b, 0xb1, 0xdc, 0x8e, 0x1f, 0x67, 0xf1, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0xb1, 0x2a, 0x51, 0x65, 0x02, 0x00, 0x00,
}
