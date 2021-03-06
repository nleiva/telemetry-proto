// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_acl_edm_ace.proto

package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_access_list_manager_accesses_access_access_list_sequences_access_list_sequence

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

// ACLE bag
type Ipv4AclEdmAce_KEYS struct {
	AccessListName       string   `protobuf:"bytes,1,opt,name=access_list_name,json=accessListName,proto3" json:"access_list_name,omitempty"`
	SequenceNumber       uint32   `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclEdmAce_KEYS) Reset()         { *m = Ipv4AclEdmAce_KEYS{} }
func (m *Ipv4AclEdmAce_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclEdmAce_KEYS) ProtoMessage()    {}
func (*Ipv4AclEdmAce_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_acl_edm_ace_8aa296eaeec7bc93, []int{0}
}
func (m *Ipv4AclEdmAce_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclEdmAce_KEYS.Unmarshal(m, b)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclEdmAce_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv4AclEdmAce_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclEdmAce_KEYS.Merge(dst, src)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclEdmAce_KEYS.Size(m)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclEdmAce_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclEdmAce_KEYS proto.InternalMessageInfo

func (m *Ipv4AclEdmAce_KEYS) GetAccessListName() string {
	if m != nil {
		return m.AccessListName
	}
	return ""
}

func (m *Ipv4AclEdmAce_KEYS) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type Ipv4AclEdmAce struct {
	// ACE type (acl, remark)
	ItemType string `protobuf:"bytes,50,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	// ACLE sequence number
	Sequence uint32 `protobuf:"varint,51,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Permit/deny
	Grant string `protobuf:"bytes,52,opt,name=grant,proto3" json:"grant,omitempty"`
	// IPv4 protocol type
	Protocol uint32 `protobuf:"varint,53,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Source address
	SourceAddress string `protobuf:"bytes,54,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Source mask
	SourceAddressMask string `protobuf:"bytes,55,opt,name=source_address_mask,json=sourceAddressMask,proto3" json:"source_address_mask,omitempty"`
	// Destination address
	DestinationAddress string `protobuf:"bytes,56,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	// Destination mask
	DestinationAddressMask string `protobuf:"bytes,57,opt,name=destination_address_mask,json=destinationAddressMask,proto3" json:"destination_address_mask,omitempty"`
	// Source operator
	SourceOperator string `protobuf:"bytes,58,opt,name=source_operator,json=sourceOperator,proto3" json:"source_operator,omitempty"`
	// Source port 1
	SourcePort1 uint32 `protobuf:"varint,59,opt,name=source_port1,json=sourcePort1,proto3" json:"source_port1,omitempty"`
	// Source port 2
	SourcePort2 uint32 `protobuf:"varint,60,opt,name=source_port2,json=sourcePort2,proto3" json:"source_port2,omitempty"`
	// Deprecated by Source operator
	SorceOperator string `protobuf:"bytes,61,opt,name=sorce_operator,json=sorceOperator,proto3" json:"sorce_operator,omitempty"`
	// Deprecated by SourcePort1
	SorcePort1 uint32 `protobuf:"varint,62,opt,name=sorce_port1,json=sorcePort1,proto3" json:"sorce_port1,omitempty"`
	// Deprecated by SourcePort2
	SorcePort2 uint32 `protobuf:"varint,63,opt,name=sorce_port2,json=sorcePort2,proto3" json:"sorce_port2,omitempty"`
	// Destination operator
	DestinationOperator string `protobuf:"bytes,64,opt,name=destination_operator,json=destinationOperator,proto3" json:"destination_operator,omitempty"`
	// Destination port 1
	DestinationPort1 uint32 `protobuf:"varint,65,opt,name=destination_port1,json=destinationPort1,proto3" json:"destination_port1,omitempty"`
	// Destination port 2
	DestinationPort2 uint32 `protobuf:"varint,66,opt,name=destination_port2,json=destinationPort2,proto3" json:"destination_port2,omitempty"`
	// Log option
	LogOption string `protobuf:"bytes,67,opt,name=log_option,json=logOption,proto3" json:"log_option,omitempty"`
	// Counter name
	CounterName string `protobuf:"bytes,68,opt,name=counter_name,json=counterName,proto3" json:"counter_name,omitempty"`
	// Capture option, TRUE if enabled
	Capture bool `protobuf:"varint,69,opt,name=capture,proto3" json:"capture,omitempty"`
	// DSCP present
	DscpPresent bool `protobuf:"varint,70,opt,name=dscp_present,json=dscpPresent,proto3" json:"dscp_present,omitempty"`
	// DSCP or DSCP range start
	Dscp uint32 `protobuf:"varint,71,opt,name=dscp,proto3" json:"dscp,omitempty"`
	// DSCP Range End
	Dscp2 uint32 `protobuf:"varint,72,opt,name=dscp2,proto3" json:"dscp2,omitempty"`
	// DSCP Operator
	DscpOperator uint32 `protobuf:"varint,73,opt,name=dscp_operator,json=dscpOperator,proto3" json:"dscp_operator,omitempty"`
	// Precedence present
	PrecedencePresent bool `protobuf:"varint,74,opt,name=precedence_present,json=precedencePresent,proto3" json:"precedence_present,omitempty"`
	// Precedence
	Precedence uint32 `protobuf:"varint,75,opt,name=precedence,proto3" json:"precedence,omitempty"`
	// TCP flags operator
	TcpFlagsOperator string `protobuf:"bytes,76,opt,name=tcp_flags_operator,json=tcpFlagsOperator,proto3" json:"tcp_flags_operator,omitempty"`
	// TCP flags
	TcpFlags uint32 `protobuf:"varint,77,opt,name=tcp_flags,json=tcpFlags,proto3" json:"tcp_flags,omitempty"`
	// TCP flags mask
	TcpFlagsMask uint32 `protobuf:"varint,78,opt,name=tcp_flags_mask,json=tcpFlagsMask,proto3" json:"tcp_flags_mask,omitempty"`
	// Fragments
	Fragments uint32 `protobuf:"varint,79,opt,name=fragments,proto3" json:"fragments,omitempty"`
	// Port length operator
	PortLengthOperator string `protobuf:"bytes,80,opt,name=port_length_operator,json=portLengthOperator,proto3" json:"port_length_operator,omitempty"`
	// Port length 1
	PortLength1 uint32 `protobuf:"varint,81,opt,name=port_length1,json=portLength1,proto3" json:"port_length1,omitempty"`
	// Port length 2
	PortLength2 uint32 `protobuf:"varint,82,opt,name=port_length2,json=portLength2,proto3" json:"port_length2,omitempty"`
	// TTL operator
	TtlOperator string `protobuf:"bytes,83,opt,name=ttl_operator,json=ttlOperator,proto3" json:"ttl_operator,omitempty"`
	// TTL 1
	Ttl1 uint32 `protobuf:"varint,84,opt,name=ttl1,proto3" json:"ttl1,omitempty"`
	// TTL 2
	Ttl2 uint32 `protobuf:"varint,85,opt,name=ttl2,proto3" json:"ttl2,omitempty"`
	// No stats
	NoStats bool `protobuf:"varint,86,opt,name=no_stats,json=noStats,proto3" json:"no_stats,omitempty"`
	// Number of hits
	Hits uint64 `protobuf:"varint,87,opt,name=hits,proto3" json:"hits,omitempty"`
	// True if ICMP off
	IsIcmpOff bool `protobuf:"varint,88,opt,name=is_icmp_off,json=isIcmpOff,proto3" json:"is_icmp_off,omitempty"`
	// Next hop type
	NextHopType string `protobuf:"bytes,89,opt,name=next_hop_type,json=nextHopType,proto3" json:"next_hop_type,omitempty"`
	// Next hop info
	NextHopInfo []*Ipv4AclBagNhInfo `protobuf:"bytes,90,rep,name=next_hop_info,json=nextHopInfo,proto3" json:"next_hop_info,omitempty"`
	// HW Next hop info
	HwNextHopInfo *Ipv4AclBagHwNhInfo `protobuf:"bytes,91,opt,name=hw_next_hop_info,json=hwNextHopInfo,proto3" json:"hw_next_hop_info,omitempty"`
	// Remark String
	Remark string `protobuf:"bytes,92,opt,name=remark,proto3" json:"remark,omitempty"`
	// Is dynamic ACE
	Dynamic bool `protobuf:"varint,93,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	// Source prefix object-group
	SourcePrefixGroup string `protobuf:"bytes,94,opt,name=source_prefix_group,json=sourcePrefixGroup,proto3" json:"source_prefix_group,omitempty"`
	// Destination prefix object-group
	DestinationPrefixGroup string `protobuf:"bytes,95,opt,name=destination_prefix_group,json=destinationPrefixGroup,proto3" json:"destination_prefix_group,omitempty"`
	// Source port object-group
	SourcePortGroup string `protobuf:"bytes,96,opt,name=source_port_group,json=sourcePortGroup,proto3" json:"source_port_group,omitempty"`
	// Destination port object-group
	DestinationPortGroup string `protobuf:"bytes,97,opt,name=destination_port_group,json=destinationPortGroup,proto3" json:"destination_port_group,omitempty"`
	// ACL Name
	AclName string `protobuf:"bytes,98,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	// Sequence String
	SequenceStr          string   `protobuf:"bytes,99,opt,name=sequence_str,json=sequenceStr,proto3" json:"sequence_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclEdmAce) Reset()         { *m = Ipv4AclEdmAce{} }
func (m *Ipv4AclEdmAce) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclEdmAce) ProtoMessage()    {}
func (*Ipv4AclEdmAce) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_acl_edm_ace_8aa296eaeec7bc93, []int{1}
}
func (m *Ipv4AclEdmAce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclEdmAce.Unmarshal(m, b)
}
func (m *Ipv4AclEdmAce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclEdmAce.Marshal(b, m, deterministic)
}
func (dst *Ipv4AclEdmAce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclEdmAce.Merge(dst, src)
}
func (m *Ipv4AclEdmAce) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclEdmAce.Size(m)
}
func (m *Ipv4AclEdmAce) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclEdmAce.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclEdmAce proto.InternalMessageInfo

func (m *Ipv4AclEdmAce) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetGrant() string {
	if m != nil {
		return m.Grant
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourceAddressMask() string {
	if m != nil {
		return m.SourceAddressMask
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationAddressMask() string {
	if m != nil {
		return m.DestinationAddressMask
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourceOperator() string {
	if m != nil {
		return m.SourceOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourcePort1() uint32 {
	if m != nil {
		return m.SourcePort1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSourcePort2() uint32 {
	if m != nil {
		return m.SourcePort2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSorceOperator() string {
	if m != nil {
		return m.SorceOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSorcePort1() uint32 {
	if m != nil {
		return m.SorcePort1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSorcePort2() uint32 {
	if m != nil {
		return m.SorcePort2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDestinationOperator() string {
	if m != nil {
		return m.DestinationOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationPort1() uint32 {
	if m != nil {
		return m.DestinationPort1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDestinationPort2() uint32 {
	if m != nil {
		return m.DestinationPort2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetLogOption() string {
	if m != nil {
		return m.LogOption
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetCounterName() string {
	if m != nil {
		return m.CounterName
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetCapture() bool {
	if m != nil {
		return m.Capture
	}
	return false
}

func (m *Ipv4AclEdmAce) GetDscpPresent() bool {
	if m != nil {
		return m.DscpPresent
	}
	return false
}

func (m *Ipv4AclEdmAce) GetDscp() uint32 {
	if m != nil {
		return m.Dscp
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDscp2() uint32 {
	if m != nil {
		return m.Dscp2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDscpOperator() uint32 {
	if m != nil {
		return m.DscpOperator
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetPrecedencePresent() bool {
	if m != nil {
		return m.PrecedencePresent
	}
	return false
}

func (m *Ipv4AclEdmAce) GetPrecedence() uint32 {
	if m != nil {
		return m.Precedence
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTcpFlagsOperator() string {
	if m != nil {
		return m.TcpFlagsOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetTcpFlags() uint32 {
	if m != nil {
		return m.TcpFlags
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTcpFlagsMask() uint32 {
	if m != nil {
		return m.TcpFlagsMask
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetFragments() uint32 {
	if m != nil {
		return m.Fragments
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetPortLengthOperator() string {
	if m != nil {
		return m.PortLengthOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetPortLength1() uint32 {
	if m != nil {
		return m.PortLength1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetPortLength2() uint32 {
	if m != nil {
		return m.PortLength2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTtlOperator() string {
	if m != nil {
		return m.TtlOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetTtl1() uint32 {
	if m != nil {
		return m.Ttl1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTtl2() uint32 {
	if m != nil {
		return m.Ttl2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetNoStats() bool {
	if m != nil {
		return m.NoStats
	}
	return false
}

func (m *Ipv4AclEdmAce) GetHits() uint64 {
	if m != nil {
		return m.Hits
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetIsIcmpOff() bool {
	if m != nil {
		return m.IsIcmpOff
	}
	return false
}

func (m *Ipv4AclEdmAce) GetNextHopType() string {
	if m != nil {
		return m.NextHopType
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetNextHopInfo() []*Ipv4AclBagNhInfo {
	if m != nil {
		return m.NextHopInfo
	}
	return nil
}

func (m *Ipv4AclEdmAce) GetHwNextHopInfo() *Ipv4AclBagHwNhInfo {
	if m != nil {
		return m.HwNextHopInfo
	}
	return nil
}

func (m *Ipv4AclEdmAce) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *Ipv4AclEdmAce) GetSourcePrefixGroup() string {
	if m != nil {
		return m.SourcePrefixGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationPrefixGroup() string {
	if m != nil {
		return m.DestinationPrefixGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourcePortGroup() string {
	if m != nil {
		return m.SourcePortGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationPortGroup() string {
	if m != nil {
		return m.DestinationPortGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSequenceStr() string {
	if m != nil {
		return m.SequenceStr
	}
	return ""
}

// NH_Info structure
type Ipv4AclBagNhInfo struct {
	// The next hop
	NextHop string `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	// Track name
	TrackName string `protobuf:"bytes,2,opt,name=track_name,json=trackName,proto3" json:"track_name,omitempty"`
	// The next hop status
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// The next hop at status
	AtStatus string `protobuf:"bytes,4,opt,name=at_status,json=atStatus,proto3" json:"at_status,omitempty"`
	// The nexthop exist
	IsAclNextHopExist    bool     `protobuf:"varint,5,opt,name=is_acl_next_hop_exist,json=isAclNextHopExist,proto3" json:"is_acl_next_hop_exist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclBagNhInfo) Reset()         { *m = Ipv4AclBagNhInfo{} }
func (m *Ipv4AclBagNhInfo) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclBagNhInfo) ProtoMessage()    {}
func (*Ipv4AclBagNhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_acl_edm_ace_8aa296eaeec7bc93, []int{2}
}
func (m *Ipv4AclBagNhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclBagNhInfo.Unmarshal(m, b)
}
func (m *Ipv4AclBagNhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclBagNhInfo.Marshal(b, m, deterministic)
}
func (dst *Ipv4AclBagNhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclBagNhInfo.Merge(dst, src)
}
func (m *Ipv4AclBagNhInfo) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclBagNhInfo.Size(m)
}
func (m *Ipv4AclBagNhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclBagNhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclBagNhInfo proto.InternalMessageInfo

func (m *Ipv4AclBagNhInfo) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetTrackName() string {
	if m != nil {
		return m.TrackName
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetAtStatus() string {
	if m != nil {
		return m.AtStatus
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetIsAclNextHopExist() bool {
	if m != nil {
		return m.IsAclNextHopExist
	}
	return false
}

// HW_NH_Info structure
type Ipv4AclBagHwNhInfo struct {
	// The Next Hop
	NextHop uint32 `protobuf:"varint,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	// the next-hop type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// VRF name
	VrfName              string   `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclBagHwNhInfo) Reset()         { *m = Ipv4AclBagHwNhInfo{} }
func (m *Ipv4AclBagHwNhInfo) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclBagHwNhInfo) ProtoMessage()    {}
func (*Ipv4AclBagHwNhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_acl_edm_ace_8aa296eaeec7bc93, []int{3}
}
func (m *Ipv4AclBagHwNhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclBagHwNhInfo.Unmarshal(m, b)
}
func (m *Ipv4AclBagHwNhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclBagHwNhInfo.Marshal(b, m, deterministic)
}
func (dst *Ipv4AclBagHwNhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclBagHwNhInfo.Merge(dst, src)
}
func (m *Ipv4AclBagHwNhInfo) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclBagHwNhInfo.Size(m)
}
func (m *Ipv4AclBagHwNhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclBagHwNhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclBagHwNhInfo proto.InternalMessageInfo

func (m *Ipv4AclBagHwNhInfo) GetNextHop() uint32 {
	if m != nil {
		return m.NextHop
	}
	return 0
}

func (m *Ipv4AclBagHwNhInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Ipv4AclBagHwNhInfo) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv4AclEdmAce_KEYS)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_edm_ace_KEYS")
	proto.RegisterType((*Ipv4AclEdmAce)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_edm_ace")
	proto.RegisterType((*Ipv4AclBagNhInfo)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_bag_nh_info")
	proto.RegisterType((*Ipv4AclBagHwNhInfo)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_bag_hw_nh_info")
}

func init() {
	proto.RegisterFile("ipv4_acl_edm_ace.proto", fileDescriptor_ipv4_acl_edm_ace_8aa296eaeec7bc93)
}

var fileDescriptor_ipv4_acl_edm_ace_8aa296eaeec7bc93 = []byte{
	// 1072 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x5b, 0x77, 0x1b, 0x35,
	0x10, 0x3e, 0xdb, 0xa6, 0x49, 0x2c, 0xe3, 0xd4, 0x51, 0xd3, 0xa0, 0x52, 0x28, 0x26, 0xc0, 0xc1,
	0x87, 0x4b, 0x68, 0xb6, 0x01, 0xca, 0x9d, 0x00, 0x69, 0x1b, 0x9a, 0x26, 0x21, 0x29, 0x97, 0x72,
	0x13, 0x8a, 0xac, 0xb5, 0x97, 0xec, 0xae, 0x16, 0x49, 0x4e, 0x93, 0x57, 0x7e, 0x05, 0xe7, 0xf0,
	0x43, 0x78, 0xe2, 0x91, 0x17, 0x7e, 0x15, 0x47, 0x33, 0x7b, 0xb3, 0x13, 0x9e, 0xe9, 0x93, 0x77,
	0xbe, 0xf9, 0x46, 0xdf, 0x68, 0x34, 0x1a, 0x99, 0x2c, 0xc7, 0xf9, 0xf1, 0x3a, 0x17, 0x32, 0xe1,
	0x6a, 0x90, 0x72, 0x21, 0xd5, 0x6a, 0x6e, 0xb4, 0xd3, 0xf4, 0xb7, 0x40, 0xc6, 0x56, 0x6a, 0x1e,
	0x6b, 0xcb, 0x4f, 0x0c, 0xaf, 0x58, 0x3a, 0x57, 0x66, 0xb5, 0xb2, 0x44, 0x36, 0xe0, 0xb9, 0x51,
	0x51, 0x7c, 0xc2, 0x93, 0xd8, 0xba, 0x55, 0x21, 0xa5, 0xb2, 0x16, 0xbe, 0x79, 0x2a, 0x32, 0x31,
	0x54, 0xa6, 0xc0, 0x94, 0x2d, 0x3e, 0x26, 0x38, 0x56, 0xfd, 0x3a, 0x56, 0x99, 0x54, 0xe7, 0xa3,
	0x2b, 0xbf, 0x90, 0xab, 0xd3, 0xe9, 0xf1, 0xfb, 0x9b, 0x8f, 0x0e, 0x68, 0x9f, 0x74, 0x9b, 0x01,
	0x99, 0x48, 0x15, 0x0b, 0x7a, 0x41, 0xbf, 0xb5, 0xbf, 0x80, 0xf8, 0x76, 0x6c, 0xdd, 0x8e, 0x48,
	0x15, 0x7d, 0x85, 0x5c, 0x2e, 0x97, 0xe3, 0xd9, 0x38, 0x3d, 0x54, 0x86, 0x5d, 0xe8, 0x05, 0xfd,
	0xce, 0xfe, 0x42, 0x09, 0xef, 0x00, 0xba, 0xf2, 0x4f, 0x97, 0x74, 0xa7, 0xc5, 0xe8, 0x75, 0xd2,
	0x8a, 0x9d, 0x4a, 0xb9, 0x3b, 0xcd, 0x15, 0x0b, 0x41, 0x60, 0xde, 0x03, 0x0f, 0x4f, 0x73, 0x45,
	0x9f, 0x21, 0xf3, 0xe5, 0x1a, 0xec, 0x16, 0xac, 0x59, 0xd9, 0x74, 0x89, 0x5c, 0x1a, 0x1a, 0x91,
	0x39, 0xb6, 0x0e, 0x41, 0x68, 0xf8, 0x08, 0xa8, 0xae, 0xd4, 0x09, 0x7b, 0x0b, 0x23, 0x4a, 0x9b,
	0xbe, 0x4c, 0x16, 0xac, 0x1e, 0x1b, 0xa9, 0xb8, 0x18, 0x0c, 0x8c, 0xb2, 0x96, 0xbd, 0x0d, 0xa1,
	0x1d, 0x44, 0x37, 0x10, 0xa4, 0xab, 0xe4, 0xca, 0x24, 0x8d, 0xa7, 0xc2, 0x1e, 0xb1, 0x77, 0x80,
	0xbb, 0x38, 0xc1, 0x7d, 0x20, 0xec, 0x11, 0x7d, 0x93, 0x5c, 0x19, 0x28, 0xeb, 0xe2, 0x4c, 0xb8,
	0x58, 0x67, 0xd5, 0xda, 0xb7, 0x81, 0x4f, 0x1b, 0xae, 0x52, 0xe0, 0x36, 0x61, 0xe7, 0x04, 0xa0,
	0xca, 0xbb, 0x10, 0xb5, 0x7c, 0x36, 0x0a, 0xa4, 0x7c, 0xa9, 0x31, 0x35, 0xdf, 0x24, 0xc2, 0x69,
	0xc3, 0xde, 0xc3, 0x33, 0x41, 0x78, 0xb7, 0x40, 0xe9, 0x0b, 0xe4, 0xa9, 0x82, 0x98, 0x6b, 0xe3,
	0xd6, 0xd8, 0xfb, 0x50, 0x8a, 0x36, 0x62, 0x7b, 0x1e, 0x9a, 0xa2, 0x84, 0xec, 0x83, 0x69, 0x4a,
	0x88, 0x05, 0x9b, 0x50, 0xfb, 0xb0, 0x2c, 0x58, 0x53, 0xec, 0x79, 0xd2, 0x46, 0x1a, 0x6a, 0x7d,
	0x04, 0x0b, 0x11, 0x80, 0x50, 0x6a, 0x82, 0x10, 0xb2, 0x8f, 0xa7, 0x08, 0x21, 0x5d, 0x23, 0x4b,
	0xcd, 0x8a, 0x54, 0x72, 0x9f, 0x80, 0x5c, 0xb3, 0xbc, 0x95, 0xe8, 0x6b, 0x64, 0xb1, 0x19, 0x82,
	0xd2, 0x1b, 0xb0, 0x72, 0xb7, 0xe1, 0xc0, 0x04, 0xce, 0x21, 0x87, 0xec, 0xd3, 0x73, 0xc9, 0x21,
	0x7d, 0x8e, 0x90, 0x44, 0x0f, 0xb9, 0xce, 0x3d, 0xc4, 0x3e, 0x83, 0x14, 0x5a, 0x89, 0x1e, 0xee,
	0x02, 0xe0, 0xeb, 0x26, 0xf5, 0x38, 0x73, 0xca, 0xe0, 0xa5, 0xf8, 0x1c, 0x08, 0xed, 0x02, 0x83,
	0x1b, 0xc1, 0xc8, 0x9c, 0x14, 0xb9, 0x1b, 0x1b, 0xc5, 0x36, 0x7b, 0x41, 0x7f, 0x7e, 0xbf, 0x34,
	0x7d, 0xf0, 0xc0, 0xca, 0xdc, 0xdf, 0x68, 0xab, 0x32, 0xc7, 0xee, 0x80, 0xbb, 0xed, 0xb1, 0x3d,
	0x84, 0x28, 0x25, 0x33, 0xde, 0x64, 0x77, 0x21, 0x3d, 0xf8, 0xf6, 0xbd, 0xee, 0x7f, 0x43, 0x76,
	0x0f, 0x40, 0x34, 0xe8, 0x8b, 0xa4, 0x03, 0x8b, 0x55, 0xe5, 0xda, 0x02, 0x2f, 0x28, 0x54, 0x75,
	0x7a, 0x83, 0xd0, 0xdc, 0x28, 0xa9, 0x06, 0x70, 0x3f, 0x4b, 0xdd, 0x2f, 0x40, 0x77, 0xb1, 0xf6,
	0x94, 0xea, 0x37, 0x08, 0xa9, 0x41, 0x76, 0x1f, 0x4f, 0xaa, 0x46, 0xe8, 0xeb, 0x84, 0x3a, 0x99,
	0xf3, 0x28, 0x11, 0x43, 0x5b, 0x0b, 0x6f, 0x43, 0x0d, 0xba, 0x4e, 0xe6, 0x77, 0xbc, 0xa3, 0x12,
	0xbf, 0x4e, 0x5a, 0x15, 0x9b, 0x3d, 0xc0, 0xeb, 0x58, 0x92, 0xe8, 0x4b, 0x64, 0xa1, 0x5e, 0x0a,
	0x9a, 0x7f, 0x07, 0xf3, 0x2f, 0x19, 0xd0, 0xf2, 0xcf, 0x92, 0x56, 0x64, 0xc4, 0x30, 0x55, 0x99,
	0xb3, 0x6c, 0x17, 0x08, 0x35, 0x40, 0x6f, 0x92, 0x25, 0x7f, 0x98, 0x3c, 0x51, 0xd9, 0xd0, 0x8d,
	0xea, 0x84, 0xf6, 0xf0, 0xf2, 0x79, 0xdf, 0x36, 0xb8, 0x9a, 0x37, 0xa3, 0x11, 0xb1, 0xc6, 0xbe,
	0xc4, 0xb6, 0xaf, 0x99, 0x6b, 0x53, 0x94, 0x90, 0xed, 0x4f, 0x53, 0x42, 0x4f, 0x71, 0x2e, 0xa9,
	0xf5, 0x0e, 0xb0, 0x09, 0x9c, 0x4b, 0x2a, 0x21, 0x4a, 0x66, 0x9c, 0x4b, 0xd6, 0xd8, 0x43, 0x3c,
	0x47, 0xff, 0x5d, 0x60, 0x21, 0xfb, 0xaa, 0xc2, 0x42, 0x7a, 0x8d, 0xcc, 0x67, 0x9a, 0x5b, 0x27,
	0x9c, 0x65, 0x5f, 0x63, 0xb7, 0x64, 0xfa, 0xc0, 0x9b, 0x9e, 0x3e, 0x8a, 0x9d, 0x65, 0xdf, 0xf4,
	0x82, 0xfe, 0xcc, 0x3e, 0x7c, 0xd3, 0x1b, 0xa4, 0x1d, 0x5b, 0x1e, 0xcb, 0x34, 0xe7, 0x3a, 0x8a,
	0xd8, 0xb7, 0x10, 0xd1, 0x8a, 0xed, 0x96, 0x4c, 0xf3, 0xdd, 0x28, 0xa2, 0x2b, 0xa4, 0x93, 0xa9,
	0x13, 0xc7, 0x47, 0x3a, 0xc7, 0x99, 0xfa, 0x08, 0x53, 0xf3, 0xe0, 0x3d, 0x9d, 0xc3, 0x58, 0xfd,
	0x2b, 0x68, 0x90, 0xe2, 0x2c, 0xd2, 0xec, 0xbb, 0xde, 0xc5, 0x7e, 0x3b, 0xfc, 0x3d, 0x58, 0xfd,
	0xff, 0x9f, 0xa4, 0x5a, 0xe7, 0x50, 0x0c, 0x79, 0x36, 0x82, 0x04, 0xab, 0xfc, 0xb7, 0xb2, 0x48,
	0xd3, 0xbf, 0x03, 0xd2, 0x1d, 0x3d, 0xe6, 0x93, 0x5b, 0xf8, 0xbe, 0x17, 0xf4, 0xdb, 0xe1, 0x1f,
	0x4f, 0xde, 0x16, 0x7c, 0xa6, 0xc5, 0x2e, 0x3a, 0xa3, 0xc7, 0x3b, 0x8d, 0x7d, 0x2c, 0x93, 0x59,
	0xa3, 0x52, 0x61, 0x8e, 0xd8, 0x0f, 0x70, 0x48, 0x85, 0xe5, 0xe7, 0xc7, 0xe0, 0x34, 0x13, 0x69,
	0x2c, 0xd9, 0x8f, 0xd8, 0x11, 0x85, 0xd9, 0x78, 0x9b, 0x8a, 0xec, 0x87, 0x46, 0x8f, 0x73, 0xf6,
	0x53, 0xf3, 0x6d, 0xda, 0x03, 0xcf, 0x5d, 0xef, 0x98, 0x7e, 0x6a, 0x26, 0x82, 0xf8, 0x99, 0xa7,
	0xa6, 0x19, 0xf9, 0x2a, 0x59, 0x6c, 0x3c, 0x0f, 0x45, 0xc8, 0xcf, 0x10, 0x72, 0xb9, 0x7e, 0x23,
	0x90, 0xbb, 0x4e, 0x96, 0xa7, 0xc7, 0x6b, 0x11, 0x20, 0x20, 0x60, 0x69, 0x6a, 0xc6, 0x62, 0xd4,
	0x35, 0x32, 0xef, 0x4b, 0x04, 0x43, 0xf4, 0x10, 0x78, 0x73, 0x42, 0x26, 0x30, 0x40, 0xfd, 0xdb,
	0x54, 0xfe, 0xa5, 0xb0, 0xce, 0x30, 0x89, 0x3d, 0x5c, 0x62, 0x07, 0xce, 0xac, 0xfc, 0x19, 0x90,
	0xa5, 0xf3, 0x3a, 0x05, 0xee, 0x53, 0xd1, 0x18, 0xc5, 0x1f, 0x96, 0xb9, 0xa2, 0x77, 0xfc, 0x64,
	0x77, 0x46, 0xc8, 0x23, 0xd4, 0xbc, 0x80, 0x93, 0x1d, 0x10, 0x50, 0x5d, 0x26, 0xb3, 0xfe, 0x1a,
	0x8e, 0x2d, 0xbb, 0x88, 0xc7, 0x81, 0x96, 0x9f, 0x62, 0xc2, 0xf1, 0xc2, 0x35, 0x83, 0x7f, 0x51,
	0x84, 0x3b, 0x40, 0xe7, 0x4d, 0x72, 0x35, 0xb6, 0x90, 0x44, 0xd5, 0x8e, 0xea, 0x24, 0xb6, 0x8e,
	0x5d, 0xc2, 0x11, 0x1b, 0xdb, 0x0d, 0x99, 0x14, 0x87, 0xbe, 0xe9, 0x1d, 0x2b, 0x92, 0x3c, 0xfd,
	0x1f, 0xfd, 0x71, 0x26, 0xf7, 0x4e, 0x9d, 0xbb, 0x1f, 0x1d, 0xfe, 0x3a, 0x63, 0xd6, 0xf0, 0xed,
	0xe9, 0xc7, 0x26, 0xc2, 0xdd, 0x60, 0xca, 0x73, 0xc7, 0x26, 0xf2, 0x7b, 0x39, 0x9c, 0x85, 0x7f,
	0x3d, 0xb7, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xde, 0xb1, 0xf8, 0x7d, 0x0a, 0x00, 0x00,
}
