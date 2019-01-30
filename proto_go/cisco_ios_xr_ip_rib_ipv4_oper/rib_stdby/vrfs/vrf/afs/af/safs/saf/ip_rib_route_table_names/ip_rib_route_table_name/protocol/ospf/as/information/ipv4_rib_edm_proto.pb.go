// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_ospf_as_information

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName       string   `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	As                   string   `protobuf:"bytes,5,opt,name=as,proto3" json:"as,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4RibEdmProto_KEYS) Reset()         { *m = Ipv4RibEdmProto_KEYS{} }
func (m *Ipv4RibEdmProto_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4RibEdmProto_KEYS) ProtoMessage()    {}
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipv4_rib_edm_proto_7822d375ec4c7cc5, []int{0}
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Unmarshal(m, b)
}
func (m *Ipv4RibEdmProto_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmProto_KEYS.Marshal(b, m, deterministic)
}
func (dst *Ipv4RibEdmProto_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmProto_KEYS.Merge(dst, src)
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

func (m *Ipv4RibEdmProto_KEYS) GetAs() string {
	if m != nil {
		return m.As
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
	return fileDescriptor_ipv4_rib_edm_proto_7822d375ec4c7cc5, []int{1}
}
func (m *Ipv4RibEdmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4RibEdmProto.Unmarshal(m, b)
}
func (m *Ipv4RibEdmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4RibEdmProto.Marshal(b, m, deterministic)
}
func (dst *Ipv4RibEdmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4RibEdmProto.Merge(dst, src)
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
	proto.RegisterType((*Ipv4RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.ospf.as.information.ipv4_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv4RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.ospf.as.information.ipv4_rib_edm_proto")
}

func init() {
	proto.RegisterFile("ipv4_rib_edm_proto.proto", fileDescriptor_ipv4_rib_edm_proto_7822d375ec4c7cc5)
}

var fileDescriptor_ipv4_rib_edm_proto_7822d375ec4c7cc5 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x14, 0x9a, 0x30, 0x90, 0x08, 0xb6, 0x85, 0x6e, 0xe0, 0x40, 0x55, 0x84, 0x94,
	0xd3, 0x0a, 0xb5, 0xe1, 0x3f, 0xe2, 0x52, 0x71, 0x42, 0x70, 0x08, 0x5c, 0x38, 0xad, 0xc6, 0xf6,
	0x5a, 0xac, 0x88, 0xbd, 0xd6, 0xce, 0xda, 0xa2, 0x2f, 0xc1, 0x43, 0xf0, 0x8a, 0xbc, 0x00, 0xf2,
	0xac, 0x1d, 0x12, 0xa2, 0x5e, 0x1c, 0xed, 0xfc, 0x7e, 0xdf, 0xe4, 0xf3, 0x1a, 0xa4, 0xad, 0xdb,
	0xa5, 0xf6, 0x36, 0xd5, 0x26, 0x2f, 0x75, 0xed, 0x5d, 0x70, 0x8a, 0x9f, 0xe2, 0x57, 0x92, 0x59,
	0xca, 0x9c, 0xb6, 0x8e, 0xf4, 0x4f, 0xaf, 0x6d, 0xcd, 0x16, 0xeb, 0xae, 0x36, 0x5e, 0x75, 0x27,
	0x0a, 0x79, 0x7a, 0xa5, 0x5a, 0x5f, 0x50, 0xf7, 0x50, 0x58, 0x90, 0xc2, 0x42, 0x51, 0xf7, 0x4b,
	0x58, 0xa8, 0x3e, 0xe3, 0x5d, 0x13, 0x8c, 0x0e, 0x98, 0xae, 0x8d, 0xae, 0xb0, 0x34, 0x74, 0x1d,
	0x88, 0xff, 0x9c, 0xb9, 0xb5, 0x72, 0x54, 0x17, 0x0a, 0x49, 0xd9, 0xaa, 0x70, 0xbe, 0xc4, 0x60,
	0x5d, 0x75, 0xf6, 0x3b, 0x81, 0x93, 0xfd, 0xb6, 0xfa, 0xe3, 0x87, 0x6f, 0x5f, 0xc4, 0x1c, 0x26,
	0xad, 0x2f, 0x78, 0x89, 0x4c, 0x4e, 0x93, 0xc5, 0xad, 0xd5, 0xb8, 0xf5, 0xc5, 0x67, 0x2c, 0x8d,
	0x38, 0x81, 0x31, 0xf6, 0x64, 0xc4, 0xe4, 0x10, 0x23, 0x98, 0xc3, 0x84, 0x06, 0x72, 0x10, 0x33,
	0xd4, 0xa3, 0x05, 0xdc, 0xfd, 0xbf, 0x9b, 0xbc, 0xc1, 0xca, 0x8c, 0xe7, 0x5f, 0xbb, 0x31, 0x9b,
	0x33, 0x18, 0x21, 0xc9, 0x9b, 0xcc, 0x46, 0x48, 0x67, 0x7f, 0x0e, 0x40, 0xec, 0x97, 0x14, 0x4f,
	0x61, 0x36, 0xbc, 0x5b, 0xbc, 0x02, 0x79, 0xce, 0x91, 0xe9, 0x30, 0xed, 0x96, 0x91, 0x78, 0x08,
	0x13, 0x5b, 0x51, 0xc0, 0x2a, 0x33, 0xf2, 0x82, 0x85, 0xcd, 0x59, 0x48, 0x18, 0xb7, 0xc6, 0x93,
	0x75, 0x95, 0x5c, 0x9e, 0x26, 0x8b, 0xe9, 0x6a, 0x38, 0x8a, 0xf7, 0xf0, 0xc8, 0x9b, 0xdc, 0x52,
	0xf0, 0x36, 0x6d, 0xba, 0xab, 0xd2, 0xd9, 0xda, 0x9a, 0x2a, 0xe8, 0xcc, 0x35, 0x55, 0x90, 0xcf,
	0xd9, 0x9e, 0xef, 0x2a, 0x97, 0x6c, 0x5c, 0x76, 0x82, 0x58, 0xc2, 0x83, 0x4d, 0xb9, 0x98, 0xa4,
	0x3e, 0xfa, 0x82, 0xa3, 0xc7, 0x03, 0x8d, 0x21, 0x8a, 0xa9, 0x27, 0x30, 0xe5, 0xbb, 0xe8, 0x5d,
	0x92, 0x2f, 0x59, 0xbe, 0x13, 0x87, 0xec, 0x90, 0x50, 0x70, 0x84, 0x59, 0xb0, 0xad, 0xd1, 0xdb,
	0xae, 0x7c, 0xc5, 0xea, 0xbd, 0x88, 0x56, 0xff, 0x02, 0xe2, 0x19, 0x1c, 0xe7, 0x66, 0x6d, 0x82,
	0xc9, 0x77, 0x03, 0xaf, 0x39, 0x20, 0x7a, 0xb6, 0x9d, 0x78, 0x0c, 0xb7, 0x6b, 0x0c, 0xdf, 0x07,
	0xf1, 0x0d, 0x8b, 0xc0, 0xa3, 0x28, 0x9c, 0xc3, 0xfd, 0xcd, 0xdb, 0xc5, 0x8f, 0x5a, 0x9a, 0xd2,
	0xf9, 0x2b, 0xf9, 0x96, 0xd5, 0xa3, 0x01, 0xf2, 0xd2, 0x4f, 0x8c, 0xba, 0xda, 0x29, 0x66, 0x3f,
	0x9a, 0x7a, 0xb7, 0xc5, 0xbb, 0x58, 0x3b, 0xa2, 0xad, 0x12, 0xe9, 0x21, 0x2f, 0xb9, 0xf8, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0x40, 0x9c, 0xa8, 0x4e, 0x03, 0x00, 0x00,
}
