// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_adj_log_struc.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_levels_level_adjacency_log

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

// Adjacency log structure
type IsisShAdjLogStruc_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShAdjLogStruc_KEYS) Reset()         { *m = IsisShAdjLogStruc_KEYS{} }
func (m *IsisShAdjLogStruc_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShAdjLogStruc_KEYS) ProtoMessage()    {}
func (*IsisShAdjLogStruc_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{0}
}
func (m *IsisShAdjLogStruc_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShAdjLogStruc_KEYS.Unmarshal(m, b)
}
func (m *IsisShAdjLogStruc_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShAdjLogStruc_KEYS.Marshal(b, m, deterministic)
}
func (dst *IsisShAdjLogStruc_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShAdjLogStruc_KEYS.Merge(dst, src)
}
func (m *IsisShAdjLogStruc_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShAdjLogStruc_KEYS.Size(m)
}
func (m *IsisShAdjLogStruc_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShAdjLogStruc_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShAdjLogStruc_KEYS proto.InternalMessageInfo

func (m *IsisShAdjLogStruc_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShAdjLogStruc_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type IsisShAdjLogStruc struct {
	// Adjacency Log entries
	LogEntry             []*IsisShAdjLogEnt `protobuf:"bytes,50,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IsisShAdjLogStruc) Reset()         { *m = IsisShAdjLogStruc{} }
func (m *IsisShAdjLogStruc) String() string { return proto.CompactTextString(m) }
func (*IsisShAdjLogStruc) ProtoMessage()    {}
func (*IsisShAdjLogStruc) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{1}
}
func (m *IsisShAdjLogStruc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShAdjLogStruc.Unmarshal(m, b)
}
func (m *IsisShAdjLogStruc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShAdjLogStruc.Marshal(b, m, deterministic)
}
func (dst *IsisShAdjLogStruc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShAdjLogStruc.Merge(dst, src)
}
func (m *IsisShAdjLogStruc) XXX_Size() int {
	return xxx_messageInfo_IsisShAdjLogStruc.Size(m)
}
func (m *IsisShAdjLogStruc) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShAdjLogStruc.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShAdjLogStruc proto.InternalMessageInfo

func (m *IsisShAdjLogStruc) GetLogEntry() []*IsisShAdjLogEnt {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

// Identification of an IS-IS topology
type IsisTopoIdType struct {
	// AF name
	AfName string `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	// Sub-AF name
	SafName string `protobuf:"bytes,2,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	// VRF Name
	VrfName string `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	// Topology Name
	TopologyName         string   `protobuf:"bytes,4,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisTopoIdType) Reset()         { *m = IsisTopoIdType{} }
func (m *IsisTopoIdType) String() string { return proto.CompactTextString(m) }
func (*IsisTopoIdType) ProtoMessage()    {}
func (*IsisTopoIdType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{2}
}
func (m *IsisTopoIdType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisTopoIdType.Unmarshal(m, b)
}
func (m *IsisTopoIdType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisTopoIdType.Marshal(b, m, deterministic)
}
func (dst *IsisTopoIdType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisTopoIdType.Merge(dst, src)
}
func (m *IsisTopoIdType) XXX_Size() int {
	return xxx_messageInfo_IsisTopoIdType.Size(m)
}
func (m *IsisTopoIdType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisTopoIdType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisTopoIdType proto.InternalMessageInfo

func (m *IsisTopoIdType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisTopoIdType) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisTopoIdType) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IsisTopoIdType) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

// Per-topology per-adjacency log data
type IsisAdjLogTopoType struct {
	// Topology ID
	Id *IsisTopoIdType `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Change
	Change               string   `protobuf:"bytes,2,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisAdjLogTopoType) Reset()         { *m = IsisAdjLogTopoType{} }
func (m *IsisAdjLogTopoType) String() string { return proto.CompactTextString(m) }
func (*IsisAdjLogTopoType) ProtoMessage()    {}
func (*IsisAdjLogTopoType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{3}
}
func (m *IsisAdjLogTopoType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisAdjLogTopoType.Unmarshal(m, b)
}
func (m *IsisAdjLogTopoType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisAdjLogTopoType.Marshal(b, m, deterministic)
}
func (dst *IsisAdjLogTopoType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisAdjLogTopoType.Merge(dst, src)
}
func (m *IsisAdjLogTopoType) XXX_Size() int {
	return xxx_messageInfo_IsisAdjLogTopoType.Size(m)
}
func (m *IsisAdjLogTopoType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisAdjLogTopoType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisAdjLogTopoType proto.InternalMessageInfo

func (m *IsisAdjLogTopoType) GetId() *IsisTopoIdType {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *IsisAdjLogTopoType) GetChange() string {
	if m != nil {
		return m.Change
	}
	return ""
}

// Timestamp for an event
type IsisShTimestampType struct {
	// Timestamp value (seconds)
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Timestamp value (nanoseconds)
	NanoSeconds          uint32   `protobuf:"varint,2,opt,name=nano_seconds,json=nanoSeconds,proto3" json:"nano_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShTimestampType) Reset()         { *m = IsisShTimestampType{} }
func (m *IsisShTimestampType) String() string { return proto.CompactTextString(m) }
func (*IsisShTimestampType) ProtoMessage()    {}
func (*IsisShTimestampType) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{4}
}
func (m *IsisShTimestampType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTimestampType.Unmarshal(m, b)
}
func (m *IsisShTimestampType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTimestampType.Marshal(b, m, deterministic)
}
func (dst *IsisShTimestampType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTimestampType.Merge(dst, src)
}
func (m *IsisShTimestampType) XXX_Size() int {
	return xxx_messageInfo_IsisShTimestampType.Size(m)
}
func (m *IsisShTimestampType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTimestampType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTimestampType proto.InternalMessageInfo

func (m *IsisShTimestampType) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *IsisShTimestampType) GetNanoSeconds() uint32 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

// Generic portion of a log entry
type IsisShGenericLogEnt struct {
	// Time in UTC relative to Jan 1st, 1970
	Timestamp            *IsisShTimestampType `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IsisShGenericLogEnt) Reset()         { *m = IsisShGenericLogEnt{} }
func (m *IsisShGenericLogEnt) String() string { return proto.CompactTextString(m) }
func (*IsisShGenericLogEnt) ProtoMessage()    {}
func (*IsisShGenericLogEnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{5}
}
func (m *IsisShGenericLogEnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShGenericLogEnt.Unmarshal(m, b)
}
func (m *IsisShGenericLogEnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShGenericLogEnt.Marshal(b, m, deterministic)
}
func (dst *IsisShGenericLogEnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShGenericLogEnt.Merge(dst, src)
}
func (m *IsisShGenericLogEnt) XXX_Size() int {
	return xxx_messageInfo_IsisShGenericLogEnt.Size(m)
}
func (m *IsisShGenericLogEnt) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShGenericLogEnt.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShGenericLogEnt proto.InternalMessageInfo

func (m *IsisShGenericLogEnt) GetTimestamp() *IsisShTimestampType {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// An adjacency log entry
type IsisShAdjLogEnt struct {
	// Generic entry data
	GenericData *IsisShGenericLogEnt `protobuf:"bytes,1,opt,name=generic_data,json=genericData,proto3" json:"generic_data,omitempty"`
	// Neighbor system ID
	AdjLogNeighborSystemId string `protobuf:"bytes,2,opt,name=adj_log_neighbor_system_id,json=adjLogNeighborSystemId,proto3" json:"adj_log_neighbor_system_id,omitempty"`
	// Interface name
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Previous adjacency state
	PrevAdjState string `protobuf:"bytes,4,opt,name=prev_adj_state,json=prevAdjState,proto3" json:"prev_adj_state,omitempty"`
	// Current adjacency state
	CurAdjState string `protobuf:"bytes,5,opt,name=cur_adj_state,json=curAdjState,proto3" json:"cur_adj_state,omitempty"`
	// Reason adjacency changed state
	StateReason string `protobuf:"bytes,6,opt,name=state_reason,json=stateReason,proto3" json:"state_reason,omitempty"`
	// Per-topology changes
	AdjacencyPerTopologyChanges []*IsisAdjLogTopoType `protobuf:"bytes,7,rep,name=adjacency_per_topology_changes,json=adjacencyPerTopologyChanges,proto3" json:"adjacency_per_topology_changes,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}              `json:"-"`
	XXX_unrecognized            []byte                `json:"-"`
	XXX_sizecache               int32                 `json:"-"`
}

func (m *IsisShAdjLogEnt) Reset()         { *m = IsisShAdjLogEnt{} }
func (m *IsisShAdjLogEnt) String() string { return proto.CompactTextString(m) }
func (*IsisShAdjLogEnt) ProtoMessage()    {}
func (*IsisShAdjLogEnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875, []int{6}
}
func (m *IsisShAdjLogEnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShAdjLogEnt.Unmarshal(m, b)
}
func (m *IsisShAdjLogEnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShAdjLogEnt.Marshal(b, m, deterministic)
}
func (dst *IsisShAdjLogEnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShAdjLogEnt.Merge(dst, src)
}
func (m *IsisShAdjLogEnt) XXX_Size() int {
	return xxx_messageInfo_IsisShAdjLogEnt.Size(m)
}
func (m *IsisShAdjLogEnt) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShAdjLogEnt.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShAdjLogEnt proto.InternalMessageInfo

func (m *IsisShAdjLogEnt) GetGenericData() *IsisShGenericLogEnt {
	if m != nil {
		return m.GenericData
	}
	return nil
}

func (m *IsisShAdjLogEnt) GetAdjLogNeighborSystemId() string {
	if m != nil {
		return m.AdjLogNeighborSystemId
	}
	return ""
}

func (m *IsisShAdjLogEnt) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IsisShAdjLogEnt) GetPrevAdjState() string {
	if m != nil {
		return m.PrevAdjState
	}
	return ""
}

func (m *IsisShAdjLogEnt) GetCurAdjState() string {
	if m != nil {
		return m.CurAdjState
	}
	return ""
}

func (m *IsisShAdjLogEnt) GetStateReason() string {
	if m != nil {
		return m.StateReason
	}
	return ""
}

func (m *IsisShAdjLogEnt) GetAdjacencyPerTopologyChanges() []*IsisAdjLogTopoType {
	if m != nil {
		return m.AdjacencyPerTopologyChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShAdjLogStruc_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_sh_adj_log_struc_KEYS")
	proto.RegisterType((*IsisShAdjLogStruc)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_sh_adj_log_struc")
	proto.RegisterType((*IsisTopoIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_topo_id_type")
	proto.RegisterType((*IsisAdjLogTopoType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_adj_log_topo_type")
	proto.RegisterType((*IsisShTimestampType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_sh_timestamp_type")
	proto.RegisterType((*IsisShGenericLogEnt)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_sh_generic_log_ent")
	proto.RegisterType((*IsisShAdjLogEnt)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.adjacency_log.isis_sh_adj_log_ent")
}

func init() {
	proto.RegisterFile("isis_sh_adj_log_struc.proto", fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875)
}

var fileDescriptor_isis_sh_adj_log_struc_89df20c27fb77875 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xba, 0xad, 0xdd, 0x5e, 0xdb, 0x49, 0x18, 0xe8, 0xc2, 0x26, 0xa1, 0x2d, 0x03, 0x69,
	0xa7, 0x1c, 0xca, 0x8d, 0x1b, 0x82, 0x1d, 0x10, 0xa8, 0x42, 0x29, 0x08, 0x71, 0xb2, 0x3c, 0xdb,
	0x4d, 0x5d, 0xa5, 0x76, 0xb0, 0xdd, 0x8a, 0xde, 0x38, 0x20, 0x21, 0xf1, 0x13, 0x38, 0x4d, 0xe2,
	0xc8, 0x9f, 0x44, 0xb1, 0xe3, 0xb4, 0x8c, 0x1e, 0xe9, 0x29, 0x79, 0xdf, 0xf7, 0xe9, 0xbd, 0xef,
	0x3d, 0x3f, 0x1b, 0xce, 0x84, 0x11, 0x06, 0x9b, 0x29, 0x26, 0x6c, 0x86, 0x0b, 0x95, 0x63, 0x63,
	0xf5, 0x82, 0xa6, 0xa5, 0x56, 0x56, 0xa1, 0x11, 0x15, 0x86, 0x2a, 0x2c, 0x94, 0xc1, 0x5f, 0x34,
	0xa6, 0x85, 0x34, 0xd8, 0xc9, 0x55, 0xc9, 0x75, 0x5a, 0xfd, 0xa5, 0x42, 0x1a, 0x4b, 0x24, 0xe5,
	0xeb, 0xbf, 0xb4, 0xe0, 0x4b, 0x5e, 0x18, 0xff, 0x49, 0x09, 0x9b, 0x11, 0xca, 0x25, 0x5d, 0x55,
	0xb9, 0x93, 0x8f, 0x70, 0xba, 0xb5, 0x1c, 0x7e, 0x73, 0xfd, 0x69, 0x8c, 0x2e, 0xa1, 0x1f, 0x92,
	0x60, 0x49, 0xe6, 0x3c, 0x8e, 0xce, 0xa3, 0xab, 0xa3, 0xac, 0x17, 0xc0, 0x11, 0x99, 0x73, 0xf4,
	0x00, 0x0e, 0x5c, 0xe6, 0xb8, 0xe5, 0x48, 0x1f, 0x24, 0x3f, 0x23, 0x78, 0xb8, 0x35, 0x33, 0xfa,
	0x1a, 0xc1, 0x51, 0x15, 0x71, 0x69, 0xf5, 0x2a, 0x1e, 0x9e, 0xef, 0x5d, 0x75, 0x87, 0x34, 0xfd,
	0xbf, 0x7d, 0xa5, 0x77, 0x4b, 0x73, 0x69, 0xb3, 0xc3, 0x42, 0xe5, 0xd7, 0x55, 0xd1, 0xe4, 0x7b,
	0x04, 0xf7, 0x9c, 0xc2, 0xaa, 0x52, 0x61, 0xc1, 0xb0, 0x5d, 0x95, 0x1c, 0x9d, 0x40, 0x87, 0x4c,
	0x36, 0xfb, 0x6c, 0x93, 0x89, 0xeb, 0xf0, 0x11, 0x1c, 0x9a, 0xc0, 0xf8, 0x26, 0x3b, 0x66, 0x4d,
	0x2d, 0x75, 0x4d, 0xed, 0x79, 0x6a, 0xa9, 0x3d, 0x75, 0x09, 0xfd, 0x2a, 0x7d, 0xa1, 0xf2, 0x95,
	0xe7, 0xf7, 0xfd, 0xf0, 0x02, 0x58, 0x89, 0x92, 0x5f, 0x11, 0x0c, 0x9c, 0x93, 0x60, 0xd4, 0x39,
	0x72, 0x76, 0x3e, 0x43, 0x4b, 0x30, 0xe7, 0xa4, 0x3b, 0x24, 0x3b, 0x99, 0xcf, 0x66, 0xf7, 0x59,
	0x4b, 0x30, 0x34, 0x80, 0x36, 0x9d, 0x12, 0x99, 0x87, 0x36, 0xeb, 0x28, 0xf9, 0x50, 0x9b, 0x34,
	0x53, 0x6c, 0xc5, 0x9c, 0x1b, 0x4b, 0xe6, 0xa5, 0x37, 0x19, 0x43, 0xc7, 0x70, 0xaa, 0x24, 0x33,
	0xce, 0x69, 0x3f, 0x0b, 0x21, 0xba, 0x80, 0x9e, 0x24, 0x52, 0xe1, 0x40, 0xb7, 0x1c, 0xdd, 0xad,
	0xb0, 0xb1, 0x87, 0x92, 0xdb, 0x08, 0x4e, 0x42, 0xde, 0x9c, 0x4b, 0xae, 0x05, 0x0d, 0x87, 0x85,
	0xbe, 0x45, 0x70, 0xd4, 0xd4, 0xaa, 0xa7, 0x30, 0xd9, 0xd5, 0x96, 0xfc, 0xdd, 0x54, 0xb6, 0x2e,
	0x9c, 0xdc, 0xee, 0xc3, 0xfd, 0x2d, 0xbb, 0x84, 0x7e, 0x44, 0xd0, 0x0b, 0x96, 0x19, 0xb1, 0xa4,
	0x76, 0x98, 0xef, 0xca, 0xe1, 0x9d, 0xf1, 0x64, 0xdd, 0x1a, 0x78, 0x45, 0x2c, 0x41, 0xcf, 0xe1,
	0x34, 0x78, 0x93, 0x5c, 0xe4, 0xd3, 0x1b, 0xa5, 0xb1, 0x59, 0x19, 0xcb, 0xe7, 0x58, 0xb0, 0xfa,
	0x28, 0x07, 0x84, 0xcd, 0xde, 0xaa, 0x7c, 0x54, 0xf3, 0x63, 0x47, 0xbf, 0x66, 0xe8, 0x29, 0x1c,
	0x0b, 0x69, 0xb9, 0x9e, 0x90, 0x70, 0xc7, 0xfd, 0x1a, 0xf7, 0x1b, 0xd4, 0x2d, 0xf3, 0x13, 0x38,
	0x2e, 0x35, 0x5f, 0xba, 0x19, 0x18, 0x4b, 0x6c, 0xb3, 0xcd, 0x15, 0xfa, 0x82, 0xcd, 0xc6, 0x15,
	0x86, 0x12, 0xe8, 0xd3, 0x85, 0xde, 0x10, 0x1d, 0x38, 0x51, 0x97, 0x2e, 0x74, 0xa3, 0xb9, 0x80,
	0x9e, 0xe3, 0xb0, 0xe6, 0xc4, 0x28, 0x19, 0xb7, 0xbd, 0xc4, 0x61, 0x99, 0x83, 0xd0, 0xef, 0x08,
	0x1e, 0xaf, 0xc7, 0x50, 0x72, 0x8d, 0x9b, 0x8b, 0xe4, 0x17, 0xd2, 0xc4, 0x1d, 0xf7, 0x6c, 0xec,
	0x66, 0x21, 0xfe, 0xb9, 0x8a, 0xd9, 0x59, 0xa3, 0x7a, 0xc7, 0xf5, 0xfb, 0xda, 0xcb, 0x4b, 0x6f,
	0xe5, 0xa6, 0xed, 0x5e, 0xe6, 0x67, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x08, 0x04, 0x05,
	0xb8, 0x05, 0x00, 0x00,
}
