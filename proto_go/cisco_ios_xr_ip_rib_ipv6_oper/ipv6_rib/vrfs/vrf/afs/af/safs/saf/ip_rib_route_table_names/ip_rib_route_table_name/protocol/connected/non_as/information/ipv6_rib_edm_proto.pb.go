// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_connected_non_as_information

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

// Information of a rib protocol
type Ipv6RibEdmProto_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmProto_KEYS) Reset()         { *m = Ipv6RibEdmProto_KEYS{} }
func (m *Ipv6RibEdmProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto_KEYS) ProtoMessage()    {}
func (*Ipv6RibEdmProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb6e91b99f67ad, []int{0}
}

func (m *Ipv6RibEdmProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmProto_KEYS.Unmarshal(m, b)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmProto_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmProto_KEYS.Merge(m, src)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmProto_KEYS.Size(m)
}
func (m *Ipv6RibEdmProto_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmProto_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmProto_KEYS proto.InternalMessageInfo

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

type Ipv6RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames,proto3" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance,proto3" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version,proto3" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount,proto3" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount,proto3" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts,proto3" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount,proto3" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount,proto3" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount,proto3" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory,proto3" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount    uint32   `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount,proto3" json:"backup_routes_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6RibEdmProto) Reset()         { *m = Ipv6RibEdmProto{} }
func (m *Ipv6RibEdmProto) String() string { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto) ProtoMessage()    {}
func (*Ipv6RibEdmProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cb6e91b99f67ad, []int{1}
}

func (m *Ipv6RibEdmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6RibEdmProto.Unmarshal(m, b)
}
func (m *Ipv6RibEdmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6RibEdmProto.Marshal(b, m, deterministic)
}
func (m *Ipv6RibEdmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6RibEdmProto.Merge(m, src)
}
func (m *Ipv6RibEdmProto) XXX_Size() int {
	return xxx_messageInfo_Ipv6RibEdmProto.Size(m)
}
func (m *Ipv6RibEdmProto) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6RibEdmProto.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6RibEdmProto proto.InternalMessageInfo

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
	proto.RegisterType((*Ipv6RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.non_as.information.ipv6_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv6RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.non_as.information.ipv6_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv6_rib_edm_proto.proto", fileDescriptor_63cb6e91b99f67ad) }

var fileDescriptor_63cb6e91b99f67ad = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x6f, 0x13, 0x3d,
	0x10, 0xc6, 0xb5, 0x6f, 0x5f, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42,
	0xca, 0xc9, 0x42, 0x6d, 0x29, 0x7f, 0xc5, 0xa5, 0xe2, 0x84, 0xe0, 0x10, 0xb8, 0x70, 0xb2, 0xbc,
	0xde, 0x59, 0x61, 0x91, 0xb5, 0x57, 0x1e, 0x67, 0x05, 0x5f, 0x83, 0x03, 0x5f, 0x94, 0x2f, 0x80,
	0x76, 0xbc, 0x1b, 0x12, 0x2a, 0x2e, 0x8e, 0x3c, 0xcf, 0xef, 0xe7, 0x3c, 0x93, 0x80, 0x74, 0x6d,
	0x77, 0xa9, 0xa3, 0x2b, 0x35, 0x56, 0x8d, 0x6e, 0x63, 0x48, 0x41, 0xf1, 0x29, 0x7e, 0x16, 0xd6,
	0x91, 0x0d, 0xda, 0x05, 0xd2, 0xdf, 0xa2, 0x76, 0x2d, 0x53, 0x8c, 0x87, 0x16, 0xa3, 0x1a, 0x45,
	0xd5, 0xc5, 0x9a, 0xfa, 0x43, 0x99, 0x9a, 0x94, 0xa9, 0x15, 0xf5, 0x9f, 0x64, 0x6a, 0x35, 0x28,
	0x31, 0xac, 0x13, 0xea, 0x64, 0xca, 0x15, 0x6a, 0x6f, 0x1a, 0xa4, 0x7f, 0x05, 0xf9, 0x8b, 0x6d,
	0x58, 0x29, 0x1b, 0xbc, 0x47, 0x9b, 0xb0, 0x52, 0x3e, 0x78, 0x6d, 0x48, 0x39, 0x5f, 0x87, 0xd8,
	0x98, 0xe4, 0x82, 0x3f, 0xfd, 0x51, 0xc0, 0xf1, 0xf5, 0xd6, 0xfa, 0xdd, 0xdb, 0xcf, 0x1f, 0xc5,
	0x1c, 0xa6, 0x5d, 0xac, 0xf9, 0x35, 0x59, 0x9c, 0x14, 0x8b, 0x1b, 0xcb, 0x49, 0x17, 0xeb, 0x0f,
	0xa6, 0x41, 0x71, 0x0c, 0x13, 0x33, 0x24, 0xff, 0x71, 0xb2, 0x6f, 0x72, 0x30, 0x87, 0x29, 0x8d,
	0xc9, 0x5e, 0x76, 0x68, 0x88, 0x16, 0x70, 0xfb, 0xef, 0x92, 0xf2, 0x7f, 0x46, 0x0e, 0x78, 0xfe,
	0xa9, 0x1f, 0xf7, 0xe4, 0xe9, 0xaf, 0x3d, 0x10, 0xd7, 0x4b, 0x89, 0xc7, 0x70, 0x30, 0x2e, 0x95,
	0x77, 0x97, 0x67, 0xac, 0xcf, 0xc6, 0x69, 0x2f, 0x93, 0xb8, 0x0f, 0x53, 0xe7, 0x29, 0x19, 0x6f,
	0x51, 0x9e, 0x33, 0xb0, 0xb9, 0x0b, 0x09, 0x93, 0x0e, 0x23, 0xb9, 0xe0, 0xe5, 0xc5, 0x49, 0xb1,
	0x98, 0x2d, 0xc7, 0xab, 0x78, 0x03, 0x0f, 0x22, 0x56, 0x8e, 0x52, 0x74, 0xe5, 0xba, 0xff, 0x69,
	0xb4, 0x5d, 0x39, 0xf4, 0x49, 0xdb, 0xb0, 0xf6, 0x49, 0x3e, 0x65, 0x7a, 0xbe, 0x8b, 0x5c, 0x31,
	0x71, 0xd5, 0x03, 0xe2, 0x02, 0xee, 0x6d, 0xca, 0x65, 0x93, 0x06, 0xf5, 0x92, 0xd5, 0xa3, 0x31,
	0xcd, 0x12, 0x65, 0xeb, 0x11, 0xcc, 0x78, 0xf7, 0x81, 0x25, 0xf9, 0x8c, 0xe1, 0x5b, 0x79, 0xc8,
	0x0c, 0x09, 0x05, 0x87, 0xc6, 0x26, 0xd7, 0xa1, 0xde, 0x66, 0xe5, 0x73, 0x46, 0xef, 0xe4, 0x68,
	0xf9, 0x47, 0x10, 0x4f, 0xe0, 0xa8, 0xc2, 0x15, 0x26, 0xac, 0x76, 0x85, 0x17, 0x2c, 0x88, 0x21,
	0xdb, 0x36, 0x1e, 0xc2, 0xcd, 0xd6, 0xa4, 0x2f, 0x23, 0xf8, 0x92, 0x41, 0xe0, 0x51, 0x06, 0xce,
	0xe0, 0xee, 0x66, 0xbb, 0xfc, 0x27, 0x36, 0xd8, 0x84, 0xf8, 0x5d, 0xbe, 0x62, 0xf4, 0x70, 0x0c,
	0xf9, 0xd1, 0xf7, 0x1c, 0xf5, 0xb5, 0x4b, 0x63, 0xbf, 0xae, 0xdb, 0xdd, 0x16, 0xaf, 0x73, 0xed,
	0x1c, 0x6d, 0x95, 0x28, 0xf7, 0xf9, 0x91, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xa4,
	0xaf, 0xed, 0x46, 0x03, 0x00, 0x00,
}
