// Code generated by protoc-gen-go.
// source: bgp_upderr_vrf_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_update_inbound_error_vrf is a generated protocol buffer package.

It is generated from these files:
	bgp_upderr_vrf_bag.proto

It has these top-level messages:
	BgpUpderrVrfBag_KEYS
	BgpUpderrVrfBag
	BgpTimespec
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_update_inbound_error_vrf

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

// BGP Update error-handling VRF information
type BgpUpderrVrfBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *BgpUpderrVrfBag_KEYS) Reset()                    { *m = BgpUpderrVrfBag_KEYS{} }
func (m *BgpUpderrVrfBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpUpderrVrfBag_KEYS) ProtoMessage()               {}
func (*BgpUpderrVrfBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpUpderrVrfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpUpderrVrfBag struct {
	// VRF Name
	UpdateVrfName string `protobuf:"bytes,50,opt,name=update_vrf_name,json=updateVrfName" json:"update_vrf_name,omitempty"`
	// Malformed messages count
	UpdateMalformedMessageCount uint32 `protobuf:"varint,51,opt,name=update_malformed_message_count,json=updateMalformedMessageCount" json:"update_malformed_message_count,omitempty"`
	// Count of neighbors that received malformed messages
	UpdateMalformedNeighborCount uint32 `protobuf:"varint,52,opt,name=update_malformed_neighbor_count,json=updateMalformedNeighborCount" json:"update_malformed_neighbor_count,omitempty"`
	// Last malformed messages received time: time elapsed since 00:00:00 UTC, January 1, 1970
	LastUpdateMalformedTimestamp *BgpTimespec `protobuf:"bytes,53,opt,name=last_update_malformed_timestamp,json=lastUpdateMalformedTimestamp" json:"last_update_malformed_timestamp,omitempty"`
	// Time since last malformed messages received event (in seconds)
	LastUpdateMalformedAge uint32 `protobuf:"varint,54,opt,name=last_update_malformed_age,json=lastUpdateMalformedAge" json:"last_update_malformed_age,omitempty"`
}

func (m *BgpUpderrVrfBag) Reset()                    { *m = BgpUpderrVrfBag{} }
func (m *BgpUpderrVrfBag) String() string            { return proto.CompactTextString(m) }
func (*BgpUpderrVrfBag) ProtoMessage()               {}
func (*BgpUpderrVrfBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpUpderrVrfBag) GetUpdateVrfName() string {
	if m != nil {
		return m.UpdateVrfName
	}
	return ""
}

func (m *BgpUpderrVrfBag) GetUpdateMalformedMessageCount() uint32 {
	if m != nil {
		return m.UpdateMalformedMessageCount
	}
	return 0
}

func (m *BgpUpderrVrfBag) GetUpdateMalformedNeighborCount() uint32 {
	if m != nil {
		return m.UpdateMalformedNeighborCount
	}
	return 0
}

func (m *BgpUpderrVrfBag) GetLastUpdateMalformedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateMalformedTimestamp
	}
	return nil
}

func (m *BgpUpderrVrfBag) GetLastUpdateMalformedAge() uint32 {
	if m != nil {
		return m.LastUpdateMalformedAge
	}
	return 0
}

type BgpTimespec struct {
	// Seconds part of time value
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Nanoseconds part of time value
	Nanoseconds uint32 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *BgpTimespec) Reset()                    { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string            { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()               {}
func (*BgpTimespec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpUpderrVrfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.update_inbound_error_vrf.bgp_upderr_vrf_bag_KEYS")
	proto.RegisterType((*BgpUpderrVrfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.update_inbound_error_vrf.bgp_upderr_vrf_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.update_inbound_error_vrf.bgp_timespec")
}

func init() { proto.RegisterFile("bgp_upderr_vrf_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xe9, 0xfb, 0xc2, 0xfb, 0x62, 0xb6, 0x22, 0xe4, 0xa0, 0x15, 0x87, 0x2b, 0x13, 0x64,
	0xa7, 0x1e, 0xb6, 0x29, 0x78, 0x11, 0x64, 0xec, 0xa2, 0x6c, 0x87, 0xfa, 0x03, 0x3c, 0x85, 0xb4,
	0xfd, 0x36, 0x16, 0xd6, 0x24, 0x24, 0xe9, 0xd0, 0xff, 0xcd, 0x93, 0x7f, 0x99, 0x24, 0x6d, 0xc7,
	0xb4, 0x3b, 0x7b, 0x6a, 0xfa, 0xe4, 0xf9, 0x7e, 0x9e, 0x27, 0x24, 0x28, 0x48, 0x98, 0x24, 0x95,
	0xcc, 0x40, 0x29, 0xb2, 0x51, 0x39, 0x49, 0x28, 0x8b, 0xa4, 0x12, 0x46, 0x60, 0x96, 0x16, 0x3a,
	0x15, 0xa4, 0x10, 0x9a, 0xbc, 0x29, 0x52, 0xc8, 0xcd, 0x8c, 0x58, 0xaf, 0x90, 0xa0, 0xa2, 0x84,
	0xc9, 0xa8, 0xe0, 0xda, 0x50, 0x9e, 0x82, 0xde, 0xae, 0xb6, 0x0b, 0x62, 0x3f, 0x59, 0xf2, 0x1e,
	0x65, 0x90, 0xd3, 0x6a, 0x6d, 0x2c, 0x39, 0xaa, 0x64, 0x46, 0x0d, 0x90, 0x82, 0x27, 0xa2, 0xe2,
	0x19, 0x01, 0xa5, 0x84, 0x8b, 0x1c, 0xdd, 0xa0, 0xe3, 0x6e, 0x09, 0x72, 0xbf, 0x78, 0x79, 0xc0,
	0xe7, 0xc8, 0xdf, 0x32, 0x39, 0x2d, 0x21, 0xf0, 0x42, 0x6f, 0x7c, 0x10, 0xf7, 0x5b, 0x71, 0x45,
	0x4b, 0x18, 0x7d, 0xfe, 0x45, 0xb8, 0x0b, 0xc0, 0x17, 0xe8, 0xb0, 0x89, 0xb4, 0x8a, 0x9b, 0x9e,
	0xb8, 0x69, 0xbf, 0x96, 0x9f, 0x55, 0x6e, 0xc7, 0xf1, 0x1c, 0x9d, 0x35, 0xbe, 0x92, 0xae, 0x73,
	0xa1, 0x4a, 0xc8, 0x48, 0x09, 0x5a, 0x53, 0x06, 0x24, 0x15, 0x15, 0x37, 0xc1, 0x34, 0xf4, 0xc6,
	0x7e, 0x7c, 0x5a, 0xbb, 0x96, 0xad, 0x69, 0x59, 0x7b, 0xe6, 0xd6, 0x82, 0x17, 0x68, 0xd8, 0x81,
	0x70, 0x28, 0xd8, 0x6b, 0x22, 0x54, 0x43, 0x99, 0x39, 0xca, 0xe0, 0x07, 0x65, 0xd5, 0x98, 0x6a,
	0xcc, 0x87, 0x87, 0x86, 0x6b, 0xaa, 0x0d, 0xe9, 0xc0, 0x4c, 0x51, 0x82, 0x36, 0xb4, 0x94, 0xc1,
	0x65, 0xe8, 0x8d, 0x7b, 0x93, 0x2a, 0xfa, 0xa5, 0xeb, 0xb1, 0xb0, 0x3a, 0x5d, 0x42, 0x1a, 0x0f,
	0x6c, 0xbb, 0xa7, 0xef, 0x47, 0x78, 0x6c, 0xab, 0xe1, 0x6b, 0x74, 0xb2, 0xbf, 0x3d, 0x65, 0x10,
	0x5c, 0xb9, 0xf3, 0x1f, 0xed, 0x01, 0xdc, 0x32, 0x18, 0xdd, 0xa1, 0xfe, 0x6e, 0x10, 0x0e, 0xd0,
	0x7f, 0x0d, 0xa9, 0xe0, 0x99, 0x76, 0x77, 0xee, 0xc7, 0xed, 0x2f, 0x0e, 0x51, 0x8f, 0x53, 0x2e,
	0xda, 0xdd, 0x3f, 0x6e, 0x77, 0x57, 0x4a, 0xfe, 0xb9, 0x07, 0x3c, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0x3c, 0xb7, 0x0b, 0xdc, 0x02, 0x00, 0x00,
}