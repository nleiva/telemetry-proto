// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dpa_dpt_telemetry_voq_block_info.proto

package cisco_ios_xr_dpt_telemetry_oper_data_path_telemetry_sensor_non_zero_voq_data

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

// VOQ Info Block
type DpaDptTelemetryVoqBlockInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaDptTelemetryVoqBlockInfo_KEYS) Reset()         { *m = DpaDptTelemetryVoqBlockInfo_KEYS{} }
func (m *DpaDptTelemetryVoqBlockInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*DpaDptTelemetryVoqBlockInfo_KEYS) ProtoMessage()    {}
func (*DpaDptTelemetryVoqBlockInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpa_dpt_telemetry_voq_block_info_e1686ca1b4873c7d, []int{0}
}
func (m *DpaDptTelemetryVoqBlockInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaDptTelemetryVoqBlockInfo_KEYS.Unmarshal(m, b)
}
func (m *DpaDptTelemetryVoqBlockInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaDptTelemetryVoqBlockInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *DpaDptTelemetryVoqBlockInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaDptTelemetryVoqBlockInfo_KEYS.Merge(dst, src)
}
func (m *DpaDptTelemetryVoqBlockInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_DpaDptTelemetryVoqBlockInfo_KEYS.Size(m)
}
func (m *DpaDptTelemetryVoqBlockInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaDptTelemetryVoqBlockInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DpaDptTelemetryVoqBlockInfo_KEYS proto.InternalMessageInfo

type DpaDptTelemetryVoqBlockInfo struct {
	// Timestamp
	TimeStamp uint64 `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	// Rack number
	RackNumber uint32 `protobuf:"varint,51,opt,name=rack_number,json=rackNumber,proto3" json:"rack_number,omitempty"`
	// Slot number
	SlotNumber uint32 `protobuf:"varint,52,opt,name=slot_number,json=slotNumber,proto3" json:"slot_number,omitempty"`
	// NPU number
	NpuNumber uint32 `protobuf:"varint,53,opt,name=npu_number,json=npuNumber,proto3" json:"npu_number,omitempty"`
	// NPU core
	NpuCoreNum uint32 `protobuf:"varint,54,opt,name=npu_core_num,json=npuCoreNum,proto3" json:"npu_core_num,omitempty"`
	// VOQ Info Array
	VoqInfoArray         []*DpaDptTelemetryVoqInfo `protobuf:"bytes,55,rep,name=voq_info_array,json=voqInfoArray,proto3" json:"voq_info_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DpaDptTelemetryVoqBlockInfo) Reset()         { *m = DpaDptTelemetryVoqBlockInfo{} }
func (m *DpaDptTelemetryVoqBlockInfo) String() string { return proto.CompactTextString(m) }
func (*DpaDptTelemetryVoqBlockInfo) ProtoMessage()    {}
func (*DpaDptTelemetryVoqBlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpa_dpt_telemetry_voq_block_info_e1686ca1b4873c7d, []int{1}
}
func (m *DpaDptTelemetryVoqBlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaDptTelemetryVoqBlockInfo.Unmarshal(m, b)
}
func (m *DpaDptTelemetryVoqBlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaDptTelemetryVoqBlockInfo.Marshal(b, m, deterministic)
}
func (dst *DpaDptTelemetryVoqBlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaDptTelemetryVoqBlockInfo.Merge(dst, src)
}
func (m *DpaDptTelemetryVoqBlockInfo) XXX_Size() int {
	return xxx_messageInfo_DpaDptTelemetryVoqBlockInfo.Size(m)
}
func (m *DpaDptTelemetryVoqBlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaDptTelemetryVoqBlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DpaDptTelemetryVoqBlockInfo proto.InternalMessageInfo

func (m *DpaDptTelemetryVoqBlockInfo) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *DpaDptTelemetryVoqBlockInfo) GetRackNumber() uint32 {
	if m != nil {
		return m.RackNumber
	}
	return 0
}

func (m *DpaDptTelemetryVoqBlockInfo) GetSlotNumber() uint32 {
	if m != nil {
		return m.SlotNumber
	}
	return 0
}

func (m *DpaDptTelemetryVoqBlockInfo) GetNpuNumber() uint32 {
	if m != nil {
		return m.NpuNumber
	}
	return 0
}

func (m *DpaDptTelemetryVoqBlockInfo) GetNpuCoreNum() uint32 {
	if m != nil {
		return m.NpuCoreNum
	}
	return 0
}

func (m *DpaDptTelemetryVoqBlockInfo) GetVoqInfoArray() []*DpaDptTelemetryVoqInfo {
	if m != nil {
		return m.VoqInfoArray
	}
	return nil
}

// VOQ Info
type DpaDptTelemetryVoqInfo struct {
	// Interface name
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// VOQ id
	VoqId uint32 `protobuf:"varint,2,opt,name=voq_id,json=voqId,proto3" json:"voq_id,omitempty"`
	// VOQ depth
	VoqDepth             uint32   `protobuf:"varint,3,opt,name=voq_depth,json=voqDepth,proto3" json:"voq_depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaDptTelemetryVoqInfo) Reset()         { *m = DpaDptTelemetryVoqInfo{} }
func (m *DpaDptTelemetryVoqInfo) String() string { return proto.CompactTextString(m) }
func (*DpaDptTelemetryVoqInfo) ProtoMessage()    {}
func (*DpaDptTelemetryVoqInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpa_dpt_telemetry_voq_block_info_e1686ca1b4873c7d, []int{2}
}
func (m *DpaDptTelemetryVoqInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaDptTelemetryVoqInfo.Unmarshal(m, b)
}
func (m *DpaDptTelemetryVoqInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaDptTelemetryVoqInfo.Marshal(b, m, deterministic)
}
func (dst *DpaDptTelemetryVoqInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaDptTelemetryVoqInfo.Merge(dst, src)
}
func (m *DpaDptTelemetryVoqInfo) XXX_Size() int {
	return xxx_messageInfo_DpaDptTelemetryVoqInfo.Size(m)
}
func (m *DpaDptTelemetryVoqInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaDptTelemetryVoqInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DpaDptTelemetryVoqInfo proto.InternalMessageInfo

func (m *DpaDptTelemetryVoqInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *DpaDptTelemetryVoqInfo) GetVoqId() uint32 {
	if m != nil {
		return m.VoqId
	}
	return 0
}

func (m *DpaDptTelemetryVoqInfo) GetVoqDepth() uint32 {
	if m != nil {
		return m.VoqDepth
	}
	return 0
}

func init() {
	proto.RegisterType((*DpaDptTelemetryVoqBlockInfo_KEYS)(nil), "cisco_ios_xr_dpt_telemetry_oper.data_path_telemetry.sensor.non_zero_voq_data.dpa_dpt_telemetry_voq_block_info_KEYS")
	proto.RegisterType((*DpaDptTelemetryVoqBlockInfo)(nil), "cisco_ios_xr_dpt_telemetry_oper.data_path_telemetry.sensor.non_zero_voq_data.dpa_dpt_telemetry_voq_block_info")
	proto.RegisterType((*DpaDptTelemetryVoqInfo)(nil), "cisco_ios_xr_dpt_telemetry_oper.data_path_telemetry.sensor.non_zero_voq_data.dpa_dpt_telemetry_voq_info")
}

func init() {
	proto.RegisterFile("dpa_dpt_telemetry_voq_block_info.proto", fileDescriptor_dpa_dpt_telemetry_voq_block_info_e1686ca1b4873c7d)
}

var fileDescriptor_dpa_dpt_telemetry_voq_block_info_e1686ca1b4873c7d = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xfb, 0xff, 0x17, 0x33, 0x7d, 0x39, 0x04, 0x84, 0xa0, 0x14, 0x43, 0xa1, 0x9a,
	0x53, 0x0e, 0xad, 0x2f, 0x67, 0x51, 0x0f, 0xa2, 0xf4, 0x90, 0x9e, 0x3c, 0x2d, 0xdb, 0x64, 0x4a,
	0x43, 0xbb, 0x3b, 0xdb, 0xcd, 0xa6, 0x5a, 0x3f, 0x84, 0xdf, 0xce, 0xef, 0x23, 0xbb, 0x6d, 0x15,
	0x44, 0xe9, 0xc5, 0xe3, 0x3e, 0xcf, 0x6f, 0xe6, 0x99, 0x19, 0x16, 0x4e, 0x73, 0xc5, 0x59, 0xae,
	0x0c, 0x33, 0xb8, 0x40, 0x81, 0x46, 0xaf, 0xd9, 0x8a, 0x96, 0x6c, 0xb2, 0xa0, 0x6c, 0xce, 0x0a,
	0x39, 0xa5, 0x44, 0x69, 0x32, 0x14, 0x3c, 0x66, 0x45, 0x99, 0x11, 0x2b, 0xa8, 0x64, 0x2f, 0xfa,
	0x5b, 0x01, 0x29, 0xd4, 0x49, 0xce, 0x0d, 0x67, 0x8a, 0x9b, 0xd9, 0x97, 0x91, 0x94, 0x28, 0x4b,
	0xd2, 0x89, 0x24, 0xc9, 0x5e, 0x51, 0x93, 0xeb, 0x6c, 0xb9, 0xde, 0x19, 0xf4, 0xf7, 0xe5, 0xb2,
	0x87, 0xbb, 0xa7, 0x71, 0xef, 0xbd, 0x06, 0xd1, 0x3e, 0x32, 0xe8, 0x02, 0x98, 0x42, 0x20, 0x2b,
	0x0d, 0x17, 0x2a, 0x1c, 0x44, 0x5e, 0xfc, 0x2f, 0xf5, 0xad, 0x32, 0xb6, 0x42, 0x70, 0x02, 0x4d,
	0xcd, 0xb3, 0x39, 0x93, 0x95, 0x98, 0xa0, 0x0e, 0x87, 0x91, 0x17, 0xb7, 0x53, 0xb0, 0xd2, 0xc8,
	0x29, 0x16, 0x28, 0x17, 0x64, 0x76, 0xc0, 0xf9, 0x06, 0xb0, 0xd2, 0x16, 0xe8, 0x02, 0x48, 0x55,
	0xed, 0xfc, 0x0b, 0xe7, 0xfb, 0x52, 0x55, 0x5b, 0x3b, 0x82, 0x96, 0xb5, 0x33, 0xd2, 0x68, 0x99,
	0xf0, 0x72, 0xd3, 0x40, 0xaa, 0xea, 0x86, 0x34, 0x8e, 0x2a, 0x11, 0xbc, 0x79, 0xd0, 0xb1, 0x43,
	0xbb, 0xc5, 0xb8, 0xd6, 0x7c, 0x1d, 0x5e, 0x45, 0xf5, 0xb8, 0x39, 0x98, 0x25, 0x7f, 0x79, 0xd7,
	0xe4, 0xe7, 0x53, 0xd9, 0xd4, 0xb4, 0xb5, 0xa2, 0xe5, 0xbd, 0x9c, 0xd2, 0xb5, 0x4d, 0xef, 0x3d,
	0xc3, 0xd1, 0xef, 0x6c, 0xd0, 0x87, 0x4e, 0x21, 0x0d, 0xea, 0x29, 0xcf, 0x90, 0x49, 0x2e, 0x30,
	0xf4, 0x22, 0x2f, 0xf6, 0xd3, 0xf6, 0xa7, 0x3a, 0xe2, 0x02, 0x83, 0x43, 0x68, 0xb8, 0x92, 0x3c,
	0xac, 0xb9, 0x8d, 0xff, 0xdb, 0x88, 0x3c, 0x38, 0x06, 0xdf, 0x0d, 0x84, 0xca, 0xcc, 0xc2, 0xba,
	0x73, 0x0e, 0x56, 0xb4, 0xbc, 0xb5, 0xef, 0x49, 0xc3, 0x7d, 0xa7, 0xe1, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0x70, 0x73, 0x29, 0x78, 0x02, 0x00, 0x00,
}
