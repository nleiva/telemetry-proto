// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ip_traffic.proto

package cisco_ios_xr_ipv4_io_oper_ipv4_network_nodes_node_statistics_traffic

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

// IP and ICMP Traffic Information
type IpTraffic_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpTraffic_KEYS) Reset()         { *m = IpTraffic_KEYS{} }
func (m *IpTraffic_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpTraffic_KEYS) ProtoMessage()    {}
func (*IpTraffic_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ip_traffic_268f01bbe9b94ef2, []int{0}
}
func (m *IpTraffic_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpTraffic_KEYS.Unmarshal(m, b)
}
func (m *IpTraffic_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpTraffic_KEYS.Marshal(b, m, deterministic)
}
func (dst *IpTraffic_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpTraffic_KEYS.Merge(dst, src)
}
func (m *IpTraffic_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpTraffic_KEYS.Size(m)
}
func (m *IpTraffic_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpTraffic_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpTraffic_KEYS proto.InternalMessageInfo

func (m *IpTraffic_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type IpTraffic struct {
	// IPv4 Network Stats
	Ipv4Stats *Ipv4IoTraffic `protobuf:"bytes,50,opt,name=ipv4_stats,json=ipv4Stats,proto3" json:"ipv4_stats,omitempty"`
	// ICMP Stats
	IcmpStats            *Ipv4IoIcmpTraffic `protobuf:"bytes,51,opt,name=icmp_stats,json=icmpStats,proto3" json:"icmp_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IpTraffic) Reset()         { *m = IpTraffic{} }
func (m *IpTraffic) String() string { return proto.CompactTextString(m) }
func (*IpTraffic) ProtoMessage()    {}
func (*IpTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_ip_traffic_268f01bbe9b94ef2, []int{1}
}
func (m *IpTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpTraffic.Unmarshal(m, b)
}
func (m *IpTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpTraffic.Marshal(b, m, deterministic)
}
func (dst *IpTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpTraffic.Merge(dst, src)
}
func (m *IpTraffic) XXX_Size() int {
	return xxx_messageInfo_IpTraffic.Size(m)
}
func (m *IpTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_IpTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_IpTraffic proto.InternalMessageInfo

func (m *IpTraffic) GetIpv4Stats() *Ipv4IoTraffic {
	if m != nil {
		return m.Ipv4Stats
	}
	return nil
}

func (m *IpTraffic) GetIcmpStats() *Ipv4IoIcmpTraffic {
	if m != nil {
		return m.IcmpStats
	}
	return nil
}

// IP Traffic Information
type Ipv4IoTraffic struct {
	// Input Packets
	InputPackets uint32 `protobuf:"varint,1,opt,name=input_packets,json=inputPackets,proto3" json:"input_packets,omitempty"`
	// Received Packets
	ReceivedPackets uint32 `protobuf:"varint,2,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	// Format Errors
	FormatErrors uint32 `protobuf:"varint,3,opt,name=format_errors,json=formatErrors,proto3" json:"format_errors,omitempty"`
	// Bad Hop Count
	BadHopCount uint32 `protobuf:"varint,4,opt,name=bad_hop_count,json=badHopCount,proto3" json:"bad_hop_count,omitempty"`
	// Bad Source Address
	BadSourceAddress uint32 `protobuf:"varint,5,opt,name=bad_source_address,json=badSourceAddress,proto3" json:"bad_source_address,omitempty"`
	// Bad Header
	BadHeader uint32 `protobuf:"varint,6,opt,name=bad_header,json=badHeader,proto3" json:"bad_header,omitempty"`
	// No Protocol
	NoProtocol uint32 `protobuf:"varint,7,opt,name=no_protocol,json=noProtocol,proto3" json:"no_protocol,omitempty"`
	// No Gateway
	NoGateway uint32 `protobuf:"varint,8,opt,name=no_gateway,json=noGateway,proto3" json:"no_gateway,omitempty"`
	// RaInput
	ReassembleInput uint32 `protobuf:"varint,9,opt,name=reassemble_input,json=reassembleInput,proto3" json:"reassemble_input,omitempty"`
	// Reassembled
	Reassembled uint32 `protobuf:"varint,10,opt,name=reassembled,proto3" json:"reassembled,omitempty"`
	// Reassembly Timeout
	ReassembleTimeout uint32 `protobuf:"varint,11,opt,name=reassemble_timeout,json=reassembleTimeout,proto3" json:"reassemble_timeout,omitempty"`
	// Reassembly Max Drop
	ReassembleMaxDrop uint32 `protobuf:"varint,12,opt,name=reassemble_max_drop,json=reassembleMaxDrop,proto3" json:"reassemble_max_drop,omitempty"`
	// Reassembly Failed
	ReassembleFailed uint32 `protobuf:"varint,13,opt,name=reassemble_failed,json=reassembleFailed,proto3" json:"reassemble_failed,omitempty"`
	// IP Options Present
	OptionsPresent uint32 `protobuf:"varint,14,opt,name=options_present,json=optionsPresent,proto3" json:"options_present,omitempty"`
	// Bad Option
	BadOption uint32 `protobuf:"varint,15,opt,name=bad_option,json=badOption,proto3" json:"bad_option,omitempty"`
	// Unknown Option
	UnknownOption uint32 `protobuf:"varint,16,opt,name=unknown_option,json=unknownOption,proto3" json:"unknown_option,omitempty"`
	// Bad Security Option
	BadSecurityOption uint32 `protobuf:"varint,17,opt,name=bad_security_option,json=badSecurityOption,proto3" json:"bad_security_option,omitempty"`
	// Basic Security Option
	BasicSecurityOption uint32 `protobuf:"varint,18,opt,name=basic_security_option,json=basicSecurityOption,proto3" json:"basic_security_option,omitempty"`
	// Extended Security Option
	ExtendedSecurityOption uint32 `protobuf:"varint,19,opt,name=extended_security_option,json=extendedSecurityOption,proto3" json:"extended_security_option,omitempty"`
	// Cipso Option
	CipsoOption uint32 `protobuf:"varint,20,opt,name=cipso_option,json=cipsoOption,proto3" json:"cipso_option,omitempty"`
	// Strict Source Route Option
	StrictSourceRouteOption uint32 `protobuf:"varint,21,opt,name=strict_source_route_option,json=strictSourceRouteOption,proto3" json:"strict_source_route_option,omitempty"`
	// Loose Source Route Option
	LooseSourceRouteOption uint32 `protobuf:"varint,22,opt,name=loose_source_route_option,json=looseSourceRouteOption,proto3" json:"loose_source_route_option,omitempty"`
	// Record Route Option
	RecordRouteOption uint32 `protobuf:"varint,23,opt,name=record_route_option,json=recordRouteOption,proto3" json:"record_route_option,omitempty"`
	// SID Option
	SidOption uint32 `protobuf:"varint,24,opt,name=sid_option,json=sidOption,proto3" json:"sid_option,omitempty"`
	// Timestamp Option
	TimestampOption uint32 `protobuf:"varint,25,opt,name=timestamp_option,json=timestampOption,proto3" json:"timestamp_option,omitempty"`
	// Router Alert Option
	RouterAlertOption uint32 `protobuf:"varint,26,opt,name=router_alert_option,json=routerAlertOption,proto3" json:"router_alert_option,omitempty"`
	// Noop Option
	NoopOption uint32 `protobuf:"varint,27,opt,name=noop_option,json=noopOption,proto3" json:"noop_option,omitempty"`
	// End Option
	EndOption uint32 `protobuf:"varint,28,opt,name=end_option,json=endOption,proto3" json:"end_option,omitempty"`
	// Packets Output
	PacketsOutput uint32 `protobuf:"varint,29,opt,name=packets_output,json=packetsOutput,proto3" json:"packets_output,omitempty"`
	// Packets Forwarded
	PacketsForwarded uint32 `protobuf:"varint,30,opt,name=packets_forwarded,json=packetsForwarded,proto3" json:"packets_forwarded,omitempty"`
	// Packets Fragmented
	PacketsFragmented uint32 `protobuf:"varint,31,opt,name=packets_fragmented,json=packetsFragmented,proto3" json:"packets_fragmented,omitempty"`
	// Fragment Count
	FragmentCount uint32 `protobuf:"varint,32,opt,name=fragment_count,json=fragmentCount,proto3" json:"fragment_count,omitempty"`
	// Encapsulation Failed
	EncapsulationFailed uint32 `protobuf:"varint,33,opt,name=encapsulation_failed,json=encapsulationFailed,proto3" json:"encapsulation_failed,omitempty"`
	// No Router
	NoRouter uint32 `protobuf:"varint,34,opt,name=no_router,json=noRouter,proto3" json:"no_router,omitempty"`
	// Packet Too Big
	PacketTooBig uint32 `protobuf:"varint,35,opt,name=packet_too_big,json=packetTooBig,proto3" json:"packet_too_big,omitempty"`
	// Multicast In
	MulticastIn uint32 `protobuf:"varint,36,opt,name=multicast_in,json=multicastIn,proto3" json:"multicast_in,omitempty"`
	// Multicast Out
	MulticastOut uint32 `protobuf:"varint,37,opt,name=multicast_out,json=multicastOut,proto3" json:"multicast_out,omitempty"`
	// Broadcast In
	BroadcastIn uint32 `protobuf:"varint,38,opt,name=broadcast_in,json=broadcastIn,proto3" json:"broadcast_in,omitempty"`
	// Broadcast Out
	BroadcastOut uint32 `protobuf:"varint,39,opt,name=broadcast_out,json=broadcastOut,proto3" json:"broadcast_out,omitempty"`
	// Lisp IPv4 encapped packets
	LispV4Encap uint32 `protobuf:"varint,40,opt,name=lisp_v4_encap,json=lispV4Encap,proto3" json:"lisp_v4_encap,omitempty"`
	// Lisp IPv4 decapped packets
	LispV4Decap uint32 `protobuf:"varint,41,opt,name=lisp_v4_decap,json=lispV4Decap,proto3" json:"lisp_v4_decap,omitempty"`
	// Lisp IPv6 encapped packets
	LispV6Encap uint32 `protobuf:"varint,42,opt,name=lisp_v6_encap,json=lispV6Encap,proto3" json:"lisp_v6_encap,omitempty"`
	// Lisp IPv6 decapped packets
	LispV6Decap uint32 `protobuf:"varint,43,opt,name=lisp_v6_decap,json=lispV6Decap,proto3" json:"lisp_v6_decap,omitempty"`
	// Lisp encap errors
	LispEncapError uint32 `protobuf:"varint,44,opt,name=lisp_encap_error,json=lispEncapError,proto3" json:"lisp_encap_error,omitempty"`
	// Lisp decap errors
	LispDecapError       uint32   `protobuf:"varint,45,opt,name=lisp_decap_error,json=lispDecapError,proto3" json:"lisp_decap_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IoTraffic) Reset()         { *m = Ipv4IoTraffic{} }
func (m *Ipv4IoTraffic) String() string { return proto.CompactTextString(m) }
func (*Ipv4IoTraffic) ProtoMessage()    {}
func (*Ipv4IoTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_ip_traffic_268f01bbe9b94ef2, []int{2}
}
func (m *Ipv4IoTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IoTraffic.Unmarshal(m, b)
}
func (m *Ipv4IoTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IoTraffic.Marshal(b, m, deterministic)
}
func (dst *Ipv4IoTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IoTraffic.Merge(dst, src)
}
func (m *Ipv4IoTraffic) XXX_Size() int {
	return xxx_messageInfo_Ipv4IoTraffic.Size(m)
}
func (m *Ipv4IoTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IoTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IoTraffic proto.InternalMessageInfo

func (m *Ipv4IoTraffic) GetInputPackets() uint32 {
	if m != nil {
		return m.InputPackets
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *Ipv4IoTraffic) GetFormatErrors() uint32 {
	if m != nil {
		return m.FormatErrors
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadHopCount() uint32 {
	if m != nil {
		return m.BadHopCount
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadSourceAddress() uint32 {
	if m != nil {
		return m.BadSourceAddress
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadHeader() uint32 {
	if m != nil {
		return m.BadHeader
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoProtocol() uint32 {
	if m != nil {
		return m.NoProtocol
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoGateway() uint32 {
	if m != nil {
		return m.NoGateway
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleInput() uint32 {
	if m != nil {
		return m.ReassembleInput
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembled() uint32 {
	if m != nil {
		return m.Reassembled
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleTimeout() uint32 {
	if m != nil {
		return m.ReassembleTimeout
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleMaxDrop() uint32 {
	if m != nil {
		return m.ReassembleMaxDrop
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleFailed() uint32 {
	if m != nil {
		return m.ReassembleFailed
	}
	return 0
}

func (m *Ipv4IoTraffic) GetOptionsPresent() uint32 {
	if m != nil {
		return m.OptionsPresent
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadOption() uint32 {
	if m != nil {
		return m.BadOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetUnknownOption() uint32 {
	if m != nil {
		return m.UnknownOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadSecurityOption() uint32 {
	if m != nil {
		return m.BadSecurityOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBasicSecurityOption() uint32 {
	if m != nil {
		return m.BasicSecurityOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetExtendedSecurityOption() uint32 {
	if m != nil {
		return m.ExtendedSecurityOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetCipsoOption() uint32 {
	if m != nil {
		return m.CipsoOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetStrictSourceRouteOption() uint32 {
	if m != nil {
		return m.StrictSourceRouteOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLooseSourceRouteOption() uint32 {
	if m != nil {
		return m.LooseSourceRouteOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetRecordRouteOption() uint32 {
	if m != nil {
		return m.RecordRouteOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetSidOption() uint32 {
	if m != nil {
		return m.SidOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetTimestampOption() uint32 {
	if m != nil {
		return m.TimestampOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetRouterAlertOption() uint32 {
	if m != nil {
		return m.RouterAlertOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoopOption() uint32 {
	if m != nil {
		return m.NoopOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetEndOption() uint32 {
	if m != nil {
		return m.EndOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketsOutput() uint32 {
	if m != nil {
		return m.PacketsOutput
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketsForwarded() uint32 {
	if m != nil {
		return m.PacketsForwarded
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketsFragmented() uint32 {
	if m != nil {
		return m.PacketsFragmented
	}
	return 0
}

func (m *Ipv4IoTraffic) GetFragmentCount() uint32 {
	if m != nil {
		return m.FragmentCount
	}
	return 0
}

func (m *Ipv4IoTraffic) GetEncapsulationFailed() uint32 {
	if m != nil {
		return m.EncapsulationFailed
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoRouter() uint32 {
	if m != nil {
		return m.NoRouter
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketTooBig() uint32 {
	if m != nil {
		return m.PacketTooBig
	}
	return 0
}

func (m *Ipv4IoTraffic) GetMulticastIn() uint32 {
	if m != nil {
		return m.MulticastIn
	}
	return 0
}

func (m *Ipv4IoTraffic) GetMulticastOut() uint32 {
	if m != nil {
		return m.MulticastOut
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBroadcastIn() uint32 {
	if m != nil {
		return m.BroadcastIn
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBroadcastOut() uint32 {
	if m != nil {
		return m.BroadcastOut
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV4Encap() uint32 {
	if m != nil {
		return m.LispV4Encap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV4Decap() uint32 {
	if m != nil {
		return m.LispV4Decap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV6Encap() uint32 {
	if m != nil {
		return m.LispV6Encap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV6Decap() uint32 {
	if m != nil {
		return m.LispV6Decap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispEncapError() uint32 {
	if m != nil {
		return m.LispEncapError
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispDecapError() uint32 {
	if m != nil {
		return m.LispDecapError
	}
	return 0
}

// ICMP Traffic Information
type Ipv4IoIcmpTraffic struct {
	// ICMP Received
	Received uint32 `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
	// ICMP Checksum Errors
	ChecksumError uint32 `protobuf:"varint,2,opt,name=checksum_error,json=checksumError,proto3" json:"checksum_error,omitempty"`
	// ICMP Unknown
	Unknown uint32 `protobuf:"varint,3,opt,name=unknown,proto3" json:"unknown,omitempty"`
	// ICMP Transmitted
	Output uint32 `protobuf:"varint,4,opt,name=output,proto3" json:"output,omitempty"`
	// ICMP Admin Unreachable Sent
	AdminUnreachableSent uint32 `protobuf:"varint,5,opt,name=admin_unreachable_sent,json=adminUnreachableSent,proto3" json:"admin_unreachable_sent,omitempty"`
	// ICMP Network Unreachable Sent
	NetworkUnreachableSent uint32 `protobuf:"varint,6,opt,name=network_unreachable_sent,json=networkUnreachableSent,proto3" json:"network_unreachable_sent,omitempty"`
	// ICMP Host Unreachable Sent
	HostUnreachableSent uint32 `protobuf:"varint,7,opt,name=host_unreachable_sent,json=hostUnreachableSent,proto3" json:"host_unreachable_sent,omitempty"`
	// ICMP Protocol Unreachable Sent
	ProtocolUnreachableSent uint32 `protobuf:"varint,8,opt,name=protocol_unreachable_sent,json=protocolUnreachableSent,proto3" json:"protocol_unreachable_sent,omitempty"`
	// ICMP Port Unreachable Sent
	PortUnreachableSent uint32 `protobuf:"varint,9,opt,name=port_unreachable_sent,json=portUnreachableSent,proto3" json:"port_unreachable_sent,omitempty"`
	// ICMP Fragment Unreachable Sent
	FragmentUnreachableSent uint32 `protobuf:"varint,10,opt,name=fragment_unreachable_sent,json=fragmentUnreachableSent,proto3" json:"fragment_unreachable_sent,omitempty"`
	// ICMP Admin Unreachable Received
	AdminUnreachableReceived uint32 `protobuf:"varint,11,opt,name=admin_unreachable_received,json=adminUnreachableReceived,proto3" json:"admin_unreachable_received,omitempty"`
	// ICMP Network Unreachable Received
	NetworkUnreachableReceived uint32 `protobuf:"varint,12,opt,name=network_unreachable_received,json=networkUnreachableReceived,proto3" json:"network_unreachable_received,omitempty"`
	// ICMP Host Unreachable Received
	HostUnreachableReceived uint32 `protobuf:"varint,13,opt,name=host_unreachable_received,json=hostUnreachableReceived,proto3" json:"host_unreachable_received,omitempty"`
	// ICMP Protocol Unreachable Received
	ProtocolUnreachableReceived uint32 `protobuf:"varint,14,opt,name=protocol_unreachable_received,json=protocolUnreachableReceived,proto3" json:"protocol_unreachable_received,omitempty"`
	// ICMP Port Unreachable Received
	PortUnreachableReceived uint32 `protobuf:"varint,15,opt,name=port_unreachable_received,json=portUnreachableReceived,proto3" json:"port_unreachable_received,omitempty"`
	// ICMP Fragment Unreachable Received
	FragmentUnreachableReceived uint32 `protobuf:"varint,16,opt,name=fragment_unreachable_received,json=fragmentUnreachableReceived,proto3" json:"fragment_unreachable_received,omitempty"`
	// ICMP Hopcount Sent
	HopcountSent uint32 `protobuf:"varint,17,opt,name=hopcount_sent,json=hopcountSent,proto3" json:"hopcount_sent,omitempty"`
	// ICMP Reassembly Sent
	ReassemblySent uint32 `protobuf:"varint,18,opt,name=reassembly_sent,json=reassemblySent,proto3" json:"reassembly_sent,omitempty"`
	// ICMP Hopcount Received
	HopcountReceived uint32 `protobuf:"varint,19,opt,name=hopcount_received,json=hopcountReceived,proto3" json:"hopcount_received,omitempty"`
	// ICMP Reassembly Received
	ReasseblyReceived uint32 `protobuf:"varint,20,opt,name=reassebly_received,json=reasseblyReceived,proto3" json:"reassebly_received,omitempty"`
	// ICMP Parameter Error Received
	ParamErrorReceived uint32 `protobuf:"varint,21,opt,name=param_error_received,json=paramErrorReceived,proto3" json:"param_error_received,omitempty"`
	// ICMP Parameter Error Sent
	ParamErrorSend uint32 `protobuf:"varint,22,opt,name=param_error_send,json=paramErrorSend,proto3" json:"param_error_send,omitempty"`
	// ICMP Echo Request Sent
	EchoRequestSent uint32 `protobuf:"varint,23,opt,name=echo_request_sent,json=echoRequestSent,proto3" json:"echo_request_sent,omitempty"`
	// ICMP Echo Request Sent
	EchoRequestReceived uint32 `protobuf:"varint,24,opt,name=echo_request_received,json=echoRequestReceived,proto3" json:"echo_request_received,omitempty"`
	// ICMP Echo Reply Sent
	EchoReplySent uint32 `protobuf:"varint,25,opt,name=echo_reply_sent,json=echoReplySent,proto3" json:"echo_reply_sent,omitempty"`
	// ICMP Echo Reply Received
	EchoReplyReceived uint32 `protobuf:"varint,26,opt,name=echo_reply_received,json=echoReplyReceived,proto3" json:"echo_reply_received,omitempty"`
	// ICMP Mask Sent
	MaskRequestSent uint32 `protobuf:"varint,27,opt,name=mask_request_sent,json=maskRequestSent,proto3" json:"mask_request_sent,omitempty"`
	// ICMP Mask Received
	MaskRequestReceived uint32 `protobuf:"varint,28,opt,name=mask_request_received,json=maskRequestReceived,proto3" json:"mask_request_received,omitempty"`
	// ICMP Mask Sent
	MaskReplySent uint32 `protobuf:"varint,29,opt,name=mask_reply_sent,json=maskReplySent,proto3" json:"mask_reply_sent,omitempty"`
	// ICMP Mask Received
	MaskReplyReceived uint32 `protobuf:"varint,30,opt,name=mask_reply_received,json=maskReplyReceived,proto3" json:"mask_reply_received,omitempty"`
	// ICMP Source Quench
	SourceQuenchReceived uint32 `protobuf:"varint,31,opt,name=source_quench_received,json=sourceQuenchReceived,proto3" json:"source_quench_received,omitempty"`
	// ICMP Redirect Received
	RedirectReceived uint32 `protobuf:"varint,32,opt,name=redirect_received,json=redirectReceived,proto3" json:"redirect_received,omitempty"`
	// ICMP Redirect Sent
	RedirectSend uint32 `protobuf:"varint,33,opt,name=redirect_send,json=redirectSend,proto3" json:"redirect_send,omitempty"`
	// ICMP Timestamp Received
	TimestampReceived uint32 `protobuf:"varint,34,opt,name=timestamp_received,json=timestampReceived,proto3" json:"timestamp_received,omitempty"`
	// ICMP Timestamp Reply Received
	TimestampReplyReceived uint32 `protobuf:"varint,35,opt,name=timestamp_reply_received,json=timestampReplyReceived,proto3" json:"timestamp_reply_received,omitempty"`
	// ICMP Router Advertisement Received
	RouterAdvertReceived uint32 `protobuf:"varint,36,opt,name=router_advert_received,json=routerAdvertReceived,proto3" json:"router_advert_received,omitempty"`
	// ICMP Router Solicited Received
	RouterSolicitReceived uint32   `protobuf:"varint,37,opt,name=router_solicit_received,json=routerSolicitReceived,proto3" json:"router_solicit_received,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Ipv4IoIcmpTraffic) Reset()         { *m = Ipv4IoIcmpTraffic{} }
func (m *Ipv4IoIcmpTraffic) String() string { return proto.CompactTextString(m) }
func (*Ipv4IoIcmpTraffic) ProtoMessage()    {}
func (*Ipv4IoIcmpTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_ip_traffic_268f01bbe9b94ef2, []int{3}
}
func (m *Ipv4IoIcmpTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IoIcmpTraffic.Unmarshal(m, b)
}
func (m *Ipv4IoIcmpTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IoIcmpTraffic.Marshal(b, m, deterministic)
}
func (dst *Ipv4IoIcmpTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IoIcmpTraffic.Merge(dst, src)
}
func (m *Ipv4IoIcmpTraffic) XXX_Size() int {
	return xxx_messageInfo_Ipv4IoIcmpTraffic.Size(m)
}
func (m *Ipv4IoIcmpTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IoIcmpTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IoIcmpTraffic proto.InternalMessageInfo

func (m *Ipv4IoIcmpTraffic) GetReceived() uint32 {
	if m != nil {
		return m.Received
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetChecksumError() uint32 {
	if m != nil {
		return m.ChecksumError
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetUnknown() uint32 {
	if m != nil {
		return m.Unknown
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetOutput() uint32 {
	if m != nil {
		return m.Output
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetAdminUnreachableSent() uint32 {
	if m != nil {
		return m.AdminUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetNetworkUnreachableSent() uint32 {
	if m != nil {
		return m.NetworkUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHostUnreachableSent() uint32 {
	if m != nil {
		return m.HostUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetProtocolUnreachableSent() uint32 {
	if m != nil {
		return m.ProtocolUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetPortUnreachableSent() uint32 {
	if m != nil {
		return m.PortUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetFragmentUnreachableSent() uint32 {
	if m != nil {
		return m.FragmentUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetAdminUnreachableReceived() uint32 {
	if m != nil {
		return m.AdminUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetNetworkUnreachableReceived() uint32 {
	if m != nil {
		return m.NetworkUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHostUnreachableReceived() uint32 {
	if m != nil {
		return m.HostUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetProtocolUnreachableReceived() uint32 {
	if m != nil {
		return m.ProtocolUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetPortUnreachableReceived() uint32 {
	if m != nil {
		return m.PortUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetFragmentUnreachableReceived() uint32 {
	if m != nil {
		return m.FragmentUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHopcountSent() uint32 {
	if m != nil {
		return m.HopcountSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetReassemblySent() uint32 {
	if m != nil {
		return m.ReassemblySent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHopcountReceived() uint32 {
	if m != nil {
		return m.HopcountReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetReasseblyReceived() uint32 {
	if m != nil {
		return m.ReasseblyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetParamErrorReceived() uint32 {
	if m != nil {
		return m.ParamErrorReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetParamErrorSend() uint32 {
	if m != nil {
		return m.ParamErrorSend
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoRequestSent() uint32 {
	if m != nil {
		return m.EchoRequestSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoRequestReceived() uint32 {
	if m != nil {
		return m.EchoRequestReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoReplySent() uint32 {
	if m != nil {
		return m.EchoReplySent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoReplyReceived() uint32 {
	if m != nil {
		return m.EchoReplyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskRequestSent() uint32 {
	if m != nil {
		return m.MaskRequestSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskRequestReceived() uint32 {
	if m != nil {
		return m.MaskRequestReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskReplySent() uint32 {
	if m != nil {
		return m.MaskReplySent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskReplyReceived() uint32 {
	if m != nil {
		return m.MaskReplyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetSourceQuenchReceived() uint32 {
	if m != nil {
		return m.SourceQuenchReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRedirectReceived() uint32 {
	if m != nil {
		return m.RedirectReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRedirectSend() uint32 {
	if m != nil {
		return m.RedirectSend
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetTimestampReceived() uint32 {
	if m != nil {
		return m.TimestampReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetTimestampReplyReceived() uint32 {
	if m != nil {
		return m.TimestampReplyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRouterAdvertReceived() uint32 {
	if m != nil {
		return m.RouterAdvertReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRouterSolicitReceived() uint32 {
	if m != nil {
		return m.RouterSolicitReceived
	}
	return 0
}

func init() {
	proto.RegisterType((*IpTraffic_KEYS)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ip_traffic_KEYS")
	proto.RegisterType((*IpTraffic)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ip_traffic")
	proto.RegisterType((*Ipv4IoTraffic)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ipv4_io_traffic")
	proto.RegisterType((*Ipv4IoIcmpTraffic)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ipv4_io_icmp_traffic")
}

func init() { proto.RegisterFile("ip_traffic.proto", fileDescriptor_ip_traffic_268f01bbe9b94ef2) }

var fileDescriptor_ip_traffic_268f01bbe9b94ef2 = []byte{
	// 1479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5b, 0x4f, 0x1c, 0x39,
	0x16, 0x16, 0x7b, 0x49, 0xe0, 0x40, 0xd3, 0x4d, 0x75, 0x03, 0x05, 0x84, 0x0d, 0x81, 0x5c, 0xc8,
	0xad, 0xb5, 0x4b, 0x10, 0xca, 0x66, 0xf7, 0x61, 0x93, 0x4d, 0xb2, 0x1b, 0x8d, 0x66, 0x92, 0x81,
	0x64, 0xa4, 0x99, 0x17, 0xcb, 0x5d, 0x65, 0xc0, 0xa2, 0xbb, 0x5c, 0xb1, 0x5d, 0x5c, 0xfe, 0xd5,
	0x3c, 0xcd, 0xbf, 0x1b, 0x69, 0xe4, 0x73, 0x6c, 0x57, 0xf5, 0xe5, 0x31, 0x2f, 0x2d, 0xf5, 0x77,
	0xbe, 0xef, 0xf3, 0xa9, 0xe3, 0xcb, 0x39, 0xd0, 0x91, 0x25, 0xb3, 0x9a, 0x9f, 0x9e, 0xca, 0xac,
	0x5f, 0x6a, 0x65, 0x55, 0xf2, 0x36, 0x93, 0x26, 0x53, 0x4c, 0x2a, 0xc3, 0xae, 0x35, 0x93, 0xe5,
	0xe5, 0x21, 0x93, 0x8a, 0xa9, 0x52, 0xe8, 0x3e, 0xfe, 0x29, 0x84, 0xbd, 0x52, 0xfa, 0xa2, 0x5f,
	0xa8, 0x5c, 0x18, 0xfc, 0xed, 0x1b, 0xcb, 0xad, 0x34, 0x56, 0x66, 0xa6, 0xef, 0xbd, 0x76, 0xfb,
	0xd0, 0xae, 0x9d, 0xd9, 0x77, 0xef, 0x7e, 0x3e, 0x49, 0xb6, 0x60, 0xc1, 0xb1, 0x59, 0xc1, 0x47,
	0x22, 0x9d, 0xdb, 0x99, 0xdb, 0x5f, 0x38, 0x9e, 0x77, 0xc0, 0x0f, 0x7c, 0x24, 0x76, 0x7f, 0x9f,
	0x03, 0xa8, 0x05, 0x89, 0x75, 0xff, 0x2e, 0x0f, 0x99, 0x73, 0x36, 0xe9, 0xc1, 0xce, 0xdc, 0xfe,
	0xe2, 0xc1, 0x97, 0xfe, 0xb7, 0xc8, 0xac, 0x1f, 0x74, 0xfe, 0xff, 0xf1, 0x82, 0x03, 0x4e, 0xdc,
	0x3a, 0xc9, 0x0d, 0x80, 0xcc, 0x46, 0xa5, 0x5f, 0xf5, 0x05, 0xae, 0xfa, 0xcb, 0xb7, 0x5d, 0x15,
	0xfd, 0xeb, 0xa5, 0xb3, 0x51, 0x89, 0x4b, 0xef, 0xfe, 0xb6, 0xec, 0x0a, 0x36, 0x96, 0x59, 0xb2,
	0x07, 0x2d, 0x59, 0x94, 0x95, 0x65, 0x25, 0xcf, 0x2e, 0x84, 0x35, 0x58, 0xb4, 0xd6, 0xf1, 0x12,
	0x82, 0x9f, 0x08, 0x4b, 0x1e, 0x43, 0x47, 0x8b, 0x4c, 0xc8, 0x4b, 0x91, 0x47, 0xde, 0x9f, 0x90,
	0xd7, 0x0e, 0x78, 0xa0, 0xee, 0x41, 0xeb, 0x54, 0xe9, 0x11, 0xb7, 0x4c, 0x68, 0xad, 0xb4, 0x49,
	0xff, 0x4c, 0x7e, 0x04, 0xbe, 0x43, 0x2c, 0xd9, 0x85, 0xd6, 0x80, 0xe7, 0xec, 0x5c, 0x95, 0x2c,
	0x53, 0x55, 0x61, 0xd3, 0xbf, 0x20, 0x69, 0x71, 0xc0, 0xf3, 0xff, 0xab, 0xf2, 0xbf, 0x0e, 0x4a,
	0x9e, 0x41, 0xe2, 0x38, 0x46, 0x55, 0x3a, 0x13, 0x8c, 0xe7, 0xb9, 0x16, 0xc6, 0xa4, 0x7f, 0x45,
	0x62, 0x67, 0xc0, 0xf3, 0x13, 0x0c, 0xbc, 0x26, 0x3c, 0xd9, 0x06, 0x40, 0x47, 0xc1, 0x73, 0xa1,
	0xd3, 0x5b, 0xc8, 0x5a, 0x70, 0x76, 0x08, 0x24, 0x77, 0x61, 0xb1, 0x50, 0x0c, 0xcf, 0x5e, 0xa6,
	0x86, 0xe9, 0x6d, 0x8c, 0x43, 0xa1, 0x3e, 0x79, 0xc4, 0xe9, 0x0b, 0xc5, 0xce, 0xb8, 0x15, 0x57,
	0xfc, 0x26, 0x9d, 0x27, 0x7d, 0xa1, 0xfe, 0x47, 0x00, 0x15, 0x80, 0x1b, 0x23, 0x46, 0x83, 0xa1,
	0x60, 0x58, 0x9b, 0x74, 0x21, 0x14, 0x20, 0xe0, 0x1f, 0x1c, 0x9c, 0xec, 0xc0, 0x62, 0x0d, 0xe5,
	0x29, 0xd0, 0x97, 0x35, 0xa0, 0xe4, 0x39, 0x24, 0x0d, 0x33, 0x2b, 0x47, 0x42, 0x55, 0x36, 0x5d,
	0x44, 0xe2, 0x4a, 0x1d, 0xf9, 0x4c, 0x81, 0xa4, 0x0f, 0xdd, 0x06, 0x7d, 0xc4, 0xaf, 0x59, 0xae,
	0x55, 0x99, 0x2e, 0x4d, 0xf2, 0xbf, 0xe7, 0xd7, 0x6f, 0xb5, 0x2a, 0x93, 0xa7, 0xd0, 0x00, 0xd9,
	0x29, 0x97, 0x2e, 0x8d, 0x16, 0xd5, 0xad, 0x0e, 0xbc, 0x47, 0x3c, 0x79, 0x04, 0x6d, 0x55, 0x5a,
	0xa9, 0x0a, 0xc3, 0x4a, 0x2d, 0x8c, 0x28, 0x6c, 0xba, 0x8c, 0xd4, 0x65, 0x0f, 0x7f, 0x22, 0x34,
	0x14, 0x98, 0xd0, 0xb4, 0x1d, 0x0b, 0xfc, 0x11, 0x81, 0xe4, 0x01, 0x2c, 0x57, 0xc5, 0x45, 0xa1,
	0xae, 0x8a, 0x40, 0xe9, 0x20, 0xa5, 0xe5, 0x51, 0x4f, 0xeb, 0x43, 0x17, 0x37, 0x55, 0x64, 0x95,
	0x96, 0xf6, 0x26, 0x70, 0x57, 0xe8, 0x5b, 0xdc, 0xae, 0xfa, 0x88, 0xe7, 0x1f, 0xc0, 0xea, 0x80,
	0x1b, 0x99, 0x4d, 0x29, 0x12, 0x54, 0x74, 0x31, 0x38, 0xa1, 0x79, 0x09, 0xa9, 0xb8, 0xb6, 0xa2,
	0xc8, 0xc5, 0xf4, 0x42, 0x5d, 0x94, 0xad, 0x85, 0xf8, 0x84, 0xf2, 0x1e, 0x2c, 0x65, 0xb2, 0x34,
	0x2a, 0xb0, 0x7b, 0xb4, 0x77, 0x88, 0x79, 0xca, 0xbf, 0x60, 0xd3, 0x58, 0x2d, 0x33, 0x1b, 0x0e,
	0xa6, 0x56, 0x95, 0x15, 0x41, 0xb0, 0x8a, 0x82, 0x75, 0x62, 0xd0, 0x01, 0x3d, 0x76, 0x71, 0x2f,
	0xfe, 0x27, 0x6c, 0x0c, 0x95, 0x32, 0x62, 0xa6, 0x76, 0x8d, 0x52, 0x43, 0xc2, 0xb4, 0x14, 0x0f,
	0x41, 0xa6, 0x74, 0x3e, 0x2e, 0x5a, 0x0f, 0x87, 0xc0, 0x85, 0x9a, 0xfc, 0x6d, 0x00, 0x23, 0xe3,
	0x76, 0xa5, 0xb4, 0x5d, 0x46, 0x86, 0xed, 0x7a, 0x0c, 0x1d, 0x77, 0xee, 0x8c, 0xe5, 0xa3, 0x32,
	0x90, 0x36, 0xe8, 0x3c, 0x47, 0xbc, 0xb1, 0xb2, 0x33, 0xd6, 0x8c, 0x0f, 0x85, 0xb6, 0x81, 0xbd,
	0xe9, 0x57, 0xc6, 0xd0, 0x6b, 0x17, 0xf1, 0x7c, 0xbc, 0x6a, 0x2a, 0xba, 0x6e, 0x85, 0xab, 0xa6,
	0xca, 0x3a, 0x35, 0x51, 0xc4, 0xd4, 0xee, 0x50, 0x6a, 0xa2, 0x68, 0x9c, 0x24, 0xff, 0xc4, 0x30,
	0x55, 0x59, 0x77, 0xd1, 0xb6, 0xe9, 0x24, 0x79, 0xf4, 0x23, 0x82, 0xee, 0x94, 0x07, 0xda, 0xa9,
	0xd2, 0x57, 0x5c, 0xe7, 0x22, 0x4f, 0xff, 0x46, 0xa7, 0xdc, 0x07, 0xde, 0x07, 0xdc, 0xdd, 0xb8,
	0x48, 0xd6, 0xfc, 0x6c, 0x24, 0x0a, 0x2b, 0xf2, 0xf4, 0x2e, 0x7d, 0x42, 0x60, 0xc7, 0x80, 0x4b,
	0x21, 0xd0, 0xfc, 0xfb, 0xb4, 0x43, 0x29, 0x04, 0x94, 0x5e, 0xa8, 0x7f, 0x40, 0x4f, 0x14, 0x19,
	0x2f, 0x4d, 0x35, 0xe4, 0x2e, 0xf5, 0x70, 0xd7, 0xee, 0xd1, 0xd9, 0x1c, 0x8b, 0xf9, 0xeb, 0x86,
	0xed, 0x89, 0xb6, 0x50, 0xa7, 0xbb, 0xc8, 0x9b, 0x2f, 0x14, 0x6e, 0x9c, 0x4e, 0xee, 0x87, 0x2f,
	0x67, 0x56, 0x29, 0x36, 0x90, 0x67, 0xe9, 0x1e, 0xbd, 0x9d, 0x84, 0x7e, 0x56, 0xea, 0x8d, 0x3c,
	0x73, 0x87, 0x74, 0x54, 0x0d, 0xad, 0xcc, 0xb8, 0xb1, 0x4c, 0x16, 0xe9, 0x7d, 0x3a, 0xa4, 0x11,
	0xfb, 0x50, 0xb8, 0x37, 0xb8, 0xa6, 0xb8, 0xb7, 0xe5, 0x01, 0xf9, 0x44, 0xf0, 0x63, 0x65, 0x9d,
	0xcf, 0x40, 0x2b, 0x9e, 0x07, 0x9f, 0x87, 0xfe, 0x09, 0x0e, 0x18, 0xf9, 0xd4, 0x14, 0xe7, 0xf3,
	0x88, 0x7c, 0x22, 0xe8, 0x7c, 0x76, 0xa1, 0x35, 0x94, 0xa6, 0x64, 0x97, 0x87, 0x0c, 0xbf, 0x38,
	0xdd, 0x27, 0x23, 0x07, 0xfe, 0x74, 0xf8, 0xce, 0x41, 0x4d, 0x4e, 0x2e, 0x1c, 0xe7, 0x71, 0x93,
	0xf3, 0x56, 0x8c, 0x71, 0x8e, 0xbc, 0xcf, 0x93, 0x06, 0xe7, 0x68, 0xc2, 0xe7, 0xc8, 0xfb, 0x3c,
	0x6d, 0x72, 0xc8, 0x67, 0x1f, 0x3a, 0xc8, 0x41, 0x13, 0x6a, 0x42, 0xe9, 0x33, 0x7a, 0xd2, 0x1c,
	0x8e, 0x46, 0xd8, 0x86, 0x22, 0x13, 0xad, 0x3c, 0xf3, 0x79, 0xcd, 0x44, 0x3b, 0x64, 0xee, 0xfe,
	0xba, 0x0c, 0xbd, 0x59, 0xcd, 0x35, 0xd9, 0x84, 0xf9, 0xd0, 0x00, 0x7d, 0xe3, 0x8c, 0xff, 0xdd,
	0x29, 0xca, 0xce, 0x45, 0x76, 0x61, 0xaa, 0x91, 0x37, 0xa7, 0x96, 0xd9, 0x0a, 0x28, 0x65, 0x91,
	0xc2, 0x6d, 0xff, 0x46, 0xfa, 0x56, 0x19, 0xfe, 0x26, 0x6b, 0x70, 0xcb, 0xdf, 0x00, 0x6a, 0x8f,
	0xfe, 0x5f, 0x72, 0x08, 0x6b, 0x3c, 0x1f, 0xc9, 0x82, 0x55, 0x85, 0x16, 0x3c, 0x3b, 0xe7, 0xee,
	0x9d, 0xc7, 0xa7, 0x9b, 0xba, 0x63, 0x0f, 0xa3, 0x5f, 0xea, 0xe0, 0x89, 0x7b, 0xc0, 0x5f, 0x42,
	0xea, 0xa7, 0x88, 0x69, 0x1d, 0xf5, 0xcb, 0x35, 0x1f, 0x9f, 0x54, 0x1e, 0xc0, 0xea, 0xb9, 0x32,
	0x76, 0x5a, 0x46, 0x6d, 0xb4, 0xeb, 0x82, 0x93, 0x9a, 0x57, 0xb0, 0x11, 0xba, 0xed, 0xb4, 0x8e,
	0xda, 0xeb, 0x7a, 0x20, 0xcc, 0x58, 0xaf, 0x54, 0x7a, 0xc6, 0x7a, 0xd4, 0x71, 0xbb, 0x2e, 0x38,
	0x63, 0xbd, 0x78, 0x65, 0xa7, 0x74, 0xd4, 0x83, 0xd7, 0x03, 0x61, 0x52, 0xfb, 0x6f, 0xd8, 0x9c,
	0xae, 0x67, 0xdc, 0x56, 0xea, 0xcb, 0xe9, 0x64, 0x4d, 0x8f, 0xc3, 0x36, 0xff, 0x07, 0xee, 0xcc,
	0xaa, 0x6b, 0xd4, 0x53, 0x9f, 0xde, 0x9c, 0xae, 0x6d, 0x74, 0x78, 0x05, 0x1b, 0x53, 0xf5, 0x8d,
	0x72, 0x6a, 0xdc, 0xeb, 0x13, 0x35, 0x8e, 0xda, 0x37, 0xb0, 0x3d, 0xb3, 0xce, 0x51, 0x4f, 0xdd,
	0x7c, 0x6b, 0x46, 0xad, 0x9b, 0xeb, 0x4f, 0xd5, 0x3b, 0xea, 0xdb, 0x7e, 0xaf, 0xc6, 0x6b, 0xde,
	0x5c, 0x7f, 0x66, 0xdd, 0xa3, 0x9e, 0xc6, 0x80, 0xad, 0x19, 0xb5, 0x8f, 0x1e, 0x7b, 0xd0, 0x3a,
	0x57, 0x25, 0x3e, 0xb4, 0xb4, 0x5f, 0x34, 0x0e, 0x2c, 0x05, 0x10, 0x37, 0xe9, 0x11, 0xd4, 0x93,
	0xd6, 0x0d, 0xd1, 0x68, 0x06, 0x58, 0xae, 0x61, 0x24, 0x3e, 0x85, 0x95, 0xe8, 0x16, 0xb3, 0xa0,
	0xbe, 0xdf, 0x09, 0x81, 0xb8, 0x74, 0x1c, 0xc5, 0x9c, 0x69, 0x64, 0xf7, 0x9a, 0xa3, 0xd5, 0x60,
	0x78, 0x13, 0xe9, 0x7f, 0x87, 0x5e, 0xc9, 0x35, 0xf7, 0xf7, 0xb9, 0x16, 0x50, 0xdf, 0x4f, 0x30,
	0x86, 0xb7, 0x3a, 0x2a, 0xf6, 0xa1, 0xd3, 0x54, 0x18, 0x51, 0xe4, 0xbe, 0xd3, 0x2f, 0xd7, 0xec,
	0x13, 0x51, 0xe4, 0xc9, 0x13, 0x58, 0x11, 0xd9, 0xb9, 0x62, 0x5a, 0x7c, 0xad, 0x84, 0xf1, 0x95,
	0xa0, 0xfe, 0xde, 0x76, 0x81, 0x63, 0xc2, 0xc3, 0x0d, 0x19, 0xe3, 0xc6, 0x44, 0x52, 0xdf, 0x7a,
	0x6a, 0x7e, 0xcc, 0xe4, 0x21, 0xb4, 0xbd, 0xa6, 0x0c, 0x05, 0xa4, 0x8e, 0xdf, 0x22, 0x76, 0xe9,
	0xeb, 0xd7, 0x87, 0x6e, 0x83, 0x17, 0x9d, 0x7d, 0xbf, 0x8f, 0xdc, 0xe8, 0xfb, 0x04, 0x56, 0x46,
	0xdc, 0x5c, 0x8c, 0xe7, 0x4d, 0x5d, 0xbf, 0xed, 0x02, 0x13, 0x79, 0x8f, 0x71, 0xa3, 0x3b, 0x4d,
	0x01, 0xdd, 0x06, 0xbf, 0x99, 0xb7, 0xd7, 0xc4, 0xbc, 0xfd, 0x40, 0x40, 0xec, 0x46, 0xde, 0x0d,
	0x5e, 0x74, 0xa6, 0x91, 0x60, 0x25, 0x72, 0xa3, 0xef, 0x21, 0xac, 0xf9, 0x31, 0xec, 0x6b, 0x25,
	0x8a, 0xec, 0xbc, 0x96, 0xd0, 0x5c, 0xd0, 0xa3, 0xe8, 0x8f, 0x18, 0x8c, 0x2a, 0x1c, 0xae, 0x73,
	0xa9, 0x45, 0xd6, 0xc8, 0x7e, 0x27, 0x0c, 0xd7, 0x14, 0x68, 0x1e, 0xec, 0x48, 0xc6, 0x9d, 0xa7,
	0xc9, 0x60, 0x29, 0x80, 0xb8, 0xef, 0xcf, 0x21, 0xa9, 0x47, 0xb1, 0x68, 0x49, 0xb3, 0xc1, 0x4a,
	0x8c, 0x44, 0xcf, 0x97, 0x90, 0x36, 0xe9, 0x63, 0xdf, 0x4a, 0xe3, 0xc2, 0x5a, 0x43, 0x34, 0xf1,
	0xc1, 0x61, 0x90, 0xcb, 0x2f, 0xdd, 0x24, 0x17, 0x75, 0x34, 0x42, 0xf4, 0xfc, 0x2c, 0x87, 0xc1,
	0xa8, 0x3a, 0x82, 0x75, 0xaf, 0x32, 0x6a, 0x28, 0x33, 0xd9, 0x90, 0xd1, 0x54, 0xb1, 0x4a, 0xe1,
	0x13, 0x8a, 0x06, 0xdd, 0xe0, 0x16, 0xbe, 0x38, 0x2f, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x68,
	0x19, 0x6b, 0x11, 0xfc, 0x0f, 0x00, 0x00,
}
