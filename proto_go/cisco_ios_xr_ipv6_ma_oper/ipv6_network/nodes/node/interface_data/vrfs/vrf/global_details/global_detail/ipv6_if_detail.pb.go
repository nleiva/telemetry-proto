// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_if_detail.proto

package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_vrfs_vrf_global_details_global_detail

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

// Detailed Info of IPv6 Interface
type Ipv6IfDetail_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6IfDetail_KEYS) Reset()         { *m = Ipv6IfDetail_KEYS{} }
func (m *Ipv6IfDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6IfDetail_KEYS) ProtoMessage()    {}
func (*Ipv6IfDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{0}
}
func (m *Ipv6IfDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6IfDetail_KEYS.Unmarshal(m, b)
}
func (m *Ipv6IfDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6IfDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv6IfDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6IfDetail_KEYS.Merge(dst, src)
}
func (m *Ipv6IfDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6IfDetail_KEYS.Size(m)
}
func (m *Ipv6IfDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6IfDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6IfDetail_KEYS proto.InternalMessageInfo

func (m *Ipv6IfDetail_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6IfDetail_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6IfDetail_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6IfDetail struct {
	// State of Interface Line
	LineState string `protobuf:"bytes,50,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	// IPv6 Multicast Group
	MulticastGroups []*Ipv6McastGroup `protobuf:"bytes,51,rep,name=multicast_groups,json=multicastGroups,proto3" json:"multicast_groups,omitempty"`
	// IPv6 MTU
	Mtu uint32 `protobuf:"varint,52,opt,name=mtu,proto3" json:"mtu,omitempty"`
	// IPv6 Operation State
	OperationState string `protobuf:"bytes,53,opt,name=operation_state,json=operationState,proto3" json:"operation_state,omitempty"`
	// Address List
	AddressList []*Ipv6AddrNode `protobuf:"bytes,54,rep,name=address_list,json=addressList,proto3" json:"address_list,omitempty"`
	// Link Local Address
	LinkLocalAddress *Ipv6AddrNode `protobuf:"bytes,55,opt,name=link_local_address,json=linkLocalAddress,proto3" json:"link_local_address,omitempty"`
	// IPv6 Access Control List
	AccessControlList *Ipv6AclConfig `protobuf:"bytes,56,opt,name=access_control_list,json=accessControlList,proto3" json:"access_control_list,omitempty"`
	// Multi IPv6 Access Control List
	MultiAccessControlList *Ipv6MultiAclConfig `protobuf:"bytes,57,opt,name=multi_access_control_list,json=multiAccessControlList,proto3" json:"multi_access_control_list,omitempty"`
	// ICMP unreach Enable
	IsIcmpUnreachEnabled bool `protobuf:"varint,58,opt,name=is_icmp_unreach_enabled,json=isIcmpUnreachEnabled,proto3" json:"is_icmp_unreach_enabled,omitempty"`
	// RPF config on the interface
	Rpf *RpfConfig `protobuf:"bytes,59,opt,name=rpf,proto3" json:"rpf,omitempty"`
	// BGP PA config on the interface
	BgpPa *BgpPaConfig `protobuf:"bytes,60,opt,name=bgp_pa,json=bgpPa,proto3" json:"bgp_pa,omitempty"`
	// Does ICCP RG ID exist on the interface?
	RgIdExists bool `protobuf:"varint,61,opt,name=rg_id_exists,json=rgIdExists,proto3" json:"rg_id_exists,omitempty"`
	// Is mLACP state Active (valid if RG ID exists)
	MLacpActive bool `protobuf:"varint,62,opt,name=m_lacp_active,json=mLacpActive,proto3" json:"m_lacp_active,omitempty"`
	// Is BGP Flow Tag Source is enable
	FlowTagSrc bool `protobuf:"varint,63,opt,name=flow_tag_src,json=flowTagSrc,proto3" json:"flow_tag_src,omitempty"`
	// Is BGP Flow Tag Destination is enable
	FlowTagDst bool `protobuf:"varint,64,opt,name=flow_tag_dst,json=flowTagDst,proto3" json:"flow_tag_dst,omitempty"`
	// Address Publish Time
	Utime *TimevalEntry `protobuf:"bytes,65,opt,name=utime,proto3" json:"utime,omitempty"`
	// IDB Create Time
	IdbUtime *TimevalEntry `protobuf:"bytes,66,opt,name=idb_utime,json=idbUtime,proto3" json:"idb_utime,omitempty"`
	// CAPS Add Time
	CapsUtime *TimevalEntry `protobuf:"bytes,67,opt,name=caps_utime,json=capsUtime,proto3" json:"caps_utime,omitempty"`
	// FWD ENABLE Time
	FwdEnUtime *TimevalEntry `protobuf:"bytes,68,opt,name=fwd_en_utime,json=fwdEnUtime,proto3" json:"fwd_en_utime,omitempty"`
	// FWD DISABLE Time
	FwdDisUtime *TimevalEntry `protobuf:"bytes,69,opt,name=fwd_dis_utime,json=fwdDisUtime,proto3" json:"fwd_dis_utime,omitempty"`
	// IPv6 Client Multicast Group
	ClientMulticastGroups []*Ipv6McastGroup `protobuf:"bytes,70,rep,name=client_multicast_groups,json=clientMulticastGroups,proto3" json:"client_multicast_groups,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *Ipv6IfDetail) Reset()         { *m = Ipv6IfDetail{} }
func (m *Ipv6IfDetail) String() string { return proto.CompactTextString(m) }
func (*Ipv6IfDetail) ProtoMessage()    {}
func (*Ipv6IfDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{1}
}
func (m *Ipv6IfDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6IfDetail.Unmarshal(m, b)
}
func (m *Ipv6IfDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6IfDetail.Marshal(b, m, deterministic)
}
func (dst *Ipv6IfDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6IfDetail.Merge(dst, src)
}
func (m *Ipv6IfDetail) XXX_Size() int {
	return xxx_messageInfo_Ipv6IfDetail.Size(m)
}
func (m *Ipv6IfDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6IfDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6IfDetail proto.InternalMessageInfo

func (m *Ipv6IfDetail) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *Ipv6IfDetail) GetMulticastGroups() []*Ipv6McastGroup {
	if m != nil {
		return m.MulticastGroups
	}
	return nil
}

func (m *Ipv6IfDetail) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Ipv6IfDetail) GetOperationState() string {
	if m != nil {
		return m.OperationState
	}
	return ""
}

func (m *Ipv6IfDetail) GetAddressList() []*Ipv6AddrNode {
	if m != nil {
		return m.AddressList
	}
	return nil
}

func (m *Ipv6IfDetail) GetLinkLocalAddress() *Ipv6AddrNode {
	if m != nil {
		return m.LinkLocalAddress
	}
	return nil
}

func (m *Ipv6IfDetail) GetAccessControlList() *Ipv6AclConfig {
	if m != nil {
		return m.AccessControlList
	}
	return nil
}

func (m *Ipv6IfDetail) GetMultiAccessControlList() *Ipv6MultiAclConfig {
	if m != nil {
		return m.MultiAccessControlList
	}
	return nil
}

func (m *Ipv6IfDetail) GetIsIcmpUnreachEnabled() bool {
	if m != nil {
		return m.IsIcmpUnreachEnabled
	}
	return false
}

func (m *Ipv6IfDetail) GetRpf() *RpfConfig {
	if m != nil {
		return m.Rpf
	}
	return nil
}

func (m *Ipv6IfDetail) GetBgpPa() *BgpPaConfig {
	if m != nil {
		return m.BgpPa
	}
	return nil
}

func (m *Ipv6IfDetail) GetRgIdExists() bool {
	if m != nil {
		return m.RgIdExists
	}
	return false
}

func (m *Ipv6IfDetail) GetMLacpActive() bool {
	if m != nil {
		return m.MLacpActive
	}
	return false
}

func (m *Ipv6IfDetail) GetFlowTagSrc() bool {
	if m != nil {
		return m.FlowTagSrc
	}
	return false
}

func (m *Ipv6IfDetail) GetFlowTagDst() bool {
	if m != nil {
		return m.FlowTagDst
	}
	return false
}

func (m *Ipv6IfDetail) GetUtime() *TimevalEntry {
	if m != nil {
		return m.Utime
	}
	return nil
}

func (m *Ipv6IfDetail) GetIdbUtime() *TimevalEntry {
	if m != nil {
		return m.IdbUtime
	}
	return nil
}

func (m *Ipv6IfDetail) GetCapsUtime() *TimevalEntry {
	if m != nil {
		return m.CapsUtime
	}
	return nil
}

func (m *Ipv6IfDetail) GetFwdEnUtime() *TimevalEntry {
	if m != nil {
		return m.FwdEnUtime
	}
	return nil
}

func (m *Ipv6IfDetail) GetFwdDisUtime() *TimevalEntry {
	if m != nil {
		return m.FwdDisUtime
	}
	return nil
}

func (m *Ipv6IfDetail) GetClientMulticastGroups() []*Ipv6McastGroup {
	if m != nil {
		return m.ClientMulticastGroups
	}
	return nil
}

type Str struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Str) Reset()         { *m = Str{} }
func (m *Str) String() string { return proto.CompactTextString(m) }
func (*Str) ProtoMessage()    {}
func (*Str) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{2}
}
func (m *Str) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Str.Unmarshal(m, b)
}
func (m *Str) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Str.Marshal(b, m, deterministic)
}
func (dst *Str) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Str.Merge(dst, src)
}
func (m *Str) XXX_Size() int {
	return xxx_messageInfo_Str.Size(m)
}
func (m *Str) XXX_DiscardUnknown() {
	xxx_messageInfo_Str.DiscardUnknown(m)
}

var xxx_messageInfo_Str proto.InternalMessageInfo

func (m *Str) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Creatation or Update Time
type TimevalEntry struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimevalEntry) Reset()         { *m = TimevalEntry{} }
func (m *TimevalEntry) String() string { return proto.CompactTextString(m) }
func (*TimevalEntry) ProtoMessage()    {}
func (*TimevalEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{3}
}
func (m *TimevalEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimevalEntry.Unmarshal(m, b)
}
func (m *TimevalEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimevalEntry.Marshal(b, m, deterministic)
}
func (dst *TimevalEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimevalEntry.Merge(dst, src)
}
func (m *TimevalEntry) XXX_Size() int {
	return xxx_messageInfo_TimevalEntry.Size(m)
}
func (m *TimevalEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TimevalEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TimevalEntry proto.InternalMessageInfo

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
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{4}
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

type Ipv6AclConfig struct {
	// ACL applied to incoming packets
	InBound string `protobuf:"bytes,1,opt,name=in_bound,json=inBound,proto3" json:"in_bound,omitempty"`
	// ACL applied to outgoing packets
	OutBound string `protobuf:"bytes,2,opt,name=out_bound,json=outBound,proto3" json:"out_bound,omitempty"`
	// Common ACL applied to incoming packets
	CommonInBound string `protobuf:"bytes,3,opt,name=common_in_bound,json=commonInBound,proto3" json:"common_in_bound,omitempty"`
	// Common ACL applied to outgoing packets
	CommonOutBound       string   `protobuf:"bytes,4,opt,name=common_out_bound,json=commonOutBound,proto3" json:"common_out_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6AclConfig) Reset()         { *m = Ipv6AclConfig{} }
func (m *Ipv6AclConfig) String() string { return proto.CompactTextString(m) }
func (*Ipv6AclConfig) ProtoMessage()    {}
func (*Ipv6AclConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{5}
}
func (m *Ipv6AclConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6AclConfig.Unmarshal(m, b)
}
func (m *Ipv6AclConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6AclConfig.Marshal(b, m, deterministic)
}
func (dst *Ipv6AclConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6AclConfig.Merge(dst, src)
}
func (m *Ipv6AclConfig) XXX_Size() int {
	return xxx_messageInfo_Ipv6AclConfig.Size(m)
}
func (m *Ipv6AclConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6AclConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6AclConfig proto.InternalMessageInfo

func (m *Ipv6AclConfig) GetInBound() string {
	if m != nil {
		return m.InBound
	}
	return ""
}

func (m *Ipv6AclConfig) GetOutBound() string {
	if m != nil {
		return m.OutBound
	}
	return ""
}

func (m *Ipv6AclConfig) GetCommonInBound() string {
	if m != nil {
		return m.CommonInBound
	}
	return ""
}

func (m *Ipv6AclConfig) GetCommonOutBound() string {
	if m != nil {
		return m.CommonOutBound
	}
	return ""
}

type Ipv6MultiAclConfig struct {
	// Inbound ACLs
	Inbound []*Str `protobuf:"bytes,1,rep,name=inbound,proto3" json:"inbound,omitempty"`
	// Outbound ACLs
	Outbound []*Str `protobuf:"bytes,2,rep,name=outbound,proto3" json:"outbound,omitempty"`
	// Common ACLs
	Common               []*Str   `protobuf:"bytes,3,rep,name=common,proto3" json:"common,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6MultiAclConfig) Reset()         { *m = Ipv6MultiAclConfig{} }
func (m *Ipv6MultiAclConfig) String() string { return proto.CompactTextString(m) }
func (*Ipv6MultiAclConfig) ProtoMessage()    {}
func (*Ipv6MultiAclConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{6}
}
func (m *Ipv6MultiAclConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6MultiAclConfig.Unmarshal(m, b)
}
func (m *Ipv6MultiAclConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6MultiAclConfig.Marshal(b, m, deterministic)
}
func (dst *Ipv6MultiAclConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6MultiAclConfig.Merge(dst, src)
}
func (m *Ipv6MultiAclConfig) XXX_Size() int {
	return xxx_messageInfo_Ipv6MultiAclConfig.Size(m)
}
func (m *Ipv6MultiAclConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6MultiAclConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6MultiAclConfig proto.InternalMessageInfo

func (m *Ipv6MultiAclConfig) GetInbound() []*Str {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *Ipv6MultiAclConfig) GetOutbound() []*Str {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *Ipv6MultiAclConfig) GetCommon() []*Str {
	if m != nil {
		return m.Common
	}
	return nil
}

// MCast Group
type Ipv6McastGroup struct {
	// IPv6 Address of Multicast Group
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6McastGroup) Reset()         { *m = Ipv6McastGroup{} }
func (m *Ipv6McastGroup) String() string { return proto.CompactTextString(m) }
func (*Ipv6McastGroup) ProtoMessage()    {}
func (*Ipv6McastGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{7}
}
func (m *Ipv6McastGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6McastGroup.Unmarshal(m, b)
}
func (m *Ipv6McastGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6McastGroup.Marshal(b, m, deterministic)
}
func (dst *Ipv6McastGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6McastGroup.Merge(dst, src)
}
func (m *Ipv6McastGroup) XXX_Size() int {
	return xxx_messageInfo_Ipv6McastGroup.Size(m)
}
func (m *Ipv6McastGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6McastGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6McastGroup proto.InternalMessageInfo

func (m *Ipv6McastGroup) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// RPF config information
type RpfConfig struct {
	// Enable RPF config
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// Allow Default Route
	AllowDefaultRoute bool `protobuf:"varint,2,opt,name=allow_default_route,json=allowDefaultRoute,proto3" json:"allow_default_route,omitempty"`
	// Allow Self Ping
	AllowSelfPing bool `protobuf:"varint,3,opt,name=allow_self_ping,json=allowSelfPing,proto3" json:"allow_self_ping,omitempty"`
	// RPF Mode (loose/strict)
	Mode                 uint32   `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpfConfig) Reset()         { *m = RpfConfig{} }
func (m *RpfConfig) String() string { return proto.CompactTextString(m) }
func (*RpfConfig) ProtoMessage()    {}
func (*RpfConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{8}
}
func (m *RpfConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpfConfig.Unmarshal(m, b)
}
func (m *RpfConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpfConfig.Marshal(b, m, deterministic)
}
func (dst *RpfConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpfConfig.Merge(dst, src)
}
func (m *RpfConfig) XXX_Size() int {
	return xxx_messageInfo_RpfConfig.Size(m)
}
func (m *RpfConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RpfConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RpfConfig proto.InternalMessageInfo

func (m *RpfConfig) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *RpfConfig) GetAllowDefaultRoute() bool {
	if m != nil {
		return m.AllowDefaultRoute
	}
	return false
}

func (m *RpfConfig) GetAllowSelfPing() bool {
	if m != nil {
		return m.AllowSelfPing
	}
	return false
}

func (m *RpfConfig) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

// BGP PA config for ingress/egress direction
type BgpPaDir struct {
	// Enable BGP PA for ingress/egress
	Enable uint32 `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// Enable source accouting
	Source bool `protobuf:"varint,2,opt,name=source,proto3" json:"source,omitempty"`
	// Enable destination accouting
	Destination          bool     `protobuf:"varint,3,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpPaDir) Reset()         { *m = BgpPaDir{} }
func (m *BgpPaDir) String() string { return proto.CompactTextString(m) }
func (*BgpPaDir) ProtoMessage()    {}
func (*BgpPaDir) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{9}
}
func (m *BgpPaDir) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPaDir.Unmarshal(m, b)
}
func (m *BgpPaDir) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPaDir.Marshal(b, m, deterministic)
}
func (dst *BgpPaDir) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPaDir.Merge(dst, src)
}
func (m *BgpPaDir) XXX_Size() int {
	return xxx_messageInfo_BgpPaDir.Size(m)
}
func (m *BgpPaDir) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPaDir.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPaDir proto.InternalMessageInfo

func (m *BgpPaDir) GetEnable() uint32 {
	if m != nil {
		return m.Enable
	}
	return 0
}

func (m *BgpPaDir) GetSource() bool {
	if m != nil {
		return m.Source
	}
	return false
}

func (m *BgpPaDir) GetDestination() bool {
	if m != nil {
		return m.Destination
	}
	return false
}

// BGP PA config information
type BgpPaConfig struct {
	// BGP PA input config
	Input *BgpPaDir `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// BGP PA output config
	Output               *BgpPaDir `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BgpPaConfig) Reset()         { *m = BgpPaConfig{} }
func (m *BgpPaConfig) String() string { return proto.CompactTextString(m) }
func (*BgpPaConfig) ProtoMessage()    {}
func (*BgpPaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv6_if_detail_903b9145e71c146c, []int{10}
}
func (m *BgpPaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPaConfig.Unmarshal(m, b)
}
func (m *BgpPaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPaConfig.Marshal(b, m, deterministic)
}
func (dst *BgpPaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPaConfig.Merge(dst, src)
}
func (m *BgpPaConfig) XXX_Size() int {
	return xxx_messageInfo_BgpPaConfig.Size(m)
}
func (m *BgpPaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPaConfig proto.InternalMessageInfo

func (m *BgpPaConfig) GetInput() *BgpPaDir {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *BgpPaConfig) GetOutput() *BgpPaDir {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6IfDetail_KEYS)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.ipv6_if_detail_KEYS")
	proto.RegisterType((*Ipv6IfDetail)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.ipv6_if_detail")
	proto.RegisterType((*Str)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.str")
	proto.RegisterType((*TimevalEntry)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.timeval_entry")
	proto.RegisterType((*Ipv6AddrNode)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.ipv6_addr_node")
	proto.RegisterType((*Ipv6AclConfig)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.ipv6_acl_config")
	proto.RegisterType((*Ipv6MultiAclConfig)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.ipv6_multi_acl_config")
	proto.RegisterType((*Ipv6McastGroup)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.ipv6_mcast_group")
	proto.RegisterType((*RpfConfig)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.rpf_config")
	proto.RegisterType((*BgpPaDir)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.bgp_pa_dir")
	proto.RegisterType((*BgpPaConfig)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.global_details.global_detail.bgp_pa_config")
}

func init() {
	proto.RegisterFile("ipv6_if_detail.proto", fileDescriptor_ipv6_if_detail_903b9145e71c146c)
}

var fileDescriptor_ipv6_if_detail_903b9145e71c146c = []byte{
	// 1090 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x97, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x35, 0x75, 0xe3, 0xda, 0x2f, 0x71, 0x92, 0x4e, 0xd2, 0x74, 0xab, 0xaa, 0x92, 0xb5,
	0x08, 0xc8, 0x01, 0xf9, 0xd0, 0xd2, 0xf2, 0xfb, 0x47, 0xda, 0x04, 0x14, 0x11, 0xa0, 0xda, 0xb4,
	0x07, 0x2e, 0x8c, 0xc6, 0xbb, 0xb3, 0xdb, 0x51, 0x66, 0x67, 0x56, 0x33, 0xb3, 0x76, 0x7a, 0xe0,
	0x04, 0x08, 0x04, 0x1c, 0x10, 0x42, 0x42, 0x15, 0x47, 0x90, 0x7a, 0xa4, 0xe2, 0xc4, 0x9f, 0x87,
	0x66, 0x66, 0xd7, 0xc1, 0x69, 0xb9, 0xda, 0x17, 0xcb, 0xf3, 0xde, 0xdb, 0xf7, 0x3e, 0xef, 0xe9,
	0xd9, 0xf3, 0x5d, 0xd8, 0xe6, 0xd5, 0xe4, 0x0e, 0xe1, 0x39, 0xc9, 0x98, 0xa5, 0x5c, 0x8c, 0x2a,
	0xad, 0xac, 0xc2, 0x79, 0xca, 0x4d, 0xaa, 0x08, 0x57, 0x86, 0x9c, 0x6a, 0xe2, 0x43, 0x4a, 0x4a,
	0x54, 0xc5, 0xf4, 0xc8, 0x1f, 0x24, 0xb3, 0x53, 0xa5, 0x4f, 0x46, 0x52, 0x65, 0xcc, 0xf8, 0xcf,
	0x11, 0x97, 0x96, 0xe9, 0x9c, 0xa6, 0x8c, 0x64, 0xd4, 0xd2, 0xd1, 0x44, 0xe7, 0xc6, 0x7d, 0x8c,
	0x0a, 0xa1, 0xc6, 0x54, 0x34, 0xe9, 0xcd, 0xfc, 0x31, 0xb6, 0xb0, 0x35, 0x5f, 0x9f, 0x7c, 0x72,
	0xf0, 0xc5, 0x31, 0xbe, 0x0e, 0x7d, 0x97, 0x93, 0x48, 0x5a, 0xb2, 0x08, 0x0d, 0xd1, 0x6e, 0x3f,
	0xe9, 0x39, 0xc3, 0x67, 0xb4, 0x64, 0xf8, 0x1a, 0xf4, 0x26, 0x3a, 0x0f, 0xbe, 0x0b, 0xde, 0x77,
	0x69, 0xa2, 0x73, 0xef, 0x7a, 0x19, 0xd6, 0xcf, 0x30, 0x7c, 0x40, 0xc7, 0x07, 0x0c, 0x66, 0x56,
	0x17, 0x16, 0xff, 0x8d, 0x61, 0x7d, 0xbe, 0x2c, 0xbe, 0x01, 0x20, 0xb8, 0x64, 0xc4, 0x58, 0x6a,
	0x59, 0x74, 0xd3, 0x3f, 0xd5, 0x77, 0x96, 0x63, 0x67, 0xc0, 0x7f, 0x20, 0xd8, 0x2c, 0x6b, 0x61,
	0x79, 0x4a, 0x8d, 0x25, 0x85, 0x56, 0x75, 0x65, 0xa2, 0x5b, 0xc3, 0xce, 0xee, 0xea, 0xcd, 0xd3,
	0xd1, 0x62, 0x66, 0x15, 0x92, 0x95, 0x67, 0x00, 0xc9, 0xc6, 0x8c, 0xe8, 0x63, 0x0f, 0x84, 0x37,
	0xa1, 0x53, 0xda, 0x3a, 0x7a, 0x7d, 0x88, 0x76, 0x07, 0x89, 0xfb, 0x8a, 0x5f, 0x85, 0x0d, 0x07,
	0x42, 0x2d, 0x57, 0xb2, 0xe9, 0xed, 0xb6, 0xef, 0x6d, 0x7d, 0x66, 0x0e, 0x0d, 0x3e, 0x41, 0xb0,
	0x46, 0xb3, 0x4c, 0x33, 0x63, 0x88, 0xe0, 0xc6, 0x46, 0x77, 0x7c, 0x73, 0x93, 0x85, 0x36, 0xe7,
	0x00, 0x88, 0xcb, 0x91, 0xac, 0x36, 0x2c, 0x47, 0xdc, 0x58, 0xfc, 0x27, 0x02, 0x2c, 0xb8, 0x3c,
	0x21, 0x42, 0xa5, 0x54, 0x90, 0xc6, 0x15, 0xbd, 0x31, 0x44, 0x4b, 0x24, 0xdc, 0x74, 0x44, 0x47,
	0x0e, 0x68, 0x2f, 0xf0, 0xe0, 0xa7, 0x08, 0xb6, 0x68, 0x9a, 0xba, 0x09, 0xa6, 0x4a, 0x5a, 0xad,
	0x44, 0x98, 0xe4, 0x9b, 0x9e, 0x73, 0xba, 0x58, 0xce, 0x54, 0x38, 0x88, 0x9c, 0x17, 0xc9, 0xe5,
	0xc0, 0x74, 0x2f, 0x20, 0xf9, 0x81, 0xfe, 0x83, 0xe0, 0x9a, 0xdf, 0x1d, 0xf2, 0x22, 0xde, 0xb7,
	0x3c, 0xef, 0x57, 0x8b, 0x5d, 0xeb, 0x86, 0x66, 0x46, 0xbd, 0xe3, 0x2d, 0x7b, 0xcf, 0xa1, 0xdf,
	0x86, 0xab, 0xdc, 0x10, 0x9e, 0x96, 0x15, 0xa9, 0xa5, 0x66, 0x34, 0x7d, 0x44, 0x98, 0xa4, 0x63,
	0xc1, 0xb2, 0xe8, 0xed, 0x21, 0xda, 0xed, 0x25, 0xdb, 0xdc, 0x1c, 0xa6, 0x65, 0xf5, 0x30, 0x38,
	0x0f, 0x82, 0x0f, 0x7f, 0x83, 0xa0, 0xa3, 0xab, 0x3c, 0x7a, 0xc7, 0xf7, 0xa6, 0x17, 0xd5, 0x9b,
	0xae, 0xf2, 0xb6, 0x21, 0x57, 0x1e, 0xff, 0x84, 0xa0, 0x3b, 0x2e, 0x2a, 0x52, 0xd1, 0xe8, 0x5d,
	0x4f, 0x52, 0x2f, 0x8a, 0x24, 0x54, 0x6d, 0x61, 0x56, 0xc6, 0x45, 0x75, 0x9f, 0xe2, 0x21, 0xac,
	0xe9, 0x82, 0xf0, 0x8c, 0xb0, 0x53, 0x6e, 0xac, 0x89, 0xde, 0xf3, 0x13, 0x04, 0x5d, 0x1c, 0x66,
	0x07, 0xde, 0x82, 0x63, 0x18, 0x94, 0x44, 0xd0, 0xb4, 0x22, 0x34, 0xb5, 0x7c, 0xc2, 0xa2, 0xf7,
	0x7d, 0xc8, 0x6a, 0x79, 0x44, 0xd3, 0x6a, 0xcf, 0x9b, 0x5c, 0x96, 0x5c, 0xa8, 0x29, 0xb1, 0xb4,
	0x20, 0x46, 0xa7, 0xd1, 0x07, 0x21, 0x8b, 0xb3, 0x3d, 0xa0, 0xc5, 0xb1, 0x4e, 0xe7, 0x22, 0x32,
	0x63, 0xa3, 0x0f, 0xe7, 0x22, 0xf6, 0x8d, 0xc5, 0x3f, 0x22, 0x58, 0xa9, 0x2d, 0x2f, 0x59, 0xb4,
	0xb7, 0xd8, 0xb9, 0xb8, 0x9a, 0x13, 0x2a, 0x08, 0x93, 0x56, 0x3f, 0x4e, 0x02, 0x03, 0xfe, 0x05,
	0x41, 0x9f, 0x67, 0x63, 0x12, 0x88, 0xee, 0x2e, 0x93, 0xa8, 0xc7, 0xb3, 0xf1, 0x43, 0x0f, 0xf5,
	0x2b, 0x02, 0x48, 0x69, 0x65, 0x1a, 0xaa, 0x7b, 0xcb, 0xa4, 0xea, 0x3b, 0x90, 0x80, 0xf5, 0x1b,
	0x82, 0xb5, 0x7c, 0x9a, 0x11, 0x26, 0x1b, 0xb0, 0xfd, 0x65, 0x82, 0x41, 0x3e, 0xcd, 0x0e, 0x64,
	0x20, 0x7b, 0x82, 0x60, 0xe0, 0xc8, 0x32, 0xde, 0xce, 0xec, 0x60, 0x99, 0x68, 0xab, 0xf9, 0x34,
	0xdb, 0xe7, 0xcd, 0xd4, 0x9e, 0x21, 0xb8, 0x9a, 0x0a, 0xce, 0xa4, 0x25, 0xcf, 0xc9, 0x8a, 0x8f,
	0x96, 0x2c, 0x2b, 0xae, 0x04, 0xb0, 0x4f, 0xe7, 0xc5, 0x45, 0x7c, 0x1d, 0x3a, 0xc6, 0x6a, 0xbc,
	0x0d, 0x2b, 0x13, 0x2a, 0xea, 0x56, 0x96, 0x85, 0x43, 0xbc, 0x01, 0x83, 0xb9, 0x6e, 0xe3, 0x67,
	0xa8, 0x91, 0x58, 0xb3, 0x1b, 0x13, 0x47, 0x70, 0xa9, 0xbd, 0xba, 0xc3, 0xb3, 0xed, 0x11, 0xbf,
	0x04, 0x83, 0x4a, 0xb3, 0x9c, 0x9f, 0x12, 0xc1, 0x64, 0x61, 0x1f, 0x79, 0x59, 0x37, 0x48, 0xd6,
	0x82, 0xf1, 0xc8, 0xdb, 0x5c, 0x50, 0x2b, 0x50, 0x82, 0x90, 0x09, 0xd2, 0xae, 0x55, 0x2d, 0x41,
	0xc6, 0xdc, 0x00, 0xe0, 0x86, 0x50, 0xf9, 0xd8, 0x81, 0x47, 0x17, 0xfd, 0xff, 0x4c, 0x9f, 0x9b,
	0xbd, 0x60, 0x70, 0xba, 0x52, 0xab, 0xda, 0x32, 0xf7, 0x4f, 0x14, 0xad, 0xf8, 0x22, 0x3d, 0x6f,
	0x78, 0x40, 0x8b, 0xf8, 0x77, 0x04, 0x1b, 0xe7, 0x2e, 0x4f, 0xa7, 0x35, 0xb9, 0x24, 0x63, 0x55,
	0xcb, 0xac, 0x85, 0xe6, 0xf2, 0xae, 0x3b, 0xba, 0x5c, 0xaa, 0xb6, 0x8d, 0x2f, 0xe8, 0xd0, 0x9e,
	0xaa, 0x6d, 0x70, 0xbe, 0x02, 0x1b, 0xa9, 0x2a, 0x4b, 0x25, 0xc9, 0xec, 0xf1, 0x46, 0x89, 0x06,
	0xf3, 0x61, 0x93, 0x64, 0x17, 0x36, 0x9b, 0xb8, 0xb3, 0x5c, 0x17, 0x83, 0x40, 0x0b, 0xf6, 0xcf,
	0x9b, 0x8c, 0xf1, 0x5f, 0x1d, 0xb8, 0xf2, 0xc2, 0xab, 0x12, 0x7f, 0x8b, 0xe0, 0x12, 0x97, 0x2d,
	0xa3, 0xdb, 0x9d, 0x93, 0x45, 0xed, 0x8e, 0xb1, 0x3a, 0x69, 0x6b, 0xe3, 0xef, 0x10, 0xb8, 0x01,
	0xb4, 0x03, 0x59, 0x38, 0xc8, 0xac, 0x38, 0xfe, 0x1a, 0x41, 0x37, 0x8c, 0x2f, 0xea, 0x2c, 0x9e,
	0xa3, 0x29, 0x1d, 0xbf, 0x06, 0x9b, 0xe7, 0x7f, 0x5b, 0xff, 0xff, 0x1b, 0x88, 0x7f, 0x46, 0x00,
	0x67, 0x72, 0x01, 0xef, 0x40, 0x37, 0xe8, 0x1a, 0x1f, 0xd7, 0x4b, 0x9a, 0x13, 0x1e, 0xc1, 0x16,
	0x15, 0xee, 0x2e, 0xcd, 0x58, 0x4e, 0x6b, 0x61, 0x89, 0x5f, 0x5f, 0xbf, 0x7f, 0xbd, 0xe4, 0xb2,
	0x77, 0xed, 0x07, 0x4f, 0xe2, 0x1c, 0x6e, 0x11, 0x43, 0xbc, 0x61, 0x22, 0x27, 0x15, 0x97, 0x85,
	0x5f, 0xc4, 0x5e, 0x32, 0xf0, 0xe6, 0x63, 0x26, 0xf2, 0xfb, 0x5c, 0x16, 0x18, 0xc3, 0xc5, 0x52,
	0x65, 0xcc, 0x2f, 0xdf, 0x20, 0xf1, 0xdf, 0xe3, 0x2f, 0x01, 0x1a, 0xd9, 0x90, 0x71, 0x7d, 0x8e,
	0x68, 0x30, 0x23, 0xda, 0x81, 0xae, 0x51, 0xb5, 0x4e, 0x5b, 0x88, 0xe6, 0x84, 0x87, 0xb0, 0x9a,
	0x31, 0x63, 0xb9, 0xf4, 0x6f, 0x19, 0x4d, 0xd5, 0xff, 0x9a, 0xe2, 0xa7, 0x17, 0x60, 0x30, 0xa7,
	0x4b, 0xf0, 0xf7, 0x08, 0x56, 0xb8, 0xac, 0x6a, 0xeb, 0x6b, 0x2c, 0x50, 0xa8, 0x9d, 0xf5, 0x99,
	0x04, 0x00, 0xfc, 0x03, 0x82, 0xae, 0xaa, 0xad, 0x63, 0xb9, 0xb0, 0x34, 0x96, 0x86, 0x60, 0xdc,
	0xf5, 0x6f, 0xe5, 0xb7, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xbd, 0x7a, 0x0e, 0xad, 0x0f,
	0x00, 0x00,
}
