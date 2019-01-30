// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm_mgr_show_alarm_brief_info.proto

package cisco_ios_xr_alarmgr_server_oper_alarms_brief_brief_system_active

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

// Alarm mgr show alarm brief info
type AlarmMgrShowAlarmBriefInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmBriefInfo_KEYS) Reset()         { *m = AlarmMgrShowAlarmBriefInfo_KEYS{} }
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefInfo_KEYS) ProtoMessage()    {}
func (*AlarmMgrShowAlarmBriefInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0631eb21234d0e7, []int{0}
}

func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Merge(m, src)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.Size(m)
}
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmBriefInfo_KEYS proto.InternalMessageInfo

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
	return fileDescriptor_a0631eb21234d0e7, []int{1}
}

func (m *AlarmMgrShowAlarmBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmBriefInfo.Merge(m, src)
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
	return fileDescriptor_a0631eb21234d0e7, []int{2}
}

func (m *AlarmMgrShowAlarmBriefData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefData.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmBriefData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmBriefData.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmBriefData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmBriefData.Merge(m, src)
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
	proto.RegisterType((*AlarmMgrShowAlarmBriefInfo_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_system.active.alarm_mgr_show_alarm_brief_info_KEYS")
	proto.RegisterType((*AlarmMgrShowAlarmBriefInfo)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_system.active.alarm_mgr_show_alarm_brief_info")
	proto.RegisterType((*AlarmMgrShowAlarmBriefData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_system.active.alarm_mgr_show_alarm_brief_data")
}

func init() {
	proto.RegisterFile("alarm_mgr_show_alarm_brief_info.proto", fileDescriptor_a0631eb21234d0e7)
}

var fileDescriptor_a0631eb21234d0e7 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3b, 0x4e, 0x03, 0x31,
	0x10, 0x86, 0xb5, 0x79, 0x67, 0xc2, 0x43, 0xb2, 0x28, 0x16, 0x24, 0xc4, 0x2a, 0xbc, 0x52, 0xb9,
	0x08, 0x27, 0xa0, 0xa0, 0x40, 0x74, 0x81, 0x86, 0xca, 0x72, 0x9c, 0x49, 0xb0, 0x14, 0xc7, 0xab,
	0xb1, 0x09, 0xa4, 0xe4, 0x18, 0x1c, 0x80, 0x7b, 0xa2, 0x1d, 0x93, 0x40, 0x95, 0x14, 0x34, 0x2b,
	0xfd, 0x8f, 0xf9, 0x66, 0x56, 0x32, 0x5c, 0xea, 0xb9, 0x26, 0xa7, 0xdc, 0x8c, 0x54, 0x78, 0xf1,
	0x6f, 0x2a, 0xc9, 0x31, 0x59, 0x9c, 0x2a, 0xbb, 0x98, 0x7a, 0x59, 0x92, 0x8f, 0x5e, 0xdc, 0x1a,
	0x1b, 0x8c, 0x57, 0xd6, 0x07, 0xf5, 0x4e, 0xa9, 0x54, 0x4d, 0x20, 0x2d, 0x91, 0x94, 0x2f, 0x91,
	0x24, 0x7b, 0x41, 0xf2, 0x64, 0xfa, 0xaa, 0xb0, 0x0a, 0x11, 0x9d, 0xd4, 0x26, 0xda, 0x25, 0xf6,
	0xaf, 0xe0, 0x62, 0xc7, 0x2e, 0xf5, 0x70, 0xf7, 0xfc, 0xd8, 0xff, 0xca, 0xe0, 0x6c, 0x47, 0x51,
	0x7c, 0x64, 0x00, 0xc9, 0xac, 0x64, 0x3e, 0x2c, 0xea, 0x83, 0xde, 0x70, 0x2c, 0xff, 0x7d, 0xa4,
	0xdc, 0xb2, 0x78, 0xa2, 0xa3, 0x1e, 0x75, 0xd9, 0xb9, 0x5f, 0x4c, 0x7d, 0xff, 0xb3, 0xb6, 0xf5,
	0xce, 0xaa, 0x2e, 0x4e, 0xa0, 0x33, 0xf7, 0x46, 0x47, 0xeb, 0x17, 0x79, 0x56, 0x64, 0x83, 0xee,
	0x68, 0xa3, 0xab, 0x2c, 0xe0, 0x12, 0xc9, 0xc6, 0x55, 0x5e, 0x4b, 0xd9, 0x5a, 0x8b, 0x23, 0x68,
	0xce, 0xc8, 0xbf, 0x96, 0x79, 0x9d, 0x83, 0x24, 0xc4, 0x71, 0x35, 0x11, 0x55, 0xb4, 0x0e, 0xf3,
	0x06, 0x07, 0xed, 0x80, 0xf1, 0xc9, 0x3a, 0x14, 0xe7, 0xb0, 0xbf, 0x8e, 0x42, 0xd4, 0xae, 0xcc,
	0x9b, 0x45, 0x36, 0x68, 0x8c, 0xf6, 0x7e, 0x72, 0xf6, 0xc4, 0x29, 0x80, 0x99, 0xa3, 0xa6, 0x44,
	0x68, 0x31, 0xa1, 0xcb, 0x0e, 0x33, 0xae, 0xe1, 0xf0, 0x37, 0x4e, 0x94, 0x36, 0x53, 0x0e, 0x36,
	0x9d, 0xc4, 0x29, 0xa0, 0x37, 0xc1, 0x60, 0xc8, 0x96, 0xfc, 0x63, 0x1d, 0x06, 0xfd, 0xb5, 0xc6,
	0x2d, 0x7e, 0x35, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe5, 0x77, 0xb6, 0x5e, 0x02,
	0x00, 0x00,
}
