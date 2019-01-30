// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dpa_sysdb_hw_resource.proto

package cisco_ios_xr_fretta_bcm_dpa_hw_resources_oper_dpa_stats_nodes_node_hw_resources_datas_hw_resources_data

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

type DpaSysdbHwResource_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaSysdbHwResource_KEYS) Reset()         { *m = DpaSysdbHwResource_KEYS{} }
func (m *DpaSysdbHwResource_KEYS) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbHwResource_KEYS) ProtoMessage()    {}
func (*DpaSysdbHwResource_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b35bbb2e1f9511c, []int{0}
}

func (m *DpaSysdbHwResource_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbHwResource_KEYS.Unmarshal(m, b)
}
func (m *DpaSysdbHwResource_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbHwResource_KEYS.Marshal(b, m, deterministic)
}
func (m *DpaSysdbHwResource_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbHwResource_KEYS.Merge(m, src)
}
func (m *DpaSysdbHwResource_KEYS) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbHwResource_KEYS.Size(m)
}
func (m *DpaSysdbHwResource_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbHwResource_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbHwResource_KEYS proto.InternalMessageInfo

func (m *DpaSysdbHwResource_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *DpaSysdbHwResource_KEYS) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type DpaSysdbHwResource struct {
	ResourceId           uint32                   `protobuf:"varint,50,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Name                 string                   `protobuf:"bytes,51,opt,name=name,proto3" json:"name,omitempty"`
	NumNpus              uint32                   `protobuf:"varint,52,opt,name=num_npus,json=numNpus,proto3" json:"num_npus,omitempty"`
	NpuHwrList           []*DpaSysdbNpuHwResource `protobuf:"bytes,53,rep,name=npu_hwr_list,json=npuHwrList,proto3" json:"npu_hwr_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DpaSysdbHwResource) Reset()         { *m = DpaSysdbHwResource{} }
func (m *DpaSysdbHwResource) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbHwResource) ProtoMessage()    {}
func (*DpaSysdbHwResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b35bbb2e1f9511c, []int{1}
}

func (m *DpaSysdbHwResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbHwResource.Unmarshal(m, b)
}
func (m *DpaSysdbHwResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbHwResource.Marshal(b, m, deterministic)
}
func (m *DpaSysdbHwResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbHwResource.Merge(m, src)
}
func (m *DpaSysdbHwResource) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbHwResource.Size(m)
}
func (m *DpaSysdbHwResource) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbHwResource.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbHwResource proto.InternalMessageInfo

func (m *DpaSysdbHwResource) GetResourceId() uint32 {
	if m != nil {
		return m.ResourceId
	}
	return 0
}

func (m *DpaSysdbHwResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DpaSysdbHwResource) GetNumNpus() uint32 {
	if m != nil {
		return m.NumNpus
	}
	return 0
}

func (m *DpaSysdbHwResource) GetNpuHwrList() []*DpaSysdbNpuHwResource {
	if m != nil {
		return m.NpuHwrList
	}
	return nil
}

type DpaSysdbLtHwResource struct {
	LtId                 uint32   `protobuf:"varint,1,opt,name=lt_id,json=ltId,proto3" json:"lt_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	HwEntries            uint32   `protobuf:"varint,3,opt,name=hw_entries,json=hwEntries,proto3" json:"hw_entries,omitempty"`
	SwEntries            uint32   `protobuf:"varint,4,opt,name=sw_entries,json=swEntries,proto3" json:"sw_entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DpaSysdbLtHwResource) Reset()         { *m = DpaSysdbLtHwResource{} }
func (m *DpaSysdbLtHwResource) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbLtHwResource) ProtoMessage()    {}
func (*DpaSysdbLtHwResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b35bbb2e1f9511c, []int{2}
}

func (m *DpaSysdbLtHwResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbLtHwResource.Unmarshal(m, b)
}
func (m *DpaSysdbLtHwResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbLtHwResource.Marshal(b, m, deterministic)
}
func (m *DpaSysdbLtHwResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbLtHwResource.Merge(m, src)
}
func (m *DpaSysdbLtHwResource) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbLtHwResource.Size(m)
}
func (m *DpaSysdbLtHwResource) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbLtHwResource.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbLtHwResource proto.InternalMessageInfo

func (m *DpaSysdbLtHwResource) GetLtId() uint32 {
	if m != nil {
		return m.LtId
	}
	return 0
}

func (m *DpaSysdbLtHwResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DpaSysdbLtHwResource) GetHwEntries() uint32 {
	if m != nil {
		return m.HwEntries
	}
	return 0
}

func (m *DpaSysdbLtHwResource) GetSwEntries() uint32 {
	if m != nil {
		return m.SwEntries
	}
	return 0
}

type DpaSysdbNpuHwResource struct {
	MaxAllowed                uint32                  `protobuf:"varint,1,opt,name=max_allowed,json=maxAllowed,proto3" json:"max_allowed,omitempty"`
	NpuId                     uint32                  `protobuf:"varint,2,opt,name=npu_id,json=npuId,proto3" json:"npu_id,omitempty"`
	MaxEntries                uint32                  `protobuf:"varint,3,opt,name=max_entries,json=maxEntries,proto3" json:"max_entries,omitempty"`
	RedOorThreshold           uint32                  `protobuf:"varint,4,opt,name=red_oor_threshold,json=redOorThreshold,proto3" json:"red_oor_threshold,omitempty"`
	RedOorThresholdPercent    uint32                  `protobuf:"varint,5,opt,name=red_oor_threshold_percent,json=redOorThresholdPercent,proto3" json:"red_oor_threshold_percent,omitempty"`
	YellowOorThreshold        uint32                  `protobuf:"varint,6,opt,name=yellow_oor_threshold,json=yellowOorThreshold,proto3" json:"yellow_oor_threshold,omitempty"`
	YellowOorThresholdPercent uint32                  `protobuf:"varint,7,opt,name=yellow_oor_threshold_percent,json=yellowOorThresholdPercent,proto3" json:"yellow_oor_threshold_percent,omitempty"`
	InuseObjects              uint32                  `protobuf:"varint,8,opt,name=inuse_objects,json=inuseObjects,proto3" json:"inuse_objects,omitempty"`
	NumLt                     uint32                  `protobuf:"varint,9,opt,name=num_lt,json=numLt,proto3" json:"num_lt,omitempty"`
	OorChangeCount            uint32                  `protobuf:"varint,10,opt,name=oor_change_count,json=oorChangeCount,proto3" json:"oor_change_count,omitempty"`
	OorStateChangeTime1       string                  `protobuf:"bytes,11,opt,name=oor_state_change_time1,json=oorStateChangeTime1,proto3" json:"oor_state_change_time1,omitempty"`
	OorStateChangeTime2       string                  `protobuf:"bytes,12,opt,name=oor_state_change_time2,json=oorStateChangeTime2,proto3" json:"oor_state_change_time2,omitempty"`
	OorState                  string                  `protobuf:"bytes,13,opt,name=oor_state,json=oorState,proto3" json:"oor_state,omitempty"`
	LtHwrList                 []*DpaSysdbLtHwResource `protobuf:"bytes,14,rep,name=lt_hwr_list,json=ltHwrList,proto3" json:"lt_hwr_list,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                `json:"-"`
	XXX_unrecognized          []byte                  `json:"-"`
	XXX_sizecache             int32                   `json:"-"`
}

func (m *DpaSysdbNpuHwResource) Reset()         { *m = DpaSysdbNpuHwResource{} }
func (m *DpaSysdbNpuHwResource) String() string { return proto.CompactTextString(m) }
func (*DpaSysdbNpuHwResource) ProtoMessage()    {}
func (*DpaSysdbNpuHwResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b35bbb2e1f9511c, []int{3}
}

func (m *DpaSysdbNpuHwResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DpaSysdbNpuHwResource.Unmarshal(m, b)
}
func (m *DpaSysdbNpuHwResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DpaSysdbNpuHwResource.Marshal(b, m, deterministic)
}
func (m *DpaSysdbNpuHwResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DpaSysdbNpuHwResource.Merge(m, src)
}
func (m *DpaSysdbNpuHwResource) XXX_Size() int {
	return xxx_messageInfo_DpaSysdbNpuHwResource.Size(m)
}
func (m *DpaSysdbNpuHwResource) XXX_DiscardUnknown() {
	xxx_messageInfo_DpaSysdbNpuHwResource.DiscardUnknown(m)
}

var xxx_messageInfo_DpaSysdbNpuHwResource proto.InternalMessageInfo

func (m *DpaSysdbNpuHwResource) GetMaxAllowed() uint32 {
	if m != nil {
		return m.MaxAllowed
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetNpuId() uint32 {
	if m != nil {
		return m.NpuId
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetMaxEntries() uint32 {
	if m != nil {
		return m.MaxEntries
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetRedOorThreshold() uint32 {
	if m != nil {
		return m.RedOorThreshold
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetRedOorThresholdPercent() uint32 {
	if m != nil {
		return m.RedOorThresholdPercent
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetYellowOorThreshold() uint32 {
	if m != nil {
		return m.YellowOorThreshold
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetYellowOorThresholdPercent() uint32 {
	if m != nil {
		return m.YellowOorThresholdPercent
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetInuseObjects() uint32 {
	if m != nil {
		return m.InuseObjects
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetNumLt() uint32 {
	if m != nil {
		return m.NumLt
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetOorChangeCount() uint32 {
	if m != nil {
		return m.OorChangeCount
	}
	return 0
}

func (m *DpaSysdbNpuHwResource) GetOorStateChangeTime1() string {
	if m != nil {
		return m.OorStateChangeTime1
	}
	return ""
}

func (m *DpaSysdbNpuHwResource) GetOorStateChangeTime2() string {
	if m != nil {
		return m.OorStateChangeTime2
	}
	return ""
}

func (m *DpaSysdbNpuHwResource) GetOorState() string {
	if m != nil {
		return m.OorState
	}
	return ""
}

func (m *DpaSysdbNpuHwResource) GetLtHwrList() []*DpaSysdbLtHwResource {
	if m != nil {
		return m.LtHwrList
	}
	return nil
}

func init() {
	proto.RegisterType((*DpaSysdbHwResource_KEYS)(nil), "cisco_ios_xr_fretta_bcm_dpa_hw_resources_oper.dpa.stats.nodes.node.hw_resources_datas.hw_resources_data.dpa_sysdb_hw_resource_KEYS")
	proto.RegisterType((*DpaSysdbHwResource)(nil), "cisco_ios_xr_fretta_bcm_dpa_hw_resources_oper.dpa.stats.nodes.node.hw_resources_datas.hw_resources_data.dpa_sysdb_hw_resource")
	proto.RegisterType((*DpaSysdbLtHwResource)(nil), "cisco_ios_xr_fretta_bcm_dpa_hw_resources_oper.dpa.stats.nodes.node.hw_resources_datas.hw_resources_data.dpa_sysdb_lt_hw_resource")
	proto.RegisterType((*DpaSysdbNpuHwResource)(nil), "cisco_ios_xr_fretta_bcm_dpa_hw_resources_oper.dpa.stats.nodes.node.hw_resources_datas.hw_resources_data.dpa_sysdb_npu_hw_resource")
}

func init() { proto.RegisterFile("dpa_sysdb_hw_resource.proto", fileDescriptor_6b35bbb2e1f9511c) }

var fileDescriptor_6b35bbb2e1f9511c = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe5, 0xb6, 0x69, 0x93, 0x9b, 0xa4, 0xff, 0x3f, 0x53, 0x5a, 0x4d, 0x5a, 0x10, 0x51,
	0xd8, 0x44, 0x2c, 0x2c, 0x48, 0x60, 0xc1, 0x0a, 0xa1, 0xaa, 0x12, 0x15, 0x55, 0x8b, 0xd2, 0xb2,
	0x60, 0x35, 0x9a, 0x78, 0x86, 0xda, 0xc8, 0x9e, 0xb1, 0xe6, 0x43, 0x4e, 0x97, 0xe5, 0x05, 0x78,
	0x07, 0x24, 0xde, 0x8b, 0x47, 0x41, 0x33, 0x8e, 0xf3, 0xd1, 0x24, 0x5b, 0xd8, 0x44, 0xc9, 0x39,
	0xe7, 0x37, 0xf7, 0xfa, 0x66, 0xae, 0xe1, 0x84, 0xe5, 0x94, 0xe8, 0x3b, 0xcd, 0xc6, 0x24, 0x2e,
	0x88, 0xe2, 0x5a, 0x5a, 0x15, 0xf1, 0x30, 0x57, 0xd2, 0x48, 0x74, 0x1b, 0x25, 0x3a, 0x92, 0x24,
	0x91, 0x9a, 0x4c, 0x14, 0xf9, 0xaa, 0xb8, 0x31, 0x94, 0x8c, 0xa3, 0x8c, 0x38, 0x68, 0x21, 0xae,
	0x89, 0xcc, 0xb9, 0x0a, 0x59, 0x4e, 0x43, 0x6d, 0xa8, 0xd1, 0xa1, 0x90, 0x8c, 0x97, 0x9f, 0xe1,
	0x52, 0x8c, 0x51, 0x43, 0xf5, 0xaa, 0xd4, 0xfb, 0x0c, 0xc7, 0x6b, 0xfb, 0x20, 0x1f, 0xcf, 0xbe,
	0x5c, 0xa3, 0x13, 0x68, 0xb8, 0xa3, 0x88, 0xa0, 0x19, 0xc7, 0x41, 0x37, 0xe8, 0x37, 0x46, 0x75,
	0x27, 0x5c, 0xd2, 0x8c, 0xa3, 0x63, 0xa8, 0x57, 0x69, 0xbc, 0x55, 0x7a, 0xd5, 0xef, 0xde, 0x8f,
	0x2d, 0x38, 0x5c, 0x7b, 0x2e, 0x7a, 0x06, 0xcd, 0x59, 0x8d, 0x84, 0xe1, 0x41, 0x37, 0xe8, 0xb7,
	0x47, 0x50, 0x49, 0xe7, 0x0c, 0x21, 0xd8, 0xf1, 0xe5, 0x86, 0xfe, 0x48, 0xff, 0x1d, 0x75, 0xa0,
	0x2e, 0x6c, 0x46, 0x44, 0x6e, 0x35, 0x7e, 0xed, 0x89, 0x3d, 0x61, 0xb3, 0xcb, 0xdc, 0x6a, 0xf4,
	0x2b, 0x80, 0x96, 0xc8, 0x2d, 0x89, 0x0b, 0x45, 0xd2, 0x44, 0x1b, 0xfc, 0xa6, 0xbb, 0xdd, 0x6f,
	0x0e, 0xbe, 0x07, 0xe1, 0x5f, 0x1a, 0x61, 0x38, 0x7f, 0xce, 0xb2, 0x8f, 0x99, 0x3f, 0x02, 0x91,
	0xdb, 0x0f, 0x85, 0xba, 0x48, 0xb4, 0xe9, 0xdd, 0x07, 0x80, 0xe7, 0xc9, 0xd4, 0x2c, 0x0d, 0xe5,
	0x00, 0x6a, 0xa9, 0x71, 0xe3, 0x08, 0xfc, 0xc3, 0xed, 0xa4, 0x66, 0x61, 0x10, 0x5b, 0x0b, 0x83,
	0x78, 0x0a, 0x10, 0x17, 0x84, 0x0b, 0xa3, 0x12, 0xae, 0xf1, 0xb6, 0x4f, 0x37, 0xe2, 0xe2, 0xac,
	0x14, 0x9c, 0xad, 0xe7, 0xf6, 0x4e, 0x69, 0xeb, 0xca, 0xee, 0xfd, 0xae, 0x41, 0x67, 0x63, 0xb7,
	0xee, 0x9f, 0xc9, 0xe8, 0x84, 0xd0, 0x34, 0x95, 0x05, 0xaf, 0x5a, 0x81, 0x8c, 0x4e, 0xde, 0x97,
	0x0a, 0x3a, 0x84, 0x5d, 0xc7, 0x24, 0xcc, 0xb7, 0xd4, 0x1e, 0xd5, 0x44, 0x6e, 0xcf, 0x59, 0xc5,
	0x2d, 0x37, 0xe5, 0xb8, 0xaa, 0xab, 0x17, 0xf0, 0x48, 0x71, 0x46, 0xa4, 0x54, 0xc4, 0xc4, 0x8a,
	0xeb, 0x58, 0xa6, 0x6c, 0xda, 0xdc, 0x7f, 0x8a, 0xb3, 0x2b, 0xa9, 0x6e, 0x2a, 0x19, 0xbd, 0x85,
	0xce, 0x4a, 0x96, 0xe4, 0x5c, 0x45, 0x5c, 0x18, 0x5c, 0xf3, 0xcc, 0xd1, 0x03, 0xe6, 0x53, 0xe9,
	0xa2, 0x97, 0xf0, 0xf8, 0x8e, 0xbb, 0x56, 0x1f, 0x54, 0xda, 0xf5, 0x14, 0x2a, 0xbd, 0xa5, 0x62,
	0xef, 0xe0, 0xc9, 0x3a, 0x62, 0x56, 0x6f, 0xcf, 0x93, 0x9d, 0x55, 0xb2, 0x2a, 0xf9, 0x1c, 0xda,
	0x89, 0xb0, 0x9a, 0x13, 0x39, 0xfe, 0xc6, 0x23, 0xa3, 0x71, 0xdd, 0x13, 0x2d, 0x2f, 0x5e, 0x95,
	0x9a, 0x1f, 0x9b, 0xcd, 0x48, 0x6a, 0x70, 0x63, 0x3a, 0x36, 0x9b, 0x5d, 0x18, 0xd4, 0x87, 0xff,
	0x5d, 0xd5, 0x28, 0xa6, 0xe2, 0x96, 0x93, 0x48, 0x5a, 0x61, 0x30, 0xf8, 0xc0, 0xbe, 0x94, 0xea,
	0xd4, 0xcb, 0xa7, 0x4e, 0x45, 0x43, 0x38, 0x72, 0x49, 0x77, 0x3b, 0x79, 0x95, 0x37, 0x49, 0xc6,
	0x5f, 0xe1, 0xa6, 0xbf, 0x1a, 0x07, 0x52, 0xaa, 0x6b, 0x67, 0x96, 0xd0, 0x8d, 0xb3, 0x36, 0x42,
	0x03, 0xdc, 0xda, 0x04, 0x0d, 0xdc, 0xbe, 0xcf, 0x20, 0xdc, 0x2e, 0x77, 0xba, 0xca, 0xa1, 0x9f,
	0x01, 0x34, 0xfd, 0xbd, 0x9d, 0x2e, 0xda, 0xbe, 0x5f, 0xb4, 0xfb, 0x7f, 0xb1, 0x68, 0xcb, 0xeb,
	0x33, 0x6a, 0xa4, 0x66, 0xba, 0x66, 0xe3, 0x5d, 0xff, 0xfe, 0x1c, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x16, 0x96, 0x36, 0x4c, 0x5e, 0x05, 0x00, 0x00,
}
