// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lldp_stats.proto

package cisco_ios_xr_ethernet_lldp_oper_lldp_nodes_node_statistics

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

// LLDP statistics
type LldpStats_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpStats_KEYS) Reset()         { *m = LldpStats_KEYS{} }
func (m *LldpStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*LldpStats_KEYS) ProtoMessage()    {}
func (*LldpStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_lldp_stats_7792e58a45181ed6, []int{0}
}
func (m *LldpStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpStats_KEYS.Unmarshal(m, b)
}
func (m *LldpStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpStats_KEYS.Marshal(b, m, deterministic)
}
func (dst *LldpStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpStats_KEYS.Merge(dst, src)
}
func (m *LldpStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_LldpStats_KEYS.Size(m)
}
func (m *LldpStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LldpStats_KEYS proto.InternalMessageInfo

func (m *LldpStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type LldpStats struct {
	// Transmitted packets
	TransmittedPackets uint32 `protobuf:"varint,50,opt,name=transmitted_packets,json=transmittedPackets,proto3" json:"transmitted_packets,omitempty"`
	// Aged out entries
	AgedOutEntries uint32 `protobuf:"varint,51,opt,name=aged_out_entries,json=agedOutEntries,proto3" json:"aged_out_entries,omitempty"`
	// Discarded packets
	DiscardedPackets uint32 `protobuf:"varint,52,opt,name=discarded_packets,json=discardedPackets,proto3" json:"discarded_packets,omitempty"`
	// Bad packet received and dropped
	BadPackets uint32 `protobuf:"varint,53,opt,name=bad_packets,json=badPackets,proto3" json:"bad_packets,omitempty"`
	// Received packets
	ReceivedPackets uint32 `protobuf:"varint,54,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	// Discarded TLVs
	DiscardedTlVs uint32 `protobuf:"varint,55,opt,name=discarded_tl_vs,json=discardedTlVs,proto3" json:"discarded_tl_vs,omitempty"`
	// Unrecognized TLVs
	UnrecognizedTlVs uint32 `protobuf:"varint,56,opt,name=unrecognized_tl_vs,json=unrecognizedTlVs,proto3" json:"unrecognized_tl_vs,omitempty"`
	// Out-of-memory conditions
	OutOfMemoryErrors uint32 `protobuf:"varint,57,opt,name=out_of_memory_errors,json=outOfMemoryErrors,proto3" json:"out_of_memory_errors,omitempty"`
	// Transmission errors
	EncapsulationErrors uint32 `protobuf:"varint,58,opt,name=encapsulation_errors,json=encapsulationErrors,proto3" json:"encapsulation_errors,omitempty"`
	// Queue overflows
	QueueOverflowErrors uint32 `protobuf:"varint,59,opt,name=queue_overflow_errors,json=queueOverflowErrors,proto3" json:"queue_overflow_errors,omitempty"`
	// Table overflows
	TableOverflowErrors  uint32   `protobuf:"varint,60,opt,name=table_overflow_errors,json=tableOverflowErrors,proto3" json:"table_overflow_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpStats) Reset()         { *m = LldpStats{} }
func (m *LldpStats) String() string { return proto.CompactTextString(m) }
func (*LldpStats) ProtoMessage()    {}
func (*LldpStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_lldp_stats_7792e58a45181ed6, []int{1}
}
func (m *LldpStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpStats.Unmarshal(m, b)
}
func (m *LldpStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpStats.Marshal(b, m, deterministic)
}
func (dst *LldpStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpStats.Merge(dst, src)
}
func (m *LldpStats) XXX_Size() int {
	return xxx_messageInfo_LldpStats.Size(m)
}
func (m *LldpStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpStats.DiscardUnknown(m)
}

var xxx_messageInfo_LldpStats proto.InternalMessageInfo

func (m *LldpStats) GetTransmittedPackets() uint32 {
	if m != nil {
		return m.TransmittedPackets
	}
	return 0
}

func (m *LldpStats) GetAgedOutEntries() uint32 {
	if m != nil {
		return m.AgedOutEntries
	}
	return 0
}

func (m *LldpStats) GetDiscardedPackets() uint32 {
	if m != nil {
		return m.DiscardedPackets
	}
	return 0
}

func (m *LldpStats) GetBadPackets() uint32 {
	if m != nil {
		return m.BadPackets
	}
	return 0
}

func (m *LldpStats) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *LldpStats) GetDiscardedTlVs() uint32 {
	if m != nil {
		return m.DiscardedTlVs
	}
	return 0
}

func (m *LldpStats) GetUnrecognizedTlVs() uint32 {
	if m != nil {
		return m.UnrecognizedTlVs
	}
	return 0
}

func (m *LldpStats) GetOutOfMemoryErrors() uint32 {
	if m != nil {
		return m.OutOfMemoryErrors
	}
	return 0
}

func (m *LldpStats) GetEncapsulationErrors() uint32 {
	if m != nil {
		return m.EncapsulationErrors
	}
	return 0
}

func (m *LldpStats) GetQueueOverflowErrors() uint32 {
	if m != nil {
		return m.QueueOverflowErrors
	}
	return 0
}

func (m *LldpStats) GetTableOverflowErrors() uint32 {
	if m != nil {
		return m.TableOverflowErrors
	}
	return 0
}

func init() {
	proto.RegisterType((*LldpStats_KEYS)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.statistics.lldp_stats_KEYS")
	proto.RegisterType((*LldpStats)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.statistics.lldp_stats")
}

func init() { proto.RegisterFile("lldp_stats.proto", fileDescriptor_lldp_stats_7792e58a45181ed6) }

var fileDescriptor_lldp_stats_7792e58a45181ed6 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4b, 0xab, 0xd3, 0x40,
	0x14, 0x80, 0x29, 0x88, 0xd8, 0x23, 0xb5, 0xe9, 0xb4, 0x42, 0xc0, 0x85, 0xa5, 0x0b, 0xa9, 0x28,
	0x29, 0xb6, 0x3e, 0xab, 0xdb, 0xae, 0x44, 0x2b, 0x55, 0x04, 0x57, 0xc3, 0x64, 0x72, 0x5a, 0x87,
	0x9b, 0xcc, 0xe4, 0xce, 0x9c, 0xf4, 0x3e, 0x96, 0xf7, 0x97, 0x5f, 0x32, 0x69, 0x1e, 0xe5, 0x6e,
	0x42, 0x38, 0xdf, 0xf7, 0x9d, 0x64, 0x60, 0x20, 0x48, 0xd3, 0x24, 0xe7, 0x8e, 0x04, 0xb9, 0x28,
	0xb7, 0x86, 0x0c, 0x5b, 0x4b, 0xe5, 0xa4, 0xe1, 0xca, 0x38, 0x7e, 0x6d, 0x39, 0xd2, 0x7f, 0xb4,
	0x1a, 0x89, 0x7b, 0xcf, 0xe4, 0x68, 0xa3, 0xf2, 0x2d, 0xd2, 0x26, 0x41, 0xe7, 0x9f, 0x51, 0x19,
	0x2b, 0x47, 0x4a, 0xba, 0x59, 0x04, 0xc3, 0x76, 0x1f, 0xff, 0xbe, 0xf9, 0xf7, 0x9b, 0xbd, 0x80,
	0x7e, 0x69, 0x71, 0x2d, 0x32, 0x0c, 0x7b, 0xd3, 0xde, 0xbc, 0xbf, 0x7b, 0x52, 0x0e, 0x7e, 0x8a,
	0x0c, 0x67, 0x77, 0x8f, 0x00, 0xda, 0x80, 0x2d, 0x60, 0x4c, 0x56, 0x68, 0x97, 0x29, 0x22, 0x4c,
	0x78, 0x2e, 0xe4, 0x05, 0x92, 0x0b, 0x97, 0xd3, 0xde, 0x7c, 0xb0, 0x63, 0x1d, 0xf4, 0xab, 0x22,
	0x6c, 0x0e, 0x81, 0x38, 0x60, 0xc2, 0x4d, 0x41, 0x1c, 0x35, 0x59, 0x85, 0x2e, 0x5c, 0x79, 0xfb,
	0x59, 0x39, 0xdf, 0x16, 0xb4, 0xa9, 0xa6, 0xec, 0x0d, 0x8c, 0x12, 0xe5, 0xa4, 0xb0, 0x49, 0x67,
	0xf1, 0x7b, 0xaf, 0x06, 0x0d, 0xa8, 0xd7, 0xbe, 0x84, 0xa7, 0xb1, 0x68, 0xb5, 0x0f, 0x5e, 0x83,
	0x58, 0x34, 0xc2, 0x6b, 0x08, 0x2c, 0x4a, 0x54, 0xc7, 0xce, 0xb2, 0x8f, 0xde, 0x1a, 0xd6, 0xf3,
	0x5a, 0x7d, 0x05, 0xc3, 0xf6, 0xc3, 0x94, 0xf2, 0xa3, 0x0b, 0x3f, 0x79, 0x73, 0xd0, 0x8c, 0xff,
	0xa4, 0x7f, 0x1d, 0x7b, 0x0b, 0xac, 0xd0, 0x16, 0xa5, 0x39, 0x68, 0x75, 0xdb, 0xa8, 0x9f, 0xab,
	0x3f, 0xec, 0x12, 0x6f, 0x2f, 0x60, 0x52, 0x9e, 0xd9, 0xec, 0x79, 0x86, 0x99, 0xb1, 0x37, 0x1c,
	0xad, 0x35, 0xd6, 0x85, 0x5f, 0xbc, 0x3f, 0x32, 0x05, 0x6d, 0xf7, 0x3f, 0x3c, 0xd9, 0x78, 0xc0,
	0xde, 0xc1, 0x04, 0xb5, 0x14, 0xb9, 0x2b, 0x52, 0x41, 0xca, 0xe8, 0x3a, 0x58, 0xfb, 0x60, 0x7c,
	0xc6, 0x4e, 0xc9, 0x12, 0x9e, 0x5f, 0x16, 0x58, 0x20, 0x37, 0x47, 0xb4, 0xfb, 0xd4, 0x5c, 0xd5,
	0xcd, 0xd7, 0xaa, 0xf1, 0x70, 0x7b, 0x62, 0x6d, 0x43, 0x22, 0x4e, 0x1f, 0x36, 0xdf, 0xaa, 0xc6,
	0xc3, 0xf3, 0x26, 0x7e, 0xec, 0xef, 0xdd, 0xea, 0x3e, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x85, 0x4f,
	0xa7, 0x8b, 0x02, 0x00, 0x00,
}
