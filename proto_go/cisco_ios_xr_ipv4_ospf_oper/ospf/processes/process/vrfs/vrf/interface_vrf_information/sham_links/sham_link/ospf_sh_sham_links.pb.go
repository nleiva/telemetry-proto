// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_sham_links.proto

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_interface_vrf_information_sham_links_sham_link is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_sham_links.proto

It has these top-level messages:
	OspfShShamLinks_KEYS
	OspfShShamLinks
	OspfShNeighborRetrans
	OspfShSlinkNeighbor
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_interface_vrf_information_sham_links_sham_link

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

// OSPF Sham Link
type OspfShShamLinks_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	ShamLinkName string `protobuf:"bytes,3,opt,name=sham_link_name,json=shamLinkName" json:"sham_link_name,omitempty"`
}

func (m *OspfShShamLinks_KEYS) Reset()                    { *m = OspfShShamLinks_KEYS{} }
func (m *OspfShShamLinks_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShShamLinks_KEYS) ProtoMessage()               {}
func (*OspfShShamLinks_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShShamLinks_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShShamLinks_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShShamLinks_KEYS) GetShamLinkName() string {
	if m != nil {
		return m.ShamLinkName
	}
	return ""
}

type OspfShShamLinks struct {
	// Sham link name
	ShamLinkName string `protobuf:"bytes,50,opt,name=sham_link_name,json=shamLinkName" json:"sham_link_name,omitempty"`
	// Neighbor on other end of this sham link
	ShamLinkNeighborId string `protobuf:"bytes,51,opt,name=sham_link_neighbor_id,json=shamLinkNeighborId" json:"sham_link_neighbor_id,omitempty"`
	// Sham-link source
	ShamLinkSourceAddress string `protobuf:"bytes,52,opt,name=sham_link_source_address,json=shamLinkSourceAddress" json:"sham_link_source_address,omitempty"`
	// Sham-link dest
	ShamLinkDestAddress string `protobuf:"bytes,53,opt,name=sham_link_dest_address,json=shamLinkDestAddress" json:"sham_link_dest_address,omitempty"`
	// OSPF interface state for the sham link
	ShamLinkState string `protobuf:"bytes,54,opt,name=sham_link_state,json=shamLinkState" json:"sham_link_state,omitempty"`
	// If true, the link runs as demand circuit
	ShamLinkDemandCircuit bool `protobuf:"varint,55,opt,name=sham_link_demand_circuit,json=shamLinkDemandCircuit" json:"sham_link_demand_circuit,omitempty"`
	// Number of LSA's with demand circuit bit not set
	ShamLinkDcBitlessLsa uint32 `protobuf:"varint,56,opt,name=sham_link_dc_bitless_lsa,json=shamLinkDcBitlessLsa" json:"sham_link_dc_bitless_lsa,omitempty"`
	// Sham-link ifindex
	ShamLinkIfindex uint32 `protobuf:"varint,57,opt,name=sham_link_ifindex,json=shamLinkIfindex" json:"sham_link_ifindex,omitempty"`
	// Area id
	ShamLinkArea string `protobuf:"bytes,58,opt,name=sham_link_area,json=shamLinkArea" json:"sham_link_area,omitempty"`
	// Cost of the sham link
	ShamLinkCost uint32 `protobuf:"varint,59,opt,name=sham_link_cost,json=shamLinkCost" json:"sham_link_cost,omitempty"`
	// Transmission delay in seconds
	ShamLinkTransmissionDelay uint32 `protobuf:"varint,60,opt,name=sham_link_transmission_delay,json=shamLinkTransmissionDelay" json:"sham_link_transmission_delay,omitempty"`
	// Hello interval (s)
	ShamLinkHelloInterval uint32 `protobuf:"varint,61,opt,name=sham_link_hello_interval,json=shamLinkHelloInterval" json:"sham_link_hello_interval,omitempty"`
	// Hello interval (ms)
	ShamLinkHelloIntervalMs uint32 `protobuf:"varint,62,opt,name=sham_link_hello_interval_ms,json=shamLinkHelloIntervalMs" json:"sham_link_hello_interval_ms,omitempty"`
	// Dead interval (s)
	ShamLinkDeadInterval uint32 `protobuf:"varint,63,opt,name=sham_link_dead_interval,json=shamLinkDeadInterval" json:"sham_link_dead_interval,omitempty"`
	// Wait interval (s)
	ShamLinkWaitInterval uint32 `protobuf:"varint,64,opt,name=sham_link_wait_interval,json=shamLinkWaitInterval" json:"sham_link_wait_interval,omitempty"`
	// Retransmission interval (s)
	ShamLinkRetransmissionInterval uint32 `protobuf:"varint,65,opt,name=sham_link_retransmission_interval,json=shamLinkRetransmissionInterval" json:"sham_link_retransmission_interval,omitempty"`
	// Time until next hello (s)
	ShamLinkNextHello uint32 `protobuf:"varint,66,opt,name=sham_link_next_hello,json=shamLinkNextHello" json:"sham_link_next_hello,omitempty"`
	// Time until next hello (ms)
	ShamLinkNextHelloMs uint32 `protobuf:"varint,67,opt,name=sham_link_next_hello_ms,json=shamLinkNextHelloMs" json:"sham_link_next_hello_ms,omitempty"`
	// If true, interface is passive
	ShamLinkPassive bool `protobuf:"varint,68,opt,name=sham_link_passive,json=shamLinkPassive" json:"sham_link_passive,omitempty"`
	// Authentication type
	ShamLinkAuthenticationType string `protobuf:"bytes,69,opt,name=sham_link_authentication_type,json=shamLinkAuthenticationType" json:"sham_link_authentication_type,omitempty"`
	// If true, MD key configured
	ShamLinkYoungestMdKey bool `protobuf:"varint,70,opt,name=sham_link_youngest_md_key,json=shamLinkYoungestMdKey" json:"sham_link_youngest_md_key,omitempty"`
	// Youngest MD key ID
	ShamLinkYoungestMdKeyId uint32 `protobuf:"varint,71,opt,name=sham_link_youngest_md_key_id,json=shamLinkYoungestMdKeyId" json:"sham_link_youngest_md_key_id,omitempty"`
	// Number of neighbors still using the old key (rollover in progress)
	ShamLinkOldMdKeyCount uint32 `protobuf:"varint,72,opt,name=sham_link_old_md_key_count,json=shamLinkOldMdKeyCount" json:"sham_link_old_md_key_count,omitempty"`
	// List of old MD keys (if any)
	ShamLinkMdKeyList []uint32 `protobuf:"varint,73,rep,packed,name=sham_link_md_key_list,json=shamLinkMdKeyList" json:"sham_link_md_key_list,omitempty"`
	// Sham Link Keychain ID
	ShamLinkKeychainId uint64 `protobuf:"varint,74,opt,name=sham_link_keychain_id,json=shamLinkKeychainId" json:"sham_link_keychain_id,omitempty"`
	// If true, NSF enabled
	ShamLinkNsfEnabled bool `protobuf:"varint,75,opt,name=sham_link_nsf_enabled,json=shamLinkNsfEnabled" json:"sham_link_nsf_enabled,omitempty"`
	// If true, NSF restart in progress on the sham link
	ShamLinkNsf bool `protobuf:"varint,76,opt,name=sham_link_nsf,json=shamLinkNsf" json:"sham_link_nsf,omitempty"`
	// Time in seconds since last NSF
	ShamLinkLastNsf uint32 `protobuf:"varint,77,opt,name=sham_link_last_nsf,json=shamLinkLastNsf" json:"sham_link_last_nsf,omitempty"`
	// Neighbor information
	ShamLinkNeighbor *OspfShSlinkNeighbor `protobuf:"bytes,78,opt,name=sham_link_neighbor,json=shamLinkNeighbor" json:"sham_link_neighbor,omitempty"`
}

func (m *OspfShShamLinks) Reset()                    { *m = OspfShShamLinks{} }
func (m *OspfShShamLinks) String() string            { return proto.CompactTextString(m) }
func (*OspfShShamLinks) ProtoMessage()               {}
func (*OspfShShamLinks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShShamLinks) GetShamLinkName() string {
	if m != nil {
		return m.ShamLinkName
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkNeighborId() string {
	if m != nil {
		return m.ShamLinkNeighborId
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkSourceAddress() string {
	if m != nil {
		return m.ShamLinkSourceAddress
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkDestAddress() string {
	if m != nil {
		return m.ShamLinkDestAddress
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkState() string {
	if m != nil {
		return m.ShamLinkState
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkDemandCircuit() bool {
	if m != nil {
		return m.ShamLinkDemandCircuit
	}
	return false
}

func (m *OspfShShamLinks) GetShamLinkDcBitlessLsa() uint32 {
	if m != nil {
		return m.ShamLinkDcBitlessLsa
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkIfindex() uint32 {
	if m != nil {
		return m.ShamLinkIfindex
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkArea() string {
	if m != nil {
		return m.ShamLinkArea
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkCost() uint32 {
	if m != nil {
		return m.ShamLinkCost
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkTransmissionDelay() uint32 {
	if m != nil {
		return m.ShamLinkTransmissionDelay
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkHelloInterval() uint32 {
	if m != nil {
		return m.ShamLinkHelloInterval
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkHelloIntervalMs() uint32 {
	if m != nil {
		return m.ShamLinkHelloIntervalMs
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkDeadInterval() uint32 {
	if m != nil {
		return m.ShamLinkDeadInterval
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkWaitInterval() uint32 {
	if m != nil {
		return m.ShamLinkWaitInterval
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkRetransmissionInterval() uint32 {
	if m != nil {
		return m.ShamLinkRetransmissionInterval
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkNextHello() uint32 {
	if m != nil {
		return m.ShamLinkNextHello
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkNextHelloMs() uint32 {
	if m != nil {
		return m.ShamLinkNextHelloMs
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkPassive() bool {
	if m != nil {
		return m.ShamLinkPassive
	}
	return false
}

func (m *OspfShShamLinks) GetShamLinkAuthenticationType() string {
	if m != nil {
		return m.ShamLinkAuthenticationType
	}
	return ""
}

func (m *OspfShShamLinks) GetShamLinkYoungestMdKey() bool {
	if m != nil {
		return m.ShamLinkYoungestMdKey
	}
	return false
}

func (m *OspfShShamLinks) GetShamLinkYoungestMdKeyId() uint32 {
	if m != nil {
		return m.ShamLinkYoungestMdKeyId
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkOldMdKeyCount() uint32 {
	if m != nil {
		return m.ShamLinkOldMdKeyCount
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkMdKeyList() []uint32 {
	if m != nil {
		return m.ShamLinkMdKeyList
	}
	return nil
}

func (m *OspfShShamLinks) GetShamLinkKeychainId() uint64 {
	if m != nil {
		return m.ShamLinkKeychainId
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkNsfEnabled() bool {
	if m != nil {
		return m.ShamLinkNsfEnabled
	}
	return false
}

func (m *OspfShShamLinks) GetShamLinkNsf() bool {
	if m != nil {
		return m.ShamLinkNsf
	}
	return false
}

func (m *OspfShShamLinks) GetShamLinkLastNsf() uint32 {
	if m != nil {
		return m.ShamLinkLastNsf
	}
	return 0
}

func (m *OspfShShamLinks) GetShamLinkNeighbor() *OspfShSlinkNeighbor {
	if m != nil {
		return m.ShamLinkNeighbor
	}
	return nil
}

// OSPF Neighbor Retransmission Information
type OspfShNeighborRetrans struct {
	// Number of DBD retransmissions during last exchange
	DbdRetransmissionCount uint32 `protobuf:"varint,1,opt,name=dbd_retransmission_count,json=dbdRetransmissionCount" json:"dbd_retransmission_count,omitempty"`
	// Total number of DBD retransmissions for this neighbor
	DbdRetransmissionTotalCount uint32 `protobuf:"varint,2,opt,name=dbd_retransmission_total_count,json=dbdRetransmissionTotalCount" json:"dbd_retransmission_total_count,omitempty"`
	// Area scope LSA's flood index
	AreaFloodingIndex uint32 `protobuf:"varint,3,opt,name=area_flooding_index,json=areaFloodingIndex" json:"area_flooding_index,omitempty"`
	// AS scope LSA's flood index
	AsFloodIndex uint32 `protobuf:"varint,4,opt,name=as_flood_index,json=asFloodIndex" json:"as_flood_index,omitempty"`
	// Retransmission queue length
	NeighborRetransmissionCount uint32 `protobuf:"varint,5,opt,name=neighbor_retransmission_count,json=neighborRetransmissionCount" json:"neighbor_retransmission_count,omitempty"`
	// Number of retransmissions for this neighbor
	NumberOfRetransmissions uint32 `protobuf:"varint,6,opt,name=number_of_retransmissions,json=numberOfRetransmissions" json:"number_of_retransmissions,omitempty"`
	// First flood item for area scope LSAs
	AreaFirstFloodInformation uint32 `protobuf:"varint,7,opt,name=area_first_flood_information,json=areaFirstFloodInformation" json:"area_first_flood_information,omitempty"`
	// Index of the first flood item for area scope LSAs
	AreaFirstFloodInformationIndex uint32 `protobuf:"varint,8,opt,name=area_first_flood_information_index,json=areaFirstFloodInformationIndex" json:"area_first_flood_information_index,omitempty"`
	// First flood item for AS scope LSAs
	AsFirstFloodInformation uint32 `protobuf:"varint,9,opt,name=as_first_flood_information,json=asFirstFloodInformation" json:"as_first_flood_information,omitempty"`
	// Index for first flood item for AS scope LSAs
	AsFirstFloodInformationIndex uint32 `protobuf:"varint,10,opt,name=as_first_flood_information_index,json=asFirstFloodInformationIndex" json:"as_first_flood_information_index,omitempty"`
	// Next flood item for area scope LSAs
	AreaNextFloodInformation uint32 `protobuf:"varint,11,opt,name=area_next_flood_information,json=areaNextFloodInformation" json:"area_next_flood_information,omitempty"`
	// Index of next flood item for Area scope LSAs
	AreaNextFloodInformationIndex uint32 `protobuf:"varint,12,opt,name=area_next_flood_information_index,json=areaNextFloodInformationIndex" json:"area_next_flood_information_index,omitempty"`
	// Next flood item for AS scope LSAs
	AsNextFloodInformation uint32 `protobuf:"varint,13,opt,name=as_next_flood_information,json=asNextFloodInformation" json:"as_next_flood_information,omitempty"`
	// Index of next flood item for AS scope LSAs
	AsNextFloodInformationIndex uint32 `protobuf:"varint,14,opt,name=as_next_flood_information_index,json=asNextFloodInformationIndex" json:"as_next_flood_information_index,omitempty"`
	// Number of LSAs sent in last retransmission
	LastRetransmissionLength uint32 `protobuf:"varint,15,opt,name=last_retransmission_length,json=lastRetransmissionLength" json:"last_retransmission_length,omitempty"`
	// Maximum number of LSAs sent in a retransmission
	MaximumRetransmissionLength uint32 `protobuf:"varint,16,opt,name=maximum_retransmission_length,json=maximumRetransmissionLength" json:"maximum_retransmission_length,omitempty"`
	// Last retransmission scan time (ms)
	LastRetransmissionTime uint32 `protobuf:"varint,17,opt,name=last_retransmission_time,json=lastRetransmissionTime" json:"last_retransmission_time,omitempty"`
	// Maximum retransmission scan time (ms)
	MaximumRetransmissionTime uint32 `protobuf:"varint,18,opt,name=maximum_retransmission_time,json=maximumRetransmissionTime" json:"maximum_retransmission_time,omitempty"`
	// Time until next LSA retransmission (ms)
	LsaRetransmissionTimer uint32 `protobuf:"varint,19,opt,name=lsa_retransmission_timer,json=lsaRetransmissionTimer" json:"lsa_retransmission_timer,omitempty"`
}

func (m *OspfShNeighborRetrans) Reset()                    { *m = OspfShNeighborRetrans{} }
func (m *OspfShNeighborRetrans) String() string            { return proto.CompactTextString(m) }
func (*OspfShNeighborRetrans) ProtoMessage()               {}
func (*OspfShNeighborRetrans) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

// Sham Link Neighbor Information
type OspfShSlinkNeighbor struct {
	// If true Hellos suppressed
	ShamLinkSuppressHello bool `protobuf:"varint,1,opt,name=sham_link_suppress_hello,json=shamLinkSuppressHello" json:"sham_link_suppress_hello,omitempty"`
	// Adjacency state
	ShamLinkState string `protobuf:"bytes,2,opt,name=sham_link_state,json=shamLinkState" json:"sham_link_state,omitempty"`
	// Neighbor retransmission info
	ShamLinkRetransmissoin *OspfShNeighborRetrans `protobuf:"bytes,3,opt,name=sham_link_retransmissoin,json=shamLinkRetransmissoin" json:"sham_link_retransmissoin,omitempty"`
}

func (m *OspfShSlinkNeighbor) Reset()                    { *m = OspfShSlinkNeighbor{} }
func (m *OspfShSlinkNeighbor) String() string            { return proto.CompactTextString(m) }
func (*OspfShSlinkNeighbor) ProtoMessage()               {}
func (*OspfShSlinkNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShSlinkNeighbor) GetShamLinkSuppressHello() bool {
	if m != nil {
		return m.ShamLinkSuppressHello
	}
	return false
}

func (m *OspfShSlinkNeighbor) GetShamLinkState() string {
	if m != nil {
		return m.ShamLinkState
	}
	return ""
}

func (m *OspfShSlinkNeighbor) GetShamLinkRetransmissoin() *OspfShNeighborRetrans {
	if m != nil {
		return m.ShamLinkRetransmissoin
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShShamLinks_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_vrf_information.sham_links.sham_link.ospf_sh_sham_links_KEYS")
	proto.RegisterType((*OspfShShamLinks)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_vrf_information.sham_links.sham_link.ospf_sh_sham_links")
	proto.RegisterType((*OspfShNeighborRetrans)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_vrf_information.sham_links.sham_link.ospf_sh_neighbor_retrans")
	proto.RegisterType((*OspfShSlinkNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_vrf_information.sham_links.sham_link.ospf_sh_slink_neighbor")
}

func init() { proto.RegisterFile("ospf_sh_sham_links.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x97, 0xdf, 0x4f, 0x1c, 0x37,
	0x10, 0xc7, 0x75, 0x24, 0x4d, 0x82, 0x81, 0x10, 0x4c, 0x02, 0x06, 0x42, 0x7a, 0x9c, 0xaa, 0x0a,
	0xb5, 0xd2, 0xb5, 0x05, 0x92, 0x40, 0x02, 0x49, 0xf9, 0x59, 0x2e, 0x1c, 0x49, 0x05, 0x48, 0x55,
	0x9e, 0x2c, 0xdf, 0xae, 0x97, 0x73, 0xd9, 0x5d, 0x9f, 0xd6, 0xbe, 0x2b, 0xf7, 0xd0, 0x97, 0x2a,
	0xff, 0x50, 0x9f, 0xfb, 0x50, 0xf5, 0x3f, 0xab, 0x3c, 0xde, 0xdb, 0x5d, 0xdf, 0xed, 0xe5, 0xb1,
	0x7d, 0x41, 0xa7, 0x9d, 0xef, 0x67, 0x66, 0x3c, 0x1e, 0x8f, 0x0d, 0x22, 0x52, 0x75, 0x02, 0xaa,
	0xda, 0x54, 0xb5, 0x59, 0x44, 0x43, 0x11, 0xdf, 0xa8, 0x7a, 0x27, 0x91, 0x5a, 0xe2, 0x5f, 0x3d,
	0xa1, 0x3c, 0x49, 0x85, 0x54, 0xf4, 0x36, 0xa1, 0xa2, 0xd3, 0xdb, 0xa2, 0xa0, 0x95, 0x1d, 0x9e,
	0xd4, 0xcd, 0x2f, 0xa3, 0xf3, 0xb8, 0x52, 0x5c, 0x0d, 0x7e, 0xd5, 0x7b, 0x49, 0x00, 0x7f, 0xea,
	0x22, 0xd6, 0x3c, 0x09, 0x98, 0xc7, 0x69, 0x2f, 0x09, 0xa8, 0x88, 0x03, 0x99, 0x44, 0x4c, 0x0b,
	0x19, 0xd7, 0x0b, 0x71, 0xb2, 0x9f, 0xb5, 0xdf, 0xd1, 0xe2, 0x68, 0x1e, 0xf4, 0xec, 0xf8, 0xe3,
	0x25, 0x5e, 0x43, 0xd3, 0xa9, 0x77, 0x1a, 0xb3, 0x88, 0x93, 0x4a, 0xb5, 0xb2, 0x3e, 0x79, 0x31,
	0x95, 0x7e, 0x7b, 0xcf, 0x22, 0x8e, 0x97, 0xd0, 0x03, 0x13, 0x05, 0xcc, 0x13, 0x60, 0xbe, 0xdf,
	0x4b, 0x02, 0x30, 0x7d, 0x85, 0x1e, 0x66, 0x0e, 0xad, 0xe0, 0x0e, 0x08, 0xa6, 0xcd, 0xd7, 0xa6,
	0x88, 0x6f, 0x8c, 0xaa, 0xf6, 0x69, 0x06, 0xe1, 0xd1, 0xf8, 0x25, 0xf0, 0xc6, 0x28, 0x8c, 0x7f,
	0x40, 0x4f, 0x0a, 0x2a, 0x2e, 0xae, 0xdb, 0x2d, 0x99, 0x50, 0xe1, 0x93, 0x4d, 0x10, 0xe3, 0x4c,
	0x9c, 0x9a, 0x1a, 0x3e, 0x7e, 0x89, 0x48, 0x8e, 0x28, 0xd9, 0x4d, 0x3c, 0x4e, 0x99, 0xef, 0x27,
	0x5c, 0x29, 0xb2, 0x05, 0xd4, 0x93, 0x01, 0x75, 0x09, 0xd6, 0x7d, 0x6b, 0xc4, 0x9b, 0x68, 0x21,
	0x07, 0x7d, 0xae, 0x74, 0x86, 0x3d, 0x07, 0x6c, 0x7e, 0x80, 0x1d, 0x71, 0xa5, 0x07, 0xd0, 0xd7,
	0x68, 0xb6, 0x10, 0x4d, 0x33, 0xcd, 0xc9, 0x0b, 0x50, 0xcf, 0x64, 0x41, 0xcc, 0x47, 0x37, 0x2b,
	0x9f, 0x47, 0x2c, 0xf6, 0xa9, 0x27, 0x12, 0xaf, 0x2b, 0x34, 0x79, 0x59, 0xad, 0xac, 0x3f, 0xc8,
	0xb3, 0x3a, 0x02, 0xeb, 0xa1, 0x35, 0xe2, 0x17, 0x0e, 0xe8, 0xd1, 0x96, 0xd0, 0xa1, 0xd9, 0xaf,
	0x50, 0x31, 0xb2, 0x5d, 0xad, 0xac, 0xcf, 0x5c, 0x3c, 0xce, 0x40, 0xef, 0xc0, 0x1a, 0x9b, 0x8a,
	0xe1, 0x6f, 0xd0, 0x5c, 0xce, 0x89, 0x40, 0xc4, 0x3e, 0xbf, 0x25, 0x3b, 0x00, 0xcc, 0x0e, 0x80,
	0x86, 0xfd, 0xec, 0xee, 0x05, 0x4b, 0x38, 0x23, 0xaf, 0xdc, 0xbd, 0xd8, 0x4f, 0x38, 0x73, 0x55,
	0x9e, 0x54, 0x9a, 0xbc, 0x06, 0x77, 0x99, 0xea, 0x50, 0x2a, 0x8d, 0xdf, 0xa2, 0xa7, 0xb9, 0x4a,
	0x27, 0x2c, 0x56, 0x91, 0x50, 0x4a, 0xc8, 0x98, 0xfa, 0x3c, 0x64, 0x7d, 0xb2, 0x0b, 0xcc, 0xd2,
	0x80, 0xb9, 0x2a, 0x28, 0x8e, 0x8c, 0xc0, 0xad, 0x54, 0x9b, 0x87, 0xa1, 0xa4, 0xd0, 0xf0, 0x3d,
	0x16, 0x92, 0x3d, 0x80, 0xb3, 0x4a, 0x9d, 0x1a, 0x6b, 0x23, 0x35, 0xe2, 0x5d, 0xb4, 0x32, 0x0e,
	0xa4, 0x91, 0x22, 0x6f, 0x80, 0x5d, 0x2c, 0x65, 0xcf, 0x15, 0x7e, 0x8e, 0x16, 0x8b, 0x1b, 0xc4,
	0xfc, 0x3c, 0xea, 0xdb, 0xa1, 0x32, 0x73, 0xe6, 0x67, 0x41, 0x1d, 0xec, 0x37, 0x26, 0x74, 0x8e,
	0xfd, 0xe8, 0x62, 0xbf, 0x30, 0xa1, 0x33, 0xac, 0x81, 0xd6, 0x72, 0x2c, 0xe1, 0x4e, 0x9d, 0x32,
	0x07, 0xfb, 0xe0, 0xe0, 0xd9, 0xc0, 0xc1, 0x85, 0x23, 0xcb, 0x5c, 0x7d, 0x87, 0x1e, 0x17, 0x8f,
	0xc8, 0xad, 0xb6, 0x6b, 0x27, 0x07, 0x40, 0xcf, 0xe5, 0x27, 0xe4, 0x56, 0xc3, 0x9a, 0xf1, 0x56,
	0x31, 0xe5, 0x1c, 0x30, 0x35, 0x3a, 0x04, 0x66, 0x7e, 0x84, 0x39, 0x57, 0x6e, 0x3f, 0x75, 0x98,
	0x52, 0xa2, 0xc7, 0xc9, 0x11, 0x74, 0x6e, 0xd6, 0x4f, 0x3f, 0xdb, 0xcf, 0x78, 0x1f, 0xad, 0x16,
	0xfa, 0xa9, 0xab, 0xdb, 0x3c, 0xd6, 0xc2, 0x83, 0x11, 0x45, 0x75, 0xbf, 0xc3, 0xc9, 0x31, 0xb4,
	0xd7, 0x72, 0xd6, 0x5e, 0x8e, 0xe4, 0xaa, 0xdf, 0xe1, 0x78, 0x1b, 0x2d, 0xe5, 0x2e, 0xfa, 0xb2,
	0x1b, 0x5f, 0x9b, 0x03, 0x19, 0xf9, 0xf4, 0x86, 0xf7, 0xc9, 0x89, 0x7b, 0x60, 0x3e, 0xa6, 0xe6,
	0x73, 0xff, 0x8c, 0xf7, 0xf1, 0x5e, 0xb1, 0x01, 0x87, 0x48, 0x33, 0x39, 0x7e, 0x72, 0xfb, 0xc0,
	0x81, 0x1b, 0x3e, 0xde, 0x41, 0xcb, 0x39, 0x2e, 0x43, 0x7f, 0x40, 0x7a, 0xb2, 0x1b, 0x6b, 0x72,
	0xea, 0x36, 0xe0, 0x87, 0xd0, 0x07, 0xee, 0xd0, 0x18, 0xf1, 0xf7, 0xc5, 0x61, 0x95, 0x62, 0xa1,
	0x50, 0x9a, 0x34, 0xaa, 0x77, 0x8a, 0x5b, 0x01, 0x48, 0x53, 0x28, 0xed, 0x8e, 0xb7, 0x1b, 0xde,
	0xf7, 0xda, 0x4c, 0xc4, 0x26, 0xc9, 0x77, 0xd5, 0xca, 0xfa, 0xdd, 0x7c, 0xbc, 0x9d, 0xa5, 0xa6,
	0x86, 0x3f, 0x34, 0x11, 0x55, 0x40, 0x79, 0xcc, 0x5a, 0x21, 0xf7, 0xc9, 0x19, 0x14, 0x25, 0x9f,
	0x88, 0x2a, 0x38, 0xb6, 0x16, 0x5c, 0x43, 0x33, 0x0e, 0x42, 0x9a, 0x20, 0x9d, 0x2a, 0x48, 0xf1,
	0xb7, 0x08, 0xe7, 0x9a, 0x90, 0x29, 0x0d, 0xc2, 0x73, 0x77, 0x5e, 0x34, 0x99, 0xd2, 0x46, 0xfc,
	0x67, 0xa5, 0xa8, 0x1e, 0x8c, 0x65, 0xf2, 0xbe, 0x5a, 0x59, 0x9f, 0xda, 0xf8, 0xa3, 0x52, 0xff,
	0xef, 0x2e, 0xb7, 0x7a, 0x76, 0xb3, 0x38, 0xa9, 0x5c, 0x3c, 0x1a, 0xbe, 0x18, 0x6a, 0xff, 0x4c,
	0xe6, 0xd7, 0x71, 0x76, 0x91, 0xa4, 0x27, 0x0f, 0x6f, 0x23, 0xe2, 0xb7, 0xfc, 0xe1, 0x83, 0x68,
	0xb7, 0xbc, 0x02, 0x35, 0x58, 0xf0, 0x5b, 0xbe, 0x7b, 0x00, 0xed, 0x9e, 0x1f, 0xa2, 0x67, 0x25,
	0xa4, 0x96, 0x9a, 0x85, 0x29, 0x3f, 0x01, 0xfc, 0xca, 0x08, 0x7f, 0x65, 0x34, 0xd6, 0x49, 0x1d,
	0xcd, 0x9b, 0xa9, 0x4b, 0x83, 0x50, 0x4a, 0x5f, 0xc4, 0xd7, 0xd4, 0x4e, 0xeb, 0x3b, 0xf6, 0x04,
	0x1b, 0xd3, 0x49, 0x6a, 0x69, 0x0c, 0xe6, 0x35, 0x53, 0x56, 0x9d, 0x4a, 0xef, 0xda, 0x49, 0xcc,
	0x14, 0x08, 0xad, 0xea, 0x00, 0xad, 0x0e, 0x2f, 0xd4, 0x5d, 0xd9, 0x17, 0x36, 0xb3, 0xac, 0x68,
	0x25, 0xcb, 0x7b, 0x85, 0x96, 0xe2, 0x6e, 0xd4, 0xe2, 0x09, 0x95, 0xc1, 0x90, 0x13, 0x45, 0xee,
	0xd9, 0x93, 0x64, 0x05, 0x1f, 0x02, 0x97, 0x57, 0xe6, 0x26, 0xb0, 0xab, 0x12, 0x89, 0xd2, 0x59,
	0xb6, 0xd9, 0x7e, 0x92, 0xfb, 0xf6, 0x26, 0x80, 0xe5, 0x19, 0x49, 0x9a, 0x7a, 0x26, 0xc0, 0xef,
	0x50, 0xed, 0x73, 0x0e, 0xd2, 0xa5, 0x3f, 0xb0, 0x53, 0x72, 0xac, 0x1b, 0x5b, 0x8c, 0xd7, 0x68,
	0xd9, 0x94, 0x6c, 0x4c, 0x2a, 0x93, 0x76, 0x25, 0x4c, 0x95, 0x27, 0x72, 0x82, 0xaa, 0xe3, 0xe1,
	0x34, 0x0d, 0x04, 0x2e, 0x9e, 0x8e, 0x71, 0x61, 0x93, 0xd8, 0x43, 0x2b, 0xb0, 0x20, 0x18, 0xba,
	0xa3, 0x59, 0x4c, 0x81, 0x0b, 0x62, 0x24, 0x66, 0xf2, 0x8e, 0xa4, 0x71, 0x8a, 0xd6, 0x3e, 0x83,
	0xa7, 0x79, 0x4c, 0x83, 0x93, 0xd5, 0x71, 0x4e, 0x6c, 0x22, 0x3b, 0x68, 0x89, 0xa9, 0x71, 0x69,
	0xcc, 0xd8, 0x86, 0x67, 0xaa, 0x34, 0x89, 0x23, 0xf4, 0xe5, 0x58, 0x34, 0x4d, 0xe1, 0xa1, 0xed,
	0xab, 0x72, 0x07, 0x36, 0x81, 0x5d, 0xb4, 0x0c, 0x43, 0x66, 0xa8, 0x2f, 0x43, 0x1e, 0x5f, 0xeb,
	0x36, 0x99, 0xb5, 0x85, 0x30, 0x0a, 0xb7, 0xa9, 0x9a, 0x60, 0x37, 0x9d, 0x1d, 0xb1, 0x5b, 0x11,
	0x75, 0xa3, 0x31, 0x0e, 0x1e, 0xd9, 0x0c, 0x52, 0x51, 0xa9, 0x8f, 0x6d, 0x44, 0xca, 0x32, 0xd0,
	0x22, 0xe2, 0x64, 0xce, 0x56, 0x60, 0x34, 0xfe, 0x95, 0x88, 0x38, 0x7e, 0x83, 0x56, 0xc6, 0x44,
	0x07, 0x18, 0xdb, 0xb6, 0x2e, 0x8d, 0x0d, 0xbc, 0x89, 0xac, 0x58, 0x19, 0x9b, 0x90, 0xf9, 0x34,
	0xb2, 0x62, 0xa3, 0x60, 0x52, 0xfb, 0x6b, 0x02, 0x2d, 0x94, 0x0f, 0xbc, 0xa1, 0x57, 0x6f, 0xb7,
	0xd3, 0x31, 0xaf, 0xd3, 0xf4, 0x25, 0x50, 0x71, 0xaf, 0xcb, 0xcb, 0xd4, 0x6a, 0x5f, 0x03, 0x25,
	0x0f, 0xd8, 0x89, 0xb2, 0x07, 0xec, 0xdf, 0x95, 0x62, 0x84, 0x42, 0xf2, 0x52, 0xc4, 0x30, 0xa9,
	0xa6, 0x36, 0x3e, 0xfd, 0x2f, 0x93, 0x7f, 0x78, 0xc6, 0x5d, 0x2c, 0x94, 0x3c, 0x98, 0xa4, 0x88,
	0x5b, 0xf7, 0xe0, 0x5f, 0xaf, 0xcd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x98, 0x53, 0x0e,
	0x96, 0x0d, 0x00, 0x00,
}
