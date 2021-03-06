// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_db_area_sum.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_database_database_summaries_database_summary

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

// OSPF Area Scope Database Summary counters
type OspfShDbAreaSum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbAreaSum_KEYS) Reset()         { *m = OspfShDbAreaSum_KEYS{} }
func (m *OspfShDbAreaSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAreaSum_KEYS) ProtoMessage()    {}
func (*OspfShDbAreaSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_db_area_sum_9bec9dbc1e73dae6, []int{0}
}
func (m *OspfShDbAreaSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAreaSum_KEYS.Unmarshal(m, b)
}
func (m *OspfShDbAreaSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAreaSum_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShDbAreaSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAreaSum_KEYS.Merge(dst, src)
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

func (m *OspfShDbAreaSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	return fileDescriptor_ospf_sh_db_area_sum_9bec9dbc1e73dae6, []int{1}
}
func (m *OspfShDbAreaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAreaSum.Unmarshal(m, b)
}
func (m *OspfShDbAreaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAreaSum.Marshal(b, m, deterministic)
}
func (dst *OspfShDbAreaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAreaSum.Merge(dst, src)
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
	return fileDescriptor_ospf_sh_db_area_sum_9bec9dbc1e73dae6, []int{2}
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
	proto.RegisterType((*OspfShDbAreaSum_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.database.database_summaries.database_summary.ospf_sh_db_area_sum_KEYS")
	proto.RegisterType((*OspfShDbAreaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.database.database_summaries.database_summary.ospf_sh_db_area_sum")
	proto.RegisterType((*OspfShDbSumCntr)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.database.database_summaries.database_summary.ospf_sh_db_sum_cntr")
}

func init() {
	proto.RegisterFile("ospf_sh_db_area_sum.proto", fileDescriptor_ospf_sh_db_area_sum_9bec9dbc1e73dae6)
}

var fileDescriptor_ospf_sh_db_area_sum_9bec9dbc1e73dae6 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xbb, 0x4e, 0xc3, 0x40,
	0x10, 0x45, 0x65, 0x02, 0x09, 0xd9, 0x84, 0x47, 0x0c, 0x12, 0x8e, 0x68, 0x42, 0x44, 0xe1, 0xca,
	0x45, 0x82, 0xe8, 0x11, 0x50, 0x20, 0x1e, 0x85, 0x4d, 0x43, 0xb5, 0x9a, 0xd8, 0xe3, 0xc4, 0x92,
	0x5f, 0xec, 0x38, 0x56, 0x52, 0xf0, 0x41, 0x7c, 0x09, 0xdf, 0x44, 0x87, 0x76, 0x62, 0x47, 0x04,
	0xa5, 0xa6, 0xb1, 0x66, 0xe7, 0x1e, 0xcf, 0xde, 0xb9, 0x5a, 0xd1, 0xcf, 0x28, 0x0f, 0x25, 0xcd,
	0x64, 0x30, 0x91, 0xa0, 0x10, 0x24, 0xcd, 0x13, 0x27, 0x57, 0x59, 0x91, 0x99, 0x33, 0x3f, 0x22,
	0x3f, 0x93, 0x51, 0x46, 0x72, 0xa1, 0x64, 0x94, 0x97, 0x57, 0x92, 0xe1, 0x2c, 0x47, 0xe5, 0xe8,
	0x4a, 0x73, 0x3e, 0x12, 0x21, 0xd5, 0x95, 0x53, 0xaa, 0x90, 0x3f, 0x4e, 0x00, 0x05, 0x4c, 0x80,
	0x70, 0x5d, 0xe8, 0xc1, 0x09, 0xa8, 0x08, 0xe9, 0x6f, 0x6b, 0x39, 0x7c, 0x17, 0xd6, 0x16, 0x1b,
	0xf2, 0xf1, 0xfe, 0xcd, 0x33, 0x2f, 0x44, 0xb7, 0x1a, 0x2e, 0x53, 0x48, 0xd0, 0x32, 0x06, 0x86,
	0xdd, 0x76, 0x3b, 0x55, 0xef, 0x05, 0x12, 0x34, 0xfb, 0x62, 0xbf, 0x54, 0xe1, 0x4a, 0xde, 0x61,
	0xb9, 0x55, 0xaa, 0x90, 0xa5, 0x33, 0xd1, 0xe2, 0x71, 0x51, 0x60, 0x35, 0x06, 0x86, 0x7d, 0xe0,
	0x36, 0xf5, 0xf1, 0x21, 0x18, 0x7e, 0x1b, 0xe2, 0x64, 0xcb, 0x9d, 0xe6, 0xa7, 0x21, 0x7a, 0x7c,
	0x88, 0x09, 0xa4, 0x9f, 0xcd, 0xd3, 0x02, 0x15, 0x59, 0xa3, 0x41, 0xc3, 0xee, 0x8c, 0x3e, 0x9c,
	0xff, 0x4a, 0xc4, 0xf9, 0x65, 0x4d, 0x27, 0xe1, 0xa7, 0x85, 0x72, 0x8f, 0xb4, 0xaf, 0x27, 0x82,
	0xdb, 0xca, 0x95, 0x79, 0x2d, 0xac, 0xa0, 0xfe, 0x4b, 0x56, 0x6b, 0x4a, 0x2a, 0x54, 0x94, 0x4e,
	0xad, 0x31, 0xe7, 0x70, 0xca, 0xba, 0x37, 0x4f, 0x6e, 0x78, 0x6b, 0x8f, 0xb5, 0xe1, 0xd7, 0xe6,
	0xee, 0xf5, 0x05, 0x3a, 0x47, 0xbd, 0x75, 0xb1, 0xcc, 0xeb, 0x98, 0x5b, 0x31, 0xc1, 0xeb, 0x32,
	0x47, 0xf3, 0x5c, 0xb4, 0xd7, 0x81, 0x70, 0xc6, 0x3d, 0x57, 0xb3, 0x6c, 0xc5, 0xb4, 0xc5, 0xb1,
	0x16, 0x03, 0x8c, 0xb1, 0xc0, 0x8a, 0x69, 0x30, 0x73, 0x18, 0x13, 0xdc, 0x71, 0x7b, 0x83, 0x4c,
	0x60, 0x01, 0xd3, 0x9a, 0xdc, 0x5d, 0x93, 0xcf, 0xdc, 0x5e, 0x91, 0x97, 0x42, 0x77, 0x24, 0x61,
	0x1c, 0x56, 0xdc, 0x1e, 0x73, 0xdd, 0x98, 0xc0, 0xc3, 0x38, 0x64, 0x6a, 0xd2, 0xe4, 0x97, 0x3a,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x03, 0xdc, 0x99, 0xc6, 0x02, 0x00, 0x00,
}
