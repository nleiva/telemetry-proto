// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_lsp.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_levels_level_detailed_lsps_detailed_lsp

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

// An IS-IS LSP
type IsisShLsp_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	LspId                string   `protobuf:"bytes,3,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShLsp_KEYS) Reset()         { *m = IsisShLsp_KEYS{} }
func (m *IsisShLsp_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShLsp_KEYS) ProtoMessage()    {}
func (*IsisShLsp_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_lsp_ffd9d8f2498b6b79, []int{0}
}
func (m *IsisShLsp_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLsp_KEYS.Unmarshal(m, b)
}
func (m *IsisShLsp_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLsp_KEYS.Marshal(b, m, deterministic)
}
func (dst *IsisShLsp_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLsp_KEYS.Merge(dst, src)
}
func (m *IsisShLsp_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShLsp_KEYS.Size(m)
}
func (m *IsisShLsp_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLsp_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLsp_KEYS proto.InternalMessageInfo

func (m *IsisShLsp_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShLsp_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShLsp_KEYS) GetLspId() string {
	if m != nil {
		return m.LspId
	}
	return ""
}

type IsisShLsp struct {
	// Information from the LSP header
	LspHeaderData *IsisShLspHeader `protobuf:"bytes,50,opt,name=lsp_header_data,json=lspHeaderData,proto3" json:"lsp_header_data,omitempty"`
	// LSP as received/sent over the wire, starting from the LSP ID field
	LspBody              []byte   `protobuf:"bytes,51,opt,name=lsp_body,json=lspBody,proto3" json:"lsp_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShLsp) Reset()         { *m = IsisShLsp{} }
func (m *IsisShLsp) String() string { return proto.CompactTextString(m) }
func (*IsisShLsp) ProtoMessage()    {}
func (*IsisShLsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_lsp_ffd9d8f2498b6b79, []int{1}
}
func (m *IsisShLsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLsp.Unmarshal(m, b)
}
func (m *IsisShLsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLsp.Marshal(b, m, deterministic)
}
func (dst *IsisShLsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLsp.Merge(dst, src)
}
func (m *IsisShLsp) XXX_Size() int {
	return xxx_messageInfo_IsisShLsp.Size(m)
}
func (m *IsisShLsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLsp.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLsp proto.InternalMessageInfo

func (m *IsisShLsp) GetLspHeaderData() *IsisShLspHeader {
	if m != nil {
		return m.LspHeaderData
	}
	return nil
}

func (m *IsisShLsp) GetLspBody() []byte {
	if m != nil {
		return m.LspBody
	}
	return nil
}

// Contents of an IS-IS LSP header
type IsisShLspHeader struct {
	// The LSP ID
	LspId string `protobuf:"bytes,1,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	// TRUE if this is a locally generated LSP
	LocalLspFlag bool `protobuf:"varint,2,opt,name=local_lsp_flag,json=localLspFlag,proto3" json:"local_lsp_flag,omitempty"`
	// TRUE if this LSP has not expired
	LspActiveFlag bool `protobuf:"varint,3,opt,name=lsp_active_flag,json=lspActiveFlag,proto3" json:"lsp_active_flag,omitempty"`
	// Time, in seconds, until LSP expiry (if active) or deletion (if expired)
	LspHoldtime uint32 `protobuf:"varint,4,opt,name=lsp_holdtime,json=lspHoldtime,proto3" json:"lsp_holdtime,omitempty"`
	// The LSP sequence number
	LspSequenceNumber uint32 `protobuf:"varint,5,opt,name=lsp_sequence_number,json=lspSequenceNumber,proto3" json:"lsp_sequence_number,omitempty"`
	// The LSP checksum
	LspChecksum uint32 `protobuf:"varint,6,opt,name=lsp_checksum,json=lspChecksum,proto3" json:"lsp_checksum,omitempty"`
	// TRUE if partition repair is supported
	LspParitionRepairSupportedFlag bool `protobuf:"varint,7,opt,name=lsp_parition_repair_supported_flag,json=lspParitionRepairSupportedFlag,proto3" json:"lsp_parition_repair_supported_flag,omitempty"`
	// TRUE if attached bit is set
	LspAttachedFlag bool `protobuf:"varint,8,opt,name=lsp_attached_flag,json=lspAttachedFlag,proto3" json:"lsp_attached_flag,omitempty"`
	// TRUE if the overload bit is set
	LspOverloadedFlag bool `protobuf:"varint,9,opt,name=lsp_overloaded_flag,json=lspOverloadedFlag,proto3" json:"lsp_overloaded_flag,omitempty"`
	// TRUE if the LSP is non-v1a, XXX for testing
	LspNonV1AFlag uint32 `protobuf:"varint,10,opt,name=lsp_non_v1_a_flag,json=lspNonV1AFlag,proto3" json:"lsp_non_v1_a_flag,omitempty"`
	// The type of the IS sourcing the LSP
	LspLevel string `protobuf:"bytes,11,opt,name=lsp_level,json=lspLevel,proto3" json:"lsp_level,omitempty"`
	// The total length of the LSP
	LspLength            uint32   `protobuf:"varint,12,opt,name=lsp_length,json=lspLength,proto3" json:"lsp_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShLspHeader) Reset()         { *m = IsisShLspHeader{} }
func (m *IsisShLspHeader) String() string { return proto.CompactTextString(m) }
func (*IsisShLspHeader) ProtoMessage()    {}
func (*IsisShLspHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_lsp_ffd9d8f2498b6b79, []int{2}
}
func (m *IsisShLspHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLspHeader.Unmarshal(m, b)
}
func (m *IsisShLspHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLspHeader.Marshal(b, m, deterministic)
}
func (dst *IsisShLspHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLspHeader.Merge(dst, src)
}
func (m *IsisShLspHeader) XXX_Size() int {
	return xxx_messageInfo_IsisShLspHeader.Size(m)
}
func (m *IsisShLspHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLspHeader.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLspHeader proto.InternalMessageInfo

func (m *IsisShLspHeader) GetLspId() string {
	if m != nil {
		return m.LspId
	}
	return ""
}

func (m *IsisShLspHeader) GetLocalLspFlag() bool {
	if m != nil {
		return m.LocalLspFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspActiveFlag() bool {
	if m != nil {
		return m.LspActiveFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspHoldtime() uint32 {
	if m != nil {
		return m.LspHoldtime
	}
	return 0
}

func (m *IsisShLspHeader) GetLspSequenceNumber() uint32 {
	if m != nil {
		return m.LspSequenceNumber
	}
	return 0
}

func (m *IsisShLspHeader) GetLspChecksum() uint32 {
	if m != nil {
		return m.LspChecksum
	}
	return 0
}

func (m *IsisShLspHeader) GetLspParitionRepairSupportedFlag() bool {
	if m != nil {
		return m.LspParitionRepairSupportedFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspAttachedFlag() bool {
	if m != nil {
		return m.LspAttachedFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspOverloadedFlag() bool {
	if m != nil {
		return m.LspOverloadedFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspNonV1AFlag() uint32 {
	if m != nil {
		return m.LspNonV1AFlag
	}
	return 0
}

func (m *IsisShLspHeader) GetLspLevel() string {
	if m != nil {
		return m.LspLevel
	}
	return ""
}

func (m *IsisShLspHeader) GetLspLength() uint32 {
	if m != nil {
		return m.LspLength
	}
	return 0
}

func init() {
	proto.RegisterType((*IsisShLsp_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.detailed_lsps.detailed_lsp.isis_sh_lsp_KEYS")
	proto.RegisterType((*IsisShLsp)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.detailed_lsps.detailed_lsp.isis_sh_lsp")
	proto.RegisterType((*IsisShLspHeader)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.detailed_lsps.detailed_lsp.isis_sh_lsp_header")
}

func init() { proto.RegisterFile("isis_sh_lsp.proto", fileDescriptor_isis_sh_lsp_ffd9d8f2498b6b79) }

var fileDescriptor_isis_sh_lsp_ffd9d8f2498b6b79 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0x19, 0xeb, 0x76, 0xdb, 0xb4, 0x55, 0x1b, 0x15, 0x46, 0x44, 0xa9, 0x55, 0x64, 0xf0,
	0x62, 0x60, 0x77, 0x9f, 0xa0, 0xfe, 0xc3, 0x3f, 0x4b, 0x95, 0x29, 0x08, 0xe2, 0x45, 0x48, 0x27,
	0xb1, 0x13, 0xcc, 0x4c, 0x62, 0x4e, 0x5a, 0xdc, 0x27, 0xf1, 0x7d, 0x7c, 0x2f, 0x41, 0x72, 0x32,
	0xb3, 0x8e, 0x78, 0xbd, 0x57, 0xcd, 0x7c, 0xe7, 0x97, 0x2f, 0xe7, 0x3b, 0x49, 0xc9, 0x5c, 0x81,
	0x02, 0x06, 0x15, 0xd3, 0x60, 0x73, 0xeb, 0x8c, 0x37, 0xf4, 0x4b, 0xa9, 0xa0, 0x34, 0x4c, 0x19,
	0x60, 0x3f, 0x1c, 0x2b, 0x75, 0x03, 0x0c, 0x21, 0x63, 0xa5, 0xcb, 0xc3, 0x2a, 0x57, 0x0d, 0x78,
	0xde, 0x94, 0xf2, 0xef, 0x2a, 0xd7, 0xf2, 0x20, 0x35, 0xc4, 0x9f, 0x5c, 0x48, 0xcf, 0x95, 0x96,
	0x22, 0x98, 0xc2, 0x3f, 0x5f, 0x4b, 0x41, 0x6e, 0xf5, 0x4e, 0x64, 0xef, 0x5f, 0x7d, 0xde, 0xd0,
	0xc7, 0x64, 0xd6, 0xf9, 0xb0, 0x86, 0xd7, 0x32, 0x4d, 0x16, 0x49, 0x36, 0x2e, 0xa6, 0x9d, 0xb8,
	0xe6, 0xb5, 0xa4, 0x77, 0xc8, 0x11, 0x9a, 0xa7, 0xd7, 0xb0, 0x18, 0x3f, 0xe8, 0x5d, 0x32, 0x0c,
	0x36, 0x4a, 0xa4, 0x83, 0x56, 0x06, 0xfb, 0x56, 0x2c, 0x7f, 0x25, 0x64, 0xd2, 0x3b, 0x86, 0xfe,
	0x4c, 0xc8, 0xcd, 0xc0, 0x55, 0x92, 0x0b, 0xe9, 0x98, 0xe0, 0x9e, 0xa7, 0xa7, 0x8b, 0x24, 0x9b,
	0x9c, 0x9a, 0xfc, 0x0a, 0xd3, 0xe6, 0xfd, 0xa8, 0xf1, 0xec, 0x62, 0xa6, 0xc1, 0xbe, 0xc1, 0xe5,
	0x4b, 0xee, 0x39, 0xbd, 0x47, 0x46, 0xa1, 0xb8, 0x35, 0xe2, 0x22, 0x3d, 0x5b, 0x24, 0xd9, 0xb4,
	0x38, 0xd6, 0x60, 0x9f, 0x1b, 0x71, 0xb1, 0xfc, 0x3d, 0x20, 0xf4, 0x7f, 0x83, 0x5e, 0xe4, 0xa4,
	0x17, 0x99, 0x3e, 0x21, 0x37, 0xb4, 0x29, 0xb9, 0x46, 0xf4, 0xab, 0xe6, 0x3b, 0x1c, 0xd4, 0xa8,
	0x98, 0xa2, 0x7a, 0x0e, 0xf6, 0xb5, 0xe6, 0x3b, 0xfa, 0x34, 0xce, 0x81, 0x97, 0x5e, 0x1d, 0x64,
	0xc4, 0x06, 0x88, 0x85, 0xb6, 0x56, 0xa8, 0x22, 0xf7, 0x88, 0x4c, 0xf1, 0x48, 0xa3, 0x85, 0x57,
	0xb5, 0x4c, 0xaf, 0x2f, 0x92, 0x6c, 0x56, 0x4c, 0x42, 0xef, 0xad, 0x44, 0x73, 0x72, 0x3b, 0x20,
	0x20, 0xbf, 0xef, 0x25, 0xde, 0xdc, 0xbe, 0xde, 0x4a, 0x97, 0x1e, 0x21, 0x39, 0xd7, 0x60, 0x37,
	0x6d, 0x65, 0x8d, 0x85, 0xce, 0xb2, 0xac, 0x64, 0xf9, 0x0d, 0xf6, 0x75, 0x3a, 0xbc, 0xb4, 0x7c,
	0xd1, 0x4a, 0xf4, 0x1d, 0x59, 0x06, 0xc4, 0x72, 0xa7, 0xbc, 0x32, 0x0d, 0x73, 0xd2, 0x72, 0xe5,
	0x18, 0xec, 0xad, 0x35, 0xce, 0x4b, 0x11, 0x1b, 0x3e, 0xc6, 0x86, 0x1f, 0x6a, 0xb0, 0x1f, 0x5b,
	0xb0, 0x40, 0x6e, 0xd3, 0x61, 0x98, 0xe0, 0x19, 0x99, 0x63, 0x52, 0xef, 0x79, 0x59, 0x75, 0x5b,
	0x47, 0xb8, 0x35, 0x8c, 0x60, 0xd5, 0xea, 0xc8, 0xb6, 0x51, 0xcc, 0x41, 0x3a, 0x6d, 0xb8, 0xe8,
	0xe8, 0x31, 0xd2, 0xc1, 0xe6, 0xc3, 0x65, 0x05, 0xf9, 0x2c, 0x7a, 0x37, 0xa6, 0x61, 0x87, 0x13,
	0xc6, 0x23, 0x4d, 0x30, 0x4f, 0x98, 0xe3, 0xda, 0x34, 0x9f, 0x4e, 0x56, 0x48, 0xde, 0x27, 0xe3,
	0x40, 0xc6, 0x97, 0x3b, 0xc1, 0xfb, 0x0a, 0xf7, 0x7d, 0x8e, 0x8f, 0xf7, 0x01, 0x21, 0xb1, 0xd8,
	0xec, 0x7c, 0x95, 0x4e, 0x71, 0xff, 0x18, 0xab, 0x41, 0xd8, 0x0e, 0xf1, 0xef, 0x78, 0xf6, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x10, 0x9e, 0xd3, 0xa3, 0x03, 0x00, 0x00,
}
