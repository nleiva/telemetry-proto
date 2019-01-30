// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im_if_summary_info.proto

package cisco_ios_xr_pfi_im_cmd_oper_interfaces_node_type_sets_node_type_set_interface_summary

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

// Interface summary bag
type ImIfSummaryInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	TypeSetName          string   `protobuf:"bytes,2,opt,name=type_set_name,json=typeSetName,proto3" json:"type_set_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImIfSummaryInfo_KEYS) Reset()         { *m = ImIfSummaryInfo_KEYS{} }
func (m *ImIfSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ImIfSummaryInfo_KEYS) ProtoMessage()    {}
func (*ImIfSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{0}
}

func (m *ImIfSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *ImIfSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ImIfSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfSummaryInfo_KEYS.Merge(m, src)
}
func (m *ImIfSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ImIfSummaryInfo_KEYS.Size(m)
}
func (m *ImIfSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfSummaryInfo_KEYS proto.InternalMessageInfo

func (m *ImIfSummaryInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ImIfSummaryInfo_KEYS) GetTypeSetName() string {
	if m != nil {
		return m.TypeSetName
	}
	return ""
}

type ImIfSummaryInfo struct {
	// List of per interface type summary information
	InterfaceTypeList []*ImIfTypeSummarySt `protobuf:"bytes,50,rep,name=interface_type_list,json=interfaceTypeList,proto3" json:"interface_type_list,omitempty"`
	// Counts for all interfaces
	InterfaceCounts      *ImIfGroupCountsSt `protobuf:"bytes,51,opt,name=interface_counts,json=interfaceCounts,proto3" json:"interface_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ImIfSummaryInfo) Reset()         { *m = ImIfSummaryInfo{} }
func (m *ImIfSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*ImIfSummaryInfo) ProtoMessage()    {}
func (*ImIfSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{1}
}

func (m *ImIfSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfSummaryInfo.Unmarshal(m, b)
}
func (m *ImIfSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfSummaryInfo.Marshal(b, m, deterministic)
}
func (m *ImIfSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfSummaryInfo.Merge(m, src)
}
func (m *ImIfSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_ImIfSummaryInfo.Size(m)
}
func (m *ImIfSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfSummaryInfo proto.InternalMessageInfo

func (m *ImIfSummaryInfo) GetInterfaceTypeList() []*ImIfTypeSummarySt {
	if m != nil {
		return m.InterfaceTypeList
	}
	return nil
}

func (m *ImIfSummaryInfo) GetInterfaceCounts() *ImIfGroupCountsSt {
	if m != nil {
		return m.InterfaceCounts
	}
	return nil
}

type ImIfGroupCountsSt struct {
	// Number of interfaces
	InterfaceCount uint32 `protobuf:"varint,1,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	// Number of interfaces in UP state
	UpInterfaceCount uint32 `protobuf:"varint,2,opt,name=up_interface_count,json=upInterfaceCount,proto3" json:"up_interface_count,omitempty"`
	// Number of interfaces in DOWN state
	DownInterfaceCount uint32 `protobuf:"varint,3,opt,name=down_interface_count,json=downInterfaceCount,proto3" json:"down_interface_count,omitempty"`
	// Number of interfaces in an ADMINDOWN state
	AdminDownInterfaceCount uint32   `protobuf:"varint,4,opt,name=admin_down_interface_count,json=adminDownInterfaceCount,proto3" json:"admin_down_interface_count,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ImIfGroupCountsSt) Reset()         { *m = ImIfGroupCountsSt{} }
func (m *ImIfGroupCountsSt) String() string { return proto.CompactTextString(m) }
func (*ImIfGroupCountsSt) ProtoMessage()    {}
func (*ImIfGroupCountsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{2}
}

func (m *ImIfGroupCountsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfGroupCountsSt.Unmarshal(m, b)
}
func (m *ImIfGroupCountsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfGroupCountsSt.Marshal(b, m, deterministic)
}
func (m *ImIfGroupCountsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfGroupCountsSt.Merge(m, src)
}
func (m *ImIfGroupCountsSt) XXX_Size() int {
	return xxx_messageInfo_ImIfGroupCountsSt.Size(m)
}
func (m *ImIfGroupCountsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfGroupCountsSt.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfGroupCountsSt proto.InternalMessageInfo

func (m *ImIfGroupCountsSt) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *ImIfGroupCountsSt) GetUpInterfaceCount() uint32 {
	if m != nil {
		return m.UpInterfaceCount
	}
	return 0
}

func (m *ImIfGroupCountsSt) GetDownInterfaceCount() uint32 {
	if m != nil {
		return m.DownInterfaceCount
	}
	return 0
}

func (m *ImIfGroupCountsSt) GetAdminDownInterfaceCount() uint32 {
	if m != nil {
		return m.AdminDownInterfaceCount
	}
	return 0
}

type ImIfTypeSummarySt struct {
	// Name of the interface type
	InterfaceTypeName string `protobuf:"bytes,1,opt,name=interface_type_name,json=interfaceTypeName,proto3" json:"interface_type_name,omitempty"`
	// Description of the interface type
	InterfaceTypeDescription string `protobuf:"bytes,2,opt,name=interface_type_description,json=interfaceTypeDescription,proto3" json:"interface_type_description,omitempty"`
	// Counts for interfaces of this type
	InterfaceCounts      *ImIfGroupCountsSt `protobuf:"bytes,3,opt,name=interface_counts,json=interfaceCounts,proto3" json:"interface_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ImIfTypeSummarySt) Reset()         { *m = ImIfTypeSummarySt{} }
func (m *ImIfTypeSummarySt) String() string { return proto.CompactTextString(m) }
func (*ImIfTypeSummarySt) ProtoMessage()    {}
func (*ImIfTypeSummarySt) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f1e7e171affb549, []int{3}
}

func (m *ImIfTypeSummarySt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImIfTypeSummarySt.Unmarshal(m, b)
}
func (m *ImIfTypeSummarySt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImIfTypeSummarySt.Marshal(b, m, deterministic)
}
func (m *ImIfTypeSummarySt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImIfTypeSummarySt.Merge(m, src)
}
func (m *ImIfTypeSummarySt) XXX_Size() int {
	return xxx_messageInfo_ImIfTypeSummarySt.Size(m)
}
func (m *ImIfTypeSummarySt) XXX_DiscardUnknown() {
	xxx_messageInfo_ImIfTypeSummarySt.DiscardUnknown(m)
}

var xxx_messageInfo_ImIfTypeSummarySt proto.InternalMessageInfo

func (m *ImIfTypeSummarySt) GetInterfaceTypeName() string {
	if m != nil {
		return m.InterfaceTypeName
	}
	return ""
}

func (m *ImIfTypeSummarySt) GetInterfaceTypeDescription() string {
	if m != nil {
		return m.InterfaceTypeDescription
	}
	return ""
}

func (m *ImIfTypeSummarySt) GetInterfaceCounts() *ImIfGroupCountsSt {
	if m != nil {
		return m.InterfaceCounts
	}
	return nil
}

func init() {
	proto.RegisterType((*ImIfSummaryInfo_KEYS)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.node_type_sets.node_type_set.interface_summary.im_if_summary_info_KEYS")
	proto.RegisterType((*ImIfSummaryInfo)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.node_type_sets.node_type_set.interface_summary.im_if_summary_info")
	proto.RegisterType((*ImIfGroupCountsSt)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.node_type_sets.node_type_set.interface_summary.im_if_group_counts_st")
	proto.RegisterType((*ImIfTypeSummarySt)(nil), "cisco_ios_xr_pfi_im_cmd_oper.interfaces.node_type_sets.node_type_set.interface_summary.im_if_type_summary_st")
}

func init() { proto.RegisterFile("im_if_summary_info.proto", fileDescriptor_1f1e7e171affb549) }

var fileDescriptor_1f1e7e171affb549 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xfa, 0xe7, 0x8f, 0xdd, 0x52, 0x5b, 0x57, 0xa5, 0xa1, 0x5e, 0x4a, 0x2e, 0xf6,
	0x20, 0x41, 0xda, 0xa3, 0xde, 0xac, 0x07, 0x51, 0x3c, 0xb4, 0x22, 0xe8, 0x65, 0x88, 0xc9, 0x46,
	0x06, 0xdc, 0x17, 0xb2, 0x1b, 0xb4, 0xdf, 0xc4, 0x8b, 0xe0, 0x37, 0xf2, 0x13, 0xf8, 0x5d, 0x24,
	0xdb, 0x97, 0x98, 0xa4, 0xd7, 0x82, 0xc7, 0xcc, 0xfe, 0x9e, 0x99, 0x67, 0x98, 0x87, 0x10, 0x0f,
	0x39, 0x60, 0x02, 0x3a, 0xe3, 0x3c, 0x4c, 0xe7, 0x80, 0x22, 0x91, 0x81, 0x4a, 0xa5, 0x91, 0xf4,
	0x3e, 0x42, 0x1d, 0x49, 0x40, 0xa9, 0xe1, 0x2d, 0x05, 0x95, 0x20, 0x20, 0x87, 0x88, 0xc7, 0x20,
	0x15, 0x4b, 0x03, 0x14, 0x86, 0xa5, 0x49, 0x18, 0x31, 0x1d, 0x08, 0x19, 0x33, 0x30, 0x73, 0xc5,
	0x40, 0x33, 0x53, 0xf9, 0x2c, 0xb8, 0xd5, 0x08, 0xff, 0x91, 0xf4, 0xea, 0x33, 0xe1, 0xfa, 0xf2,
	0x61, 0x46, 0x8f, 0x48, 0xd3, 0xaa, 0x45, 0xc8, 0x99, 0xe7, 0x0c, 0x9c, 0x61, 0x73, 0xba, 0x93,
	0x17, 0x6e, 0x43, 0xce, 0xa8, 0x4f, 0xda, 0xab, 0xae, 0x0b, 0xc0, 0xb5, 0x40, 0x2b, 0x2f, 0xce,
	0x98, 0xc9, 0x19, 0xff, 0xcb, 0x25, 0xb4, 0xde, 0x9c, 0x7e, 0x38, 0x64, 0xbf, 0x30, 0x62, 0xbb,
	0xbc, 0xa0, 0x36, 0xde, 0x68, 0xd0, 0x18, 0xb6, 0x46, 0x3c, 0xd8, 0xce, 0xa6, 0xc1, 0xc2, 0xc9,
	0x02, 0x58, 0xda, 0xd1, 0x66, 0xba, 0xb7, 0x06, 0xef, 0xe6, 0x8a, 0xdd, 0xa0, 0x36, 0xf4, 0xdd,
	0x21, 0xdd, 0x42, 0x1e, 0xc9, 0x4c, 0x18, 0xed, 0x8d, 0x07, 0xce, 0xf6, 0xcd, 0x3d, 0xa7, 0x32,
	0x53, 0xcb, 0x89, 0xb9, 0xb9, 0xce, 0x1a, 0xbc, 0xb0, 0x35, 0xff, 0xdb, 0x21, 0x87, 0x1b, 0x51,
	0x7a, 0x4c, 0x3a, 0x15, 0xcf, 0xf6, 0x64, 0xed, 0xe9, 0x6e, 0xb9, 0x07, 0x3d, 0x21, 0x34, 0x53,
	0x50, 0x65, 0x5d, 0xcb, 0x76, 0x33, 0x75, 0x55, 0xa6, 0x4f, 0xc9, 0x41, 0x2c, 0x5f, 0x45, 0x8d,
	0x6f, 0x58, 0x9e, 0xe6, 0x6f, 0x15, 0xc5, 0x19, 0xe9, 0x87, 0x31, 0x47, 0x01, 0x1b, 0x75, 0xff,
	0xac, 0xae, 0x67, 0x89, 0x49, 0x4d, 0xec, 0x7f, 0xba, 0xab, 0xfd, 0x2a, 0x77, 0xa2, 0x41, 0x2d,
	0x33, 0xbf, 0x62, 0x59, 0x3e, 0xa2, 0xcd, 0xe7, 0x39, 0xe9, 0x57, 0xf8, 0x98, 0xe9, 0x28, 0x45,
	0x65, 0x50, 0x8a, 0x65, 0x58, 0xbd, 0x92, 0x6c, 0x52, 0xbc, 0x6f, 0x8e, 0x40, 0xe3, 0x2f, 0x44,
	0xe0, 0xe9, 0xbf, 0xfd, 0x1f, 0x8c, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x91, 0xeb, 0x8d,
	0x2b, 0x04, 0x00, 0x00,
}
