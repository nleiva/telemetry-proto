// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ifstatsbag_proto.proto

package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_cache_protocols_protocol

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

// Interface counters per protocol
type IfstatsbagProto_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfstatsbagProto_KEYS) Reset()         { *m = IfstatsbagProto_KEYS{} }
func (m *IfstatsbagProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*IfstatsbagProto_KEYS) ProtoMessage()    {}
func (*IfstatsbagProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e305d2d18a94ee6, []int{0}
}

func (m *IfstatsbagProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Unmarshal(m, b)
}
func (m *IfstatsbagProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Marshal(b, m, deterministic)
}
func (m *IfstatsbagProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagProto_KEYS.Merge(m, src)
}
func (m *IfstatsbagProto_KEYS) XXX_Size() int {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Size(m)
}
func (m *IfstatsbagProto_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IfstatsbagProto_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IfstatsbagProto_KEYS proto.InternalMessageInfo

func (m *IfstatsbagProto_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IfstatsbagProto_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

type IfstatsbagProto struct {
	// Bytes received
	BytesReceived uint64 `protobuf:"varint,50,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	// Packets received
	PacketsReceived uint64 `protobuf:"varint,51,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	// Bytes sent
	BytesSent uint64 `protobuf:"varint,52,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	// Packets sent
	PacketsSent uint64 `protobuf:"varint,53,opt,name=packets_sent,json=packetsSent,proto3" json:"packets_sent,omitempty"`
	// Protocol number
	Protocol uint32 `protobuf:"varint,54,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Time when counters were last written (in seconds)
	LastDataTime uint32 `protobuf:"varint,55,opt,name=last_data_time,json=lastDataTime,proto3" json:"last_data_time,omitempty"`
	// Input data rate in 1000's of bps
	InputDataRate uint64 `protobuf:"varint,56,opt,name=input_data_rate,json=inputDataRate,proto3" json:"input_data_rate,omitempty"`
	// Input packets per second
	InputPacketRate uint64 `protobuf:"varint,57,opt,name=input_packet_rate,json=inputPacketRate,proto3" json:"input_packet_rate,omitempty"`
	// Output data rate in 1000's of bps
	OutputDataRate uint64 `protobuf:"varint,58,opt,name=output_data_rate,json=outputDataRate,proto3" json:"output_data_rate,omitempty"`
	// Output packets per second
	OutputPacketRate     uint64   `protobuf:"varint,59,opt,name=output_packet_rate,json=outputPacketRate,proto3" json:"output_packet_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfstatsbagProto) Reset()         { *m = IfstatsbagProto{} }
func (m *IfstatsbagProto) String() string { return proto.CompactTextString(m) }
func (*IfstatsbagProto) ProtoMessage()    {}
func (*IfstatsbagProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e305d2d18a94ee6, []int{1}
}

func (m *IfstatsbagProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagProto.Unmarshal(m, b)
}
func (m *IfstatsbagProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagProto.Marshal(b, m, deterministic)
}
func (m *IfstatsbagProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagProto.Merge(m, src)
}
func (m *IfstatsbagProto) XXX_Size() int {
	return xxx_messageInfo_IfstatsbagProto.Size(m)
}
func (m *IfstatsbagProto) XXX_DiscardUnknown() {
	xxx_messageInfo_IfstatsbagProto.DiscardUnknown(m)
}

var xxx_messageInfo_IfstatsbagProto proto.InternalMessageInfo

func (m *IfstatsbagProto) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *IfstatsbagProto) GetPacketsReceived() uint64 {
	if m != nil {
		return m.PacketsReceived
	}
	return 0
}

func (m *IfstatsbagProto) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *IfstatsbagProto) GetPacketsSent() uint64 {
	if m != nil {
		return m.PacketsSent
	}
	return 0
}

func (m *IfstatsbagProto) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *IfstatsbagProto) GetLastDataTime() uint32 {
	if m != nil {
		return m.LastDataTime
	}
	return 0
}

func (m *IfstatsbagProto) GetInputDataRate() uint64 {
	if m != nil {
		return m.InputDataRate
	}
	return 0
}

func (m *IfstatsbagProto) GetInputPacketRate() uint64 {
	if m != nil {
		return m.InputPacketRate
	}
	return 0
}

func (m *IfstatsbagProto) GetOutputDataRate() uint64 {
	if m != nil {
		return m.OutputDataRate
	}
	return 0
}

func (m *IfstatsbagProto) GetOutputPacketRate() uint64 {
	if m != nil {
		return m.OutputPacketRate
	}
	return 0
}

func init() {
	proto.RegisterType((*IfstatsbagProto_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.cache.protocols.protocol.ifstatsbag_proto_KEYS")
	proto.RegisterType((*IfstatsbagProto)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.cache.protocols.protocol.ifstatsbag_proto")
}

func init() { proto.RegisterFile("ifstatsbag_proto.proto", fileDescriptor_4e305d2d18a94ee6) }

var fileDescriptor_4e305d2d18a94ee6 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4b, 0x4b, 0xeb, 0x40,
	0x14, 0x80, 0xc9, 0xbd, 0x97, 0xcb, 0xed, 0xdc, 0xa6, 0xad, 0x03, 0x4a, 0x10, 0x84, 0x5a, 0x1f,
	0x44, 0x91, 0x2c, 0xac, 0x6f, 0xb7, 0xba, 0x12, 0x44, 0x52, 0x37, 0x2e, 0x64, 0x38, 0x9d, 0x9e,
	0xea, 0x60, 0xf3, 0x20, 0x73, 0x2a, 0xfa, 0x3f, 0xfc, 0xc1, 0x92, 0x33, 0x49, 0x5a, 0xbb, 0x09,
	0x93, 0x2f, 0x5f, 0xbe, 0x33, 0x47, 0x6c, 0x98, 0xa9, 0x25, 0x20, 0x3b, 0x86, 0x17, 0x95, 0x17,
	0x19, 0x65, 0x11, 0x3f, 0xe5, 0xb3, 0x36, 0x56, 0x67, 0xca, 0x64, 0x56, 0x7d, 0x14, 0xca, 0xa4,
	0xd3, 0x02, 0x14, 0x8b, 0x13, 0x95, 0xe5, 0x58, 0x44, 0x0b, 0x62, 0x2c, 0x19, 0x6d, 0x23, 0x93,
	0x12, 0x16, 0x53, 0xd0, 0xb8, 0x74, 0x8c, 0x34, 0xe8, 0x57, 0x74, 0x45, 0x9d, 0xcd, 0x6c, 0x73,
	0x1a, 0x68, 0xb1, 0xbe, 0x3a, 0x58, 0xdd, 0xdd, 0x3e, 0x8d, 0xe4, 0x9e, 0xe8, 0x34, 0xbf, 0xab,
	0x14, 0x12, 0x0c, 0xbc, 0xbe, 0x17, 0xb6, 0x62, 0xbf, 0xa1, 0xf7, 0x90, 0xa0, 0xdc, 0x11, 0x7e,
	0xdd, 0x72, 0xd6, 0x2f, 0xb6, 0xda, 0x35, 0x2c, 0xa5, 0xc1, 0xd7, 0x6f, 0xd1, 0x5b, 0x9d, 0x52,
	0x0e, 0x18, 0x7f, 0x12, 0x5a, 0x55, 0xa0, 0x46, 0xf3, 0x8e, 0x93, 0xe0, 0xb8, 0xef, 0x85, 0x7f,
	0x62, 0x9f, 0x69, 0x5c, 0x41, 0x79, 0x20, 0x7a, 0x39, 0xe8, 0x37, 0xa4, 0x25, 0x71, 0xc8, 0x62,
	0xb7, 0xe2, 0x8d, 0xba, 0x25, 0x84, 0x2b, 0x5a, 0x4c, 0x29, 0x38, 0x61, 0xa9, 0xc5, 0x64, 0x84,
	0x29, 0xc9, 0x6d, 0xd1, 0xae, 0x4b, 0x2c, 0x9c, 0xb2, 0xf0, 0xbf, 0x62, 0xac, 0x6c, 0x8a, 0x7f,
	0xf5, 0xc5, 0x83, 0xb3, 0xbe, 0x17, 0xfa, 0x71, 0xf3, 0x2e, 0x77, 0x45, 0x67, 0x06, 0x96, 0xd4,
	0x04, 0x08, 0x14, 0x99, 0x04, 0x83, 0x73, 0x36, 0xda, 0x25, 0xbd, 0x01, 0x82, 0x47, 0x93, 0xa0,
	0xdc, 0x17, 0x5d, 0x93, 0xe6, 0xf3, 0x4a, 0x2b, 0x80, 0x30, 0xb8, 0x70, 0x6b, 0x31, 0x2e, 0xbd,
	0x18, 0x08, 0xe5, 0xa1, 0x58, 0x73, 0x9e, 0x1b, 0xef, 0xcc, 0x4b, 0xb7, 0x17, 0x7f, 0x78, 0x60,
	0xce, 0x6e, 0x28, 0x7a, 0xd9, 0x9c, 0x7e, 0x46, 0xaf, 0x58, 0xed, 0x38, 0xde, 0x54, 0x8f, 0x84,
	0xac, 0xcc, 0xe5, 0xec, 0x35, 0xbb, 0x55, 0x63, 0xd1, 0x1d, 0xff, 0xe5, 0xdd, 0x86, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0xd8, 0x28, 0x98, 0x7b, 0x02, 0x00, 0x00,
}
