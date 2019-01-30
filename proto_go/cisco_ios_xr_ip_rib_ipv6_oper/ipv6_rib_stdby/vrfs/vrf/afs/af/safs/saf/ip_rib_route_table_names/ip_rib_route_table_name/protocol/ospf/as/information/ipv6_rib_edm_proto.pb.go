// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_rib_edm_proto.proto

/*
Package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_ospf_as_information is a generated protocol buffer package.

It is generated from these files:
	ipv6_rib_edm_proto.proto

It has these top-level messages:
	Ipv6RibEdmProto_KEYS
	Ipv6RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_ospf_as_information

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
type Ipv6RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	As             string `protobuf:"bytes,5,opt,name=as" json:"as,omitempty"`
}

func (m *Ipv6RibEdmProto_KEYS) Reset()                    { *m = Ipv6RibEdmProto_KEYS{} }
func (m *Ipv6RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv6RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetAs() string {
	if m != nil {
		return m.As
	}
	return ""
}

type Ipv6RibEdmProto struct {
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

func (m *Ipv6RibEdmProto) Reset()                    { *m = Ipv6RibEdmProto{} }
func (m *Ipv6RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto) ProtoMessage()               {}
func (*Ipv6RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv6RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.ospf.as.information.ipv6_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv6RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.ospf.as.information.ipv6_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv6_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x14, 0x9a, 0x30, 0x90, 0x08, 0xb6, 0x85, 0x6e, 0xe0, 0x40, 0x55, 0x84, 0x94,
	0xd3, 0x0a, 0xb5, 0xa5, 0xfc, 0x15, 0x97, 0x8a, 0x13, 0x82, 0x43, 0xe0, 0xc2, 0x69, 0x35, 0xb6,
	0xd7, 0x62, 0x45, 0xbc, 0x6b, 0xed, 0x6c, 0x2c, 0xfa, 0x1c, 0xbc, 0x01, 0xaf, 0xc8, 0x0b, 0x20,
	0xcf, 0xda, 0x21, 0xa1, 0xe2, 0xe2, 0xc8, 0xf3, 0xfd, 0x7e, 0x93, 0x6f, 0xd7, 0x20, 0x6d, 0xd3,
	0x5e, 0xe8, 0x60, 0x73, 0x6d, 0xca, 0x5a, 0x37, 0xc1, 0x47, 0xaf, 0xf8, 0x29, 0x7e, 0x66, 0x85,
	0xa5, 0xc2, 0x6b, 0xeb, 0x49, 0xff, 0x08, 0xda, 0x36, 0x4c, 0x31, 0xee, 0x1b, 0x13, 0xd4, 0x46,
	0xa4, 0x58, 0xe6, 0x57, 0xaa, 0x0d, 0x15, 0x75, 0x0f, 0x85, 0x15, 0x29, 0xac, 0x14, 0x75, 0xbf,
	0x84, 0x95, 0xea, 0xc5, 0xe0, 0xd7, 0xd1, 0xe8, 0x88, 0xf9, 0xca, 0x68, 0x87, 0xb5, 0xa1, 0xff,
	0x05, 0xe9, 0xef, 0x0b, 0xbf, 0x52, 0x9e, 0x9a, 0x4a, 0x21, 0x29, 0xeb, 0x2a, 0x1f, 0x6a, 0x8c,
	0xd6, 0xbb, 0x93, 0x5f, 0x19, 0x1c, 0x5d, 0xaf, 0xac, 0x3f, 0xbc, 0xff, 0xfa, 0x59, 0xcc, 0x61,
	0xd2, 0x86, 0x8a, 0x97, 0xc8, 0xec, 0x38, 0x5b, 0xdc, 0x5a, 0x8e, 0xdb, 0x50, 0x7d, 0xc2, 0xda,
	0x88, 0x23, 0x18, 0x63, 0x9f, 0x8c, 0x38, 0xd9, 0xc7, 0x14, 0xcc, 0x61, 0x42, 0x43, 0xb2, 0x97,
	0x1c, 0xea, 0xa3, 0x05, 0xdc, 0xfd, 0xb7, 0x9b, 0xbc, 0xc1, 0xc8, 0x8c, 0xe7, 0x5f, 0xba, 0x31,
	0x93, 0x33, 0x18, 0x21, 0xc9, 0x9b, 0x9c, 0x8d, 0x90, 0x4e, 0x7e, 0xef, 0x81, 0xb8, 0x5e, 0x52,
	0x3c, 0x85, 0xd9, 0x70, 0xb6, 0x74, 0x05, 0xf2, 0x94, 0x95, 0xe9, 0x30, 0xed, 0x96, 0x91, 0x78,
	0x08, 0x13, 0xeb, 0x28, 0xa2, 0x2b, 0x8c, 0x3c, 0x63, 0x60, 0xf3, 0x2e, 0x24, 0x8c, 0x5b, 0x13,
	0xc8, 0x7a, 0x27, 0xcf, 0x8f, 0xb3, 0xc5, 0x74, 0x39, 0xbc, 0x8a, 0x77, 0xf0, 0x28, 0x98, 0xd2,
	0x52, 0x0c, 0x36, 0x5f, 0x77, 0x57, 0xa5, 0x8b, 0x95, 0x35, 0x2e, 0xea, 0xc2, 0xaf, 0x5d, 0x94,
	0xcf, 0x99, 0x9e, 0xef, 0x22, 0x97, 0x4c, 0x5c, 0x76, 0x80, 0x38, 0x87, 0x07, 0x9b, 0x72, 0xc9,
	0xa4, 0x5e, 0xbd, 0x60, 0xf5, 0x70, 0x48, 0x93, 0x44, 0xc9, 0x7a, 0x02, 0x53, 0xbe, 0x8b, 0x9e,
	0x25, 0xf9, 0x82, 0xe1, 0x3b, 0x69, 0xc8, 0x0c, 0x09, 0x05, 0x07, 0x58, 0x44, 0xdb, 0x1a, 0xbd,
	0xcd, 0xca, 0x97, 0x8c, 0xde, 0x4b, 0xd1, 0xf2, 0xaf, 0x20, 0x9e, 0xc1, 0x61, 0x69, 0x56, 0x26,
	0x9a, 0x72, 0x57, 0x78, 0xc5, 0x82, 0xe8, 0xb3, 0x6d, 0xe3, 0x31, 0xdc, 0x6e, 0x30, 0x7e, 0x1b,
	0xc0, 0xd7, 0x0c, 0x02, 0x8f, 0x12, 0x70, 0x0a, 0xf7, 0x37, 0xa7, 0x4b, 0x1f, 0xb5, 0x36, 0xb5,
	0x0f, 0x57, 0xf2, 0x0d, 0xa3, 0x07, 0x43, 0xc8, 0x4b, 0x3f, 0x72, 0xd4, 0xd5, 0xce, 0xb1, 0xf8,
	0xbe, 0x6e, 0x76, 0x5b, 0xbc, 0x4d, 0xb5, 0x53, 0xb4, 0x55, 0x22, 0xdf, 0xe7, 0x25, 0x67, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x6f, 0x13, 0xb8, 0x53, 0x03, 0x00, 0x00,
}
