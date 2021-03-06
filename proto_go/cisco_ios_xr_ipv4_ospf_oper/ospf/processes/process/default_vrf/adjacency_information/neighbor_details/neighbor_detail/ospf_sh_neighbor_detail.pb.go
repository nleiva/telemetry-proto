// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_neighbor_detail.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_neighbor_details_neighbor_detail

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

// OSPF Neighbor Detailed Information
type OspfShNeighborDetail_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighborDetail_KEYS) Reset()         { *m = OspfShNeighborDetail_KEYS{} }
func (m *OspfShNeighborDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborDetail_KEYS) ProtoMessage()    {}
func (*OspfShNeighborDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e, []int{0}
}
func (m *OspfShNeighborDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborDetail_KEYS.Unmarshal(m, b)
}
func (m *OspfShNeighborDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighborDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborDetail_KEYS.Merge(dst, src)
}
func (m *OspfShNeighborDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborDetail_KEYS.Size(m)
}
func (m *OspfShNeighborDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborDetail_KEYS proto.InternalMessageInfo

func (m *OspfShNeighborDetail_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShNeighborDetail_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShNeighborDetail_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShNeighborDetail struct {
	// Neighbor summary information
	NeighborSummary *OspfShNeighbor `protobuf:"bytes,50,opt,name=neighbor_summary,json=neighborSummary,proto3" json:"neighbor_summary,omitempty"`
	// Area ID string in decimal or dotted-decimal format
	NeighborAreaId string `protobuf:"bytes,51,opt,name=neighbor_area_id,json=neighborAreaId,proto3" json:"neighbor_area_id,omitempty"`
	// Number of state changes
	StateChangeCount uint32 `protobuf:"varint,52,opt,name=state_change_count,json=stateChangeCount,proto3" json:"state_change_count,omitempty"`
	// Cost of path to this neighbor
	NeighborCost uint32 `protobuf:"varint,53,opt,name=neighbor_cost,json=neighborCost,proto3" json:"neighbor_cost,omitempty"`
	// If true, filter outgoing LSAs
	NeighborFilter bool `protobuf:"varint,54,opt,name=neighbor_filter,json=neighborFilter,proto3" json:"neighbor_filter,omitempty"`
	// Address of designated router
	NeighborDesignatedRouterAddress string `protobuf:"bytes,55,opt,name=neighbor_designated_router_address,json=neighborDesignatedRouterAddress,proto3" json:"neighbor_designated_router_address,omitempty"`
	// Address of backup designated router
	NeighborBackupDesignatedRouterAddress string `protobuf:"bytes,56,opt,name=neighbor_backup_designated_router_address,json=neighborBackupDesignatedRouterAddress,proto3" json:"neighbor_backup_designated_router_address,omitempty"`
	// Interface type
	InterfaceType string `protobuf:"bytes,57,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	// Poll interval (s)
	PollInterval uint32 `protobuf:"varint,58,opt,name=poll_interval,json=pollInterval,proto3" json:"poll_interval,omitempty"`
	// For NBMA networks, amount of time remaining in seconds before the next poll interval expires and Hello is sent (s)
	NextPollInterval uint32 `protobuf:"varint,59,opt,name=next_poll_interval,json=nextPollInterval,proto3" json:"next_poll_interval,omitempty"`
	//  This is bitmaks of neighbor's option field received
	NeighborOption uint32 `protobuf:"varint,60,opt,name=neighbor_option,json=neighborOption,proto3" json:"neighbor_option,omitempty"`
	// Number of pending events
	PendingEvents uint32 `protobuf:"varint,61,opt,name=pending_events,json=pendingEvents,proto3" json:"pending_events,omitempty"`
	// This is a bitmask of Link Local signalling options received from the neighbor
	NeighborLlsOption uint32 `protobuf:"varint,62,opt,name=neighbor_lls_option,json=neighborLlsOption,proto3" json:"neighbor_lls_option,omitempty"`
	// Out-Of-Bound resynchronization in progress
	OobResynchronization bool `protobuf:"varint,63,opt,name=oob_resynchronization,json=oobResynchronization,proto3" json:"oob_resynchronization,omitempty"`
	// For cisco NSF, the router is either Requester or Receiver
	NsfRouterState string `protobuf:"bytes,64,opt,name=nsf_router_state,json=nsfRouterState,proto3" json:"nsf_router_state,omitempty"`
	// The amount of time in seconds since last time Out-Of-Band resynchronization was done with this neighbor
	LastOobTime uint32 `protobuf:"varint,65,opt,name=last_oob_time,json=lastOobTime,proto3" json:"last_oob_time,omitempty"`
	// Neighbor BFD information
	NeighborBfdInformation *OspfShNeighborBfd `protobuf:"bytes,66,opt,name=neighbor_bfd_information,json=neighborBfdInformation,proto3" json:"neighbor_bfd_information,omitempty"`
	// Retransmission information with this neighbor
	NeighborRetransmissionInformation *OspfShNeighborRetrans `protobuf:"bytes,67,opt,name=neighbor_retransmission_information,json=neighborRetransmissionInformation,proto3" json:"neighbor_retransmission_information,omitempty"`
	// LFA Interface
	LfaInterface string `protobuf:"bytes,68,opt,name=lfa_interface,json=lfaInterface,proto3" json:"lfa_interface,omitempty"`
	// LFA Next Hop
	LfaNextHop string `protobuf:"bytes,69,opt,name=lfa_next_hop,json=lfaNextHop,proto3" json:"lfa_next_hop,omitempty"`
	// LFA Neighbor ID
	LfaNeighborId string `protobuf:"bytes,70,opt,name=lfa_neighbor_id,json=lfaNeighborId,proto3" json:"lfa_neighbor_id,omitempty"`
	// LFA Neighbor Revision
	LfaNeighborRevision uint32 `protobuf:"varint,71,opt,name=lfa_neighbor_revision,json=lfaNeighborRevision,proto3" json:"lfa_neighbor_revision,omitempty"`
	// Ack List Count
	NeighborAckListCount uint32 `protobuf:"varint,72,opt,name=neighbor_ack_list_count,json=neighborAckListCount,proto3" json:"neighbor_ack_list_count,omitempty"`
	// Ack List High Watermark
	NeighborAckListHighWatermark uint32 `protobuf:"varint,73,opt,name=neighbor_ack_list_high_watermark,json=neighborAckListHighWatermark,proto3" json:"neighbor_ack_list_high_watermark,omitempty"`
	// SR Adjacency SID Label
	AdjacencySidLabel uint32 `protobuf:"varint,74,opt,name=adjacency_sid_label,json=adjacencySidLabel,proto3" json:"adjacency_sid_label,omitempty"`
	// SR Adjacency SID Protected
	AdjacencySidProtected bool `protobuf:"varint,75,opt,name=adjacency_sid_protected,json=adjacencySidProtected,proto3" json:"adjacency_sid_protected,omitempty"`
	// SR Adjacency SID Unprotected Label
	AdjacencySidUnprotectedLabel uint32   `protobuf:"varint,76,opt,name=adjacency_sid_unprotected_label,json=adjacencySidUnprotectedLabel,proto3" json:"adjacency_sid_unprotected_label,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *OspfShNeighborDetail) Reset()         { *m = OspfShNeighborDetail{} }
func (m *OspfShNeighborDetail) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborDetail) ProtoMessage()    {}
func (*OspfShNeighborDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e, []int{1}
}
func (m *OspfShNeighborDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborDetail.Unmarshal(m, b)
}
func (m *OspfShNeighborDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborDetail.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighborDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborDetail.Merge(dst, src)
}
func (m *OspfShNeighborDetail) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborDetail.Size(m)
}
func (m *OspfShNeighborDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborDetail.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborDetail proto.InternalMessageInfo

func (m *OspfShNeighborDetail) GetNeighborSummary() *OspfShNeighbor {
	if m != nil {
		return m.NeighborSummary
	}
	return nil
}

func (m *OspfShNeighborDetail) GetNeighborAreaId() string {
	if m != nil {
		return m.NeighborAreaId
	}
	return ""
}

func (m *OspfShNeighborDetail) GetStateChangeCount() uint32 {
	if m != nil {
		return m.StateChangeCount
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborCost() uint32 {
	if m != nil {
		return m.NeighborCost
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborFilter() bool {
	if m != nil {
		return m.NeighborFilter
	}
	return false
}

func (m *OspfShNeighborDetail) GetNeighborDesignatedRouterAddress() string {
	if m != nil {
		return m.NeighborDesignatedRouterAddress
	}
	return ""
}

func (m *OspfShNeighborDetail) GetNeighborBackupDesignatedRouterAddress() string {
	if m != nil {
		return m.NeighborBackupDesignatedRouterAddress
	}
	return ""
}

func (m *OspfShNeighborDetail) GetInterfaceType() string {
	if m != nil {
		return m.InterfaceType
	}
	return ""
}

func (m *OspfShNeighborDetail) GetPollInterval() uint32 {
	if m != nil {
		return m.PollInterval
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNextPollInterval() uint32 {
	if m != nil {
		return m.NextPollInterval
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborOption() uint32 {
	if m != nil {
		return m.NeighborOption
	}
	return 0
}

func (m *OspfShNeighborDetail) GetPendingEvents() uint32 {
	if m != nil {
		return m.PendingEvents
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborLlsOption() uint32 {
	if m != nil {
		return m.NeighborLlsOption
	}
	return 0
}

func (m *OspfShNeighborDetail) GetOobResynchronization() bool {
	if m != nil {
		return m.OobResynchronization
	}
	return false
}

func (m *OspfShNeighborDetail) GetNsfRouterState() string {
	if m != nil {
		return m.NsfRouterState
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLastOobTime() uint32 {
	if m != nil {
		return m.LastOobTime
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborBfdInformation() *OspfShNeighborBfd {
	if m != nil {
		return m.NeighborBfdInformation
	}
	return nil
}

func (m *OspfShNeighborDetail) GetNeighborRetransmissionInformation() *OspfShNeighborRetrans {
	if m != nil {
		return m.NeighborRetransmissionInformation
	}
	return nil
}

func (m *OspfShNeighborDetail) GetLfaInterface() string {
	if m != nil {
		return m.LfaInterface
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLfaNextHop() string {
	if m != nil {
		return m.LfaNextHop
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLfaNeighborId() string {
	if m != nil {
		return m.LfaNeighborId
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLfaNeighborRevision() uint32 {
	if m != nil {
		return m.LfaNeighborRevision
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborAckListCount() uint32 {
	if m != nil {
		return m.NeighborAckListCount
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborAckListHighWatermark() uint32 {
	if m != nil {
		return m.NeighborAckListHighWatermark
	}
	return 0
}

func (m *OspfShNeighborDetail) GetAdjacencySidLabel() uint32 {
	if m != nil {
		return m.AdjacencySidLabel
	}
	return 0
}

func (m *OspfShNeighborDetail) GetAdjacencySidProtected() bool {
	if m != nil {
		return m.AdjacencySidProtected
	}
	return false
}

func (m *OspfShNeighborDetail) GetAdjacencySidUnprotectedLabel() uint32 {
	if m != nil {
		return m.AdjacencySidUnprotectedLabel
	}
	return 0
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
	return fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e, []int{2}
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

// OSPF Neighbor BFD information
type OspfShNeighborBfd struct {
	// BFD enable mode - Default/Strict
	BfdIntfEnableMode uint32 `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode,proto3" json:"bfd_intf_enable_mode,omitempty"`
	// Status of the BFD Session
	BfdStatusFlag        uint32   `protobuf:"varint,2,opt,name=bfd_status_flag,json=bfdStatusFlag,proto3" json:"bfd_status_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighborBfd) Reset()         { *m = OspfShNeighborBfd{} }
func (m *OspfShNeighborBfd) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborBfd) ProtoMessage()    {}
func (*OspfShNeighborBfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e, []int{3}
}
func (m *OspfShNeighborBfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborBfd.Unmarshal(m, b)
}
func (m *OspfShNeighborBfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborBfd.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighborBfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborBfd.Merge(dst, src)
}
func (m *OspfShNeighborBfd) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborBfd.Size(m)
}
func (m *OspfShNeighborBfd) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborBfd.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborBfd proto.InternalMessageInfo

func (m *OspfShNeighborBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *OspfShNeighborBfd) GetBfdStatusFlag() uint32 {
	if m != nil {
		return m.BfdStatusFlag
	}
	return 0
}

// OSPF Neighbor Summary Information
type OspfShNeighbor struct {
	// Neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId,proto3" json:"neighbor_id,omitempty"`
	// Neighbor IP Address
	NeighborAddress string `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	// Interface via which the neighbor is seen
	NeighborInterfaceName string `protobuf:"bytes,3,opt,name=neighbor_interface_name,json=neighborInterfaceName,proto3" json:"neighbor_interface_name,omitempty"`
	// Neighbor's DR priority
	NeighborDrPriority uint32 `protobuf:"varint,4,opt,name=neighbor_dr_priority,json=neighborDrPriority,proto3" json:"neighbor_dr_priority,omitempty"`
	// Neighbor's state
	NeighborState string `protobuf:"bytes,5,opt,name=neighbor_state,json=neighborState,proto3" json:"neighbor_state,omitempty"`
	// Designated Router state
	DrBdrState string `protobuf:"bytes,6,opt,name=dr_bdr_state,json=drBdrState,proto3" json:"dr_bdr_state,omitempty"`
	// Time until neighbor's dead timer expires (s)
	NeighborDeadTimer uint32 `protobuf:"varint,7,opt,name=neighbor_dead_timer,json=neighborDeadTimer,proto3" json:"neighbor_dead_timer,omitempty"`
	// Amount of time since the adjacency is up (s)
	NeighborUpTime uint32 `protobuf:"varint,8,opt,name=neighbor_up_time,json=neighborUpTime,proto3" json:"neighbor_up_time,omitempty"`
	// Interface is MADJ
	NeighborMadjInterface bool `protobuf:"varint,9,opt,name=neighbor_madj_interface,json=neighborMadjInterface,proto3" json:"neighbor_madj_interface,omitempty"`
	// Neighbor BFD information
	NeighborBfdInformation *OspfShNeighborBfd `protobuf:"bytes,10,opt,name=neighbor_bfd_information,json=neighborBfdInformation,proto3" json:"neighbor_bfd_information,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *OspfShNeighbor) Reset()         { *m = OspfShNeighbor{} }
func (m *OspfShNeighbor) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighbor) ProtoMessage()    {}
func (*OspfShNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e, []int{4}
}
func (m *OspfShNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighbor.Unmarshal(m, b)
}
func (m *OspfShNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighbor.Marshal(b, m, deterministic)
}
func (dst *OspfShNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighbor.Merge(dst, src)
}
func (m *OspfShNeighbor) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighbor.Size(m)
}
func (m *OspfShNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighbor proto.InternalMessageInfo

func (m *OspfShNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborInterfaceName() string {
	if m != nil {
		return m.NeighborInterfaceName
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDrPriority() uint32 {
	if m != nil {
		return m.NeighborDrPriority
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *OspfShNeighbor) GetDrBdrState() string {
	if m != nil {
		return m.DrBdrState
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDeadTimer() uint32 {
	if m != nil {
		return m.NeighborDeadTimer
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborUpTime() uint32 {
	if m != nil {
		return m.NeighborUpTime
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborMadjInterface() bool {
	if m != nil {
		return m.NeighborMadjInterface
	}
	return false
}

func (m *OspfShNeighbor) GetNeighborBfdInformation() *OspfShNeighborBfd {
	if m != nil {
		return m.NeighborBfdInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShNeighborDetail_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_detail_KEYS")
	proto.RegisterType((*OspfShNeighborDetail)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_detail")
	proto.RegisterType((*OspfShNeighborRetrans)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_retrans")
	proto.RegisterType((*OspfShNeighborBfd)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_bfd")
	proto.RegisterType((*OspfShNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor")
}

func init() {
	proto.RegisterFile("ospf_sh_neighbor_detail.proto", fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e)
}

var fileDescriptor_ospf_sh_neighbor_detail_7d195356d2c2101e = []byte{
	// 1380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xc9, 0x72, 0x1b, 0x37,
	0x10, 0x2d, 0x3a, 0xde, 0x04, 0x89, 0xb2, 0x0c, 0x4b, 0x16, 0xbc, 0x8a, 0xa6, 0xe3, 0x44, 0xae,
	0x4a, 0x31, 0x29, 0x6f, 0xf1, 0xee, 0x58, 0x5b, 0x24, 0x5b, 0x5e, 0x8a, 0xb2, 0x2b, 0xc9, 0x09,
	0x05, 0x12, 0x18, 0x12, 0xd6, 0xcc, 0x60, 0x0a, 0x00, 0x15, 0x29, 0x3f, 0x90, 0x43, 0x0e, 0x39,
	0xa4, 0x2a, 0x3f, 0x92, 0x2f, 0xc8, 0x25, 0xdf, 0x91, 0x0f, 0xc8, 0x47, 0xa4, 0xd0, 0x00, 0x67,
	0xa1, 0x48, 0x5d, 0xe3, 0xdc, 0xc8, 0xee, 0xd7, 0x8d, 0xd7, 0x3d, 0x8d, 0xee, 0x06, 0xba, 0xa2,
	0x4c, 0x16, 0x51, 0xd3, 0xa7, 0xa9, 0x90, 0xbd, 0x7e, 0x47, 0x69, 0xca, 0x85, 0x65, 0x32, 0x6e,
	0x65, 0x5a, 0x59, 0x85, 0x07, 0x5d, 0x69, 0xba, 0x8a, 0x4a, 0x65, 0xe8, 0xbe, 0xa6, 0x32, 0xdb,
	0xbb, 0x43, 0xc1, 0x40, 0x65, 0x42, 0xb7, 0xdc, 0x2f, 0x87, 0xeb, 0x0a, 0x63, 0x84, 0x19, 0xfe,
	0x6a, 0x71, 0x11, 0xb1, 0x41, 0x6c, 0xe9, 0x9e, 0x8e, 0x5a, 0x8c, 0x7f, 0x60, 0x5d, 0x91, 0x76,
	0x0f, 0xa8, 0x4c, 0x23, 0xa5, 0x13, 0x66, 0xa5, 0x4a, 0x5b, 0x23, 0x67, 0x99, 0x51, 0x41, 0xf3,
	0xb7, 0x1a, 0xba, 0x3c, 0x81, 0x18, 0x7d, 0xb9, 0xfe, 0xc3, 0x0e, 0xbe, 0x86, 0x66, 0xc2, 0x71,
	0x34, 0x65, 0x89, 0x20, 0xb5, 0x46, 0x6d, 0x79, 0xaa, 0x3d, 0x1d, 0x64, 0xaf, 0x59, 0x22, 0xf0,
	0x0d, 0x34, 0x2b, 0x53, 0x2b, 0x74, 0xc4, 0xba, 0xc2, 0x83, 0x8e, 0x01, 0xa8, 0x9e, 0x4b, 0x01,
	0x76, 0x13, 0xcd, 0xe5, 0x27, 0x30, 0xce, 0xb5, 0x30, 0x86, 0x7c, 0x02, 0xc0, 0x33, 0x43, 0xf9,
	0x73, 0x2f, 0x6e, 0xfe, 0x3e, 0x8b, 0x16, 0x27, 0xb0, 0xc2, 0x7f, 0xd4, 0x4a, 0x7e, 0xcc, 0x20,
	0x49, 0x98, 0x3e, 0x20, 0xb7, 0x1a, 0xb5, 0xe5, 0xe9, 0x5b, 0x3f, 0xd7, 0x5a, 0xff, 0x49, 0x16,
	0x5b, 0xa3, 0x5c, 0x8b, 0x88, 0x76, 0x3c, 0x41, 0xbc, 0x5c, 0x0e, 0x5e, 0x0b, 0x46, 0x25, 0x27,
	0xb7, 0x21, 0xf8, 0xd9, 0x3c, 0x78, 0x2d, 0xd8, 0x16, 0xc7, 0x5f, 0x20, 0x6c, 0x2c, 0xb3, 0x82,
	0x76, 0xfb, 0x2c, 0xed, 0x09, 0xda, 0x55, 0x83, 0xd4, 0x92, 0x3b, 0x8d, 0xda, 0x72, 0xbd, 0x3d,
	0x07, 0x9a, 0x55, 0x50, 0xac, 0x3a, 0x39, 0xbe, 0x8e, 0xea, 0xb9, 0xdf, 0xae, 0x32, 0x96, 0xdc,
	0x05, 0xe0, 0xcc, 0x50, 0xb8, 0xaa, 0x8c, 0xc5, 0x9f, 0xa3, 0x9c, 0x0f, 0x8d, 0x64, 0x6c, 0x85,
	0x26, 0xf7, 0x1a, 0xb5, 0xe5, 0xd3, 0xc5, 0xd9, 0x1b, 0x20, 0xc5, 0x2f, 0x51, 0xb3, 0x14, 0x9a,
	0x91, 0xbd, 0x94, 0x59, 0xc1, 0xa9, 0x56, 0x03, 0x2b, 0x8a, 0x8f, 0xf6, 0x35, 0xf0, 0x5e, 0x1a,
	0x22, 0xd7, 0x72, 0x60, 0x1b, 0x70, 0xe1, 0x23, 0xe2, 0xef, 0xd1, 0xcd, 0xdc, 0x59, 0x87, 0x75,
	0x77, 0x07, 0xd9, 0x11, 0x3e, 0xef, 0x83, 0xcf, 0x1b, 0x43, 0x83, 0x15, 0xc0, 0x4f, 0xf2, 0x5c,
	0x29, 0x38, 0x7b, 0x90, 0x09, 0xf2, 0x60, 0xa4, 0xe0, 0xde, 0x1d, 0x64, 0xc2, 0xe5, 0x26, 0x53,
	0x71, 0x4c, 0x41, 0xba, 0xc7, 0x62, 0xf2, 0xd0, 0xe7, 0xc6, 0x09, 0xb7, 0x82, 0xcc, 0xa5, 0x3b,
	0x15, 0xfb, 0x96, 0x56, 0x91, 0x8f, 0x7c, 0xba, 0x9d, 0xe6, 0x6d, 0x19, 0x5d, 0xce, 0xa4, 0xca,
	0x5c, 0x71, 0x90, 0xc7, 0x00, 0xcd, 0x33, 0xf9, 0x06, 0xa4, 0x8e, 0x62, 0x26, 0x52, 0x2e, 0xd3,
	0x1e, 0x15, 0x7b, 0x22, 0xb5, 0x86, 0x3c, 0x01, 0x5c, 0x3d, 0x48, 0xd7, 0x41, 0x88, 0x5b, 0xe8,
	0x5c, 0xee, 0x2f, 0x8e, 0xcd, 0xd0, 0xe7, 0x53, 0xc0, 0x9e, 0x1d, 0xaa, 0xb6, 0x63, 0x13, 0xdc,
	0xde, 0x46, 0x0b, 0x4a, 0x75, 0xa8, 0x16, 0xe6, 0x20, 0xed, 0xf6, 0xb5, 0x4a, 0xe5, 0x4f, 0x50,
	0xa2, 0xe4, 0x19, 0x7c, 0xcf, 0x79, 0xa5, 0x3a, 0xed, 0x51, 0x1d, 0xd4, 0x9e, 0x89, 0x86, 0x19,
	0x87, 0x12, 0x22, 0xdf, 0x84, 0xda, 0x33, 0x91, 0x4f, 0xed, 0x8e, 0x93, 0xe2, 0x26, 0xaa, 0xc7,
	0xcc, 0x58, 0xea, 0xce, 0xb0, 0x32, 0x11, 0xe4, 0x39, 0x10, 0x99, 0x76, 0xc2, 0x37, 0xaa, 0xf3,
	0x4e, 0x26, 0x02, 0xff, 0x55, 0x43, 0xa4, 0xf8, 0xae, 0x11, 0x2f, 0xdf, 0x14, 0xb2, 0x02, 0xf7,
	0xf0, 0x97, 0x8f, 0xe5, 0x1e, 0x3a, 0x7e, 0xed, 0xf3, 0x79, 0x51, 0x45, 0x7c, 0xab, 0xf0, 0x86,
	0xff, 0xae, 0xa1, 0xeb, 0x39, 0x50, 0x0b, 0xab, 0x59, 0x6a, 0x12, 0x69, 0x8c, 0x54, 0x69, 0x25,
	0xa6, 0x55, 0x88, 0xe9, 0xd7, 0x8f, 0x26, 0xa6, 0x40, 0xb5, 0x7d, 0x2d, 0xef, 0x36, 0x15, 0xee,
	0xe5, 0x10, 0xaf, 0xa3, 0x7a, 0x1c, 0x31, 0x9a, 0x5f, 0x0b, 0xb2, 0x06, 0x9f, 0x7d, 0x26, 0x8e,
	0xd8, 0xd6, 0x50, 0x86, 0x1b, 0xc8, 0xfd, 0xa7, 0x70, 0x0b, 0xfa, 0x2a, 0x23, 0xeb, 0x80, 0x41,
	0x71, 0xc4, 0x5e, 0x8b, 0x7d, 0xbb, 0xa9, 0x32, 0xfc, 0x19, 0x3a, 0xe3, 0x11, 0x81, 0x81, 0xe4,
	0x64, 0xc3, 0x5f, 0x38, 0x00, 0x79, 0xe9, 0x16, 0xc7, 0xb7, 0xd0, 0x42, 0x05, 0xa7, 0xc5, 0x9e,
	0x74, 0x94, 0xc8, 0xb7, 0x50, 0x46, 0xe7, 0x4a, 0xe8, 0x76, 0x50, 0xe1, 0xbb, 0x68, 0xb1, 0x68,
	0x8c, 0xdd, 0x5d, 0x1a, 0x4b, 0x63, 0x43, 0xcf, 0xdb, 0x04, 0xab, 0xf9, 0xbc, 0x3f, 0x76, 0x77,
	0xb7, 0xa5, 0xb1, 0xbe, 0xef, 0x6d, 0xa0, 0xc6, 0x61, 0xb3, 0xbe, 0xec, 0xf5, 0xe9, 0x8f, 0xcc,
	0x0a, 0x9d, 0x30, 0xbd, 0x4b, 0xb6, 0xc0, 0xfe, 0xf2, 0x88, 0xfd, 0xa6, 0xec, 0xf5, 0xbf, 0x1b,
	0x62, 0xdc, 0x05, 0x2c, 0xbe, 0x8b, 0x91, 0x9c, 0xc6, 0xac, 0x23, 0x62, 0xf2, 0xc2, 0x5f, 0xc0,
	0x5c, 0xb5, 0x23, 0xf9, 0xb6, 0x53, 0xe0, 0x7b, 0x68, 0xb1, 0x8a, 0x77, 0xd3, 0x5b, 0x74, 0xad,
	0xe0, 0xe4, 0x25, 0x5c, 0xc1, 0x85, 0xb2, 0xcd, 0xdb, 0xa1, 0x12, 0xaf, 0xa3, 0xa5, 0xaa, 0xdd,
	0x20, 0xcd, 0x2d, 0xc3, 0x99, 0xdb, 0x9e, 0x6e, 0xd9, 0xfe, 0x7d, 0x01, 0x82, 0xe3, 0x9b, 0x7f,
	0x4e, 0x21, 0x32, 0xa9, 0x20, 0xf0, 0x7d, 0x44, 0x78, 0x87, 0x8f, 0x96, 0xb2, 0xcf, 0x65, 0x0d,
	0x9c, 0x9f, 0xe7, 0x1d, 0x5e, 0xad, 0x16, 0x9f, 0xcd, 0x55, 0x74, 0x75, 0x8c, 0xa5, 0x55, 0x96,
	0xc5, 0xc1, 0xfe, 0x18, 0xd8, 0x5f, 0x3a, 0x64, 0xff, 0xce, 0x61, 0xbc, 0x13, 0x97, 0x4a, 0x37,
	0xd9, 0xa2, 0x58, 0x29, 0x68, 0x7c, 0x32, 0xe5, 0x62, 0x1f, 0x46, 0xbc, 0x4b, 0xa5, 0x16, 0x6c,
	0x23, 0x68, 0xb6, 0x9c, 0x02, 0x7f, 0x8a, 0x66, 0x99, 0xf1, 0xe8, 0x00, 0x3d, 0xee, 0xfb, 0x33,
	0x33, 0x00, 0xf4, 0xa8, 0x15, 0x74, 0x65, 0xd2, 0x25, 0xf5, 0xcc, 0x4e, 0x78, 0x66, 0xe3, 0x2f,
	0x83, 0x67, 0xf6, 0x10, 0x5d, 0x48, 0x07, 0x49, 0x47, 0x68, 0xaa, 0xa2, 0x11, 0x27, 0x86, 0x9c,
	0x04, 0xfb, 0x45, 0x0f, 0x78, 0x13, 0x55, 0xed, 0x0d, 0x7e, 0x86, 0x2e, 0xfb, 0xa8, 0xa4, 0x36,
	0x36, 0x67, 0x5b, 0x74, 0x87, 0x53, 0x60, 0x7e, 0x01, 0xc2, 0x73, 0x90, 0x40, 0xbd, 0xb8, 0x83,
	0x2f, 0x50, 0xf3, 0x28, 0x07, 0x21, 0xf4, 0xd3, 0xe0, 0xe6, 0xea, 0x44, 0x37, 0x3e, 0x19, 0x8f,
	0xd0, 0x45, 0x97, 0xb2, 0x09, 0x54, 0xa6, 0x7c, 0x24, 0xcc, 0x8c, 0x27, 0xb2, 0x81, 0x1a, 0x93,
	0x8d, 0x03, 0x0d, 0x14, 0x6a, 0xd0, 0x1c, 0x41, 0xe2, 0x09, 0xba, 0x04, 0x01, 0x41, 0xc3, 0x38,
	0xcc, 0x62, 0x1a, 0x5c, 0x10, 0x07, 0x71, 0xfd, 0xe3, 0x10, 0x8d, 0x4d, 0x74, 0xed, 0x08, 0xf3,
	0xc0, 0x63, 0x06, 0x9c, 0x5c, 0x99, 0xe4, 0xc4, 0x13, 0x79, 0x80, 0x2e, 0x30, 0x33, 0x89, 0x46,
	0xdd, 0x17, 0x3c, 0x33, 0x63, 0x49, 0xac, 0xa1, 0xa5, 0x89, 0xa6, 0x81, 0xc2, 0xac, 0xaf, 0xab,
	0xf1, 0x0e, 0x3c, 0x81, 0xc7, 0xe8, 0x22, 0x8c, 0xcb, 0x91, 0xba, 0x8c, 0x45, 0xda, 0xb3, 0x7d,
	0x72, 0xc6, 0x27, 0xc2, 0x21, 0xaa, 0x45, 0xb5, 0x0d, 0x7a, 0x57, 0xd9, 0x09, 0xdb, 0x97, 0xc9,
	0x20, 0x99, 0xe0, 0x60, 0xce, 0x33, 0x08, 0xa0, 0xb1, 0x3e, 0xee, 0x23, 0x32, 0x8e, 0x01, 0xcc,
	0xee, 0xb3, 0x3e, 0x03, 0x87, 0xcf, 0x87, 0x31, 0xfe, 0x14, 0x5d, 0x9a, 0x70, 0x3a, 0x18, 0x63,
	0x5f, 0xd6, 0x63, 0xcf, 0x06, 0x7b, 0x77, 0xb2, 0x61, 0xe3, 0x6c, 0x35, 0x39, 0x17, 0x4e, 0x36,
	0xec, 0xb0, 0xa1, 0x6e, 0x2a, 0x34, 0x3f, 0x6e, 0x4e, 0xe3, 0x2f, 0xd1, 0xbc, 0x5f, 0x27, 0x6c,
	0x44, 0x45, 0xca, 0x3a, 0xb1, 0xa0, 0x89, 0xe2, 0x22, 0xb4, 0xae, 0xb3, 0x1d, 0x37, 0xbd, 0x6d,
	0xb4, 0x0e, 0x9a, 0x57, 0x8a, 0x0b, 0x37, 0x96, 0x9c, 0x81, 0x5b, 0x68, 0x06, 0xae, 0x91, 0xb0,
	0x5e, 0x68, 0x53, 0xf5, 0x4e, 0xc4, 0x77, 0x40, 0xba, 0x11, 0xb3, 0x5e, 0xf3, 0x9f, 0xe3, 0x68,
	0x6e, 0xf4, 0x44, 0xbc, 0x84, 0xa6, 0xcb, 0xf3, 0xcc, 0x3f, 0x6b, 0x50, 0x5a, 0x0c, 0xb3, 0x71,
	0xcf, 0x95, 0x63, 0x63, 0x9f, 0x2b, 0x6e, 0x28, 0x14, 0xbe, 0xaa, 0x2f, 0x21, 0xff, 0xc0, 0x59,
	0xc8, 0xfd, 0x56, 0x5e, 0x44, 0x5f, 0xa1, 0xf9, 0x62, 0xda, 0x6b, 0x9a, 0x69, 0xa9, 0xb4, 0xb4,
	0x07, 0xa1, 0x0f, 0xe2, 0x7c, 0xc1, 0xd6, 0x6f, 0x83, 0xc6, 0xad, 0x95, 0xc5, 0xdb, 0x07, 0x16,
	0xb9, 0x13, 0x7e, 0x10, 0xe7, 0xef, 0x0d, 0xd8, 0xe3, 0x1a, 0x68, 0x86, 0x6b, 0xda, 0xe1, 0x43,
	0xd0, 0x49, 0x1f, 0x1d, 0xd7, 0x2b, 0x3c, 0x20, 0xca, 0x8b, 0x27, 0x17, 0x8c, 0x87, 0x2f, 0x77,
	0xaa, 0xba, 0x78, 0xae, 0x09, 0xc6, 0xe1, 0xa3, 0x55, 0xde, 0x2f, 0x83, 0xcc, 0xd7, 0xc8, 0xe9,
	0xea, 0xe6, 0xfb, 0x3e, 0x83, 0xc2, 0x28, 0x27, 0x23, 0x61, 0xfc, 0x43, 0x69, 0xfb, 0x98, 0xf2,
	0x13, 0x72, 0xa8, 0x7e, 0xc5, 0xf8, 0x87, 0x62, 0x0d, 0x39, 0x72, 0xaf, 0x44, 0xff, 0x9f, 0xbd,
	0xb2, 0x73, 0x12, 0x1e, 0xf4, 0xb7, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x85, 0xa3, 0xcb, 0x61,
	0xf1, 0x0f, 0x00, 0x00,
}
