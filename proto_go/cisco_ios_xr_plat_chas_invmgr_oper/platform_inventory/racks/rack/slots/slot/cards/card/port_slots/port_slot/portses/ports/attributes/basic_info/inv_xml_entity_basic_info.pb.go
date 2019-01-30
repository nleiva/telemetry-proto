// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inv_xml_entity_basic_info.proto

/*
Package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_portses_ports_attributes_basic_info is a generated protocol buffer package.

It is generated from these files:
	inv_xml_entity_basic_info.proto

It has these top-level messages:
	InvXmlEntityBasicInfo_KEYS
	InvXmlEntityBasicInfo
*/
package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_slots_slot_cards_card_portses_ports_attributes_basic_info

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

// Bag contains all the basic inventory information for each entity
type InvXmlEntityBasicInfo_KEYS struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Name_1 string `protobuf:"bytes,2,opt,name=name_1,json=name1" json:"name_1,omitempty"`
	Name_2 string `protobuf:"bytes,3,opt,name=name_2,json=name2" json:"name_2,omitempty"`
	Name_3 string `protobuf:"bytes,4,opt,name=name_3,json=name3" json:"name_3,omitempty"`
}

func (m *InvXmlEntityBasicInfo_KEYS) Reset()                    { *m = InvXmlEntityBasicInfo_KEYS{} }
func (m *InvXmlEntityBasicInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*InvXmlEntityBasicInfo_KEYS) ProtoMessage()               {}
func (*InvXmlEntityBasicInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InvXmlEntityBasicInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *InvXmlEntityBasicInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

type InvXmlEntityBasicInfo struct {
	// name string for the entity
	Name string `protobuf:"bytes,50,opt,name=name" json:"name,omitempty"`
	// describes in user-readable terms what the entity in question does
	Description string `protobuf:"bytes,51,opt,name=description" json:"description,omitempty"`
	// model name
	ModelName string `protobuf:"bytes,52,opt,name=model_name,json=modelName" json:"model_name,omitempty"`
	// hw revision string
	HardwareRevision string `protobuf:"bytes,53,opt,name=hardware_revision,json=hardwareRevision" json:"hardware_revision,omitempty"`
	// serial number
	SerialNumber string `protobuf:"bytes,54,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	// firmware revision string
	FirmwareRevision string `protobuf:"bytes,55,opt,name=firmware_revision,json=firmwareRevision" json:"firmware_revision,omitempty"`
	// software revision string
	SoftwareRevision string `protobuf:"bytes,56,opt,name=software_revision,json=softwareRevision" json:"software_revision,omitempty"`
	// maps to the vendor OID string
	VendorType string `protobuf:"bytes,57,opt,name=vendor_type,json=vendorType" json:"vendor_type,omitempty"`
	// 1 if Field Replaceable Unit 0, if not
	IsFieldReplaceableUnit bool `protobuf:"varint,58,opt,name=is_field_replaceable_unit,json=isFieldReplaceableUnit" json:"is_field_replaceable_unit,omitempty"`
}

func (m *InvXmlEntityBasicInfo) Reset()                    { *m = InvXmlEntityBasicInfo{} }
func (m *InvXmlEntityBasicInfo) String() string            { return proto.CompactTextString(m) }
func (*InvXmlEntityBasicInfo) ProtoMessage()               {}
func (*InvXmlEntityBasicInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InvXmlEntityBasicInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetHardwareRevision() string {
	if m != nil {
		return m.HardwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetFirmwareRevision() string {
	if m != nil {
		return m.FirmwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetSoftwareRevision() string {
	if m != nil {
		return m.SoftwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetVendorType() string {
	if m != nil {
		return m.VendorType
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetIsFieldReplaceableUnit() bool {
	if m != nil {
		return m.IsFieldReplaceableUnit
	}
	return false
}

func init() {
	proto.RegisterType((*InvXmlEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.portses.ports.attributes.basic_info.inv_xml_entity_basic_info_KEYS")
	proto.RegisterType((*InvXmlEntityBasicInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.slots.slot.cards.card.portses.ports.attributes.basic_info.inv_xml_entity_basic_info")
}

func init() { proto.RegisterFile("inv_xml_entity_basic_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0x14, 0x31,
	0x10, 0xc5, 0x75, 0x21, 0x44, 0xc4, 0x01, 0x89, 0xac, 0x04, 0xda, 0x14, 0x90, 0x53, 0x68, 0x22,
	0x45, 0x5a, 0x29, 0x77, 0xfc, 0x0b, 0x3d, 0x34, 0x48, 0x29, 0x0e, 0x28, 0xa8, 0x46, 0x5e, 0xef,
	0x2c, 0x19, 0xe1, 0xb5, 0xad, 0x19, 0xdf, 0x91, 0xe5, 0x93, 0x53, 0x22, 0x7b, 0x73, 0xbb, 0x49,
	0x71, 0xcd, 0x8c, 0xfd, 0x7b, 0xcf, 0x6f, 0xa6, 0xb0, 0x3a, 0x25, 0xb7, 0x81, 0xdb, 0xce, 0x02,
	0xba, 0x48, 0xb1, 0x87, 0x5a, 0x0b, 0x19, 0x20, 0xd7, 0xfa, 0x2a, 0xb0, 0x8f, 0xbe, 0xf8, 0x6b,
	0x48, 0x8c, 0x07, 0xf2, 0x02, 0xb7, 0x0c, 0xc1, 0xea, 0x08, 0xe6, 0x46, 0x0b, 0x90, 0xdb, 0x74,
	0xbf, 0x18, 0x7c, 0x40, 0xae, 0x12, 0x6d, 0x3d, 0x77, 0x09, 0xa2, 0x8b, 0x9e, 0xfb, 0x8a, 0xb5,
	0xf9, 0x2d, 0xb9, 0x56, 0x62, 0x7d, 0x94, 0x5c, 0x2b, 0xa3, 0xb9, 0x91, 0x5c, 0xab, 0xe0, 0x39,
	0x0a, 0xca, 0xd0, 0x2b, 0x1d, 0x23, 0x53, 0xbd, 0x8e, 0x28, 0xd5, 0xb4, 0xc1, 0x59, 0xaf, 0x5e,
	0xef, 0x5c, 0x0f, 0xbe, 0x7e, 0xfe, 0xf9, 0xad, 0x28, 0xd4, 0xbe, 0xd3, 0x1d, 0x96, 0xb3, 0xf9,
	0xec, 0xfc, 0x70, 0x95, 0xcf, 0xc5, 0x0b, 0x75, 0x90, 0x3a, 0x5c, 0x96, 0x7b, 0x99, 0x3e, 0x4e,
	0xb7, 0xcb, 0x11, 0x2f, 0xca, 0x47, 0x13, 0x5e, 0x8c, 0x78, 0x59, 0xee, 0x4f, 0x78, 0x79, 0xf6,
	0x6f, 0x4f, 0x9d, 0xec, 0x9c, 0x3d, 0x8e, 0x5d, 0xdc, 0x1b, 0x3b, 0x57, 0x47, 0x0d, 0x8a, 0x61,
	0x0a, 0x91, 0xbc, 0x2b, 0x97, 0x59, 0xba, 0x8f, 0x8a, 0x57, 0x4a, 0x75, 0xbe, 0x41, 0x0b, 0xf9,
	0xed, 0xdb, 0x6c, 0x38, 0xcc, 0xe4, 0x3a, 0x05, 0x5c, 0xa8, 0xe3, 0x1b, 0xcd, 0xcd, 0x1f, 0xcd,
	0x08, 0x8c, 0x1b, 0x92, 0x14, 0xf3, 0x2e, 0xbb, 0x9e, 0x6f, 0x85, 0xd5, 0x1d, 0x2f, 0xde, 0xa8,
	0x67, 0x82, 0x4c, 0xda, 0x82, 0x5b, 0x77, 0x35, 0x72, 0xf9, 0x3e, 0x1b, 0x9f, 0x0e, 0xf0, 0x3a,
	0xb3, 0x94, 0xd8, 0x12, 0x77, 0x0f, 0x13, 0x3f, 0x0c, 0x89, 0x5b, 0x61, 0x4c, 0xbc, 0x50, 0xc7,
	0xe2, 0xdb, 0xf8, 0xd0, 0xfc, 0x71, 0x30, 0x6f, 0x85, 0xd1, 0x7c, 0xaa, 0x8e, 0x36, 0xe8, 0x1a,
	0xcf, 0x10, 0xfb, 0x80, 0xe5, 0x55, 0xb6, 0xa9, 0x01, 0x7d, 0xef, 0x03, 0x16, 0x57, 0xea, 0x84,
	0x04, 0x5a, 0x42, 0xdb, 0x00, 0x63, 0xb0, 0xda, 0xa0, 0xae, 0x2d, 0xc2, 0xda, 0x51, 0x2c, 0x3f,
	0xcd, 0x67, 0xe7, 0x4f, 0x56, 0x2f, 0x49, 0xbe, 0x24, 0x7d, 0x35, 0xc9, 0x3f, 0x1c, 0xc5, 0xfa,
	0x20, 0x7f, 0xbc, 0xe5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x3c, 0xf1, 0x54, 0x9b, 0x02,
	0x00, 0x00,
}
