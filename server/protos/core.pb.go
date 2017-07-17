// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/core.proto

/*
Package serving is a generated protocol buffer package.

It is generated from these files:
	protos/core.proto
	protos/evaluator.proto
	protos/features.proto
	protos/models.proto

It has these top-level messages:
	ValueInterface
	EmptyResponse
	EvaluationRequest
	EvaluationResponse
	FeatureSet
	FeatureSetSchema
	FeatureSetSchemaField
	ListFeatureSetsRequest
	FindFeatureSetRequest
	CreateFeatureSetRequest
	DeleteFeatureSetRequest
	ListFeatureSetSchemasRequest
	FindFeatureSetSchemaRequest
	CreateFeatureSetSchemaRequest
	DeleteFeatureSetSchemaRequest
	Model
	ModelVersion
	ModelFeature
	ListModelsRequest
	FindModelRequest
	CreateModelRequest
	DeleteModelRequest
	ListVersionsRequest
	FindVersionRequest
	SetPrimaryVersionRequest
	CreateVersionRequest
	PrecomputedFeaturesSet
	DeleteVersionRequest
*/
package serving

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

type ValueInterface struct {
	// Types that are valid to be assigned to Value:
	//	*ValueInterface_ValueBoolean
	//	*ValueInterface_ValueInt32
	//	*ValueInterface_ValueInt64
	//	*ValueInterface_ValueFloat32
	//	*ValueInterface_ValueFloat64
	//	*ValueInterface_ValueString
	//	*ValueInterface_ValueBytes
	Value isValueInterface_Value `protobuf_oneof:"value"`
}

func (m *ValueInterface) Reset()                    { *m = ValueInterface{} }
func (m *ValueInterface) String() string            { return proto.CompactTextString(m) }
func (*ValueInterface) ProtoMessage()               {}
func (*ValueInterface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isValueInterface_Value interface {
	isValueInterface_Value()
}

type ValueInterface_ValueBoolean struct {
	ValueBoolean bool `protobuf:"varint,1,opt,name=value_boolean,json=valueBoolean,oneof"`
}
type ValueInterface_ValueInt32 struct {
	ValueInt32 int32 `protobuf:"varint,2,opt,name=value_int32,json=valueInt32,oneof"`
}
type ValueInterface_ValueInt64 struct {
	ValueInt64 int64 `protobuf:"varint,3,opt,name=value_int64,json=valueInt64,oneof"`
}
type ValueInterface_ValueFloat32 struct {
	ValueFloat32 float32 `protobuf:"fixed32,4,opt,name=value_float32,json=valueFloat32,oneof"`
}
type ValueInterface_ValueFloat64 struct {
	ValueFloat64 float64 `protobuf:"fixed64,5,opt,name=value_float64,json=valueFloat64,oneof"`
}
type ValueInterface_ValueString struct {
	ValueString string `protobuf:"bytes,6,opt,name=value_string,json=valueString,oneof"`
}
type ValueInterface_ValueBytes struct {
	ValueBytes []byte `protobuf:"bytes,7,opt,name=value_bytes,json=valueBytes,proto3,oneof"`
}

func (*ValueInterface_ValueBoolean) isValueInterface_Value() {}
func (*ValueInterface_ValueInt32) isValueInterface_Value()   {}
func (*ValueInterface_ValueInt64) isValueInterface_Value()   {}
func (*ValueInterface_ValueFloat32) isValueInterface_Value() {}
func (*ValueInterface_ValueFloat64) isValueInterface_Value() {}
func (*ValueInterface_ValueString) isValueInterface_Value()  {}
func (*ValueInterface_ValueBytes) isValueInterface_Value()   {}

func (m *ValueInterface) GetValue() isValueInterface_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueInterface) GetValueBoolean() bool {
	if x, ok := m.GetValue().(*ValueInterface_ValueBoolean); ok {
		return x.ValueBoolean
	}
	return false
}

func (m *ValueInterface) GetValueInt32() int32 {
	if x, ok := m.GetValue().(*ValueInterface_ValueInt32); ok {
		return x.ValueInt32
	}
	return 0
}

func (m *ValueInterface) GetValueInt64() int64 {
	if x, ok := m.GetValue().(*ValueInterface_ValueInt64); ok {
		return x.ValueInt64
	}
	return 0
}

func (m *ValueInterface) GetValueFloat32() float32 {
	if x, ok := m.GetValue().(*ValueInterface_ValueFloat32); ok {
		return x.ValueFloat32
	}
	return 0
}

func (m *ValueInterface) GetValueFloat64() float64 {
	if x, ok := m.GetValue().(*ValueInterface_ValueFloat64); ok {
		return x.ValueFloat64
	}
	return 0
}

func (m *ValueInterface) GetValueString() string {
	if x, ok := m.GetValue().(*ValueInterface_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *ValueInterface) GetValueBytes() []byte {
	if x, ok := m.GetValue().(*ValueInterface_ValueBytes); ok {
		return x.ValueBytes
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ValueInterface) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ValueInterface_OneofMarshaler, _ValueInterface_OneofUnmarshaler, _ValueInterface_OneofSizer, []interface{}{
		(*ValueInterface_ValueBoolean)(nil),
		(*ValueInterface_ValueInt32)(nil),
		(*ValueInterface_ValueInt64)(nil),
		(*ValueInterface_ValueFloat32)(nil),
		(*ValueInterface_ValueFloat64)(nil),
		(*ValueInterface_ValueString)(nil),
		(*ValueInterface_ValueBytes)(nil),
	}
}

func _ValueInterface_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ValueInterface)
	// value
	switch x := m.Value.(type) {
	case *ValueInterface_ValueBoolean:
		t := uint64(0)
		if x.ValueBoolean {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ValueInterface_ValueInt32:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ValueInt32))
	case *ValueInterface_ValueInt64:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ValueInt64))
	case *ValueInterface_ValueFloat32:
		b.EncodeVarint(4<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.ValueFloat32)))
	case *ValueInterface_ValueFloat64:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.ValueFloat64))
	case *ValueInterface_ValueString:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ValueString)
	case *ValueInterface_ValueBytes:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ValueBytes)
	case nil:
	default:
		return fmt.Errorf("ValueInterface.Value has unexpected type %T", x)
	}
	return nil
}

func _ValueInterface_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ValueInterface)
	switch tag {
	case 1: // value.value_boolean
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ValueInterface_ValueBoolean{x != 0}
		return true, err
	case 2: // value.value_int32
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ValueInterface_ValueInt32{int32(x)}
		return true, err
	case 3: // value.value_int64
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ValueInterface_ValueInt64{int64(x)}
		return true, err
	case 4: // value.value_float32
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &ValueInterface_ValueFloat32{math.Float32frombits(uint32(x))}
		return true, err
	case 5: // value.value_float64
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &ValueInterface_ValueFloat64{math.Float64frombits(x)}
		return true, err
	case 6: // value.value_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &ValueInterface_ValueString{x}
		return true, err
	case 7: // value.value_bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &ValueInterface_ValueBytes{x}
		return true, err
	default:
		return false, nil
	}
}

func _ValueInterface_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ValueInterface)
	// value
	switch x := m.Value.(type) {
	case *ValueInterface_ValueBoolean:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *ValueInterface_ValueInt32:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ValueInt32))
	case *ValueInterface_ValueInt64:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ValueInt64))
	case *ValueInterface_ValueFloat32:
		n += proto.SizeVarint(4<<3 | proto.WireFixed32)
		n += 4
	case *ValueInterface_ValueFloat64:
		n += proto.SizeVarint(5<<3 | proto.WireFixed64)
		n += 8
	case *ValueInterface_ValueString:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ValueString)))
		n += len(x.ValueString)
	case *ValueInterface_ValueBytes:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ValueBytes)))
		n += len(x.ValueBytes)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ValueInterface)(nil), "serving.ValueInterface")
	proto.RegisterType((*EmptyResponse)(nil), "serving.EmptyResponse")
}

func init() { proto.RegisterFile("protos/core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x7b, 0x3a, 0xbb, 0xea, 0x71, 0x53, 0xec, 0x55, 0x2e, 0xe3, 0x44, 0xc8, 0x95, 0xc2,
	0x5a, 0xfa, 0x00, 0x05, 0xa5, 0xde, 0x56, 0xf0, 0x76, 0xa4, 0xe3, 0x6c, 0x14, 0x6a, 0x52, 0x92,
	0x58, 0xd8, 0xbb, 0xf8, 0xb0, 0xd2, 0x44, 0x36, 0xe7, 0x5d, 0xce, 0xc7, 0xc7, 0xff, 0xff, 0x04,
	0xef, 0x06, 0xa3, 0x9d, 0xb6, 0xcf, 0x5b, 0x6d, 0xe8, 0xc9, 0xbf, 0xb3, 0xd4, 0x92, 0x19, 0x3b,
	0xb5, 0x5f, 0x7d, 0xc7, 0x78, 0xf3, 0x21, 0xfb, 0x2f, 0x7a, 0x53, 0x8e, 0xcc, 0x4e, 0x6e, 0x29,
	0x7b, 0xc4, 0xe5, 0x38, 0x91, 0x4d, 0xab, 0x75, 0x4f, 0x52, 0x31, 0xe0, 0x20, 0x2e, 0xeb, 0xa8,
	0x59, 0x78, 0x5c, 0x05, 0x9a, 0xdd, 0xe3, 0x75, 0xd0, 0x3a, 0xe5, 0xf2, 0x35, 0x8b, 0x39, 0x88,
	0xa4, 0x8e, 0x1a, 0x1c, 0x7f, 0xd3, 0xf2, 0xf5, 0x99, 0x52, 0x16, 0x6c, 0xc6, 0x41, 0xcc, 0xfe,
	0x2a, 0x65, 0x71, 0x2a, 0xdb, 0xf5, 0x5a, 0x4e, 0x39, 0x17, 0x1c, 0x44, 0x7c, 0x2c, 0x7b, 0x0d,
	0xf4, 0x9f, 0x56, 0x16, 0x2c, 0xe1, 0x20, 0xe0, 0x5c, 0x2b, 0x8b, 0xec, 0x01, 0xc3, 0xbd, 0xb1,
	0xce, 0x74, 0x6a, 0xcf, 0xe6, 0x1c, 0xc4, 0x55, 0x1d, 0x35, 0x61, 0xc6, 0xbb, 0x87, 0xa7, 0x55,
	0xed, 0xc1, 0x91, 0x65, 0x29, 0x07, 0xb1, 0x38, 0xae, 0xaa, 0x26, 0x56, 0xa5, 0x98, 0xf8, 0x6b,
	0x75, 0x8b, 0xcb, 0x97, 0xcf, 0xc1, 0x1d, 0x1a, 0xb2, 0x83, 0x56, 0x96, 0xda, 0xb9, 0xff, 0xbf,
	0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x77, 0x8a, 0xfb, 0xf3, 0x54, 0x01, 0x00, 0x00,
}
