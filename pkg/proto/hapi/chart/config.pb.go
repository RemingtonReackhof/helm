// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/chart/config.proto

package chart

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConfType int32

const (
	ConfType_VALUES  ConfType = 0
	ConfType_SECRETS ConfType = 1
)

var ConfType_name = map[int32]string{
	0: "VALUES",
	1: "SECRETS",
}
var ConfType_value = map[string]int32{
	"VALUES":  0,
	"SECRETS": 1,
}

func (x ConfType) String() string {
	return proto.EnumName(ConfType_name, int32(x))
}
func (ConfType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Config supplies values to the parametrizable templates of a chart.
type Config struct {
	Raw    string            `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
	Values map[string]*Value `protobuf:"bytes,2,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type   ConfType          `protobuf:"varint,3,opt,name=type,enum=hapi.chart.ConfType" json:"type,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Config) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

func (m *Config) GetValues() map[string]*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Config) GetType() ConfType {
	if m != nil {
		return m.Type
	}
	return ConfType_VALUES
}

// Value describes a configuration value as a string.
type Value struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "hapi.chart.Config")
	proto.RegisterType((*Value)(nil), "hapi.chart.Value")
	proto.RegisterEnum("hapi.chart.ConfType", ConfType_name, ConfType_value)
}

func init() { proto.RegisterFile("hapi/chart/config.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x48, 0x2c, 0xc8,
	0xd4, 0x4f, 0xce, 0x48, 0x2c, 0x2a, 0xd1, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x49, 0xe8, 0x81, 0x25, 0x94, 0x4e, 0x31, 0x72, 0xb1, 0x39,
	0x83, 0x25, 0x85, 0x04, 0xb8, 0x98, 0x8b, 0x12, 0xcb, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x40, 0x4c, 0x21, 0x33, 0x2e, 0xb6, 0xb2, 0xc4, 0x9c, 0xd2, 0xd4, 0x62, 0x09, 0x26, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x39, 0x3d, 0x84, 0x4e, 0x3d, 0x88, 0x2e, 0xbd, 0x30, 0xb0, 0x02, 0xd7, 0xbc,
	0x92, 0xa2, 0xca, 0x20, 0xa8, 0x6a, 0x21, 0x0d, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x66,
	0x05, 0x46, 0x0d, 0x3e, 0x23, 0x11, 0x64, 0x5d, 0x20, 0x87, 0x84, 0x54, 0x16, 0xa4, 0x06, 0x81,
	0x55, 0x48, 0xf9, 0x70, 0x71, 0x23, 0x19, 0x00, 0x72, 0x42, 0x76, 0x6a, 0x25, 0xcc, 0x09, 0xd9,
	0xa9, 0x95, 0x42, 0xea, 0x5c, 0xac, 0x60, 0x43, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x04,
	0x91, 0xcd, 0x02, 0xeb, 0x0c, 0x82, 0xc8, 0x5b, 0x31, 0x59, 0x30, 0x2a, 0xc9, 0x72, 0xb1, 0x82,
	0xc5, 0x84, 0x44, 0x60, 0xba, 0x20, 0x26, 0x41, 0x38, 0x5a, 0xca, 0x5c, 0x1c, 0x30, 0xeb, 0x85,
	0xb8, 0xb8, 0xd8, 0xc2, 0x1c, 0x7d, 0x42, 0x5d, 0x83, 0x05, 0x18, 0x84, 0xb8, 0xb9, 0xd8, 0x83,
	0x5d, 0x9d, 0x83, 0x5c, 0x43, 0x82, 0x05, 0x18, 0x9d, 0xd8, 0xa3, 0x58, 0xc1, 0xa6, 0x27, 0xb1,
	0x81, 0x03, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x99, 0xd8, 0x60, 0x3b, 0x47, 0x01, 0x00,
	0x00,
}
