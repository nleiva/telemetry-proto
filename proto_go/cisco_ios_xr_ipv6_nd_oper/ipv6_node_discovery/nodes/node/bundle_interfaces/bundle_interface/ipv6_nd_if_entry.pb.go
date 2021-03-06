// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_nd_if_entry.proto

package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_bundle_interfaces_bundle_interface

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

// Detailed Info of ND IPv6 Interface entry
type Ipv6NdIfEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdIfEntry_KEYS) Reset()         { *m = Ipv6NdIfEntry_KEYS{} }
func (m *Ipv6NdIfEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdIfEntry_KEYS) ProtoMessage()    {}
func (*Ipv6NdIfEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b, []int{0}
}
func (m *Ipv6NdIfEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdIfEntry_KEYS.Unmarshal(m, b)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdIfEntry_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv6NdIfEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdIfEntry_KEYS.Merge(dst, src)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdIfEntry_KEYS.Size(m)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdIfEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdIfEntry_KEYS proto.InternalMessageInfo

func (m *Ipv6NdIfEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdIfEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6NdIfEntry struct {
	// Parent interface name
	ParentInterfaceName string `protobuf:"bytes,50,opt,name=parent_interface_name,json=parentInterfaceName,proto3" json:"parent_interface_name,omitempty"`
	// Interface type
	Iftype uint32 `protobuf:"varint,51,opt,name=iftype,proto3" json:"iftype,omitempty"`
	// MTU
	Mtu uint32 `protobuf:"varint,52,opt,name=mtu,proto3" json:"mtu,omitempty"`
	// etype
	Etype uint32 `protobuf:"varint,53,opt,name=etype,proto3" json:"etype,omitempty"`
	// vlan tag/id/ucv
	VlanTag uint32 `protobuf:"varint,54,opt,name=vlan_tag,json=vlanTag,proto3" json:"vlan_tag,omitempty"`
	// mac address size
	MacAddrSize uint32 `protobuf:"varint,55,opt,name=mac_addr_size,json=macAddrSize,proto3" json:"mac_addr_size,omitempty"`
	// mac address
	MacAddr string `protobuf:"bytes,56,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	// If true, interface is enabled
	IsInterfaceEnabled bool `protobuf:"varint,57,opt,name=is_interface_enabled,json=isInterfaceEnabled,proto3" json:"is_interface_enabled,omitempty"`
	// If true, IPv6 is enabled
	IsIpv6Enabled bool `protobuf:"varint,58,opt,name=is_ipv6_enabled,json=isIpv6Enabled,proto3" json:"is_ipv6_enabled,omitempty"`
	// If true, MPLS is enabled
	IsMplsEnabled bool `protobuf:"varint,59,opt,name=is_mpls_enabled,json=isMplsEnabled,proto3" json:"is_mpls_enabled,omitempty"`
	// ND interface parameters
	NdParameters *Ipv6NdIfParams `protobuf:"bytes,60,opt,name=nd_parameters,json=ndParameters,proto3" json:"nd_parameters,omitempty"`
	// List of ND global addresses
	GlobalAddressList []*Ipv6NdAddr `protobuf:"bytes,61,rep,name=global_address_list,json=globalAddressList,proto3" json:"global_address_list,omitempty"`
	// Link local address
	LocalAddress *Ipv6NdAddr `protobuf:"bytes,62,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	// List of member links
	MemberLinkList []uint32 `protobuf:"varint,63,rep,packed,name=member_link_list,json=memberLinkList,proto3" json:"member_link_list,omitempty"`
	// List of member nodes
	MemberNodeList       []*Ipv6NdGspnode `protobuf:"bytes,64,rep,name=member_node_list,json=memberNodeList,proto3" json:"member_node_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Ipv6NdIfEntry) Reset()         { *m = Ipv6NdIfEntry{} }
func (m *Ipv6NdIfEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdIfEntry) ProtoMessage()    {}
func (*Ipv6NdIfEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b, []int{1}
}
func (m *Ipv6NdIfEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdIfEntry.Unmarshal(m, b)
}
func (m *Ipv6NdIfEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdIfEntry.Marshal(b, m, deterministic)
}
func (dst *Ipv6NdIfEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdIfEntry.Merge(dst, src)
}
func (m *Ipv6NdIfEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdIfEntry.Size(m)
}
func (m *Ipv6NdIfEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdIfEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdIfEntry proto.InternalMessageInfo

func (m *Ipv6NdIfEntry) GetParentInterfaceName() string {
	if m != nil {
		return m.ParentInterfaceName
	}
	return ""
}

func (m *Ipv6NdIfEntry) GetIftype() uint32 {
	if m != nil {
		return m.Iftype
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetEtype() uint32 {
	if m != nil {
		return m.Etype
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetVlanTag() uint32 {
	if m != nil {
		return m.VlanTag
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetMacAddrSize() uint32 {
	if m != nil {
		return m.MacAddrSize
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *Ipv6NdIfEntry) GetIsInterfaceEnabled() bool {
	if m != nil {
		return m.IsInterfaceEnabled
	}
	return false
}

func (m *Ipv6NdIfEntry) GetIsIpv6Enabled() bool {
	if m != nil {
		return m.IsIpv6Enabled
	}
	return false
}

func (m *Ipv6NdIfEntry) GetIsMplsEnabled() bool {
	if m != nil {
		return m.IsMplsEnabled
	}
	return false
}

func (m *Ipv6NdIfEntry) GetNdParameters() *Ipv6NdIfParams {
	if m != nil {
		return m.NdParameters
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetGlobalAddressList() []*Ipv6NdAddr {
	if m != nil {
		return m.GlobalAddressList
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetLocalAddress() *Ipv6NdAddr {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetMemberLinkList() []uint32 {
	if m != nil {
		return m.MemberLinkList
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetMemberNodeList() []*Ipv6NdGspnode {
	if m != nil {
		return m.MemberNodeList
	}
	return nil
}

// Detailed Info of ND IPv6 Interface
type Ipv6NdIfParams struct {
	// If true, DAD (D.. A.. D..) is enabled otherwise it is disabled
	IsDadEnabled bool `protobuf:"varint,1,opt,name=is_dad_enabled,json=isDadEnabled,proto3" json:"is_dad_enabled,omitempty"`
	// DAD attempt count
	DadAttempts uint32 `protobuf:"varint,2,opt,name=dad_attempts,json=dadAttempts,proto3" json:"dad_attempts,omitempty"`
	// ICMP redirect flag
	IsIcmPv6Redirect bool `protobuf:"varint,3,opt,name=is_icm_pv6_redirect,json=isIcmPv6Redirect,proto3" json:"is_icm_pv6_redirect,omitempty"`
	// Flag used for utilising DHCP
	IsDhcpManaged bool `protobuf:"varint,4,opt,name=is_dhcp_managed,json=isDhcpManaged,proto3" json:"is_dhcp_managed,omitempty"`
	// Flag used to manage routable address
	IsRouteAddressManaged bool `protobuf:"varint,5,opt,name=is_route_address_managed,json=isRouteAddressManaged,proto3" json:"is_route_address_managed,omitempty"`
	// Suppress flag
	IsSuppressed bool `protobuf:"varint,6,opt,name=is_suppressed,json=isSuppressed,proto3" json:"is_suppressed,omitempty"`
	// ND retransmit interval in msec
	NdRetransmitInterval uint32 `protobuf:"varint,7,opt,name=nd_retransmit_interval,json=ndRetransmitInterval,proto3" json:"nd_retransmit_interval,omitempty"`
	// ND router advertisement minimum transmit interval in sec
	NdMinTransmitInterval uint32 `protobuf:"varint,8,opt,name=nd_min_transmit_interval,json=ndMinTransmitInterval,proto3" json:"nd_min_transmit_interval,omitempty"`
	// ND router advertisement maximum transmit interval in sec
	NdMaxTransmitInterval uint32 `protobuf:"varint,9,opt,name=nd_max_transmit_interval,json=ndMaxTransmitInterval,proto3" json:"nd_max_transmit_interval,omitempty"`
	// ND router advertisement life time in sec
	NdAdvertisementLifetime uint32 `protobuf:"varint,10,opt,name=nd_advertisement_lifetime,json=ndAdvertisementLifetime,proto3" json:"nd_advertisement_lifetime,omitempty"`
	// Time to reach ND in msec
	NdReachableTime uint32 `protobuf:"varint,11,opt,name=nd_reachable_time,json=ndReachableTime,proto3" json:"nd_reachable_time,omitempty"`
	// Completed adjacency limit per interface
	NdCacheLimit uint32 `protobuf:"varint,12,opt,name=nd_cache_limit,json=ndCacheLimit,proto3" json:"nd_cache_limit,omitempty"`
	// Completed PROTO entry Count
	CompleteProtocolCount uint32 `protobuf:"varint,13,opt,name=complete_protocol_count,json=completeProtocolCount,proto3" json:"complete_protocol_count,omitempty"`
	// Completed GLEAN entry count
	CompleteGleanCount uint32 `protobuf:"varint,14,opt,name=complete_glean_count,json=completeGleanCount,proto3" json:"complete_glean_count,omitempty"`
	// Incomplete PROTO entry count
	IncompleteProtocolCount uint32 `protobuf:"varint,15,opt,name=incomplete_protocol_count,json=incompleteProtocolCount,proto3" json:"incomplete_protocol_count,omitempty"`
	// Incomplete GLEAN entry count
	IncompleteGleanCount uint32 `protobuf:"varint,16,opt,name=incomplete_glean_count,json=incompleteGleanCount,proto3" json:"incomplete_glean_count,omitempty"`
	// Dropped PROTO entry request count
	DroppedProtocolReqCount uint32 `protobuf:"varint,17,opt,name=dropped_protocol_req_count,json=droppedProtocolReqCount,proto3" json:"dropped_protocol_req_count,omitempty"`
	// Dropped GLEAN entry lequest count
	DroppedGleanReqCount uint32   `protobuf:"varint,18,opt,name=dropped_glean_req_count,json=droppedGleanReqCount,proto3" json:"dropped_glean_req_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdIfParams) Reset()         { *m = Ipv6NdIfParams{} }
func (m *Ipv6NdIfParams) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdIfParams) ProtoMessage()    {}
func (*Ipv6NdIfParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b, []int{2}
}
func (m *Ipv6NdIfParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdIfParams.Unmarshal(m, b)
}
func (m *Ipv6NdIfParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdIfParams.Marshal(b, m, deterministic)
}
func (dst *Ipv6NdIfParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdIfParams.Merge(dst, src)
}
func (m *Ipv6NdIfParams) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdIfParams.Size(m)
}
func (m *Ipv6NdIfParams) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdIfParams.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdIfParams proto.InternalMessageInfo

func (m *Ipv6NdIfParams) GetIsDadEnabled() bool {
	if m != nil {
		return m.IsDadEnabled
	}
	return false
}

func (m *Ipv6NdIfParams) GetDadAttempts() uint32 {
	if m != nil {
		return m.DadAttempts
	}
	return 0
}

func (m *Ipv6NdIfParams) GetIsIcmPv6Redirect() bool {
	if m != nil {
		return m.IsIcmPv6Redirect
	}
	return false
}

func (m *Ipv6NdIfParams) GetIsDhcpManaged() bool {
	if m != nil {
		return m.IsDhcpManaged
	}
	return false
}

func (m *Ipv6NdIfParams) GetIsRouteAddressManaged() bool {
	if m != nil {
		return m.IsRouteAddressManaged
	}
	return false
}

func (m *Ipv6NdIfParams) GetIsSuppressed() bool {
	if m != nil {
		return m.IsSuppressed
	}
	return false
}

func (m *Ipv6NdIfParams) GetNdRetransmitInterval() uint32 {
	if m != nil {
		return m.NdRetransmitInterval
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdMinTransmitInterval() uint32 {
	if m != nil {
		return m.NdMinTransmitInterval
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdMaxTransmitInterval() uint32 {
	if m != nil {
		return m.NdMaxTransmitInterval
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdAdvertisementLifetime() uint32 {
	if m != nil {
		return m.NdAdvertisementLifetime
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdReachableTime() uint32 {
	if m != nil {
		return m.NdReachableTime
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdCacheLimit() uint32 {
	if m != nil {
		return m.NdCacheLimit
	}
	return 0
}

func (m *Ipv6NdIfParams) GetCompleteProtocolCount() uint32 {
	if m != nil {
		return m.CompleteProtocolCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetCompleteGleanCount() uint32 {
	if m != nil {
		return m.CompleteGleanCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetIncompleteProtocolCount() uint32 {
	if m != nil {
		return m.IncompleteProtocolCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetIncompleteGleanCount() uint32 {
	if m != nil {
		return m.IncompleteGleanCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetDroppedProtocolReqCount() uint32 {
	if m != nil {
		return m.DroppedProtocolReqCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetDroppedGleanReqCount() uint32 {
	if m != nil {
		return m.DroppedGleanReqCount
	}
	return 0
}

// List of IPv6 ND Addresses
type Ipv6NdAddr struct {
	// IPv6 address
	Ipv6Address          string   `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdAddr) Reset()         { *m = Ipv6NdAddr{} }
func (m *Ipv6NdAddr) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdAddr) ProtoMessage()    {}
func (*Ipv6NdAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b, []int{3}
}
func (m *Ipv6NdAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdAddr.Unmarshal(m, b)
}
func (m *Ipv6NdAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdAddr.Marshal(b, m, deterministic)
}
func (dst *Ipv6NdAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdAddr.Merge(dst, src)
}
func (m *Ipv6NdAddr) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdAddr.Size(m)
}
func (m *Ipv6NdAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdAddr.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdAddr proto.InternalMessageInfo

func (m *Ipv6NdAddr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

// GSP node info
type Ipv6NdGspnode struct {
	// Node Name
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Number of links on the node
	TotalLinks           uint32   `protobuf:"varint,2,opt,name=total_links,json=totalLinks,proto3" json:"total_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdGspnode) Reset()         { *m = Ipv6NdGspnode{} }
func (m *Ipv6NdGspnode) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdGspnode) ProtoMessage()    {}
func (*Ipv6NdGspnode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b, []int{4}
}
func (m *Ipv6NdGspnode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdGspnode.Unmarshal(m, b)
}
func (m *Ipv6NdGspnode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdGspnode.Marshal(b, m, deterministic)
}
func (dst *Ipv6NdGspnode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdGspnode.Merge(dst, src)
}
func (m *Ipv6NdGspnode) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdGspnode.Size(m)
}
func (m *Ipv6NdGspnode) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdGspnode.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdGspnode proto.InternalMessageInfo

func (m *Ipv6NdGspnode) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdGspnode) GetTotalLinks() uint32 {
	if m != nil {
		return m.TotalLinks
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6NdIfEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_if_entry_KEYS")
	proto.RegisterType((*Ipv6NdIfEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_if_entry")
	proto.RegisterType((*Ipv6NdIfParams)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_if_params")
	proto.RegisterType((*Ipv6NdAddr)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_addr")
	proto.RegisterType((*Ipv6NdGspnode)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_gspnode")
}

func init() {
	proto.RegisterFile("ipv6_nd_if_entry.proto", fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b)
}

var fileDescriptor_ipv6_nd_if_entry_21927ea4d3fda74b = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0x12, 0x9a, 0xcb, 0xb1, 0xd7, 0x71, 0x26, 0x97, 0x6e, 0xe0, 0x01, 0x63, 0x0a, 0xb2,
	0x90, 0xb0, 0x20, 0x4d, 0x5d, 0x48, 0xb9, 0x45, 0x6d, 0x85, 0x22, 0x92, 0x12, 0x6d, 0xf2, 0x82,
	0xfa, 0x30, 0x1a, 0xef, 0x9c, 0x38, 0xa3, 0xee, 0xce, 0x6e, 0x67, 0xc6, 0x56, 0xd2, 0x1f, 0xc1,
	0x0b, 0x12, 0x88, 0x7f, 0xc8, 0xcf, 0x40, 0x73, 0xf6, 0x62, 0x37, 0xa1, 0xbc, 0x91, 0x97, 0x28,
	0xf3, 0x5d, 0xe6, 0x9c, 0x33, 0x73, 0xe6, 0xac, 0x61, 0x47, 0x15, 0xb3, 0x11, 0xd7, 0x92, 0xab,
	0x0b, 0x8e, 0xda, 0x99, 0xeb, 0x61, 0x61, 0x72, 0x97, 0xb3, 0x97, 0x89, 0xb2, 0x49, 0xce, 0x55,
	0x6e, 0xf9, 0x95, 0xe1, 0xb5, 0x28, 0x2f, 0xd0, 0x0c, 0xcb, 0x45, 0x2e, 0x91, 0x4b, 0xaf, 0x99,
	0xa1, 0xb9, 0x1e, 0xfa, 0xa5, 0xa5, 0xbf, 0xc3, 0xf1, 0x54, 0xcb, 0x14, 0xb9, 0xd2, 0x0e, 0xcd,
	0x85, 0x48, 0xd0, 0xde, 0x42, 0xfa, 0x2f, 0x61, 0xfb, 0x66, 0x58, 0xfe, 0xf3, 0xf3, 0x5f, 0xcf,
	0xd8, 0x87, 0xb0, 0x46, 0x1b, 0x6b, 0x91, 0x61, 0x14, 0xf4, 0x82, 0xc1, 0x5a, 0xbc, 0xea, 0x81,
	0x17, 0x22, 0x43, 0xf6, 0x29, 0x74, 0x9a, 0x2d, 0x4a, 0xc5, 0x7b, 0xa4, 0x08, 0x1b, 0xd4, 0xcb,
	0xfa, 0x7f, 0xaf, 0x40, 0xf7, 0xe6, 0xee, 0x6c, 0x0f, 0xb6, 0x0b, 0x61, 0x50, 0x3b, 0x7e, 0x63,
	0x8b, 0x3d, 0xda, 0x62, 0xb3, 0x24, 0x8f, 0x16, 0x37, 0x62, 0x3b, 0xb0, 0xac, 0x2e, 0xdc, 0x75,
	0x81, 0xd1, 0xc3, 0x5e, 0x30, 0x08, 0xe3, 0x6a, 0xc5, 0xba, 0xb0, 0x94, 0xb9, 0x69, 0xb4, 0x4f,
	0xa0, 0xff, 0x97, 0x6d, 0xc1, 0x3d, 0x24, 0xe1, 0x23, 0xc2, 0xca, 0x05, 0xdb, 0x85, 0xd5, 0x59,
	0x2a, 0x34, 0x77, 0x62, 0x12, 0x8d, 0x88, 0x58, 0xf1, 0xeb, 0x73, 0x31, 0x61, 0x7d, 0x08, 0x33,
	0x91, 0x70, 0x21, 0xa5, 0xe1, 0x56, 0xbd, 0xc1, 0xe8, 0x31, 0xf1, 0xad, 0x4c, 0x24, 0x87, 0x52,
	0x9a, 0x33, 0xf5, 0x86, 0xec, 0xb5, 0x26, 0xfa, 0x9a, 0xb2, 0x5c, 0xa9, 0x68, 0xf6, 0x25, 0x6c,
	0x29, 0xbb, 0x50, 0x09, 0x6a, 0x31, 0x4e, 0x51, 0x46, 0xdf, 0xf4, 0x82, 0xc1, 0x6a, 0xcc, 0x94,
	0x6d, 0x0a, 0x79, 0x5e, 0x32, 0xec, 0x33, 0x58, 0xf7, 0x0e, 0x7f, 0x2c, 0xb5, 0xf8, 0x80, 0xc4,
	0xa1, 0xb2, 0x47, 0xc5, 0x6c, 0xf4, 0xb6, 0x2e, 0x2b, 0x52, 0xdb, 0xe8, 0x9e, 0xd4, 0xba, 0x93,
	0x22, 0xb5, 0xb5, 0xee, 0xf7, 0x00, 0x42, 0x2d, 0x79, 0x21, 0x8c, 0xc8, 0xd0, 0xa1, 0xb1, 0xd1,
	0xb7, 0xbd, 0x60, 0xd0, 0xda, 0xd3, 0xc3, 0xff, 0xb1, 0x6f, 0x86, 0x0b, 0xd7, 0x4a, 0x91, 0x6d,
	0xdc, 0xd6, 0xf2, 0xb4, 0xc9, 0x81, 0xfd, 0x15, 0xc0, 0xe6, 0x24, 0xcd, 0xc7, 0x22, 0xa5, 0x63,
	0x43, 0x6b, 0x79, 0xaa, 0xac, 0x8b, 0xbe, 0xeb, 0x2d, 0x0d, 0x5a, 0x7b, 0xea, 0x4e, 0x72, 0xf3,
	0x81, 0xe3, 0x8d, 0x32, 0x8b, 0xc3, 0x32, 0x89, 0x63, 0x65, 0x1d, 0xfb, 0x2d, 0x80, 0x30, 0xcd,
	0x93, 0x79, 0x6a, 0xd1, 0xf7, 0x74, 0x62, 0x77, 0x98, 0x55, 0x9b, 0xe2, 0x57, 0x49, 0xb1, 0x01,
	0x74, 0x33, 0xcc, 0xc6, 0x68, 0x78, 0xaa, 0xf4, 0xab, 0xf2, 0xa0, 0x7e, 0xe8, 0x2d, 0x0d, 0xc2,
	0xb8, 0x53, 0xe2, 0xc7, 0x4a, 0xbf, 0xa2, 0xd4, 0xff, 0x08, 0x1a, 0x29, 0x25, 0x43, 0xd2, 0x1f,
	0xe9, 0x4c, 0xd3, 0x3b, 0xc9, 0x7e, 0x62, 0x0b, 0x6f, 0xa9, 0x13, 0x7b, 0x91, 0x4b, 0xf4, 0x89,
	0xf5, 0xff, 0x5c, 0x81, 0x8d, 0x5b, 0x3d, 0xc1, 0x1e, 0x40, 0x47, 0x59, 0x2e, 0x85, 0x6c, 0x5a,
	0x38, 0xa0, 0x16, 0x6e, 0x2b, 0xfb, 0x4c, 0xc8, 0xba, 0x83, 0x3f, 0x86, 0xb6, 0x97, 0x08, 0xe7,
	0x30, 0x2b, 0x9c, 0xa5, 0x59, 0x12, 0xc6, 0x2d, 0x29, 0xe4, 0x61, 0x05, 0xb1, 0x2f, 0x60, 0xd3,
	0x3f, 0x9a, 0x24, 0xe3, 0x3e, 0x86, 0x41, 0xa9, 0x0c, 0x26, 0x2e, 0x5a, 0xa2, 0xdd, 0xba, 0xca,
	0x1e, 0x25, 0xd9, 0xe9, 0x6c, 0x14, 0x57, 0x78, 0xf5, 0x76, 0xe4, 0x65, 0x52, 0xf0, 0x4c, 0x68,
	0x31, 0x41, 0x19, 0xbd, 0x5f, 0xbf, 0x9d, 0x67, 0x97, 0x49, 0x71, 0x52, 0x82, 0xec, 0x31, 0x44,
	0xca, 0x72, 0x93, 0x4f, 0x1d, 0x36, 0x6d, 0x5a, 0x1b, 0xee, 0x91, 0x61, 0x5b, 0xd9, 0xd8, 0xd3,
	0xd5, 0x55, 0xd5, 0xc6, 0x4f, 0x20, 0x54, 0x96, 0xdb, 0x69, 0x51, 0x78, 0x14, 0x65, 0xb4, 0x5c,
	0xd7, 0x75, 0xd6, 0x60, 0x6c, 0x1f, 0x76, 0xb4, 0xe4, 0x06, 0x9d, 0x11, 0xda, 0x66, 0xaa, 0x1a,
	0x78, 0x33, 0x91, 0x46, 0x2b, 0x54, 0xe1, 0x96, 0x96, 0x71, 0x43, 0x1e, 0x55, 0x9c, 0xcf, 0x49,
	0x4b, 0x9e, 0x29, 0xcd, 0x6f, 0xfb, 0x56, 0xc9, 0xb7, 0xad, 0xe5, 0x89, 0xd2, 0xe7, 0xef, 0x30,
	0x8a, 0xab, 0x7f, 0x31, 0xae, 0x35, 0x46, 0x71, 0x75, 0xcb, 0x78, 0x00, 0xbb, 0xd4, 0x97, 0x33,
	0x34, 0x4e, 0x59, 0xcc, 0xfc, 0x6c, 0x4e, 0xd5, 0x05, 0x3a, 0x95, 0x61, 0x04, 0xe4, 0xbc, 0xaf,
	0xe5, 0xe1, 0x22, 0x7f, 0x5c, 0xd1, 0xec, 0x73, 0xd8, 0xa0, 0x1a, 0x45, 0x72, 0xe9, 0x2f, 0x93,
	0x93, 0xa7, 0x45, 0x9e, 0x75, 0x5f, 0x5e, 0x85, 0x9f, 0x7b, 0xed, 0x03, 0xe8, 0x68, 0xc9, 0x13,
	0x91, 0x5c, 0xfa, 0xc6, 0xcd, 0x94, 0x8b, 0xda, 0x24, 0x6c, 0x6b, 0xf9, 0xd4, 0x83, 0xc7, 0x1e,
	0x63, 0x23, 0xb8, 0x9f, 0xe4, 0x59, 0x91, 0xa2, 0x43, 0x4e, 0x1f, 0xc0, 0x24, 0x4f, 0x79, 0x92,
	0x4f, 0xb5, 0x8b, 0xc2, 0xb2, 0x8a, 0x9a, 0x3e, 0xad, 0xd8, 0xa7, 0x9e, 0xf4, 0x93, 0xb8, 0xf1,
	0x4d, 0x52, 0x14, 0xba, 0x32, 0x75, 0xc8, 0xc4, 0x6a, 0xee, 0x27, 0x4f, 0x95, 0x8e, 0x03, 0xd8,
	0x55, 0xfa, 0x5d, 0xb1, 0xd6, 0xcb, 0xba, 0xe7, 0x82, 0xb7, 0xa3, 0xed, 0xc3, 0xce, 0x82, 0x77,
	0x31, 0x5e, 0xb7, 0xbc, 0xdb, 0x39, 0xbb, 0x10, 0xf1, 0x09, 0x7c, 0x20, 0x4d, 0x5e, 0x14, 0x28,
	0xe7, 0xe1, 0x0c, 0xbe, 0xae, 0x9c, 0x1b, 0x65, 0xc8, 0x4a, 0x51, 0xc7, 0x8b, 0xf1, 0x75, 0x69,
	0x7e, 0x04, 0x35, 0x55, 0xc5, 0x9b, 0x3b, 0x59, 0x19, 0xb3, 0xa2, 0x29, 0x60, 0x6d, 0xeb, 0x7f,
	0x05, 0xed, 0xc5, 0xd1, 0xe3, 0x5f, 0x1b, 0xad, 0xeb, 0xd9, 0x57, 0x7e, 0xdb, 0x5b, 0x1e, 0xab,
	0x9a, 0xbc, 0xff, 0x0b, 0xac, 0xdf, 0x78, 0xef, 0xff, 0xfd, 0x73, 0xe0, 0x23, 0x68, 0xb9, 0xdc,
	0x89, 0x94, 0xc6, 0x57, 0xfd, 0x7e, 0x81, 0x20, 0x3f, 0xb9, 0xec, 0x78, 0x99, 0xaa, 0x7d, 0xf8,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x28, 0x10, 0x30, 0xe3, 0x08, 0x00, 0x00,
}
