// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_area.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_area_informations_area_information

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

// OSPFv3 area summary information
type Ospfv3EdmArea_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmArea_KEYS) Reset()         { *m = Ospfv3EdmArea_KEYS{} }
func (m *Ospfv3EdmArea_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmArea_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmArea_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_area_430859e84a9192a0, []int{0}
}
func (m *Ospfv3EdmArea_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmArea_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmArea_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmArea_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmArea_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmArea_KEYS.Merge(dst, src)
}
func (m *Ospfv3EdmArea_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmArea_KEYS.Size(m)
}
func (m *Ospfv3EdmArea_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmArea_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmArea_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmArea_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmArea_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmArea_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

type Ospfv3EdmArea struct {
	// If true, Backbone area is active
	IsBackboneAreaActive bool `protobuf:"varint,50,opt,name=is_backbone_area_active,json=isBackboneAreaActive,proto3" json:"is_backbone_area_active,omitempty"`
	// Number of interfaces in the area
	AreaInterfaces uint32 `protobuf:"varint,51,opt,name=area_interfaces,json=areaInterfaces,proto3" json:"area_interfaces,omitempty"`
	// If true, stub area
	IsAreaStubbed bool `protobuf:"varint,52,opt,name=is_area_stubbed,json=isAreaStubbed,proto3" json:"is_area_stubbed,omitempty"`
	// If true, totally stubby area
	IsAreaTotalStubbed bool `protobuf:"varint,53,opt,name=is_area_total_stubbed,json=isAreaTotalStubbed,proto3" json:"is_area_total_stubbed,omitempty"`
	// Default cost for Stub or NSSA area
	StubDefaultCost uint32 `protobuf:"varint,54,opt,name=stub_default_cost,json=stubDefaultCost,proto3" json:"stub_default_cost,omitempty"`
	// If true, area is a NSSA
	IsAreaNssa bool `protobuf:"varint,55,opt,name=is_area_nssa,json=isAreaNssa,proto3" json:"is_area_nssa,omitempty"`
	// If true, No redistribution into this NSSA area
	NssaNoRedistribution bool `protobuf:"varint,56,opt,name=nssa_no_redistribution,json=nssaNoRedistribution,proto3" json:"nssa_no_redistribution,omitempty"`
	// If true, perform 7/5 translation
	IsNssaTranslated bool `protobuf:"varint,57,opt,name=is_nssa_translated,json=isNssaTranslated,proto3" json:"is_nssa_translated,omitempty"`
	// If true, generate NSSA default route
	IsNssaDefault bool `protobuf:"varint,58,opt,name=is_nssa_default,json=isNssaDefault,proto3" json:"is_nssa_default,omitempty"`
	// If true, RRR is enabled
	IsRrrEnabled bool `protobuf:"varint,59,opt,name=is_rrr_enabled,json=isRrrEnabled,proto3" json:"is_rrr_enabled,omitempty"`
	// Number of SPF calculations run
	SpFs uint32 `protobuf:"varint,60,opt,name=sp_fs,json=spFs,proto3" json:"sp_fs,omitempty"`
	// List of ranges to summarize
	AreaRangeList []*Ospfv3EdmAreaRange `protobuf:"bytes,61,rep,name=area_range_list,json=areaRangeList,proto3" json:"area_range_list,omitempty"`
	// Number of opaque LSAs in the area
	AreaOpaqueLsAs uint32 `protobuf:"varint,62,opt,name=area_opaque_ls_as,json=areaOpaqueLsAs,proto3" json:"area_opaque_ls_as,omitempty"`
	// Sum of opaque LSA checksums
	AreaOpaqueLsaChecksum uint32 `protobuf:"varint,63,opt,name=area_opaque_lsa_checksum,json=areaOpaqueLsaChecksum,proto3" json:"area_opaque_lsa_checksum,omitempty"`
	// Number of LSA with demand circuit bit not set
	AreaDcBitlessLsAs uint32 `protobuf:"varint,64,opt,name=area_dc_bitless_ls_as,json=areaDcBitlessLsAs,proto3" json:"area_dc_bitless_ls_as,omitempty"`
	// Number of indication LSAs
	IndicationLsAs uint32 `protobuf:"varint,65,opt,name=indication_ls_as,json=indicationLsAs,proto3" json:"indication_ls_as,omitempty"`
	// Number of do not age LSAs
	DoNotAgeLsAs uint32 `protobuf:"varint,66,opt,name=do_not_age_ls_as,json=doNotAgeLsAs,proto3" json:"do_not_age_ls_as,omitempty"`
	// Number of LSAs which need to be flooded
	FloodListLength uint32 `protobuf:"varint,67,opt,name=flood_list_length,json=floodListLength,proto3" json:"flood_list_length,omitempty"`
	// Number of LFA enabled interfaces
	AreaLfaInterfaceCount uint32 `protobuf:"varint,68,opt,name=area_lfa_interface_count,json=areaLfaInterfaceCount,proto3" json:"area_lfa_interface_count,omitempty"`
	// Number of Per Prefix LFA enabled interfaces
	AreaPerPrefixLfaInterfaceCount uint32 `protobuf:"varint,69,opt,name=area_per_prefix_lfa_interface_count,json=areaPerPrefixLfaInterfaceCount,proto3" json:"area_per_prefix_lfa_interface_count,omitempty"`
	// Area LFA revision
	AreaLfaRevision      uint32   `protobuf:"varint,70,opt,name=area_lfa_revision,json=areaLfaRevision,proto3" json:"area_lfa_revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmArea) Reset()         { *m = Ospfv3EdmArea{} }
func (m *Ospfv3EdmArea) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmArea) ProtoMessage()    {}
func (*Ospfv3EdmArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_area_430859e84a9192a0, []int{1}
}
func (m *Ospfv3EdmArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmArea.Unmarshal(m, b)
}
func (m *Ospfv3EdmArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmArea.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmArea.Merge(dst, src)
}
func (m *Ospfv3EdmArea) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmArea.Size(m)
}
func (m *Ospfv3EdmArea) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmArea.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmArea proto.InternalMessageInfo

func (m *Ospfv3EdmArea) GetIsBackboneAreaActive() bool {
	if m != nil {
		return m.IsBackboneAreaActive
	}
	return false
}

func (m *Ospfv3EdmArea) GetAreaInterfaces() uint32 {
	if m != nil {
		return m.AreaInterfaces
	}
	return 0
}

func (m *Ospfv3EdmArea) GetIsAreaStubbed() bool {
	if m != nil {
		return m.IsAreaStubbed
	}
	return false
}

func (m *Ospfv3EdmArea) GetIsAreaTotalStubbed() bool {
	if m != nil {
		return m.IsAreaTotalStubbed
	}
	return false
}

func (m *Ospfv3EdmArea) GetStubDefaultCost() uint32 {
	if m != nil {
		return m.StubDefaultCost
	}
	return 0
}

func (m *Ospfv3EdmArea) GetIsAreaNssa() bool {
	if m != nil {
		return m.IsAreaNssa
	}
	return false
}

func (m *Ospfv3EdmArea) GetNssaNoRedistribution() bool {
	if m != nil {
		return m.NssaNoRedistribution
	}
	return false
}

func (m *Ospfv3EdmArea) GetIsNssaTranslated() bool {
	if m != nil {
		return m.IsNssaTranslated
	}
	return false
}

func (m *Ospfv3EdmArea) GetIsNssaDefault() bool {
	if m != nil {
		return m.IsNssaDefault
	}
	return false
}

func (m *Ospfv3EdmArea) GetIsRrrEnabled() bool {
	if m != nil {
		return m.IsRrrEnabled
	}
	return false
}

func (m *Ospfv3EdmArea) GetSpFs() uint32 {
	if m != nil {
		return m.SpFs
	}
	return 0
}

func (m *Ospfv3EdmArea) GetAreaRangeList() []*Ospfv3EdmAreaRange {
	if m != nil {
		return m.AreaRangeList
	}
	return nil
}

func (m *Ospfv3EdmArea) GetAreaOpaqueLsAs() uint32 {
	if m != nil {
		return m.AreaOpaqueLsAs
	}
	return 0
}

func (m *Ospfv3EdmArea) GetAreaOpaqueLsaChecksum() uint32 {
	if m != nil {
		return m.AreaOpaqueLsaChecksum
	}
	return 0
}

func (m *Ospfv3EdmArea) GetAreaDcBitlessLsAs() uint32 {
	if m != nil {
		return m.AreaDcBitlessLsAs
	}
	return 0
}

func (m *Ospfv3EdmArea) GetIndicationLsAs() uint32 {
	if m != nil {
		return m.IndicationLsAs
	}
	return 0
}

func (m *Ospfv3EdmArea) GetDoNotAgeLsAs() uint32 {
	if m != nil {
		return m.DoNotAgeLsAs
	}
	return 0
}

func (m *Ospfv3EdmArea) GetFloodListLength() uint32 {
	if m != nil {
		return m.FloodListLength
	}
	return 0
}

func (m *Ospfv3EdmArea) GetAreaLfaInterfaceCount() uint32 {
	if m != nil {
		return m.AreaLfaInterfaceCount
	}
	return 0
}

func (m *Ospfv3EdmArea) GetAreaPerPrefixLfaInterfaceCount() uint32 {
	if m != nil {
		return m.AreaPerPrefixLfaInterfaceCount
	}
	return 0
}

func (m *Ospfv3EdmArea) GetAreaLfaRevision() uint32 {
	if m != nil {
		return m.AreaLfaRevision
	}
	return 0
}

// OSPFv3 area range information
type Ospfv3EdmAreaRange struct {
	// IP prefix for summarization
	RangePrefix string `protobuf:"bytes,1,opt,name=range_prefix,json=rangePrefix,proto3" json:"range_prefix,omitempty"`
	// IP prefix length for summarization
	RangePrefixLength uint32 `protobuf:"varint,2,opt,name=range_prefix_length,json=rangePrefixLength,proto3" json:"range_prefix_length,omitempty"`
	// Net cost
	NetCost uint32 `protobuf:"varint,3,opt,name=net_cost,json=netCost,proto3" json:"net_cost,omitempty"`
	// Area range status
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// If true, cost is configured
	IsCostConfigured     bool     `protobuf:"varint,5,opt,name=is_cost_configured,json=isCostConfigured,proto3" json:"is_cost_configured,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmAreaRange) Reset()         { *m = Ospfv3EdmAreaRange{} }
func (m *Ospfv3EdmAreaRange) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmAreaRange) ProtoMessage()    {}
func (*Ospfv3EdmAreaRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_ospfv3_edm_area_430859e84a9192a0, []int{2}
}
func (m *Ospfv3EdmAreaRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmAreaRange.Unmarshal(m, b)
}
func (m *Ospfv3EdmAreaRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmAreaRange.Marshal(b, m, deterministic)
}
func (dst *Ospfv3EdmAreaRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmAreaRange.Merge(dst, src)
}
func (m *Ospfv3EdmAreaRange) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmAreaRange.Size(m)
}
func (m *Ospfv3EdmAreaRange) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmAreaRange.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmAreaRange proto.InternalMessageInfo

func (m *Ospfv3EdmAreaRange) GetRangePrefix() string {
	if m != nil {
		return m.RangePrefix
	}
	return ""
}

func (m *Ospfv3EdmAreaRange) GetRangePrefixLength() uint32 {
	if m != nil {
		return m.RangePrefixLength
	}
	return 0
}

func (m *Ospfv3EdmAreaRange) GetNetCost() uint32 {
	if m != nil {
		return m.NetCost
	}
	return 0
}

func (m *Ospfv3EdmAreaRange) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Ospfv3EdmAreaRange) GetIsCostConfigured() bool {
	if m != nil {
		return m.IsCostConfigured
	}
	return false
}

func init() {
	proto.RegisterType((*Ospfv3EdmArea_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.area_informations.area_information.ospfv3_edm_area_KEYS")
	proto.RegisterType((*Ospfv3EdmArea)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.area_informations.area_information.ospfv3_edm_area")
	proto.RegisterType((*Ospfv3EdmAreaRange)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.area_informations.area_information.ospfv3_edm_area_range")
}

func init() {
	proto.RegisterFile("ospfv3_edm_area.proto", fileDescriptor_ospfv3_edm_area_430859e84a9192a0)
}

var fileDescriptor_ospfv3_edm_area_430859e84a9192a0 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x4f, 0x23, 0x37,
	0x14, 0x55, 0xf6, 0x03, 0x52, 0x6f, 0xd8, 0x80, 0x77, 0xd9, 0xf5, 0xbe, 0x54, 0x29, 0xad, 0xb6,
	0x29, 0xaa, 0xa2, 0x96, 0xcf, 0x7e, 0xb7, 0x21, 0x80, 0x54, 0x11, 0xa5, 0x68, 0xe0, 0xa5, 0x4f,
	0x96, 0x67, 0xc6, 0x13, 0x2c, 0x26, 0xf6, 0xd4, 0xd7, 0x13, 0xd1, 0x9f, 0xd3, 0xdf, 0xc3, 0x9f,
	0xaa, 0x7c, 0xed, 0x49, 0x02, 0xf4, 0xb5, 0x2f, 0x23, 0xcf, 0x39, 0xe7, 0x9e, 0x7b, 0x6d, 0xdf,
	0x6b, 0xb2, 0x6d, 0xa0, 0x2a, 0xe6, 0xfb, 0x5c, 0xe6, 0x33, 0x2e, 0xac, 0x14, 0x83, 0xca, 0x1a,
	0x67, 0x68, 0x9a, 0x29, 0xc8, 0x0c, 0x57, 0x06, 0xf8, 0x9d, 0xe5, 0xaa, 0x9a, 0x1f, 0xf1, 0x28,
	0x34, 0x95, 0xb4, 0x83, 0xb0, 0xf6, 0xda, 0x4c, 0x02, 0x48, 0x68, 0x56, 0x83, 0xb9, 0x2d, 0xf0,
	0x33, 0xf0, 0x5e, 0x5c, 0xe9, 0xc2, 0xd8, 0x99, 0x70, 0xca, 0x68, 0x78, 0x82, 0xec, 0xcc, 0xc8,
	0xdb, 0x47, 0xc9, 0xf9, 0xc5, 0xd9, 0x9f, 0x57, 0xf4, 0x33, 0xd2, 0x89, 0x76, 0x5c, 0x8b, 0x99,
	0x64, 0xad, 0x5e, 0xab, 0xff, 0x49, 0xf2, 0x2a, 0x62, 0x13, 0x31, 0x93, 0xf4, 0x03, 0x69, 0xcf,
	0x6d, 0x11, 0xe8, 0x67, 0x48, 0xaf, 0xcf, 0x6d, 0x81, 0xd4, 0x7b, 0xb2, 0x1e, 0x32, 0xe5, 0xec,
	0x79, 0xaf, 0xd5, 0xdf, 0x48, 0xd6, 0xfc, 0xef, 0xef, 0xf9, 0xce, 0x7d, 0x9b, 0x74, 0x1f, 0xe5,
	0xa3, 0x87, 0xe4, 0xbd, 0x02, 0x9e, 0x8a, 0xec, 0x36, 0x35, 0x5a, 0x86, 0x1a, 0x44, 0xe6, 0xd4,
	0x5c, 0xb2, 0xbd, 0x5e, 0xab, 0xdf, 0x4e, 0xde, 0x2a, 0x38, 0x89, 0xec, 0xd0, 0x4a, 0x31, 0x44,
	0x8e, 0x7e, 0x49, 0xba, 0x71, 0x37, 0x4e, 0xda, 0x42, 0x64, 0x12, 0xd8, 0x3e, 0xe6, 0x7a, 0x8d,
	0xb9, 0x16, 0x28, 0xfd, 0x48, 0xba, 0x0a, 0x82, 0x2d, 0xb8, 0x3a, 0x4d, 0x65, 0xce, 0x0e, 0xd0,
	0x77, 0x43, 0x81, 0xf7, 0xbb, 0x0a, 0x20, 0xfd, 0x96, 0x6c, 0x37, 0x3a, 0x67, 0x9c, 0x28, 0x17,
	0xea, 0x43, 0x54, 0xd3, 0xa0, 0xbe, 0xf6, 0x54, 0x13, 0xb2, 0x4b, 0xb6, 0xbc, 0x88, 0xe7, 0xb2,
	0x10, 0x75, 0xe9, 0x78, 0x66, 0xc0, 0xb1, 0x23, 0xac, 0xa2, 0xeb, 0x89, 0xd3, 0x80, 0x8f, 0x0c,
	0x38, 0xda, 0x23, 0x9d, 0xc6, 0x5e, 0x03, 0x08, 0x76, 0x8c, 0xae, 0x24, 0xb8, 0x4e, 0x00, 0x04,
	0x3d, 0x20, 0xef, 0x3c, 0xc3, 0xb5, 0xe1, 0x56, 0xe6, 0x0a, 0x9c, 0x55, 0x69, 0xed, 0x6f, 0x89,
	0x7d, 0x17, 0xce, 0xc1, 0xb3, 0x13, 0x93, 0x3c, 0xe0, 0xe8, 0xd7, 0x84, 0x2a, 0x40, 0x4b, 0xee,
	0xac, 0xd0, 0x50, 0x0a, 0x27, 0x73, 0xf6, 0x3d, 0x46, 0x6c, 0x2a, 0xf0, 0xce, 0xd7, 0x0b, 0x3c,
	0x1e, 0x06, 0xaa, 0x63, 0xd1, 0xec, 0x87, 0xe6, 0x30, 0xbc, 0x34, 0x56, 0x4c, 0xbf, 0x20, 0xaf,
	0x15, 0x70, 0x6b, 0x2d, 0x97, 0x5a, 0xa4, 0xa5, 0xcc, 0xd9, 0x8f, 0x28, 0xeb, 0x28, 0x48, 0xac,
	0x3d, 0x0b, 0x18, 0x7d, 0x43, 0x5e, 0x42, 0xc5, 0x0b, 0x60, 0x3f, 0xe1, 0x9e, 0x5f, 0x40, 0x75,
	0x0e, 0xf4, 0x9f, 0x56, 0xbc, 0x19, 0x2b, 0xf4, 0x54, 0xf2, 0x52, 0x81, 0x63, 0x3f, 0xf7, 0x9e,
	0xf7, 0x5f, 0xed, 0xfd, 0x3d, 0xf8, 0xff, 0x3b, 0x7a, 0xf0, 0xb8, 0x9d, 0xb1, 0x8a, 0x64, 0xc3,
	0xaf, 0x13, 0xbf, 0x1c, 0x2b, 0x70, 0xf4, 0x2b, 0xb2, 0x85, 0xa4, 0xa9, 0xc4, 0x5f, 0xb5, 0xe4,
	0x25, 0x70, 0x01, 0xec, 0x97, 0x65, 0xfb, 0xfc, 0x81, 0xf8, 0x18, 0x86, 0x40, 0x8f, 0x09, 0x7b,
	0x28, 0x15, 0x3c, 0xbb, 0x91, 0xd9, 0x2d, 0xd4, 0x33, 0xf6, 0x2b, 0x46, 0x6c, 0xaf, 0x46, 0x88,
	0x51, 0x24, 0xe9, 0x37, 0x04, 0x09, 0x9e, 0x67, 0x3c, 0x55, 0xae, 0xf4, 0xa3, 0x14, 0xf2, 0xfc,
	0x86, 0x51, 0x58, 0xc0, 0x69, 0x76, 0x12, 0x28, 0x4c, 0xd5, 0x27, 0x9b, 0x4a, 0xe7, 0x2a, 0xc3,
	0x8d, 0x44, 0xf1, 0x30, 0x14, 0xb5, 0xc4, 0x51, 0xf9, 0x91, 0x6c, 0xe6, 0x86, 0x6b, 0xe3, 0xb8,
	0x98, 0x36, 0xe5, 0x9f, 0xa0, 0xb2, 0x93, 0x9b, 0x89, 0x71, 0xc3, 0x69, 0x28, 0x7e, 0x97, 0x6c,
	0x15, 0xa5, 0x31, 0x39, 0xde, 0x02, 0x2f, 0xa5, 0x9e, 0xba, 0x1b, 0x36, 0x0a, 0x0d, 0x8a, 0x84,
	0x3f, 0x8d, 0x31, 0xc2, 0x8b, 0x8d, 0x96, 0xc5, 0xca, 0x50, 0xf1, 0xcc, 0xd4, 0xda, 0xb1, 0xd3,
	0xe5, 0x46, 0xc7, 0xc5, 0x72, 0xb8, 0x46, 0x9e, 0xa4, 0x17, 0xe4, 0x73, 0x0c, 0xac, 0xa4, 0xe5,
	0x95, 0x95, 0x85, 0xba, 0xfb, 0x4f, 0x8f, 0x33, 0xf4, 0xf8, 0xd4, 0x4b, 0x2f, 0xa5, 0xbd, 0x44,
	0xe1, 0x53, 0xb3, 0xdd, 0x78, 0x33, 0xde, 0xc1, 0xca, 0xb9, 0x02, 0xdf, 0xff, 0xe7, 0xa1, 0xe2,
	0x98, 0x3e, 0x89, 0xf0, 0xce, 0x7d, 0xeb, 0xc9, 0xd3, 0x19, 0xae, 0xdb, 0x3f, 0x5f, 0xa1, 0xfb,
	0x42, 0x3d, 0xcd, 0xf3, 0x85, 0x58, 0xc8, 0x4c, 0x07, 0xe4, 0xcd, 0xaa, 0xa4, 0x39, 0x9c, 0x67,
	0xe1, 0x72, 0x56, 0x94, 0xf1, 0x78, 0x3e, 0x90, 0xb6, 0x96, 0x71, 0xc4, 0xc3, 0xa3, 0xb6, 0xae,
	0x65, 0x18, 0xed, 0x77, 0x64, 0x0d, 0x9c, 0x70, 0x35, 0xb0, 0x17, 0x98, 0x27, 0xfe, 0xc5, 0xd1,
	0xf4, 0x11, 0x3c, 0x33, 0xba, 0x50, 0xd3, 0xda, 0xca, 0x9c, 0xbd, 0x6c, 0x46, 0xd3, 0xc7, 0x8e,
	0x16, 0x78, 0xba, 0x86, 0xaf, 0xfe, 0xfe, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x60, 0xa3,
	0x4d, 0x0e, 0x06, 0x00, 0x00,
}
