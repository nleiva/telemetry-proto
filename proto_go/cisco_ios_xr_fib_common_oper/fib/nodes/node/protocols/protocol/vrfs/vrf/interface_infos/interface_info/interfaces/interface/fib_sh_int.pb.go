// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fib_sh_int.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_vrfs_vrf_interface_infos_interface_info_interfaces_interface

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

// FIB per interface information
type FibShInt_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	VrfName              string   `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	LinkType             string   `protobuf:"bytes,4,opt,name=link_type,json=linkType,proto3" json:"link_type,omitempty"`
	InterfaceName        string   `protobuf:"bytes,5,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShInt_KEYS) Reset()         { *m = FibShInt_KEYS{} }
func (m *FibShInt_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibShInt_KEYS) ProtoMessage()    {}
func (*FibShInt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_int_1b0c7d47e5664c2f, []int{0}
}
func (m *FibShInt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShInt_KEYS.Unmarshal(m, b)
}
func (m *FibShInt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShInt_KEYS.Marshal(b, m, deterministic)
}
func (dst *FibShInt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShInt_KEYS.Merge(dst, src)
}
func (m *FibShInt_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibShInt_KEYS.Size(m)
}
func (m *FibShInt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShInt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibShInt_KEYS proto.InternalMessageInfo

func (m *FibShInt_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibShInt_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *FibShInt_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FibShInt_KEYS) GetLinkType() string {
	if m != nil {
		return m.LinkType
	}
	return ""
}

func (m *FibShInt_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type FibShInt struct {
	// Interface handle
	PerInterface string `protobuf:"bytes,50,opt,name=per_interface,json=perInterface,proto3" json:"per_interface,omitempty"`
	// FIB Interface type
	FibInterfaceType uint32 `protobuf:"varint,51,opt,name=fib_interface_type,json=fibInterfaceType,proto3" json:"fib_interface_type,omitempty"`
	// Pointer to fibidb
	FibIdPointer uint32 `protobuf:"varint,52,opt,name=fib_id_pointer,json=fibIdPointer,proto3" json:"fib_id_pointer,omitempty"`
	// Flags on fibidb
	FibIdFlags uint32 `protobuf:"varint,53,opt,name=fib_id_flags,json=fibIdFlags,proto3" json:"fib_id_flags,omitempty"`
	// Pointer to fibidb extension
	FibIdExtensionPointer uint32 `protobuf:"varint,54,opt,name=fib_id_extension_pointer,json=fibIdExtensionPointer,proto3" json:"fib_id_extension_pointer,omitempty"`
	// Flags on fibidb extension
	FibIdExtensionFlags uint32 `protobuf:"varint,55,opt,name=fib_id_extension_flags,json=fibIdExtensionFlags,proto3" json:"fib_id_extension_flags,omitempty"`
	// Number of dependent nhinfo's
	NumberOfDependentNextHopInformation uint32 `protobuf:"varint,56,opt,name=number_of_dependent_next_hop_information,json=numberOfDependentNextHopInformation,proto3" json:"number_of_dependent_next_hop_information,omitempty"`
	// Vrf local cef info ptr
	VrfLocalCefInformationPointer uint32 `protobuf:"varint,57,opt,name=vrf_local_cef_information_pointer,json=vrfLocalCefInformationPointer,proto3" json:"vrf_local_cef_information_pointer,omitempty"`
	// Reference count
	ReferenceCount uint32 `protobuf:"varint,58,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	// Time last modified
	LastModifiedTime uint32 `protobuf:"varint,59,opt,name=last_modified_time,json=lastModifiedTime,proto3" json:"last_modified_time,omitempty"`
	// Last Oper
	LastOperation string `protobuf:"bytes,60,opt,name=last_operation,json=lastOperation,proto3" json:"last_operation,omitempty"`
	// Is the protocol configured?
	ProtocolEnabled bool `protobuf:"varint,61,opt,name=protocol_enabled,json=protocolEnabled,proto3" json:"protocol_enabled,omitempty"`
	// Detailed FIB interface information
	DetailFibIntInformation *FibShIntDet `protobuf:"bytes,62,opt,name=detail_fib_int_information,json=detailFibIntInformation,proto3" json:"detail_fib_int_information,omitempty"`
	// Reference count for the protocol
	ReferanceCountForProtocol uint32 `protobuf:"varint,63,opt,name=referance_count_for_protocol,json=referanceCountForProtocol,proto3" json:"referance_count_for_protocol,omitempty"`
	// Number of input packets
	NumberOfInputPackets uint64 `protobuf:"varint,64,opt,name=number_of_input_packets,json=numberOfInputPackets,proto3" json:"number_of_input_packets,omitempty"`
	// Number of input bytes
	NumberOfInputBytes uint64 `protobuf:"varint,65,opt,name=number_of_input_bytes,json=numberOfInputBytes,proto3" json:"number_of_input_bytes,omitempty"`
	// Number of output packets
	NumberOfOutputPackets uint64 `protobuf:"varint,66,opt,name=number_of_output_packets,json=numberOfOutputPackets,proto3" json:"number_of_output_packets,omitempty"`
	// Number of output bytes
	NumberOfOutputBytes uint64 `protobuf:"varint,67,opt,name=number_of_output_bytes,json=numberOfOutputBytes,proto3" json:"number_of_output_bytes,omitempty"`
	// Interface up flag
	InterfaceUpFlag bool `protobuf:"varint,68,opt,name=interface_up_flag,json=interfaceUpFlag,proto3" json:"interface_up_flag,omitempty"`
	// Per packet loadbalancing flag
	PerPacketLoadBalancingFlag bool `protobuf:"varint,69,opt,name=per_packet_load_balancing_flag,json=perPacketLoadBalancingFlag,proto3" json:"per_packet_load_balancing_flag,omitempty"`
	// P2P interface flag
	P2PInterfaceFlag bool `protobuf:"varint,70,opt,name=p2_p_interface_flag,json=p2PInterfaceFlag,proto3" json:"p2_p_interface_flag,omitempty"`
	// Loopback interface flag
	LoopbackInterfaceFlag bool `protobuf:"varint,71,opt,name=loopback_interface_flag,json=loopbackInterfaceFlag,proto3" json:"loopback_interface_flag,omitempty"`
	// Null interface flag
	NullInterfaceFlag bool `protobuf:"varint,72,opt,name=null_interface_flag,json=nullInterfaceFlag,proto3" json:"null_interface_flag,omitempty"`
	// Tunnel interface flag
	TunnelInterfaceFlag bool `protobuf:"varint,73,opt,name=tunnel_interface_flag,json=tunnelInterfaceFlag,proto3" json:"tunnel_interface_flag,omitempty"`
	// GRE Tunnel interface flag
	GreTunnelInterfaceFlag bool `protobuf:"varint,74,opt,name=gre_tunnel_interface_flag,json=greTunnelInterfaceFlag,proto3" json:"gre_tunnel_interface_flag,omitempty"`
	// Punt packets from FIB switching flag
	PuntPacketsFromFibSwitchingFlag bool `protobuf:"varint,75,opt,name=punt_packets_from_fib_switching_flag,json=puntPacketsFromFibSwitchingFlag,proto3" json:"punt_packets_from_fib_switching_flag,omitempty"`
	// Drop packets while FIB switching flag
	DropPacketsWhileFibSwitchingFlag bool `protobuf:"varint,76,opt,name=drop_packets_while_fib_switching_flag,json=dropPacketsWhileFibSwitchingFlag,proto3" json:"drop_packets_while_fib_switching_flag,omitempty"`
	// Punt packets from linecard flag
	PuntPacketsFromLinecardFlag bool `protobuf:"varint,77,opt,name=punt_packets_from_linecard_flag,json=puntPacketsFromLinecardFlag,proto3" json:"punt_packets_from_linecard_flag,omitempty"`
	// Pimary local v4 address for the interface
	PrimaryIpv4Address string `protobuf:"bytes,78,opt,name=primary_ipv4_address,json=primaryIpv4Address,proto3" json:"primary_ipv4_address,omitempty"`
	// Pimary local v6 address for the interface
	PrimaryIpv6Address string `protobuf:"bytes,79,opt,name=primary_ipv6_address,json=primaryIpv6Address,proto3" json:"primary_ipv6_address,omitempty"`
	// Internal Information
	SiInternal           *FibShIntInternal `protobuf:"bytes,80,opt,name=si_internal,json=siInternal,proto3" json:"si_internal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FibShInt) Reset()         { *m = FibShInt{} }
func (m *FibShInt) String() string { return proto.CompactTextString(m) }
func (*FibShInt) ProtoMessage()    {}
func (*FibShInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_int_1b0c7d47e5664c2f, []int{1}
}
func (m *FibShInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShInt.Unmarshal(m, b)
}
func (m *FibShInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShInt.Marshal(b, m, deterministic)
}
func (dst *FibShInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShInt.Merge(dst, src)
}
func (m *FibShInt) XXX_Size() int {
	return xxx_messageInfo_FibShInt.Size(m)
}
func (m *FibShInt) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShInt.DiscardUnknown(m)
}

var xxx_messageInfo_FibShInt proto.InternalMessageInfo

func (m *FibShInt) GetPerInterface() string {
	if m != nil {
		return m.PerInterface
	}
	return ""
}

func (m *FibShInt) GetFibInterfaceType() uint32 {
	if m != nil {
		return m.FibInterfaceType
	}
	return 0
}

func (m *FibShInt) GetFibIdPointer() uint32 {
	if m != nil {
		return m.FibIdPointer
	}
	return 0
}

func (m *FibShInt) GetFibIdFlags() uint32 {
	if m != nil {
		return m.FibIdFlags
	}
	return 0
}

func (m *FibShInt) GetFibIdExtensionPointer() uint32 {
	if m != nil {
		return m.FibIdExtensionPointer
	}
	return 0
}

func (m *FibShInt) GetFibIdExtensionFlags() uint32 {
	if m != nil {
		return m.FibIdExtensionFlags
	}
	return 0
}

func (m *FibShInt) GetNumberOfDependentNextHopInformation() uint32 {
	if m != nil {
		return m.NumberOfDependentNextHopInformation
	}
	return 0
}

func (m *FibShInt) GetVrfLocalCefInformationPointer() uint32 {
	if m != nil {
		return m.VrfLocalCefInformationPointer
	}
	return 0
}

func (m *FibShInt) GetReferenceCount() uint32 {
	if m != nil {
		return m.ReferenceCount
	}
	return 0
}

func (m *FibShInt) GetLastModifiedTime() uint32 {
	if m != nil {
		return m.LastModifiedTime
	}
	return 0
}

func (m *FibShInt) GetLastOperation() string {
	if m != nil {
		return m.LastOperation
	}
	return ""
}

func (m *FibShInt) GetProtocolEnabled() bool {
	if m != nil {
		return m.ProtocolEnabled
	}
	return false
}

func (m *FibShInt) GetDetailFibIntInformation() *FibShIntDet {
	if m != nil {
		return m.DetailFibIntInformation
	}
	return nil
}

func (m *FibShInt) GetReferanceCountForProtocol() uint32 {
	if m != nil {
		return m.ReferanceCountForProtocol
	}
	return 0
}

func (m *FibShInt) GetNumberOfInputPackets() uint64 {
	if m != nil {
		return m.NumberOfInputPackets
	}
	return 0
}

func (m *FibShInt) GetNumberOfInputBytes() uint64 {
	if m != nil {
		return m.NumberOfInputBytes
	}
	return 0
}

func (m *FibShInt) GetNumberOfOutputPackets() uint64 {
	if m != nil {
		return m.NumberOfOutputPackets
	}
	return 0
}

func (m *FibShInt) GetNumberOfOutputBytes() uint64 {
	if m != nil {
		return m.NumberOfOutputBytes
	}
	return 0
}

func (m *FibShInt) GetInterfaceUpFlag() bool {
	if m != nil {
		return m.InterfaceUpFlag
	}
	return false
}

func (m *FibShInt) GetPerPacketLoadBalancingFlag() bool {
	if m != nil {
		return m.PerPacketLoadBalancingFlag
	}
	return false
}

func (m *FibShInt) GetP2PInterfaceFlag() bool {
	if m != nil {
		return m.P2PInterfaceFlag
	}
	return false
}

func (m *FibShInt) GetLoopbackInterfaceFlag() bool {
	if m != nil {
		return m.LoopbackInterfaceFlag
	}
	return false
}

func (m *FibShInt) GetNullInterfaceFlag() bool {
	if m != nil {
		return m.NullInterfaceFlag
	}
	return false
}

func (m *FibShInt) GetTunnelInterfaceFlag() bool {
	if m != nil {
		return m.TunnelInterfaceFlag
	}
	return false
}

func (m *FibShInt) GetGreTunnelInterfaceFlag() bool {
	if m != nil {
		return m.GreTunnelInterfaceFlag
	}
	return false
}

func (m *FibShInt) GetPuntPacketsFromFibSwitchingFlag() bool {
	if m != nil {
		return m.PuntPacketsFromFibSwitchingFlag
	}
	return false
}

func (m *FibShInt) GetDropPacketsWhileFibSwitchingFlag() bool {
	if m != nil {
		return m.DropPacketsWhileFibSwitchingFlag
	}
	return false
}

func (m *FibShInt) GetPuntPacketsFromLinecardFlag() bool {
	if m != nil {
		return m.PuntPacketsFromLinecardFlag
	}
	return false
}

func (m *FibShInt) GetPrimaryIpv4Address() string {
	if m != nil {
		return m.PrimaryIpv4Address
	}
	return ""
}

func (m *FibShInt) GetPrimaryIpv6Address() string {
	if m != nil {
		return m.PrimaryIpv6Address
	}
	return ""
}

func (m *FibShInt) GetSiInternal() *FibShIntInternal {
	if m != nil {
		return m.SiInternal
	}
	return nil
}

// Event history Entry
type EvtHistEntry struct {
	// The timestamp of the event
	EvtTimestamp string `protobuf:"bytes,1,opt,name=evt_timestamp,json=evtTimestamp,proto3" json:"evt_timestamp,omitempty"`
	// Event name
	EvtName string `protobuf:"bytes,2,opt,name=evt_name,json=evtName,proto3" json:"evt_name,omitempty"`
	// Event type
	EvtType uint32 `protobuf:"varint,3,opt,name=evt_type,json=evtType,proto3" json:"evt_type,omitempty"`
	// Multiple instance flag
	EvtMany bool `protobuf:"varint,4,opt,name=evt_many,json=evtMany,proto3" json:"evt_many,omitempty"`
	// Sticky flag
	EvtSticky bool `protobuf:"varint,5,opt,name=evt_sticky,json=evtSticky,proto3" json:"evt_sticky,omitempty"`
	// Optional data
	EvtData              []uint32 `protobuf:"varint,6,rep,packed,name=evt_data,json=evtData,proto3" json:"evt_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvtHistEntry) Reset()         { *m = EvtHistEntry{} }
func (m *EvtHistEntry) String() string { return proto.CompactTextString(m) }
func (*EvtHistEntry) ProtoMessage()    {}
func (*EvtHistEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_int_1b0c7d47e5664c2f, []int{2}
}
func (m *EvtHistEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvtHistEntry.Unmarshal(m, b)
}
func (m *EvtHistEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvtHistEntry.Marshal(b, m, deterministic)
}
func (dst *EvtHistEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvtHistEntry.Merge(dst, src)
}
func (m *EvtHistEntry) XXX_Size() int {
	return xxx_messageInfo_EvtHistEntry.Size(m)
}
func (m *EvtHistEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_EvtHistEntry.DiscardUnknown(m)
}

var xxx_messageInfo_EvtHistEntry proto.InternalMessageInfo

func (m *EvtHistEntry) GetEvtTimestamp() string {
	if m != nil {
		return m.EvtTimestamp
	}
	return ""
}

func (m *EvtHistEntry) GetEvtName() string {
	if m != nil {
		return m.EvtName
	}
	return ""
}

func (m *EvtHistEntry) GetEvtType() uint32 {
	if m != nil {
		return m.EvtType
	}
	return 0
}

func (m *EvtHistEntry) GetEvtMany() bool {
	if m != nil {
		return m.EvtMany
	}
	return false
}

func (m *EvtHistEntry) GetEvtSticky() bool {
	if m != nil {
		return m.EvtSticky
	}
	return false
}

func (m *EvtHistEntry) GetEvtData() []uint32 {
	if m != nil {
		return m.EvtData
	}
	return nil
}

// Per object event history
type EvtHistInfo struct {
	// Class name string
	EvtClassName string `protobuf:"bytes,1,opt,name=evt_class_name,json=evtClassName,proto3" json:"evt_class_name,omitempty"`
	// Array of event entries
	EvtEntry             []*EvtHistEntry `protobuf:"bytes,2,rep,name=evt_entry,json=evtEntry,proto3" json:"evt_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EvtHistInfo) Reset()         { *m = EvtHistInfo{} }
func (m *EvtHistInfo) String() string { return proto.CompactTextString(m) }
func (*EvtHistInfo) ProtoMessage()    {}
func (*EvtHistInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_int_1b0c7d47e5664c2f, []int{3}
}
func (m *EvtHistInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvtHistInfo.Unmarshal(m, b)
}
func (m *EvtHistInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvtHistInfo.Marshal(b, m, deterministic)
}
func (dst *EvtHistInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvtHistInfo.Merge(dst, src)
}
func (m *EvtHistInfo) XXX_Size() int {
	return xxx_messageInfo_EvtHistInfo.Size(m)
}
func (m *EvtHistInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EvtHistInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EvtHistInfo proto.InternalMessageInfo

func (m *EvtHistInfo) GetEvtClassName() string {
	if m != nil {
		return m.EvtClassName
	}
	return ""
}

func (m *EvtHistInfo) GetEvtEntry() []*EvtHistEntry {
	if m != nil {
		return m.EvtEntry
	}
	return nil
}

// FIB per interface detail information
type FibShIntDet struct {
	// Interface Protocol MTU
	InterfaceMtu uint32 `protobuf:"varint,1,opt,name=interface_mtu,json=interfaceMtu,proto3" json:"interface_mtu,omitempty"`
	// Forwarding enabled/disabled flag
	ForwardingFlag bool `protobuf:"varint,2,opt,name=forwarding_flag,json=forwardingFlag,proto3" json:"forwarding_flag,omitempty"`
	// RPF configured flag
	RpfConfiguredFlag bool `protobuf:"varint,3,opt,name=rpf_configured_flag,json=rpfConfiguredFlag,proto3" json:"rpf_configured_flag,omitempty"`
	// RPF mode
	RpfMode string `protobuf:"bytes,4,opt,name=rpf_mode,json=rpfMode,proto3" json:"rpf_mode,omitempty"`
	// Allow default route with RPF
	DefaultRouteWithRpf bool `protobuf:"varint,5,opt,name=default_route_with_rpf,json=defaultRouteWithRpf,proto3" json:"default_route_with_rpf,omitempty"`
	// Allow selfping with RPF
	SelfPingWithRpf bool `protobuf:"varint,6,opt,name=self_ping_with_rpf,json=selfPingWithRpf,proto3" json:"self_ping_with_rpf,omitempty"`
	// BGP PA configured flag
	BgpPaInputConfiguredFlag bool `protobuf:"varint,7,opt,name=bgp_pa_input_configured_flag,json=bgpPaInputConfiguredFlag,proto3" json:"bgp_pa_input_configured_flag,omitempty"`
	// src BGP PA configured flag
	SourceBgpPaInputConfiguredFlag bool `protobuf:"varint,8,opt,name=source_bgp_pa_input_configured_flag,json=sourceBgpPaInputConfiguredFlag,proto3" json:"source_bgp_pa_input_configured_flag,omitempty"`
	// dst BGP PA configured flag
	DestinationBgpPaInputConfiguredFlag bool `protobuf:"varint,9,opt,name=destination_bgp_pa_input_configured_flag,json=destinationBgpPaInputConfiguredFlag,proto3" json:"destination_bgp_pa_input_configured_flag,omitempty"`
	// BGP PA configured flag
	BgpPaOutputConfiguredFlag bool `protobuf:"varint,10,opt,name=bgp_pa_output_configured_flag,json=bgpPaOutputConfiguredFlag,proto3" json:"bgp_pa_output_configured_flag,omitempty"`
	// src BGP PA configured flag
	SourceBgpPaOutputConfiguredFlag bool `protobuf:"varint,11,opt,name=source_bgp_pa_output_configured_flag,json=sourceBgpPaOutputConfiguredFlag,proto3" json:"source_bgp_pa_output_configured_flag,omitempty"`
	// dst BGP PA configured flag
	DestinationBgpPaOutputConfiguredFlag bool `protobuf:"varint,12,opt,name=destination_bgp_pa_output_configured_flag,json=destinationBgpPaOutputConfiguredFlag,proto3" json:"destination_bgp_pa_output_configured_flag,omitempty"`
	// ICMP  configured flag
	IcmpFlag uint32 `protobuf:"varint,13,opt,name=icmp_flag,json=icmpFlag,proto3" json:"icmp_flag,omitempty"`
	// Drop packets with multiple-label-stack if set
	MultiLabelDropFlag   bool     `protobuf:"varint,14,opt,name=multi_label_drop_flag,json=multiLabelDropFlag,proto3" json:"multi_label_drop_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShIntDet) Reset()         { *m = FibShIntDet{} }
func (m *FibShIntDet) String() string { return proto.CompactTextString(m) }
func (*FibShIntDet) ProtoMessage()    {}
func (*FibShIntDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_int_1b0c7d47e5664c2f, []int{4}
}
func (m *FibShIntDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShIntDet.Unmarshal(m, b)
}
func (m *FibShIntDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShIntDet.Marshal(b, m, deterministic)
}
func (dst *FibShIntDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShIntDet.Merge(dst, src)
}
func (m *FibShIntDet) XXX_Size() int {
	return xxx_messageInfo_FibShIntDet.Size(m)
}
func (m *FibShIntDet) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShIntDet.DiscardUnknown(m)
}

var xxx_messageInfo_FibShIntDet proto.InternalMessageInfo

func (m *FibShIntDet) GetInterfaceMtu() uint32 {
	if m != nil {
		return m.InterfaceMtu
	}
	return 0
}

func (m *FibShIntDet) GetForwardingFlag() bool {
	if m != nil {
		return m.ForwardingFlag
	}
	return false
}

func (m *FibShIntDet) GetRpfConfiguredFlag() bool {
	if m != nil {
		return m.RpfConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetRpfMode() string {
	if m != nil {
		return m.RpfMode
	}
	return ""
}

func (m *FibShIntDet) GetDefaultRouteWithRpf() bool {
	if m != nil {
		return m.DefaultRouteWithRpf
	}
	return false
}

func (m *FibShIntDet) GetSelfPingWithRpf() bool {
	if m != nil {
		return m.SelfPingWithRpf
	}
	return false
}

func (m *FibShIntDet) GetBgpPaInputConfiguredFlag() bool {
	if m != nil {
		return m.BgpPaInputConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetSourceBgpPaInputConfiguredFlag() bool {
	if m != nil {
		return m.SourceBgpPaInputConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetDestinationBgpPaInputConfiguredFlag() bool {
	if m != nil {
		return m.DestinationBgpPaInputConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetBgpPaOutputConfiguredFlag() bool {
	if m != nil {
		return m.BgpPaOutputConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetSourceBgpPaOutputConfiguredFlag() bool {
	if m != nil {
		return m.SourceBgpPaOutputConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetDestinationBgpPaOutputConfiguredFlag() bool {
	if m != nil {
		return m.DestinationBgpPaOutputConfiguredFlag
	}
	return false
}

func (m *FibShIntDet) GetIcmpFlag() uint32 {
	if m != nil {
		return m.IcmpFlag
	}
	return 0
}

func (m *FibShIntDet) GetMultiLabelDropFlag() bool {
	if m != nil {
		return m.MultiLabelDropFlag
	}
	return false
}

// FIB per interface internal information
type FibShIntInternal struct {
	// Event History for IDB
	FibIdbHist *EvtHistInfo `protobuf:"bytes,1,opt,name=fib_idb_hist,json=fibIdbHist,proto3" json:"fib_idb_hist,omitempty"`
	// Event History for Srtehead
	FibSrteHeadHist      *EvtHistInfo `protobuf:"bytes,2,opt,name=fib_srte_head_hist,json=fibSrteHeadHist,proto3" json:"fib_srte_head_hist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FibShIntInternal) Reset()         { *m = FibShIntInternal{} }
func (m *FibShIntInternal) String() string { return proto.CompactTextString(m) }
func (*FibShIntInternal) ProtoMessage()    {}
func (*FibShIntInternal) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_sh_int_1b0c7d47e5664c2f, []int{5}
}
func (m *FibShIntInternal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShIntInternal.Unmarshal(m, b)
}
func (m *FibShIntInternal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShIntInternal.Marshal(b, m, deterministic)
}
func (dst *FibShIntInternal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShIntInternal.Merge(dst, src)
}
func (m *FibShIntInternal) XXX_Size() int {
	return xxx_messageInfo_FibShIntInternal.Size(m)
}
func (m *FibShIntInternal) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShIntInternal.DiscardUnknown(m)
}

var xxx_messageInfo_FibShIntInternal proto.InternalMessageInfo

func (m *FibShIntInternal) GetFibIdbHist() *EvtHistInfo {
	if m != nil {
		return m.FibIdbHist
	}
	return nil
}

func (m *FibShIntInternal) GetFibSrteHeadHist() *EvtHistInfo {
	if m != nil {
		return m.FibSrteHeadHist
	}
	return nil
}

func init() {
	proto.RegisterType((*FibShInt_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.interface_infos.interface_info.interfaces.interface.fib_sh_int_KEYS")
	proto.RegisterType((*FibShInt)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.interface_infos.interface_info.interfaces.interface.fib_sh_int")
	proto.RegisterType((*EvtHistEntry)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.interface_infos.interface_info.interfaces.interface.evt_hist_entry")
	proto.RegisterType((*EvtHistInfo)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.interface_infos.interface_info.interfaces.interface.evt_hist_info")
	proto.RegisterType((*FibShIntDet)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.interface_infos.interface_info.interfaces.interface.fib_sh_int_det")
	proto.RegisterType((*FibShIntInternal)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.interface_infos.interface_info.interfaces.interface.fib_sh_int_internal")
}

func init() { proto.RegisterFile("fib_sh_int.proto", fileDescriptor_fib_sh_int_1b0c7d47e5664c2f) }

var fileDescriptor_fib_sh_int_1b0c7d47e5664c2f = []byte{
	// 1451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xc9, 0x53, 0x1c, 0x37,
	0x17, 0xaf, 0x06, 0x7f, 0x30, 0x88, 0xcd, 0x6e, 0x8c, 0xdd, 0x18, 0x2f, 0xf3, 0x01, 0xae, 0x0f,
	0x7f, 0x49, 0xa6, 0x62, 0xb0, 0x71, 0x9c, 0xc5, 0x0b, 0x5b, 0x20, 0x06, 0x33, 0x35, 0xe0, 0x72,
	0xe5, 0xa4, 0x52, 0x77, 0xab, 0x67, 0x54, 0x74, 0x4b, 0x2a, 0xb5, 0x7a, 0x60, 0x2a, 0x97, 0x9c,
	0x93, 0x4b, 0xfe, 0x0d, 0x1f, 0x72, 0xf1, 0x29, 0xb9, 0xe5, 0x9a, 0x3f, 0x23, 0xff, 0x49, 0x4a,
	0x4f, 0xdd, 0x3d, 0x0b, 0x83, 0xaf, 0xf6, 0x85, 0x1a, 0xde, 0x6f, 0xd1, 0x7b, 0x5a, 0x9e, 0xd4,
	0xe8, 0x6a, 0xc4, 0x7c, 0x9c, 0xb6, 0x30, 0xe3, 0xba, 0x26, 0x95, 0xd0, 0xc2, 0xfd, 0x29, 0x60,
	0x69, 0x20, 0x30, 0x13, 0x29, 0x3e, 0x57, 0xd8, 0xc0, 0x81, 0x48, 0x12, 0xc1, 0xb1, 0x90, 0x54,
	0xd5, 0x22, 0xe6, 0xd7, 0xb8, 0x08, 0x69, 0x0a, 0x7f, 0xad, 0x24, 0x10, 0x71, 0x5a, 0xfe, 0xaa,
	0xb5, 0x55, 0x94, 0x9a, 0x3f, 0x35, 0xc6, 0x35, 0x55, 0x11, 0x09, 0x28, 0x66, 0x3c, 0x12, 0xe9,
	0xc0, 0xff, 0xdd, 0x7f, 0x7b, 0x90, 0xa5, 0xf7, 0x0e, 0x9a, 0xed, 0x66, 0x84, 0x5f, 0xed, 0xfc,
	0x78, 0xec, 0x2e, 0xa2, 0x09, 0x33, 0x16, 0xe6, 0x24, 0xa1, 0x9e, 0x53, 0x75, 0x56, 0x27, 0x1a,
	0x15, 0x13, 0x78, 0x4d, 0x12, 0xea, 0x2e, 0xa3, 0xe9, 0x62, 0x64, 0x4b, 0x18, 0x01, 0xc2, 0x54,
	0x11, 0x04, 0xd2, 0x02, 0xaa, 0xb4, 0x55, 0x64, 0xf1, 0x51, 0xc0, 0xc7, 0xdb, 0x2a, 0x02, 0x68,
	0x11, 0x4d, 0xc4, 0x8c, 0x9f, 0x62, 0xdd, 0x91, 0xd4, 0xbb, 0x62, 0xcd, 0x4d, 0xe0, 0xa4, 0x23,
	0xa9, 0x7b, 0x1f, 0xcd, 0x74, 0x93, 0x06, 0xf5, 0x7f, 0x80, 0x31, 0x5d, 0x46, 0x8d, 0xc7, 0xd2,
	0x9f, 0xb3, 0x08, 0x75, 0x93, 0x86, 0x94, 0xa8, 0xc2, 0x25, 0xc7, 0x5b, 0xcb, 0x53, 0xa2, 0x6a,
	0xbf, 0x88, 0xb9, 0x9f, 0x23, 0xd7, 0x48, 0xba, 0xf6, 0x90, 0xc0, 0x7a, 0xd5, 0x59, 0x9d, 0x6e,
	0x98, 0x35, 0x29, 0x99, 0x90, 0xc8, 0x0a, 0x9a, 0x01, 0x76, 0x88, 0xa5, 0x00, 0x89, 0xf7, 0x08,
	0x98, 0x53, 0x86, 0x19, 0xd6, 0x6d, 0xcc, 0xad, 0xa2, 0xa9, 0x9c, 0x15, 0xc5, 0xa4, 0x99, 0x7a,
	0x8f, 0x81, 0x83, 0x80, 0xb3, 0x6b, 0x22, 0xee, 0x13, 0xe4, 0xe5, 0x0c, 0x7a, 0xae, 0x29, 0x4f,
	0x99, 0xe0, 0xa5, 0xe3, 0x06, 0xb0, 0xe7, 0x81, 0xbd, 0x53, 0xa0, 0x85, 0xf5, 0x3a, 0xba, 0x71,
	0x41, 0x68, 0x07, 0x79, 0x02, 0xb2, 0xb9, 0x7e, 0x99, 0x1d, 0xed, 0x0d, 0x5a, 0xe5, 0x59, 0xe2,
	0x53, 0x85, 0x45, 0x84, 0x43, 0x2a, 0x29, 0x0f, 0x29, 0xd7, 0x98, 0xd3, 0x73, 0x8d, 0x5b, 0x42,
	0xc2, 0x36, 0x50, 0x09, 0xd1, 0x4c, 0x70, 0xef, 0x2b, 0xb0, 0x59, 0xb6, 0xfc, 0xa3, 0x68, 0xbb,
	0x60, 0xbf, 0xa6, 0xe7, 0x7a, 0x4f, 0xc8, 0xfd, 0x2e, 0xd5, 0xdd, 0x43, 0xff, 0x35, 0xab, 0x19,
	0x8b, 0x80, 0xc4, 0x38, 0xa0, 0x51, 0xaf, 0x4f, 0x59, 0xcd, 0x53, 0xf0, 0xbb, 0xd3, 0x56, 0xd1,
	0x81, 0xe1, 0x6d, 0xd1, 0xa8, 0xc7, 0xa2, 0xa8, 0xea, 0x7f, 0x68, 0x56, 0xd1, 0x88, 0x2a, 0xca,
	0x03, 0x8a, 0x03, 0x91, 0x71, 0xed, 0x7d, 0x0d, 0xba, 0x99, 0x32, 0xbc, 0x65, 0xa2, 0x66, 0xb5,
	0x62, 0x92, 0x6a, 0x9c, 0x88, 0x90, 0x45, 0x8c, 0x86, 0x58, 0xb3, 0x84, 0x7a, 0xdf, 0xd8, 0xd5,
	0x32, 0xc8, 0x61, 0x0e, 0x9c, 0xb0, 0x04, 0xb6, 0x0d, 0xb0, 0xcd, 0x81, 0xb1, 0xd5, 0x7d, 0x6b,
	0xb7, 0x8d, 0x89, 0x1e, 0x15, 0x41, 0xf7, 0x01, 0xba, 0x5a, 0x6e, 0x5d, 0xca, 0x89, 0x1f, 0xd3,
	0xd0, 0xfb, 0xae, 0xea, 0xac, 0x56, 0x1a, 0xb3, 0x45, 0x7c, 0xc7, 0x86, 0xdd, 0xbf, 0x1d, 0x74,
	0x2b, 0xa4, 0x9a, 0xb0, 0x18, 0xe7, 0xbb, 0xa6, 0x6f, 0xf2, 0x9e, 0x55, 0x9d, 0xd5, 0xc9, 0xb5,
	0x5f, 0x9d, 0xda, 0x47, 0x3c, 0xba, 0xb5, 0x9e, 0x63, 0x1b, 0x52, 0xdd, 0xb8, 0x69, 0xf3, 0xdd,
	0x85, 0xad, 0xdc, 0xbb, 0x7c, 0xcf, 0xd1, 0x6d, 0x98, 0x5d, 0x52, 0x4e, 0x3a, 0x8e, 0x84, 0xc2,
	0x45, 0x12, 0xde, 0x73, 0x98, 0xd5, 0x85, 0x92, 0x03, 0x2b, 0xb0, 0x2b, 0x54, 0x3d, 0x27, 0xb8,
	0x8f, 0xd1, 0xcd, 0xee, 0xb6, 0x62, 0x5c, 0x66, 0x1a, 0x4b, 0x12, 0x9c, 0x52, 0x9d, 0x7a, 0x2f,
	0xaa, 0xce, 0xea, 0x95, 0xc6, 0xf5, 0x62, 0x17, 0xed, 0x1b, 0xb0, 0x6e, 0x31, 0xf7, 0x21, 0x9a,
	0x1f, 0x94, 0xf9, 0x1d, 0x4d, 0x53, 0xef, 0x25, 0x88, 0xdc, 0x3e, 0xd1, 0xa6, 0x41, 0xcc, 0x71,
	0xe9, 0x4a, 0x44, 0xa6, 0x7b, 0x87, 0xda, 0x04, 0xd5, 0x7c, 0xa1, 0x3a, 0x02, 0xb4, 0x18, 0x6b,
	0x1d, 0xdd, 0xb8, 0x20, 0xb4, 0x83, 0x6d, 0x81, 0x6c, 0xae, 0x5f, 0x66, 0x47, 0xfb, 0x3f, 0xba,
	0xd6, 0x9d, 0xe7, 0x4c, 0xc2, 0xf9, 0xf2, 0xb6, 0xed, 0x86, 0x28, 0x81, 0x37, 0xd2, 0x9c, 0x2d,
	0x77, 0x13, 0xdd, 0x35, 0x3d, 0xc6, 0x26, 0x83, 0x63, 0x41, 0x42, 0xec, 0x93, 0x98, 0xf0, 0x80,
	0xf1, 0xa6, 0x15, 0xee, 0x80, 0xf0, 0x96, 0xa4, 0xca, 0x26, 0x75, 0x20, 0x48, 0xb8, 0x59, 0x50,
	0xc0, 0xe3, 0x0b, 0x34, 0x27, 0xd7, 0xb0, 0xec, 0xe9, 0x41, 0x20, 0xdc, 0x05, 0xe1, 0x55, 0xb9,
	0x56, 0x2f, 0x7b, 0x10, 0xd0, 0x37, 0xd0, 0xcd, 0x58, 0x08, 0xe9, 0x93, 0xe0, 0x74, 0x50, 0xf2,
	0x3d, 0x48, 0xe6, 0x0b, 0xb8, 0x5f, 0x57, 0x43, 0x73, 0x3c, 0x8b, 0xe3, 0x41, 0xcd, 0x1e, 0x68,
	0xae, 0x19, 0xa8, 0x9f, 0xbf, 0x86, 0xe6, 0x75, 0xc6, 0x39, 0xbd, 0xa0, 0xd8, 0x07, 0xc5, 0x9c,
	0x05, 0xfb, 0x35, 0x4f, 0xd1, 0x42, 0x53, 0x51, 0x3c, 0x5c, 0xf7, 0x03, 0xe8, 0x6e, 0x34, 0x15,
	0x3d, 0x19, 0x22, 0x3d, 0x44, 0x2b, 0xd2, 0xec, 0xc1, 0x7c, 0x5d, 0x71, 0xa4, 0x44, 0x02, 0x47,
	0x27, 0x3d, 0x63, 0x3a, 0x68, 0x95, 0xf3, 0xf9, 0x0a, 0x5c, 0xee, 0x19, 0x6e, 0xbe, 0xca, 0xbb,
	0x4a, 0x24, 0xbb, 0xcc, 0x3f, 0x2e, 0x78, 0x60, 0x77, 0x84, 0xee, 0x87, 0x4a, 0xc8, 0xd2, 0xee,
	0xac, 0xc5, 0x62, 0x3a, 0xcc, 0xef, 0x00, 0xfc, 0xaa, 0x86, 0x9c, 0xfb, 0xbd, 0x35, 0xd4, 0x0b,
	0x86, 0xdb, 0xe8, 0xde, 0xc5, 0xfc, 0x62, 0xc6, 0x69, 0x40, 0x94, 0x6d, 0xf4, 0xde, 0x21, 0x58,
	0x2d, 0x0e, 0xa4, 0x76, 0x90, 0x73, 0xc0, 0xe5, 0x4b, 0x74, 0x5d, 0x2a, 0x96, 0x10, 0xd5, 0xc1,
	0x4c, 0xb6, 0x1f, 0x61, 0x12, 0x86, 0x8a, 0xa6, 0xa9, 0xf7, 0x1a, 0x1a, 0x93, 0x9b, 0x63, 0xfb,
	0xb2, 0xfd, 0xe8, 0xa5, 0x45, 0x06, 0x14, 0x1b, 0xa5, 0xe2, 0x68, 0x50, 0xb1, 0x51, 0x28, 0xde,
	0x3b, 0x68, 0x32, 0x65, 0x76, 0xf6, 0x39, 0x89, 0xbd, 0x3a, 0x74, 0xa5, 0xdf, 0x3e, 0x99, 0xae,
	0x54, 0x24, 0xd6, 0x40, 0x29, 0xdb, 0xcf, 0x7f, 0x2f, 0xfd, 0xe5, 0xa0, 0x19, 0xda, 0xd6, 0xb8,
	0xc5, 0x52, 0x8d, 0x29, 0xd7, 0xaa, 0x63, 0xee, 0x6f, 0x13, 0x31, 0x2d, 0x3e, 0xd5, 0x24, 0x91,
	0xf9, 0x9b, 0x63, 0x8a, 0xb6, 0xf5, 0x49, 0x11, 0x33, 0x4f, 0x0a, 0x43, 0xea, 0x79, 0x72, 0x8c,
	0xd3, 0xb6, 0x2e, 0x5e, 0x1b, 0xa0, 0x37, 0x17, 0xfa, 0x28, 0x34, 0x33, 0x03, 0xc1, 0x3d, 0x9e,
	0x43, 0x09, 0xe1, 0x1d, 0x78, 0x6c, 0x54, 0x00, 0x3a, 0x24, 0xbc, 0xe3, 0xde, 0x41, 0xc8, 0x40,
	0xa9, 0x66, 0xc1, 0x69, 0x07, 0xde, 0x19, 0x95, 0xc6, 0x04, 0x6d, 0xeb, 0x63, 0x08, 0x14, 0xca,
	0x90, 0x68, 0xe2, 0x8d, 0x55, 0x47, 0x73, 0xd3, 0x6d, 0xa2, 0xc9, 0xd2, 0x3f, 0x8e, 0x4d, 0x18,
	0x4a, 0x30, 0x33, 0x61, 0x9e, 0x0b, 0x26, 0x10, 0xc4, 0x24, 0x4d, 0x7b, 0x9f, 0x4d, 0xa6, 0x84,
	0x2d, 0x13, 0x84, 0x3c, 0xdf, 0x39, 0xc8, 0x0c, 0x60, 0xab, 0xf6, 0x46, 0xaa, 0xa3, 0x1f, 0xff,
	0x0e, 0xe9, 0x5f, 0x89, 0x86, 0x29, 0x79, 0xc7, 0xfc, 0x5a, 0x7a, 0x37, 0x66, 0x5f, 0x40, 0xdd,
	0x0b, 0xc6, 0x2c, 0x53, 0xd7, 0x32, 0xd1, 0x19, 0xd4, 0x38, 0xdd, 0x98, 0x2a, 0x83, 0x87, 0x3a,
	0x33, 0x37, 0x7c, 0x24, 0xd4, 0x19, 0x51, 0x61, 0x79, 0xf0, 0x46, 0x60, 0x6a, 0x67, 0xba, 0xe1,
	0xa2, 0x4b, 0x29, 0x19, 0xe1, 0x40, 0xf0, 0x88, 0x35, 0x33, 0x45, 0xf3, 0xa3, 0x35, 0x6a, 0xbb,
	0x94, 0x92, 0xd1, 0x56, 0x89, 0x00, 0x7f, 0x01, 0x55, 0x0c, 0x3f, 0x11, 0x61, 0xf1, 0x6c, 0x1c,
	0x57, 0x32, 0x3a, 0x14, 0x21, 0x35, 0xcd, 0x3f, 0xa4, 0x11, 0xc9, 0x62, 0x8d, 0x95, 0xc8, 0x34,
	0xc5, 0x67, 0x4c, 0xb7, 0xb0, 0x92, 0x51, 0xbe, 0xaa, 0x73, 0x39, 0xda, 0x30, 0xe0, 0x5b, 0xa6,
	0x5b, 0x0d, 0x19, 0xb9, 0x9f, 0x21, 0x37, 0xa5, 0x71, 0x84, 0xa5, 0xc9, 0xb3, 0x14, 0x8c, 0xd9,
	0xee, 0x6f, 0x90, 0x3a, 0xe3, 0xcd, 0x82, 0xfc, 0x0c, 0xdd, 0xf6, 0x9b, 0xa6, 0xc7, 0xe4, 0xf7,
	0xd8, 0x60, 0xd6, 0xe3, 0x20, 0xf3, 0xfc, 0xa6, 0xac, 0x13, 0xb8, 0xce, 0x06, 0x92, 0x7f, 0x85,
	0x96, 0x53, 0x91, 0xa9, 0x80, 0xe2, 0x0f, 0xda, 0x54, 0xc0, 0xe6, 0xae, 0xa5, 0x6e, 0x5e, 0x66,
	0xf6, 0x06, 0xad, 0x86, 0x34, 0xd5, 0x8c, 0xdb, 0x07, 0xd8, 0x07, 0x1d, 0x27, 0xc0, 0x71, 0xb9,
	0x87, 0x7f, 0xa9, 0xed, 0x0b, 0x74, 0x27, 0xb7, 0xca, 0xef, 0xcf, 0x41, 0x2f, 0x04, 0x5e, 0x0b,
	0x50, 0xa4, 0xbd, 0x46, 0x07, 0x1c, 0x0e, 0xd1, 0x4a, 0x7f, 0x95, 0x97, 0x18, 0x4d, 0xda, 0xce,
	0xde, 0x53, 0xe6, 0x50, 0xbb, 0xb7, 0xe8, 0xc1, 0x90, 0x3a, 0x2f, 0xf1, 0x9c, 0x02, 0xcf, 0x95,
	0xc1, 0x42, 0x87, 0x1a, 0x2f, 0xa2, 0x09, 0x16, 0x24, 0xf9, 0x7d, 0x3f, 0x0d, 0x9b, 0xb8, 0x62,
	0x02, 0x00, 0x3e, 0x44, 0xf3, 0x49, 0x16, 0x6b, 0x86, 0x63, 0xe2, 0xd3, 0x18, 0xc3, 0xdd, 0x02,
	0xc4, 0x19, 0x18, 0xc1, 0x05, 0xf0, 0xc0, 0x60, 0xdb, 0x4a, 0x80, 0x64, 0xe9, 0xe7, 0x51, 0x34,
	0x37, 0xa4, 0xed, 0xb9, 0xbf, 0x3b, 0xc5, 0xf7, 0x81, 0x0f, 0x87, 0x0c, 0x0e, 0xcc, 0xe4, 0xda,
	0x2f, 0x9f, 0xc8, 0x91, 0x37, 0x94, 0xfc, 0x63, 0xc5, 0xdf, 0x63, 0xa9, 0x76, 0xff, 0x70, 0xec,
	0x37, 0x52, 0xaa, 0x34, 0xc5, 0x2d, 0x4a, 0x42, 0x9b, 0xf5, 0xc8, 0xa7, 0x97, 0xb5, 0xf9, 0x64,
	0x3d, 0x56, 0x9a, 0xee, 0x51, 0x12, 0x9a, 0xd4, 0xfd, 0x31, 0x18, 0x65, 0xfd, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0xad, 0xf6, 0xe8, 0x5e, 0x0f, 0x00, 0x00,
}
