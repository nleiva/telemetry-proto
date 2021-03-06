// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_sh_leaf.proto

package cisco_ios_xr_fib_common_oper_mpls_forwarding_nodes_node_label_fib_informations_information

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

// Information about label leaf
type MplsShLeaf_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	LabelValue           uint32   `protobuf:"varint,2,opt,name=label_value,json=labelValue,proto3" json:"label_value,omitempty"`
	Eos                  string   `protobuf:"bytes,3,opt,name=eos,proto3" json:"eos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsShLeaf_KEYS) Reset()         { *m = MplsShLeaf_KEYS{} }
func (m *MplsShLeaf_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsShLeaf_KEYS) ProtoMessage()    {}
func (*MplsShLeaf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc, []int{0}
}
func (m *MplsShLeaf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsShLeaf_KEYS.Unmarshal(m, b)
}
func (m *MplsShLeaf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsShLeaf_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsShLeaf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsShLeaf_KEYS.Merge(dst, src)
}
func (m *MplsShLeaf_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsShLeaf_KEYS.Size(m)
}
func (m *MplsShLeaf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsShLeaf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsShLeaf_KEYS proto.InternalMessageInfo

func (m *MplsShLeaf_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsShLeaf_KEYS) GetLabelValue() uint32 {
	if m != nil {
		return m.LabelValue
	}
	return 0
}

func (m *MplsShLeaf_KEYS) GetEos() string {
	if m != nil {
		return m.Eos
	}
	return ""
}

type MplsShLeaf struct {
	// Local label
	LeafLocalLabel uint32 `protobuf:"varint,50,opt,name=leaf_local_label,json=leafLocalLabel,proto3" json:"leaf_local_label,omitempty"`
	// EOS bit
	EosBit uint32 `protobuf:"varint,51,opt,name=eos_bit,json=eosBit,proto3" json:"eos_bit,omitempty"`
	// Label-infos in FIB leaf
	LabelInformation []*MplsAdjInfo `protobuf:"bytes,52,rep,name=label_information,json=labelInformation,proto3" json:"label_information,omitempty"`
	// LDI-info in FIB leaf
	LdiInformation *MplsLdiInfo `protobuf:"bytes,53,opt,name=ldi_information,json=ldiInformation,proto3" json:"ldi_information,omitempty"`
	// Hardware info
	HardwareInformation []byte `protobuf:"bytes,54,opt,name=hardware_information,json=hardwareInformation,proto3" json:"hardware_information,omitempty"`
	// Number of references to the leaf
	LeafReferanceCount uint32 `protobuf:"varint,55,opt,name=leaf_referance_count,json=leafReferanceCount,proto3" json:"leaf_referance_count,omitempty"`
	// The leaf flags
	LeafFlags uint32 `protobuf:"varint,56,opt,name=leaf_flags,json=leafFlags,proto3" json:"leaf_flags,omitempty"`
	// Number of references to the pathlist
	PathListReferanceCount uint32 `protobuf:"varint,57,opt,name=path_list_referance_count,json=pathListReferanceCount,proto3" json:"path_list_referance_count,omitempty"`
	// The pathlist flags
	PathListFlags uint32 `protobuf:"varint,58,opt,name=path_list_flags,json=pathListFlags,proto3" json:"path_list_flags,omitempty"`
	// Number of references to the LDI
	LdiReferanceCount uint32 `protobuf:"varint,59,opt,name=ldi_referance_count,json=ldiReferanceCount,proto3" json:"ldi_referance_count,omitempty"`
	// The LDI flags
	LdiFlags uint32 `protobuf:"varint,60,opt,name=ldi_flags,json=ldiFlags,proto3" json:"ldi_flags,omitempty"`
	// The LDI type
	LdiType uint32 `protobuf:"varint,61,opt,name=ldi_type,json=ldiType,proto3" json:"ldi_type,omitempty"`
	// The pointer to the LDI
	LdiPointer uint32 `protobuf:"varint,62,opt,name=ldi_pointer,json=ldiPointer,proto3" json:"ldi_pointer,omitempty"`
	// The LW-LDI type
	LwLdiType uint32 `protobuf:"varint,63,opt,name=lw_ldi_type,json=lwLdiType,proto3" json:"lw_ldi_type,omitempty"`
	// The pointer to the LW-LDI
	LwLdiPointer uint32 `protobuf:"varint,64,opt,name=lw_ldi_pointer,json=lwLdiPointer,proto3" json:"lw_ldi_pointer,omitempty"`
	// The LW-LDI refcounter
	LwLdiRefernaceCount uint32 `protobuf:"varint,65,opt,name=lw_ldi_refernace_count,json=lwLdiRefernaceCount,proto3" json:"lw_ldi_refernace_count,omitempty"`
	// The pointer to the shared LDI in LW-LDI
	LwSharedLdiPointer uint32 `protobuf:"varint,66,opt,name=lw_shared_ldi_pointer,json=lwSharedLdiPointer,proto3" json:"lw_shared_ldi_pointer,omitempty"`
	// The LSPA flags
	LspaFlags uint32 `protobuf:"varint,67,opt,name=lspa_flags,json=lspaFlags,proto3" json:"lspa_flags,omitempty"`
	// The AFI table ID
	AfiTableId uint32 `protobuf:"varint,68,opt,name=afi_table_id,json=afiTableId,proto3" json:"afi_table_id,omitempty"`
	// The unicast or multicast label
	MulticastLabel bool `protobuf:"varint,69,opt,name=multicast_label,json=multicastLabel,proto3" json:"multicast_label,omitempty"`
	// The multicast info
	MulticastInformation *MplsMcastInfo `protobuf:"bytes,70,opt,name=multicast_information,json=multicastInformation,proto3" json:"multicast_information,omitempty"`
	// The time of last update in msec
	LeafTimeInMilliSeconds uint64   `protobuf:"varint,71,opt,name=leaf_time_in_milli_seconds,json=leafTimeInMilliSeconds,proto3" json:"leaf_time_in_milli_seconds,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *MplsShLeaf) Reset()         { *m = MplsShLeaf{} }
func (m *MplsShLeaf) String() string { return proto.CompactTextString(m) }
func (*MplsShLeaf) ProtoMessage()    {}
func (*MplsShLeaf) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc, []int{1}
}
func (m *MplsShLeaf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsShLeaf.Unmarshal(m, b)
}
func (m *MplsShLeaf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsShLeaf.Marshal(b, m, deterministic)
}
func (dst *MplsShLeaf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsShLeaf.Merge(dst, src)
}
func (m *MplsShLeaf) XXX_Size() int {
	return xxx_messageInfo_MplsShLeaf.Size(m)
}
func (m *MplsShLeaf) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsShLeaf.DiscardUnknown(m)
}

var xxx_messageInfo_MplsShLeaf proto.InternalMessageInfo

func (m *MplsShLeaf) GetLeafLocalLabel() uint32 {
	if m != nil {
		return m.LeafLocalLabel
	}
	return 0
}

func (m *MplsShLeaf) GetEosBit() uint32 {
	if m != nil {
		return m.EosBit
	}
	return 0
}

func (m *MplsShLeaf) GetLabelInformation() []*MplsAdjInfo {
	if m != nil {
		return m.LabelInformation
	}
	return nil
}

func (m *MplsShLeaf) GetLdiInformation() *MplsLdiInfo {
	if m != nil {
		return m.LdiInformation
	}
	return nil
}

func (m *MplsShLeaf) GetHardwareInformation() []byte {
	if m != nil {
		return m.HardwareInformation
	}
	return nil
}

func (m *MplsShLeaf) GetLeafReferanceCount() uint32 {
	if m != nil {
		return m.LeafReferanceCount
	}
	return 0
}

func (m *MplsShLeaf) GetLeafFlags() uint32 {
	if m != nil {
		return m.LeafFlags
	}
	return 0
}

func (m *MplsShLeaf) GetPathListReferanceCount() uint32 {
	if m != nil {
		return m.PathListReferanceCount
	}
	return 0
}

func (m *MplsShLeaf) GetPathListFlags() uint32 {
	if m != nil {
		return m.PathListFlags
	}
	return 0
}

func (m *MplsShLeaf) GetLdiReferanceCount() uint32 {
	if m != nil {
		return m.LdiReferanceCount
	}
	return 0
}

func (m *MplsShLeaf) GetLdiFlags() uint32 {
	if m != nil {
		return m.LdiFlags
	}
	return 0
}

func (m *MplsShLeaf) GetLdiType() uint32 {
	if m != nil {
		return m.LdiType
	}
	return 0
}

func (m *MplsShLeaf) GetLdiPointer() uint32 {
	if m != nil {
		return m.LdiPointer
	}
	return 0
}

func (m *MplsShLeaf) GetLwLdiType() uint32 {
	if m != nil {
		return m.LwLdiType
	}
	return 0
}

func (m *MplsShLeaf) GetLwLdiPointer() uint32 {
	if m != nil {
		return m.LwLdiPointer
	}
	return 0
}

func (m *MplsShLeaf) GetLwLdiRefernaceCount() uint32 {
	if m != nil {
		return m.LwLdiRefernaceCount
	}
	return 0
}

func (m *MplsShLeaf) GetLwSharedLdiPointer() uint32 {
	if m != nil {
		return m.LwSharedLdiPointer
	}
	return 0
}

func (m *MplsShLeaf) GetLspaFlags() uint32 {
	if m != nil {
		return m.LspaFlags
	}
	return 0
}

func (m *MplsShLeaf) GetAfiTableId() uint32 {
	if m != nil {
		return m.AfiTableId
	}
	return 0
}

func (m *MplsShLeaf) GetMulticastLabel() bool {
	if m != nil {
		return m.MulticastLabel
	}
	return false
}

func (m *MplsShLeaf) GetMulticastInformation() *MplsMcastInfo {
	if m != nil {
		return m.MulticastInformation
	}
	return nil
}

func (m *MplsShLeaf) GetLeafTimeInMilliSeconds() uint64 {
	if m != nil {
		return m.LeafTimeInMilliSeconds
	}
	return 0
}

type MplsFwdInfo struct {
	// L3 MTU
	L3Mtu uint32 `protobuf:"varint,1,opt,name=l3_mtu,json=l3Mtu,proto3" json:"l3_mtu,omitempty"`
	// Total encapsulation size: L2 + MPLS
	TotalEncapsulationSize uint32 `protobuf:"varint,2,opt,name=total_encapsulation_size,json=totalEncapsulationSize,proto3" json:"total_encapsulation_size,omitempty"`
	// Length of L2 encapsulation
	MacSize uint32 `protobuf:"varint,3,opt,name=mac_size,json=macSize,proto3" json:"mac_size,omitempty"`
	// Label stack
	LabelStack []uint32 `protobuf:"varint,4,rep,packed,name=label_stack,json=labelStack,proto3" json:"label_stack,omitempty"`
	// Number of packets switched
	TransmitNumberOfPacketsSwitched uint64 `protobuf:"varint,5,opt,name=transmit_number_of_packets_switched,json=transmitNumberOfPacketsSwitched,proto3" json:"transmit_number_of_packets_switched,omitempty"`
	// Number of Bytes switched
	TransmitNumberOfBytesSwitched uint64 `protobuf:"varint,6,opt,name=transmit_number_of_bytes_switched,json=transmitNumberOfBytesSwitched,proto3" json:"transmit_number_of_bytes_switched,omitempty"`
	// Status
	Status int32 `protobuf:"zigzag32,7,opt,name=status,proto3" json:"status,omitempty"`
	// Next hop interface
	NextHopInterface string `protobuf:"bytes,8,opt,name=next_hop_interface,json=nextHopInterface,proto3" json:"next_hop_interface,omitempty"`
	// The address family (V4/V6)
	NextHopProtocol string `protobuf:"bytes,9,opt,name=next_hop_protocol,json=nextHopProtocol,proto3" json:"next_hop_protocol,omitempty"`
	// Next hop address in string format
	NextHopString        string   `protobuf:"bytes,10,opt,name=next_hop_string,json=nextHopString,proto3" json:"next_hop_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsFwdInfo) Reset()         { *m = MplsFwdInfo{} }
func (m *MplsFwdInfo) String() string { return proto.CompactTextString(m) }
func (*MplsFwdInfo) ProtoMessage()    {}
func (*MplsFwdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc, []int{2}
}
func (m *MplsFwdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsFwdInfo.Unmarshal(m, b)
}
func (m *MplsFwdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsFwdInfo.Marshal(b, m, deterministic)
}
func (dst *MplsFwdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsFwdInfo.Merge(dst, src)
}
func (m *MplsFwdInfo) XXX_Size() int {
	return xxx_messageInfo_MplsFwdInfo.Size(m)
}
func (m *MplsFwdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsFwdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsFwdInfo proto.InternalMessageInfo

func (m *MplsFwdInfo) GetL3Mtu() uint32 {
	if m != nil {
		return m.L3Mtu
	}
	return 0
}

func (m *MplsFwdInfo) GetTotalEncapsulationSize() uint32 {
	if m != nil {
		return m.TotalEncapsulationSize
	}
	return 0
}

func (m *MplsFwdInfo) GetMacSize() uint32 {
	if m != nil {
		return m.MacSize
	}
	return 0
}

func (m *MplsFwdInfo) GetLabelStack() []uint32 {
	if m != nil {
		return m.LabelStack
	}
	return nil
}

func (m *MplsFwdInfo) GetTransmitNumberOfPacketsSwitched() uint64 {
	if m != nil {
		return m.TransmitNumberOfPacketsSwitched
	}
	return 0
}

func (m *MplsFwdInfo) GetTransmitNumberOfBytesSwitched() uint64 {
	if m != nil {
		return m.TransmitNumberOfBytesSwitched
	}
	return 0
}

func (m *MplsFwdInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MplsFwdInfo) GetNextHopInterface() string {
	if m != nil {
		return m.NextHopInterface
	}
	return ""
}

func (m *MplsFwdInfo) GetNextHopProtocol() string {
	if m != nil {
		return m.NextHopProtocol
	}
	return ""
}

func (m *MplsFwdInfo) GetNextHopString() string {
	if m != nil {
		return m.NextHopString
	}
	return ""
}

type MplsAdjInfo struct {
	// Label-Info type
	LabelInformationType uint32 `protobuf:"varint,1,opt,name=label_information_type,json=labelInformationType,proto3" json:"label_information_type,omitempty"`
	// Local label
	LocalLabel uint32 `protobuf:"varint,2,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	// Outgoing label
	OutgoingLabel uint32 `protobuf:"varint,3,opt,name=outgoing_label,json=outgoingLabel,proto3" json:"outgoing_label,omitempty"`
	// MPLS Adjacency flags
	MplsAdjacencyFlags uint32 `protobuf:"varint,4,opt,name=mpls_adjacency_flags,json=mplsAdjacencyFlags,proto3" json:"mpls_adjacency_flags,omitempty"`
	// Tunnel id present?
	TunnelIdPresent bool `protobuf:"varint,5,opt,name=tunnel_id_present,json=tunnelIdPresent,proto3" json:"tunnel_id_present,omitempty"`
	// Outgoing interface
	OutgoingInterface string `protobuf:"bytes,6,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	// Outgoing Physical Interface
	OutgoingPhysicalInterface string `protobuf:"bytes,7,opt,name=outgoing_physical_interface,json=outgoingPhysicalInterface,proto3" json:"outgoing_physical_interface,omitempty"`
	// Tunnel Interface
	TunnelInterface string `protobuf:"bytes,8,opt,name=tunnel_interface,json=tunnelInterface,proto3" json:"tunnel_interface,omitempty"`
	// Detail label info
	LabelInformationDetail    *MplsFwdInfo `protobuf:"bytes,9,opt,name=label_information_detail,json=labelInformationDetail,proto3" json:"label_information_detail,omitempty"`
	LabelInformationPathIndex uint32       `protobuf:"varint,10,opt,name=label_information_path_index,json=labelInformationPathIndex,proto3" json:"label_information_path_index,omitempty"`
	// NHinfo Type
	LabelInformationNextHopType string `protobuf:"bytes,11,opt,name=label_information_next_hop_type,json=labelInformationNextHopType,proto3" json:"label_information_next_hop_type,omitempty"`
	// The address family (v4/v6)
	LabelInformationNextHopProtocol string `protobuf:"bytes,12,opt,name=label_information_next_hop_protocol,json=labelInformationNextHopProtocol,proto3" json:"label_information_next_hop_protocol,omitempty"`
	// Bytes transmitted per LSP
	TxBytes uint64 `protobuf:"varint,13,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
	// Packets transmitted per LSP
	TxPackets uint64 `protobuf:"varint,14,opt,name=tx_packets,json=txPackets,proto3" json:"tx_packets,omitempty"`
	// Output Interface in string format
	OutgoingInterfaceString string `protobuf:"bytes,15,opt,name=outgoing_interface_string,json=outgoingInterfaceString,proto3" json:"outgoing_interface_string,omitempty"`
	// Output Label in string format
	OutgoingLabelString string `protobuf:"bytes,16,opt,name=outgoing_label_string,json=outgoingLabelString,proto3" json:"outgoing_label_string,omitempty"`
	// Prefix Or ID
	PrefixOrId string `protobuf:"bytes,17,opt,name=prefix_or_id,json=prefixOrId,proto3" json:"prefix_or_id,omitempty"`
	// Next hop address in string format
	LabelInformationNextHopString string `protobuf:"bytes,18,opt,name=label_information_next_hop_string,json=labelInformationNextHopString,proto3" json:"label_information_next_hop_string,omitempty"`
	// The version of the route
	LabelInformationRouteVersion uint64 `protobuf:"varint,19,opt,name=label_information_route_version,json=labelInformationRouteVersion,proto3" json:"label_information_route_version,omitempty"`
	// The time of last update in msec
	LabelInformationTimeInMilliSeconds uint64   `protobuf:"varint,20,opt,name=label_information_time_in_milli_seconds,json=labelInformationTimeInMilliSeconds,proto3" json:"label_information_time_in_milli_seconds,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *MplsAdjInfo) Reset()         { *m = MplsAdjInfo{} }
func (m *MplsAdjInfo) String() string { return proto.CompactTextString(m) }
func (*MplsAdjInfo) ProtoMessage()    {}
func (*MplsAdjInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc, []int{3}
}
func (m *MplsAdjInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsAdjInfo.Unmarshal(m, b)
}
func (m *MplsAdjInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsAdjInfo.Marshal(b, m, deterministic)
}
func (dst *MplsAdjInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsAdjInfo.Merge(dst, src)
}
func (m *MplsAdjInfo) XXX_Size() int {
	return xxx_messageInfo_MplsAdjInfo.Size(m)
}
func (m *MplsAdjInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsAdjInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsAdjInfo proto.InternalMessageInfo

func (m *MplsAdjInfo) GetLabelInformationType() uint32 {
	if m != nil {
		return m.LabelInformationType
	}
	return 0
}

func (m *MplsAdjInfo) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *MplsAdjInfo) GetOutgoingLabel() uint32 {
	if m != nil {
		return m.OutgoingLabel
	}
	return 0
}

func (m *MplsAdjInfo) GetMplsAdjacencyFlags() uint32 {
	if m != nil {
		return m.MplsAdjacencyFlags
	}
	return 0
}

func (m *MplsAdjInfo) GetTunnelIdPresent() bool {
	if m != nil {
		return m.TunnelIdPresent
	}
	return false
}

func (m *MplsAdjInfo) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *MplsAdjInfo) GetOutgoingPhysicalInterface() string {
	if m != nil {
		return m.OutgoingPhysicalInterface
	}
	return ""
}

func (m *MplsAdjInfo) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationDetail() *MplsFwdInfo {
	if m != nil {
		return m.LabelInformationDetail
	}
	return nil
}

func (m *MplsAdjInfo) GetLabelInformationPathIndex() uint32 {
	if m != nil {
		return m.LabelInformationPathIndex
	}
	return 0
}

func (m *MplsAdjInfo) GetLabelInformationNextHopType() string {
	if m != nil {
		return m.LabelInformationNextHopType
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationNextHopProtocol() string {
	if m != nil {
		return m.LabelInformationNextHopProtocol
	}
	return ""
}

func (m *MplsAdjInfo) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *MplsAdjInfo) GetTxPackets() uint64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *MplsAdjInfo) GetOutgoingInterfaceString() string {
	if m != nil {
		return m.OutgoingInterfaceString
	}
	return ""
}

func (m *MplsAdjInfo) GetOutgoingLabelString() string {
	if m != nil {
		return m.OutgoingLabelString
	}
	return ""
}

func (m *MplsAdjInfo) GetPrefixOrId() string {
	if m != nil {
		return m.PrefixOrId
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationNextHopString() string {
	if m != nil {
		return m.LabelInformationNextHopString
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationRouteVersion() uint64 {
	if m != nil {
		return m.LabelInformationRouteVersion
	}
	return 0
}

func (m *MplsAdjInfo) GetLabelInformationTimeInMilliSeconds() uint64 {
	if m != nil {
		return m.LabelInformationTimeInMilliSeconds
	}
	return 0
}

// Detailed load sharing information for mpls table entries
type MplsLdiInfo struct {
	// Hardware info
	LdiHardwareInformation []byte   `protobuf:"bytes,1,opt,name=ldi_hardware_information,json=ldiHardwareInformation,proto3" json:"ldi_hardware_information,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *MplsLdiInfo) Reset()         { *m = MplsLdiInfo{} }
func (m *MplsLdiInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLdiInfo) ProtoMessage()    {}
func (*MplsLdiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc, []int{4}
}
func (m *MplsLdiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLdiInfo.Unmarshal(m, b)
}
func (m *MplsLdiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLdiInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLdiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLdiInfo.Merge(dst, src)
}
func (m *MplsLdiInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLdiInfo.Size(m)
}
func (m *MplsLdiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLdiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLdiInfo proto.InternalMessageInfo

func (m *MplsLdiInfo) GetLdiHardwareInformation() []byte {
	if m != nil {
		return m.LdiHardwareInformation
	}
	return nil
}

// Information for mpls multicast entries
type MplsMcastInfo struct {
	// MOL base flags
	MulticastMolBaseFlags uint32 `protobuf:"varint,1,opt,name=multicast_mol_base_flags,json=multicastMolBaseFlags,proto3" json:"multicast_mol_base_flags,omitempty"`
	// MOL flags
	MulticastMolFlags uint32 `protobuf:"varint,2,opt,name=multicast_mol_flags,json=multicastMolFlags,proto3" json:"multicast_mol_flags,omitempty"`
	// MOL refcount
	MulticastMolReferanceCount uint32 `protobuf:"varint,3,opt,name=multicast_mol_referance_count,json=multicastMolReferanceCount,proto3" json:"multicast_mol_referance_count,omitempty"`
	// multicast mpls tunnel
	MulticastTunnelInterfaceHandler string `protobuf:"bytes,4,opt,name=multicast_tunnel_interface_handler,json=multicastTunnelInterfaceHandler,proto3" json:"multicast_tunnel_interface_handler,omitempty"`
	// multicast mpls P2MP-TE tunnel id or MLDP Tunnel LSMID on all nodes
	MulticastTunnelId uint32 `protobuf:"varint,5,opt,name=multicast_tunnel_id,json=multicastTunnelId,proto3" json:"multicast_tunnel_id,omitempty"`
	// multicast nhinfo for p2mp TE Head
	MulticastTunnelNextHopInformation uint32 `protobuf:"varint,6,opt,name=multicast_tunnel_next_hop_information,json=multicastTunnelNextHopInformation,proto3" json:"multicast_tunnel_next_hop_information,omitempty"`
	// multicast LSPVIF for MLDP Tunnels
	MulticastTunnelLspvif uint32 `protobuf:"varint,7,opt,name=multicast_tunnel_lspvif,json=multicastTunnelLspvif,proto3" json:"multicast_tunnel_lspvif,omitempty"`
	// num multicast mpls output paths
	MulticastMplsOutputPaths uint32 `protobuf:"varint,8,opt,name=multicast_mpls_output_paths,json=multicastMplsOutputPaths,proto3" json:"multicast_mpls_output_paths,omitempty"`
	// num multicast mpls prot output paths
	MulticastMplsProtocolOutputPaths uint32 `protobuf:"varint,9,opt,name=multicast_mpls_protocol_output_paths,json=multicastMplsProtocolOutputPaths,proto3" json:"multicast_mpls_protocol_output_paths,omitempty"`
	// num multicast mpls local output paths
	MulticastMplsLocalOutputPaths uint32 `protobuf:"varint,10,opt,name=multicast_mpls_local_output_paths,json=multicastMplsLocalOutputPaths,proto3" json:"multicast_mpls_local_output_paths,omitempty"`
	// The multicast RPF-ID
	MulticastRpfId uint32 `protobuf:"varint,11,opt,name=multicast_rpf_id,json=multicastRpfId,proto3" json:"multicast_rpf_id,omitempty"`
	// The multicast ENCAP-ID
	MulticastEncapId uint32 `protobuf:"varint,12,opt,name=multicast_encap_id,json=multicastEncapId,proto3" json:"multicast_encap_id,omitempty"`
	// The multicast platform data len
	MulticastPlatformDataLength uint32 `protobuf:"varint,13,opt,name=multicast_platform_data_length,json=multicastPlatformDataLength,proto3" json:"multicast_platform_data_length,omitempty"`
	// The multicast platform data
	MulticastPlatformData []byte   `protobuf:"bytes,14,opt,name=multicast_platform_data,json=multicastPlatformData,proto3" json:"multicast_platform_data,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *MplsMcastInfo) Reset()         { *m = MplsMcastInfo{} }
func (m *MplsMcastInfo) String() string { return proto.CompactTextString(m) }
func (*MplsMcastInfo) ProtoMessage()    {}
func (*MplsMcastInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc, []int{5}
}
func (m *MplsMcastInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsMcastInfo.Unmarshal(m, b)
}
func (m *MplsMcastInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsMcastInfo.Marshal(b, m, deterministic)
}
func (dst *MplsMcastInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsMcastInfo.Merge(dst, src)
}
func (m *MplsMcastInfo) XXX_Size() int {
	return xxx_messageInfo_MplsMcastInfo.Size(m)
}
func (m *MplsMcastInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsMcastInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsMcastInfo proto.InternalMessageInfo

func (m *MplsMcastInfo) GetMulticastMolBaseFlags() uint32 {
	if m != nil {
		return m.MulticastMolBaseFlags
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMolFlags() uint32 {
	if m != nil {
		return m.MulticastMolFlags
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMolReferanceCount() uint32 {
	if m != nil {
		return m.MulticastMolReferanceCount
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastTunnelInterfaceHandler() string {
	if m != nil {
		return m.MulticastTunnelInterfaceHandler
	}
	return ""
}

func (m *MplsMcastInfo) GetMulticastTunnelId() uint32 {
	if m != nil {
		return m.MulticastTunnelId
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastTunnelNextHopInformation() uint32 {
	if m != nil {
		return m.MulticastTunnelNextHopInformation
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastTunnelLspvif() uint32 {
	if m != nil {
		return m.MulticastTunnelLspvif
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMplsOutputPaths() uint32 {
	if m != nil {
		return m.MulticastMplsOutputPaths
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMplsProtocolOutputPaths() uint32 {
	if m != nil {
		return m.MulticastMplsProtocolOutputPaths
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMplsLocalOutputPaths() uint32 {
	if m != nil {
		return m.MulticastMplsLocalOutputPaths
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastRpfId() uint32 {
	if m != nil {
		return m.MulticastRpfId
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastEncapId() uint32 {
	if m != nil {
		return m.MulticastEncapId
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastPlatformDataLength() uint32 {
	if m != nil {
		return m.MulticastPlatformDataLength
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastPlatformData() []byte {
	if m != nil {
		return m.MulticastPlatformData
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsShLeaf_KEYS)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.informations.information.mpls_sh_leaf_KEYS")
	proto.RegisterType((*MplsShLeaf)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.informations.information.mpls_sh_leaf")
	proto.RegisterType((*MplsFwdInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.informations.information.mpls_fwd_info")
	proto.RegisterType((*MplsAdjInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.informations.information.mpls_adj_info")
	proto.RegisterType((*MplsLdiInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.informations.information.mpls_ldi_info")
	proto.RegisterType((*MplsMcastInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.informations.information.mpls_mcast_info")
}

func init() { proto.RegisterFile("mpls_sh_leaf.proto", fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc) }

var fileDescriptor_mpls_sh_leaf_dae78e05eeb3c0bc = []byte{
	// 1561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x5d, 0x72, 0x1b, 0xc7,
	0x11, 0x2e, 0x84, 0x32, 0x48, 0x34, 0x01, 0x92, 0x18, 0x92, 0xd0, 0xd2, 0x34, 0x4d, 0x88, 0xb6,
	0x63, 0xc4, 0x95, 0xb0, 0x6c, 0xd1, 0xf1, 0x8f, 0x12, 0xdb, 0x11, 0x45, 0x3a, 0x44, 0x19, 0xa4,
	0x50, 0x0b, 0x96, 0xab, 0x92, 0x97, 0xa9, 0xc1, 0xee, 0x80, 0x98, 0x68, 0xf6, 0xa7, 0x76, 0x06,
	0x02, 0xe8, 0x63, 0xe4, 0x21, 0xcf, 0xa9, 0xca, 0x0d, 0x72, 0x90, 0x1c, 0x23, 0x07, 0xc8, 0x09,
	0x52, 0xd3, 0xb3, 0x7f, 0x58, 0xd0, 0x7e, 0x93, 0x5e, 0x54, 0x50, 0xf7, 0xd7, 0xdf, 0xf4, 0x76,
	0x4f, 0x7f, 0x3d, 0x04, 0x12, 0xc4, 0x52, 0x51, 0x35, 0xa5, 0x92, 0xb3, 0xc9, 0x69, 0x9c, 0x44,
	0x3a, 0x22, 0x7f, 0xf5, 0x84, 0xf2, 0x22, 0x2a, 0x22, 0x45, 0x17, 0x09, 0x9d, 0x88, 0x31, 0xf5,
	0xa2, 0x20, 0x88, 0x42, 0x1a, 0xc5, 0x3c, 0x39, 0xc5, 0x80, 0x49, 0x94, 0xcc, 0x59, 0xe2, 0x8b,
	0xf0, 0xee, 0x34, 0x8c, 0x7c, 0xae, 0xf0, 0xdf, 0x53, 0xc9, 0xc6, 0x5c, 0x9a, 0x80, 0x53, 0x11,
	0x4e, 0xa2, 0x24, 0x60, 0x5a, 0x44, 0xa1, 0x2a, 0xff, 0xe7, 0xc4, 0x83, 0x76, 0xf9, 0x44, 0xfa,
	0xc3, 0xe5, 0x5f, 0x46, 0xe4, 0x10, 0x1a, 0x26, 0x9e, 0x86, 0x2c, 0xe0, 0x4e, 0xad, 0x5b, 0xeb,
	0x35, 0xdc, 0x0d, 0x63, 0xb8, 0x61, 0x01, 0x27, 0xc7, 0xb0, 0x69, 0x79, 0x5f, 0x33, 0x39, 0xe3,
	0xce, 0xaf, 0xba, 0xb5, 0x5e, 0xcb, 0x05, 0x34, 0xfd, 0x68, 0x2c, 0x64, 0x07, 0xd6, 0x78, 0xa4,
	0x9c, 0x35, 0x8c, 0x33, 0x3f, 0x4f, 0xfe, 0xdb, 0x80, 0x66, 0xf9, 0x14, 0xd2, 0x83, 0x1d, 0x3c,
	0x4d, 0x46, 0x1e, 0x93, 0x14, 0x63, 0x9d, 0xa7, 0x48, 0xb4, 0x65, 0xec, 0x03, 0x63, 0x1e, 0x18,
	0x2b, 0x79, 0x0c, 0xeb, 0x3c, 0x52, 0x74, 0x2c, 0xb4, 0x73, 0x86, 0x80, 0x3a, 0x8f, 0xd4, 0xb9,
	0xd0, 0xe4, 0x1f, 0x35, 0x68, 0xdb, 0x3c, 0x4a, 0x9f, 0xe3, 0x7c, 0xde, 0x5d, 0xeb, 0x6d, 0x3e,
	0x15, 0xa7, 0x6f, 0xae, 0x62, 0x36, 0x9a, 0xf9, 0x7f, 0xc3, 0x73, 0xdd, 0x1d, 0x8c, 0xe8, 0x17,
	0x7e, 0xf2, 0xf7, 0x1a, 0x6c, 0x4b, 0x5f, 0x2c, 0xa5, 0xf5, 0xfb, 0x6e, 0xed, 0xad, 0xa4, 0x95,
	0x9d, 0xeb, 0x6e, 0x49, 0x5f, 0x94, 0x93, 0xfa, 0x0c, 0xf6, 0xa6, 0x2c, 0xf1, 0xe7, 0x2c, 0xe1,
	0x4b, 0x89, 0x7d, 0xd1, 0xad, 0xf5, 0x9a, 0xee, 0x6e, 0xe6, 0x2b, 0x87, 0x7c, 0x0a, 0x7b, 0xd8,
	0xa3, 0x84, 0x4f, 0x78, 0xc2, 0x42, 0x8f, 0x53, 0x2f, 0x9a, 0x85, 0xda, 0xf9, 0x12, 0xdb, 0x40,
	0x8c, 0xcf, 0xcd, 0x5c, 0x2f, 0x8c, 0x87, 0x1c, 0x01, 0x60, 0xc4, 0x44, 0xb2, 0x3b, 0xe5, 0x7c,
	0x85, 0xb8, 0x86, 0xb1, 0x7c, 0x6f, 0x0c, 0xe4, 0x6b, 0x38, 0x88, 0x99, 0x9e, 0x52, 0x29, 0x94,
	0x5e, 0x61, 0xfd, 0x1a, 0xd1, 0x1d, 0x03, 0x18, 0x08, 0xa5, 0x2b, 0xcc, 0xbf, 0x86, 0xed, 0x22,
	0xd4, 0xd2, 0x3f, 0xc3, 0x80, 0x56, 0x16, 0x60, 0x8f, 0x38, 0x85, 0x5d, 0x53, 0x82, 0x2a, 0xf9,
	0x1f, 0x10, 0xdb, 0x96, 0xbe, 0xa8, 0xf0, 0x1e, 0x42, 0xc3, 0xe0, 0x2d, 0xe3, 0x1f, 0x11, 0xb5,
	0x21, 0x7d, 0x61, 0xc9, 0x0e, 0xc0, 0xfc, 0xa6, 0xfa, 0x3e, 0xe6, 0xce, 0x37, 0xe8, 0x5b, 0x97,
	0xbe, 0xb8, 0xbd, 0x8f, 0xed, 0x0c, 0xf8, 0x82, 0xc6, 0x91, 0x08, 0x35, 0x4f, 0x9c, 0x6f, 0xd3,
	0x19, 0xf0, 0xc5, 0xd0, 0x5a, 0xc8, 0xfb, 0xb0, 0x29, 0xe7, 0x34, 0x0f, 0xff, 0x2e, 0xad, 0xc5,
	0x7c, 0x90, 0x12, 0x7c, 0x08, 0x5b, 0xa9, 0x3f, 0xe3, 0xf8, 0x13, 0x42, 0x9a, 0x08, 0xc9, 0x58,
	0xce, 0xa0, 0x93, 0xa2, 0xf0, 0x8b, 0x42, 0x96, 0x7f, 0xd1, 0x73, 0x44, 0xef, 0x22, 0xda, 0xcd,
	0x7c, 0xf6, 0x9b, 0x3e, 0x83, 0x7d, 0x39, 0xa7, 0x6a, 0xca, 0x12, 0xee, 0x2f, 0x9d, 0x70, 0x9e,
	0x36, 0x6e, 0x3e, 0x42, 0x5f, 0xe9, 0x1c, 0xd3, 0x38, 0x15, 0xb3, 0xb4, 0x0e, 0x2f, 0xd2, 0x64,
	0x55, 0xcc, 0x6c, 0x21, 0xba, 0xd0, 0x64, 0x13, 0x41, 0x35, 0x1b, 0x4b, 0x4e, 0x85, 0xef, 0x5c,
	0xd8, 0xcf, 0x65, 0x13, 0x71, 0x6b, 0x4c, 0x7d, 0x9f, 0x7c, 0x0c, 0xdb, 0xc1, 0x4c, 0x6a, 0xe1,
	0x31, 0xa5, 0xd3, 0x71, 0xbe, 0xec, 0xd6, 0x7a, 0x1b, 0xee, 0x56, 0x6e, 0xb6, 0xe3, 0xfc, 0xcf,
	0x1a, 0xec, 0x17, 0xc8, 0xf2, 0x4d, 0xfc, 0x1e, 0x47, 0xe4, 0xd5, 0x1b, 0x1f, 0x91, 0x20, 0x3f,
	0xd9, 0xdd, 0xcb, 0x33, 0x29, 0xdf, 0xfb, 0x67, 0xf0, 0x2e, 0xde, 0x62, 0x2d, 0x02, 0x33, 0x2b,
	0x34, 0x10, 0x52, 0x0a, 0xaa, 0xb8, 0x17, 0x85, 0xbe, 0x72, 0xfe, 0xdc, 0xad, 0xf5, 0x1e, 0xb9,
	0x1d, 0x83, 0xb8, 0x15, 0x01, 0xef, 0x87, 0xd7, 0xc6, 0x3d, 0xb2, 0xde, 0x93, 0xff, 0xac, 0x41,
	0xcb, 0xe6, 0x38, 0xf7, 0xf1, 0x0c, 0xb2, 0x0f, 0x75, 0x79, 0x46, 0x03, 0x3d, 0x43, 0x1d, 0x6d,
	0xb9, 0xef, 0xc8, 0xb3, 0x6b, 0x3d, 0x23, 0x5f, 0x81, 0xa3, 0x23, 0xcd, 0x24, 0xe5, 0xa1, 0xc7,
	0x62, 0x35, 0x93, 0x78, 0x36, 0x55, 0xe2, 0xa7, 0x4c, 0x51, 0x3b, 0xe8, 0xbf, 0x2c, 0xbb, 0x47,
	0xe2, 0x27, 0x6e, 0x6e, 0x65, 0xc0, 0x3c, 0x8b, 0x5c, 0xb3, 0xb7, 0x32, 0x60, 0x1e, 0xba, 0x72,
	0x65, 0x56, 0x9a, 0x79, 0xaf, 0x9c, 0x47, 0xdd, 0xb5, 0x5c, 0x99, 0x47, 0xc6, 0x42, 0x06, 0xf0,
	0x81, 0x4e, 0x58, 0xa8, 0x02, 0xa1, 0x69, 0x38, 0x0b, 0xc6, 0x3c, 0xa1, 0xd1, 0x84, 0xc6, 0xcc,
	0x7b, 0xc5, 0xb5, 0xa2, 0x6a, 0x2e, 0xb4, 0x37, 0xe5, 0xbe, 0xf3, 0x0e, 0x7e, 0xe3, 0x71, 0x06,
	0xbd, 0x41, 0xe4, 0xcb, 0xc9, 0xd0, 0xe2, 0x46, 0x29, 0x8c, 0x5c, 0xc1, 0x93, 0x07, 0xd8, 0xc6,
	0xf7, 0x9a, 0x97, 0xb8, 0xea, 0xc8, 0x75, 0x54, 0xe5, 0x3a, 0x37, 0xa8, 0x9c, 0xa9, 0x03, 0x75,
	0xa5, 0x99, 0x9e, 0x29, 0x67, 0xbd, 0x5b, 0xeb, 0xb5, 0xdd, 0xf4, 0x7f, 0xe4, 0xb7, 0x40, 0x42,
	0xbe, 0xd0, 0x74, 0x1a, 0xc5, 0x14, 0x6f, 0xea, 0x84, 0x79, 0xdc, 0xd9, 0xc0, 0xc5, 0xb2, 0x63,
	0x3c, 0x57, 0x51, 0xdc, 0xcf, 0xec, 0xe4, 0x13, 0x68, 0xe7, 0x68, 0x5c, 0x9c, 0x5e, 0x24, 0x9d,
	0x06, 0x82, 0xb7, 0x53, 0xf0, 0x30, 0x35, 0x1b, 0x41, 0xc9, 0xb1, 0x4a, 0x27, 0x22, 0xbc, 0x73,
	0x00, 0x91, 0xad, 0x14, 0x39, 0x42, 0xe3, 0xc9, 0xbf, 0x1b, 0x69, 0x43, 0x33, 0xc1, 0x27, 0x9f,
	0x43, 0x67, 0x65, 0xed, 0xd8, 0x21, 0xb7, 0x0d, 0xde, 0xab, 0x2e, 0x84, 0x5c, 0x30, 0x4a, 0xbb,
	0x2e, 0x5b, 0x9a, 0xc5, 0x9e, 0xfb, 0x08, 0xb6, 0xa2, 0x99, 0xbe, 0x8b, 0x44, 0x78, 0x97, 0x62,
	0x6c, 0x73, 0x5b, 0x99, 0xd5, 0xc2, 0x3e, 0x85, 0xbd, 0x2c, 0x1d, 0xe6, 0xf1, 0xd0, 0xbb, 0x4f,
	0x67, 0xf6, 0x91, 0x9d, 0x6d, 0xe3, 0x7b, 0x9e, 0xb9, 0xec, 0xf0, 0x7e, 0x02, 0x6d, 0x3d, 0x0b,
	0x43, 0x93, 0xb0, 0x4f, 0xe3, 0x84, 0x2b, 0x1e, 0x6a, 0xec, 0xf0, 0x86, 0xbb, 0x6d, 0x1d, 0x7d,
	0x7f, 0x68, 0xcd, 0xe4, 0x77, 0x40, 0xf2, 0x24, 0x8a, 0x7a, 0xd7, 0xb1, 0x30, 0xed, 0xcc, 0x53,
	0x14, 0xfc, 0x5b, 0x38, 0xcc, 0xe1, 0xf1, 0xf4, 0x5e, 0x09, 0xf3, 0x81, 0x45, 0xdc, 0x3a, 0xc6,
	0x1d, 0x64, 0x90, 0x61, 0x8a, 0x28, 0xe2, 0x7f, 0x03, 0x3b, 0x59, 0x6a, 0x95, 0xe6, 0x66, 0x99,
	0xe5, 0xd0, 0x7f, 0xd5, 0xc0, 0x59, 0x2d, 0xbb, 0xcf, 0x35, 0x13, 0xb6, 0xc7, 0x6f, 0x63, 0xbb,
	0x66, 0x43, 0xed, 0x76, 0xaa, 0x3d, 0xbe, 0xc0, 0x44, 0xc8, 0x77, 0xf0, 0xde, 0x6a, 0x92, 0xb8,
	0xb8, 0x44, 0xe8, 0xf3, 0x05, 0x5e, 0xb1, 0x96, 0x7b, 0x50, 0x8d, 0x1e, 0x32, 0x3d, 0xed, 0x1b,
	0x00, 0xb9, 0x80, 0xe3, 0x55, 0x82, 0xfc, 0xa2, 0xe2, 0x2d, 0xdb, 0xc4, 0x02, 0x1d, 0x56, 0x39,
	0x6e, 0xec, 0xb5, 0xc5, 0xcb, 0x36, 0x80, 0x0f, 0x7e, 0x81, 0x25, 0x1f, 0x8d, 0x26, 0x32, 0x1d,
	0xff, 0x0c, 0x53, 0x3e, 0x2a, 0x07, 0xb0, 0xa1, 0x17, 0x76, 0xac, 0x9d, 0x16, 0x4e, 0xf3, 0xba,
	0x5e, 0xe0, 0xfc, 0x9a, 0xbd, 0xa1, 0x17, 0x99, 0x7e, 0x38, 0x5b, 0xe8, 0x6c, 0xe8, 0x45, 0x2a,
	0x14, 0xe4, 0x19, 0x1c, 0xac, 0x5e, 0xa7, 0x6c, 0xdc, 0xb6, 0xf1, 0xf4, 0xc7, 0x2b, 0xb7, 0xca,
	0x0e, 0x1e, 0x79, 0x0a, 0xfb, 0xcb, 0xf3, 0x90, 0xc5, 0xed, 0x60, 0xdc, 0xee, 0xd2, 0x58, 0xa4,
	0x31, 0x5d, 0x68, 0xc6, 0x09, 0x9f, 0x88, 0x05, 0x8d, 0x12, 0xb3, 0xa7, 0xda, 0x08, 0x05, 0x6b,
	0x7b, 0x99, 0xf4, 0x51, 0xb2, 0x7e, 0xa1, 0x32, 0xe9, 0x09, 0x04, 0xc3, 0x8e, 0x7e, 0xa6, 0x2e,
	0xe9, 0x59, 0x97, 0x0f, 0x75, 0x2a, 0x89, 0x66, 0x9a, 0xd3, 0xd7, 0x3c, 0x51, 0x66, 0xa3, 0xed,
	0x62, 0x3d, 0xde, 0xab, 0xf2, 0xb8, 0x06, 0xf4, 0xa3, 0xc5, 0x90, 0x11, 0x7c, 0xfc, 0x80, 0x9a,
	0x3c, 0xb8, 0x79, 0xf6, 0x90, 0xee, 0x64, 0x45, 0x5e, 0x56, 0xb7, 0x50, 0x3f, 0xd5, 0xac, 0xec,
	0x35, 0x68, 0xb6, 0x8d, 0xf9, 0xfd, 0xe0, 0x0b, 0xb0, 0x86, 0x2f, 0xc0, 0x8e, 0xf4, 0xc5, 0xd5,
	0xea, 0x23, 0xf0, 0xe4, 0x7f, 0x75, 0xd8, 0xae, 0xac, 0x4d, 0xf2, 0x25, 0x38, 0xc5, 0x0a, 0x0f,
	0x22, 0x49, 0xc7, 0x4c, 0xf1, 0x54, 0x87, 0xac, 0x06, 0x16, 0x2b, 0xfe, 0x3a, 0x92, 0xe7, 0x4c,
	0xf1, 0xfc, 0x75, 0xb6, 0x1c, 0x68, 0x63, 0xac, 0x18, 0xb6, 0xcb, 0x31, 0x16, 0xff, 0x1c, 0x8e,
	0x96, 0xf1, 0xd5, 0x77, 0x9d, 0x95, 0xc8, 0x77, 0xcb, 0x91, 0x95, 0x07, 0xde, 0x0f, 0x70, 0x52,
	0x50, 0x54, 0xc5, 0x86, 0x4e, 0x59, 0xe8, 0x4b, 0x9e, 0xa0, 0x7a, 0x36, 0xdc, 0xe3, 0x1c, 0x79,
	0xbb, 0xac, 0x3e, 0x57, 0x16, 0xb6, 0x9c, 0x7f, 0x2e, 0xaa, 0x28, 0xa6, 0xe5, 0xfc, 0xd3, 0x68,
	0x9f, 0x0c, 0xe1, 0xa3, 0x15, 0x7c, 0x69, 0x9f, 0x15, 0x3d, 0xa8, 0x23, 0xc3, 0x93, 0x0a, 0xc3,
	0x4d, 0xb6, 0xe0, 0x8a, 0xb7, 0xc9, 0x17, 0xf0, 0x78, 0x85, 0x51, 0xaa, 0xf8, 0xb5, 0x98, 0xa0,
	0xda, 0x96, 0x2b, 0x6f, 0x39, 0x06, 0xe8, 0x24, 0xdf, 0xc0, 0x61, 0xa9, 0x92, 0xa6, 0x9f, 0xd1,
	0x4c, 0xc7, 0x33, 0x8d, 0xe2, 0xa4, 0x50, 0x74, 0x5b, 0x6e, 0xd1, 0xd5, 0xeb, 0x58, 0xaa, 0x97,
	0x08, 0x30, 0xd2, 0xa4, 0xc8, 0x0d, 0x7c, 0x58, 0x09, 0xcf, 0x44, 0x64, 0x99, 0xa7, 0x81, 0x3c,
	0xdd, 0x25, 0x9e, 0x4c, 0x47, 0xca, 0x7c, 0x57, 0xf0, 0xa4, 0xc2, 0x67, 0x97, 0xe3, 0x12, 0x99,
	0x15, 0xcb, 0xa3, 0x25, 0x32, 0xfc, 0xc3, 0xb0, 0xcc, 0xd4, 0x83, 0x9d, 0x82, 0x29, 0x89, 0x27,
	0xa6, 0x1f, 0x9b, 0xf6, 0x0f, 0xc9, 0xdc, 0xee, 0xc6, 0x93, 0xbe, 0x6f, 0xde, 0x12, 0x05, 0x12,
	0x5f, 0x5d, 0x06, 0xdb, 0x44, 0x6c, 0xc1, 0x81, 0xef, 0xad, 0xbe, 0x4f, 0x5e, 0xc0, 0xfb, 0x05,
	0x3a, 0x96, 0x4c, 0x9b, 0x1e, 0x50, 0x9f, 0x69, 0x46, 0x25, 0x0f, 0xef, 0xf4, 0x14, 0xa5, 0xb0,
	0xe5, 0x16, 0x65, 0x1d, 0xa6, 0xa0, 0x0b, 0xa6, 0xd9, 0x00, 0x21, 0xcb, 0xdd, 0x5a, 0x22, 0x41,
	0xad, 0x6c, 0x96, 0xba, 0x55, 0x8e, 0x1e, 0xd7, 0xb1, 0xba, 0x67, 0xff, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0x12, 0xd6, 0x7d, 0x0c, 0x10, 0x00, 0x00,
}
