// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_virtual_links.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_interface_information_virtual_links_virtual_link

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

// OSPF Virtual Link
type OspfShVirtualLinks_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	VirtualLinkName      string   `protobuf:"bytes,3,opt,name=virtual_link_name,json=virtualLinkName,proto3" json:"virtual_link_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShVirtualLinks_KEYS) Reset()         { *m = OspfShVirtualLinks_KEYS{} }
func (m *OspfShVirtualLinks_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShVirtualLinks_KEYS) ProtoMessage()    {}
func (*OspfShVirtualLinks_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_virtual_links_a6ab197f955dc787, []int{0}
}
func (m *OspfShVirtualLinks_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShVirtualLinks_KEYS.Unmarshal(m, b)
}
func (m *OspfShVirtualLinks_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShVirtualLinks_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShVirtualLinks_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShVirtualLinks_KEYS.Merge(dst, src)
}
func (m *OspfShVirtualLinks_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShVirtualLinks_KEYS.Size(m)
}
func (m *OspfShVirtualLinks_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShVirtualLinks_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShVirtualLinks_KEYS proto.InternalMessageInfo

func (m *OspfShVirtualLinks_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShVirtualLinks_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShVirtualLinks_KEYS) GetVirtualLinkName() string {
	if m != nil {
		return m.VirtualLinkName
	}
	return ""
}

type OspfShVirtualLinks struct {
	// Virtual link name
	VirtualLinkName string `protobuf:"bytes,50,opt,name=virtual_link_name,json=virtualLinkName,proto3" json:"virtual_link_name,omitempty"`
	// Neighbor on other end of this virtual link
	VirtualLinkNeighborId string `protobuf:"bytes,51,opt,name=virtual_link_neighbor_id,json=virtualLinkNeighborId,proto3" json:"virtual_link_neighbor_id,omitempty"`
	// OSPF interface state for the virtual link
	VirtualLinkState string `protobuf:"bytes,52,opt,name=virtual_link_state,json=virtualLinkState,proto3" json:"virtual_link_state,omitempty"`
	// If true, the link runs as demand circuit
	VirtualLinkDemandCircuit bool `protobuf:"varint,53,opt,name=virtual_link_demand_circuit,json=virtualLinkDemandCircuit,proto3" json:"virtual_link_demand_circuit,omitempty"`
	// Number of LSA's with demand circuit bit not set
	VirtualLinkDcBitlessLsa uint32 `protobuf:"varint,54,opt,name=virtual_link_dc_bitless_lsa,json=virtualLinkDcBitlessLsa,proto3" json:"virtual_link_dc_bitless_lsa,omitempty"`
	// Transit area id
	TransitArea string `protobuf:"bytes,55,opt,name=transit_area,json=transitArea,proto3" json:"transit_area,omitempty"`
	// Interface on which this virtual link is formed
	VirtualLinkInterfaceName string `protobuf:"bytes,56,opt,name=virtual_link_interface_name,json=virtualLinkInterfaceName,proto3" json:"virtual_link_interface_name,omitempty"`
	// Cost of the virtual link
	VirtualLinkCost uint32 `protobuf:"varint,57,opt,name=virtual_link_cost,json=virtualLinkCost,proto3" json:"virtual_link_cost,omitempty"`
	// Transmission delay in seconds
	VirualLinkTransmissionDelay uint32 `protobuf:"varint,58,opt,name=virual_link_transmission_delay,json=virualLinkTransmissionDelay,proto3" json:"virual_link_transmission_delay,omitempty"`
	// Hello interval (s)
	VirtualLinkHelloInterval uint32 `protobuf:"varint,59,opt,name=virtual_link_hello_interval,json=virtualLinkHelloInterval,proto3" json:"virtual_link_hello_interval,omitempty"`
	// Hello interval (ms)
	VirtualLinkHelloIntervalMs uint32 `protobuf:"varint,60,opt,name=virtual_link_hello_interval_ms,json=virtualLinkHelloIntervalMs,proto3" json:"virtual_link_hello_interval_ms,omitempty"`
	// Dead interval (s)
	VirtualLinkDeadInterval uint32 `protobuf:"varint,61,opt,name=virtual_link_dead_interval,json=virtualLinkDeadInterval,proto3" json:"virtual_link_dead_interval,omitempty"`
	// Wait interval (s)
	VirtualLinkWaitInterval uint32 `protobuf:"varint,62,opt,name=virtual_link_wait_interval,json=virtualLinkWaitInterval,proto3" json:"virtual_link_wait_interval,omitempty"`
	// Retransmission interval (s)
	VirtaulLinkRetransmissionInterval uint32 `protobuf:"varint,63,opt,name=virtaul_link_retransmission_interval,json=virtaulLinkRetransmissionInterval,proto3" json:"virtaul_link_retransmission_interval,omitempty"`
	// Time until next hello (s)
	VirtualLinkNextHello uint32 `protobuf:"varint,64,opt,name=virtual_link_next_hello,json=virtualLinkNextHello,proto3" json:"virtual_link_next_hello,omitempty"`
	// Time until next hello (ms)
	VirtualLinkNextHelloMs uint32 `protobuf:"varint,65,opt,name=virtual_link_next_hello_ms,json=virtualLinkNextHelloMs,proto3" json:"virtual_link_next_hello_ms,omitempty"`
	// If true, interface is passive
	VirtualLinkPassive bool `protobuf:"varint,66,opt,name=virtual_link_passive,json=virtualLinkPassive,proto3" json:"virtual_link_passive,omitempty"`
	// Authentication type
	VirtualLinkAuthenticationType string `protobuf:"bytes,67,opt,name=virtual_link_authentication_type,json=virtualLinkAuthenticationType,proto3" json:"virtual_link_authentication_type,omitempty"`
	// If true, MD key configured
	VirtualLinkYoungestMdKey bool `protobuf:"varint,68,opt,name=virtual_link_youngest_md_key,json=virtualLinkYoungestMdKey,proto3" json:"virtual_link_youngest_md_key,omitempty"`
	// Youngest MD key ID
	VirtualLinkYoungestMdKeyId uint32 `protobuf:"varint,69,opt,name=virtual_link_youngest_md_key_id,json=virtualLinkYoungestMdKeyId,proto3" json:"virtual_link_youngest_md_key_id,omitempty"`
	// Number of neighbors still using the old key (rollover in progress)
	VirtualLinkOldMdKeyCount uint32 `protobuf:"varint,70,opt,name=virtual_link_old_md_key_count,json=virtualLinkOldMdKeyCount,proto3" json:"virtual_link_old_md_key_count,omitempty"`
	// List of old MD keys (if any)
	VirtualLinkMdKeyList []uint32 `protobuf:"varint,71,rep,packed,name=virtual_link_md_key_list,json=virtualLinkMdKeyList,proto3" json:"virtual_link_md_key_list,omitempty"`
	// Virtual Link Keychain ID
	VirtualLinkKeychainId uint64 `protobuf:"varint,72,opt,name=virtual_link_keychain_id,json=virtualLinkKeychainId,proto3" json:"virtual_link_keychain_id,omitempty"`
	// If true, NSF enabled
	VirtualLinkNsfEnabled bool `protobuf:"varint,73,opt,name=virtual_link_nsf_enabled,json=virtualLinkNsfEnabled,proto3" json:"virtual_link_nsf_enabled,omitempty"`
	// If true, NSF restart in progress on the virtual link
	VirtualLinkNsf bool `protobuf:"varint,74,opt,name=virtual_link_nsf,json=virtualLinkNsf,proto3" json:"virtual_link_nsf,omitempty"`
	// Time in seconds since last NSF
	VirtualLinkLastNsf uint32 `protobuf:"varint,75,opt,name=virtual_link_last_nsf,json=virtualLinkLastNsf,proto3" json:"virtual_link_last_nsf,omitempty"`
	// Neighbor information
	VirtualLinkNeighbor  *OspfShVlinkNeighbor `protobuf:"bytes,76,opt,name=virtual_link_neighbor,json=virtualLinkNeighbor,proto3" json:"virtual_link_neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OspfShVirtualLinks) Reset()         { *m = OspfShVirtualLinks{} }
func (m *OspfShVirtualLinks) String() string { return proto.CompactTextString(m) }
func (*OspfShVirtualLinks) ProtoMessage()    {}
func (*OspfShVirtualLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_virtual_links_a6ab197f955dc787, []int{1}
}
func (m *OspfShVirtualLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShVirtualLinks.Unmarshal(m, b)
}
func (m *OspfShVirtualLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShVirtualLinks.Marshal(b, m, deterministic)
}
func (dst *OspfShVirtualLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShVirtualLinks.Merge(dst, src)
}
func (m *OspfShVirtualLinks) XXX_Size() int {
	return xxx_messageInfo_OspfShVirtualLinks.Size(m)
}
func (m *OspfShVirtualLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShVirtualLinks.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShVirtualLinks proto.InternalMessageInfo

func (m *OspfShVirtualLinks) GetVirtualLinkName() string {
	if m != nil {
		return m.VirtualLinkName
	}
	return ""
}

func (m *OspfShVirtualLinks) GetVirtualLinkNeighborId() string {
	if m != nil {
		return m.VirtualLinkNeighborId
	}
	return ""
}

func (m *OspfShVirtualLinks) GetVirtualLinkState() string {
	if m != nil {
		return m.VirtualLinkState
	}
	return ""
}

func (m *OspfShVirtualLinks) GetVirtualLinkDemandCircuit() bool {
	if m != nil {
		return m.VirtualLinkDemandCircuit
	}
	return false
}

func (m *OspfShVirtualLinks) GetVirtualLinkDcBitlessLsa() uint32 {
	if m != nil {
		return m.VirtualLinkDcBitlessLsa
	}
	return 0
}

func (m *OspfShVirtualLinks) GetTransitArea() string {
	if m != nil {
		return m.TransitArea
	}
	return ""
}

func (m *OspfShVirtualLinks) GetVirtualLinkInterfaceName() string {
	if m != nil {
		return m.VirtualLinkInterfaceName
	}
	return ""
}

func (m *OspfShVirtualLinks) GetVirtualLinkCost() uint32 {
	if m != nil {
		return m.VirtualLinkCost
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirualLinkTransmissionDelay() uint32 {
	if m != nil {
		return m.VirualLinkTransmissionDelay
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkHelloInterval() uint32 {
	if m != nil {
		return m.VirtualLinkHelloInterval
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkHelloIntervalMs() uint32 {
	if m != nil {
		return m.VirtualLinkHelloIntervalMs
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkDeadInterval() uint32 {
	if m != nil {
		return m.VirtualLinkDeadInterval
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkWaitInterval() uint32 {
	if m != nil {
		return m.VirtualLinkWaitInterval
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtaulLinkRetransmissionInterval() uint32 {
	if m != nil {
		return m.VirtaulLinkRetransmissionInterval
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkNextHello() uint32 {
	if m != nil {
		return m.VirtualLinkNextHello
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkNextHelloMs() uint32 {
	if m != nil {
		return m.VirtualLinkNextHelloMs
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkPassive() bool {
	if m != nil {
		return m.VirtualLinkPassive
	}
	return false
}

func (m *OspfShVirtualLinks) GetVirtualLinkAuthenticationType() string {
	if m != nil {
		return m.VirtualLinkAuthenticationType
	}
	return ""
}

func (m *OspfShVirtualLinks) GetVirtualLinkYoungestMdKey() bool {
	if m != nil {
		return m.VirtualLinkYoungestMdKey
	}
	return false
}

func (m *OspfShVirtualLinks) GetVirtualLinkYoungestMdKeyId() uint32 {
	if m != nil {
		return m.VirtualLinkYoungestMdKeyId
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkOldMdKeyCount() uint32 {
	if m != nil {
		return m.VirtualLinkOldMdKeyCount
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkMdKeyList() []uint32 {
	if m != nil {
		return m.VirtualLinkMdKeyList
	}
	return nil
}

func (m *OspfShVirtualLinks) GetVirtualLinkKeychainId() uint64 {
	if m != nil {
		return m.VirtualLinkKeychainId
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkNsfEnabled() bool {
	if m != nil {
		return m.VirtualLinkNsfEnabled
	}
	return false
}

func (m *OspfShVirtualLinks) GetVirtualLinkNsf() bool {
	if m != nil {
		return m.VirtualLinkNsf
	}
	return false
}

func (m *OspfShVirtualLinks) GetVirtualLinkLastNsf() uint32 {
	if m != nil {
		return m.VirtualLinkLastNsf
	}
	return 0
}

func (m *OspfShVirtualLinks) GetVirtualLinkNeighbor() *OspfShVlinkNeighbor {
	if m != nil {
		return m.VirtualLinkNeighbor
	}
	return nil
}

// OSPF Neighbor Retransmission Information
type OspfShNeighborRetrans struct {
	// Number of DBD retransmissions during last exchange
	DbdRetransmissionCount uint32 `protobuf:"varint,1,opt,name=dbd_retransmission_count,json=dbdRetransmissionCount,proto3" json:"dbd_retransmission_count,omitempty"`
	// Total number of DBD retransmissions for this neighbor
	DbdRetransmissionTotalCount uint32 `protobuf:"varint,2,opt,name=dbd_retransmission_total_count,json=dbdRetransmissionTotalCount,proto3" json:"dbd_retransmission_total_count,omitempty"`
	// Area scope LSA's flood index
	AreaFloodingIndex uint32 `protobuf:"varint,3,opt,name=area_flooding_index,json=areaFloodingIndex,proto3" json:"area_flooding_index,omitempty"`
	// AS scope LSA's flood index
	AsFloodIndex uint32 `protobuf:"varint,4,opt,name=as_flood_index,json=asFloodIndex,proto3" json:"as_flood_index,omitempty"`
	// Retransmission queue length
	NeighborRetransmissionCount uint32 `protobuf:"varint,5,opt,name=neighbor_retransmission_count,json=neighborRetransmissionCount,proto3" json:"neighbor_retransmission_count,omitempty"`
	// Number of retransmissions for this neighbor
	NumberOfRetransmissions uint32 `protobuf:"varint,6,opt,name=number_of_retransmissions,json=numberOfRetransmissions,proto3" json:"number_of_retransmissions,omitempty"`
	// First flood item for area scope LSAs
	AreaFirstFloodInformation uint32 `protobuf:"varint,7,opt,name=area_first_flood_information,json=areaFirstFloodInformation,proto3" json:"area_first_flood_information,omitempty"`
	// Index of the first flood item for area scope LSAs
	AreaFirstFloodInformationIndex uint32 `protobuf:"varint,8,opt,name=area_first_flood_information_index,json=areaFirstFloodInformationIndex,proto3" json:"area_first_flood_information_index,omitempty"`
	// First flood item for AS scope LSAs
	AsFirstFloodInformation uint32 `protobuf:"varint,9,opt,name=as_first_flood_information,json=asFirstFloodInformation,proto3" json:"as_first_flood_information,omitempty"`
	// Index for first flood item for AS scope LSAs
	AsFirstFloodInformationIndex uint32 `protobuf:"varint,10,opt,name=as_first_flood_information_index,json=asFirstFloodInformationIndex,proto3" json:"as_first_flood_information_index,omitempty"`
	// Next flood item for area scope LSAs
	AreaNextFloodInformation uint32 `protobuf:"varint,11,opt,name=area_next_flood_information,json=areaNextFloodInformation,proto3" json:"area_next_flood_information,omitempty"`
	// Index of next flood item for Area scope LSAs
	AreaNextFloodInformationIndex uint32 `protobuf:"varint,12,opt,name=area_next_flood_information_index,json=areaNextFloodInformationIndex,proto3" json:"area_next_flood_information_index,omitempty"`
	// Next flood item for AS scope LSAs
	AsNextFloodInformation uint32 `protobuf:"varint,13,opt,name=as_next_flood_information,json=asNextFloodInformation,proto3" json:"as_next_flood_information,omitempty"`
	// Index of next flood item for AS scope LSAs
	AsNextFloodInformationIndex uint32 `protobuf:"varint,14,opt,name=as_next_flood_information_index,json=asNextFloodInformationIndex,proto3" json:"as_next_flood_information_index,omitempty"`
	// Number of LSAs sent in last retransmission
	LastRetransmissionLength uint32 `protobuf:"varint,15,opt,name=last_retransmission_length,json=lastRetransmissionLength,proto3" json:"last_retransmission_length,omitempty"`
	// Maximum number of LSAs sent in a retransmission
	MaximumRetransmissionLength uint32 `protobuf:"varint,16,opt,name=maximum_retransmission_length,json=maximumRetransmissionLength,proto3" json:"maximum_retransmission_length,omitempty"`
	// Last retransmission scan time (ms)
	LastRetransmissionTime uint32 `protobuf:"varint,17,opt,name=last_retransmission_time,json=lastRetransmissionTime,proto3" json:"last_retransmission_time,omitempty"`
	// Maximum retransmission scan time (ms)
	MaximumRetransmissionTime uint32 `protobuf:"varint,18,opt,name=maximum_retransmission_time,json=maximumRetransmissionTime,proto3" json:"maximum_retransmission_time,omitempty"`
	// Time until next LSA retransmission (ms)
	LsaRetransmissionTimer uint32   `protobuf:"varint,19,opt,name=lsa_retransmission_timer,json=lsaRetransmissionTimer,proto3" json:"lsa_retransmission_timer,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *OspfShNeighborRetrans) Reset()         { *m = OspfShNeighborRetrans{} }
func (m *OspfShNeighborRetrans) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborRetrans) ProtoMessage()    {}
func (*OspfShNeighborRetrans) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_virtual_links_a6ab197f955dc787, []int{2}
}
func (m *OspfShNeighborRetrans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborRetrans.Unmarshal(m, b)
}
func (m *OspfShNeighborRetrans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborRetrans.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighborRetrans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborRetrans.Merge(dst, src)
}
func (m *OspfShNeighborRetrans) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborRetrans.Size(m)
}
func (m *OspfShNeighborRetrans) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborRetrans.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborRetrans proto.InternalMessageInfo

func (m *OspfShNeighborRetrans) GetDbdRetransmissionCount() uint32 {
	if m != nil {
		return m.DbdRetransmissionCount
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetDbdRetransmissionTotalCount() uint32 {
	if m != nil {
		return m.DbdRetransmissionTotalCount
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaFloodingIndex() uint32 {
	if m != nil {
		return m.AreaFloodingIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsFloodIndex() uint32 {
	if m != nil {
		return m.AsFloodIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetNeighborRetransmissionCount() uint32 {
	if m != nil {
		return m.NeighborRetransmissionCount
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetNumberOfRetransmissions() uint32 {
	if m != nil {
		return m.NumberOfRetransmissions
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaFirstFloodInformation() uint32 {
	if m != nil {
		return m.AreaFirstFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaFirstFloodInformationIndex() uint32 {
	if m != nil {
		return m.AreaFirstFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsFirstFloodInformation() uint32 {
	if m != nil {
		return m.AsFirstFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsFirstFloodInformationIndex() uint32 {
	if m != nil {
		return m.AsFirstFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaNextFloodInformation() uint32 {
	if m != nil {
		return m.AreaNextFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaNextFloodInformationIndex() uint32 {
	if m != nil {
		return m.AreaNextFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsNextFloodInformation() uint32 {
	if m != nil {
		return m.AsNextFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsNextFloodInformationIndex() uint32 {
	if m != nil {
		return m.AsNextFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetLastRetransmissionLength() uint32 {
	if m != nil {
		return m.LastRetransmissionLength
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetMaximumRetransmissionLength() uint32 {
	if m != nil {
		return m.MaximumRetransmissionLength
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetLastRetransmissionTime() uint32 {
	if m != nil {
		return m.LastRetransmissionTime
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetMaximumRetransmissionTime() uint32 {
	if m != nil {
		return m.MaximumRetransmissionTime
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetLsaRetransmissionTimer() uint32 {
	if m != nil {
		return m.LsaRetransmissionTimer
	}
	return 0
}

// Virtual Link Neighbor Information
type OspfShVlinkNeighbor struct {
	// If true Hellos suppressed
	VirtualLinkSuppressHello bool `protobuf:"varint,1,opt,name=virtual_link_suppress_hello,json=virtualLinkSuppressHello,proto3" json:"virtual_link_suppress_hello,omitempty"`
	// Adjacency state
	VirtualLinkState string `protobuf:"bytes,2,opt,name=virtual_link_state,json=virtualLinkState,proto3" json:"virtual_link_state,omitempty"`
	// Neighbor retransmission info
	VirtualLinkRetransmissoin *OspfShNeighborRetrans `protobuf:"bytes,3,opt,name=virtual_link_retransmissoin,json=virtualLinkRetransmissoin,proto3" json:"virtual_link_retransmissoin,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}               `json:"-"`
	XXX_unrecognized          []byte                 `json:"-"`
	XXX_sizecache             int32                  `json:"-"`
}

func (m *OspfShVlinkNeighbor) Reset()         { *m = OspfShVlinkNeighbor{} }
func (m *OspfShVlinkNeighbor) String() string { return proto.CompactTextString(m) }
func (*OspfShVlinkNeighbor) ProtoMessage()    {}
func (*OspfShVlinkNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_virtual_links_a6ab197f955dc787, []int{3}
}
func (m *OspfShVlinkNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShVlinkNeighbor.Unmarshal(m, b)
}
func (m *OspfShVlinkNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShVlinkNeighbor.Marshal(b, m, deterministic)
}
func (dst *OspfShVlinkNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShVlinkNeighbor.Merge(dst, src)
}
func (m *OspfShVlinkNeighbor) XXX_Size() int {
	return xxx_messageInfo_OspfShVlinkNeighbor.Size(m)
}
func (m *OspfShVlinkNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShVlinkNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShVlinkNeighbor proto.InternalMessageInfo

func (m *OspfShVlinkNeighbor) GetVirtualLinkSuppressHello() bool {
	if m != nil {
		return m.VirtualLinkSuppressHello
	}
	return false
}

func (m *OspfShVlinkNeighbor) GetVirtualLinkState() string {
	if m != nil {
		return m.VirtualLinkState
	}
	return ""
}

func (m *OspfShVlinkNeighbor) GetVirtualLinkRetransmissoin() *OspfShNeighborRetrans {
	if m != nil {
		return m.VirtualLinkRetransmissoin
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShVirtualLinks_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.virtual_links.virtual_link.ospf_sh_virtual_links_KEYS")
	proto.RegisterType((*OspfShVirtualLinks)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.virtual_links.virtual_link.ospf_sh_virtual_links")
	proto.RegisterType((*OspfShNeighborRetrans)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.virtual_links.virtual_link.ospf_sh_neighbor_retrans")
	proto.RegisterType((*OspfShVlinkNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.virtual_links.virtual_link.ospf_sh_vlink_neighbor")
}

func init() {
	proto.RegisterFile("ospf_sh_virtual_links.proto", fileDescriptor_ospf_sh_virtual_links_a6ab197f955dc787)
}

var fileDescriptor_ospf_sh_virtual_links_a6ab197f955dc787 = []byte{
	// 1174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdb, 0x4f, 0x1c, 0xb7,
	0x17, 0xd6, 0x90, 0xfc, 0x12, 0x30, 0x97, 0x80, 0x09, 0x89, 0xb9, 0x66, 0x41, 0x79, 0x58, 0xfd,
	0x54, 0xad, 0x5a, 0x72, 0x23, 0x5c, 0x42, 0x61, 0x81, 0xb0, 0x61, 0x29, 0xd5, 0x82, 0x54, 0xe5,
	0x69, 0xe4, 0x9d, 0xf1, 0xb2, 0x16, 0x73, 0x59, 0x8d, 0xbd, 0x5b, 0xf6, 0xbd, 0xca, 0xff, 0xd4,
	0xa7, 0xaa, 0x4f, 0x7d, 0xec, 0xbf, 0x54, 0xf9, 0xd8, 0x2c, 0xe3, 0xb9, 0xf0, 0xd8, 0xbe, 0x44,
	0x9b, 0x39, 0xdf, 0xf7, 0xf9, 0x3b, 0xc7, 0xc7, 0xc7, 0x06, 0x2d, 0xc7, 0xa2, 0xd7, 0x71, 0x45,
	0xd7, 0x1d, 0xf0, 0x44, 0xf6, 0x69, 0xe0, 0x06, 0x3c, 0xba, 0x11, 0xb5, 0x5e, 0x12, 0xcb, 0x18,
	0x07, 0x1e, 0x17, 0x5e, 0xec, 0xf2, 0x58, 0xb8, 0xb7, 0x89, 0xcb, 0x7b, 0x83, 0xb7, 0x2e, 0xc0,
	0xe3, 0x1e, 0x4b, 0x6a, 0xea, 0x97, 0xc2, 0x79, 0x4c, 0x08, 0x26, 0xee, 0x7e, 0xd5, 0x06, 0x49,
	0x07, 0xfe, 0xa9, 0xf1, 0x48, 0xb2, 0xa4, 0x43, 0x3d, 0xe6, 0xf2, 0xa8, 0x13, 0x27, 0x21, 0x95,
	0x3c, 0x8e, 0x6a, 0xf6, 0x32, 0xe9, 0xff, 0x6d, 0x7c, 0x73, 0xd0, 0x52, 0xa1, 0x1b, 0xf7, 0xec,
	0xf8, 0xeb, 0x25, 0x5e, 0x47, 0x53, 0x66, 0x0d, 0x37, 0xa2, 0x21, 0x23, 0x4e, 0xc5, 0xa9, 0x4e,
	0xb4, 0x26, 0xcd, 0xb7, 0x9f, 0x68, 0xc8, 0xf0, 0x22, 0x1a, 0x1f, 0x24, 0x1d, 0x1d, 0x1e, 0x83,
	0xf0, 0xd3, 0x41, 0xd2, 0x81, 0xd0, 0xff, 0xd1, 0x5c, 0x5a, 0x53, 0x63, 0x1e, 0x01, 0xe6, 0x99,
	0x09, 0x34, 0x79, 0x74, 0xa3, 0xb0, 0x1b, 0x7f, 0x4c, 0xa3, 0x85, 0x42, 0x23, 0xc5, 0x2a, 0x9b,
	0x85, 0x2a, 0xf8, 0x03, 0x22, 0x36, 0x96, 0xf1, 0xeb, 0x6e, 0x3b, 0x4e, 0x5c, 0xee, 0x93, 0x37,
	0x40, 0x59, 0x48, 0x53, 0x4c, 0xb4, 0xe1, 0xe3, 0xef, 0x10, 0xb6, 0x88, 0x42, 0x52, 0xc9, 0xc8,
	0x5b, 0xa0, 0xcc, 0xa6, 0x28, 0x97, 0xea, 0x3b, 0xde, 0x43, 0xcb, 0x16, 0xda, 0x67, 0x21, 0x8d,
	0x7c, 0xd7, 0xe3, 0x89, 0xd7, 0xe7, 0x92, 0xbc, 0xab, 0x38, 0xd5, 0xf1, 0x16, 0x49, 0xd1, 0x8e,
	0x00, 0x50, 0xd7, 0x71, 0xbc, 0x9b, 0xa5, 0x7b, 0x6e, 0x9b, 0xcb, 0x40, 0x55, 0x39, 0x10, 0x94,
	0xbc, 0xaf, 0x38, 0xd5, 0xe9, 0xd6, 0xcb, 0x34, 0xdd, 0x3b, 0xd4, 0xf1, 0xa6, 0xa0, 0x6a, 0x4f,
	0x64, 0x42, 0x23, 0xc1, 0xa5, 0x4b, 0x13, 0x46, 0xc9, 0x07, 0xbd, 0x27, 0xe6, 0xdb, 0x41, 0xc2,
	0x68, 0xce, 0xdf, 0x7d, 0x5b, 0x40, 0xf1, 0xb6, 0x80, 0x91, 0xf6, 0xd7, 0xb8, 0x03, 0x14, 0xee,
	0x9b, 0x17, 0x0b, 0x49, 0x3e, 0x82, 0xab, 0x74, 0xc5, 0xeb, 0xb1, 0x90, 0xb8, 0x8e, 0xd6, 0x06,
	0x3c, 0x19, 0x41, 0xc1, 0x45, 0xc8, 0x85, 0xe0, 0x71, 0xe4, 0xfa, 0x2c, 0xa0, 0x43, 0xb2, 0x0d,
	0xc4, 0x65, 0x8d, 0x52, 0xbc, 0xab, 0x14, 0xe6, 0x48, 0x41, 0x72, 0x7e, 0xbb, 0x2c, 0x08, 0x62,
	0xed, 0x7a, 0x40, 0x03, 0xb2, 0x03, 0x0a, 0x69, 0xbf, 0xa7, 0x0a, 0xd0, 0x30, 0x71, 0x7c, 0x08,
	0x1e, 0xca, 0xe8, 0x6e, 0x28, 0xc8, 0x2e, 0x28, 0x2c, 0x95, 0x29, 0x9c, 0x0b, 0xbc, 0x83, 0x96,
	0x32, 0x5b, 0x4a, 0xfd, 0x7b, 0x07, 0x7b, 0xf9, 0x2d, 0x61, 0xd4, 0x1f, 0x19, 0xc8, 0x92, 0x7f,
	0xa5, 0x5c, 0xde, 0x93, 0x3f, 0xe5, 0xc8, 0xbf, 0x50, 0x2e, 0x47, 0xe4, 0x0b, 0xf4, 0x5a, 0x85,
	0x68, 0xdf, 0x90, 0x13, 0x66, 0x15, 0x71, 0x24, 0xb3, 0x0f, 0x32, 0xeb, 0x06, 0xab, 0x64, 0x5a,
	0x16, 0x72, 0x24, 0xf8, 0x0e, 0xbd, 0xcc, 0x1c, 0x82, 0x5b, 0xa9, 0x6b, 0x42, 0x7e, 0x04, 0x8d,
	0xe7, 0xd6, 0x19, 0xb8, 0x95, 0x50, 0x0b, 0xbc, 0x9d, 0x49, 0xe2, 0x9e, 0xa6, 0x2a, 0x78, 0x00,
	0xcc, 0x17, 0x45, 0xcc, 0x73, 0x81, 0xbf, 0x47, 0xcf, 0x2d, 0x6e, 0x8f, 0x0a, 0xc1, 0x07, 0x8c,
	0x1c, 0xc2, 0x49, 0xc0, 0x29, 0xd6, 0xcf, 0x3a, 0x82, 0x3f, 0xa3, 0x8a, 0xc5, 0xa0, 0x7d, 0xd9,
	0x65, 0x91, 0xe4, 0x1e, 0x8c, 0x2c, 0x57, 0x0e, 0x7b, 0x8c, 0xd4, 0xa1, 0x4f, 0x57, 0x53, 0xec,
	0x03, 0x0b, 0x75, 0x35, 0xec, 0x31, 0xfc, 0x09, 0xad, 0x58, 0x42, 0xc3, 0xb8, 0x1f, 0x5d, 0x33,
	0x21, 0xdd, 0xd0, 0x77, 0x6f, 0xd8, 0x90, 0x1c, 0xe5, 0x0e, 0xe3, 0x57, 0x83, 0x38, 0xf7, 0xcf,
	0xd8, 0x10, 0xd7, 0xd1, 0xab, 0x87, 0xf8, 0x6a, 0x72, 0x1c, 0xe7, 0xba, 0xc7, 0x92, 0x68, 0xf8,
	0x78, 0x1f, 0xad, 0x5a, 0x22, 0x71, 0xe0, 0xdf, 0xf1, 0xbd, 0xb8, 0x1f, 0x49, 0x72, 0x92, 0x6b,
	0xe1, 0x8b, 0xc0, 0x07, 0x76, 0x5d, 0xc5, 0xf1, 0xfb, 0xcc, 0xe0, 0x32, 0xe4, 0x80, 0x0b, 0x49,
	0x3e, 0x57, 0x1e, 0x65, 0x36, 0x0d, 0x88, 0x4d, 0x2e, 0x64, 0x6e, 0xe0, 0xdd, 0xb0, 0xa1, 0xd7,
	0xa5, 0x3c, 0x52, 0xb6, 0x4f, 0x2b, 0x4e, 0xf5, 0xb1, 0x35, 0xf0, 0xce, 0x4c, 0xb4, 0xe1, 0xe7,
	0x27, 0xa5, 0xe8, 0xb8, 0x2c, 0xa2, 0xed, 0x80, 0xf9, 0xa4, 0x01, 0x25, 0xb3, 0x26, 0xa5, 0xe8,
	0x1c, 0xeb, 0x20, 0xae, 0xa2, 0xd9, 0x2c, 0x91, 0x7c, 0x01, 0xc2, 0x8c, 0x4d, 0xc0, 0x3f, 0xa0,
	0x05, 0x0b, 0x19, 0x50, 0x21, 0x01, 0x7e, 0x06, 0xc5, 0x48, 0x77, 0x45, 0x93, 0x0a, 0xa9, 0x28,
	0xbf, 0x3b, 0x19, 0xce, 0xdd, 0x00, 0x27, 0xcd, 0x8a, 0x53, 0x9d, 0xdc, 0xfc, 0xcd, 0xa9, 0xfd,
	0x9b, 0xd7, 0x63, 0x6d, 0x74, 0x23, 0x59, 0x66, 0x5a, 0xf3, 0x05, 0x97, 0xc8, 0xc6, 0x9f, 0x13,
	0x88, 0xdc, 0xe1, 0x47, 0xf7, 0x8e, 0x39, 0xcc, 0x78, 0x0b, 0x11, 0xbf, 0xed, 0x67, 0xcf, 0xb6,
	0xee, 0x0d, 0x47, 0x1f, 0x2d, 0xbf, 0xed, 0xdb, 0x07, 0x5a, 0x77, 0x46, 0x1d, 0xad, 0x15, 0x30,
	0x65, 0x2c, 0x69, 0x60, 0xf8, 0x63, 0x7a, 0xc0, 0xe6, 0xf8, 0x57, 0x0a, 0xa3, 0x45, 0x6a, 0x68,
	0x5e, 0xdd, 0x15, 0x6e, 0x27, 0x88, 0x63, 0x9f, 0x47, 0xd7, 0x2e, 0x8f, 0x7c, 0x76, 0x0b, 0x77,
	0xf1, 0x74, 0x6b, 0x4e, 0x85, 0x4e, 0x4c, 0xa4, 0xa1, 0x02, 0xf8, 0x35, 0x9a, 0xa1, 0x42, 0xa3,
	0x0d, 0xf4, 0x31, 0x40, 0xa7, 0xa8, 0x00, 0xa0, 0x46, 0x1d, 0xa2, 0xd5, 0x6c, 0xa2, 0x76, 0x66,
	0xff, 0xd3, 0xce, 0x46, 0x75, 0x2b, 0x48, 0x6f, 0x1b, 0x2d, 0x46, 0xfd, 0xb0, 0xcd, 0x12, 0x37,
	0xee, 0x64, 0x44, 0x04, 0x79, 0xa2, 0x27, 0xa7, 0x06, 0x5c, 0x74, 0x6c, 0xbe, 0xc0, 0xfb, 0x68,
	0x45, 0x67, 0xc5, 0x13, 0x21, 0x47, 0x6e, 0x47, 0x1b, 0x4b, 0x9e, 0x02, 0x7d, 0x11, 0xd2, 0x53,
	0x10, 0x63, 0x7d, 0x04, 0xc0, 0x5f, 0xd0, 0xc6, 0x43, 0x02, 0x26, 0xf5, 0x71, 0x90, 0x59, 0x2b,
	0x95, 0xd1, 0xc5, 0xd8, 0x41, 0x4b, 0xaa, 0x64, 0x25, 0x56, 0x26, 0x74, 0x26, 0x54, 0x14, 0x1b,
	0x39, 0x41, 0x95, 0x72, 0xb2, 0xb1, 0x81, 0x40, 0x62, 0xa5, 0x44, 0x42, 0x9b, 0xd8, 0x43, 0xcb,
	0x90, 0x10, 0xcc, 0xee, 0xbc, 0x8b, 0x49, 0x3d, 0x85, 0x14, 0x44, 0x4d, 0xef, 0x9c, 0x8d, 0x53,
	0xb4, 0xfe, 0x00, 0xdd, 0xf8, 0x98, 0x02, 0x91, 0xd5, 0x32, 0x11, 0x6d, 0xe4, 0x23, 0x5a, 0xa4,
	0xa2, 0xcc, 0xc6, 0xb4, 0x6e, 0x78, 0x2a, 0x0a, 0x4d, 0x1c, 0xa1, 0x57, 0xa5, 0x54, 0x63, 0x61,
	0x46, 0xf7, 0x55, 0xb1, 0x80, 0x36, 0xb0, 0x8b, 0x96, 0x60, 0xde, 0x64, 0xfa, 0x32, 0x60, 0xd1,
	0xb5, 0xec, 0x92, 0x67, 0xba, 0x10, 0x0a, 0x61, 0x37, 0x55, 0x13, 0xe2, 0xaa, 0xb3, 0x43, 0x7a,
	0xcb, 0xc3, 0x7e, 0x58, 0x22, 0x30, 0xab, 0x1d, 0x18, 0x50, 0xa1, 0xc6, 0x16, 0x22, 0x45, 0x0e,
	0x24, 0x0f, 0x19, 0x99, 0xd3, 0x15, 0xc8, 0xaf, 0x7f, 0xc5, 0x43, 0x75, 0xa5, 0x2d, 0x97, 0xac,
	0x0e, 0x64, 0xac, 0xdb, 0xba, 0x70, 0x6d, 0xe0, 0xab, 0x95, 0x05, 0x2d, 0xe2, 0x26, 0x64, 0xde,
	0xac, 0x2c, 0x68, 0x9e, 0x98, 0x6c, 0xfc, 0x3d, 0x86, 0x5e, 0x14, 0xcf, 0xbc, 0xdc, 0x1b, 0x4d,
	0xf4, 0x7b, 0xbd, 0x44, 0x3d, 0x59, 0xf5, 0xcb, 0xc2, 0xc9, 0x5d, 0xb3, 0x97, 0x06, 0xa0, 0x5f,
	0x17, 0xc5, 0x0f, 0xec, 0xb1, 0x92, 0x07, 0xf6, 0x5f, 0x4e, 0x66, 0xb5, 0x54, 0x2e, 0x31, 0x8f,
	0x60, 0x70, 0x4d, 0x6e, 0x7e, 0xfb, 0x8f, 0x6e, 0x83, 0xec, 0xd0, 0x6b, 0x2d, 0xa6, 0x12, 0x68,
	0x59, 0x4e, 0xdb, 0x4f, 0xe0, 0xaf, 0xba, 0x37, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x37,
	0xb2, 0x4d, 0xf4, 0x0d, 0x00, 0x00,
}
