// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_summary_address.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_summary_prefixes_summary_prefix

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

// OSPF Summary Address
type OspfShSummaryAddress_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Netmask              string   `protobuf:"bytes,4,opt,name=netmask,proto3" json:"netmask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShSummaryAddress_KEYS) Reset()         { *m = OspfShSummaryAddress_KEYS{} }
func (m *OspfShSummaryAddress_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShSummaryAddress_KEYS) ProtoMessage()    {}
func (*OspfShSummaryAddress_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_summary_address_1035bde12270272f, []int{0}
}
func (m *OspfShSummaryAddress_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShSummaryAddress_KEYS.Unmarshal(m, b)
}
func (m *OspfShSummaryAddress_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShSummaryAddress_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShSummaryAddress_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShSummaryAddress_KEYS.Merge(dst, src)
}
func (m *OspfShSummaryAddress_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShSummaryAddress_KEYS.Size(m)
}
func (m *OspfShSummaryAddress_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShSummaryAddress_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShSummaryAddress_KEYS proto.InternalMessageInfo

func (m *OspfShSummaryAddress_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShSummaryAddress_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShSummaryAddress_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShSummaryAddress_KEYS) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

type OspfShSummaryAddress struct {
	// Summary prefix
	SummaryPrefix string `protobuf:"bytes,50,opt,name=summary_prefix,json=summaryPrefix,proto3" json:"summary_prefix,omitempty"`
	// Summary Netmask
	SummaryMask string `protobuf:"bytes,51,opt,name=summary_mask,json=summaryMask,proto3" json:"summary_mask,omitempty"`
	// Cost of Summary
	SummaryMetric uint32 `protobuf:"varint,52,opt,name=summary_metric,json=summaryMetric,proto3" json:"summary_metric,omitempty"`
	// Type of Metric
	SummaryMetricType string `protobuf:"bytes,53,opt,name=summary_metric_type,json=summaryMetricType,proto3" json:"summary_metric_type,omitempty"`
	// Tag associated with this summary prefix
	SummaryTag           uint32   `protobuf:"varint,54,opt,name=summary_tag,json=summaryTag,proto3" json:"summary_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShSummaryAddress) Reset()         { *m = OspfShSummaryAddress{} }
func (m *OspfShSummaryAddress) String() string { return proto.CompactTextString(m) }
func (*OspfShSummaryAddress) ProtoMessage()    {}
func (*OspfShSummaryAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_summary_address_1035bde12270272f, []int{1}
}
func (m *OspfShSummaryAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShSummaryAddress.Unmarshal(m, b)
}
func (m *OspfShSummaryAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShSummaryAddress.Marshal(b, m, deterministic)
}
func (dst *OspfShSummaryAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShSummaryAddress.Merge(dst, src)
}
func (m *OspfShSummaryAddress) XXX_Size() int {
	return xxx_messageInfo_OspfShSummaryAddress.Size(m)
}
func (m *OspfShSummaryAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShSummaryAddress.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShSummaryAddress proto.InternalMessageInfo

func (m *OspfShSummaryAddress) GetSummaryPrefix() string {
	if m != nil {
		return m.SummaryPrefix
	}
	return ""
}

func (m *OspfShSummaryAddress) GetSummaryMask() string {
	if m != nil {
		return m.SummaryMask
	}
	return ""
}

func (m *OspfShSummaryAddress) GetSummaryMetric() uint32 {
	if m != nil {
		return m.SummaryMetric
	}
	return 0
}

func (m *OspfShSummaryAddress) GetSummaryMetricType() string {
	if m != nil {
		return m.SummaryMetricType
	}
	return ""
}

func (m *OspfShSummaryAddress) GetSummaryTag() uint32 {
	if m != nil {
		return m.SummaryTag
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShSummaryAddress_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.summary_prefixes.summary_prefix.ospf_sh_summary_address_KEYS")
	proto.RegisterType((*OspfShSummaryAddress)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.summary_prefixes.summary_prefix.ospf_sh_summary_address")
}

func init() {
	proto.RegisterFile("ospf_sh_summary_address.proto", fileDescriptor_ospf_sh_summary_address_1035bde12270272f)
}

var fileDescriptor_ospf_sh_summary_address_1035bde12270272f = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x4a, 0xab, 0x63, 0x2b, 0xb8, 0x82, 0xae, 0xa0, 0xd8, 0x16, 0x84, 0x9e, 0x72,
	0xb0, 0xd5, 0x37, 0xf0, 0x24, 0x15, 0xd1, 0x5e, 0xc4, 0xc3, 0xb0, 0xa6, 0x93, 0x1a, 0x4a, 0xba,
	0xcb, 0x4e, 0x0c, 0xed, 0x4b, 0xf8, 0x82, 0xbe, 0x8c, 0x74, 0xb2, 0x11, 0x23, 0xf4, 0x12, 0x76,
	0xfe, 0xff, 0x9f, 0xef, 0x1f, 0x08, 0x5c, 0x5a, 0x76, 0x29, 0xf2, 0x07, 0xf2, 0x67, 0x9e, 0x1b,
	0xbf, 0x46, 0x33, 0x9b, 0x79, 0x62, 0x8e, 0x9d, 0xb7, 0x85, 0x55, 0x6f, 0x49, 0xc6, 0x89, 0xc5,
	0xcc, 0x32, 0xae, 0x3c, 0x66, 0xae, 0x1c, 0xa3, 0x2c, 0x58, 0x47, 0x3e, 0xde, 0xbc, 0x36, 0xb9,
	0x84, 0x98, 0x89, 0xeb, 0x57, 0x5c, 0xfa, 0x54, 0x3e, 0x71, 0x8d, 0x74, 0x9e, 0xd2, 0x6c, 0x45,
	0xfc, 0x4f, 0x18, 0x7c, 0x45, 0x70, 0xb1, 0xa5, 0x1e, 0x1f, 0xee, 0x5f, 0x5f, 0x54, 0x1f, 0x3a,
	0x01, 0x8a, 0x4b, 0x93, 0x93, 0x8e, 0x7a, 0xd1, 0xf0, 0xe0, 0xf9, 0x30, 0x68, 0x8f, 0x26, 0x27,
	0x75, 0x0e, 0xfb, 0xa5, 0x4f, 0x2b, 0x7b, 0x47, 0xec, 0x76, 0xe9, 0x53, 0xb1, 0x4e, 0xa1, 0x55,
	0x15, 0xe9, 0x5d, 0x31, 0xc2, 0xa4, 0x34, 0xb4, 0x97, 0x54, 0xe4, 0x86, 0x17, 0x7a, 0xaf, 0xda,
	0x08, 0xe3, 0xe0, 0x3b, 0x82, 0xb3, 0x2d, 0x07, 0xa9, 0x6b, 0x38, 0x6a, 0x9e, 0xaf, 0x6f, 0x64,
	0xb9, 0x1b, 0xd4, 0xa7, 0x0a, 0xde, 0x87, 0x4e, 0x1d, 0x93, 0x86, 0x51, 0x75, 0x72, 0xd0, 0x26,
	0x86, 0x17, 0x7f, 0x49, 0x39, 0x15, 0x3e, 0x4b, 0xf4, 0xb8, 0x17, 0x0d, 0xbb, 0xbf, 0xa4, 0x89,
	0x88, 0x2a, 0x86, 0x93, 0x66, 0x0c, 0x8b, 0xb5, 0x23, 0x7d, 0x2b, 0xc0, 0xe3, 0x46, 0x76, 0xba,
	0x76, 0xa4, 0xae, 0xa0, 0x6e, 0xc1, 0xc2, 0xcc, 0xf5, 0x9d, 0x30, 0x21, 0x48, 0x53, 0x33, 0x7f,
	0x6f, 0xc9, 0x2f, 0x1d, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0x62, 0xa0, 0x20, 0xf3, 0x01,
	0x00, 0x00,
}
