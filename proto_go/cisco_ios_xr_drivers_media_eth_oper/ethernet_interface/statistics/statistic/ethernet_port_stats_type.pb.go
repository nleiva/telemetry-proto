// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ethernet_port_stats_type.proto

package cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_statistics_statistic

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

// Port stats counters
type EthernetPortStatsType_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthernetPortStatsType_KEYS) Reset()         { *m = EthernetPortStatsType_KEYS{} }
func (m *EthernetPortStatsType_KEYS) String() string { return proto.CompactTextString(m) }
func (*EthernetPortStatsType_KEYS) ProtoMessage()    {}
func (*EthernetPortStatsType_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982f68fad60746a, []int{0}
}

func (m *EthernetPortStatsType_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthernetPortStatsType_KEYS.Unmarshal(m, b)
}
func (m *EthernetPortStatsType_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthernetPortStatsType_KEYS.Marshal(b, m, deterministic)
}
func (m *EthernetPortStatsType_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthernetPortStatsType_KEYS.Merge(m, src)
}
func (m *EthernetPortStatsType_KEYS) XXX_Size() int {
	return xxx_messageInfo_EthernetPortStatsType_KEYS.Size(m)
}
func (m *EthernetPortStatsType_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EthernetPortStatsType_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EthernetPortStatsType_KEYS proto.InternalMessageInfo

func (m *EthernetPortStatsType_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type EthernetPortStatsType struct {
	// Total octets of all frames
	ReceivedTotalBytes uint64 `protobuf:"varint,50,opt,name=received_total_bytes,json=receivedTotalBytes,proto3" json:"received_total_bytes,omitempty"`
	// Total octets of all good frames
	ReceivedGoodBytes uint64 `protobuf:"varint,51,opt,name=received_good_bytes,json=receivedGoodBytes,proto3" json:"received_good_bytes,omitempty"`
	// All frames, good or bad
	ReceivedTotalFrames uint64 `protobuf:"varint,52,opt,name=received_total_frames,json=receivedTotalFrames,proto3" json:"received_total_frames,omitempty"`
	// All 802.1Q frames
	Received8021QFrames uint64 `protobuf:"varint,53,opt,name=received8021_q_frames,json=received8021QFrames,proto3" json:"received8021_q_frames,omitempty"`
	// All pause frames
	ReceivedPauseFrames uint64 `protobuf:"varint,54,opt,name=received_pause_frames,json=receivedPauseFrames,proto3" json:"received_pause_frames,omitempty"`
	// Unsupported MAC Control frames
	ReceivedUnknownOpcodes uint64 `protobuf:"varint,55,opt,name=received_unknown_opcodes,json=receivedUnknownOpcodes,proto3" json:"received_unknown_opcodes,omitempty"`
	// All 64 Octet Frame Count
	ReceivedTotal64OctetFrames uint64 `protobuf:"varint,56,opt,name=received_total64_octet_frames,json=receivedTotal64OctetFrames,proto3" json:"received_total64_octet_frames,omitempty"`
	// All 65-127 Octet Frame Count
	ReceivedTotalOctetFramesFrom65To127 uint64 `protobuf:"varint,57,opt,name=received_total_octet_frames_from65_to127,json=receivedTotalOctetFramesFrom65To127,proto3" json:"received_total_octet_frames_from65_to127,omitempty"`
	// All 128-255 Octet Frame Count
	ReceivedTotalOctetFramesFrom128To255 uint64 `protobuf:"varint,58,opt,name=received_total_octet_frames_from128_to255,json=receivedTotalOctetFramesFrom128To255,proto3" json:"received_total_octet_frames_from128_to255,omitempty"`
	// All 256-511 Octet Frame Count
	ReceivedTotalOctetFramesFrom256To511 uint64 `protobuf:"varint,59,opt,name=received_total_octet_frames_from256_to511,json=receivedTotalOctetFramesFrom256To511,proto3" json:"received_total_octet_frames_from256_to511,omitempty"`
	// All 512-1023 Octet Frame Count
	ReceivedTotalOctetFramesFrom512To1023 uint64 `protobuf:"varint,60,opt,name=received_total_octet_frames_from512_to1023,json=receivedTotalOctetFramesFrom512To1023,proto3" json:"received_total_octet_frames_from512_to1023,omitempty"`
	// All 1024-1518 Octet Frame Count
	ReceivedTotalOctetFramesFrom1024To1518 uint64 `protobuf:"varint,61,opt,name=received_total_octet_frames_from1024_to1518,json=receivedTotalOctetFramesFrom1024To1518,proto3" json:"received_total_octet_frames_from1024_to1518,omitempty"`
	// All > 1518 Octet Frame Count
	ReceivedTotalOctetFramesFrom1519ToMax uint64 `protobuf:"varint,62,opt,name=received_total_octet_frames_from1519_to_max,json=receivedTotalOctetFramesFrom1519ToMax,proto3" json:"received_total_octet_frames_from1519_to_max,omitempty"`
	// Received Good Frames
	ReceivedGoodFrames uint64 `protobuf:"varint,63,opt,name=received_good_frames,json=receivedGoodFrames,proto3" json:"received_good_frames,omitempty"`
	// Received unicast Frames
	ReceivedUnicastFrames uint64 `protobuf:"varint,64,opt,name=received_unicast_frames,json=receivedUnicastFrames,proto3" json:"received_unicast_frames,omitempty"`
	// Received multicast Frames
	ReceivedMulticastFrames uint64 `protobuf:"varint,65,opt,name=received_multicast_frames,json=receivedMulticastFrames,proto3" json:"received_multicast_frames,omitempty"`
	// Received broadcast Frames
	ReceivedBroadcastFrames uint64 `protobuf:"varint,66,opt,name=received_broadcast_frames,json=receivedBroadcastFrames,proto3" json:"received_broadcast_frames,omitempty"`
	// Drops due to buffer overrun
	NumberOfBufferOverrunPacketsDropped uint64 `protobuf:"varint,67,opt,name=number_of_buffer_overrun_packets_dropped,json=numberOfBufferOverrunPacketsDropped,proto3" json:"number_of_buffer_overrun_packets_dropped,omitempty"`
	// Drops due to packet abort
	NumberOfAbortedPacketsDropped uint64 `protobuf:"varint,68,opt,name=number_of_aborted_packets_dropped,json=numberOfAbortedPacketsDropped,proto3" json:"number_of_aborted_packets_dropped,omitempty"`
	// Drops due to invalid VLAN id
	NumberofInvalidVlanIdPacketsDropped uint64 `protobuf:"varint,69,opt,name=numberof_invalid_vlan_id_packets_dropped,json=numberofInvalidVlanIdPacketsDropped,proto3" json:"numberof_invalid_vlan_id_packets_dropped,omitempty"`
	// Drops due to the destination MAC not matching
	InvalidDestMacDropPackets uint64 `protobuf:"varint,70,opt,name=invalid_dest_mac_drop_packets,json=invalidDestMacDropPackets,proto3" json:"invalid_dest_mac_drop_packets,omitempty"`
	// Drops due to the encapsulation or ether type not matching
	InvalidEncapDropPackets uint64 `protobuf:"varint,71,opt,name=invalid_encap_drop_packets,json=invalidEncapDropPackets,proto3" json:"invalid_encap_drop_packets,omitempty"`
	// Any other drops not counted
	NumberOfMiscellaneousPacketsDropped uint64 `protobuf:"varint,72,opt,name=number_of_miscellaneous_packets_dropped,json=numberOfMiscellaneousPacketsDropped,proto3" json:"number_of_miscellaneous_packets_dropped,omitempty"`
	// Good frames > MRU, dropped
	DroppedGiantPacketsGreaterthanMru uint64 `protobuf:"varint,73,opt,name=dropped_giant_packets_greaterthan_mru,json=droppedGiantPacketsGreaterthanMru,proto3" json:"dropped_giant_packets_greaterthan_mru,omitempty"`
	// Good frames < 64 Octet, dropped
	DroppedEtherStatsUndersizePkts uint64 `protobuf:"varint,74,opt,name=dropped_ether_stats_undersize_pkts,json=droppedEtherStatsUndersizePkts,proto3" json:"dropped_ether_stats_undersize_pkts,omitempty"`
	// Bad Frames > MRU, dropped
	DroppedJabbersPacketsGreaterthanMru uint64 `protobuf:"varint,75,opt,name=dropped_jabbers_packets_greaterthan_mru,json=droppedJabbersPacketsGreaterthanMru,proto3" json:"dropped_jabbers_packets_greaterthan_mru,omitempty"`
	// Bad Frames < 64 Octet, dropped
	DroppedEtherStatsFragments uint64 `protobuf:"varint,76,opt,name=dropped_ether_stats_fragments,json=droppedEtherStatsFragments,proto3" json:"dropped_ether_stats_fragments,omitempty"`
	// Frames 64 - MRU with CRC error
	DroppedPacketsWithCrcAlignErrors uint64 `protobuf:"varint,77,opt,name=dropped_packets_with_crc_align_errors,json=droppedPacketsWithCrcAlignErrors,proto3" json:"dropped_packets_with_crc_align_errors,omitempty"`
	// All collision events
	EtherStatsCollisions uint64 `protobuf:"varint,78,opt,name=ether_stats_collisions,json=etherStatsCollisions,proto3" json:"ether_stats_collisions,omitempty"`
	// Symbol errors
	SymbolErrors uint64 `protobuf:"varint,79,opt,name=symbol_errors,json=symbolErrors,proto3" json:"symbol_errors,omitempty"`
	// Any other errors not counted
	DroppedMiscellaneousErrorPackets uint64 `protobuf:"varint,80,opt,name=dropped_miscellaneous_error_packets,json=droppedMiscellaneousErrorPackets,proto3" json:"dropped_miscellaneous_error_packets,omitempty"`
	// RFC2819 etherStatsOversizedPkts
	Rfc2819EtherStatsOversizedPkts uint64 `protobuf:"varint,81,opt,name=rfc2819_ether_stats_oversized_pkts,json=rfc2819EtherStatsOversizedPkts,proto3" json:"rfc2819_ether_stats_oversized_pkts,omitempty"`
	// RFC2819 etherStatsJabbers
	Rfc2819EtherStatsJabbers uint64 `protobuf:"varint,82,opt,name=rfc2819_ether_stats_jabbers,json=rfc2819EtherStatsJabbers,proto3" json:"rfc2819_ether_stats_jabbers,omitempty"`
	// RFC2819 etherStatsCRCAlignErrors
	Rfc2819EtherStatsCrcAlignErrors uint64 `protobuf:"varint,83,opt,name=rfc2819_ether_stats_crc_align_errors,json=rfc2819EtherStatsCrcAlignErrors,proto3" json:"rfc2819_ether_stats_crc_align_errors,omitempty"`
	// RFC3635 dot3StatsAlignmentErrors
	Rfc3635Dot3StatsAlignmentErrors uint64 `protobuf:"varint,84,opt,name=rfc3635dot3_stats_alignment_errors,json=rfc3635dot3StatsAlignmentErrors,proto3" json:"rfc3635dot3_stats_alignment_errors,omitempty"`
	// Total octets of all frames
	TotalBytesTransmitted uint64 `protobuf:"varint,85,opt,name=total_bytes_transmitted,json=totalBytesTransmitted,proto3" json:"total_bytes_transmitted,omitempty"`
	// Total octets of all good frames
	TotalGoodBytesTransmitted uint64 `protobuf:"varint,86,opt,name=total_good_bytes_transmitted,json=totalGoodBytesTransmitted,proto3" json:"total_good_bytes_transmitted,omitempty"`
	// All frames, good or bad
	TotalFramesTransmitted uint64 `protobuf:"varint,87,opt,name=total_frames_transmitted,json=totalFramesTransmitted,proto3" json:"total_frames_transmitted,omitempty"`
	// All 802.1Q frames
	Transmitted8021QFrames uint64 `protobuf:"varint,88,opt,name=transmitted8021_q_frames,json=transmitted8021QFrames,proto3" json:"transmitted8021_q_frames,omitempty"`
	// All pause frames
	TransmittedTotalPauseFrames uint64 `protobuf:"varint,89,opt,name=transmitted_total_pause_frames,json=transmittedTotalPauseFrames,proto3" json:"transmitted_total_pause_frames,omitempty"`
	// All 64 Octet Frame Count
	TransmittedTotal64OctetFrames uint64 `protobuf:"varint,90,opt,name=transmitted_total64_octet_frames,json=transmittedTotal64OctetFrames,proto3" json:"transmitted_total64_octet_frames,omitempty"`
	// All 65-127 Octet Frame Count
	TransmittedTotalOctetFramesFrom65To127 uint64 `protobuf:"varint,91,opt,name=transmitted_total_octet_frames_from65_to127,json=transmittedTotalOctetFramesFrom65To127,proto3" json:"transmitted_total_octet_frames_from65_to127,omitempty"`
	// All 128-255 Octet Frame Count
	TransmittedTotalOctetFramesFrom128To255 uint64 `protobuf:"varint,92,opt,name=transmitted_total_octet_frames_from128_to255,json=transmittedTotalOctetFramesFrom128To255,proto3" json:"transmitted_total_octet_frames_from128_to255,omitempty"`
	// All 256-511 Octet Frame Count
	TransmittedTotalOctetFramesFrom256To511 uint64 `protobuf:"varint,93,opt,name=transmitted_total_octet_frames_from256_to511,json=transmittedTotalOctetFramesFrom256To511,proto3" json:"transmitted_total_octet_frames_from256_to511,omitempty"`
	// All 512-1023 Octet Frame Count
	TransmittedTotalOctetFramesFrom512To1023 uint64 `protobuf:"varint,94,opt,name=transmitted_total_octet_frames_from512_to1023,json=transmittedTotalOctetFramesFrom512To1023,proto3" json:"transmitted_total_octet_frames_from512_to1023,omitempty"`
	// All 1024-1518 Octet Frame Count
	TransmittedTotalOctetFramesFrom1024To1518 uint64 `protobuf:"varint,95,opt,name=transmitted_total_octet_frames_from1024_to1518,json=transmittedTotalOctetFramesFrom1024To1518,proto3" json:"transmitted_total_octet_frames_from1024_to1518,omitempty"`
	// All > 1518 Octet Frame Count
	TransmittedTotalOctetFramesFrom1518ToMax uint64 `protobuf:"varint,96,opt,name=transmitted_total_octet_frames_from1518_to_max,json=transmittedTotalOctetFramesFrom1518ToMax,proto3" json:"transmitted_total_octet_frames_from1518_to_max,omitempty"`
	// Transmitted Good Frames
	TransmittedGoodFrames uint64 `protobuf:"varint,97,opt,name=transmitted_good_frames,json=transmittedGoodFrames,proto3" json:"transmitted_good_frames,omitempty"`
	// Transmitted unicast Frames
	TransmittedUnicastFrames uint64 `protobuf:"varint,98,opt,name=transmitted_unicast_frames,json=transmittedUnicastFrames,proto3" json:"transmitted_unicast_frames,omitempty"`
	// Transmitted multicast Frames
	TransmittedMulticastFrames uint64 `protobuf:"varint,99,opt,name=transmitted_multicast_frames,json=transmittedMulticastFrames,proto3" json:"transmitted_multicast_frames,omitempty"`
	// Transmitted broadcast Frames
	TransmittedBroadcastFrames uint64 `protobuf:"varint,100,opt,name=transmitted_broadcast_frames,json=transmittedBroadcastFrames,proto3" json:"transmitted_broadcast_frames,omitempty"`
	// Drops due to buffer underrun
	BufferUnderrunPacketDrops uint64 `protobuf:"varint,101,opt,name=buffer_underrun_packet_drops,json=bufferUnderrunPacketDrops,proto3" json:"buffer_underrun_packet_drops,omitempty"`
	// Drops due to packet abort
	AbortedPacketDrops uint64 `protobuf:"varint,102,opt,name=aborted_packet_drops,json=abortedPacketDrops,proto3" json:"aborted_packet_drops,omitempty"`
	// Any other drops not counted
	UncountedDroppedFrames uint64 `protobuf:"varint,103,opt,name=uncounted_dropped_frames,json=uncountedDroppedFrames,proto3" json:"uncounted_dropped_frames,omitempty"`
	// Any other errors not counted
	MiscellaneousOutputErrors uint64   `protobuf:"varint,104,opt,name=miscellaneous_output_errors,json=miscellaneousOutputErrors,proto3" json:"miscellaneous_output_errors,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *EthernetPortStatsType) Reset()         { *m = EthernetPortStatsType{} }
func (m *EthernetPortStatsType) String() string { return proto.CompactTextString(m) }
func (*EthernetPortStatsType) ProtoMessage()    {}
func (*EthernetPortStatsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982f68fad60746a, []int{1}
}

func (m *EthernetPortStatsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthernetPortStatsType.Unmarshal(m, b)
}
func (m *EthernetPortStatsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthernetPortStatsType.Marshal(b, m, deterministic)
}
func (m *EthernetPortStatsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthernetPortStatsType.Merge(m, src)
}
func (m *EthernetPortStatsType) XXX_Size() int {
	return xxx_messageInfo_EthernetPortStatsType.Size(m)
}
func (m *EthernetPortStatsType) XXX_DiscardUnknown() {
	xxx_messageInfo_EthernetPortStatsType.DiscardUnknown(m)
}

var xxx_messageInfo_EthernetPortStatsType proto.InternalMessageInfo

func (m *EthernetPortStatsType) GetReceivedTotalBytes() uint64 {
	if m != nil {
		return m.ReceivedTotalBytes
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedGoodBytes() uint64 {
	if m != nil {
		return m.ReceivedGoodBytes
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalFrames() uint64 {
	if m != nil {
		return m.ReceivedTotalFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceived8021QFrames() uint64 {
	if m != nil {
		return m.Received8021QFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedPauseFrames() uint64 {
	if m != nil {
		return m.ReceivedPauseFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedUnknownOpcodes() uint64 {
	if m != nil {
		return m.ReceivedUnknownOpcodes
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotal64OctetFrames() uint64 {
	if m != nil {
		return m.ReceivedTotal64OctetFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalOctetFramesFrom65To127() uint64 {
	if m != nil {
		return m.ReceivedTotalOctetFramesFrom65To127
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalOctetFramesFrom128To255() uint64 {
	if m != nil {
		return m.ReceivedTotalOctetFramesFrom128To255
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalOctetFramesFrom256To511() uint64 {
	if m != nil {
		return m.ReceivedTotalOctetFramesFrom256To511
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalOctetFramesFrom512To1023() uint64 {
	if m != nil {
		return m.ReceivedTotalOctetFramesFrom512To1023
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalOctetFramesFrom1024To1518() uint64 {
	if m != nil {
		return m.ReceivedTotalOctetFramesFrom1024To1518
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedTotalOctetFramesFrom1519ToMax() uint64 {
	if m != nil {
		return m.ReceivedTotalOctetFramesFrom1519ToMax
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedGoodFrames() uint64 {
	if m != nil {
		return m.ReceivedGoodFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedUnicastFrames() uint64 {
	if m != nil {
		return m.ReceivedUnicastFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedMulticastFrames() uint64 {
	if m != nil {
		return m.ReceivedMulticastFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetReceivedBroadcastFrames() uint64 {
	if m != nil {
		return m.ReceivedBroadcastFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetNumberOfBufferOverrunPacketsDropped() uint64 {
	if m != nil {
		return m.NumberOfBufferOverrunPacketsDropped
	}
	return 0
}

func (m *EthernetPortStatsType) GetNumberOfAbortedPacketsDropped() uint64 {
	if m != nil {
		return m.NumberOfAbortedPacketsDropped
	}
	return 0
}

func (m *EthernetPortStatsType) GetNumberofInvalidVlanIdPacketsDropped() uint64 {
	if m != nil {
		return m.NumberofInvalidVlanIdPacketsDropped
	}
	return 0
}

func (m *EthernetPortStatsType) GetInvalidDestMacDropPackets() uint64 {
	if m != nil {
		return m.InvalidDestMacDropPackets
	}
	return 0
}

func (m *EthernetPortStatsType) GetInvalidEncapDropPackets() uint64 {
	if m != nil {
		return m.InvalidEncapDropPackets
	}
	return 0
}

func (m *EthernetPortStatsType) GetNumberOfMiscellaneousPacketsDropped() uint64 {
	if m != nil {
		return m.NumberOfMiscellaneousPacketsDropped
	}
	return 0
}

func (m *EthernetPortStatsType) GetDroppedGiantPacketsGreaterthanMru() uint64 {
	if m != nil {
		return m.DroppedGiantPacketsGreaterthanMru
	}
	return 0
}

func (m *EthernetPortStatsType) GetDroppedEtherStatsUndersizePkts() uint64 {
	if m != nil {
		return m.DroppedEtherStatsUndersizePkts
	}
	return 0
}

func (m *EthernetPortStatsType) GetDroppedJabbersPacketsGreaterthanMru() uint64 {
	if m != nil {
		return m.DroppedJabbersPacketsGreaterthanMru
	}
	return 0
}

func (m *EthernetPortStatsType) GetDroppedEtherStatsFragments() uint64 {
	if m != nil {
		return m.DroppedEtherStatsFragments
	}
	return 0
}

func (m *EthernetPortStatsType) GetDroppedPacketsWithCrcAlignErrors() uint64 {
	if m != nil {
		return m.DroppedPacketsWithCrcAlignErrors
	}
	return 0
}

func (m *EthernetPortStatsType) GetEtherStatsCollisions() uint64 {
	if m != nil {
		return m.EtherStatsCollisions
	}
	return 0
}

func (m *EthernetPortStatsType) GetSymbolErrors() uint64 {
	if m != nil {
		return m.SymbolErrors
	}
	return 0
}

func (m *EthernetPortStatsType) GetDroppedMiscellaneousErrorPackets() uint64 {
	if m != nil {
		return m.DroppedMiscellaneousErrorPackets
	}
	return 0
}

func (m *EthernetPortStatsType) GetRfc2819EtherStatsOversizedPkts() uint64 {
	if m != nil {
		return m.Rfc2819EtherStatsOversizedPkts
	}
	return 0
}

func (m *EthernetPortStatsType) GetRfc2819EtherStatsJabbers() uint64 {
	if m != nil {
		return m.Rfc2819EtherStatsJabbers
	}
	return 0
}

func (m *EthernetPortStatsType) GetRfc2819EtherStatsCrcAlignErrors() uint64 {
	if m != nil {
		return m.Rfc2819EtherStatsCrcAlignErrors
	}
	return 0
}

func (m *EthernetPortStatsType) GetRfc3635Dot3StatsAlignmentErrors() uint64 {
	if m != nil {
		return m.Rfc3635Dot3StatsAlignmentErrors
	}
	return 0
}

func (m *EthernetPortStatsType) GetTotalBytesTransmitted() uint64 {
	if m != nil {
		return m.TotalBytesTransmitted
	}
	return 0
}

func (m *EthernetPortStatsType) GetTotalGoodBytesTransmitted() uint64 {
	if m != nil {
		return m.TotalGoodBytesTransmitted
	}
	return 0
}

func (m *EthernetPortStatsType) GetTotalFramesTransmitted() uint64 {
	if m != nil {
		return m.TotalFramesTransmitted
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmitted8021QFrames() uint64 {
	if m != nil {
		return m.Transmitted8021QFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalPauseFrames() uint64 {
	if m != nil {
		return m.TransmittedTotalPauseFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotal64OctetFrames() uint64 {
	if m != nil {
		return m.TransmittedTotal64OctetFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalOctetFramesFrom65To127() uint64 {
	if m != nil {
		return m.TransmittedTotalOctetFramesFrom65To127
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalOctetFramesFrom128To255() uint64 {
	if m != nil {
		return m.TransmittedTotalOctetFramesFrom128To255
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalOctetFramesFrom256To511() uint64 {
	if m != nil {
		return m.TransmittedTotalOctetFramesFrom256To511
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalOctetFramesFrom512To1023() uint64 {
	if m != nil {
		return m.TransmittedTotalOctetFramesFrom512To1023
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalOctetFramesFrom1024To1518() uint64 {
	if m != nil {
		return m.TransmittedTotalOctetFramesFrom1024To1518
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedTotalOctetFramesFrom1518ToMax() uint64 {
	if m != nil {
		return m.TransmittedTotalOctetFramesFrom1518ToMax
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedGoodFrames() uint64 {
	if m != nil {
		return m.TransmittedGoodFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedUnicastFrames() uint64 {
	if m != nil {
		return m.TransmittedUnicastFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedMulticastFrames() uint64 {
	if m != nil {
		return m.TransmittedMulticastFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetTransmittedBroadcastFrames() uint64 {
	if m != nil {
		return m.TransmittedBroadcastFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetBufferUnderrunPacketDrops() uint64 {
	if m != nil {
		return m.BufferUnderrunPacketDrops
	}
	return 0
}

func (m *EthernetPortStatsType) GetAbortedPacketDrops() uint64 {
	if m != nil {
		return m.AbortedPacketDrops
	}
	return 0
}

func (m *EthernetPortStatsType) GetUncountedDroppedFrames() uint64 {
	if m != nil {
		return m.UncountedDroppedFrames
	}
	return 0
}

func (m *EthernetPortStatsType) GetMiscellaneousOutputErrors() uint64 {
	if m != nil {
		return m.MiscellaneousOutputErrors
	}
	return 0
}

func init() {
	proto.RegisterType((*EthernetPortStatsType_KEYS)(nil), "cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.statistics.statistic.ethernet_port_stats_type_KEYS")
	proto.RegisterType((*EthernetPortStatsType)(nil), "cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.statistics.statistic.ethernet_port_stats_type")
}

func init() { proto.RegisterFile("ethernet_port_stats_type.proto", fileDescriptor_1982f68fad60746a) }

var fileDescriptor_1982f68fad60746a = []byte{
	// 1346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x97, 0x79, 0x73, 0x13, 0x37,
	0x18, 0xc6, 0x87, 0x99, 0x4e, 0x67, 0xaa, 0x29, 0x9d, 0xe9, 0x72, 0x89, 0x23, 0x69, 0x08, 0x50,
	0x42, 0x69, 0x33, 0xf1, 0x26, 0x6b, 0xcc, 0x4d, 0x12, 0x92, 0x40, 0x52, 0xe3, 0x00, 0x0e, 0x14,
	0x28, 0x15, 0xf2, 0xae, 0x1c, 0x6f, 0xe3, 0x5d, 0xb9, 0x92, 0x36, 0x40, 0xbf, 0x6b, 0xbf, 0x4b,
	0x47, 0xd7, 0xae, 0xb4, 0x4e, 0x62, 0xff, 0x17, 0xfc, 0x3e, 0xcf, 0x4f, 0xaf, 0x8e, 0x95, 0x1e,
	0xc0, 0x2c, 0x11, 0x03, 0xc2, 0x72, 0x22, 0xd0, 0x88, 0x32, 0x81, 0xb8, 0xc0, 0x82, 0x23, 0xf1,
	0x75, 0x44, 0x16, 0x47, 0x8c, 0x0a, 0x1a, 0xec, 0xc4, 0x29, 0x8f, 0x29, 0x4a, 0x29, 0x47, 0x5f,
	0x18, 0x4a, 0x58, 0x7a, 0x48, 0x18, 0x47, 0x19, 0x49, 0x52, 0x8c, 0x88, 0x18, 0x20, 0x3a, 0x22,
	0x6c, 0xb1, 0x64, 0xa4, 0xb9, 0x20, 0xac, 0x8f, 0x63, 0xb2, 0x28, 0x41, 0x29, 0x17, 0x69, 0xcc,
	0xab, 0x3f, 0xe7, 0x37, 0xc1, 0xcc, 0x71, 0xc3, 0xa1, 0x9d, 0x8d, 0x77, 0xaf, 0x83, 0x1b, 0xe0,
	0x87, 0x12, 0x81, 0x72, 0x9c, 0x11, 0x78, 0x6a, 0xee, 0xd4, 0xc2, 0x77, 0xaf, 0x4e, 0x97, 0xbf,
	0xbe, 0xc0, 0x19, 0x99, 0xff, 0x6f, 0x16, 0xc0, 0xe3, 0x40, 0xc1, 0x12, 0x38, 0xcb, 0x48, 0x4c,
	0xd2, 0x43, 0x92, 0x20, 0x41, 0x05, 0x1e, 0xa2, 0xde, 0x57, 0x41, 0x38, 0x0c, 0xe7, 0x4e, 0x2d,
	0x7c, 0xf3, 0x2a, 0xb0, 0xb5, 0xae, 0x2c, 0xad, 0xc9, 0x4a, 0xb0, 0x08, 0xce, 0x94, 0x8e, 0x7d,
	0x4a, 0x13, 0x63, 0x58, 0x56, 0x86, 0x1f, 0x6d, 0x69, 0x8b, 0xd2, 0x44, 0xeb, 0x43, 0x70, 0xae,
	0x36, 0x42, 0x9f, 0xe1, 0x8c, 0x70, 0xb8, 0xa2, 0x1c, 0x67, 0xbc, 0x21, 0x36, 0x55, 0xc9, 0xf5,
	0xb4, 0x96, 0xc2, 0x06, 0xfa, 0xc7, 0x7a, 0x22, 0xdf, 0x23, 0x8b, 0x2f, 0xc7, 0x3d, 0x68, 0x84,
	0x0b, 0x4e, 0xac, 0xa7, 0xe9, 0x7b, 0x76, 0x65, 0xcd, 0x78, 0x5a, 0x00, 0x96, 0x9e, 0x22, 0x3f,
	0xc8, 0xe9, 0xe7, 0x1c, 0xd1, 0x51, 0x4c, 0x13, 0xc2, 0xe1, 0x1d, 0x65, 0x3b, 0x6f, 0xeb, 0x7b,
	0xba, 0xdc, 0xd1, 0xd5, 0x60, 0x15, 0xcc, 0xf8, 0xb3, 0x6a, 0xae, 0x20, 0x1a, 0x0b, 0x22, 0xec,
	0xa8, 0x2d, 0x65, 0xbf, 0xe4, 0xcd, 0xae, 0xb9, 0xd2, 0x91, 0x12, 0x33, 0xf8, 0x1e, 0x58, 0xa8,
	0x2d, 0x8c, 0x0b, 0x40, 0x7d, 0x46, 0xb3, 0x66, 0x84, 0x04, 0x6d, 0x84, 0x77, 0xe0, 0x5d, 0x45,
	0xbb, 0xe6, 0xd1, 0x1c, 0xd6, 0xa6, 0xd2, 0x76, 0xa5, 0x34, 0x78, 0x0b, 0x6e, 0x4d, 0xc2, 0x36,
	0xc2, 0x16, 0x12, 0x34, 0x8c, 0x22, 0x78, 0x4f, 0x71, 0xaf, 0x9f, 0xc4, 0x6d, 0x84, 0xad, 0xae,
	0xd4, 0x4e, 0x03, 0x0e, 0xa3, 0x26, 0x12, 0x34, 0x6a, 0x34, 0xe0, 0xfd, 0xc9, 0xe0, 0x30, 0x6a,
	0x76, 0xa5, 0x36, 0x78, 0x07, 0x7e, 0x99, 0x04, 0x8e, 0x1a, 0xa1, 0x5c, 0x89, 0xa5, 0x70, 0x19,
	0x3e, 0x50, 0xe4, 0x1b, 0x27, 0x91, 0xa3, 0x46, 0xd8, 0x55, 0xe2, 0xe0, 0x03, 0xb8, 0x3d, 0x71,
	0x31, 0x96, 0xc2, 0x15, 0xc9, 0x8e, 0x1a, 0x2d, 0xf8, 0x50, 0xb1, 0x7f, 0x3e, 0x71, 0x39, 0x96,
	0xc2, 0x95, 0xae, 0x52, 0x07, 0xef, 0xa7, 0x80, 0x47, 0x8d, 0xbb, 0x48, 0x50, 0x94, 0xe1, 0x2f,
	0xf0, 0xd1, 0xe4, 0xc6, 0xa5, 0xbc, 0x4b, 0xdb, 0xf8, 0x8b, 0xf7, 0x5d, 0xaa, 0xaf, 0xcc, 0x1c,
	0xab, 0xc7, 0xfe, 0x77, 0x29, 0x3f, 0x33, 0x73, 0x9c, 0x9a, 0xe0, 0x82, 0x73, 0x96, 0xd3, 0x18,
	0xf3, 0xf2, 0x2c, 0x3e, 0x51, 0xa6, 0x73, 0xd5, 0x51, 0x56, 0x55, 0xe3, 0xbb, 0x07, 0x2e, 0x96,
	0xbe, 0xac, 0x18, 0x0a, 0xcf, 0xb9, 0xaa, 0x9c, 0x25, 0xb8, 0x6d, 0xeb, 0x47, 0x78, 0x7b, 0x8c,
	0xe2, 0xc4, 0xf5, 0xae, 0xf9, 0xde, 0x35, 0x5b, 0xaf, 0x8e, 0x7f, 0x5e, 0x64, 0x3d, 0xc2, 0x10,
	0xed, 0xa3, 0x5e, 0xd1, 0xef, 0xcb, 0xbf, 0x0e, 0x09, 0x63, 0x45, 0x8e, 0x46, 0x38, 0x3e, 0x20,
	0x82, 0xa3, 0x84, 0xd1, 0xd1, 0x88, 0x24, 0x70, 0x5d, 0x1f, 0x7f, 0xad, 0xef, 0xf4, 0xd7, 0x94,
	0xba, 0xa3, 0xc5, 0xbb, 0x5a, 0xfb, 0x54, 0x4b, 0x83, 0x67, 0xe0, 0x6a, 0x85, 0xc5, 0x3d, 0xca,
	0x84, 0xba, 0x0f, 0x7c, 0xde, 0x53, 0xc5, 0x9b, 0xb1, 0xbc, 0x55, 0x2d, 0xab, 0x91, 0xca, 0x06,
	0x69, 0x1f, 0xa5, 0xf9, 0x21, 0x1e, 0xa6, 0x09, 0x3a, 0x1c, 0xe2, 0x1c, 0xa5, 0xe3, 0xc0, 0x0d,
	0xb7, 0x41, 0xda, 0x7f, 0xae, 0xe5, 0x6f, 0x86, 0x38, 0x7f, 0x5e, 0xc7, 0x3e, 0x01, 0x33, 0x96,
	0x96, 0x10, 0x2e, 0x50, 0x86, 0x63, 0x85, 0xb1, 0x4c, 0xb8, 0xa9, 0x58, 0x17, 0x8d, 0xe8, 0x29,
	0xe1, 0xa2, 0x8d, 0x63, 0xe9, 0x36, 0xa0, 0xe0, 0x3e, 0xb8, 0x64, 0x09, 0x24, 0x8f, 0xf1, 0xc8,
	0xb7, 0x6f, 0xe9, 0x65, 0x37, 0x8a, 0x0d, 0x29, 0x70, 0xcd, 0x5d, 0x70, 0xb3, 0x5a, 0x9f, 0x2c,
	0xe5, 0x31, 0x19, 0x0e, 0x71, 0x4e, 0x68, 0xc1, 0xc7, 0x26, 0xf5, 0xcc, 0x5f, 0xf5, 0xb6, 0x2b,
	0xae, 0x4d, 0x6a, 0x17, 0xdc, 0x30, 0x2e, 0xb4, 0x9f, 0xe2, 0x5c, 0x94, 0xac, 0x7d, 0x46, 0xb0,
	0x20, 0x4c, 0x0c, 0x70, 0x8e, 0x32, 0x56, 0xc0, 0xe7, 0x8a, 0x79, 0xd5, 0x88, 0xb7, 0xa4, 0xd6,
	0xa0, 0xb6, 0x2a, 0x65, 0x9b, 0x15, 0xc1, 0x36, 0x98, 0xb7, 0x44, 0xf5, 0x78, 0x99, 0x47, 0xab,
	0xc8, 0x13, 0xc2, 0x78, 0xfa, 0x2f, 0x41, 0xa3, 0x03, 0xc1, 0xe1, 0xb6, 0xc2, 0xcd, 0x1a, 0xe5,
	0x86, 0x14, 0xbe, 0x96, 0xba, 0x3d, 0x2b, 0xdb, 0x3d, 0xd0, 0x73, 0xb6, 0xac, 0xbf, 0x71, 0xaf,
	0x27, 0xdf, 0xe4, 0xe3, 0xfa, 0xdb, 0xd1, 0x73, 0x36, 0xf2, 0x6d, 0xad, 0x3e, 0xba, 0xc3, 0x55,
	0x30, 0x73, 0x54, 0x87, 0x7d, 0x86, 0xf7, 0x33, 0x92, 0x0b, 0x0e, 0x7f, 0xd7, 0x4f, 0xc0, 0x58,
	0x73, 0x9b, 0x56, 0x11, 0x74, 0xaa, 0x65, 0xb3, 0x0d, 0x7d, 0x4e, 0xc5, 0x00, 0xc5, 0x2c, 0x46,
	0x78, 0x98, 0xee, 0xe7, 0x88, 0x30, 0x46, 0x19, 0x87, 0x6d, 0x85, 0x9a, 0x33, 0x62, 0xd3, 0xcf,
	0xdb, 0x54, 0x0c, 0xd6, 0x59, 0xbc, 0x2a, 0x85, 0x1b, 0x4a, 0x17, 0xac, 0x80, 0xf3, 0x6e, 0x2f,
	0x31, 0x1d, 0x0e, 0x53, 0x9e, 0xd2, 0x9c, 0xc3, 0x17, 0x8a, 0x70, 0x96, 0x94, 0x5d, 0xac, 0x97,
	0xb5, 0xe0, 0x1a, 0x38, 0xcd, 0xbf, 0x66, 0x3d, 0x3a, 0xb4, 0xc3, 0x75, 0x94, 0xf8, 0x7b, 0xfd,
	0xa3, 0x41, 0xb7, 0x81, 0x5d, 0x95, 0xda, 0xb1, 0x51, 0x9e, 0xf2, 0xf8, 0xed, 0x7a, 0x9d, 0x7a,
	0x67, 0x46, 0x81, 0xec, 0x39, 0xdc, 0x06, 0xf3, 0xac, 0x1f, 0x87, 0xad, 0xc6, 0x5d, 0x6f, 0xf5,
	0xe4, 0x0d, 0x20, 0xf7, 0x2d, 0xd1, 0xfb, 0xfb, 0x52, 0xef, 0xaf, 0x51, 0x56, 0x4b, 0xd8, 0xb1,
	0x32, 0xb5, 0xbf, 0x0f, 0xc1, 0xe5, 0xa3, 0x58, 0x66, 0xaf, 0xe1, 0x2b, 0x05, 0x81, 0x63, 0x10,
	0xb3, 0xbb, 0x41, 0x1b, 0x5c, 0x3f, 0xca, 0x3e, 0xb6, 0x09, 0xaf, 0x15, 0xe7, 0xa7, 0x31, 0x4e,
	0x6d, 0x0f, 0x76, 0xd4, 0xcc, 0x96, 0x9b, 0xcb, 0x51, 0x42, 0xc5, 0xb2, 0x81, 0x29, 0x90, 0xdc,
	0x74, 0x0b, 0xeb, 0x96, 0x30, 0xab, 0x54, 0xac, 0x55, 0xab, 0x33, 0xb0, 0x26, 0xb8, 0xe0, 0xc4,
	0x32, 0x24, 0x18, 0xce, 0x79, 0x96, 0x0a, 0x41, 0x12, 0xb8, 0xa7, 0x6f, 0x75, 0x51, 0x46, 0xb3,
	0x6e, 0x55, 0x0c, 0x1e, 0x83, 0x2b, 0xda, 0x57, 0x45, 0x34, 0xcf, 0xfc, 0x46, 0x5f, 0x32, 0x4a,
	0x53, 0x66, 0x35, 0x17, 0xd0, 0x02, 0xd0, 0x4d, 0x6b, 0x9e, 0xf9, 0xad, 0x8e, 0x46, 0xa2, 0x4a,
	0x6c, 0x75, 0x67, 0xf5, 0x4f, 0x3f, 0xbf, 0xfd, 0x61, 0x9c, 0x7e, 0xdd, 0x46, 0xb8, 0x75, 0x30,
	0xeb, 0x54, 0xcc, 0x9b, 0xea, 0x65, 0xb9, 0x77, 0xca, 0x7f, 0xd9, 0x51, 0xa9, 0x67, 0xd4, 0xcd,
	0x74, 0x5b, 0x60, 0x6e, 0x0c, 0x52, 0x0f, 0x67, 0xef, 0xf5, 0xfd, 0x5f, 0xc7, 0xf8, 0xf9, 0xec,
	0x03, 0xb8, 0x3d, 0xde, 0xcd, 0xf1, 0x11, 0xed, 0x83, 0xce, 0x0e, 0x75, 0xe6, 0x31, 0x29, 0xed,
	0x23, 0xf8, 0x75, 0x0a, 0x78, 0x15, 0xd4, 0xfe, 0x54, 0xf4, 0x9b, 0x13, 0xe8, 0x65, 0x56, 0x9b,
	0x0e, 0x5f, 0xc5, 0xb5, 0x8f, 0x53, 0xe1, 0xcb, 0xc4, 0x86, 0xc0, 0x6f, 0x53, 0xe0, 0x9d, 0xd0,
	0xf6, 0x97, 0xe2, 0x2f, 0x4c, 0xe0, 0x57, 0xb9, 0x0d, 0x83, 0xc5, 0x69, 0x96, 0xc7, 0x89, 0x6e,
	0x48, 0x8d, 0x70, 0x6b, 0xd2, 0x02, 0x55, 0xe9, 0xed, 0xd3, 0x74, 0x43, 0x44, 0x8d, 0x96, 0x0d,
	0x70, 0x9f, 0xa6, 0x9a, 0x84, 0x74, 0xe8, 0x0c, 0x27, 0xbf, 0x5d, 0x67, 0x04, 0x37, 0xc6, 0x61,
	0xf3, 0xed, 0x56, 0x65, 0x27, 0xc9, 0x3d, 0x00, 0x97, 0x5c, 0x5f, 0x2d, 0xcc, 0xf5, 0xf4, 0x6d,
	0xe6, 0x28, 0xfc, 0x3c, 0xf7, 0x04, 0x5c, 0x71, 0xdd, 0x63, 0x91, 0x2e, 0xd6, 0xaf, 0x92, 0xa3,
	0xa9, 0xa7, 0xba, 0x1a, 0x61, 0x2c, 0xd8, 0x25, 0x63, 0x84, 0x7a, 0xb6, 0x7b, 0x0c, 0xae, 0x98,
	0x44, 0xa7, 0xde, 0xeb, 0x2a, 0xd2, 0xa9, 0x6c, 0xc1, 0x21, 0xd1, 0xb7, 0x8f, 0xd6, 0xec, 0x19,
	0x89, 0x7e, 0x19, 0x64, 0xa2, 0xe0, 0x32, 0xfe, 0xfa, 0xd9, 0xcd, 0x18, 0xfb, 0x3a, 0xfe, 0x62,
	0x37, 0xb0, 0x69, 0x47, 0x0b, 0xc0, 0x22, 0x8f, 0x69, 0x91, 0x4b, 0x8f, 0x7d, 0xa8, 0x4c, 0xc3,
	0xfb, 0xfa, 0xd6, 0x29, 0xeb, 0x26, 0xb5, 0x98, 0x66, 0x1f, 0x81, 0xcb, 0xfe, 0x83, 0x46, 0x0b,
	0x31, 0x2a, 0xca, 0x8b, 0x7a, 0xa0, 0x7b, 0xf5, 0x24, 0x1d, 0xa5, 0xd0, 0x57, 0x74, 0xef, 0x5b,
	0xf5, 0x7f, 0xff, 0xe5, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xe0, 0x65, 0xf7, 0x1d, 0x10,
	0x00, 0x00,
}
