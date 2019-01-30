// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statsdbag_datarate.proto

/*
Package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_cache_data_rate is a generated protocol buffer package.

It is generated from these files:
	statsdbag_datarate.proto

It has these top-level messages:
	StatsdbagDatarate_KEYS
	StatsdbagDatarate
*/
package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_cache_data_rate

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
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *StatsdbagDatarate_KEYS) Reset()                    { *m = StatsdbagDatarate_KEYS{} }
func (m *StatsdbagDatarate_KEYS) String() string            { return proto.CompactTextString(m) }
func (*StatsdbagDatarate_KEYS) ProtoMessage()               {}
func (*StatsdbagDatarate_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatsdbagDatarate_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type StatsdbagDatarate struct {
	// Input data rate in 1000's of bps
	InputDataRate uint64 `protobuf:"varint,50,opt,name=input_data_rate,json=inputDataRate" json:"input_data_rate,omitempty"`
	// Input packets per second
	InputPacketRate uint64 `protobuf:"varint,51,opt,name=input_packet_rate,json=inputPacketRate" json:"input_packet_rate,omitempty"`
	// Output data rate in 1000's of bps
	OutputDataRate uint64 `protobuf:"varint,52,opt,name=output_data_rate,json=outputDataRate" json:"output_data_rate,omitempty"`
	// Output packets per second
	OutputPacketRate uint64 `protobuf:"varint,53,opt,name=output_packet_rate,json=outputPacketRate" json:"output_packet_rate,omitempty"`
	// Peak input data rate
	PeakInputDataRate uint64 `protobuf:"varint,54,opt,name=peak_input_data_rate,json=peakInputDataRate" json:"peak_input_data_rate,omitempty"`
	// Peak input packet rate
	PeakInputPacketRate uint64 `protobuf:"varint,55,opt,name=peak_input_packet_rate,json=peakInputPacketRate" json:"peak_input_packet_rate,omitempty"`
	// Peak output data rate
	PeakOutputDataRate uint64 `protobuf:"varint,56,opt,name=peak_output_data_rate,json=peakOutputDataRate" json:"peak_output_data_rate,omitempty"`
	// Peak output packet rate
	PeakOutputPacketRate uint64 `protobuf:"varint,57,opt,name=peak_output_packet_rate,json=peakOutputPacketRate" json:"peak_output_packet_rate,omitempty"`
	// Bandwidth (in kbps)
	Bandwidth uint32 `protobuf:"varint,58,opt,name=bandwidth" json:"bandwidth,omitempty"`
	// Number of 30-sec intervals less one
	LoadInterval uint32 `protobuf:"varint,59,opt,name=load_interval,json=loadInterval" json:"load_interval,omitempty"`
	// Output load as fraction of 255
	OutputLoad uint32 `protobuf:"varint,60,opt,name=output_load,json=outputLoad" json:"output_load,omitempty"`
	// Input load as fraction of 255
	InputLoad uint32 `protobuf:"varint,61,opt,name=input_load,json=inputLoad" json:"input_load,omitempty"`
	// Reliability coefficient
	Reliability uint32 `protobuf:"varint,62,opt,name=reliability" json:"reliability,omitempty"`
}

func (m *StatsdbagDatarate) Reset()                    { *m = StatsdbagDatarate{} }
func (m *StatsdbagDatarate) String() string            { return proto.CompactTextString(m) }
func (*StatsdbagDatarate) ProtoMessage()               {}
func (*StatsdbagDatarate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	proto.RegisterType((*StatsdbagDatarate_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.cache.data_rate.statsdbag_datarate_KEYS")
	proto.RegisterType((*StatsdbagDatarate)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.cache.data_rate.statsdbag_datarate")
}

func init() { proto.RegisterFile("statsdbag_datarate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x6b, 0xdb, 0x40,
	0x14, 0xc4, 0x11, 0x98, 0x82, 0x9f, 0x6b, 0xb7, 0xde, 0xba, 0xb5, 0x0e, 0x2d, 0x15, 0x2e, 0x2d,
	0xa2, 0x14, 0x95, 0xd6, 0x75, 0xff, 0x27, 0xe4, 0x90, 0x1c, 0x4c, 0x42, 0x12, 0x9c, 0x5c, 0x72,
	0x5a, 0x9e, 0xa4, 0x75, 0xbc, 0x58, 0xd6, 0x0a, 0x69, 0x9d, 0x3f, 0x1f, 0x2c, 0xdf, 0x2f, 0xec,
	0x5b, 0x21, 0xaf, 0xe3, 0x9b, 0x99, 0xf9, 0xcd, 0xcc, 0x3e, 0x2c, 0xf0, 0x2b, 0x8d, 0xba, 0x4a,
	0x63, 0xbc, 0xe6, 0x29, 0x6a, 0x2c, 0x51, 0x8b, 0xa8, 0x28, 0x95, 0x56, 0xec, 0x32, 0x91, 0x55,
	0xa2, 0xb8, 0x54, 0x15, 0xbf, 0x2b, 0xb9, 0xcc, 0xe7, 0x25, 0x72, 0x0b, 0x73, 0x55, 0x88, 0x32,
	0xda, 0x28, 0xb2, 0xd2, 0x32, 0xa9, 0x22, 0x99, 0x6b, 0x51, 0xce, 0x31, 0x11, 0xce, 0xcf, 0x28,
	0xc1, 0x64, 0x21, 0x22, 0x53, 0xcd, 0x4d, 0xf7, 0xe8, 0x00, 0x86, 0xbb, 0x8b, 0xfc, 0xf8, 0xe8,
	0xea, 0x82, 0x7d, 0x84, 0x5e, 0x93, 0xe3, 0x39, 0xae, 0x84, 0xef, 0x05, 0x5e, 0xd8, 0x9e, 0x75,
	0x1b, 0xf5, 0x14, 0x57, 0x62, 0xf4, 0xd0, 0x02, 0xb6, 0x5b, 0xc1, 0x3e, 0xc1, 0x0b, 0x99, 0x17,
	0x6b, 0xcd, 0x9b, 0x2d, 0xff, 0x7b, 0xe0, 0x85, 0x2d, 0x13, 0x2f, 0xd6, 0xfa, 0x10, 0x35, 0xce,
	0x0c, 0xf7, 0x19, 0xfa, 0x96, 0x2b, 0x30, 0x59, 0x0a, 0x6d, 0xc9, 0x31, 0x91, 0xb6, 0xe0, 0x9c,
	0x74, 0x62, 0x43, 0x78, 0xa9, 0xd6, 0x7a, 0xbb, 0xf4, 0x07, 0xa1, 0x3d, 0xab, 0x37, 0xad, 0x5f,
	0x80, 0xd5, 0xa4, 0x5b, 0x3b, 0x21, 0xb6, 0xee, 0x70, 0x7a, 0xbf, 0xc2, 0xa0, 0x10, 0xb8, 0xe4,
	0x4f, 0x1f, 0xfc, 0x93, 0xf8, 0xbe, 0xf1, 0xa6, 0x5b, 0x8f, 0x1e, 0xc3, 0x1b, 0x27, 0xe0, 0x4e,
	0xfc, 0xa2, 0xc8, 0xab, 0x26, 0xe2, 0xac, 0x7c, 0x83, 0xd7, 0x14, 0xda, 0x39, 0xe1, 0x37, 0x65,
	0x98, 0x31, 0xcf, 0xb6, 0xcf, 0x98, 0xc0, 0xd0, 0x8d, 0xb8, 0x43, 0x7f, 0x28, 0x34, 0xd8, 0x84,
	0x9c, 0xa5, 0xb7, 0xd0, 0x8e, 0x31, 0x4f, 0x6f, 0x65, 0xaa, 0x17, 0xfe, 0xdf, 0xc0, 0x0b, 0xbb,
	0xb3, 0x8d, 0xc0, 0x3e, 0x40, 0x37, 0x53, 0x98, 0x72, 0xfa, 0x1b, 0x6f, 0x30, 0xf3, 0xff, 0x11,
	0xf1, 0xdc, 0x88, 0xd3, 0x5a, 0x63, 0xef, 0xa1, 0x53, 0x8f, 0x1a, 0xd9, 0xff, 0x4f, 0x08, 0x58,
	0xe9, 0x44, 0x61, 0xca, 0xde, 0x01, 0xd8, 0xeb, 0xc9, 0xdf, 0xb3, 0x23, 0xa4, 0x90, 0x1d, 0x40,
	0xa7, 0x14, 0x99, 0xc4, 0x58, 0x66, 0x52, 0xdf, 0xfb, 0xfb, 0xe4, 0xbb, 0x52, 0xfc, 0x8c, 0x3e,
	0xeb, 0xf1, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x18, 0x59, 0x03, 0xf2, 0x02, 0x00, 0x00,
}
