// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_proto.proto

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_eigrp_as_information is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_proto.proto

It has these top-level messages:
	Ipv4RibEdmProto_KEYS
	Ipv4RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_eigrp_as_information

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

// Information of a rib protocol
type Ipv4RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	As             string `protobuf:"bytes,5,opt,name=as" json:"as,omitempty"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()                    { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmProto_KEYS) GetAs() string {
	if m != nil {
		return m.As
	}
	return ""
}

type Ipv4RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount" json:"backup_routes_count,omitempty"`
}

func (m *Ipv4RibEdmProto) Reset()                    { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()               {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.eigrp.as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.eigrp.as.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0x2e, 0x6c, 0xcb, 0x40, 0x2b, 0xf0, 0x2e, 0xac, 0x0b, 0x07, 0x56, 0x8b, 0x90,
	0x7a, 0xb2, 0xd0, 0x6e, 0xf9, 0x8f, 0xb8, 0xac, 0x38, 0x21, 0x38, 0x14, 0x2e, 0x9c, 0x2c, 0x27,
	0x99, 0x80, 0x45, 0x13, 0x47, 0x1e, 0x27, 0x62, 0x9f, 0x82, 0x77, 0xe0, 0x15, 0x79, 0x01, 0x94,
	0x71, 0x52, 0x5a, 0x2a, 0x2e, 0xa9, 0x3c, 0xbf, 0xdf, 0x37, 0xfd, 0xe2, 0x80, 0xb4, 0x75, 0xbb,
	0xd4, 0xde, 0xa6, 0x1a, 0xf3, 0x52, 0xd7, 0xde, 0x05, 0xa7, 0xf8, 0x29, 0x7e, 0x26, 0x99, 0xa5,
	0xcc, 0x69, 0xeb, 0x48, 0xff, 0xf0, 0xda, 0xd6, 0x6c, 0xb1, 0xee, 0x6a, 0xf4, 0xaa, 0x3b, 0x51,
	0xc8, 0xd3, 0x2b, 0xd5, 0xfa, 0x82, 0xba, 0x87, 0x32, 0x05, 0x29, 0x53, 0x28, 0xea, 0x7e, 0xc9,
	0x14, 0xaa, 0xcf, 0x78, 0xd7, 0x04, 0xd4, 0xc1, 0xa4, 0x6b, 0xd4, 0x95, 0x29, 0x91, 0xfe, 0x07,
	0xe2, 0x3f, 0x67, 0x6e, 0xad, 0xd0, 0x7e, 0xf5, 0xb5, 0x32, 0xa4, 0x6c, 0x55, 0x38, 0x5f, 0x9a,
	0x60, 0x5d, 0x75, 0xf6, 0x2b, 0x81, 0x93, 0xfd, 0xba, 0xfa, 0xfd, 0xbb, 0x2f, 0x9f, 0xc4, 0x1c,
	0x26, 0xad, 0x2f, 0x78, 0x8b, 0x4c, 0x4e, 0x93, 0xc5, 0x8d, 0xd5, 0xb8, 0xf5, 0xc5, 0x47, 0x53,
	0xa2, 0x38, 0x81, 0xb1, 0xe9, 0xc9, 0x88, 0xc9, 0xa1, 0x89, 0x60, 0x0e, 0x13, 0x1a, 0xc8, 0x41,
	0xcc, 0x50, 0x8f, 0x16, 0x70, 0xfb, 0xdf, 0x72, 0xf2, 0x1a, 0x2b, 0x33, 0x9e, 0x7f, 0xee, 0xc6,
	0x6c, 0xce, 0x60, 0x64, 0x48, 0x5e, 0x67, 0x36, 0x32, 0x74, 0xf6, 0xfb, 0x00, 0xc4, 0x7e, 0x49,
	0xf1, 0x18, 0x66, 0xc3, 0xcb, 0xc5, 0x3b, 0x90, 0xe7, 0x1c, 0x99, 0x0e, 0xd3, 0x6e, 0x19, 0x89,
	0xfb, 0x30, 0xb1, 0x15, 0x05, 0x53, 0x65, 0x28, 0x2f, 0x58, 0xd8, 0x9c, 0x85, 0x84, 0x71, 0x8b,
	0x9e, 0xac, 0xab, 0xe4, 0xf2, 0x34, 0x59, 0x4c, 0x57, 0xc3, 0x51, 0xbc, 0x85, 0x07, 0x1e, 0x73,
	0x4b, 0xc1, 0xdb, 0xb4, 0xe9, 0xae, 0x4a, 0x67, 0x6b, 0x8b, 0x55, 0xd0, 0x99, 0x6b, 0xaa, 0x20,
	0x9f, 0xb2, 0x3d, 0xdf, 0x55, 0x2e, 0xd9, 0xb8, 0xec, 0x04, 0xb1, 0x84, 0x7b, 0x9b, 0x72, 0x31,
	0x49, 0x7d, 0xf4, 0x19, 0x47, 0x8f, 0x07, 0x1a, 0x43, 0x14, 0x53, 0x8f, 0x60, 0xca, 0x77, 0xd1,
	0xbb, 0x24, 0x9f, 0xb3, 0x7c, 0x2b, 0x0e, 0xd9, 0x21, 0xa1, 0xe0, 0xc8, 0x64, 0xc1, 0xb6, 0xa8,
	0xb7, 0x5d, 0xf9, 0x82, 0xd5, 0x3b, 0x11, 0xad, 0xfe, 0x06, 0xc4, 0x13, 0x38, 0xce, 0x71, 0x8d,
	0x01, 0xf3, 0xdd, 0xc0, 0x4b, 0x0e, 0x88, 0x9e, 0x6d, 0x27, 0x1e, 0xc2, 0xcd, 0xda, 0x84, 0x6f,
	0x83, 0xf8, 0x8a, 0x45, 0xe0, 0x51, 0x14, 0xce, 0xe1, 0xee, 0xe6, 0xed, 0xe2, 0x47, 0x2d, 0xb1,
	0x74, 0xfe, 0x4a, 0xbe, 0x66, 0xf5, 0x68, 0x80, 0xbc, 0xf4, 0x03, 0xa3, 0xae, 0x76, 0x6a, 0xb2,
	0xef, 0x4d, 0xbd, 0xdb, 0xe2, 0x4d, 0xac, 0x1d, 0xd1, 0x56, 0x89, 0xf4, 0x90, 0x97, 0x5c, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x2b, 0xfb, 0x0e, 0x4f, 0x03, 0x00, 0x00,
}
