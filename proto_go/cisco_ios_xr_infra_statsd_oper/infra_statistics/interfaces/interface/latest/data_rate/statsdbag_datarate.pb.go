// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statsdbag_datarate.proto

package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_data_rate

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

// Datarate information
type StatsdbagDatarate_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsdbagDatarate_KEYS) Reset()         { *m = StatsdbagDatarate_KEYS{} }
func (m *StatsdbagDatarate_KEYS) String() string { return proto.CompactTextString(m) }
func (*StatsdbagDatarate_KEYS) ProtoMessage()    {}
func (*StatsdbagDatarate_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_statsdbag_datarate_c2eebaf5966f1a00, []int{0}
}
func (m *StatsdbagDatarate_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsdbagDatarate_KEYS.Unmarshal(m, b)
}
func (m *StatsdbagDatarate_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsdbagDatarate_KEYS.Marshal(b, m, deterministic)
}
func (dst *StatsdbagDatarate_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsdbagDatarate_KEYS.Merge(dst, src)
}
func (m *StatsdbagDatarate_KEYS) XXX_Size() int {
	return xxx_messageInfo_StatsdbagDatarate_KEYS.Size(m)
}
func (m *StatsdbagDatarate_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsdbagDatarate_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_StatsdbagDatarate_KEYS proto.InternalMessageInfo

func (m *StatsdbagDatarate_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type StatsdbagDatarate struct {
	// Input data rate in 1000's of bps
	InputDataRate uint64 `protobuf:"varint,50,opt,name=input_data_rate,json=inputDataRate,proto3" json:"input_data_rate,omitempty"`
	// Input packets per second
	InputPacketRate uint64 `protobuf:"varint,51,opt,name=input_packet_rate,json=inputPacketRate,proto3" json:"input_packet_rate,omitempty"`
	// Output data rate in 1000's of bps
	OutputDataRate uint64 `protobuf:"varint,52,opt,name=output_data_rate,json=outputDataRate,proto3" json:"output_data_rate,omitempty"`
	// Output packets per second
	OutputPacketRate uint64 `protobuf:"varint,53,opt,name=output_packet_rate,json=outputPacketRate,proto3" json:"output_packet_rate,omitempty"`
	// Peak input data rate
	PeakInputDataRate uint64 `protobuf:"varint,54,opt,name=peak_input_data_rate,json=peakInputDataRate,proto3" json:"peak_input_data_rate,omitempty"`
	// Peak input packet rate
	PeakInputPacketRate uint64 `protobuf:"varint,55,opt,name=peak_input_packet_rate,json=peakInputPacketRate,proto3" json:"peak_input_packet_rate,omitempty"`
	// Peak output data rate
	PeakOutputDataRate uint64 `protobuf:"varint,56,opt,name=peak_output_data_rate,json=peakOutputDataRate,proto3" json:"peak_output_data_rate,omitempty"`
	// Peak output packet rate
	PeakOutputPacketRate uint64 `protobuf:"varint,57,opt,name=peak_output_packet_rate,json=peakOutputPacketRate,proto3" json:"peak_output_packet_rate,omitempty"`
	// Bandwidth (in kbps)
	Bandwidth uint32 `protobuf:"varint,58,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Number of 30-sec intervals less one
	LoadInterval uint32 `protobuf:"varint,59,opt,name=load_interval,json=loadInterval,proto3" json:"load_interval,omitempty"`
	// Output load as fraction of 255
	OutputLoad uint32 `protobuf:"varint,60,opt,name=output_load,json=outputLoad,proto3" json:"output_load,omitempty"`
	// Input load as fraction of 255
	InputLoad uint32 `protobuf:"varint,61,opt,name=input_load,json=inputLoad,proto3" json:"input_load,omitempty"`
	// Reliability coefficient
	Reliability          uint32   `protobuf:"varint,62,opt,name=reliability,proto3" json:"reliability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsdbagDatarate) Reset()         { *m = StatsdbagDatarate{} }
func (m *StatsdbagDatarate) String() string { return proto.CompactTextString(m) }
func (*StatsdbagDatarate) ProtoMessage()    {}
func (*StatsdbagDatarate) Descriptor() ([]byte, []int) {
	return fileDescriptor_statsdbag_datarate_c2eebaf5966f1a00, []int{1}
}
func (m *StatsdbagDatarate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsdbagDatarate.Unmarshal(m, b)
}
func (m *StatsdbagDatarate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsdbagDatarate.Marshal(b, m, deterministic)
}
func (dst *StatsdbagDatarate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsdbagDatarate.Merge(dst, src)
}
func (m *StatsdbagDatarate) XXX_Size() int {
	return xxx_messageInfo_StatsdbagDatarate.Size(m)
}
func (m *StatsdbagDatarate) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsdbagDatarate.DiscardUnknown(m)
}

var xxx_messageInfo_StatsdbagDatarate proto.InternalMessageInfo

func (m *StatsdbagDatarate) GetInputDataRate() uint64 {
	if m != nil {
		return m.InputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetInputPacketRate() uint64 {
	if m != nil {
		return m.InputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputDataRate() uint64 {
	if m != nil {
		return m.OutputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputPacketRate() uint64 {
	if m != nil {
		return m.OutputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakInputDataRate() uint64 {
	if m != nil {
		return m.PeakInputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakInputPacketRate() uint64 {
	if m != nil {
		return m.PeakInputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakOutputDataRate() uint64 {
	if m != nil {
		return m.PeakOutputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakOutputPacketRate() uint64 {
	if m != nil {
		return m.PeakOutputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *StatsdbagDatarate) GetLoadInterval() uint32 {
	if m != nil {
		return m.LoadInterval
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputLoad() uint32 {
	if m != nil {
		return m.OutputLoad
	}
	return 0
}

func (m *StatsdbagDatarate) GetInputLoad() uint32 {
	if m != nil {
		return m.InputLoad
	}
	return 0
}

func (m *StatsdbagDatarate) GetReliability() uint32 {
	if m != nil {
		return m.Reliability
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsdbagDatarate_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.latest.data_rate.statsdbag_datarate_KEYS")
	proto.RegisterType((*StatsdbagDatarate)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.latest.data_rate.statsdbag_datarate")
}

func init() {
	proto.RegisterFile("statsdbag_datarate.proto", fileDescriptor_statsdbag_datarate_c2eebaf5966f1a00)
}

var fileDescriptor_statsdbag_datarate_c2eebaf5966f1a00 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x09, 0x14, 0xa1, 0xb7, 0xb6, 0xda, 0xb1, 0xda, 0x2c, 0x14, 0x43, 0x45, 0x09, 0x22,
	0x11, 0xad, 0xf5, 0xbf, 0xe2, 0x42, 0x17, 0x45, 0x51, 0xa9, 0xb8, 0x70, 0x35, 0xdc, 0x24, 0x53,
	0xdf, 0xd0, 0x34, 0x13, 0x92, 0xdb, 0xf7, 0xe7, 0x83, 0xbd, 0xef, 0xf7, 0x98, 0x3b, 0x21, 0x9d,
	0xbe, 0xee, 0xca, 0x39, 0xbf, 0x73, 0xce, 0x5c, 0x1a, 0x08, 0x1b, 0x42, 0x6a, 0xf2, 0x14, 0xff,
	0xcb, 0x1c, 0x09, 0x6b, 0x24, 0x95, 0x54, 0xb5, 0x21, 0x23, 0xfe, 0x66, 0xba, 0xc9, 0x8c, 0xd4,
	0xa6, 0x91, 0xe7, 0xb5, 0xd4, 0xe5, 0xba, 0x46, 0xe9, 0x60, 0x69, 0x2a, 0x55, 0x27, 0x7b, 0x45,
	0x37, 0xa4, 0xb3, 0x26, 0xd1, 0x25, 0xa9, 0x7a, 0x8d, 0x99, 0xf2, 0x7e, 0x26, 0x05, 0x92, 0x6a,
	0x28, 0xb1, 0xdd, 0xd2, 0x96, 0xcf, 0xbe, 0xc0, 0xf4, 0x78, 0x52, 0x7e, 0xff, 0xf6, 0xef, 0x8f,
	0x78, 0x0c, 0xa3, 0x2e, 0x28, 0x4b, 0xdc, 0xaa, 0x30, 0x88, 0x82, 0xb8, 0xbf, 0x1a, 0x76, 0xea,
	0x4f, 0xdc, 0xaa, 0xd9, 0x65, 0x0f, 0xc4, 0x71, 0x85, 0x78, 0x02, 0xb7, 0x74, 0x59, 0xed, 0x48,
	0x76, 0x5b, 0xe1, 0xcb, 0x28, 0x88, 0x7b, 0x36, 0x5e, 0xed, 0xe8, 0x2b, 0x12, 0xae, 0x2c, 0xf7,
	0x14, 0xc6, 0x8e, 0xab, 0x30, 0xdb, 0x28, 0x72, 0xe4, 0x9c, 0x49, 0x57, 0xf0, 0x9b, 0x75, 0x66,
	0x63, 0xb8, 0x6d, 0x76, 0x74, 0x58, 0xfa, 0x8a, 0xd1, 0x91, 0xd3, 0xbb, 0xd6, 0x67, 0x20, 0x5a,
	0xd2, 0xaf, 0x5d, 0x30, 0xdb, 0x76, 0x78, 0xbd, 0xcf, 0x61, 0x52, 0x29, 0xdc, 0xc8, 0xeb, 0x0f,
	0x7e, 0xcd, 0xfc, 0xd8, 0x7a, 0xcb, 0x83, 0x47, 0xcf, 0xe1, 0x9e, 0x17, 0xf0, 0x27, 0xde, 0x70,
	0xe4, 0x4e, 0x17, 0xf1, 0x56, 0x5e, 0xc0, 0x5d, 0x0e, 0x1d, 0x9d, 0xf0, 0x96, 0x33, 0xc2, 0x9a,
	0xbf, 0x0e, 0xcf, 0x58, 0xc0, 0xd4, 0x8f, 0xf8, 0x43, 0xef, 0x38, 0x34, 0xd9, 0x87, 0xbc, 0xa5,
	0xfb, 0xd0, 0x4f, 0xb1, 0xcc, 0xcf, 0x74, 0x4e, 0x27, 0xe1, 0xfb, 0x28, 0x88, 0x87, 0xab, 0xbd,
	0x20, 0x1e, 0xc1, 0xb0, 0x30, 0x98, 0x4b, 0xfe, 0x1b, 0x4f, 0xb1, 0x08, 0x3f, 0x30, 0x71, 0xd3,
	0x8a, 0xcb, 0x56, 0x13, 0x0f, 0x61, 0xd0, 0x8e, 0x5a, 0x39, 0xfc, 0xc8, 0x08, 0x38, 0xe9, 0x87,
	0xc1, 0x5c, 0x3c, 0x00, 0x70, 0xd7, 0xb3, 0xff, 0xc9, 0x8d, 0xb0, 0xc2, 0x76, 0x04, 0x83, 0x5a,
	0x15, 0x1a, 0x53, 0x5d, 0x68, 0xba, 0x08, 0x3f, 0xb3, 0xef, 0x4b, 0xe9, 0x0d, 0xfe, 0xae, 0xe7,
	0x57, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0xf0, 0xef, 0x94, 0xf3, 0x02, 0x00, 0x00,
}
