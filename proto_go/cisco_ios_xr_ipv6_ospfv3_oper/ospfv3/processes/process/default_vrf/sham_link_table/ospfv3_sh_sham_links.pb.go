// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_sh_sham_links.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_sham_link_table

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

// OSPFv3 Sham Link
type Ospfv3ShShamLinks_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3ShShamLinks_KEYS) Reset()         { *m = Ospfv3ShShamLinks_KEYS{} }
func (m *Ospfv3ShShamLinks_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShShamLinks_KEYS) ProtoMessage()    {}
func (*Ospfv3ShShamLinks_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_sh_sham_links_1e5e05a2f9613144, []int{0}
}
func (m *Ospfv3ShShamLinks_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShShamLinks_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3ShShamLinks_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShShamLinks_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ospfv3ShShamLinks_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShShamLinks_KEYS.Merge(dst, src)
}
func (m *Ospfv3ShShamLinks_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShShamLinks_KEYS.Size(m)
}
func (m *Ospfv3ShShamLinks_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShShamLinks_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShShamLinks_KEYS proto.InternalMessageInfo

func (m *Ospfv3ShShamLinks_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type Ospfv3ShShamLinks struct {
	// Neighbor on other end of this sham link
	ShamLinkNeighborId string `protobuf:"bytes,50,opt,name=sham_link_neighbor_id,json=shamLinkNeighborId,proto3" json:"sham_link_neighbor_id,omitempty"`
	// Sham-link source
	ShamLinkSourceAddress string `protobuf:"bytes,51,opt,name=sham_link_source_address,json=shamLinkSourceAddress,proto3" json:"sham_link_source_address,omitempty"`
	// Sham-link dest
	ShamLinkDestAddress string `protobuf:"bytes,52,opt,name=sham_link_dest_address,json=shamLinkDestAddress,proto3" json:"sham_link_dest_address,omitempty"`
	// OSPF interface state for the sham link
	ShamLinkState string `protobuf:"bytes,53,opt,name=sham_link_state,json=shamLinkState,proto3" json:"sham_link_state,omitempty"`
	// If true, the link runs as demand circuit
	ShamLinkDemandCircuit bool `protobuf:"varint,54,opt,name=sham_link_demand_circuit,json=shamLinkDemandCircuit,proto3" json:"sham_link_demand_circuit,omitempty"`
	// Number of LSA's with demand circuit bit not set
	ShamLinkDcBitlessLsa uint32 `protobuf:"varint,55,opt,name=sham_link_dc_bitless_lsa,json=shamLinkDcBitlessLsa,proto3" json:"sham_link_dc_bitless_lsa,omitempty"`
	// Sham-link ifindex
	ShamLinkIfindex uint32 `protobuf:"varint,56,opt,name=sham_link_ifindex,json=shamLinkIfindex,proto3" json:"sham_link_ifindex,omitempty"`
	// Area id
	ShamLinkArea string `protobuf:"bytes,57,opt,name=sham_link_area,json=shamLinkArea,proto3" json:"sham_link_area,omitempty"`
	// Cost of the sham link
	ShamLinkCost uint32 `protobuf:"varint,58,opt,name=sham_link_cost,json=shamLinkCost,proto3" json:"sham_link_cost,omitempty"`
	// Transmission delay in seconds
	ShamLinkTransmissionDelay uint32 `protobuf:"varint,59,opt,name=sham_link_transmission_delay,json=shamLinkTransmissionDelay,proto3" json:"sham_link_transmission_delay,omitempty"`
	// Hello interval (s)
	ShamLinkHelloInterval uint32 `protobuf:"varint,60,opt,name=sham_link_hello_interval,json=shamLinkHelloInterval,proto3" json:"sham_link_hello_interval,omitempty"`
	// Dead interval (s)
	ShamLinkDeadInterval uint32 `protobuf:"varint,61,opt,name=sham_link_dead_interval,json=shamLinkDeadInterval,proto3" json:"sham_link_dead_interval,omitempty"`
	// Wait interval (s)
	ShamLinkWaitInterval uint32 `protobuf:"varint,62,opt,name=sham_link_wait_interval,json=shamLinkWaitInterval,proto3" json:"sham_link_wait_interval,omitempty"`
	// Retransmission interval (s)
	ShamLinkRetransmissionInterval uint32 `protobuf:"varint,63,opt,name=sham_link_retransmission_interval,json=shamLinkRetransmissionInterval,proto3" json:"sham_link_retransmission_interval,omitempty"`
	// Time until next hello (s)
	ShamLinkNextHello uint32 `protobuf:"varint,64,opt,name=sham_link_next_hello,json=shamLinkNextHello,proto3" json:"sham_link_next_hello,omitempty"`
	// If true, interface is passive
	ShamLinkPassive bool `protobuf:"varint,65,opt,name=sham_link_passive,json=shamLinkPassive,proto3" json:"sham_link_passive,omitempty"`
	// If true, sham link IP security is required
	IsShamLinkIpSecurityRequired bool `protobuf:"varint,66,opt,name=is_sham_link_ip_security_required,json=isShamLinkIpSecurityRequired,proto3" json:"is_sham_link_ip_security_required,omitempty"`
	// If true, Sham link IP security is active
	IsShamLinkIpSecurityActive bool `protobuf:"varint,67,opt,name=is_sham_link_ip_security_active,json=isShamLinkIpSecurityActive,proto3" json:"is_sham_link_ip_security_active,omitempty"`
	// If true, sham link authentication is enabled
	IsShamLinkAuthenticationEnabled bool `protobuf:"varint,68,opt,name=is_sham_link_authentication_enabled,json=isShamLinkAuthenticationEnabled,proto3" json:"is_sham_link_authentication_enabled,omitempty"`
	// Sham link authentication spi
	VirtualLinkAuthenticationSpi uint32 `protobuf:"varint,69,opt,name=virtual_link_authentication_spi,json=virtualLinkAuthenticationSpi,proto3" json:"virtual_link_authentication_spi,omitempty"`
	// Sham link authentication transmit
	ShamLinkAuthenticationTransmit uint32 `protobuf:"varint,70,opt,name=sham_link_authentication_transmit,json=shamLinkAuthenticationTransmit,proto3" json:"sham_link_authentication_transmit,omitempty"`
	// If true, sham link encryption is enabled
	IsShamLinkEncryptionEnabled bool `protobuf:"varint,71,opt,name=is_sham_link_encryption_enabled,json=isShamLinkEncryptionEnabled,proto3" json:"is_sham_link_encryption_enabled,omitempty"`
	// Sham link encryption spi
	ShamLinkEncryptionSpi uint32 `protobuf:"varint,72,opt,name=sham_link_encryption_spi,json=shamLinkEncryptionSpi,proto3" json:"sham_link_encryption_spi,omitempty"`
	// Sham link encryption transmitted
	ShamLinkEncryptionTransmitted uint32 `protobuf:"varint,73,opt,name=sham_link_encryption_transmitted,json=shamLinkEncryptionTransmitted,proto3" json:"sham_link_encryption_transmitted,omitempty"`
	// Sham link encrypted authentication transmitted
	ShamLinkEncryptedAuthenticationTransmitted uint32 `protobuf:"varint,74,opt,name=sham_link_encrypted_authentication_transmitted,json=shamLinkEncryptedAuthenticationTransmitted,proto3" json:"sham_link_encrypted_authentication_transmitted,omitempty"`
	// If true,  enabled
	ShamLinkGrEnabled bool `protobuf:"varint,75,opt,name=sham_link_gr_enabled,json=shamLinkGrEnabled,proto3" json:"sham_link_gr_enabled,omitempty"`
	// If true, Gracefule restart in progress
	ShamLinkGr bool `protobuf:"varint,76,opt,name=sham_link_gr,json=shamLinkGr,proto3" json:"sham_link_gr,omitempty"`
	// Time in seconds since last GR
	ShamLinkLastGr uint32 `protobuf:"varint,77,opt,name=sham_link_last_gr,json=shamLinkLastGr,proto3" json:"sham_link_last_gr,omitempty"`
	// Neighbor information
	ShamLinkNeighbor     *Ospfv3ShSlinkNeighbor `protobuf:"bytes,78,opt,name=sham_link_neighbor,json=shamLinkNeighbor,proto3" json:"sham_link_neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Ospfv3ShShamLinks) Reset()         { *m = Ospfv3ShShamLinks{} }
func (m *Ospfv3ShShamLinks) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShShamLinks) ProtoMessage()    {}
func (*Ospfv3ShShamLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_sh_sham_links_1e5e05a2f9613144, []int{1}
}
func (m *Ospfv3ShShamLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShShamLinks.Unmarshal(m, b)
}
func (m *Ospfv3ShShamLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShShamLinks.Marshal(b, m, deterministic)
}
func (dst *Ospfv3ShShamLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShShamLinks.Merge(dst, src)
}
func (m *Ospfv3ShShamLinks) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShShamLinks.Size(m)
}
func (m *Ospfv3ShShamLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShShamLinks.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShShamLinks proto.InternalMessageInfo

func (m *Ospfv3ShShamLinks) GetShamLinkNeighborId() string {
	if m != nil {
		return m.ShamLinkNeighborId
	}
	return ""
}

func (m *Ospfv3ShShamLinks) GetShamLinkSourceAddress() string {
	if m != nil {
		return m.ShamLinkSourceAddress
	}
	return ""
}

func (m *Ospfv3ShShamLinks) GetShamLinkDestAddress() string {
	if m != nil {
		return m.ShamLinkDestAddress
	}
	return ""
}

func (m *Ospfv3ShShamLinks) GetShamLinkState() string {
	if m != nil {
		return m.ShamLinkState
	}
	return ""
}

func (m *Ospfv3ShShamLinks) GetShamLinkDemandCircuit() bool {
	if m != nil {
		return m.ShamLinkDemandCircuit
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetShamLinkDcBitlessLsa() uint32 {
	if m != nil {
		return m.ShamLinkDcBitlessLsa
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkIfindex() uint32 {
	if m != nil {
		return m.ShamLinkIfindex
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkArea() string {
	if m != nil {
		return m.ShamLinkArea
	}
	return ""
}

func (m *Ospfv3ShShamLinks) GetShamLinkCost() uint32 {
	if m != nil {
		return m.ShamLinkCost
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkTransmissionDelay() uint32 {
	if m != nil {
		return m.ShamLinkTransmissionDelay
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkHelloInterval() uint32 {
	if m != nil {
		return m.ShamLinkHelloInterval
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkDeadInterval() uint32 {
	if m != nil {
		return m.ShamLinkDeadInterval
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkWaitInterval() uint32 {
	if m != nil {
		return m.ShamLinkWaitInterval
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkRetransmissionInterval() uint32 {
	if m != nil {
		return m.ShamLinkRetransmissionInterval
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkNextHello() uint32 {
	if m != nil {
		return m.ShamLinkNextHello
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkPassive() bool {
	if m != nil {
		return m.ShamLinkPassive
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetIsShamLinkIpSecurityRequired() bool {
	if m != nil {
		return m.IsShamLinkIpSecurityRequired
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetIsShamLinkIpSecurityActive() bool {
	if m != nil {
		return m.IsShamLinkIpSecurityActive
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetIsShamLinkAuthenticationEnabled() bool {
	if m != nil {
		return m.IsShamLinkAuthenticationEnabled
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetVirtualLinkAuthenticationSpi() uint32 {
	if m != nil {
		return m.VirtualLinkAuthenticationSpi
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkAuthenticationTransmit() uint32 {
	if m != nil {
		return m.ShamLinkAuthenticationTransmit
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetIsShamLinkEncryptionEnabled() bool {
	if m != nil {
		return m.IsShamLinkEncryptionEnabled
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetShamLinkEncryptionSpi() uint32 {
	if m != nil {
		return m.ShamLinkEncryptionSpi
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkEncryptionTransmitted() uint32 {
	if m != nil {
		return m.ShamLinkEncryptionTransmitted
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkEncryptedAuthenticationTransmitted() uint32 {
	if m != nil {
		return m.ShamLinkEncryptedAuthenticationTransmitted
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkGrEnabled() bool {
	if m != nil {
		return m.ShamLinkGrEnabled
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetShamLinkGr() bool {
	if m != nil {
		return m.ShamLinkGr
	}
	return false
}

func (m *Ospfv3ShShamLinks) GetShamLinkLastGr() uint32 {
	if m != nil {
		return m.ShamLinkLastGr
	}
	return 0
}

func (m *Ospfv3ShShamLinks) GetShamLinkNeighbor() *Ospfv3ShSlinkNeighbor {
	if m != nil {
		return m.ShamLinkNeighbor
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
	return fileDescriptor_ospfv3_sh_sham_links_1e5e05a2f9613144, []int{2}
}
func (m *Ospfv3EdmNeighborRetrans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNeighborRetrans.Unmarshal(m, b)
}
func (m *Ospfv3EdmNeighborRetrans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNeighborRetrans.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmNeighborRetrans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNeighborRetrans.Merge(dst, src)
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

// Sham Link Neighbor Information
type Ospfv3ShSlinkNeighbor struct {
	// If true Hellos suppressed
	ShamLinkSuppressHello bool `protobuf:"varint,1,opt,name=sham_link_suppress_hello,json=shamLinkSuppressHello,proto3" json:"sham_link_suppress_hello,omitempty"`
	// Adjacency state
	ShamLinkState string `protobuf:"bytes,2,opt,name=sham_link_state,json=shamLinkState,proto3" json:"sham_link_state,omitempty"`
	// Neighbor retransmission info
	ShamLinkRetransmission *Ospfv3EdmNeighborRetrans `protobuf:"bytes,3,opt,name=sham_link_retransmission,json=shamLinkRetransmission,proto3" json:"sham_link_retransmission,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                  `json:"-"`
	XXX_unrecognized       []byte                    `json:"-"`
	XXX_sizecache          int32                     `json:"-"`
}

func (m *Ospfv3ShSlinkNeighbor) Reset()         { *m = Ospfv3ShSlinkNeighbor{} }
func (m *Ospfv3ShSlinkNeighbor) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShSlinkNeighbor) ProtoMessage()    {}
func (*Ospfv3ShSlinkNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_sh_sham_links_1e5e05a2f9613144, []int{3}
}
func (m *Ospfv3ShSlinkNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShSlinkNeighbor.Unmarshal(m, b)
}
func (m *Ospfv3ShSlinkNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShSlinkNeighbor.Marshal(b, m, deterministic)
}
func (dst *Ospfv3ShSlinkNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShSlinkNeighbor.Merge(dst, src)
}
func (m *Ospfv3ShSlinkNeighbor) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShSlinkNeighbor.Size(m)
}
func (m *Ospfv3ShSlinkNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShSlinkNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShSlinkNeighbor proto.InternalMessageInfo

func (m *Ospfv3ShSlinkNeighbor) GetShamLinkSuppressHello() bool {
	if m != nil {
		return m.ShamLinkSuppressHello
	}
	return false
}

func (m *Ospfv3ShSlinkNeighbor) GetShamLinkState() string {
	if m != nil {
		return m.ShamLinkState
	}
	return ""
}

func (m *Ospfv3ShSlinkNeighbor) GetShamLinkRetransmission() *Ospfv3EdmNeighborRetrans {
	if m != nil {
		return m.ShamLinkRetransmission
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3ShShamLinks_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.sham_link_table.ospfv3_sh_sham_links_KEYS")
	proto.RegisterType((*Ospfv3ShShamLinks)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.sham_link_table.ospfv3_sh_sham_links")
	proto.RegisterType((*Ospfv3EdmNeighborRetrans)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.sham_link_table.ospfv3_edm_neighbor_retrans")
	proto.RegisterType((*Ospfv3ShSlinkNeighbor)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.sham_link_table.ospfv3_sh_slink_neighbor")
}

func init() {
	proto.RegisterFile("ospfv3_sh_sham_links.proto", fileDescriptor_ospfv3_sh_sham_links_1e5e05a2f9613144)
}

var fileDescriptor_ospfv3_sh_sham_links_1e5e05a2f9613144 = []byte{
	// 1180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x97, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0xc7, 0xc7, 0x05, 0x4a, 0xa3, 0xc4, 0xf9, 0x50, 0xbe, 0xd4, 0x24, 0x25, 0x4e, 0xca, 0x74,
	0x4c, 0x67, 0x70, 0xa7, 0x09, 0x4d, 0x5b, 0x28, 0x29, 0x49, 0x9c, 0xa6, 0xa6, 0x26, 0xc3, 0xd8,
	0x99, 0x61, 0xb8, 0xda, 0x91, 0x77, 0x95, 0x58, 0xd3, 0xf5, 0xee, 0x22, 0xc9, 0x26, 0x79, 0x08,
	0x5e, 0x80, 0x2b, 0xee, 0x18, 0xde, 0x92, 0xd1, 0x91, 0xf6, 0x43, 0xeb, 0xf5, 0x1d, 0xdc, 0x65,
	0xa4, 0xdf, 0xff, 0xaf, 0x73, 0xce, 0xca, 0xe7, 0x28, 0x68, 0x2b, 0x96, 0xc9, 0xf5, 0xe4, 0xd0,
	0x93, 0x43, 0x4f, 0x0e, 0xe9, 0xc8, 0x0b, 0x79, 0xf4, 0x51, 0xb6, 0x12, 0x11, 0xab, 0x18, 0xf7,
	0x7c, 0x2e, 0xfd, 0xd8, 0xe3, 0xb1, 0xf4, 0x6e, 0x85, 0xc7, 0x93, 0xc9, 0x91, 0x67, 0xe9, 0x38,
	0x61, 0xa2, 0x65, 0xfe, 0xd6, 0xac, 0xcf, 0xa4, 0x64, 0x32, 0xfd, 0xab, 0x15, 0xb0, 0x6b, 0x3a,
	0x0e, 0x95, 0x37, 0x11, 0xd7, 0xad, 0xcc, 0xd4, 0x53, 0x74, 0x10, 0xb2, 0xfd, 0x63, 0xf4, 0xb0,
	0xea, 0x44, 0xef, 0xc3, 0xf9, 0xaf, 0x7d, 0xbc, 0x87, 0x16, 0xac, 0x87, 0x17, 0xd1, 0x11, 0x23,
	0xb5, 0x46, 0xad, 0x39, 0xd7, 0x9b, 0xb7, 0x6b, 0x97, 0x74, 0xc4, 0xf6, 0xff, 0x58, 0x44, 0x6b,
	0x55, 0x06, 0xf8, 0x39, 0x5a, 0xcf, 0xcf, 0x8a, 0x18, 0xbf, 0x19, 0x0e, 0x62, 0xe1, 0xf1, 0x80,
	0x1c, 0x80, 0x09, 0xd6, 0x9b, 0x5d, 0x1e, 0x7d, 0xbc, 0xb4, 0x5b, 0x9d, 0x00, 0xbf, 0x44, 0x24,
	0x97, 0xc8, 0x78, 0x2c, 0x7c, 0xe6, 0xd1, 0x20, 0x10, 0x4c, 0x4a, 0x72, 0x08, 0xaa, 0xf5, 0x54,
	0xd5, 0x87, 0xdd, 0x13, 0xb3, 0x89, 0x0f, 0xd1, 0x46, 0x2e, 0x0c, 0x98, 0x54, 0x99, 0xec, 0x1b,
	0x90, 0xad, 0xa6, 0xb2, 0x36, 0x93, 0x2a, 0x15, 0x3d, 0x41, 0x4b, 0x85, 0xd3, 0x14, 0x55, 0x8c,
	0xbc, 0x00, 0xba, 0x9e, 0x1d, 0xa2, 0x17, 0xdd, 0xa8, 0x02, 0x36, 0xa2, 0x51, 0xe0, 0xf9, 0x5c,
	0xf8, 0x63, 0xae, 0xc8, 0x51, 0xa3, 0xd6, 0x7c, 0x90, 0x47, 0xd5, 0x86, 0xdd, 0x33, 0xb3, 0x89,
	0x8f, 0x1c, 0xa1, 0xef, 0x0d, 0xb8, 0x0a, 0x75, 0x29, 0x43, 0x49, 0xc9, 0xcb, 0x46, 0xad, 0x59,
	0xef, 0xad, 0x65, 0x42, 0xff, 0xd4, 0x6c, 0x76, 0x25, 0xc5, 0x4f, 0xd1, 0x4a, 0xae, 0xe3, 0xd7,
	0x3c, 0x0a, 0xd8, 0x2d, 0x79, 0x05, 0x82, 0xa5, 0x54, 0xd0, 0x31, 0xcb, 0xf8, 0x4b, 0xb4, 0x98,
	0xb3, 0x54, 0x30, 0x4a, 0x5e, 0x43, 0x0e, 0x0b, 0x29, 0x78, 0x22, 0x18, 0x75, 0x29, 0x3f, 0x96,
	0x8a, 0x7c, 0x0b, 0x76, 0x19, 0x75, 0x16, 0x4b, 0x85, 0xdf, 0xa2, 0x9d, 0xc2, 0xed, 0x10, 0x34,
	0x92, 0x23, 0x2e, 0x25, 0x8f, 0x23, 0x2f, 0x60, 0x21, 0xbd, 0x23, 0xdf, 0x81, 0xe6, 0x61, 0xaa,
	0xb9, 0x2a, 0x10, 0x6d, 0x0d, 0xb8, 0x95, 0x1a, 0xb2, 0x30, 0x8c, 0x3d, 0x1e, 0x29, 0x26, 0x26,
	0x34, 0x24, 0x6f, 0x40, 0x9c, 0x55, 0xea, 0xbd, 0xde, 0xed, 0xd8, 0x4d, 0xfc, 0x02, 0x6d, 0x16,
	0x4b, 0x4c, 0x83, 0x5c, 0xf7, 0x7d, 0xa9, 0x50, 0x8c, 0x06, 0xd5, 0xb2, 0xdf, 0x29, 0x57, 0xb9,
	0xec, 0xd8, 0x95, 0xfd, 0x42, 0xb9, 0xca, 0x64, 0x1d, 0xb4, 0x97, 0xcb, 0x04, 0x73, 0x32, 0xcd,
	0x0c, 0xde, 0x82, 0xc1, 0x17, 0xa9, 0x41, 0xcf, 0xc1, 0x32, 0xab, 0x67, 0x68, 0xad, 0x78, 0xc9,
	0x6f, 0x95, 0x49, 0x9b, 0xfc, 0x00, 0xea, 0x95, 0xfc, 0x8e, 0xdf, 0x2a, 0xc8, 0xd8, 0xfd, 0xb6,
	0x09, 0x95, 0x92, 0x4f, 0x18, 0x39, 0x81, 0x5b, 0x94, 0x7d, 0xdb, 0x9f, 0xcd, 0x32, 0xbe, 0x40,
	0x7b, 0x5c, 0x7a, 0x85, 0xab, 0x90, 0x78, 0x92, 0xf9, 0x63, 0xc1, 0xd5, 0x9d, 0x27, 0xd8, 0x6f,
	0x63, 0x2e, 0x58, 0x40, 0x4e, 0x41, 0xbb, 0xc3, 0x65, 0x3f, 0xbd, 0x19, 0x49, 0xdf, 0x42, 0x3d,
	0xcb, 0xe0, 0x33, 0xb4, 0x3b, 0xd3, 0x88, 0xfa, 0x4a, 0x87, 0x70, 0x06, 0x36, 0x5b, 0x55, 0x36,
	0x27, 0x40, 0xe0, 0x2e, 0x7a, 0xec, 0x98, 0xd0, 0xb1, 0x1a, 0xb2, 0x48, 0x71, 0x9f, 0x2a, 0x5d,
	0x38, 0x16, 0xe9, 0x76, 0x12, 0x90, 0x36, 0x18, 0xed, 0xe6, 0x46, 0x27, 0x0e, 0x77, 0x6e, 0x30,
	0x7c, 0x8e, 0x76, 0x27, 0x5c, 0xa8, 0x31, 0x0d, 0x2b, 0xdd, 0x64, 0xc2, 0xc9, 0x39, 0xd4, 0x70,
	0xc7, 0x62, 0xd3, 0x56, 0xfd, 0x84, 0xbb, 0x9f, 0xb2, 0xe4, 0x61, 0x3f, 0x98, 0x22, 0xef, 0xdc,
	0x4f, 0xe9, 0xba, 0xd8, 0x5b, 0xac, 0x70, 0xbb, 0x54, 0x24, 0x16, 0xf9, 0xe2, 0x2e, 0x71, 0x72,
	0xbb, 0x80, 0xdc, 0xb6, 0xf3, 0xdc, 0xce, 0x33, 0x26, 0xcd, 0xcb, 0xf9, 0x09, 0x14, 0x2c, 0x74,
	0x42, 0xef, 0xdd, 0x9f, 0x40, 0x2e, 0xd6, 0x99, 0x5c, 0xa0, 0x46, 0xa5, 0x30, 0xcd, 0x42, 0xb1,
	0x80, 0x74, 0xc0, 0xe0, 0xd1, 0xb4, 0xc1, 0x55, 0x0e, 0xe1, 0x01, 0x6a, 0x4d, 0x19, 0xb1, 0x60,
	0x56, 0x71, 0xb4, 0xed, 0x8f, 0x60, 0xfb, 0xb4, 0x64, 0xcb, 0x82, 0xea, 0x42, 0xe9, 0x33, 0x9c,
	0x6b, 0x7f, 0x23, 0xb2, 0x02, 0x7d, 0x80, 0x02, 0x65, 0xd7, 0xfe, 0x42, 0xa4, 0x65, 0x69, 0xa0,
	0x85, 0xa2, 0x80, 0x74, 0x01, 0x44, 0x39, 0x88, 0xbf, 0x2a, 0xfe, 0x30, 0x42, 0x2a, 0x95, 0xc6,
	0x7e, 0x82, 0xc8, 0x16, 0x53, 0xac, 0x4b, 0xa5, 0xba, 0x10, 0xf8, 0xcf, 0x1a, 0xc2, 0xd3, 0xa3,
	0x85, 0x5c, 0x36, 0x6a, 0xcd, 0xf9, 0x83, 0xb0, 0xf5, 0xdf, 0x0f, 0xc9, 0x56, 0x61, 0xc0, 0x39,
	0x67, 0xf6, 0x96, 0xcb, 0x53, 0x6c, 0xff, 0xaf, 0x39, 0xb4, 0x6d, 0x71, 0x16, 0x8c, 0xf2, 0xc1,
	0x67, 0xfb, 0x0c, 0xbe, 0x44, 0x8f, 0x03, 0xaa, 0xe8, 0x80, 0x4a, 0xa6, 0x27, 0x95, 0x2f, 0x78,
	0xa2, 0xf2, 0x6d, 0xdb, 0x5f, 0x24, 0x4c, 0xda, 0x7a, 0x6f, 0x2f, 0x45, 0xdb, 0x19, 0xe9, 0x36,
	0x22, 0x89, 0x9b, 0x68, 0x59, 0xb7, 0x7d, 0xef, 0x3a, 0x8c, 0x63, 0xdd, 0x36, 0xf5, 0xac, 0xb8,
	0x67, 0xca, 0xa6, 0xd7, 0xdf, 0xe9, 0xe5, 0x4e, 0x3a, 0x2a, 0xa8, 0x74, 0xb8, 0x4f, 0xcc, 0x10,
	0xa0, 0xb2, 0x40, 0x35, 0xd1, 0x32, 0xa4, 0x58, 0xe4, 0x3e, 0x35, 0x7e, 0x7a, 0xbd, 0x40, 0xbe,
	0x46, 0xa4, 0x9c, 0x5d, 0x16, 0xfe, 0x67, 0xa0, 0xd8, 0xcc, 0xea, 0x34, 0x15, 0xf4, 0x52, 0x59,
	0x71, 0xdf, 0xcc, 0x37, 0x31, 0x2b, 0x3d, 0x2e, 0xa4, 0x32, 0x41, 0x91, 0xcf, 0x0b, 0xe9, 0xe9,
	0x65, 0x88, 0x49, 0xbf, 0x01, 0xca, 0xa4, 0x0d, 0xff, 0x01, 0xf0, 0xab, 0x2e, 0xef, 0xd4, 0xa4,
	0x60, 0x3e, 0x97, 0xd5, 0x24, 0xb7, 0x7e, 0x86, 0xd6, 0x5c, 0xca, 0x1a, 0x23, 0xd3, 0xe5, 0x8b,
	0x6c, 0xa9, 0x88, 0x05, 0xe3, 0xf9, 0x42, 0x11, 0x9d, 0xa8, 0xcb, 0xa4, 0x35, 0x5f, 0x30, 0x51,
	0xbb, 0xbc, 0xb1, 0x7f, 0x82, 0x96, 0x20, 0x55, 0x18, 0x38, 0xc6, 0xbd, 0x0e, 0x74, 0x5d, 0x2f,
	0xeb, 0x61, 0x63, 0xcc, 0x9f, 0xa3, 0xf5, 0x12, 0x67, 0xbd, 0x17, 0x81, 0xc6, 0x0e, 0x6d, 0xac,
	0xf7, 0x51, 0x9d, 0xca, 0xa2, 0xf1, 0x12, 0xa0, 0xf3, 0x54, 0xe6, 0xb6, 0x5f, 0xa3, 0x55, 0x87,
	0xb1, 0xa6, 0xcb, 0x40, 0x2e, 0x17, 0xc8, 0x2c, 0xda, 0x7c, 0x3c, 0x1a, 0xd3, 0x15, 0x13, 0x6d,
	0x68, 0x47, 0x63, 0x16, 0x6d, 0x89, 0xb3, 0xc6, 0xd8, 0x44, 0xeb, 0xd0, 0xc6, 0xfa, 0x0d, 0xda,
	0x82, 0x56, 0x51, 0x1a, 0xe2, 0x21, 0x8b, 0x6e, 0xd4, 0x90, 0xac, 0x82, 0x8e, 0x68, 0xc2, 0xbd,
	0x80, 0x5d, 0xd8, 0xc7, 0xa7, 0xe8, 0xd1, 0x88, 0xde, 0xf2, 0xd1, 0x78, 0x34, 0xc3, 0x60, 0x0d,
	0x0c, 0xb6, 0x2d, 0x54, 0xe9, 0xf1, 0x0a, 0x91, 0xaa, 0x08, 0x14, 0x1f, 0x31, 0xb2, 0x0e, 0xf2,
	0x8d, 0xe9, 0xf3, 0xaf, 0xf8, 0x88, 0xe1, 0x63, 0xb4, 0x3d, 0xe3, 0x74, 0x10, 0x6f, 0x98, 0xc7,
	0x56, 0xe5, 0xd9, 0xa0, 0xd7, 0x27, 0x4b, 0x5a, 0xa5, 0x15, 0x64, 0xd3, 0x9e, 0x2c, 0xe9, 0xb4,
	0x50, 0xec, 0xff, 0x7d, 0x0f, 0x91, 0x59, 0x1d, 0xad, 0xf4, 0x06, 0x1f, 0x27, 0x89, 0x7e, 0x2b,
	0xdb, 0x57, 0x4d, 0xcd, 0x7d, 0xed, 0xf6, 0xed, 0xae, 0x79, 0xd9, 0x54, 0x3c, 0xa7, 0xef, 0x55,
	0x3d, 0xa7, 0xff, 0xa9, 0x15, 0x4f, 0x70, 0xc3, 0x87, 0x8e, 0x34, 0x7f, 0x10, 0xff, 0x8f, 0x3d,
	0xbc, 0xaa, 0x29, 0xf7, 0x36, 0xaa, 0x9f, 0x79, 0x83, 0xfb, 0xf0, 0x7f, 0xd7, 0xe1, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0xfa, 0xdc, 0x85, 0x95, 0x0d, 0x00, 0x00,
}
