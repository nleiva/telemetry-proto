// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_sysname.proto

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_system_name is a generated protocol buffer package.

It is generated from these files:
	snmp_sysname.proto

It has these top-level messages:
	SnmpSysname_KEYS
	SnmpSysname
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_system_name

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

// SNMP sysName
type SnmpSysname_KEYS struct {
}

func (m *SnmpSysname_KEYS) Reset()                    { *m = SnmpSysname_KEYS{} }
func (m *SnmpSysname_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpSysname_KEYS) ProtoMessage()               {}
func (*SnmpSysname_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SnmpSysname struct {
	// sysName  1.3.6.1.2.1.1.5
	SystemName string `protobuf:"bytes,50,opt,name=system_name,json=systemName" json:"system_name,omitempty"`
}

func (m *SnmpSysname) Reset()                    { *m = SnmpSysname{} }
func (m *SnmpSysname) String() string            { return proto.CompactTextString(m) }
func (*SnmpSysname) ProtoMessage()               {}
func (*SnmpSysname) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpSysname) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

func init() {
	proto.RegisterType((*SnmpSysname_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.system_name.snmp_sysname_KEYS")
	proto.RegisterType((*SnmpSysname)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.system_name.snmp_sysname")
}

func init() { proto.RegisterFile("snmp_sysname.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xce, 0xcb, 0x2d,
	0x88, 0x2f, 0xae, 0x2c, 0xce, 0x4b, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2,
	0x4c, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x07, 0x2b, 0x48,
	0x4c, 0x4f, 0xcd, 0x2b, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xf1, 0xf5, 0x32, 0xf3, 0xd2,
	0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x8a, 0x2b, 0x8b, 0x4b, 0x52, 0x73, 0xe3,
	0x41, 0x06, 0x28, 0x09, 0x73, 0x09, 0x22, 0x1b, 0x18, 0xef, 0xed, 0x1a, 0x19, 0xac, 0xa4, 0xcf,
	0xc5, 0x83, 0x2c, 0x28, 0x24, 0xcf, 0xc5, 0x8d, 0xa4, 0x47, 0xc2, 0x48, 0x81, 0x51, 0x83, 0x33,
	0x88, 0x0b, 0x22, 0xe4, 0x97, 0x98, 0x9b, 0x9a, 0xc4, 0x06, 0x76, 0x87, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x29, 0xe7, 0xd4, 0xd1, 0x9d, 0x00, 0x00, 0x00,
}
