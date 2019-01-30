// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm_mgr_show_alarm_brief_info.proto

package cisco_ios_xr_alarmgr_server_oper_alarms_brief_brief_card_brief_locations_brief_location_history

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

// Alarm mgr show alarm brief info
type AlarmMgrShowAlarmBriefInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmBriefInfo_KEYS) Reset()         { *m = AlarmMgrShowAlarmBriefInfo_KEYS{} }
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefInfo_KEYS) ProtoMessage()    {}
func (*AlarmMgrShowAlarmBriefInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_brief_info_cea0ebe96079d962, []int{0}
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Merge(dst, src)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Size(m)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmBriefInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type AlarmMgrShowAlarmBriefInfo struct {
	// Alarm List
	AlarmInfo            []*AlarmMgrShowAlarmBriefData `protobuf:"bytes,50,rep,name=alarm_info,json=alarmInfo,proto3" json:"alarm_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AlarmMgrShowAlarmBriefInfo) Reset()         { *m = AlarmMgrShowAlarmBriefInfo{} }
func (m *AlarmMgrShowAlarmBriefInfo) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefInfo) ProtoMessage()    {}
func (*AlarmMgrShowAlarmBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_brief_info_cea0ebe96079d962, []int{1}
}
func (m *AlarmMgrShowAlarmBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Marshal(b, m, deterministic)
}
func (dst *AlarmMgrShowAlarmBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Merge(dst, src)
}
func (m *AlarmMgrShowAlarmBriefInfo) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Size(m)
}
func (m *AlarmMgrShowAlarmBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmBriefInfo proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmBriefInfo) GetAlarmInfo() []*AlarmMgrShowAlarmBriefData {
	if m != nil {
		return m.AlarmInfo
	}
	return nil
}

// Alarm mgr show alarm brief data
type AlarmMgrShowAlarmBriefData struct {
	// Alarm location
	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// Alarm severity
	Severity string `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	// Alarm group
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// Alarm set time
	SetTime string `protobuf:"bytes,4,opt,name=set_time,json=setTime,proto3" json:"set_time,omitempty"`
	// Alarm set time(timestamp format)
	SetTimestamp uint64 `protobuf:"varint,5,opt,name=set_timestamp,json=setTimestamp,proto3" json:"set_timestamp,omitempty"`
	// Alarm clear time
	ClearTime string `protobuf:"bytes,6,opt,name=clear_time,json=clearTime,proto3" json:"clear_time,omitempty"`
	// Alarm clear time(timestamp format)
	ClearTimestamp uint64 `protobuf:"varint,7,opt,name=clear_timestamp,json=clearTimestamp,proto3" json:"clear_timestamp,omitempty"`
	// Alarm description
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmBriefData) Reset()         { *m = AlarmMgrShowAlarmBriefData{} }
func (m *AlarmMgrShowAlarmBriefData) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefData) ProtoMessage()    {}
func (*AlarmMgrShowAlarmBriefData) Descriptor() ([]byte, []int) {
	return fileDescriptor_alarm_mgr_show_alarm_brief_info_cea0ebe96079d962, []int{2}
}
func (m *AlarmMgrShowAlarmBriefData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefData.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmBriefData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefData.Marshal(b, m, deterministic)
}
func (dst *AlarmMgrShowAlarmBriefData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmBriefData.Merge(dst, src)
}
func (m *AlarmMgrShowAlarmBriefData) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefData.Size(m)
}
func (m *AlarmMgrShowAlarmBriefData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmBriefData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmBriefData proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmBriefData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetSetTime() string {
	if m != nil {
		return m.SetTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetSetTimestamp() uint64 {
	if m != nil {
		return m.SetTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmBriefData) GetClearTime() string {
	if m != nil {
		return m.ClearTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetClearTimestamp() uint64 {
	if m != nil {
		return m.ClearTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmBriefData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*AlarmMgrShowAlarmBriefInfo_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_card.brief_locations.brief_location.history.alarm_mgr_show_alarm_brief_info_KEYS")
	proto.RegisterType((*AlarmMgrShowAlarmBriefInfo)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_card.brief_locations.brief_location.history.alarm_mgr_show_alarm_brief_info")
	proto.RegisterType((*AlarmMgrShowAlarmBriefData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_card.brief_locations.brief_location.history.alarm_mgr_show_alarm_brief_data")
}

func init() {
	proto.RegisterFile("alarm_mgr_show_alarm_brief_info.proto", fileDescriptor_alarm_mgr_show_alarm_brief_info_cea0ebe96079d962)
}

var fileDescriptor_alarm_mgr_show_alarm_brief_info_cea0ebe96079d962 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xfe, 0xef, 0xf5, 0x7d, 0x41, 0xb2, 0x90, 0x08, 0x48, 0x88, 0xa8, 0x80, 0xe8,
	0x94, 0xa1, 0x7c, 0x00, 0x26, 0x86, 0x8a, 0xad, 0xb0, 0x30, 0x59, 0x6e, 0xe2, 0xb6, 0x96, 0x9a,
	0x5c, 0x74, 0x36, 0x85, 0x6e, 0x7c, 0x0d, 0x36, 0x3e, 0x0c, 0x1f, 0x0c, 0xf9, 0xdc, 0x16, 0xc4,
	0xd0, 0x4e, 0x2c, 0x51, 0x9e, 0xe7, 0xb9, 0xfb, 0xf9, 0xce, 0x32, 0x5c, 0xa9, 0x85, 0xa2, 0x42,
	0x16, 0x33, 0x92, 0x76, 0x8e, 0x2f, 0x32, 0xc8, 0x09, 0x19, 0x3d, 0x95, 0xa6, 0x9c, 0x62, 0x5a,
	0x11, 0x3a, 0x14, 0x32, 0x33, 0x36, 0x43, 0x69, 0xd0, 0xca, 0x57, 0x0a, 0x45, 0xbe, 0x43, 0xd3,
	0x52, 0x93, 0xc4, 0x4a, 0x53, 0xca, 0x9e, 0x4d, 0xb9, 0x33, 0x7c, 0x65, 0xa6, 0x28, 0x5f, 0xff,
	0x2e, 0x30, 0x53, 0xce, 0x60, 0x69, 0x7f, 0xe9, 0x74, 0x6e, 0xac, 0x43, 0x5a, 0xf5, 0x6f, 0xe1,
	0x72, 0xcf, 0x24, 0xf2, 0xfe, 0xee, 0xe9, 0x41, 0x1c, 0x43, 0xbb, 0xc4, 0x5c, 0x4b, 0x93, 0xc7,
	0x51, 0x12, 0x0d, 0xba, 0xe3, 0x96, 0x97, 0xa3, 0xbc, 0xff, 0x19, 0xc1, 0xf9, 0x1e, 0x82, 0xf8,
	0x88, 0x00, 0x82, 0xe9, 0x65, 0x3c, 0x4c, 0xea, 0x83, 0xde, 0xf0, 0x2d, 0x4a, 0xff, 0x78, 0xb9,
	0x74, 0xc7, 0x5c, 0xb9, 0x72, 0x6a, 0xdc, 0x65, 0x67, 0x54, 0x4e, 0xb1, 0xff, 0x5e, 0xdb, 0xb9,
	0x86, 0x2f, 0x17, 0xa7, 0xd0, 0xd9, 0x1c, 0xb1, 0xbe, 0x84, 0xad, 0xf6, 0x99, 0xd5, 0x4b, 0x4d,
	0xc6, 0xad, 0xe2, 0x5a, 0xc8, 0x36, 0x5a, 0x1c, 0x41, 0x73, 0x46, 0xf8, 0x5c, 0xc5, 0x75, 0x0e,
	0x82, 0x10, 0x27, 0xbe, 0xc3, 0x49, 0x67, 0x0a, 0x1d, 0x37, 0x38, 0x68, 0x5b, 0xed, 0x1e, 0x4d,
	0xa1, 0xc5, 0x05, 0xfc, 0xdf, 0x44, 0xd6, 0xa9, 0xa2, 0x8a, 0x9b, 0x49, 0x34, 0x68, 0x8c, 0xff,
	0xad, 0x73, 0xf6, 0xc4, 0x19, 0x40, 0xb6, 0xd0, 0x8a, 0x02, 0xa1, 0xc5, 0x84, 0x2e, 0x3b, 0xcc,
	0xb8, 0x86, 0xc3, 0xef, 0x38, 0x50, 0xda, 0x4c, 0x39, 0xd8, 0xd6, 0x04, 0x4e, 0x02, 0xbd, 0x5c,
	0xdb, 0x8c, 0x4c, 0xc5, 0x8b, 0x75, 0x18, 0xf4, 0xd3, 0x9a, 0xb4, 0xf8, 0x2d, 0xde, 0x7c, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0x8d, 0xf0, 0x8d, 0xb4, 0x02, 0x00, 0x00,
}
