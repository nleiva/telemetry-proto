// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_corr_trap_buffer_bag.proto

package cisco_ios_xr_snmp_agent_oper_snmp_correlator_traps_trap

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

// Trap Buffer Record
type SnmpCorrTrapBufferBag_KEYS struct {
	EntryId              uint32   `protobuf:"varint,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpCorrTrapBufferBag_KEYS) Reset()         { *m = SnmpCorrTrapBufferBag_KEYS{} }
func (m *SnmpCorrTrapBufferBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpCorrTrapBufferBag_KEYS) ProtoMessage()    {}
func (*SnmpCorrTrapBufferBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_snmp_corr_trap_buffer_bag_a3c3f78d3369ba7a, []int{0}
}
func (m *SnmpCorrTrapBufferBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpCorrTrapBufferBag_KEYS.Unmarshal(m, b)
}
func (m *SnmpCorrTrapBufferBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpCorrTrapBufferBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *SnmpCorrTrapBufferBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpCorrTrapBufferBag_KEYS.Merge(dst, src)
}
func (m *SnmpCorrTrapBufferBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpCorrTrapBufferBag_KEYS.Size(m)
}
func (m *SnmpCorrTrapBufferBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpCorrTrapBufferBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpCorrTrapBufferBag_KEYS proto.InternalMessageInfo

func (m *SnmpCorrTrapBufferBag_KEYS) GetEntryId() uint32 {
	if m != nil {
		return m.EntryId
	}
	return 0
}

type SnmpCorrTrapBufferBag struct {
	// Correlation ID
	CorrelationId uint32 `protobuf:"varint,50,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// True if this is the rootcause
	IsRootcause bool `protobuf:"varint,51,opt,name=is_rootcause,json=isRootcause,proto3" json:"is_rootcause,omitempty"`
	// Correlation rule name
	RuleName string `protobuf:"bytes,52,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	// Correlated trap information
	TrapInfo             *SnmpCorrTrapBag `protobuf:"bytes,53,opt,name=trap_info,json=trapInfo,proto3" json:"trap_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SnmpCorrTrapBufferBag) Reset()         { *m = SnmpCorrTrapBufferBag{} }
func (m *SnmpCorrTrapBufferBag) String() string { return proto.CompactTextString(m) }
func (*SnmpCorrTrapBufferBag) ProtoMessage()    {}
func (*SnmpCorrTrapBufferBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_snmp_corr_trap_buffer_bag_a3c3f78d3369ba7a, []int{1}
}
func (m *SnmpCorrTrapBufferBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpCorrTrapBufferBag.Unmarshal(m, b)
}
func (m *SnmpCorrTrapBufferBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpCorrTrapBufferBag.Marshal(b, m, deterministic)
}
func (dst *SnmpCorrTrapBufferBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpCorrTrapBufferBag.Merge(dst, src)
}
func (m *SnmpCorrTrapBufferBag) XXX_Size() int {
	return xxx_messageInfo_SnmpCorrTrapBufferBag.Size(m)
}
func (m *SnmpCorrTrapBufferBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpCorrTrapBufferBag.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpCorrTrapBufferBag proto.InternalMessageInfo

func (m *SnmpCorrTrapBufferBag) GetCorrelationId() uint32 {
	if m != nil {
		return m.CorrelationId
	}
	return 0
}

func (m *SnmpCorrTrapBufferBag) GetIsRootcause() bool {
	if m != nil {
		return m.IsRootcause
	}
	return false
}

func (m *SnmpCorrTrapBufferBag) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func (m *SnmpCorrTrapBufferBag) GetTrapInfo() *SnmpCorrTrapBag {
	if m != nil {
		return m.TrapInfo
	}
	return nil
}

type SnmpCorrVbind struct {
	// OID of the varbind
	Oid string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	// Value of the varbind
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpCorrVbind) Reset()         { *m = SnmpCorrVbind{} }
func (m *SnmpCorrVbind) String() string { return proto.CompactTextString(m) }
func (*SnmpCorrVbind) ProtoMessage()    {}
func (*SnmpCorrVbind) Descriptor() ([]byte, []int) {
	return fileDescriptor_snmp_corr_trap_buffer_bag_a3c3f78d3369ba7a, []int{2}
}
func (m *SnmpCorrVbind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpCorrVbind.Unmarshal(m, b)
}
func (m *SnmpCorrVbind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpCorrVbind.Marshal(b, m, deterministic)
}
func (dst *SnmpCorrVbind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpCorrVbind.Merge(dst, src)
}
func (m *SnmpCorrVbind) XXX_Size() int {
	return xxx_messageInfo_SnmpCorrVbind.Size(m)
}
func (m *SnmpCorrVbind) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpCorrVbind.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpCorrVbind proto.InternalMessageInfo

func (m *SnmpCorrVbind) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *SnmpCorrVbind) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Trap Information
type SnmpCorrTrapBag struct {
	// Object ID
	Oid string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	// Number of hsecs elapsed since snmpd was started
	RelativeTimestamp uint32 `protobuf:"varint,2,opt,name=relative_timestamp,json=relativeTimestamp,proto3" json:"relative_timestamp,omitempty"`
	// Time when the trap was generated. It is expressed in number of milliseconds since 00:00:00 UTC, January 1, 1970
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// VarBinds on the trap
	VarBind              []*SnmpCorrVbind `protobuf:"bytes,4,rep,name=var_bind,json=varBind,proto3" json:"var_bind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SnmpCorrTrapBag) Reset()         { *m = SnmpCorrTrapBag{} }
func (m *SnmpCorrTrapBag) String() string { return proto.CompactTextString(m) }
func (*SnmpCorrTrapBag) ProtoMessage()    {}
func (*SnmpCorrTrapBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_snmp_corr_trap_buffer_bag_a3c3f78d3369ba7a, []int{3}
}
func (m *SnmpCorrTrapBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpCorrTrapBag.Unmarshal(m, b)
}
func (m *SnmpCorrTrapBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpCorrTrapBag.Marshal(b, m, deterministic)
}
func (dst *SnmpCorrTrapBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpCorrTrapBag.Merge(dst, src)
}
func (m *SnmpCorrTrapBag) XXX_Size() int {
	return xxx_messageInfo_SnmpCorrTrapBag.Size(m)
}
func (m *SnmpCorrTrapBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpCorrTrapBag.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpCorrTrapBag proto.InternalMessageInfo

func (m *SnmpCorrTrapBag) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *SnmpCorrTrapBag) GetRelativeTimestamp() uint32 {
	if m != nil {
		return m.RelativeTimestamp
	}
	return 0
}

func (m *SnmpCorrTrapBag) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SnmpCorrTrapBag) GetVarBind() []*SnmpCorrVbind {
	if m != nil {
		return m.VarBind
	}
	return nil
}

func init() {
	proto.RegisterType((*SnmpCorrTrapBufferBag_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.correlator.traps.trap.snmp_corr_trap_buffer_bag_KEYS")
	proto.RegisterType((*SnmpCorrTrapBufferBag)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.correlator.traps.trap.snmp_corr_trap_buffer_bag")
	proto.RegisterType((*SnmpCorrVbind)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.correlator.traps.trap.snmp_corr_vbind")
	proto.RegisterType((*SnmpCorrTrapBag)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.correlator.traps.trap.snmp_corr_trap_bag")
}

func init() {
	proto.RegisterFile("snmp_corr_trap_buffer_bag.proto", fileDescriptor_snmp_corr_trap_buffer_bag_a3c3f78d3369ba7a)
}

var fileDescriptor_snmp_corr_trap_buffer_bag_a3c3f78d3369ba7a = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x45, 0x9b, 0xec, 0xc6, 0x9e, 0x6c, 0xf6, 0x43, 0xec, 0xc1, 0x61, 0x97, 0x5d, 0xaf, 0xa1,
	0xe0, 0x4b, 0x7d, 0x48, 0x5a, 0x4a, 0xe9, 0xad, 0x50, 0x68, 0x08, 0xf4, 0xa0, 0xf6, 0xd2, 0x93,
	0x90, 0x6d, 0x39, 0x15, 0xc4, 0x92, 0x91, 0x64, 0xd3, 0xfe, 0xd3, 0xfe, 0x88, 0xfe, 0x88, 0x22,
	0xa5, 0x26, 0xa5, 0x6d, 0x2e, 0xa5, 0x17, 0x33, 0xf3, 0xde, 0xbc, 0x37, 0x9e, 0x87, 0xe0, 0x9f,
	0x91, 0x75, 0x43, 0x0b, 0xa5, 0x35, 0xb5, 0x9a, 0x35, 0x34, 0x6f, 0xab, 0x8a, 0x6b, 0x9a, 0xb3,
	0x55, 0xd6, 0x68, 0x65, 0x15, 0x3e, 0x2a, 0x84, 0x29, 0x14, 0x15, 0xca, 0xd0, 0x5b, 0x4d, 0xfd,
	0x34, 0x5b, 0x71, 0x69, 0xa9, 0x6a, 0xb8, 0xce, 0x5c, 0x9f, 0x39, 0x35, 0x5f, 0x33, 0xab, 0x74,
	0xe6, 0x3c, 0x8c, 0xff, 0x26, 0x27, 0xf0, 0x77, 0xa7, 0x37, 0x5d, 0x9e, 0x5d, 0x5f, 0xe2, 0x29,
	0x04, 0x5c, 0x5a, 0x7d, 0x47, 0x45, 0x19, 0xa1, 0x18, 0xa5, 0x13, 0x32, 0xf2, 0xfd, 0xa2, 0x4c,
	0x1e, 0x10, 0x4c, 0x77, 0xaa, 0xf1, 0x1e, 0x7c, 0xeb, 0x77, 0x0a, 0x25, 0x9d, 0x7c, 0xe6, 0xe5,
	0x93, 0x67, 0xe8, 0xa2, 0xc4, 0xff, 0xe1, 0xab, 0x30, 0x54, 0x2b, 0x65, 0x0b, 0xd6, 0x1a, 0x1e,
	0xcd, 0x63, 0x94, 0x06, 0x64, 0x2c, 0x0c, 0xe9, 0x21, 0xfc, 0x1b, 0x42, 0xdd, 0xae, 0x39, 0x95,
	0xac, 0xe6, 0xd1, 0x41, 0x8c, 0xd2, 0x90, 0x04, 0x0e, 0xb8, 0x60, 0x35, 0xc7, 0x37, 0x10, 0xfa,
	0xcd, 0x42, 0x56, 0x2a, 0x3a, 0x8c, 0x51, 0x3a, 0x9e, 0x2d, 0xb3, 0x77, 0xc6, 0x91, 0xbd, 0xbc,
	0x86, 0xad, 0x48, 0xe0, 0xaa, 0x85, 0xac, 0x54, 0x72, 0x0c, 0xdf, 0xb7, 0x7c, 0x97, 0x0b, 0x59,
	0xe2, 0x1f, 0x30, 0x50, 0x4f, 0xb9, 0x84, 0xc4, 0x95, 0xf8, 0x17, 0x7c, 0xee, 0xd8, 0xba, 0xe5,
	0xd1, 0x27, 0x8f, 0x6d, 0x9a, 0xe4, 0x1e, 0x01, 0x7e, 0xed, 0xfd, 0x86, 0x7c, 0x1f, 0xf0, 0x26,
	0x9b, 0x8e, 0x53, 0x2b, 0x6a, 0x6e, 0x2c, 0xab, 0x1b, 0xef, 0x35, 0x21, 0x3f, 0x7b, 0xe6, 0xaa,
	0x27, 0xf0, 0x1f, 0x08, 0xb7, 0x53, 0x83, 0x18, 0xa5, 0x43, 0xb2, 0x05, 0x70, 0x01, 0x41, 0xc7,
	0x34, 0x75, 0x7f, 0x1a, 0x0d, 0xe3, 0x41, 0x3a, 0x9e, 0x9d, 0x7f, 0x40, 0x32, 0xfe, 0x72, 0x32,
	0xea, 0x98, 0x3e, 0x15, 0xb2, 0xcc, 0xbf, 0xf8, 0x17, 0x38, 0x7f, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x0c, 0xe0, 0xa4, 0xe3, 0xa4, 0x02, 0x00, 0x00,
}
