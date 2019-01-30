// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_lspv_information

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
type Ipv4RibEdmProto_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()         { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()    {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4183d9ee6d1994, []int{0}
}

func (m *Ipv4RibEdmProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Unmarshal(m, b)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmProto_KEYS.Merge(m, src)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Size(m)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmProto_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmProto_KEYS proto.InternalMessageInfo

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

type Ipv4RibEdmProto struct {
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

func (m *Ipv4RibEdmProto) Reset()         { *m = Ipv4RibEdmProto{} }
func (m *Ipv4RibEdmProto) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto) ProtoMessage()    {}
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d4183d9ee6d1994, []int{1}
}

func (m *Ipv4RibEdmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmProto.Unmarshal(m, b)
}
func (m *Ipv4RibEdmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmProto.Marshal(b, m, deterministic)
}
func (m *Ipv4RibEdmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmProto.Merge(m, src)
}
func (m *Ipv4RibEdmProto) XXX_Size() int {
	return xxx_messageInfo_Ipv4RibEdmProto.Size(m)
}
func (m *Ipv4RibEdmProto) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4RibEdmProto.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4RibEdmProto proto.InternalMessageInfo

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
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv4_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor_9d4183d9ee6d1994) }

var fileDescriptor_9d4183d9ee6d1994 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0xf8, 0x8f, 0xb8, 0x54, 0x9c, 0x10, 0x1c, 0x02, 0x17, 0x4e, 0x96, 0xd7, 0xeb,
	0x15, 0x16, 0xbb, 0xf6, 0xca, 0xe3, 0xac, 0xe8, 0x63, 0xc0, 0xab, 0xf2, 0x02, 0xc8, 0xe3, 0xdd,
	0x90, 0x50, 0xf5, 0xe2, 0xc8, 0xf3, 0xfd, 0x7e, 0xce, 0x37, 0x09, 0x70, 0xdb, 0xf5, 0x4b, 0x19,
	0x6c, 0x29, 0x4d, 0xd5, 0xca, 0x2e, 0xf8, 0xe8, 0x05, 0x9d, 0xec, 0x57, 0xa1, 0x2d, 0x6a, 0x2f,
	0xad, 0x47, 0xf9, 0x33, 0x48, 0xdb, 0x11, 0x45, 0xb8, 0xef, 0x4c, 0x10, 0xe9, 0x86, 0xb1, 0x2a,
	0x2f, 0x45, 0x1f, 0x6a, 0x4c, 0x87, 0x50, 0x35, 0x0a, 0x55, 0x0b, 0x4c, 0x9f, 0xa8, 0x6a, 0x31,
	0x38, 0xc1, 0xaf, 0xa3, 0x91, 0x51, 0x95, 0x8d, 0x91, 0x4e, 0xb5, 0x06, 0xaf, 0x0b, 0xf2, 0x37,
	0x6b, 0xdf, 0x88, 0xc6, 0x6b, 0xd5, 0x88, 0x06, 0xbb, 0x5e, 0x58, 0x57, 0xfb, 0xd0, 0xaa, 0x68,
	0xbd, 0x3b, 0xfd, 0x5d, 0xc0, 0xf1, 0xd5, 0xc2, 0xf2, 0xe3, 0x87, 0x6f, 0x5f, 0xd8, 0x1c, 0xa6,
	0x7d, 0xa8, 0xe9, 0x1d, 0x5e, 0x9c, 0x14, 0x8b, 0x5b, 0xab, 0x49, 0x1f, 0xea, 0xcf, 0xaa, 0x35,
	0xec, 0x18, 0x26, 0x6a, 0x48, 0x6e, 0x50, 0xb2, 0xaf, 0x72, 0x30, 0x87, 0x29, 0x8e, 0xc9, 0x5e,
	0x76, 0x70, 0x88, 0x16, 0x70, 0xf7, 0xff, 0x7a, 0xfc, 0x26, 0x21, 0x07, 0x34, 0xff, 0x9a, 0xc6,
	0x89, 0x3c, 0xfd, 0xb3, 0x07, 0xec, 0x6a, 0x29, 0xf6, 0x14, 0x0e, 0xc6, 0x75, 0xf2, 0xd6, 0xfc,
	0x8c, 0xf4, 0xd9, 0x38, 0x4d, 0x32, 0xb2, 0x87, 0x30, 0xb5, 0x0e, 0xa3, 0x72, 0xda, 0xf0, 0x73,
	0x02, 0x36, 0x77, 0xc6, 0x61, 0xd2, 0x9b, 0x80, 0xd6, 0x3b, 0xbe, 0x3c, 0x29, 0x16, 0xb3, 0xd5,
	0x78, 0x65, 0xef, 0xe1, 0x51, 0x30, 0x95, 0xc5, 0x18, 0x6c, 0xb9, 0x4e, 0x3f, 0x8d, 0xd4, 0x8d,
	0x35, 0x2e, 0x4a, 0xed, 0xd7, 0x2e, 0xf2, 0xe7, 0x44, 0xcf, 0x77, 0x91, 0x0b, 0x22, 0x2e, 0x12,
	0xc0, 0x96, 0xf0, 0x60, 0x53, 0x2e, 0x9b, 0x38, 0xa8, 0x2f, 0x48, 0x3d, 0x1a, 0xd3, 0x2c, 0x61,
	0xb6, 0x9e, 0xc0, 0x8c, 0x76, 0x1f, 0x58, 0xe4, 0x2f, 0x09, 0xbe, 0x93, 0x87, 0xc4, 0x20, 0x13,
	0x70, 0xa8, 0x74, 0xb4, 0xbd, 0x91, 0xdb, 0x2c, 0x7f, 0x45, 0xe8, 0xbd, 0x1c, 0xad, 0xfe, 0x09,
	0xec, 0x19, 0x1c, 0x55, 0xa6, 0x31, 0xd1, 0x54, 0xbb, 0xc2, 0x6b, 0x12, 0xd8, 0x90, 0x6d, 0x1b,
	0x8f, 0xe1, 0x76, 0xa7, 0xe2, 0xf7, 0x11, 0x7c, 0x43, 0x20, 0xd0, 0x28, 0x03, 0x67, 0x70, 0x7f,
	0xb3, 0x5d, 0xfe, 0x13, 0x5b, 0xd3, 0xfa, 0x70, 0xc9, 0xdf, 0x12, 0x7a, 0x38, 0x86, 0xf4, 0xe8,
	0x27, 0x8a, 0x52, 0xed, 0x52, 0xe9, 0x1f, 0xeb, 0x6e, 0xb7, 0xc5, 0xbb, 0x5c, 0x3b, 0x47, 0x5b,
	0x25, 0xca, 0x7d, 0x7a, 0xe4, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x19, 0x3c, 0x93,
	0x41, 0x03, 0x00, 0x00,
}
