// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_db_area_sum.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_database_database_summaries_database_summary

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

// OSPF Area Scope Database Summary counters
type OspfShDbAreaSum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbAreaSum_KEYS) Reset()         { *m = OspfShDbAreaSum_KEYS{} }
func (m *OspfShDbAreaSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAreaSum_KEYS) ProtoMessage()    {}
func (*OspfShDbAreaSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_31a96c912d26ffa4, []int{0}
}

func (m *OspfShDbAreaSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAreaSum_KEYS.Unmarshal(m, b)
}
func (m *OspfShDbAreaSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAreaSum_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShDbAreaSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAreaSum_KEYS.Merge(m, src)
}
func (m *OspfShDbAreaSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShDbAreaSum_KEYS.Size(m)
}
func (m *OspfShDbAreaSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbAreaSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbAreaSum_KEYS proto.InternalMessageInfo

func (m *OspfShDbAreaSum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShDbAreaSum_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

type OspfShDbAreaSum struct {
	// Summary counters for Area scope LSAs
	AreaLsaCounters []*OspfShDbSumCntr `protobuf:"bytes,50,rep,name=area_lsa_counters,json=areaLsaCounters,proto3" json:"area_lsa_counters,omitempty"`
	// Area id
	DbaseSumAreaIdString string   `protobuf:"bytes,51,opt,name=dbase_sum_area_id_string,json=dbaseSumAreaIdString,proto3" json:"dbase_sum_area_id_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbAreaSum) Reset()         { *m = OspfShDbAreaSum{} }
func (m *OspfShDbAreaSum) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAreaSum) ProtoMessage()    {}
func (*OspfShDbAreaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_31a96c912d26ffa4, []int{1}
}

func (m *OspfShDbAreaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAreaSum.Unmarshal(m, b)
}
func (m *OspfShDbAreaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAreaSum.Marshal(b, m, deterministic)
}
func (m *OspfShDbAreaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAreaSum.Merge(m, src)
}
func (m *OspfShDbAreaSum) XXX_Size() int {
	return xxx_messageInfo_OspfShDbAreaSum.Size(m)
}
func (m *OspfShDbAreaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbAreaSum.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbAreaSum proto.InternalMessageInfo

func (m *OspfShDbAreaSum) GetAreaLsaCounters() []*OspfShDbSumCntr {
	if m != nil {
		return m.AreaLsaCounters
	}
	return nil
}

func (m *OspfShDbAreaSum) GetDbaseSumAreaIdString() string {
	if m != nil {
		return m.DbaseSumAreaIdString
	}
	return ""
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
	return fileDescriptor_31a96c912d26ffa4, []int{2}
}

func (m *OspfShDbSumCntr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbSumCntr.Unmarshal(m, b)
}
func (m *OspfShDbSumCntr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbSumCntr.Marshal(b, m, deterministic)
}
func (m *OspfShDbSumCntr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbSumCntr.Merge(m, src)
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
	proto.RegisterType((*OspfShDbAreaSum_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summaries.database_summary.ospf_sh_db_area_sum_KEYS")
	proto.RegisterType((*OspfShDbAreaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summaries.database_summary.ospf_sh_db_area_sum")
	proto.RegisterType((*OspfShDbSumCntr)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summaries.database_summary.ospf_sh_db_sum_cntr")
}

func init() { proto.RegisterFile("ospf_sh_db_area_sum.proto", fileDescriptor_31a96c912d26ffa4) }

var fileDescriptor_31a96c912d26ffa4 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xbf, 0x4e, 0xe3, 0x40,
	0x10, 0xc6, 0xe5, 0xcb, 0x5d, 0x72, 0xd9, 0xe4, 0x0e, 0xb2, 0x20, 0xe1, 0x88, 0x26, 0x44, 0x14,
	0xae, 0x5c, 0x24, 0x88, 0x1e, 0x01, 0x05, 0xe2, 0x4f, 0x61, 0x23, 0x24, 0xaa, 0xd5, 0xc4, 0x1e,
	0x07, 0x8b, 0xf5, 0x1f, 0xed, 0xac, 0xa3, 0xa4, 0xe3, 0x85, 0x78, 0x0e, 0x9e, 0x0a, 0x09, 0xed,
	0x62, 0x47, 0x04, 0xa5, 0xa6, 0x9b, 0xfd, 0xe6, 0x37, 0xe3, 0xef, 0x1b, 0x99, 0x0d, 0x0b, 0x2a,
	0x13, 0x41, 0x4f, 0x22, 0x9e, 0x09, 0x50, 0x08, 0x82, 0xaa, 0xcc, 0x2f, 0x55, 0xa1, 0x0b, 0xfe,
	0x1c, 0xa5, 0x14, 0x15, 0x22, 0x2d, 0x48, 0x2c, 0x95, 0x48, 0xcb, 0xc5, 0x89, 0xb0, 0x70, 0x51,
	0xa2, 0xf2, 0x4d, 0x65, 0xb8, 0x08, 0x89, 0x90, 0x9a, 0xca, 0x8f, 0x31, 0x81, 0x4a, 0x6a, 0xb1,
	0x50, 0x89, 0x1f, 0x83, 0x86, 0x19, 0x10, 0xae, 0x0b, 0xb3, 0x3b, 0x03, 0x95, 0x22, 0x7d, 0x97,
	0x56, 0xe3, 0x07, 0xe6, 0x6e, 0x71, 0x22, 0xae, 0x2f, 0x1f, 0x43, 0x7e, 0xc4, 0xfa, 0xf5, 0x7e,
	0x91, 0x43, 0x86, 0xae, 0x33, 0x72, 0xbc, 0x6e, 0xd0, 0xab, 0xb5, 0x3b, 0xc8, 0x90, 0x1f, 0xb0,
	0x8e, 0x9d, 0x49, 0x63, 0xf7, 0xd7, 0xc8, 0xf1, 0xfe, 0x05, 0x6d, 0xf3, 0xbc, 0x8a, 0xc7, 0xef,
	0x0e, 0xdb, 0xdb, 0xb2, 0x98, 0xbf, 0x3a, 0x6c, 0x60, 0x1f, 0x92, 0x40, 0x44, 0x45, 0x95, 0x6b,
	0x54, 0xe4, 0x4e, 0x46, 0x2d, 0xaf, 0x37, 0x79, 0x71, 0xfc, 0x1f, 0x8c, 0xee, 0x7f, 0xb1, 0x67,
	0x22, 0x47, 0xb9, 0x56, 0xc1, 0x8e, 0xf1, 0x76, 0x43, 0x70, 0x5e, 0x3b, 0xe3, 0xa7, 0xcc, 0x8d,
	0x9b, 0x29, 0x51, 0x47, 0x15, 0xa4, 0x55, 0x9a, 0xcf, 0xdd, 0xa9, 0xbd, 0xc7, 0xbe, 0xed, 0x87,
	0x55, 0x76, 0x66, 0x93, 0x87, 0xb6, 0x37, 0x7e, 0xdb, 0xcc, 0xdf, 0x7c, 0x80, 0x0f, 0xd9, 0x5f,
	0x93, 0x5c, 0xaf, 0xca, 0xe6, 0x9e, 0x1d, 0x49, 0x70, 0xbf, 0x2a, 0x91, 0x1f, 0xb2, 0xee, 0xfa,
	0x28, 0xf6, 0x9a, 0x83, 0xc0, 0xb0, 0xd6, 0x0a, 0xf7, 0xd8, 0xae, 0x69, 0xc6, 0x28, 0x51, 0x63,
	0xcd, 0xb4, 0x2c, 0xf3, 0x5f, 0x12, 0x5c, 0x58, 0x79, 0x83, 0xcc, 0x60, 0x09, 0xf3, 0x86, 0xfc,
	0xbd, 0x26, 0x6f, 0xad, 0xfc, 0x49, 0x1e, 0x33, 0xa3, 0x08, 0x42, 0x99, 0xd4, 0xdc, 0x1f, 0xcb,
	0xf5, 0x25, 0x41, 0x88, 0x32, 0xb1, 0xd4, 0xac, 0x6d, 0xff, 0xca, 0xe9, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x75, 0xc7, 0x46, 0x08, 0xb2, 0x02, 0x00, 0x00,
}
