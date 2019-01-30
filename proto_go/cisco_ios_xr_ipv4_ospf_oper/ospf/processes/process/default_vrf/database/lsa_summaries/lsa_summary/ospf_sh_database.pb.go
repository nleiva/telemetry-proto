// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_database.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_database_lsa_summaries_lsa_summary

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

// OSPF Database Information
type OspfShDatabase_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	LsType               string   `protobuf:"bytes,3,opt,name=ls_type,json=lsType,proto3" json:"ls_type,omitempty"`
	LsId                 string   `protobuf:"bytes,4,opt,name=ls_id,json=lsId,proto3" json:"ls_id,omitempty"`
	AdvertisingRouter    string   `protobuf:"bytes,5,opt,name=advertising_router,json=advertisingRouter,proto3" json:"advertising_router,omitempty"`
	InterfaceName        string   `protobuf:"bytes,6,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDatabase_KEYS) Reset()         { *m = OspfShDatabase_KEYS{} }
func (m *OspfShDatabase_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShDatabase_KEYS) ProtoMessage()    {}
func (*OspfShDatabase_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_database_12fe31aec9b89849, []int{0}
}
func (m *OspfShDatabase_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDatabase_KEYS.Unmarshal(m, b)
}
func (m *OspfShDatabase_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDatabase_KEYS.Marshal(b, m, deterministic)
}
func (dst *OspfShDatabase_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDatabase_KEYS.Merge(dst, src)
}
func (m *OspfShDatabase_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShDatabase_KEYS.Size(m)
}
func (m *OspfShDatabase_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDatabase_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDatabase_KEYS proto.InternalMessageInfo

func (m *OspfShDatabase_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShDatabase_KEYS) GetLsType() string {
	if m != nil {
		return m.LsType
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetLsId() string {
	if m != nil {
		return m.LsId
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetAdvertisingRouter() string {
	if m != nil {
		return m.AdvertisingRouter
	}
	return ""
}

func (m *OspfShDatabase_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShDatabase struct {
	// LSA header information
	LsaHeader *OspfShDbHeader `protobuf:"bytes,50,opt,name=lsa_header,json=lsaHeader,proto3" json:"lsa_header,omitempty"`
	// Route Tag
	ExternalTag uint32 `protobuf:"varint,51,opt,name=external_tag,json=externalTag,proto3" json:"external_tag,omitempty"`
	// Number of links
	LinkCount            uint32   `protobuf:"varint,52,opt,name=link_count,json=linkCount,proto3" json:"link_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDatabase) Reset()         { *m = OspfShDatabase{} }
func (m *OspfShDatabase) String() string { return proto.CompactTextString(m) }
func (*OspfShDatabase) ProtoMessage()    {}
func (*OspfShDatabase) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_database_12fe31aec9b89849, []int{1}
}
func (m *OspfShDatabase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDatabase.Unmarshal(m, b)
}
func (m *OspfShDatabase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDatabase.Marshal(b, m, deterministic)
}
func (dst *OspfShDatabase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDatabase.Merge(dst, src)
}
func (m *OspfShDatabase) XXX_Size() int {
	return xxx_messageInfo_OspfShDatabase.Size(m)
}
func (m *OspfShDatabase) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDatabase.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDatabase proto.InternalMessageInfo

func (m *OspfShDatabase) GetLsaHeader() *OspfShDbHeader {
	if m != nil {
		return m.LsaHeader
	}
	return nil
}

func (m *OspfShDatabase) GetExternalTag() uint32 {
	if m != nil {
		return m.ExternalTag
	}
	return 0
}

func (m *OspfShDatabase) GetLinkCount() uint32 {
	if m != nil {
		return m.LinkCount
	}
	return 0
}

// OSPF LSA Database Header
type OspfShDbHeader struct {
	// LSA type
	LsType string `protobuf:"bytes,1,opt,name=ls_type,json=lsType,proto3" json:"ls_type,omitempty"`
	// LS ID
	Lsid string `protobuf:"bytes,2,opt,name=lsid,proto3" json:"lsid,omitempty"`
	// Router ID of Advertising Router
	AdvertisingRouter string `protobuf:"bytes,3,opt,name=advertising_router,json=advertisingRouter,proto3" json:"advertising_router,omitempty"`
	// Area ID in decimal or dotted-decimal format
	LsaAreaId string `protobuf:"bytes,4,opt,name=lsa_area_id,json=lsaAreaId,proto3" json:"lsa_area_id,omitempty"`
	// LSA's Age (s)
	LsaAge uint32 `protobuf:"varint,5,opt,name=lsa_age,json=lsaAge,proto3" json:"lsa_age,omitempty"`
	// If true, Do Not Age this LSA
	DnAgeLsa bool `protobuf:"varint,6,opt,name=dn_age_lsa,json=dnAgeLsa,proto3" json:"dn_age_lsa,omitempty"`
	// If true,  LSA received from neighbor during NSF
	Nsf bool `protobuf:"varint,7,opt,name=nsf,proto3" json:"nsf,omitempty"`
	// Current Sequence number
	SequenceNumber uint32 `protobuf:"varint,8,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	// Checksum value
	Checksum             uint32   `protobuf:"varint,9,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbHeader) Reset()         { *m = OspfShDbHeader{} }
func (m *OspfShDbHeader) String() string { return proto.CompactTextString(m) }
func (*OspfShDbHeader) ProtoMessage()    {}
func (*OspfShDbHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospf_sh_database_12fe31aec9b89849, []int{2}
}
func (m *OspfShDbHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbHeader.Unmarshal(m, b)
}
func (m *OspfShDbHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbHeader.Marshal(b, m, deterministic)
}
func (dst *OspfShDbHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbHeader.Merge(dst, src)
}
func (m *OspfShDbHeader) XXX_Size() int {
	return xxx_messageInfo_OspfShDbHeader.Size(m)
}
func (m *OspfShDbHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbHeader.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbHeader proto.InternalMessageInfo

func (m *OspfShDbHeader) GetLsType() string {
	if m != nil {
		return m.LsType
	}
	return ""
}

func (m *OspfShDbHeader) GetLsid() string {
	if m != nil {
		return m.Lsid
	}
	return ""
}

func (m *OspfShDbHeader) GetAdvertisingRouter() string {
	if m != nil {
		return m.AdvertisingRouter
	}
	return ""
}

func (m *OspfShDbHeader) GetLsaAreaId() string {
	if m != nil {
		return m.LsaAreaId
	}
	return ""
}

func (m *OspfShDbHeader) GetLsaAge() uint32 {
	if m != nil {
		return m.LsaAge
	}
	return 0
}

func (m *OspfShDbHeader) GetDnAgeLsa() bool {
	if m != nil {
		return m.DnAgeLsa
	}
	return false
}

func (m *OspfShDbHeader) GetNsf() bool {
	if m != nil {
		return m.Nsf
	}
	return false
}

func (m *OspfShDbHeader) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *OspfShDbHeader) GetChecksum() uint32 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShDatabase_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.lsa_summaries.lsa_summary.ospf_sh_database_KEYS")
	proto.RegisterType((*OspfShDatabase)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.lsa_summaries.lsa_summary.ospf_sh_database")
	proto.RegisterType((*OspfShDbHeader)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.lsa_summaries.lsa_summary.ospf_sh_db_header")
}

func init() {
	proto.RegisterFile("ospf_sh_database.proto", fileDescriptor_ospf_sh_database_12fe31aec9b89849)
}

var fileDescriptor_ospf_sh_database_12fe31aec9b89849 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x26, 0x4d, 0x93, 0x09, 0x29, 0xed, 0x22, 0x60, 0x85, 0x00, 0x85, 0x48, 0x88,
	0x5c, 0xf0, 0xa1, 0xed, 0x0b, 0x44, 0x08, 0x89, 0x0a, 0xd4, 0x43, 0xe8, 0x85, 0xd3, 0x6a, 0x62,
	0x8f, 0x9d, 0x55, 0xd7, 0x6b, 0xb3, 0xb3, 0x8e, 0x9a, 0xe7, 0xe0, 0xcc, 0x73, 0xf1, 0x0e, 0x3c,
	0x05, 0xda, 0x6d, 0x12, 0x85, 0xa2, 0x1e, 0xb9, 0xcd, 0x7e, 0xf3, 0x65, 0xfe, 0xfd, 0x62, 0x78,
	0x56, 0x73, 0x53, 0x28, 0x5e, 0xaa, 0x1c, 0x3d, 0x2e, 0x90, 0x29, 0x6d, 0x5c, 0xed, 0x6b, 0x81,
	0x99, 0xe6, 0xac, 0x56, 0xba, 0x66, 0x75, 0xeb, 0x94, 0x6e, 0x56, 0x17, 0x2a, 0x3a, 0xeb, 0x86,
	0x5c, 0x1a, 0xa2, 0xe0, 0xcb, 0x88, 0x99, 0x78, 0x1b, 0xa5, 0x39, 0x15, 0xd8, 0x1a, 0xaf, 0x56,
	0xae, 0x48, 0x77, 0xe5, 0x0c, 0xa3, 0xe2, 0xb6, 0xaa, 0xd0, 0x69, 0xe2, 0xbd, 0xd7, 0x7a, 0xf2,
	0x2b, 0x81, 0xa7, 0xf7, 0xbb, 0xab, 0xcf, 0x1f, 0xbf, 0x7d, 0x15, 0x6f, 0xe0, 0xd1, 0xa6, 0xa6,
	0xb2, 0x58, 0x91, 0x4c, 0xc6, 0xc9, 0x74, 0x30, 0x1f, 0x6e, 0xb4, 0x2b, 0xac, 0x48, 0x3c, 0x87,
	0x23, 0x74, 0x84, 0x4a, 0xe7, 0xf2, 0x60, 0x9c, 0x4c, 0x47, 0xf3, 0x5e, 0x78, 0x5e, 0xe6, 0x21,
	0x61, 0x58, 0xf9, 0x75, 0x43, 0xb2, 0x13, 0x7f, 0xd6, 0x33, 0x7c, 0xbd, 0x6e, 0x48, 0x3c, 0x81,
	0x43, 0xc3, 0xc1, 0xdf, 0x8d, 0x72, 0xd7, 0xf0, 0x65, 0x2e, 0xde, 0x83, 0xc0, 0x7c, 0x45, 0xce,
	0x6b, 0xd6, 0xb6, 0x54, 0xae, 0x6e, 0x3d, 0x39, 0x79, 0x18, 0x1d, 0xa7, 0x7b, 0x99, 0x79, 0x4c,
	0x88, 0xb7, 0x70, 0xac, 0xad, 0x27, 0x57, 0x60, 0x46, 0x77, 0xa3, 0xf5, 0xa2, 0x75, 0xb4, 0x53,
	0xc3, 0x70, 0x93, 0xdf, 0x09, 0x9c, 0xdc, 0xdf, 0x4c, 0xfc, 0x48, 0x00, 0xc2, 0xfa, 0x4b, 0xc2,
	0x9c, 0x9c, 0x3c, 0x1b, 0x27, 0xd3, 0xe1, 0x99, 0x4f, 0xff, 0xfb, 0x9d, 0xd3, 0xdd, 0x24, 0x8b,
	0x4d, 0xef, 0xf9, 0xc0, 0x30, 0x7e, 0x8a, 0x61, 0x38, 0x35, 0xdd, 0x7a, 0x72, 0x16, 0x8d, 0xf2,
	0x58, 0xca, 0xf3, 0x78, 0xcc, 0xe1, 0x56, 0xbb, 0xc6, 0x52, 0xbc, 0x02, 0x30, 0xda, 0xde, 0xa8,
	0xac, 0x6e, 0xad, 0x97, 0x17, 0xd1, 0x30, 0x08, 0xca, 0x87, 0x20, 0x4c, 0x7e, 0x1e, 0xc0, 0xe9,
	0x3f, 0x2d, 0xf6, 0x31, 0x24, 0x7f, 0x61, 0x10, 0xd0, 0x35, 0xbc, 0xa1, 0x16, 0x29, 0xe8, 0x87,
	0x28, 0x74, 0x1e, 0xa2, 0xf0, 0x1a, 0x86, 0x61, 0xbf, 0x2d, 0xff, 0x3b, 0x9e, 0x61, 0xa7, 0xd9,
	0xde, 0x5f, 0x00, 0x15, 0x96, 0x14, 0x49, 0x8e, 0x42, 0x6f, 0x9c, 0x95, 0x24, 0x5e, 0x02, 0xe4,
	0x36, 0xe8, 0xca, 0x30, 0x46, 0x74, 0xfd, 0x79, 0x3f, 0xb7, 0xb3, 0x92, 0xbe, 0x30, 0x8a, 0x13,
	0xe8, 0x58, 0x2e, 0xe4, 0x51, 0x94, 0x43, 0x28, 0xde, 0xc1, 0x63, 0xa6, 0xef, 0x2d, 0xd9, 0x40,
	0xbb, 0xad, 0x16, 0xe4, 0x64, 0x3f, 0x16, 0x3c, 0xde, 0xca, 0x57, 0x51, 0x15, 0x2f, 0xa0, 0x9f,
	0x2d, 0x29, 0xbb, 0xe1, 0xb6, 0x92, 0x83, 0xe8, 0xd8, 0xbd, 0x17, 0xbd, 0xf8, 0x41, 0x9d, 0xff,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x02, 0x0e, 0x1e, 0x79, 0x6a, 0x03, 0x00, 0x00,
}
