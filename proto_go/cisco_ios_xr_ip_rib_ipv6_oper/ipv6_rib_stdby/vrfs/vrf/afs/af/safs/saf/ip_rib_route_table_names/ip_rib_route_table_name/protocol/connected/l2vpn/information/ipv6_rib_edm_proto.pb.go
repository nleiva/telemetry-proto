// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_connected_l2vpn_information

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
	proto.RegisterType((*Ipv6RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.l2vpn.information.ipv6_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv6RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.connected.l2vpn.information.ipv6_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv6_rib_edm_proto.proto", fileDescriptor_63cb6e91b99f67ad) }

var fileDescriptor_63cb6e91b99f67ad = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x69, 0x29, 0x7f, 0xc5, 0xa5, 0xe2, 0x84, 0xe0, 0x10, 0xb8, 0x70, 0xb2, 0xbc, 0x5e,
	0xaf, 0xb0, 0xd8, 0xb5, 0x57, 0x1e, 0x67, 0x45, 0x5f, 0x83, 0x2b, 0x8f, 0xc9, 0x0b, 0x20, 0x8f,
	0x77, 0x43, 0x42, 0xc5, 0xc5, 0x91, 0xe7, 0xfb, 0xfd, 0x9c, 0x6f, 0x12, 0xe0, 0xb6, 0xeb, 0x2f,
	0x64, 0xb0, 0xa5, 0x34, 0x55, 0x2b, 0xbb, 0xe0, 0xa3, 0x17, 0x74, 0xb2, 0x5f, 0x85, 0xb6, 0xa8,
	0xbd, 0xb4, 0x1e, 0xe5, 0x8f, 0x20, 0x6d, 0x47, 0x14, 0xe1, 0xbe, 0x33, 0x41, 0x6c, 0x44, 0x8c,
	0x55, 0x79, 0x25, 0xfa, 0x50, 0x63, 0x3a, 0x84, 0xaa, 0x51, 0xa8, 0x5a, 0x60, 0xfa, 0x44, 0x55,
	0x8b, 0x41, 0x0c, 0x7e, 0x1d, 0x8d, 0x8c, 0xaa, 0x6c, 0x8c, 0x74, 0xaa, 0x35, 0xf8, 0xbf, 0x20,
	0x7f, 0xbd, 0xf6, 0x8d, 0xd0, 0xde, 0x39, 0xa3, 0xa3, 0xa9, 0x44, 0xb3, 0xec, 0x3b, 0x27, 0xac,
	0xab, 0x7d, 0x68, 0x55, 0xb4, 0xde, 0x9d, 0xfe, 0x2c, 0xe0, 0xf8, 0x7a, 0x75, 0xf9, 0xe1, 0xfd,
	0xd7, 0xcf, 0x6c, 0x0e, 0xd3, 0x3e, 0xd4, 0xf4, 0x18, 0x2f, 0x4e, 0x8a, 0xc5, 0xad, 0xd5, 0xa4,
	0x0f, 0xf5, 0x27, 0xd5, 0x1a, 0x76, 0x0c, 0x13, 0x35, 0x24, 0x37, 0x28, 0xd9, 0x57, 0x39, 0x98,
	0xc3, 0x14, 0xc7, 0x64, 0x2f, 0x3b, 0x38, 0x44, 0x0b, 0xb8, 0xfb, 0x6f, 0x47, 0x7e, 0x93, 0x90,
	0x03, 0x9a, 0x7f, 0x49, 0xe3, 0x44, 0x9e, 0xfe, 0xde, 0x03, 0x76, 0xbd, 0x14, 0x7b, 0x0a, 0x07,
	0xe3, 0x4e, 0x79, 0x75, 0xbe, 0x24, 0x7d, 0x36, 0x4e, 0x93, 0x8c, 0xec, 0x21, 0x4c, 0xad, 0xc3,
	0xa8, 0x9c, 0x36, 0xfc, 0x8c, 0x80, 0xcd, 0x9d, 0x71, 0x98, 0xf4, 0x26, 0xa0, 0xf5, 0x8e, 0x9f,
	0x9f, 0x14, 0x8b, 0xd9, 0x6a, 0xbc, 0xb2, 0x77, 0xf0, 0x28, 0x98, 0xca, 0x62, 0x0c, 0xb6, 0x5c,
	0xa7, 0x9f, 0x46, 0xea, 0xc6, 0x1a, 0x17, 0xa5, 0xf6, 0x6b, 0x17, 0xf9, 0x73, 0xa2, 0xe7, 0xbb,
	0xc8, 0x25, 0x11, 0x97, 0x09, 0x60, 0xe7, 0xf0, 0x60, 0x53, 0x2e, 0x9b, 0x38, 0xa8, 0x17, 0xa4,
	0x1e, 0x8d, 0x69, 0x96, 0x30, 0x5b, 0x4f, 0x60, 0x46, 0xbb, 0x0f, 0x2c, 0xf2, 0x17, 0x04, 0xdf,
	0xc9, 0x43, 0x62, 0x90, 0x09, 0x38, 0x54, 0x3a, 0xda, 0xde, 0xc8, 0x6d, 0x96, 0xbf, 0x24, 0xf4,
	0x5e, 0x8e, 0x56, 0x7f, 0x05, 0xf6, 0x0c, 0x8e, 0x2a, 0xd3, 0x98, 0x68, 0xaa, 0x5d, 0xe1, 0x15,
	0x09, 0x6c, 0xc8, 0xb6, 0x8d, 0xc7, 0x70, 0xbb, 0x53, 0xf1, 0xdb, 0x08, 0xbe, 0x26, 0x10, 0x68,
	0x94, 0x81, 0x25, 0xdc, 0xdf, 0x6c, 0x97, 0xff, 0xc4, 0xd6, 0xb4, 0x3e, 0x5c, 0xf1, 0x37, 0x84,
	0x1e, 0x8e, 0x21, 0x3d, 0xfa, 0x91, 0xa2, 0x54, 0xbb, 0x54, 0xfa, 0xfb, 0xba, 0xdb, 0x6d, 0xf1,
	0x36, 0xd7, 0xce, 0xd1, 0x56, 0x89, 0x72, 0x9f, 0x1e, 0x39, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x35, 0xfe, 0x23, 0x7c, 0x4b, 0x03, 0x00, 0x00,
}
