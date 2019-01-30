// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_engineid.proto

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_engine_id is a generated protocol buffer package.

It is generated from these files:
	snmp_engineid.proto

It has these top-level messages:
	SnmpEngineid_KEYS
	SnmpEngineid
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_engine_id

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

// SNMPv3 engineID Information
type SnmpEngineid_KEYS struct {
}

func (m *SnmpEngineid_KEYS) Reset()                    { *m = SnmpEngineid_KEYS{} }
func (m *SnmpEngineid_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpEngineid_KEYS) ProtoMessage()               {}
func (*SnmpEngineid_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SnmpEngineid struct {
	// SNMPv3 engineID
	EngineId string `protobuf:"bytes,50,opt,name=engine_id,json=engineId" json:"engine_id,omitempty"`
}

func (m *SnmpEngineid) Reset()                    { *m = SnmpEngineid{} }
func (m *SnmpEngineid) String() string            { return proto.CompactTextString(m) }
func (*SnmpEngineid) ProtoMessage()               {}
func (*SnmpEngineid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpEngineid) GetEngineId() string {
	if m != nil {
		return m.EngineId
	}
	return ""
}

func init() {
	proto.RegisterType((*SnmpEngineid_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.engine_id.snmp_engineid_KEYS")
	proto.RegisterType((*SnmpEngineid)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.engine_id.snmp_engineid")
}

func init() { proto.RegisterFile("snmp_engineid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0xcb, 0x2d,
	0x88, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xcd, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x32, 0x4f, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x07, 0xab,
	0x48, 0x4c, 0x4f, 0xcd, 0x2b, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xf1, 0xf5, 0x32, 0xf3,
	0xd2, 0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x20, 0xba, 0xe3, 0x33, 0x53, 0x94,
	0x44, 0xb8, 0x84, 0x50, 0xcc, 0x8b, 0xf7, 0x76, 0x8d, 0x0c, 0x56, 0xd2, 0xe1, 0xe2, 0x45, 0x11,
	0x15, 0x92, 0xe6, 0xe2, 0x84, 0xeb, 0x91, 0x30, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x80, 0x08,
	0x78, 0xa6, 0x24, 0xb1, 0x81, 0xdd, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xf4, 0x2b,
	0xa5, 0x9a, 0x00, 0x00, 0x00,
}
