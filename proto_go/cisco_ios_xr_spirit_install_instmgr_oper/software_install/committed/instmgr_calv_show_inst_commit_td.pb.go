// Code generated by protoc-gen-go. DO NOT EDIT.
// source: instmgr_calv_show_inst_commit_td.proto

package cisco_ios_xr_spirit_install_instmgr_oper_software_install_committed

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

// Calvados show install commit info
type InstmgrCalvShowInstCommitTd_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrCalvShowInstCommitTd_KEYS) Reset()         { *m = InstmgrCalvShowInstCommitTd_KEYS{} }
func (m *InstmgrCalvShowInstCommitTd_KEYS) String() string { return proto.CompactTextString(m) }
func (*InstmgrCalvShowInstCommitTd_KEYS) ProtoMessage()    {}
func (*InstmgrCalvShowInstCommitTd_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9dc69a01bf98677, []int{0}
}

func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Unmarshal(m, b)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Marshal(b, m, deterministic)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Merge(m, src)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_Size() int {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.Size(m)
}
func (m *InstmgrCalvShowInstCommitTd_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrCalvShowInstCommitTd_KEYS proto.InternalMessageInfo

type InstmgrCalvShowInstCommitTd struct {
	CommittedPackageInfo []*InstmgrCalvShowInstCommitRow `protobuf:"bytes,50,rep,name=committed_package_info,json=committedPackageInfo,proto3" json:"committed_package_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InstmgrCalvShowInstCommitTd) Reset()         { *m = InstmgrCalvShowInstCommitTd{} }
func (m *InstmgrCalvShowInstCommitTd) String() string { return proto.CompactTextString(m) }
func (*InstmgrCalvShowInstCommitTd) ProtoMessage()    {}
func (*InstmgrCalvShowInstCommitTd) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9dc69a01bf98677, []int{1}
}

func (m *InstmgrCalvShowInstCommitTd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd.Unmarshal(m, b)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd.Marshal(b, m, deterministic)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd.Merge(m, src)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_Size() int {
	return xxx_messageInfo_InstmgrCalvShowInstCommitTd.Size(m)
}
func (m *InstmgrCalvShowInstCommitTd) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrCalvShowInstCommitTd.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrCalvShowInstCommitTd proto.InternalMessageInfo

func (m *InstmgrCalvShowInstCommitTd) GetCommittedPackageInfo() []*InstmgrCalvShowInstCommitRow {
	if m != nil {
		return m.CommittedPackageInfo
	}
	return nil
}

type InstmgrCalvShowInstCommitRow struct {
	ErrorMessage              string   `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Location                  string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	NodeType                  string   `protobuf:"bytes,3,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	BootPartitionName         string   `protobuf:"bytes,4,opt,name=boot_partition_name,json=bootPartitionName,proto3" json:"boot_partition_name,omitempty"`
	NumberOfCommittedPackages uint32   `protobuf:"varint,5,opt,name=number_of_committed_packages,json=numberOfCommittedPackages,proto3" json:"number_of_committed_packages,omitempty"`
	CommittedPackages         string   `protobuf:"bytes,6,opt,name=committed_packages,json=committedPackages,proto3" json:"committed_packages,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *InstmgrCalvShowInstCommitRow) Reset()         { *m = InstmgrCalvShowInstCommitRow{} }
func (m *InstmgrCalvShowInstCommitRow) String() string { return proto.CompactTextString(m) }
func (*InstmgrCalvShowInstCommitRow) ProtoMessage()    {}
func (*InstmgrCalvShowInstCommitRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9dc69a01bf98677, []int{2}
}

func (m *InstmgrCalvShowInstCommitRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrCalvShowInstCommitRow.Unmarshal(m, b)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrCalvShowInstCommitRow.Marshal(b, m, deterministic)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrCalvShowInstCommitRow.Merge(m, src)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_Size() int {
	return xxx_messageInfo_InstmgrCalvShowInstCommitRow.Size(m)
}
func (m *InstmgrCalvShowInstCommitRow) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrCalvShowInstCommitRow.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrCalvShowInstCommitRow proto.InternalMessageInfo

func (m *InstmgrCalvShowInstCommitRow) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetBootPartitionName() string {
	if m != nil {
		return m.BootPartitionName
	}
	return ""
}

func (m *InstmgrCalvShowInstCommitRow) GetNumberOfCommittedPackages() uint32 {
	if m != nil {
		return m.NumberOfCommittedPackages
	}
	return 0
}

func (m *InstmgrCalvShowInstCommitRow) GetCommittedPackages() string {
	if m != nil {
		return m.CommittedPackages
	}
	return ""
}

func init() {
	proto.RegisterType((*InstmgrCalvShowInstCommitTd_KEYS)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.committed.instmgr_calv_show_inst_commit_td_KEYS")
	proto.RegisterType((*InstmgrCalvShowInstCommitTd)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.committed.instmgr_calv_show_inst_commit_td")
	proto.RegisterType((*InstmgrCalvShowInstCommitRow)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.committed.instmgr_calv_show_inst_commit_row")
}

func init() {
	proto.RegisterFile("instmgr_calv_show_inst_commit_td.proto", fileDescriptor_a9dc69a01bf98677)
}

var fileDescriptor_a9dc69a01bf98677 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x4e, 0xf3, 0x40,
	0x10, 0xc5, 0xe5, 0xe4, 0xfb, 0x22, 0xb2, 0x90, 0x82, 0x05, 0x21, 0xf3, 0xa7, 0x30, 0x41, 0x40,
	0x1a, 0x5c, 0x84, 0x03, 0x50, 0x44, 0x14, 0x08, 0x01, 0x51, 0xa0, 0xa1, 0x1a, 0x6d, 0x9c, 0x71,
	0x58, 0xe1, 0xf5, 0xac, 0x66, 0x17, 0x42, 0xee, 0x42, 0xc3, 0x51, 0xb8, 0x19, 0xb2, 0xad, 0xa4,
	0x08, 0x08, 0x37, 0x94, 0x7e, 0xef, 0x37, 0xf3, 0x9e, 0x47, 0x2b, 0x4e, 0x74, 0xee, 0xbc, 0x99,
	0x32, 0x24, 0x2a, 0x7b, 0x05, 0xf7, 0x44, 0x33, 0x28, 0x14, 0x48, 0xc8, 0x18, 0xed, 0xc1, 0x4f,
	0x62, 0xcb, 0xe4, 0x49, 0x0e, 0x12, 0xed, 0x12, 0x02, 0x4d, 0x0e, 0xde, 0x18, 0x9c, 0xd5, 0xac,
	0x7d, 0x49, 0xaa, 0x2c, 0x83, 0xc5, 0x0e, 0xb2, 0xc8, 0xb1, 0xa3, 0xd4, 0xcf, 0x14, 0xe3, 0xc2,
	0x8d, 0xab, 0x55, 0x1e, 0x27, 0xdd, 0x53, 0x71, 0x5c, 0x17, 0x07, 0xd7, 0x97, 0x8f, 0xf7, 0xdd,
	0xcf, 0x40, 0x44, 0x75, 0xa4, 0x7c, 0x0f, 0xc4, 0xce, 0x72, 0x37, 0x58, 0x95, 0x3c, 0xab, 0x69,
	0x91, 0x9a, 0x52, 0xd8, 0x8f, 0x9a, 0xbd, 0xf5, 0x7e, 0x1a, 0xff, 0x41, 0xe9, 0xf8, 0xf7, 0x1e,
	0x4c, 0xb3, 0xd1, 0xf6, 0x12, 0x1e, 0x56, 0x25, 0xae, 0xf2, 0x94, 0xba, 0x1f, 0x0d, 0x71, 0x58,
	0x3b, 0x2b, 0x8f, 0x44, 0x07, 0x99, 0x89, 0xc1, 0xa0, 0x73, 0x6a, 0x8a, 0x61, 0x10, 0x05, 0xbd,
	0xf6, 0x68, 0xa3, 0x14, 0x6f, 0x2a, 0x4d, 0xee, 0x89, 0xb5, 0x8c, 0x12, 0xe5, 0x35, 0xe5, 0x61,
	0xa3, 0xf4, 0x97, 0xdf, 0x72, 0x5f, 0xb4, 0x73, 0x9a, 0x20, 0xf8, 0xb9, 0xc5, 0xb0, 0x59, 0x99,
	0x85, 0xf0, 0x30, 0xb7, 0x28, 0x63, 0xb1, 0x35, 0x26, 0xf2, 0x60, 0x15, 0x7b, 0x5d, 0xe0, 0x90,
	0x2b, 0x83, 0xe1, 0xbf, 0x12, 0xdb, 0x2c, 0xac, 0xe1, 0xc2, 0xb9, 0x55, 0x06, 0xe5, 0x85, 0x38,
	0xc8, 0x5f, 0xcc, 0x18, 0x19, 0x28, 0x85, 0x6f, 0xb7, 0x75, 0xe1, 0xff, 0x28, 0xe8, 0x75, 0x46,
	0xbb, 0x15, 0x73, 0x97, 0x0e, 0x56, 0xfe, 0xdb, 0xc9, 0x33, 0x21, 0x7f, 0x18, 0x6b, 0x55, 0x79,
	0xab, 0x67, 0x72, 0xe3, 0x56, 0xf9, 0xb8, 0xce, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0xee,
	0x02, 0x63, 0x86, 0x02, 0x00, 0x00,
}
