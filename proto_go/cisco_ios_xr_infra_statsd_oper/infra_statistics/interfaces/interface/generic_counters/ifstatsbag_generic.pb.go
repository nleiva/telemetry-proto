// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ifstatsbag_generic.proto

package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_generic_counters

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

// Generic set of interface counters
type IfstatsbagGeneric_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfstatsbagGeneric_KEYS) Reset()         { *m = IfstatsbagGeneric_KEYS{} }
func (m *IfstatsbagGeneric_KEYS) String() string { return proto.CompactTextString(m) }
func (*IfstatsbagGeneric_KEYS) ProtoMessage()    {}
func (*IfstatsbagGeneric_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ifstatsbag_generic_28d250fe93cd2383, []int{0}
}
func (m *IfstatsbagGeneric_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagGeneric_KEYS.Unmarshal(m, b)
}
func (m *IfstatsbagGeneric_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagGeneric_KEYS.Marshal(b, m, deterministic)
}
func (dst *IfstatsbagGeneric_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagGeneric_KEYS.Merge(dst, src)
}
func (m *IfstatsbagGeneric_KEYS) XXX_Size() int {
	return xxx_messageInfo_IfstatsbagGeneric_KEYS.Size(m)
}
func (m *IfstatsbagGeneric_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IfstatsbagGeneric_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IfstatsbagGeneric_KEYS proto.InternalMessageInfo

func (m *IfstatsbagGeneric_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IfstatsbagGeneric struct {
	// Packets received
	PacketsReceived uint64 `protobuf:"varint,50,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	// Bytes received
	BytesReceived uint64 `protobuf:"varint,51,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	// Packets sent
	PacketsSent uint64 `protobuf:"varint,52,opt,name=packets_sent,json=packetsSent,proto3" json:"packets_sent,omitempty"`
	// Bytes sent
	BytesSent uint64 `protobuf:"varint,53,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	// Multicast packets received
	MulticastPacketsReceived uint64 `protobuf:"varint,54,opt,name=multicast_packets_received,json=multicastPacketsReceived,proto3" json:"multicast_packets_received,omitempty"`
	// Broadcast packets received
	BroadcastPacketsReceived uint64 `protobuf:"varint,55,opt,name=broadcast_packets_received,json=broadcastPacketsReceived,proto3" json:"broadcast_packets_received,omitempty"`
	// Multicast packets sent
	MulticastPacketsSent uint64 `protobuf:"varint,56,opt,name=multicast_packets_sent,json=multicastPacketsSent,proto3" json:"multicast_packets_sent,omitempty"`
	// Broadcast packets sent
	BroadcastPacketsSent uint64 `protobuf:"varint,57,opt,name=broadcast_packets_sent,json=broadcastPacketsSent,proto3" json:"broadcast_packets_sent,omitempty"`
	// Total output drops
	OutputDrops uint32 `protobuf:"varint,58,opt,name=output_drops,json=outputDrops,proto3" json:"output_drops,omitempty"`
	// Output queue drops
	OutputQueueDrops uint32 `protobuf:"varint,59,opt,name=output_queue_drops,json=outputQueueDrops,proto3" json:"output_queue_drops,omitempty"`
	// Total input drops
	InputDrops uint32 `protobuf:"varint,60,opt,name=input_drops,json=inputDrops,proto3" json:"input_drops,omitempty"`
	// Input queue drops
	InputQueueDrops uint32 `protobuf:"varint,61,opt,name=input_queue_drops,json=inputQueueDrops,proto3" json:"input_queue_drops,omitempty"`
	// Received runt packets
	RuntPacketsReceived uint32 `protobuf:"varint,62,opt,name=runt_packets_received,json=runtPacketsReceived,proto3" json:"runt_packets_received,omitempty"`
	// Received giant packets
	GiantPacketsReceived uint32 `protobuf:"varint,63,opt,name=giant_packets_received,json=giantPacketsReceived,proto3" json:"giant_packets_received,omitempty"`
	// Received throttled packets
	ThrottledPacketsReceived uint32 `protobuf:"varint,64,opt,name=throttled_packets_received,json=throttledPacketsReceived,proto3" json:"throttled_packets_received,omitempty"`
	// Received parity packets
	ParityPacketsReceived uint32 `protobuf:"varint,65,opt,name=parity_packets_received,json=parityPacketsReceived,proto3" json:"parity_packets_received,omitempty"`
	// Unknown protocol packets received
	UnknownProtocolPacketsReceived uint32 `protobuf:"varint,66,opt,name=unknown_protocol_packets_received,json=unknownProtocolPacketsReceived,proto3" json:"unknown_protocol_packets_received,omitempty"`
	// Total input errors
	InputErrors uint32 `protobuf:"varint,67,opt,name=input_errors,json=inputErrors,proto3" json:"input_errors,omitempty"`
	// Input CRC errors
	CrcErrors uint32 `protobuf:"varint,68,opt,name=crc_errors,json=crcErrors,proto3" json:"crc_errors,omitempty"`
	// Input overruns
	InputOverruns uint32 `protobuf:"varint,69,opt,name=input_overruns,json=inputOverruns,proto3" json:"input_overruns,omitempty"`
	// Framing-errors received
	FramingErrorsReceived uint32 `protobuf:"varint,70,opt,name=framing_errors_received,json=framingErrorsReceived,proto3" json:"framing_errors_received,omitempty"`
	// Input ignored packets
	InputIgnoredPackets uint32 `protobuf:"varint,71,opt,name=input_ignored_packets,json=inputIgnoredPackets,proto3" json:"input_ignored_packets,omitempty"`
	// Input aborts
	InputAborts uint32 `protobuf:"varint,72,opt,name=input_aborts,json=inputAborts,proto3" json:"input_aborts,omitempty"`
	// Total output errors
	OutputErrors uint32 `protobuf:"varint,73,opt,name=output_errors,json=outputErrors,proto3" json:"output_errors,omitempty"`
	// Output underruns
	OutputUnderruns uint32 `protobuf:"varint,74,opt,name=output_underruns,json=outputUnderruns,proto3" json:"output_underruns,omitempty"`
	// Output buffer failures
	OutputBufferFailures uint32 `protobuf:"varint,75,opt,name=output_buffer_failures,json=outputBufferFailures,proto3" json:"output_buffer_failures,omitempty"`
	// Output buffers swapped out
	OutputBuffersSwappedOut uint32 `protobuf:"varint,76,opt,name=output_buffers_swapped_out,json=outputBuffersSwappedOut,proto3" json:"output_buffers_swapped_out,omitempty"`
	// Applique
	Applique uint32 `protobuf:"varint,77,opt,name=applique,proto3" json:"applique,omitempty"`
	// Number of board resets
	Resets uint32 `protobuf:"varint,78,opt,name=resets,proto3" json:"resets,omitempty"`
	// Carrier transitions
	CarrierTransitions uint32 `protobuf:"varint,79,opt,name=carrier_transitions,json=carrierTransitions,proto3" json:"carrier_transitions,omitempty"`
	// Availability bit mask
	AvailabilityFlag uint32 `protobuf:"varint,80,opt,name=availability_flag,json=availabilityFlag,proto3" json:"availability_flag,omitempty"`
	// Time when counters were last written (in seconds)
	LastDataTime uint32 `protobuf:"varint,81,opt,name=last_data_time,json=lastDataTime,proto3" json:"last_data_time,omitempty"`
	// Number of seconds since last clear counters
	SecondsSinceLastClearCounters uint32 `protobuf:"varint,82,opt,name=seconds_since_last_clear_counters,json=secondsSinceLastClearCounters,proto3" json:"seconds_since_last_clear_counters,omitempty"`
	// SysUpTime when counters were last reset (in seconds)
	LastDiscontinuityTime uint32 `protobuf:"varint,83,opt,name=last_discontinuity_time,json=lastDiscontinuityTime,proto3" json:"last_discontinuity_time,omitempty"`
	// Seconds since packet received
	SecondsSincePacketReceived uint32 `protobuf:"varint,84,opt,name=seconds_since_packet_received,json=secondsSincePacketReceived,proto3" json:"seconds_since_packet_received,omitempty"`
	// Seconds since packet sent
	SecondsSincePacketSent uint32   `protobuf:"varint,85,opt,name=seconds_since_packet_sent,json=secondsSincePacketSent,proto3" json:"seconds_since_packet_sent,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *IfstatsbagGeneric) Reset()         { *m = IfstatsbagGeneric{} }
func (m *IfstatsbagGeneric) String() string { return proto.CompactTextString(m) }
func (*IfstatsbagGeneric) ProtoMessage()    {}
func (*IfstatsbagGeneric) Descriptor() ([]byte, []int) {
	return fileDescriptor_ifstatsbag_generic_28d250fe93cd2383, []int{1}
}
func (m *IfstatsbagGeneric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagGeneric.Unmarshal(m, b)
}
func (m *IfstatsbagGeneric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagGeneric.Marshal(b, m, deterministic)
}
func (dst *IfstatsbagGeneric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagGeneric.Merge(dst, src)
}
func (m *IfstatsbagGeneric) XXX_Size() int {
	return xxx_messageInfo_IfstatsbagGeneric.Size(m)
}
func (m *IfstatsbagGeneric) XXX_DiscardUnknown() {
	xxx_messageInfo_IfstatsbagGeneric.DiscardUnknown(m)
}

var xxx_messageInfo_IfstatsbagGeneric proto.InternalMessageInfo

func (m *IfstatsbagGeneric) GetPacketsReceived() uint64 {
	if m != nil {
		return m.PacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetPacketsSent() uint64 {
	if m != nil {
		return m.PacketsSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetMulticastPacketsReceived() uint64 {
	if m != nil {
		return m.MulticastPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBroadcastPacketsReceived() uint64 {
	if m != nil {
		return m.BroadcastPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetMulticastPacketsSent() uint64 {
	if m != nil {
		return m.MulticastPacketsSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBroadcastPacketsSent() uint64 {
	if m != nil {
		return m.BroadcastPacketsSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputDrops() uint32 {
	if m != nil {
		return m.OutputDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputQueueDrops() uint32 {
	if m != nil {
		return m.OutputQueueDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputDrops() uint32 {
	if m != nil {
		return m.InputDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputQueueDrops() uint32 {
	if m != nil {
		return m.InputQueueDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetRuntPacketsReceived() uint32 {
	if m != nil {
		return m.RuntPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetGiantPacketsReceived() uint32 {
	if m != nil {
		return m.GiantPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetThrottledPacketsReceived() uint32 {
	if m != nil {
		return m.ThrottledPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetParityPacketsReceived() uint32 {
	if m != nil {
		return m.ParityPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetUnknownProtocolPacketsReceived() uint32 {
	if m != nil {
		return m.UnknownProtocolPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputErrors() uint32 {
	if m != nil {
		return m.InputErrors
	}
	return 0
}

func (m *IfstatsbagGeneric) GetCrcErrors() uint32 {
	if m != nil {
		return m.CrcErrors
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputOverruns() uint32 {
	if m != nil {
		return m.InputOverruns
	}
	return 0
}

func (m *IfstatsbagGeneric) GetFramingErrorsReceived() uint32 {
	if m != nil {
		return m.FramingErrorsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputIgnoredPackets() uint32 {
	if m != nil {
		return m.InputIgnoredPackets
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputAborts() uint32 {
	if m != nil {
		return m.InputAborts
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputErrors() uint32 {
	if m != nil {
		return m.OutputErrors
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputUnderruns() uint32 {
	if m != nil {
		return m.OutputUnderruns
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputBufferFailures() uint32 {
	if m != nil {
		return m.OutputBufferFailures
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputBuffersSwappedOut() uint32 {
	if m != nil {
		return m.OutputBuffersSwappedOut
	}
	return 0
}

func (m *IfstatsbagGeneric) GetApplique() uint32 {
	if m != nil {
		return m.Applique
	}
	return 0
}

func (m *IfstatsbagGeneric) GetResets() uint32 {
	if m != nil {
		return m.Resets
	}
	return 0
}

func (m *IfstatsbagGeneric) GetCarrierTransitions() uint32 {
	if m != nil {
		return m.CarrierTransitions
	}
	return 0
}

func (m *IfstatsbagGeneric) GetAvailabilityFlag() uint32 {
	if m != nil {
		return m.AvailabilityFlag
	}
	return 0
}

func (m *IfstatsbagGeneric) GetLastDataTime() uint32 {
	if m != nil {
		return m.LastDataTime
	}
	return 0
}

func (m *IfstatsbagGeneric) GetSecondsSinceLastClearCounters() uint32 {
	if m != nil {
		return m.SecondsSinceLastClearCounters
	}
	return 0
}

func (m *IfstatsbagGeneric) GetLastDiscontinuityTime() uint32 {
	if m != nil {
		return m.LastDiscontinuityTime
	}
	return 0
}

func (m *IfstatsbagGeneric) GetSecondsSincePacketReceived() uint32 {
	if m != nil {
		return m.SecondsSincePacketReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetSecondsSincePacketSent() uint32 {
	if m != nil {
		return m.SecondsSincePacketSent
	}
	return 0
}

func init() {
	proto.RegisterType((*IfstatsbagGeneric_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.generic_counters.ifstatsbag_generic_KEYS")
	proto.RegisterType((*IfstatsbagGeneric)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.generic_counters.ifstatsbag_generic")
}

func init() {
	proto.RegisterFile("ifstatsbag_generic.proto", fileDescriptor_ifstatsbag_generic_28d250fe93cd2383)
}

var fileDescriptor_ifstatsbag_generic_28d250fe93cd2383 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0xed, 0x6f, 0x1c, 0x35,
	0x10, 0xc6, 0x15, 0x09, 0x55, 0x64, 0x92, 0xeb, 0x8b, 0xdb, 0x24, 0xe6, 0xa4, 0x40, 0xae, 0x80,
	0x94, 0x02, 0x0a, 0x52, 0x1a, 0x0a, 0xa5, 0x01, 0x9a, 0xe6, 0x85, 0x86, 0x96, 0x26, 0xbd, 0x4b,
	0x3e, 0xf0, 0xc9, 0xf2, 0x79, 0xbd, 0x87, 0xd5, 0x3d, 0x7b, 0xeb, 0x97, 0x94, 0xfe, 0x8f, 0xfc,
	0x51, 0x95, 0xc7, 0xce, 0xde, 0xe6, 0xf6, 0xbe, 0x25, 0xcf, 0xf3, 0x7b, 0x66, 0xc6, 0xde, 0xd9,
	0x3d, 0xa0, 0xaa, 0x74, 0x9e, 0x7b, 0x37, 0xe6, 0x13, 0x36, 0x91, 0x5a, 0x5a, 0x25, 0x76, 0x6a,
	0x6b, 0xbc, 0x21, 0x97, 0x42, 0x39, 0x61, 0x98, 0x32, 0x8e, 0xfd, 0x67, 0x99, 0xd2, 0xa5, 0xe5,
	0x0c, 0xd1, 0x82, 0x99, 0x5a, 0xda, 0x9d, 0x99, 0xa2, 0x9c, 0x57, 0xc2, 0xed, 0x28, 0xed, 0xa5,
	0x2d, 0xb9, 0x90, 0xad, 0x3f, 0x77, 0x72, 0x4d, 0x26, 0x4c, 0x88, 0x9a, 0x7b, 0xf8, 0x1c, 0x36,
	0xba, 0x2d, 0xd9, 0xab, 0xe3, 0x7f, 0x46, 0xe4, 0x5b, 0xb8, 0xdd, 0x04, 0x99, 0xe6, 0x53, 0x49,
	0x97, 0xb6, 0x96, 0xb6, 0x97, 0x87, 0xbd, 0x46, 0x7d, 0xc3, 0xa7, 0xf2, 0xe1, 0xff, 0x3d, 0x20,
	0xdd, 0x12, 0xe4, 0x11, 0xdc, 0xad, 0xb9, 0x78, 0x27, 0xbd, 0x63, 0x56, 0x0a, 0xa9, 0xae, 0x64,
	0x41, 0x77, 0xb7, 0x96, 0xb6, 0x3f, 0x1b, 0xde, 0xc9, 0xfa, 0x30, 0xcb, 0xb1, 0xd1, 0xf8, 0xa3,
	0x97, 0x2d, 0xf0, 0x31, 0x82, 0x3d, 0x54, 0x1b, 0x6c, 0x00, 0xab, 0xd7, 0x15, 0x9d, 0xd4, 0x9e,
	0xee, 0x21, 0xb4, 0x92, 0xb5, 0x91, 0xd4, 0x9e, 0x6c, 0x02, 0xa4, 0x4a, 0x08, 0xfc, 0x84, 0xc0,
	0x32, 0x2a, 0x68, 0xef, 0x43, 0x7f, 0x1a, 0x2a, 0xaf, 0x04, 0x77, 0x9e, 0x75, 0xa6, 0x7b, 0x82,
	0x38, 0x6d, 0x88, 0xf3, 0xb9, 0x31, 0xf7, 0xa1, 0x3f, 0xb6, 0x86, 0x17, 0x8b, 0xd3, 0x3f, 0xa7,
	0x74, 0x43, 0xcc, 0xa7, 0xf7, 0x60, 0xbd, 0xdb, 0x1b, 0xc7, 0xfc, 0x05, 0x93, 0x0f, 0xe6, 0xfb,
	0xe2, 0xc4, 0x7b, 0xb0, 0xde, 0xed, 0x89, 0xa9, 0xa7, 0x29, 0x35, 0xdf, 0x0f, 0x53, 0x03, 0x58,
	0x35, 0xc1, 0xd7, 0xc1, 0xb3, 0xc2, 0x9a, 0xda, 0xd1, 0x5f, 0xb7, 0x96, 0xb6, 0x7b, 0xc3, 0x95,
	0xa4, 0x1d, 0x45, 0x89, 0xfc, 0x00, 0x24, 0x23, 0xef, 0x83, 0x0c, 0x32, 0x83, 0xcf, 0x10, 0xbc,
	0x9b, 0x9c, 0xb7, 0xd1, 0x48, 0xf4, 0x57, 0xb0, 0xa2, 0xf4, 0xac, 0xde, 0x3e, 0x62, 0x80, 0x52,
	0x02, 0xbe, 0x83, 0x7b, 0x09, 0x68, 0x57, 0xfb, 0x0d, 0xb1, 0x3b, 0x68, 0xb4, 0x8a, 0xed, 0xc2,
	0x9a, 0x0d, 0x7a, 0xc1, 0x15, 0xfe, 0x8e, 0xfc, 0xfd, 0x68, 0x2e, 0xb8, 0xbd, 0x89, 0xe2, 0x8b,
	0x42, 0x7f, 0x60, 0xe8, 0x01, 0xba, 0x0b, 0x9e, 0x98, 0xff, 0xd7, 0x1a, 0xef, 0x2b, 0x59, 0x74,
	0x93, 0xcf, 0x31, 0x49, 0x1b, 0x62, 0x3e, 0xfd, 0x04, 0x36, 0x6a, 0x6e, 0x95, 0xff, 0xd8, 0x8d,
	0x1e, 0x60, 0x74, 0x2d, 0xd9, 0xf3, 0xb9, 0x53, 0x18, 0x04, 0xfd, 0x4e, 0x9b, 0x0f, 0x9a, 0xe1,
	0xab, 0x2b, 0x4c, 0xd5, 0xad, 0xf0, 0x02, 0x2b, 0x7c, 0x99, 0xc1, 0xf3, 0xcc, 0xcd, 0x97, 0x1a,
	0xc0, 0x6a, 0xba, 0x56, 0x69, 0xad, 0xb1, 0x8e, 0x1e, 0xa6, 0x07, 0x89, 0xda, 0x31, 0x4a, 0x71,
	0xe5, 0x85, 0x15, 0xd7, 0xc0, 0x11, 0x02, 0xcb, 0xc2, 0x8a, 0x6c, 0xe3, 0x4b, 0x1c, 0x2b, 0x98,
	0x2b, 0x69, 0x6d, 0xd0, 0x8e, 0x1e, 0x23, 0xd2, 0x43, 0xf5, 0x2c, 0x8b, 0xf1, 0xac, 0xa5, 0xe5,
	0x53, 0xa5, 0x27, 0xb9, 0xd2, 0x6c, 0xd2, 0x93, 0x74, 0xd6, 0x6c, 0xa7, 0xb2, 0xcd, 0x80, 0xbb,
	0xb0, 0x96, 0xca, 0xab, 0x89, 0x36, 0x76, 0x76, 0xcb, 0xf4, 0xcf, 0xf4, 0x2c, 0xd1, 0x3c, 0x4d,
	0x5e, 0x3e, 0xdc, 0xec, 0x50, 0x7c, 0x6c, 0xac, 0x77, 0xf4, 0x65, 0xeb, 0x50, 0x07, 0x28, 0x91,
	0xaf, 0xa1, 0x97, 0xb7, 0x33, 0x9f, 0xeb, 0x14, 0x99, 0xbc, 0xd5, 0xf9, 0x68, 0x8f, 0x20, 0x2f,
	0x2a, 0x0b, 0xba, 0xc8, 0x87, 0xfb, 0x2b, 0xad, 0x5c, 0xd2, 0x2f, 0xaf, 0xe5, 0xb8, 0x3e, 0x19,
	0x1d, 0x87, 0xb2, 0x94, 0x96, 0x95, 0x5c, 0x55, 0xc1, 0x4a, 0x47, 0x5f, 0xa5, 0xf5, 0x49, 0xee,
	0x0b, 0x34, 0x4f, 0xb2, 0x47, 0x9e, 0x41, 0xff, 0x46, 0xca, 0x31, 0xf7, 0x81, 0xd7, 0xb5, 0x2c,
	0x98, 0x09, 0x9e, 0xbe, 0xc6, 0xe4, 0x46, 0x3b, 0xe9, 0x46, 0xc9, 0x3f, 0x0b, 0x9e, 0xf4, 0xe1,
	0x73, 0x5e, 0xd7, 0x95, 0x7a, 0x1f, 0x24, 0xfd, 0x1b, 0xd1, 0xe6, 0x7f, 0xb2, 0x0e, 0xb7, 0xac,
	0x74, 0xf1, 0x9a, 0xde, 0xa0, 0x93, 0xff, 0x23, 0x3f, 0xc2, 0x7d, 0xc1, 0xad, 0x55, 0xd2, 0x32,
	0x6f, 0xb9, 0x76, 0xca, 0x2b, 0xa3, 0x1d, 0x3d, 0x43, 0x88, 0x64, 0xeb, 0x62, 0xe6, 0x90, 0xef,
	0xe1, 0x1e, 0xbf, 0xe2, 0xaa, 0xe2, 0x63, 0x55, 0xc5, 0x45, 0x2d, 0x2b, 0x3e, 0xa1, 0xe7, 0xe9,
	0x25, 0x6e, 0x1b, 0x27, 0x15, 0x9f, 0x90, 0x6f, 0xe0, 0x76, 0x15, 0x3f, 0x23, 0x05, 0xf7, 0x9c,
	0x79, 0x35, 0x95, 0xf4, 0x6d, 0xba, 0xd5, 0xa8, 0x1e, 0x71, 0xcf, 0x2f, 0xd4, 0x54, 0x92, 0x97,
	0x30, 0x70, 0x52, 0x18, 0x5d, 0x38, 0xe6, 0x94, 0x16, 0x92, 0x61, 0x46, 0x54, 0x92, 0xdb, 0xe6,
	0x57, 0x83, 0x0e, 0x31, 0xb8, 0x99, 0xc1, 0x51, 0xe4, 0x5e, 0x73, 0xe7, 0x0f, 0x23, 0x75, 0x98,
	0xa1, 0xb8, 0x53, 0xa9, 0x5f, 0xfc, 0xe1, 0xd2, 0x5e, 0xe9, 0x10, 0x47, 0xc4, 0xc6, 0xa3, 0xb4,
	0x53, 0xd8, 0xb8, 0xed, 0xe2, 0x04, 0x07, 0xb0, 0x79, 0x73, 0x82, 0xb4, 0x53, 0xb3, 0x8d, 0xbc,
	0xc0, 0x74, 0xbf, 0xdd, 0x3d, 0xed, 0x56, 0xb3, 0x96, 0x4f, 0xe1, 0x8b, 0x85, 0x25, 0xf0, 0xcb,
	0x79, 0x89, 0xf1, 0xf5, 0x6e, 0x3c, 0x7e, 0x3b, 0xc7, 0xb7, 0xf0, 0x9d, 0x7d, 0xfc, 0x29, 0x00,
	0x00, 0xff, 0xff, 0xed, 0xe2, 0x6d, 0x72, 0x8a, 0x07, 0x00, 0x00,
}
