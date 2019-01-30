// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_dhcpv6d_server_bindings_summary.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_summary

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

// DHCPv6 server bindings summary
type Ipv6Dhcpv6DServerBindingsSummary_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) Reset()         { *m = Ipv6Dhcpv6DServerBindingsSummary_KEYS{} }
func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBindingsSummary_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBindingsSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e546b90cab1ea30d, []int{0}
}

func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBindingsSummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv6Dhcpv6DServerBindingsSummary struct {
	// Total number of clients
	Clients uint32 `protobuf:"varint,50,opt,name=clients,proto3" json:"clients,omitempty"`
	// IANA server binding summary
	Iana *BagDhcpv6DServerBindingsSummary `protobuf:"bytes,51,opt,name=iana,proto3" json:"iana,omitempty"`
	// IAPD server binding summary
	Iapd                 *BagDhcpv6DServerBindingsSummary `protobuf:"bytes,52,opt,name=iapd,proto3" json:"iapd,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBindingsSummary) Reset()         { *m = Ipv6Dhcpv6DServerBindingsSummary{} }
func (m *Ipv6Dhcpv6DServerBindingsSummary) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBindingsSummary) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBindingsSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_e546b90cab1ea30d, []int{1}
}

func (m *Ipv6Dhcpv6DServerBindingsSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary.Size(m)
}
func (m *Ipv6Dhcpv6DServerBindingsSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBindingsSummary proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBindingsSummary) GetClients() uint32 {
	if m != nil {
		return m.Clients
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBindingsSummary) GetIana() *BagDhcpv6DServerBindingsSummary {
	if m != nil {
		return m.Iana
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerBindingsSummary) GetIapd() *BagDhcpv6DServerBindingsSummary {
	if m != nil {
		return m.Iapd
	}
	return nil
}

// DHCPv6 server bindings summary for IANA/IAPD
type BagDhcpv6DServerBindingsSummary struct {
	// Number of clients in init state
	InitializingClients uint32 `protobuf:"varint,1,opt,name=initializing_clients,json=initializingClients,proto3" json:"initializing_clients,omitempty"`
	// Number of clients waiting on DPM to validate subscriber
	DpmWaitingClients uint32 `protobuf:"varint,2,opt,name=dpm_waiting_clients,json=dpmWaitingClients,proto3" json:"dpm_waiting_clients,omitempty"`
	// Number of clients waiting on DAPS to assign/free addr/prefix
	DapsWaitingClients uint32 `protobuf:"varint,3,opt,name=daps_waiting_clients,json=dapsWaitingClients,proto3" json:"daps_waiting_clients,omitempty"`
	// Number of clients waiting for a request from the client
	RequestWaitingClients uint32 `protobuf:"varint,4,opt,name=request_waiting_clients,json=requestWaitingClients,proto3" json:"request_waiting_clients,omitempty"`
	// Number of clients waiting for iedge for the session
	IedgeWaitingClients uint32 `protobuf:"varint,5,opt,name=iedge_waiting_clients,json=iedgeWaitingClients,proto3" json:"iedge_waiting_clients,omitempty"`
	// Number of clients in waiting for RIB response
	RibWaitingClients uint32 `protobuf:"varint,6,opt,name=rib_waiting_clients,json=ribWaitingClients,proto3" json:"rib_waiting_clients,omitempty"`
	// Number of clients in bound state
	BoundClients         uint32   `protobuf:"varint,7,opt,name=bound_clients,json=boundClients,proto3" json:"bound_clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BagDhcpv6DServerBindingsSummary) Reset()         { *m = BagDhcpv6DServerBindingsSummary{} }
func (m *BagDhcpv6DServerBindingsSummary) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DServerBindingsSummary) ProtoMessage()    {}
func (*BagDhcpv6DServerBindingsSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_e546b90cab1ea30d, []int{2}
}

func (m *BagDhcpv6DServerBindingsSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagDhcpv6DServerBindingsSummary.Unmarshal(m, b)
}
func (m *BagDhcpv6DServerBindingsSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagDhcpv6DServerBindingsSummary.Marshal(b, m, deterministic)
}
func (m *BagDhcpv6DServerBindingsSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagDhcpv6DServerBindingsSummary.Merge(m, src)
}
func (m *BagDhcpv6DServerBindingsSummary) XXX_Size() int {
	return xxx_messageInfo_BagDhcpv6DServerBindingsSummary.Size(m)
}
func (m *BagDhcpv6DServerBindingsSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BagDhcpv6DServerBindingsSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BagDhcpv6DServerBindingsSummary proto.InternalMessageInfo

func (m *BagDhcpv6DServerBindingsSummary) GetInitializingClients() uint32 {
	if m != nil {
		return m.InitializingClients
	}
	return 0
}

func (m *BagDhcpv6DServerBindingsSummary) GetDpmWaitingClients() uint32 {
	if m != nil {
		return m.DpmWaitingClients
	}
	return 0
}

func (m *BagDhcpv6DServerBindingsSummary) GetDapsWaitingClients() uint32 {
	if m != nil {
		return m.DapsWaitingClients
	}
	return 0
}

func (m *BagDhcpv6DServerBindingsSummary) GetRequestWaitingClients() uint32 {
	if m != nil {
		return m.RequestWaitingClients
	}
	return 0
}

func (m *BagDhcpv6DServerBindingsSummary) GetIedgeWaitingClients() uint32 {
	if m != nil {
		return m.IedgeWaitingClients
	}
	return 0
}

func (m *BagDhcpv6DServerBindingsSummary) GetRibWaitingClients() uint32 {
	if m != nil {
		return m.RibWaitingClients
	}
	return 0
}

func (m *BagDhcpv6DServerBindingsSummary) GetBoundClients() uint32 {
	if m != nil {
		return m.BoundClients
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DServerBindingsSummary_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.summary.ipv6_dhcpv6d_server_bindings_summary_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DServerBindingsSummary)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.summary.ipv6_dhcpv6d_server_bindings_summary")
	proto.RegisterType((*BagDhcpv6DServerBindingsSummary)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.summary.bag_dhcpv6d_server_bindings_summary")
}

func init() {
	proto.RegisterFile("ipv6_dhcpv6d_server_bindings_summary.proto", fileDescriptor_e546b90cab1ea30d)
}

var fileDescriptor_e546b90cab1ea30d = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xbf, 0x7e, 0xad, 0xbd, 0xda, 0x85, 0x69, 0x8b, 0x01, 0x37, 0xa5, 0x75, 0x51,
	0x5d, 0x0c, 0xda, 0x4a, 0x5f, 0x40, 0x04, 0xa1, 0xe0, 0xa2, 0x2e, 0xc4, 0xd5, 0x30, 0xc9, 0x0c,
	0xf1, 0x42, 0x33, 0x19, 0x67, 0xd2, 0x56, 0x5d, 0xfb, 0x7c, 0xbe, 0x90, 0x1b, 0xc9, 0x64, 0x5a,
	0x4a, 0xb2, 0xb0, 0x0b, 0xc1, 0x4d, 0x60, 0xee, 0x39, 0xbf, 0x73, 0xff, 0x40, 0xe0, 0x02, 0xd5,
	0x6a, 0x4a, 0xf9, 0x73, 0xa4, 0x56, 0x53, 0x4e, 0x8d, 0xd0, 0x2b, 0xa1, 0x69, 0x88, 0x92, 0xa3,
	0x8c, 0x0d, 0x35, 0xcb, 0x24, 0x61, 0xfa, 0x8d, 0x28, 0x9d, 0x66, 0xa9, 0x3f, 0x8b, 0xd0, 0x44,
	0x29, 0xc5, 0xd4, 0xd0, 0x57, 0x4d, 0x2d, 0x28, 0xc5, 0x7a, 0x0b, 0xa7, 0x4a, 0x68, 0x52, 0x3c,
	0x88, 0x4c, 0xb9, 0x30, 0xf6, 0x4b, 0x8a, 0x4c, 0xe2, 0x32, 0x89, 0x8b, 0x1c, 0xdc, 0xc1, 0xf9,
	0x3e, 0xad, 0xe9, 0xec, 0xf6, 0xe9, 0xc1, 0x3f, 0x85, 0x56, 0x9e, 0x45, 0x25, 0x4b, 0x44, 0xe0,
	0xf5, 0xbd, 0x51, 0x6b, 0x7e, 0x90, 0x17, 0xee, 0x59, 0x22, 0x06, 0x9f, 0x35, 0x38, 0xdb, 0x27,
	0xca, 0x0f, 0xa0, 0x19, 0x2d, 0x50, 0xc8, 0xcc, 0x04, 0xe3, 0xbe, 0x37, 0x6a, 0xcf, 0x37, 0x4f,
	0xff, 0xc3, 0x83, 0x3a, 0x32, 0xc9, 0x82, 0x49, 0xdf, 0x1b, 0x1d, 0x8e, 0x15, 0xf9, 0xc5, 0x4d,
	0x49, 0xc8, 0xe2, 0x9f, 0x46, 0x9b, 0xdb, 0xee, 0x6e, 0x0c, 0xc5, 0x83, 0xeb, 0xbf, 0x1b, 0x43,
	0xf1, 0xc1, 0x57, 0x0d, 0x86, 0x7b, 0xb8, 0xfd, 0x2b, 0xe8, 0xa2, 0xc4, 0x0c, 0xd9, 0x02, 0xdf,
	0x51, 0xc6, 0x74, 0x73, 0x5c, 0xcf, 0x1e, 0xb7, 0xb3, 0xab, 0xdd, 0xb8, 0x43, 0x13, 0xe8, 0x70,
	0x95, 0xd0, 0x35, 0xc3, 0x6c, 0x97, 0xa8, 0x59, 0xe2, 0x98, 0xab, 0xe4, 0xb1, 0x50, 0x36, 0xfe,
	0x4b, 0xe8, 0x72, 0xa6, 0x4c, 0x05, 0xf8, 0x67, 0x01, 0x3f, 0xd7, 0x4a, 0xc4, 0x14, 0x4e, 0xb4,
	0x78, 0x59, 0x0a, 0x93, 0x55, 0xa0, 0xba, 0x85, 0x7a, 0x4e, 0x2e, 0x71, 0x63, 0xe8, 0xa1, 0xe0,
	0xb1, 0xa8, 0x50, 0xff, 0xdd, 0x36, 0xb9, 0x58, 0x62, 0x08, 0x74, 0x34, 0x86, 0x15, 0xa2, 0x51,
	0x6c, 0xa3, 0x31, 0x2c, 0xf9, 0x87, 0xd0, 0x0e, 0xd3, 0xa5, 0xe4, 0x5b, 0x67, 0xd3, 0x3a, 0x8f,
	0x6c, 0xd1, 0x99, 0xc2, 0x86, 0xfd, 0xd9, 0x26, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x33,
	0xfc, 0x27, 0x9a, 0x03, 0x00, 0x00,
}
