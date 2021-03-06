// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_db_as_sum.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_database_database_summary_as

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

// OSPF AS Scope LSA Database Summary counters
type OspfShDbAsSum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbAsSum_KEYS) Reset()         { *m = OspfShDbAsSum_KEYS{} }
func (m *OspfShDbAsSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAsSum_KEYS) ProtoMessage()    {}
func (*OspfShDbAsSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_db_as_sum_ff7dca0c2a3a0a7a, []int{0}
}
func (m *OspfShDbAsSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAsSum_KEYS.Unmarshal(m, b)
}
func (m *OspfShDbAsSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAsSum_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShDbAsSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAsSum_KEYS.Merge(dst, src)
}
func (m *OspfShDbAsSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShDbAsSum_KEYS.Size(m)
}
func (m *OspfShDbAsSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbAsSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbAsSum_KEYS proto.InternalMessageInfo

func (m *OspfShDbAsSum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type OspfShDbAsSum struct {
	// Summary counter for AS scope LSAs
	AsLsaCounters        []*OspfShDbSumCntr `protobuf:"bytes,50,rep,name=as_lsa_counters,json=asLsaCounters,proto3" json:"as_lsa_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OspfShDbAsSum) Reset()         { *m = OspfShDbAsSum{} }
func (m *OspfShDbAsSum) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAsSum) ProtoMessage()    {}
func (*OspfShDbAsSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_db_as_sum_ff7dca0c2a3a0a7a, []int{1}
}
func (m *OspfShDbAsSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAsSum.Unmarshal(m, b)
}
func (m *OspfShDbAsSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAsSum.Marshal(b, m, deterministic)
}
func (dst *OspfShDbAsSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAsSum.Merge(dst, src)
}
func (m *OspfShDbAsSum) XXX_Size() int {
	return xxx_messageInfo_OspfShDbAsSum.Size(m)
}
func (m *OspfShDbAsSum) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbAsSum.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbAsSum proto.InternalMessageInfo

func (m *OspfShDbAsSum) GetAsLsaCounters() []*OspfShDbSumCntr {
	if m != nil {
		return m.AsLsaCounters
	}
	return nil
}

// OSPF Summary counters for a type
type OspfShDbSumCntr struct {
	// Type of LSAs
	LsaType string `protobuf:"bytes,1,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	// Number of LSAs of this type
	LsaCount int32 `protobuf:"zigzag32,2,opt,name=lsa_count,json=lsaCount,proto3" json:"lsa_count,omitempty"`
	// Number of deleted LSAs of this type
	LsaDeleteCount int32 `protobuf:"zigzag32,3,opt,name=lsa_delete_count,json=lsaDeleteCount,proto3" json:"lsa_delete_count,omitempty"`
	// Number of MaxAged LSAs of this type
	LsaMaxageCount int32 `protobuf:"zigzag32,4,opt,name=lsa_maxage_count,json=lsaMaxageCount,proto3" json:"lsa_maxage_count,omitempty"`
	// Number of self-generated LSAs
	LsaSelfCount         int32    `protobuf:"zigzag32,5,opt,name=lsa_self_count,json=lsaSelfCount,proto3" json:"lsa_self_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbSumCntr) Reset()         { *m = OspfShDbSumCntr{} }
func (m *OspfShDbSumCntr) String() string { return proto.CompactTextString(m) }
func (*OspfShDbSumCntr) ProtoMessage()    {}
func (*OspfShDbSumCntr) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_db_as_sum_ff7dca0c2a3a0a7a, []int{2}
}
func (m *OspfShDbSumCntr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbSumCntr.Unmarshal(m, b)
}
func (m *OspfShDbSumCntr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbSumCntr.Marshal(b, m, deterministic)
}
func (dst *OspfShDbSumCntr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbSumCntr.Merge(dst, src)
}
func (m *OspfShDbSumCntr) XXX_Size() int {
	return xxx_messageInfo_OspfShDbSumCntr.Size(m)
}
func (m *OspfShDbSumCntr) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbSumCntr.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbSumCntr proto.InternalMessageInfo

func (m *OspfShDbSumCntr) GetLsaType() string {
	if m != nil {
		return m.LsaType
	}
	return ""
}

func (m *OspfShDbSumCntr) GetLsaCount() int32 {
	if m != nil {
		return m.LsaCount
	}
	return 0
}

func (m *OspfShDbSumCntr) GetLsaDeleteCount() int32 {
	if m != nil {
		return m.LsaDeleteCount
	}
	return 0
}

func (m *OspfShDbSumCntr) GetLsaMaxageCount() int32 {
	if m != nil {
		return m.LsaMaxageCount
	}
	return 0
}

func (m *OspfShDbSumCntr) GetLsaSelfCount() int32 {
	if m != nil {
		return m.LsaSelfCount
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShDbAsSum_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summary_as.ospf_sh_db_as_sum_KEYS")
	proto.RegisterType((*OspfShDbAsSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summary_as.ospf_sh_db_as_sum")
	proto.RegisterType((*OspfShDbSumCntr)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summary_as.ospf_sh_db_sum_cntr")
}

func init() {
	proto.RegisterFile("ospf_sh_db_as_sum.proto", fileDescriptor_ospf_sh_db_as_sum_ff7dca0c2a3a0a7a)
}

var fileDescriptor_ospf_sh_db_as_sum_ff7dca0c2a3a0a7a = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xd2, 0xbb, 0x4e, 0xeb, 0x40,
	0x10, 0x06, 0x60, 0xf9, 0xe4, 0x70, 0xc9, 0x26, 0x5c, 0xb2, 0x48, 0x10, 0x44, 0x13, 0x22, 0x0a,
	0x57, 0x5b, 0x04, 0x3a, 0x4a, 0xa0, 0xe2, 0x52, 0x24, 0x34, 0x88, 0x62, 0x34, 0xb1, 0xc7, 0x10,
	0x69, 0x9d, 0x5d, 0xed, 0x6c, 0xa2, 0xe4, 0x4d, 0x78, 0x0a, 0x5e, 0x83, 0xd7, 0x42, 0xbb, 0x38,
	0x16, 0x28, 0xb4, 0x74, 0xa3, 0xdf, 0x9f, 0xc7, 0xf3, 0x4b, 0x16, 0x47, 0x86, 0x6d, 0x01, 0xfc,
	0x0a, 0xf9, 0x18, 0x90, 0x81, 0x67, 0xa5, 0xb2, 0xce, 0x78, 0x23, 0x9f, 0xb3, 0x09, 0x67, 0x06,
	0x26, 0x86, 0x61, 0xe1, 0x60, 0x62, 0xe7, 0x17, 0x10, 0xa9, 0xb1, 0xe4, 0x54, 0x98, 0x82, 0xcb,
	0x88, 0x99, 0x78, 0x35, 0xa9, 0x9c, 0x0a, 0x9c, 0x69, 0x0f, 0x73, 0x57, 0xa8, 0x1c, 0x3d, 0x8e,
	0x91, 0xa9, 0x1e, 0xc2, 0xee, 0x12, 0xdd, 0x12, 0x90, 0xfb, 0x97, 0xe2, 0x70, 0xed, 0xbb, 0x70,
	0x7b, 0xf3, 0x34, 0x92, 0xa7, 0xa2, 0x5d, 0x6d, 0x83, 0x29, 0x96, 0xd4, 0x4d, 0x7a, 0x49, 0xda,
	0x1c, 0xb6, 0xaa, 0xec, 0x01, 0x4b, 0xea, 0xbf, 0x27, 0xa2, 0xb3, 0xf6, 0xb6, 0x7c, 0x4b, 0xc4,
	0x1e, 0x32, 0x68, 0x46, 0xc8, 0xcc, 0x6c, 0xea, 0xc9, 0x71, 0x77, 0xd0, 0x6b, 0xa4, 0xad, 0x81,
	0x55, 0x7f, 0x58, 0x45, 0x7d, 0xbb, 0x24, 0x94, 0xc8, 0xa6, 0xde, 0x0d, 0x77, 0x90, 0xef, 0x18,
	0xaf, 0xaa, 0x33, 0xfa, 0x1f, 0x89, 0x38, 0xf8, 0x85, 0xc9, 0x63, 0xb1, 0x1d, 0xce, 0xf5, 0x4b,
	0xbb, 0xea, 0xb9, 0xa5, 0x19, 0x1f, 0x97, 0x96, 0xe4, 0x89, 0x68, 0xd6, 0x4d, 0xba, 0xff, 0x7a,
	0x49, 0xda, 0x19, 0x06, 0x1b, 0x57, 0xca, 0x54, 0xec, 0x87, 0x87, 0x39, 0x69, 0xf2, 0x54, 0x99,
	0x46, 0x34, 0xbb, 0x9a, 0xf1, 0x3a, 0xc6, 0x3f, 0x64, 0x89, 0x0b, 0x7c, 0x59, 0xc9, 0xff, 0xb5,
	0xbc, 0x8f, 0xf1, 0x97, 0x3c, 0x13, 0x21, 0x01, 0x26, 0x5d, 0x54, 0x6e, 0x23, 0xba, 0xb6, 0x66,
	0x1c, 0x91, 0x2e, 0xa2, 0x1a, 0x6f, 0xc6, 0x7f, 0xe3, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0x6e, 0x1b, 0x58, 0x36, 0x02, 0x00, 0x00,
}
