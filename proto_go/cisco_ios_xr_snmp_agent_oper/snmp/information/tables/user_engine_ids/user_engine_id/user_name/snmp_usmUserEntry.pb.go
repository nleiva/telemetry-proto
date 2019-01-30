// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_usmUserEntry.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_tables_user_engine_ids_user_engine_id_user_name

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

// SNMP usmUserTable Information
type SnmpUsmUserEntry_KEYS struct {
	EngineId             string   `protobuf:"bytes,1,opt,name=engine_id,json=engineId,proto3" json:"engine_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpUsmUserEntry_KEYS) Reset()         { *m = SnmpUsmUserEntry_KEYS{} }
func (m *SnmpUsmUserEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpUsmUserEntry_KEYS) ProtoMessage()    {}
func (*SnmpUsmUserEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a45a4fc5d843690, []int{0}
}

func (m *SnmpUsmUserEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpUsmUserEntry_KEYS.Unmarshal(m, b)
}
func (m *SnmpUsmUserEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpUsmUserEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpUsmUserEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpUsmUserEntry_KEYS.Merge(m, src)
}
func (m *SnmpUsmUserEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpUsmUserEntry_KEYS.Size(m)
}
func (m *SnmpUsmUserEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpUsmUserEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpUsmUserEntry_KEYS proto.InternalMessageInfo

func (m *SnmpUsmUserEntry_KEYS) GetEngineId() string {
	if m != nil {
		return m.EngineId
	}
	return ""
}

func (m *SnmpUsmUserEntry_KEYS) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type SnmpUsmUserEntry struct {
	// Storage type
	UsmUserStorageType uint32 `protobuf:"varint,50,opt,name=usm_user_storage_type,json=usmUserStorageType,proto3" json:"usm_user_storage_type,omitempty"`
	// Status of this user
	UsmUserStatus        uint32   `protobuf:"varint,51,opt,name=usm_user_status,json=usmUserStatus,proto3" json:"usm_user_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpUsmUserEntry) Reset()         { *m = SnmpUsmUserEntry{} }
func (m *SnmpUsmUserEntry) String() string { return proto.CompactTextString(m) }
func (*SnmpUsmUserEntry) ProtoMessage()    {}
func (*SnmpUsmUserEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a45a4fc5d843690, []int{1}
}

func (m *SnmpUsmUserEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpUsmUserEntry.Unmarshal(m, b)
}
func (m *SnmpUsmUserEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpUsmUserEntry.Marshal(b, m, deterministic)
}
func (m *SnmpUsmUserEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpUsmUserEntry.Merge(m, src)
}
func (m *SnmpUsmUserEntry) XXX_Size() int {
	return xxx_messageInfo_SnmpUsmUserEntry.Size(m)
}
func (m *SnmpUsmUserEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpUsmUserEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpUsmUserEntry proto.InternalMessageInfo

func (m *SnmpUsmUserEntry) GetUsmUserStorageType() uint32 {
	if m != nil {
		return m.UsmUserStorageType
	}
	return 0
}

func (m *SnmpUsmUserEntry) GetUsmUserStatus() uint32 {
	if m != nil {
		return m.UsmUserStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpUsmUserEntry_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.tables.user_engine_ids.user_engine_id.user_name.snmp_usmUserEntry_KEYS")
	proto.RegisterType((*SnmpUsmUserEntry)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.tables.user_engine_ids.user_engine_id.user_name.snmp_usmUserEntry")
}

func init() { proto.RegisterFile("snmp_usmUserEntry.proto", fileDescriptor_6a45a4fc5d843690) }

var fileDescriptor_6a45a4fc5d843690 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x07, 0xb1, 0x0b, 0x45, 0x5c, 0x50, 0x03, 0x5e, 0x8a, 0x07, 0xe9, 0x69, 0x41,
	0xfb, 0x1b, 0x7a, 0x10, 0xc1, 0x43, 0xaa, 0x07, 0x0f, 0x32, 0x6c, 0xdb, 0x31, 0x2c, 0xb8, 0xb3,
	0xcb, 0xcc, 0x04, 0xcc, 0xbf, 0x97, 0x4d, 0x24, 0x48, 0x73, 0x7c, 0xef, 0x7d, 0xef, 0x0d, 0x8c,
	0xb9, 0x15, 0x8a, 0x19, 0x3a, 0x89, 0xef, 0x82, 0xbc, 0x25, 0xe5, 0xde, 0x65, 0x4e, 0x9a, 0xec,
	0xe7, 0x21, 0xc8, 0x21, 0x41, 0x48, 0x02, 0x3f, 0x0c, 0x03, 0xe5, 0x5b, 0x24, 0x85, 0x94, 0x91,
	0x5d, 0xd1, 0x2e, 0xd0, 0x57, 0xe2, 0xe8, 0x35, 0x24, 0x72, 0xea, 0xf7, 0xdf, 0x28, 0xae, 0x13,
	0x64, 0x40, 0x6a, 0x03, 0x21, 0x84, 0xe3, 0xa9, 0x1e, 0x25, 0xf9, 0x88, 0xf7, 0x8d, 0xb9, 0x99,
	0x5d, 0x86, 0x97, 0xed, 0xc7, 0xce, 0xde, 0x99, 0xc5, 0x54, 0xa8, 0xab, 0x55, 0xb5, 0x5e, 0x34,
	0x17, 0xa3, 0xf1, 0x7c, 0x2c, 0xe1, 0xb4, 0x51, 0x9f, 0x8d, 0x61, 0x31, 0x5e, 0xcb, 0x26, 0x99,
	0xab, 0xd9, 0xa6, 0x7d, 0x34, 0xd7, 0x9d, 0x44, 0x18, 0x5a, 0xa2, 0x89, 0x7d, 0x8b, 0xa0, 0x7d,
	0xc6, 0xfa, 0x69, 0x55, 0xad, 0x97, 0x8d, 0xfd, 0x83, 0x77, 0x63, 0xf4, 0xd6, 0x67, 0xb4, 0x0f,
	0xe6, 0xf2, 0x5f, 0xc5, 0x6b, 0x27, 0xf5, 0x66, 0x80, 0x97, 0x13, 0x5c, 0xcc, 0xfd, 0xf9, 0xf0,
	0xa9, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x4e, 0x30, 0xca, 0x44, 0x01, 0x00, 0x00,
}
