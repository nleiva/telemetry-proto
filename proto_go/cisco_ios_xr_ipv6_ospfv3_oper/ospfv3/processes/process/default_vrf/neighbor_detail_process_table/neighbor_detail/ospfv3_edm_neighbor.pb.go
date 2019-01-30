// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_neighbor.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_neighbor_detail_process_table_neighbor_detail

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

// OSPFv3 neighbor summary information
type Ospfv3EdmNeighbor_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmNeighbor_KEYS) Reset()         { *m = Ospfv3EdmNeighbor_KEYS{} }
func (m *Ospfv3EdmNeighbor_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNeighbor_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmNeighbor_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afebde32b9237de, []int{0}
}

func (m *Ospfv3EdmNeighbor_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNeighbor_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmNeighbor_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNeighbor_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNeighbor_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNeighbor_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmNeighbor_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNeighbor_KEYS.Size(m)
}
func (m *Ospfv3EdmNeighbor_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNeighbor_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNeighbor_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmNeighbor_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmNeighbor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmNeighbor_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type Ospfv3EdmNeighbor struct {
	// Neighbor IP Address
	NeighborAddress string `protobuf:"bytes,50,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	// Neighbor interface ID
	NeighborInterfaceId uint32 `protobuf:"varint,51,opt,name=neighbor_interface_id,json=neighborInterfaceId,proto3" json:"neighbor_interface_id,omitempty"`
	// Neighbor's DR priority
	NeighborDrPriority uint32 `protobuf:"varint,52,opt,name=neighbor_dr_priority,json=neighborDrPriority,proto3" json:"neighbor_dr_priority,omitempty"`
	// Neighbor's state
	NeighborState string `protobuf:"bytes,53,opt,name=neighbor_state,json=neighborState,proto3" json:"neighbor_state,omitempty"`
	// Designated router
	NeighborDesignatedRouter string `protobuf:"bytes,54,opt,name=neighbor_designated_router,json=neighborDesignatedRouter,proto3" json:"neighbor_designated_router,omitempty"`
	// Time until neighbor's dead timer expires (seconds)
	NeighborDeadTimer uint32 `protobuf:"varint,55,opt,name=neighbor_dead_timer,json=neighborDeadTimer,proto3" json:"neighbor_dead_timer,omitempty"`
	// Amount of time since the adjacency is up (seconds)
	NeighborUpTime uint32 `protobuf:"varint,56,opt,name=neighbor_up_time,json=neighborUpTime,proto3" json:"neighbor_up_time,omitempty"`
	// Neighbor virtual link id
	NeighborVirtualLinkId uint32 `protobuf:"varint,57,opt,name=neighbor_virtual_link_id,json=neighborVirtualLinkId,proto3" json:"neighbor_virtual_link_id,omitempty"`
	// If true, neighbor is on a virtual link
	IsNeighborVirtualLink bool `protobuf:"varint,58,opt,name=is_neighbor_virtual_link,json=isNeighborVirtualLink,proto3" json:"is_neighbor_virtual_link,omitempty"`
	// Neighbor sham link id
	NeighborShamLinkId uint32 `protobuf:"varint,59,opt,name=neighbor_sham_link_id,json=neighborShamLinkId,proto3" json:"neighbor_sham_link_id,omitempty"`
	// If true, neighbor is on a sham link
	IsNeighborShamLink bool `protobuf:"varint,60,opt,name=is_neighbor_sham_link,json=isNeighborShamLink,proto3" json:"is_neighbor_sham_link,omitempty"`
	// Detailed OSPFv3 neighbor information
	NeighborDetail *Ospfv3EdmNeighborDetail `protobuf:"bytes,61,opt,name=neighbor_detail,json=neighborDetail,proto3" json:"neighbor_detail,omitempty"`
	// Neighbor BFD information
	NeighborBfdInfo      *Ospfv3EdmNeighborBfd `protobuf:"bytes,62,opt,name=neighbor_bfd_info,json=neighborBfdInfo,proto3" json:"neighbor_bfd_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ospfv3EdmNeighbor) Reset()         { *m = Ospfv3EdmNeighbor{} }
func (m *Ospfv3EdmNeighbor) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNeighbor) ProtoMessage()    {}
func (*Ospfv3EdmNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afebde32b9237de, []int{1}
}

func (m *Ospfv3EdmNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNeighbor.Unmarshal(m, b)
}
func (m *Ospfv3EdmNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNeighbor.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNeighbor.Merge(m, src)
}
func (m *Ospfv3EdmNeighbor) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNeighbor.Size(m)
}
func (m *Ospfv3EdmNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNeighbor proto.InternalMessageInfo

func (m *Ospfv3EdmNeighbor) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *Ospfv3EdmNeighbor) GetNeighborInterfaceId() uint32 {
	if m != nil {
		return m.NeighborInterfaceId
	}
	return 0
}

func (m *Ospfv3EdmNeighbor) GetNeighborDrPriority() uint32 {
	if m != nil {
		return m.NeighborDrPriority
	}
	return 0
}

func (m *Ospfv3EdmNeighbor) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *Ospfv3EdmNeighbor) GetNeighborDesignatedRouter() string {
	if m != nil {
		return m.NeighborDesignatedRouter
	}
	return ""
}

func (m *Ospfv3EdmNeighbor) GetNeighborDeadTimer() uint32 {
	if m != nil {
		return m.NeighborDeadTimer
	}
	return 0
}

func (m *Ospfv3EdmNeighbor) GetNeighborUpTime() uint32 {
	if m != nil {
		return m.NeighborUpTime
	}
	return 0
}

func (m *Ospfv3EdmNeighbor) GetNeighborVirtualLinkId() uint32 {
	if m != nil {
		return m.NeighborVirtualLinkId
	}
	return 0
}

func (m *Ospfv3EdmNeighbor) GetIsNeighborVirtualLink() bool {
	if m != nil {
		return m.IsNeighborVirtualLink
	}
	return false
}

func (m *Ospfv3EdmNeighbor) GetNeighborShamLinkId() uint32 {
	if m != nil {
		return m.NeighborShamLinkId
	}
	return 0
}

func (m *Ospfv3EdmNeighbor) GetIsNeighborShamLink() bool {
	if m != nil {
		return m.IsNeighborShamLink
	}
	return false
}

func (m *Ospfv3EdmNeighbor) GetNeighborDetail() *Ospfv3EdmNeighborDetail {
	if m != nil {
		return m.NeighborDetail
	}
	return nil
}

func (m *Ospfv3EdmNeighbor) GetNeighborBfdInfo() *Ospfv3EdmNeighborBfd {
	if m != nil {
		return m.NeighborBfdInfo
	}
	return nil
}

// OSPFv3 neighbor retransmission information
type Ospfv3EdmNeighborRetrans struct {
	// Number of database descriptor retransmissions during last exchange
	DatabaseDescriptorRetransmissions uint32 `protobuf:"varint,1,opt,name=database_descriptor_retransmissions,json=databaseDescriptorRetransmissions,proto3" json:"database_descriptor_retransmissions,omitempty"`
	// Area scope LSA's flood index
	AreaFloodIndex uint32 `protobuf:"varint,2,opt,name=area_flood_index,json=areaFloodIndex,proto3" json:"area_flood_index,omitempty"`
	// AS scope LSA's flood index
	AsFloodIndex uint32 `protobuf:"varint,3,opt,name=as_flood_index,json=asFloodIndex,proto3" json:"as_flood_index,omitempty"`
	// Link flood index
	LinkFloodIndex uint32 `protobuf:"varint,4,opt,name=link_flood_index,json=linkFloodIndex,proto3" json:"link_flood_index,omitempty"`
	// Number of neighbor retransmissions
	NeighborRetransmissions uint32 `protobuf:"varint,5,opt,name=neighbor_retransmissions,json=neighborRetransmissions,proto3" json:"neighbor_retransmissions,omitempty"`
	// Number of retransmissions for this neighbor
	Retransmissions uint32 `protobuf:"varint,6,opt,name=retransmissions,proto3" json:"retransmissions,omitempty"`
	// First flood item for area scope LSAs
	AreaFirstFlood uint32 `protobuf:"varint,7,opt,name=area_first_flood,json=areaFirstFlood,proto3" json:"area_first_flood,omitempty"`
	// Index of the first flood item for area scope LSAs
	AreaFirstFloodIndex uint32 `protobuf:"varint,8,opt,name=area_first_flood_index,json=areaFirstFloodIndex,proto3" json:"area_first_flood_index,omitempty"`
	// First flood item for AS scope LSAs
	AsFirstFlood uint32 `protobuf:"varint,9,opt,name=as_first_flood,json=asFirstFlood,proto3" json:"as_first_flood,omitempty"`
	// Index for first flood item for AS scope LSAs
	AsFirstFloodIndex uint32 `protobuf:"varint,10,opt,name=as_first_flood_index,json=asFirstFloodIndex,proto3" json:"as_first_flood_index,omitempty"`
	// Link first flood information
	LinkFirstFlood uint32 `protobuf:"varint,11,opt,name=link_first_flood,json=linkFirstFlood,proto3" json:"link_first_flood,omitempty"`
	// Link first flood information index
	LinkFirstFloodIndex uint32 `protobuf:"varint,12,opt,name=link_first_flood_index,json=linkFirstFloodIndex,proto3" json:"link_first_flood_index,omitempty"`
	// Next flood item for area scope LSAs
	AreaNextFlood uint32 `protobuf:"varint,13,opt,name=area_next_flood,json=areaNextFlood,proto3" json:"area_next_flood,omitempty"`
	// Index of next flood item for Area scope LSAs
	AreaNextFloodIndex uint32 `protobuf:"varint,14,opt,name=area_next_flood_index,json=areaNextFloodIndex,proto3" json:"area_next_flood_index,omitempty"`
	// Next flood item for AS scope LSAs
	AsNextFlood uint32 `protobuf:"varint,15,opt,name=as_next_flood,json=asNextFlood,proto3" json:"as_next_flood,omitempty"`
	// Index of next flood item for AS scope LSAs
	AsNextFloodIndex uint32 `protobuf:"varint,16,opt,name=as_next_flood_index,json=asNextFloodIndex,proto3" json:"as_next_flood_index,omitempty"`
	// Link next flood information
	LinkNextFlood uint32 `protobuf:"varint,17,opt,name=link_next_flood,json=linkNextFlood,proto3" json:"link_next_flood,omitempty"`
	// Link next flood information index
	LinkNextFloodIndex uint32 `protobuf:"varint,18,opt,name=link_next_flood_index,json=linkNextFloodIndex,proto3" json:"link_next_flood_index,omitempty"`
	// Number of LSAs sent in last retransmission
	LastRetransmissionLength uint32 `protobuf:"varint,19,opt,name=last_retransmission_length,json=lastRetransmissionLength,proto3" json:"last_retransmission_length,omitempty"`
	// Maximum number of LSAs sent in a retransmission
	MaximumRetransmissionLength uint32 `protobuf:"varint,20,opt,name=maximum_retransmission_length,json=maximumRetransmissionLength,proto3" json:"maximum_retransmission_length,omitempty"`
	// Last retransmission scan time (ms)
	LastRetransmissionTime uint32 `protobuf:"varint,21,opt,name=last_retransmission_time,json=lastRetransmissionTime,proto3" json:"last_retransmission_time,omitempty"`
	// Maximum retransmission scan time (ms)
	MaximumRetransmissionTime uint32 `protobuf:"varint,22,opt,name=maximum_retransmission_time,json=maximumRetransmissionTime,proto3" json:"maximum_retransmission_time,omitempty"`
	// Time until next LSA retransmission (ms)
	LsaRetransmissionTimer uint32   `protobuf:"varint,23,opt,name=lsa_retransmission_timer,json=lsaRetransmissionTimer,proto3" json:"lsa_retransmission_timer,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Ospfv3EdmNeighborRetrans) Reset()         { *m = Ospfv3EdmNeighborRetrans{} }
func (m *Ospfv3EdmNeighborRetrans) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNeighborRetrans) ProtoMessage()    {}
func (*Ospfv3EdmNeighborRetrans) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afebde32b9237de, []int{2}
}

func (m *Ospfv3EdmNeighborRetrans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNeighborRetrans.Unmarshal(m, b)
}
func (m *Ospfv3EdmNeighborRetrans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNeighborRetrans.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNeighborRetrans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNeighborRetrans.Merge(m, src)
}
func (m *Ospfv3EdmNeighborRetrans) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNeighborRetrans.Size(m)
}
func (m *Ospfv3EdmNeighborRetrans) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNeighborRetrans.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNeighborRetrans proto.InternalMessageInfo

func (m *Ospfv3EdmNeighborRetrans) GetDatabaseDescriptorRetransmissions() uint32 {
	if m != nil {
		return m.DatabaseDescriptorRetransmissions
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAreaFloodIndex() uint32 {
	if m != nil {
		return m.AreaFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAsFloodIndex() uint32 {
	if m != nil {
		return m.AsFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLinkFloodIndex() uint32 {
	if m != nil {
		return m.LinkFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetNeighborRetransmissions() uint32 {
	if m != nil {
		return m.NeighborRetransmissions
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetRetransmissions() uint32 {
	if m != nil {
		return m.Retransmissions
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAreaFirstFlood() uint32 {
	if m != nil {
		return m.AreaFirstFlood
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAreaFirstFloodIndex() uint32 {
	if m != nil {
		return m.AreaFirstFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAsFirstFlood() uint32 {
	if m != nil {
		return m.AsFirstFlood
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAsFirstFloodIndex() uint32 {
	if m != nil {
		return m.AsFirstFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLinkFirstFlood() uint32 {
	if m != nil {
		return m.LinkFirstFlood
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLinkFirstFloodIndex() uint32 {
	if m != nil {
		return m.LinkFirstFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAreaNextFlood() uint32 {
	if m != nil {
		return m.AreaNextFlood
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAreaNextFloodIndex() uint32 {
	if m != nil {
		return m.AreaNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAsNextFlood() uint32 {
	if m != nil {
		return m.AsNextFlood
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetAsNextFloodIndex() uint32 {
	if m != nil {
		return m.AsNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLinkNextFlood() uint32 {
	if m != nil {
		return m.LinkNextFlood
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLinkNextFloodIndex() uint32 {
	if m != nil {
		return m.LinkNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLastRetransmissionLength() uint32 {
	if m != nil {
		return m.LastRetransmissionLength
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetMaximumRetransmissionLength() uint32 {
	if m != nil {
		return m.MaximumRetransmissionLength
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLastRetransmissionTime() uint32 {
	if m != nil {
		return m.LastRetransmissionTime
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetMaximumRetransmissionTime() uint32 {
	if m != nil {
		return m.MaximumRetransmissionTime
	}
	return 0
}

func (m *Ospfv3EdmNeighborRetrans) GetLsaRetransmissionTimer() uint32 {
	if m != nil {
		return m.LsaRetransmissionTimer
	}
	return 0
}

// OSPFv3 Neighbor BFD information
type Ospfv3EdmNeighborBfd struct {
	// BFD enable mode - Default/Strict
	BfdIntfEnableMode uint32 `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode,proto3" json:"bfd_intf_enable_mode,omitempty"`
	// Status of the BFD Session
	BfdStatusFlag        uint32   `protobuf:"varint,2,opt,name=bfd_status_flag,json=bfdStatusFlag,proto3" json:"bfd_status_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmNeighborBfd) Reset()         { *m = Ospfv3EdmNeighborBfd{} }
func (m *Ospfv3EdmNeighborBfd) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNeighborBfd) ProtoMessage()    {}
func (*Ospfv3EdmNeighborBfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afebde32b9237de, []int{3}
}

func (m *Ospfv3EdmNeighborBfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNeighborBfd.Unmarshal(m, b)
}
func (m *Ospfv3EdmNeighborBfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNeighborBfd.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNeighborBfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNeighborBfd.Merge(m, src)
}
func (m *Ospfv3EdmNeighborBfd) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNeighborBfd.Size(m)
}
func (m *Ospfv3EdmNeighborBfd) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNeighborBfd.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNeighborBfd proto.InternalMessageInfo

func (m *Ospfv3EdmNeighborBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *Ospfv3EdmNeighborBfd) GetBfdStatusFlag() uint32 {
	if m != nil {
		return m.BfdStatusFlag
	}
	return 0
}

// Detailed OSPFv3 neighbor information
type Ospfv3EdmNeighborDetail struct {
	// Number of state changes
	StateChanges uint32 `protobuf:"varint,1,opt,name=state_changes,json=stateChanges,proto3" json:"state_changes,omitempty"`
	// Cost of path to this neighbor
	NeighborCost uint32 `protobuf:"varint,2,opt,name=neighbor_cost,json=neighborCost,proto3" json:"neighbor_cost,omitempty"`
	// If true, filter outgoing LSAs
	IsNeighborFiltered bool `protobuf:"varint,3,opt,name=is_neighbor_filtered,json=isNeighborFiltered,proto3" json:"is_neighbor_filtered,omitempty"`
	// Address of designated router
	NeighborDesignatedRouterAddress string `protobuf:"bytes,4,opt,name=neighbor_designated_router_address,json=neighborDesignatedRouterAddress,proto3" json:"neighbor_designated_router_address,omitempty"`
	// Address of backup designated router
	NeighborBackupDesignatedRouterAddress string `protobuf:"bytes,5,opt,name=neighbor_backup_designated_router_address,json=neighborBackupDesignatedRouterAddress,proto3" json:"neighbor_backup_designated_router_address,omitempty"`
	// Interface type
	InterfaceType string `protobuf:"bytes,6,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	// Poll interval (s)
	PollInterval uint32 `protobuf:"varint,7,opt,name=poll_interval,json=pollInterval,proto3" json:"poll_interval,omitempty"`
	// For NBMA networks, amount of time remaining in seconds before the next poll interval expires and Hello is sent (s)
	NextPollInterval uint32 `protobuf:"varint,8,opt,name=next_poll_interval,json=nextPollInterval,proto3" json:"next_poll_interval,omitempty"`
	// Remaining time when ignore timer is running
	NeighborIgnoreTimer uint32 `protobuf:"varint,9,opt,name=neighbor_ignore_timer,json=neighborIgnoreTimer,proto3" json:"neighbor_ignore_timer,omitempty"`
	//  This is bitmask of neighbor's option field received
	NeighborOption uint32 `protobuf:"varint,10,opt,name=neighbor_option,json=neighborOption,proto3" json:"neighbor_option,omitempty"`
	// Number of pending events
	PendingEvents uint32 `protobuf:"varint,11,opt,name=pending_events,json=pendingEvents,proto3" json:"pending_events,omitempty"`
	// Retransmission information with this neighbor
	NeighborRetransmission *Ospfv3EdmNeighborRetrans `protobuf:"bytes,12,opt,name=neighbor_retransmission,json=neighborRetransmission,proto3" json:"neighbor_retransmission,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                  `json:"-"`
	XXX_unrecognized       []byte                    `json:"-"`
	XXX_sizecache          int32                     `json:"-"`
}

func (m *Ospfv3EdmNeighborDetail) Reset()         { *m = Ospfv3EdmNeighborDetail{} }
func (m *Ospfv3EdmNeighborDetail) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNeighborDetail) ProtoMessage()    {}
func (*Ospfv3EdmNeighborDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afebde32b9237de, []int{4}
}

func (m *Ospfv3EdmNeighborDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNeighborDetail.Unmarshal(m, b)
}
func (m *Ospfv3EdmNeighborDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNeighborDetail.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNeighborDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNeighborDetail.Merge(m, src)
}
func (m *Ospfv3EdmNeighborDetail) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNeighborDetail.Size(m)
}
func (m *Ospfv3EdmNeighborDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNeighborDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNeighborDetail proto.InternalMessageInfo

func (m *Ospfv3EdmNeighborDetail) GetStateChanges() uint32 {
	if m != nil {
		return m.StateChanges
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetNeighborCost() uint32 {
	if m != nil {
		return m.NeighborCost
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetIsNeighborFiltered() bool {
	if m != nil {
		return m.IsNeighborFiltered
	}
	return false
}

func (m *Ospfv3EdmNeighborDetail) GetNeighborDesignatedRouterAddress() string {
	if m != nil {
		return m.NeighborDesignatedRouterAddress
	}
	return ""
}

func (m *Ospfv3EdmNeighborDetail) GetNeighborBackupDesignatedRouterAddress() string {
	if m != nil {
		return m.NeighborBackupDesignatedRouterAddress
	}
	return ""
}

func (m *Ospfv3EdmNeighborDetail) GetInterfaceType() string {
	if m != nil {
		return m.InterfaceType
	}
	return ""
}

func (m *Ospfv3EdmNeighborDetail) GetPollInterval() uint32 {
	if m != nil {
		return m.PollInterval
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetNextPollInterval() uint32 {
	if m != nil {
		return m.NextPollInterval
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetNeighborIgnoreTimer() uint32 {
	if m != nil {
		return m.NeighborIgnoreTimer
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetNeighborOption() uint32 {
	if m != nil {
		return m.NeighborOption
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetPendingEvents() uint32 {
	if m != nil {
		return m.PendingEvents
	}
	return 0
}

func (m *Ospfv3EdmNeighborDetail) GetNeighborRetransmission() *Ospfv3EdmNeighborRetrans {
	if m != nil {
		return m.NeighborRetransmission
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmNeighbor_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.neighbor_detail_process_table.neighbor_detail.ospfv3_edm_neighbor_KEYS")
	proto.RegisterType((*Ospfv3EdmNeighbor)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.neighbor_detail_process_table.neighbor_detail.ospfv3_edm_neighbor")
	proto.RegisterType((*Ospfv3EdmNeighborRetrans)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.neighbor_detail_process_table.neighbor_detail.ospfv3_edm_neighbor_retrans")
	proto.RegisterType((*Ospfv3EdmNeighborBfd)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.neighbor_detail_process_table.neighbor_detail.ospfv3_edm_neighbor_bfd")
	proto.RegisterType((*Ospfv3EdmNeighborDetail)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.neighbor_detail_process_table.neighbor_detail.ospfv3_edm_neighbor_detail")
}

func init() { proto.RegisterFile("ospfv3_edm_neighbor.proto", fileDescriptor_6afebde32b9237de) }

var fileDescriptor_6afebde32b9237de = []byte{
	// 1119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xdd, 0x6e, 0x1c, 0x35,
	0x14, 0xd6, 0xd2, 0xdf, 0x38, 0xbb, 0xf9, 0x71, 0xfe, 0xa6, 0x89, 0x10, 0xe9, 0x96, 0xc2, 0x56,
	0x82, 0x85, 0x24, 0xd0, 0x1f, 0x08, 0x95, 0x68, 0x93, 0x48, 0xab, 0x96, 0x50, 0x6d, 0x0a, 0x82,
	0x2b, 0xcb, 0xbb, 0xe3, 0xd9, 0x58, 0x99, 0x19, 0x8f, 0x6c, 0xef, 0x2a, 0x79, 0x02, 0x84, 0x04,
	0xe2, 0x15, 0x78, 0x03, 0xee, 0x78, 0x05, 0x5e, 0x0b, 0xf9, 0xd8, 0x33, 0xe3, 0x99, 0xcc, 0x72,
	0x89, 0x7a, 0xb7, 0xf9, 0xce, 0xf7, 0x7d, 0x3e, 0x3e, 0x63, 0x1f, 0x9f, 0xa0, 0x7b, 0x42, 0x65,
	0xd1, 0xec, 0x80, 0xb0, 0x30, 0x21, 0x29, 0xe3, 0x93, 0xf3, 0x91, 0x90, 0xfd, 0x4c, 0x0a, 0x2d,
	0x70, 0x36, 0xe6, 0x6a, 0x2c, 0x08, 0x17, 0x8a, 0x5c, 0x4a, 0xc2, 0xb3, 0xd9, 0x63, 0xe2, 0xc8,
	0x22, 0x63, 0xb2, 0x6f, 0x7f, 0x1b, 0xee, 0x98, 0x29, 0xc5, 0x54, 0xfe, 0xab, 0x1f, 0xb2, 0x88,
	0x4e, 0x63, 0x4d, 0x66, 0x32, 0xea, 0xe7, 0x96, 0x24, 0x64, 0x9a, 0xf2, 0x98, 0x38, 0x0e, 0xd1,
	0x74, 0x14, 0xb3, 0x7a, 0xb4, 0xfb, 0x47, 0x0b, 0x05, 0x0d, 0xf9, 0x90, 0x57, 0xc7, 0x3f, 0x9f,
	0xe1, 0xfb, 0xa8, 0x9d, 0xab, 0x53, 0x9a, 0xb0, 0xa0, 0xb5, 0xdb, 0xea, 0x2d, 0x0c, 0x17, 0x1d,
	0x76, 0x4a, 0x13, 0x86, 0x1f, 0xa2, 0x25, 0x9e, 0x6a, 0x26, 0x23, 0x3a, 0x66, 0x96, 0xf4, 0x1e,
	0x90, 0x3a, 0x05, 0x0a, 0xb4, 0x47, 0x68, 0xa5, 0xb0, 0xa6, 0x61, 0x28, 0x99, 0x52, 0xc1, 0x0d,
	0x20, 0x2e, 0xe7, 0xf8, 0xb7, 0x16, 0xee, 0xfe, 0x75, 0x07, 0xad, 0x35, 0x64, 0xd4, 0x68, 0xb1,
	0xdf, 0x68, 0x81, 0xf7, 0xd1, 0x46, 0x41, 0x2d, 0xb3, 0xe3, 0x61, 0x70, 0xb0, 0xdb, 0xea, 0x75,
	0x86, 0x6b, 0x79, 0x70, 0x90, 0xc7, 0x06, 0x21, 0xfe, 0x1c, 0xad, 0x97, 0xb5, 0x91, 0x24, 0x93,
	0x5c, 0x48, 0xae, 0xaf, 0x82, 0x2f, 0x40, 0x82, 0xf3, 0xd8, 0x91, 0x7c, 0xe3, 0x22, 0x66, 0xeb,
	0x85, 0x42, 0x69, 0xaa, 0x59, 0xf0, 0xa5, 0xdd, 0x7a, 0x8e, 0x9e, 0x19, 0x10, 0x1f, 0xa2, 0x6d,
	0xaf, 0xe8, 0x8a, 0x4f, 0x52, 0xaa, 0x59, 0x48, 0xa4, 0x98, 0x6a, 0x26, 0x83, 0xc7, 0x20, 0x09,
	0x0a, 0xfb, 0x82, 0x30, 0x84, 0x38, 0xee, 0xa3, 0x35, 0x4f, 0x4d, 0x43, 0xa2, 0x79, 0xc2, 0x64,
	0xf0, 0x04, 0xb2, 0x5a, 0x2d, 0x65, 0x34, 0x7c, 0x6b, 0x02, 0xb8, 0xe7, 0x55, 0x69, 0x9a, 0x01,
	0x3b, 0x78, 0x0a, 0xe4, 0x22, 0xd9, 0x1f, 0x32, 0x43, 0xc5, 0x4f, 0x50, 0xb1, 0x2a, 0x99, 0x71,
	0xa9, 0xa7, 0x34, 0x26, 0x31, 0x4f, 0x2f, 0x4c, 0x9d, 0x9e, 0x81, 0xa2, 0x28, 0xe2, 0x8f, 0x36,
	0xfc, 0x9a, 0xa7, 0x17, 0x83, 0xd0, 0x08, 0xb9, 0x22, 0x8d, 0xda, 0xe0, 0xab, 0xdd, 0x56, 0xef,
	0xee, 0x70, 0x83, 0xab, 0xd3, 0xeb, 0x52, 0xbc, 0xe7, 0x7d, 0x16, 0x75, 0x4e, 0x93, 0x62, 0xb9,
	0xaf, 0xab, 0x35, 0x3e, 0x3b, 0xa7, 0x89, 0x5b, 0x6b, 0x0f, 0x6d, 0xf8, 0x6b, 0x15, 0xaa, 0xe0,
	0x10, 0x16, 0xc2, 0xe5, 0x42, 0xb9, 0x08, 0xff, 0xdd, 0x42, 0xcb, 0xb5, 0x53, 0x1e, 0x7c, 0xb3,
	0xdb, 0xea, 0x2d, 0xee, 0xff, 0xd6, 0xea, 0xff, 0xdf, 0xf7, 0xab, 0xdf, 0x74, 0xb7, 0x6c, 0xa8,
	0xfc, 0x20, 0x47, 0xf0, 0xb7, 0x49, 0xbc, 0xf8, 0xa0, 0x64, 0x14, 0x85, 0x84, 0xa7, 0x91, 0x08,
	0x9e, 0x43, 0xea, 0xbf, 0xbe, 0x23, 0xa9, 0x8f, 0xa2, 0xb0, 0xbc, 0x6e, 0x2f, 0xa2, 0x70, 0x90,
	0x46, 0xa2, 0xfb, 0xe7, 0x02, 0xda, 0x69, 0x22, 0x4b, 0xa6, 0x25, 0x4d, 0x15, 0x3e, 0x45, 0x0f,
	0x42, 0xaa, 0xe9, 0x88, 0x2a, 0x66, 0x6e, 0xc0, 0x58, 0xf2, 0x4c, 0x97, 0xe1, 0x84, 0x2b, 0xc5,
	0x45, 0xaa, 0xa0, 0xbb, 0x74, 0x86, 0xf7, 0x73, 0xea, 0x51, 0xc1, 0x1c, 0x56, 0x89, 0xe6, 0x8c,
	0x53, 0xc9, 0x28, 0x89, 0x62, 0x21, 0x4c, 0x95, 0x42, 0x76, 0x09, 0x5d, 0xa7, 0x33, 0x5c, 0x32,
	0xf8, 0x89, 0x81, 0x07, 0x06, 0xc5, 0x1f, 0xa2, 0x25, 0xaa, 0x2a, 0xbc, 0x1b, 0xc0, 0x6b, 0x53,
	0xe5, 0xb1, 0x7a, 0x68, 0x05, 0x4e, 0xa2, 0xcf, 0xbb, 0x69, 0xfd, 0x0c, 0xee, 0x31, 0x9f, 0x79,
	0x77, 0xa6, 0x9e, 0xfe, 0x2d, 0x50, 0x6c, 0xe5, 0xf1, 0xeb, 0x49, 0x2f, 0xd7, 0x15, 0xb7, 0x41,
	0x51, 0x87, 0xcb, 0xed, 0x71, 0xa9, 0xb4, 0x4d, 0x2a, 0xb8, 0xe3, 0x6d, 0xcf, 0xc0, 0x90, 0x13,
	0x3e, 0x40, 0x9b, 0x75, 0xa6, 0x4b, 0xff, 0xae, 0x6d, 0x74, 0x55, 0x7e, 0xa5, 0x26, 0x9e, 0xf9,
	0x42, 0x51, 0x93, 0xd2, 0xfa, 0x33, 0xb4, 0x5e, 0x65, 0x39, 0x63, 0x64, 0x1b, 0x8f, 0xcf, 0xad,
	0x15, 0xd1, 0x33, 0x5e, 0xf4, 0x8a, 0x58, 0xc9, 0xba, 0xce, 0x74, 0xe6, 0x6d, 0x9b, 0x75, 0x95,
	0x6f, 0xed, 0x3f, 0x42, 0xcb, 0xb0, 0xd5, 0x94, 0x5d, 0xe6, 0xee, 0x1d, 0x60, 0x77, 0x0c, 0x7c,
	0xca, 0x2e, 0x9d, 0xf9, 0x1e, 0xda, 0xa8, 0xf1, 0x9c, 0xf7, 0x92, 0xed, 0x31, 0x15, 0xb6, 0xb5,
	0xee, 0xa2, 0x0e, 0x55, 0xbe, 0xf1, 0x32, 0x50, 0x17, 0xa9, 0x2a, 0x6d, 0x3f, 0x45, 0x6b, 0x15,
	0x8e, 0x33, 0x5d, 0x01, 0xe6, 0x8a, 0xc7, 0x2c, 0xb2, 0x85, 0x2d, 0x7a, 0xa6, 0xab, 0x36, 0x5b,
	0x03, 0x57, 0xb2, 0xad, 0xf1, 0x9c, 0x31, 0xb6, 0xd9, 0x56, 0xd8, 0xd6, 0xfa, 0x10, 0x6d, 0xc7,
	0x54, 0xe9, 0xda, 0xf1, 0x23, 0x31, 0x4b, 0x27, 0xfa, 0x3c, 0x58, 0x03, 0x5d, 0x60, 0x18, 0xd5,
	0x03, 0xf8, 0x1a, 0xe2, 0xf8, 0x05, 0x7a, 0x3f, 0xa1, 0x97, 0x3c, 0x99, 0x26, 0x73, 0x0c, 0xd6,
	0xc1, 0x60, 0xc7, 0x91, 0x1a, 0x3d, 0x9e, 0xa2, 0xa0, 0x29, 0x03, 0x78, 0x6a, 0x36, 0x40, 0xbe,
	0x79, 0x7d, 0x7d, 0x78, 0x72, 0x9e, 0xa3, 0x9d, 0x39, 0xab, 0x83, 0x78, 0x13, 0xc4, 0xf7, 0x1a,
	0xd7, 0x06, 0xbd, 0x59, 0x59, 0xd1, 0x26, 0xad, 0x0c, 0xb6, 0xdc, 0xca, 0x8a, 0x5e, 0x17, 0xca,
	0xae, 0x44, 0x5b, 0x73, 0xda, 0x99, 0x39, 0xe9, 0xb6, 0xd9, 0xea, 0x88, 0xb0, 0xd4, 0x74, 0x42,
	0x92, 0x88, 0x90, 0xb9, 0x76, 0xb4, 0x3a, 0x32, 0x4d, 0x4e, 0x47, 0xc7, 0x10, 0xf9, 0x4e, 0x84,
	0xcc, 0x7c, 0x5c, 0x23, 0x30, 0x4f, 0xfe, 0xd4, 0x34, 0x17, 0x3a, 0x71, 0xdd, 0xa7, 0x33, 0x8a,
	0xc2, 0x33, 0x40, 0x4f, 0x62, 0x3a, 0xe9, 0xfe, 0x72, 0x1b, 0x6d, 0xcf, 0x6f, 0xff, 0xf8, 0x01,
	0xea, 0xc0, 0xd4, 0x40, 0xc6, 0xe7, 0x34, 0x9d, 0xb0, 0xbc, 0xff, 0xb5, 0x01, 0x7c, 0x69, 0x31,
	0x43, 0x2a, 0x74, 0x63, 0xa1, 0xb4, 0x5b, 0xa9, 0x9d, 0x83, 0x2f, 0x85, 0xd2, 0x66, 0x74, 0xf1,
	0x1f, 0xc9, 0x88, 0xc7, 0x9a, 0x49, 0x16, 0x42, 0xaf, 0xab, 0xbc, 0x91, 0x27, 0x2e, 0x82, 0x5f,
	0xa1, 0xee, 0xfc, 0x99, 0xa4, 0x98, 0xae, 0x6e, 0xc2, 0x6c, 0xf2, 0xc1, 0xbc, 0xd9, 0x24, 0x9f,
	0xb6, 0x7e, 0x42, 0x8f, 0xca, 0x82, 0xd2, 0xf1, 0xc5, 0x34, 0xfb, 0x0f, 0xcf, 0x5b, 0xe0, 0xf9,
	0xb0, 0x78, 0x42, 0x80, 0x3f, 0xcf, 0xb9, 0x32, 0x5c, 0xea, 0xab, 0x8c, 0x41, 0xcb, 0xf4, 0x87,
	0xcb, 0xb7, 0x57, 0x19, 0x33, 0x45, 0xca, 0x44, 0x1c, 0xdb, 0x51, 0x6f, 0x46, 0x63, 0xd7, 0x2d,
	0xdb, 0x06, 0x1c, 0x38, 0x0c, 0x7f, 0x82, 0x30, 0xdc, 0xb2, 0x2a, 0xd3, 0xf6, 0xc9, 0x15, 0x13,
	0x79, 0xe3, 0xb3, 0x2b, 0x13, 0xe4, 0x24, 0x15, 0x92, 0xb9, 0x63, 0xb6, 0x50, 0x9b, 0x20, 0x21,
	0x66, 0x47, 0xaf, 0x8f, 0xbd, 0xb9, 0x43, 0x64, 0x9a, 0x8b, 0xd4, 0x75, 0xcb, 0xe2, 0xa1, 0xff,
	0x1e, 0x50, 0xb3, 0xad, 0x8c, 0xa5, 0x21, 0x4f, 0x27, 0x84, 0xcd, 0x58, 0xaa, 0x95, 0x6b, 0x94,
	0x1d, 0x87, 0x1e, 0x03, 0x88, 0xff, 0x69, 0xa1, 0xad, 0x39, 0xaf, 0x0d, 0x74, 0xca, 0xc5, 0xfd,
	0xdf, 0xdf, 0x91, 0xa9, 0xc0, 0x25, 0x37, 0xdc, 0x6c, 0x7e, 0xfc, 0x46, 0xb7, 0xe1, 0xbf, 0x9b,
	0x83, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xb2, 0xaf, 0x8c, 0xfa, 0x0c, 0x00, 0x00,
}
