// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ifstatsbag_proto.proto

package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol

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
	return fileDescriptor_ifstatsbag_proto_627af1be7f507808, []int{0}
}
func (m *IfstatsbagProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Unmarshal(m, b)
}
func (m *IfstatsbagProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagProto_KEYS.Marshal(b, m, deterministic)
}
func (dst *IfstatsbagProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagProto_KEYS.Merge(dst, src)
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
	return fileDescriptor_ifstatsbag_proto_627af1be7f507808, []int{1}
}
func (m *IfstatsbagProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfstatsbagProto.Unmarshal(m, b)
}
func (m *IfstatsbagProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfstatsbagProto.Marshal(b, m, deterministic)
}
func (dst *IfstatsbagProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfstatsbagProto.Merge(dst, src)
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
	proto.RegisterType((*IfstatsbagProto_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.protocols.protocol.ifstatsbag_proto_KEYS")
	proto.RegisterType((*IfstatsbagProto)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.protocols.protocol.ifstatsbag_proto")
}

func init() {
	proto.RegisterFile("ifstatsbag_proto.proto", fileDescriptor_ifstatsbag_proto_627af1be7f507808)
}

var fileDescriptor_ifstatsbag_proto_627af1be7f507808 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4b, 0x4b, 0xc3, 0x40,
	0x10, 0x80, 0x89, 0x8a, 0xd8, 0xb1, 0x69, 0xeb, 0x82, 0x12, 0x04, 0xa1, 0xd6, 0x07, 0x51, 0xa4,
	0x07, 0xeb, 0xdb, 0xab, 0x9e, 0x04, 0x91, 0x54, 0x10, 0x4f, 0xcb, 0x76, 0x3b, 0x95, 0xc5, 0xe6,
	0x41, 0x76, 0x2a, 0xfa, 0x3f, 0xfc, 0xc1, 0x92, 0xd9, 0x24, 0xad, 0xbd, 0x84, 0xcd, 0x97, 0x2f,
	0xdf, 0xec, 0xc0, 0x8e, 0x99, 0x58, 0x52, 0x64, 0x47, 0xea, 0x43, 0x66, 0x79, 0x4a, 0x69, 0x9f,
	0x9f, 0xe2, 0x4d, 0x1b, 0xab, 0x53, 0x69, 0x52, 0x2b, 0xbf, 0x73, 0x69, 0x92, 0x49, 0xae, 0x24,
	0x8b, 0x63, 0x99, 0x66, 0x98, 0xf7, 0xe7, 0xc4, 0x58, 0x32, 0xda, 0xf6, 0x4d, 0x42, 0x98, 0x4f,
	0x94, 0xc6, 0x85, 0xa3, 0x6b, 0xe9, 0x74, 0x6a, 0xeb, 0x53, 0x4f, 0xc3, 0xf6, 0xf2, 0x48, 0xf9,
	0xf4, 0xf8, 0x3e, 0x14, 0x47, 0xd0, 0xaa, 0x7f, 0x94, 0x89, 0x8a, 0x31, 0xf0, 0xba, 0x5e, 0xd8,
	0x88, 0xfc, 0x9a, 0x3e, 0xab, 0x18, 0xc5, 0x01, 0xf8, 0x55, 0xcb, 0x59, 0x2b, 0x6c, 0x35, 0x2b,
	0x58, 0x48, 0xbd, 0xdf, 0x55, 0xe8, 0x2c, 0x4f, 0x29, 0x06, 0x8c, 0x7e, 0x08, 0xad, 0xcc, 0x51,
	0xa3, 0xf9, 0xc2, 0x71, 0x70, 0xde, 0xf5, 0xc2, 0xb5, 0xc8, 0x67, 0x1a, 0x95, 0x50, 0x9c, 0x40,
	0x27, 0x53, 0xfa, 0x13, 0x69, 0x41, 0x1c, 0xb0, 0xd8, 0x2e, 0x79, 0xad, 0xee, 0x01, 0xb8, 0xa2,
	0xc5, 0x84, 0x82, 0x0b, 0x96, 0x1a, 0x4c, 0x86, 0x98, 0x90, 0xd8, 0x87, 0x66, 0x55, 0x62, 0xe1,
	0x92, 0x85, 0xcd, 0x92, 0xb1, 0xb2, 0x0b, 0x1b, 0xd5, 0xc5, 0x83, 0xab, 0xae, 0x17, 0xfa, 0x51,
	0xfd, 0x2e, 0x0e, 0xa1, 0x35, 0x55, 0x96, 0xe4, 0x58, 0x91, 0x92, 0x64, 0x62, 0x0c, 0xae, 0xd9,
	0x68, 0x16, 0xf4, 0x41, 0x91, 0x7a, 0x35, 0x31, 0x8a, 0x63, 0x68, 0x9b, 0x24, 0x9b, 0x95, 0x5a,
	0xae, 0x08, 0x83, 0x1b, 0xb7, 0x16, 0xe3, 0xc2, 0x8b, 0x14, 0xa1, 0x38, 0x85, 0x2d, 0xe7, 0xb9,
	0xf1, 0xce, 0xbc, 0x75, 0x7b, 0xf1, 0x87, 0x17, 0xe6, 0xec, 0x86, 0xd0, 0x49, 0x67, 0xf4, 0x3f,
	0x7a, 0xc7, 0x6a, 0xcb, 0xf1, 0xba, 0x7a, 0x06, 0xa2, 0x34, 0x17, 0xb3, 0xf7, 0xec, 0x96, 0x8d,
	0x79, 0x77, 0xb4, 0xce, 0xbb, 0x0d, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x22, 0x97, 0x62,
	0x75, 0x02, 0x00, 0x00,
}
