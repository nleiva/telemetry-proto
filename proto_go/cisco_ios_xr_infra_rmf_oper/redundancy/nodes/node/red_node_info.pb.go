// Code generated by protoc-gen-go. DO NOT EDIT.
// source: red_node_info.proto

package cisco_ios_xr_infra_rmf_oper_redundancy_nodes_node

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

// Node Information
type RedNodeInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedNodeInfo_KEYS) Reset()         { *m = RedNodeInfo_KEYS{} }
func (m *RedNodeInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RedNodeInfo_KEYS) ProtoMessage()    {}
func (*RedNodeInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8501aa16198dc9, []int{0}
}

func (m *RedNodeInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedNodeInfo_KEYS.Unmarshal(m, b)
}
func (m *RedNodeInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedNodeInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RedNodeInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedNodeInfo_KEYS.Merge(m, src)
}
func (m *RedNodeInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RedNodeInfo_KEYS.Size(m)
}
func (m *RedNodeInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RedNodeInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RedNodeInfo_KEYS proto.InternalMessageInfo

func (m *RedNodeInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type RedNodeInfo struct {
	// Row information
	Redundancy *RedSummaryPair `protobuf:"bytes,50,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	// Reload and boot logs
	Log string `protobuf:"bytes,51,opt,name=log,proto3" json:"log,omitempty"`
	// Active node reload
	ActiveRebootReason string `protobuf:"bytes,52,opt,name=active_reboot_reason,json=activeRebootReason,proto3" json:"active_reboot_reason,omitempty"`
	// Standby node reload
	StandbyRebootReason string `protobuf:"bytes,53,opt,name=standby_reboot_reason,json=standbyRebootReason,proto3" json:"standby_reboot_reason,omitempty"`
	// Error Log
	ErrLog               string   `protobuf:"bytes,54,opt,name=err_log,json=errLog,proto3" json:"err_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedNodeInfo) Reset()         { *m = RedNodeInfo{} }
func (m *RedNodeInfo) String() string { return proto.CompactTextString(m) }
func (*RedNodeInfo) ProtoMessage()    {}
func (*RedNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8501aa16198dc9, []int{1}
}

func (m *RedNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedNodeInfo.Unmarshal(m, b)
}
func (m *RedNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedNodeInfo.Marshal(b, m, deterministic)
}
func (m *RedNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedNodeInfo.Merge(m, src)
}
func (m *RedNodeInfo) XXX_Size() int {
	return xxx_messageInfo_RedNodeInfo.Size(m)
}
func (m *RedNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RedNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RedNodeInfo proto.InternalMessageInfo

func (m *RedNodeInfo) GetRedundancy() *RedSummaryPair {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RedNodeInfo) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *RedNodeInfo) GetActiveRebootReason() string {
	if m != nil {
		return m.ActiveRebootReason
	}
	return ""
}

func (m *RedNodeInfo) GetStandbyRebootReason() string {
	if m != nil {
		return m.StandbyRebootReason
	}
	return ""
}

func (m *RedNodeInfo) GetErrLog() string {
	if m != nil {
		return m.ErrLog
	}
	return ""
}

type RedGroupInfo struct {
	Active               string   `protobuf:"bytes,1,opt,name=active,proto3" json:"active,omitempty"`
	Standby              string   `protobuf:"bytes,2,opt,name=standby,proto3" json:"standby,omitempty"`
	HaState              string   `protobuf:"bytes,3,opt,name=ha_state,json=haState,proto3" json:"ha_state,omitempty"`
	NsrState             string   `protobuf:"bytes,4,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedGroupInfo) Reset()         { *m = RedGroupInfo{} }
func (m *RedGroupInfo) String() string { return proto.CompactTextString(m) }
func (*RedGroupInfo) ProtoMessage()    {}
func (*RedGroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8501aa16198dc9, []int{2}
}

func (m *RedGroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedGroupInfo.Unmarshal(m, b)
}
func (m *RedGroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedGroupInfo.Marshal(b, m, deterministic)
}
func (m *RedGroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedGroupInfo.Merge(m, src)
}
func (m *RedGroupInfo) XXX_Size() int {
	return xxx_messageInfo_RedGroupInfo.Size(m)
}
func (m *RedGroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RedGroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RedGroupInfo proto.InternalMessageInfo

func (m *RedGroupInfo) GetActive() string {
	if m != nil {
		return m.Active
	}
	return ""
}

func (m *RedGroupInfo) GetStandby() string {
	if m != nil {
		return m.Standby
	}
	return ""
}

func (m *RedGroupInfo) GetHaState() string {
	if m != nil {
		return m.HaState
	}
	return ""
}

func (m *RedGroupInfo) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

// rmf row show
type RedSummaryPair struct {
	// Active node name R/S/I
	Active string `protobuf:"bytes,1,opt,name=active,proto3" json:"active,omitempty"`
	// Standby node name R/S/I
	Standby string `protobuf:"bytes,2,opt,name=standby,proto3" json:"standby,omitempty"`
	// High Availability state Ready/Not Ready
	HaState string `protobuf:"bytes,3,opt,name=ha_state,json=haState,proto3" json:"ha_state,omitempty"`
	// NSR state Configured/Not Configured
	NsrState             string          `protobuf:"bytes,4,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	Groupinfo            []*RedGroupInfo `protobuf:"bytes,5,rep,name=groupinfo,proto3" json:"groupinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RedSummaryPair) Reset()         { *m = RedSummaryPair{} }
func (m *RedSummaryPair) String() string { return proto.CompactTextString(m) }
func (*RedSummaryPair) ProtoMessage()    {}
func (*RedSummaryPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8501aa16198dc9, []int{3}
}

func (m *RedSummaryPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedSummaryPair.Unmarshal(m, b)
}
func (m *RedSummaryPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedSummaryPair.Marshal(b, m, deterministic)
}
func (m *RedSummaryPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedSummaryPair.Merge(m, src)
}
func (m *RedSummaryPair) XXX_Size() int {
	return xxx_messageInfo_RedSummaryPair.Size(m)
}
func (m *RedSummaryPair) XXX_DiscardUnknown() {
	xxx_messageInfo_RedSummaryPair.DiscardUnknown(m)
}

var xxx_messageInfo_RedSummaryPair proto.InternalMessageInfo

func (m *RedSummaryPair) GetActive() string {
	if m != nil {
		return m.Active
	}
	return ""
}

func (m *RedSummaryPair) GetStandby() string {
	if m != nil {
		return m.Standby
	}
	return ""
}

func (m *RedSummaryPair) GetHaState() string {
	if m != nil {
		return m.HaState
	}
	return ""
}

func (m *RedSummaryPair) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func (m *RedSummaryPair) GetGroupinfo() []*RedGroupInfo {
	if m != nil {
		return m.Groupinfo
	}
	return nil
}

func init() {
	proto.RegisterType((*RedNodeInfo_KEYS)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.nodes.node.red_node_info_KEYS")
	proto.RegisterType((*RedNodeInfo)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.nodes.node.red_node_info")
	proto.RegisterType((*RedGroupInfo)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.nodes.node.red_group_info")
	proto.RegisterType((*RedSummaryPair)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.nodes.node.red_summary_pair")
}

func init() { proto.RegisterFile("red_node_info.proto", fileDescriptor_2f8501aa16198dc9) }

var fileDescriptor_2f8501aa16198dc9 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xf6, 0x7d, 0xfb, 0xe7, 0x2a, 0x50, 0xe5, 0x02, 0x35, 0x62, 0xa9, 0x32, 0x75,
	0x21, 0x82, 0x14, 0xd8, 0x11, 0x62, 0x40, 0x30, 0xa5, 0x13, 0xd3, 0xc9, 0x8d, 0xdd, 0x36, 0x12,
	0xb5, 0xa3, 0x73, 0x8a, 0x28, 0x3b, 0x1f, 0x93, 0xef, 0x82, 0x6c, 0xb7, 0x2a, 0xed, 0x06, 0x03,
	0x4b, 0x94, 0xf3, 0xcf, 0xcf, 0x73, 0x77, 0x4f, 0x02, 0x3d, 0x52, 0x12, 0xb5, 0x91, 0x0a, 0x0b,
	0x3d, 0x35, 0x49, 0x49, 0xa6, 0x32, 0xec, 0x32, 0x2f, 0x6c, 0x6e, 0xb0, 0x30, 0x16, 0xdf, 0xc8,
	0x01, 0x12, 0x48, 0x8b, 0x29, 0x9a, 0x52, 0x51, 0x42, 0x4a, 0x2e, 0xb5, 0x14, 0x3a, 0x5f, 0x25,
	0x4e, 0x67, 0xfd, 0x33, 0x3e, 0x07, 0xb6, 0xe3, 0x84, 0x8f, 0xf7, 0xcf, 0x63, 0xd6, 0x87, 0x66,
	0x38, 0x91, 0x3c, 0x1a, 0x44, 0xc3, 0x76, 0xd6, 0x70, 0xe5, 0x83, 0x8c, 0x3f, 0x6a, 0x70, 0xb0,
	0x73, 0x9f, 0xe5, 0x00, 0x5b, 0x67, 0x9e, 0x0e, 0xa2, 0x61, 0x27, 0xbd, 0x4b, 0x7e, 0x3c, 0x88,
	0x3b, 0x45, 0xbb, 0x5c, 0x2c, 0x04, 0xad, 0xb0, 0x14, 0x05, 0x65, 0xdf, 0x6c, 0x59, 0x17, 0xea,
	0x2f, 0x66, 0xc6, 0x47, 0x7e, 0x16, 0xf7, 0xca, 0x2e, 0xe0, 0x48, 0xe4, 0x55, 0xf1, 0xaa, 0x90,
	0xd4, 0xc4, 0x98, 0x0a, 0x49, 0x09, 0x6b, 0x34, 0xbf, 0xf2, 0x57, 0x58, 0x60, 0x99, 0x47, 0x99,
	0x27, 0x2c, 0x85, 0x63, 0x5b, 0x09, 0x2d, 0x27, 0xab, 0x3d, 0xc9, 0xb5, 0x97, 0xf4, 0xd6, 0x70,
	0x47, 0xd3, 0x87, 0xa6, 0x22, 0x42, 0xd7, 0xfb, 0x26, 0xe4, 0xa0, 0x88, 0x9e, 0xcc, 0x2c, 0x7e,
	0x87, 0x43, 0x37, 0xf0, 0x8c, 0xcc, 0xb2, 0x0c, 0x39, 0x9c, 0x40, 0x23, 0x34, 0xdd, 0x24, 0x16,
	0x2a, 0xc6, 0xa1, 0xb9, 0x76, 0xe6, 0x35, 0x0f, 0x36, 0x25, 0x3b, 0x85, 0xd6, 0x5c, 0xa0, 0xad,
	0x44, 0xa5, 0x78, 0x3d, 0xa0, 0xb9, 0x18, 0xbb, 0x92, 0x9d, 0x41, 0x5b, 0x5b, 0x5a, 0xb3, 0x7f,
	0x9e, 0xb5, 0xb4, 0x25, 0x0f, 0xe3, 0xcf, 0x08, 0xba, 0xfb, 0x69, 0xfd, 0x5d, 0x7b, 0x86, 0xd0,
	0xf6, 0x6b, 0xbb, 0xad, 0xf9, 0xff, 0x41, 0x7d, 0xd8, 0x49, 0x6f, 0x7f, 0xf9, 0xbd, 0xb7, 0xf1,
	0x65, 0x5b, 0xcf, 0x49, 0xc3, 0xff, 0xcc, 0xa3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xa5,
	0x38, 0x41, 0xe3, 0x02, 0x00, 0x00,
}
