// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_dhcpv6d_relay_bindings_summary.proto

/*
Package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_summary is a generated protocol buffer package.

It is generated from these files:
	ipv6_dhcpv6d_relay_bindings_summary.proto

It has these top-level messages:
	Ipv6Dhcpv6DRelayBindingsSummary_KEYS
	Ipv6Dhcpv6DRelayBindingsSummary
*/
package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_summary

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

// DHCPv6 relay binding summary
type Ipv6Dhcpv6DRelayBindingsSummary_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *Ipv6Dhcpv6DRelayBindingsSummary_KEYS) Reset()         { *m = Ipv6Dhcpv6DRelayBindingsSummary_KEYS{} }
func (m *Ipv6Dhcpv6DRelayBindingsSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayBindingsSummary_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DRelayBindingsSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *Ipv6Dhcpv6DRelayBindingsSummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv6Dhcpv6DRelayBindingsSummary struct {
	// Total number of clients
	Clients uint32 `protobuf:"varint,50,opt,name=clients" json:"clients,omitempty"`
}

func (m *Ipv6Dhcpv6DRelayBindingsSummary) Reset()                    { *m = Ipv6Dhcpv6DRelayBindingsSummary{} }
func (m *Ipv6Dhcpv6DRelayBindingsSummary) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayBindingsSummary) ProtoMessage()               {}
func (*Ipv6Dhcpv6DRelayBindingsSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6Dhcpv6DRelayBindingsSummary) GetClients() uint32 {
	if m != nil {
		return m.Clients
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DRelayBindingsSummary_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.binding.summary.ipv6_dhcpv6d_relay_bindings_summary_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DRelayBindingsSummary)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.binding.summary.ipv6_dhcpv6d_relay_bindings_summary")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_relay_bindings_summary.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0x2c, 0x28, 0x33,
	0x8b, 0x4f, 0xc9, 0x48, 0x2e, 0x28, 0x33, 0x4b, 0x89, 0x2f, 0x4a, 0xcd, 0x49, 0xac, 0x8c, 0x4f,
	0xca, 0xcc, 0x4b, 0xc9, 0xcc, 0x4b, 0x2f, 0x8e, 0x2f, 0x2e, 0xcd, 0xcd, 0x4d, 0x2c, 0xaa, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xf2, 0x4a, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f,
	0x8e, 0xaf, 0x28, 0x8a, 0x07, 0xeb, 0xcb, 0x4b, 0x2d, 0x87, 0xeb, 0xcd, 0x2f, 0x48, 0x2d, 0xd2,
	0x83, 0x70, 0xf4, 0xf2, 0xf2, 0x53, 0x52, 0x8b, 0xc1, 0xa4, 0x1e, 0xd8, 0x48, 0x3d, 0xa8, 0x91,
	0x7a, 0x50, 0x13, 0x95, 0xdc, 0xb9, 0x34, 0x88, 0xb0, 0x38, 0xde, 0xdb, 0x35, 0x32, 0x58, 0x48,
	0x9a, 0x8b, 0x13, 0x64, 0x52, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x07, 0x48, 0xc0, 0x2f, 0x31, 0x37, 0x55, 0xc9, 0x9e, 0x4b, 0x99, 0x08, 0x83, 0x84, 0x24,
	0xb8, 0xd8, 0x93, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x8a, 0x25, 0x8c, 0x14, 0x18, 0x35, 0x78, 0x83,
	0x60, 0xdc, 0x24, 0x36, 0xb0, 0xe7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x6a, 0x33,
	0x2b, 0x09, 0x01, 0x00, 0x00,
}
