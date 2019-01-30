// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform_info.proto

package cisco_ios_xr_plat_chas_invmgr_oper_platform_racks_rack_slots_slot_state

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

// Platform information
type PlatformInfo_KEYS struct {
	RackName             string   `protobuf:"bytes,1,opt,name=rack_name,json=rackName,proto3" json:"rack_name,omitempty"`
	SlotName             string   `protobuf:"bytes,2,opt,name=slot_name,json=slotName,proto3" json:"slot_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformInfo_KEYS) Reset()         { *m = PlatformInfo_KEYS{} }
func (m *PlatformInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PlatformInfo_KEYS) ProtoMessage()    {}
func (*PlatformInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_83c645cda97a4b41, []int{0}
}

func (m *PlatformInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformInfo_KEYS.Unmarshal(m, b)
}
func (m *PlatformInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PlatformInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformInfo_KEYS.Merge(m, src)
}
func (m *PlatformInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PlatformInfo_KEYS.Size(m)
}
func (m *PlatformInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformInfo_KEYS proto.InternalMessageInfo

func (m *PlatformInfo_KEYS) GetRackName() string {
	if m != nil {
		return m.RackName
	}
	return ""
}

func (m *PlatformInfo_KEYS) GetSlotName() string {
	if m != nil {
		return m.SlotName
	}
	return ""
}

type PlatformInfo struct {
	// Card type
	CardType string `protobuf:"bytes,50,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	// Redundancy state
	CardRedundancyState string `protobuf:"bytes,51,opt,name=card_redundancy_state,json=cardRedundancyState,proto3" json:"card_redundancy_state,omitempty"`
	// PLIM
	Plim string `protobuf:"bytes,52,opt,name=plim,proto3" json:"plim,omitempty"`
	// State
	State string `protobuf:"bytes,53,opt,name=state,proto3" json:"state,omitempty"`
	// True if power state is active
	IsMonitored bool `protobuf:"varint,54,opt,name=is_monitored,json=isMonitored,proto3" json:"is_monitored,omitempty"`
	// True if monitor state is active
	IsPowered bool `protobuf:"varint,55,opt,name=is_powered,json=isPowered,proto3" json:"is_powered,omitempty"`
	// True if shutdown state is active
	IsShutdown bool `protobuf:"varint,56,opt,name=is_shutdown,json=isShutdown,proto3" json:"is_shutdown,omitempty"`
	// Admin state
	AdminState           string   `protobuf:"bytes,57,opt,name=admin_state,json=adminState,proto3" json:"admin_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformInfo) Reset()         { *m = PlatformInfo{} }
func (m *PlatformInfo) String() string { return proto.CompactTextString(m) }
func (*PlatformInfo) ProtoMessage()    {}
func (*PlatformInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_83c645cda97a4b41, []int{1}
}

func (m *PlatformInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformInfo.Unmarshal(m, b)
}
func (m *PlatformInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformInfo.Marshal(b, m, deterministic)
}
func (m *PlatformInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformInfo.Merge(m, src)
}
func (m *PlatformInfo) XXX_Size() int {
	return xxx_messageInfo_PlatformInfo.Size(m)
}
func (m *PlatformInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformInfo proto.InternalMessageInfo

func (m *PlatformInfo) GetCardType() string {
	if m != nil {
		return m.CardType
	}
	return ""
}

func (m *PlatformInfo) GetCardRedundancyState() string {
	if m != nil {
		return m.CardRedundancyState
	}
	return ""
}

func (m *PlatformInfo) GetPlim() string {
	if m != nil {
		return m.Plim
	}
	return ""
}

func (m *PlatformInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *PlatformInfo) GetIsMonitored() bool {
	if m != nil {
		return m.IsMonitored
	}
	return false
}

func (m *PlatformInfo) GetIsPowered() bool {
	if m != nil {
		return m.IsPowered
	}
	return false
}

func (m *PlatformInfo) GetIsShutdown() bool {
	if m != nil {
		return m.IsShutdown
	}
	return false
}

func (m *PlatformInfo) GetAdminState() string {
	if m != nil {
		return m.AdminState
	}
	return ""
}

func init() {
	proto.RegisterType((*PlatformInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform.racks.rack.slots.slot.state.platform_info_KEYS")
	proto.RegisterType((*PlatformInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform.racks.rack.slots.slot.state.platform_info")
}

func init() { proto.RegisterFile("platform_info.proto", fileDescriptor_83c645cda97a4b41) }

var fileDescriptor_83c645cda97a4b41 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcb, 0x4e, 0x2a, 0x41,
	0x10, 0x40, 0x03, 0xb9, 0xf7, 0x06, 0x9a, 0xeb, 0xa6, 0xd1, 0x64, 0x12, 0x63, 0x44, 0x56, 0xac,
	0x66, 0x01, 0x3e, 0x3f, 0xc0, 0xb8, 0x30, 0x12, 0x03, 0x6e, 0x5c, 0x55, 0xda, 0x99, 0x46, 0x2a,
	0xd2, 0x8f, 0x74, 0x35, 0x22, 0x3f, 0xe1, 0x37, 0x9b, 0xaa, 0x01, 0x13, 0x36, 0x9d, 0x99, 0x73,
	0x4e, 0xd7, 0x54, 0x32, 0xaa, 0x1f, 0x57, 0x26, 0x2f, 0x42, 0x72, 0x80, 0x7e, 0x11, 0xca, 0x98,
	0x42, 0x0e, 0xfa, 0xa1, 0x42, 0xaa, 0x02, 0x60, 0x20, 0xf8, 0x4a, 0xc0, 0x05, 0x54, 0x4b, 0x43,
	0x80, 0xfe, 0xd3, 0xbd, 0x27, 0x08, 0xd1, 0xa6, 0x72, 0x7f, 0xaf, 0x4c, 0xa6, 0xfa, 0x20, 0x39,
	0x4b, 0x5a, 0x85, 0x4c, 0x72, 0x96, 0x94, 0x4d, 0xb6, 0xc3, 0xa9, 0xd2, 0x07, 0xf3, 0xe1, 0xf1,
	0xfe, 0x75, 0xae, 0x4f, 0x55, 0x97, 0x73, 0xf0, 0xc6, 0xd9, 0xa2, 0x35, 0x68, 0x8d, 0xba, 0xb3,
	0x0e, 0x83, 0xa9, 0x71, 0x96, 0x25, 0x0f, 0x68, 0x64, 0xbb, 0x91, 0x0c, 0x58, 0x0e, 0xbf, 0xdb,
	0xea, 0xe8, 0x60, 0x20, 0xe7, 0x95, 0x49, 0x35, 0xe4, 0x6d, 0xb4, 0xc5, 0xb8, 0xc9, 0x19, 0xbc,
	0x6c, 0xa3, 0xd5, 0x63, 0x75, 0x22, 0x32, 0xd9, 0x7a, 0xed, 0x6b, 0xe3, 0xab, 0x2d, 0xc8, 0x5e,
	0xc5, 0x44, 0xc2, 0x3e, 0xcb, 0xd9, 0xaf, 0x9b, 0xb3, 0xd2, 0x5a, 0xfd, 0x89, 0x2b, 0x74, 0xc5,
	0xa5, 0x24, 0xf2, 0xac, 0x8f, 0xd5, 0xdf, 0xe6, 0xde, 0x95, 0xc0, 0xe6, 0x45, 0x5f, 0xa8, 0xff,
	0x48, 0xe0, 0x82, 0xc7, 0x1c, 0x92, 0xad, 0x8b, 0xeb, 0x41, 0x6b, 0xd4, 0x99, 0xf5, 0x90, 0x9e,
	0xf6, 0x48, 0x9f, 0x29, 0x85, 0x04, 0x31, 0x6c, 0x2c, 0x07, 0x37, 0x12, 0x74, 0x91, 0x9e, 0x1b,
	0xa0, 0xcf, 0x55, 0x0f, 0x09, 0x68, 0xb9, 0xce, 0x75, 0xd8, 0xf8, 0xe2, 0x56, 0xbc, 0x42, 0x9a,
	0xef, 0x08, 0x07, 0xa6, 0x76, 0xe8, 0x77, 0x6b, 0xdf, 0xc9, 0xe7, 0x95, 0x20, 0xd9, 0xf6, 0xed,
	0x9f, 0xfc, 0xb0, 0xc9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x61, 0x6a, 0x09, 0xc7, 0x01,
	0x00, 0x00,
}
