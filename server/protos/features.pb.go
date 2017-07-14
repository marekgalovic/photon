// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/features.proto

package serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FeatureSet struct {
	Uid       string   `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Keys      []string `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	CreatedAt int32    `protobuf:"varint,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt int32    `protobuf:"varint,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *FeatureSet) Reset()                    { *m = FeatureSet{} }
func (m *FeatureSet) String() string            { return proto.CompactTextString(m) }
func (*FeatureSet) ProtoMessage()               {}
func (*FeatureSet) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *FeatureSet) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FeatureSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSet) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *FeatureSet) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *FeatureSet) GetUpdatedAt() int32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type FeatureSetSchema struct {
	Uid       string                   `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	CreatedAt int32                    `protobuf:"varint,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Fields    []*FeatureSetSchemaField `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
}

func (m *FeatureSetSchema) Reset()                    { *m = FeatureSetSchema{} }
func (m *FeatureSetSchema) String() string            { return proto.CompactTextString(m) }
func (*FeatureSetSchema) ProtoMessage()               {}
func (*FeatureSetSchema) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *FeatureSetSchema) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FeatureSetSchema) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *FeatureSetSchema) GetFields() []*FeatureSetSchemaField {
	if m != nil {
		return m.Fields
	}
	return nil
}

type FeatureSetSchemaField struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ValueType string `protobuf:"bytes,2,opt,name=value_type,json=valueType" json:"value_type,omitempty"`
}

func (m *FeatureSetSchemaField) Reset()                    { *m = FeatureSetSchemaField{} }
func (m *FeatureSetSchemaField) String() string            { return proto.CompactTextString(m) }
func (*FeatureSetSchemaField) ProtoMessage()               {}
func (*FeatureSetSchemaField) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *FeatureSetSchemaField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSetSchemaField) GetValueType() string {
	if m != nil {
		return m.ValueType
	}
	return ""
}

type ListFeatureSetsRequest struct {
}

func (m *ListFeatureSetsRequest) Reset()                    { *m = ListFeatureSetsRequest{} }
func (m *ListFeatureSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFeatureSetsRequest) ProtoMessage()               {}
func (*ListFeatureSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type FindFeatureSetRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *FindFeatureSetRequest) Reset()                    { *m = FindFeatureSetRequest{} }
func (m *FindFeatureSetRequest) String() string            { return proto.CompactTextString(m) }
func (*FindFeatureSetRequest) ProtoMessage()               {}
func (*FindFeatureSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *FindFeatureSetRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type CreateFeatureSetRequest struct {
	Name string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Keys []string `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
}

func (m *CreateFeatureSetRequest) Reset()                    { *m = CreateFeatureSetRequest{} }
func (m *CreateFeatureSetRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFeatureSetRequest) ProtoMessage()               {}
func (*CreateFeatureSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CreateFeatureSetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateFeatureSetRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DeleteFeatureSetRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *DeleteFeatureSetRequest) Reset()                    { *m = DeleteFeatureSetRequest{} }
func (m *DeleteFeatureSetRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteFeatureSetRequest) ProtoMessage()               {}
func (*DeleteFeatureSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *DeleteFeatureSetRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ListFeatureSetSchemasRequest struct {
	FeatureSetUid string `protobuf:"bytes,1,opt,name=feature_set_uid,json=featureSetUid" json:"feature_set_uid,omitempty"`
}

func (m *ListFeatureSetSchemasRequest) Reset()                    { *m = ListFeatureSetSchemasRequest{} }
func (m *ListFeatureSetSchemasRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFeatureSetSchemasRequest) ProtoMessage()               {}
func (*ListFeatureSetSchemasRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ListFeatureSetSchemasRequest) GetFeatureSetUid() string {
	if m != nil {
		return m.FeatureSetUid
	}
	return ""
}

type FindFeatureSetSchemaRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *FindFeatureSetSchemaRequest) Reset()                    { *m = FindFeatureSetSchemaRequest{} }
func (m *FindFeatureSetSchemaRequest) String() string            { return proto.CompactTextString(m) }
func (*FindFeatureSetSchemaRequest) ProtoMessage()               {}
func (*FindFeatureSetSchemaRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *FindFeatureSetSchemaRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type CreateFeatureSetSchemaRequest struct {
	FeatureSetUid string            `protobuf:"bytes,1,opt,name=feature_set_uid,json=featureSetUid" json:"feature_set_uid,omitempty"`
	Schema        map[string]string `protobuf:"bytes,2,rep,name=schema" json:"schema,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateFeatureSetSchemaRequest) Reset()                    { *m = CreateFeatureSetSchemaRequest{} }
func (m *CreateFeatureSetSchemaRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFeatureSetSchemaRequest) ProtoMessage()               {}
func (*CreateFeatureSetSchemaRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *CreateFeatureSetSchemaRequest) GetFeatureSetUid() string {
	if m != nil {
		return m.FeatureSetUid
	}
	return ""
}

func (m *CreateFeatureSetSchemaRequest) GetSchema() map[string]string {
	if m != nil {
		return m.Schema
	}
	return nil
}

type DeleteFeatureSetSchemaRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *DeleteFeatureSetSchemaRequest) Reset()                    { *m = DeleteFeatureSetSchemaRequest{} }
func (m *DeleteFeatureSetSchemaRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteFeatureSetSchemaRequest) ProtoMessage()               {}
func (*DeleteFeatureSetSchemaRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *DeleteFeatureSetSchemaRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*FeatureSet)(nil), "serving.FeatureSet")
	proto.RegisterType((*FeatureSetSchema)(nil), "serving.FeatureSetSchema")
	proto.RegisterType((*FeatureSetSchemaField)(nil), "serving.FeatureSetSchemaField")
	proto.RegisterType((*ListFeatureSetsRequest)(nil), "serving.ListFeatureSetsRequest")
	proto.RegisterType((*FindFeatureSetRequest)(nil), "serving.FindFeatureSetRequest")
	proto.RegisterType((*CreateFeatureSetRequest)(nil), "serving.CreateFeatureSetRequest")
	proto.RegisterType((*DeleteFeatureSetRequest)(nil), "serving.DeleteFeatureSetRequest")
	proto.RegisterType((*ListFeatureSetSchemasRequest)(nil), "serving.ListFeatureSetSchemasRequest")
	proto.RegisterType((*FindFeatureSetSchemaRequest)(nil), "serving.FindFeatureSetSchemaRequest")
	proto.RegisterType((*CreateFeatureSetSchemaRequest)(nil), "serving.CreateFeatureSetSchemaRequest")
	proto.RegisterType((*DeleteFeatureSetSchemaRequest)(nil), "serving.DeleteFeatureSetSchemaRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FeaturesService service

type FeaturesServiceClient interface {
	List(ctx context.Context, in *ListFeatureSetsRequest, opts ...grpc.CallOption) (FeaturesService_ListClient, error)
	Find(ctx context.Context, in *FindFeatureSetRequest, opts ...grpc.CallOption) (*FeatureSet, error)
	Create(ctx context.Context, in *CreateFeatureSetRequest, opts ...grpc.CallOption) (*FeatureSet, error)
	Delete(ctx context.Context, in *DeleteFeatureSetRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Schemas
	ListSchemas(ctx context.Context, in *ListFeatureSetSchemasRequest, opts ...grpc.CallOption) (FeaturesService_ListSchemasClient, error)
	FindSchema(ctx context.Context, in *FindFeatureSetSchemaRequest, opts ...grpc.CallOption) (*FeatureSetSchema, error)
	CreateSchema(ctx context.Context, in *CreateFeatureSetSchemaRequest, opts ...grpc.CallOption) (*FeatureSetSchema, error)
	DeleteSchema(ctx context.Context, in *DeleteFeatureSetSchemaRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type featuresServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeaturesServiceClient(cc *grpc.ClientConn) FeaturesServiceClient {
	return &featuresServiceClient{cc}
}

func (c *featuresServiceClient) List(ctx context.Context, in *ListFeatureSetsRequest, opts ...grpc.CallOption) (FeaturesService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FeaturesService_serviceDesc.Streams[0], c.cc, "/serving.FeaturesService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &featuresServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeaturesService_ListClient interface {
	Recv() (*FeatureSet, error)
	grpc.ClientStream
}

type featuresServiceListClient struct {
	grpc.ClientStream
}

func (x *featuresServiceListClient) Recv() (*FeatureSet, error) {
	m := new(FeatureSet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *featuresServiceClient) Find(ctx context.Context, in *FindFeatureSetRequest, opts ...grpc.CallOption) (*FeatureSet, error) {
	out := new(FeatureSet)
	err := grpc.Invoke(ctx, "/serving.FeaturesService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuresServiceClient) Create(ctx context.Context, in *CreateFeatureSetRequest, opts ...grpc.CallOption) (*FeatureSet, error) {
	out := new(FeatureSet)
	err := grpc.Invoke(ctx, "/serving.FeaturesService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuresServiceClient) Delete(ctx context.Context, in *DeleteFeatureSetRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/serving.FeaturesService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuresServiceClient) ListSchemas(ctx context.Context, in *ListFeatureSetSchemasRequest, opts ...grpc.CallOption) (FeaturesService_ListSchemasClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FeaturesService_serviceDesc.Streams[1], c.cc, "/serving.FeaturesService/ListSchemas", opts...)
	if err != nil {
		return nil, err
	}
	x := &featuresServiceListSchemasClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeaturesService_ListSchemasClient interface {
	Recv() (*FeatureSetSchema, error)
	grpc.ClientStream
}

type featuresServiceListSchemasClient struct {
	grpc.ClientStream
}

func (x *featuresServiceListSchemasClient) Recv() (*FeatureSetSchema, error) {
	m := new(FeatureSetSchema)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *featuresServiceClient) FindSchema(ctx context.Context, in *FindFeatureSetSchemaRequest, opts ...grpc.CallOption) (*FeatureSetSchema, error) {
	out := new(FeatureSetSchema)
	err := grpc.Invoke(ctx, "/serving.FeaturesService/FindSchema", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuresServiceClient) CreateSchema(ctx context.Context, in *CreateFeatureSetSchemaRequest, opts ...grpc.CallOption) (*FeatureSetSchema, error) {
	out := new(FeatureSetSchema)
	err := grpc.Invoke(ctx, "/serving.FeaturesService/CreateSchema", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featuresServiceClient) DeleteSchema(ctx context.Context, in *DeleteFeatureSetSchemaRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/serving.FeaturesService/DeleteSchema", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FeaturesService service

type FeaturesServiceServer interface {
	List(*ListFeatureSetsRequest, FeaturesService_ListServer) error
	Find(context.Context, *FindFeatureSetRequest) (*FeatureSet, error)
	Create(context.Context, *CreateFeatureSetRequest) (*FeatureSet, error)
	Delete(context.Context, *DeleteFeatureSetRequest) (*EmptyResponse, error)
	// Schemas
	ListSchemas(*ListFeatureSetSchemasRequest, FeaturesService_ListSchemasServer) error
	FindSchema(context.Context, *FindFeatureSetSchemaRequest) (*FeatureSetSchema, error)
	CreateSchema(context.Context, *CreateFeatureSetSchemaRequest) (*FeatureSetSchema, error)
	DeleteSchema(context.Context, *DeleteFeatureSetSchemaRequest) (*EmptyResponse, error)
}

func RegisterFeaturesServiceServer(s *grpc.Server, srv FeaturesServiceServer) {
	s.RegisterService(&_FeaturesService_serviceDesc, srv)
}

func _FeaturesService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFeatureSetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeaturesServiceServer).List(m, &featuresServiceListServer{stream})
}

type FeaturesService_ListServer interface {
	Send(*FeatureSet) error
	grpc.ServerStream
}

type featuresServiceListServer struct {
	grpc.ServerStream
}

func (x *featuresServiceListServer) Send(m *FeatureSet) error {
	return x.ServerStream.SendMsg(m)
}

func _FeaturesService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFeatureSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturesServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.FeaturesService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturesServiceServer).Find(ctx, req.(*FindFeatureSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeatureSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.FeaturesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturesServiceServer).Create(ctx, req.(*CreateFeatureSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeatureSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.FeaturesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturesServiceServer).Delete(ctx, req.(*DeleteFeatureSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturesService_ListSchemas_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFeatureSetSchemasRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeaturesServiceServer).ListSchemas(m, &featuresServiceListSchemasServer{stream})
}

type FeaturesService_ListSchemasServer interface {
	Send(*FeatureSetSchema) error
	grpc.ServerStream
}

type featuresServiceListSchemasServer struct {
	grpc.ServerStream
}

func (x *featuresServiceListSchemasServer) Send(m *FeatureSetSchema) error {
	return x.ServerStream.SendMsg(m)
}

func _FeaturesService_FindSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFeatureSetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturesServiceServer).FindSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.FeaturesService/FindSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturesServiceServer).FindSchema(ctx, req.(*FindFeatureSetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturesService_CreateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeatureSetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturesServiceServer).CreateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.FeaturesService/CreateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturesServiceServer).CreateSchema(ctx, req.(*CreateFeatureSetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeaturesService_DeleteSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeatureSetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeaturesServiceServer).DeleteSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.FeaturesService/DeleteSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeaturesServiceServer).DeleteSchema(ctx, req.(*DeleteFeatureSetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeaturesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serving.FeaturesService",
	HandlerType: (*FeaturesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _FeaturesService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FeaturesService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FeaturesService_Delete_Handler,
		},
		{
			MethodName: "FindSchema",
			Handler:    _FeaturesService_FindSchema_Handler,
		},
		{
			MethodName: "CreateSchema",
			Handler:    _FeaturesService_CreateSchema_Handler,
		},
		{
			MethodName: "DeleteSchema",
			Handler:    _FeaturesService_DeleteSchema_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _FeaturesService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListSchemas",
			Handler:       _FeaturesService_ListSchemas_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/features.proto",
}

func init() { proto.RegisterFile("protos/features.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0x13, 0xc7, 0x9f, 0x32, 0xe9, 0xa7, 0x96, 0x85, 0xb4, 0xc6, 0x10, 0x88, 0x2c, 0xa8,
	0x8a, 0x90, 0xd2, 0x12, 0x24, 0xc4, 0x8f, 0x04, 0x0a, 0xd0, 0x5c, 0x54, 0x80, 0x54, 0x07, 0xae,
	0x23, 0x13, 0x4f, 0xc0, 0x4a, 0xe2, 0x18, 0xef, 0xba, 0x92, 0xc5, 0x2d, 0xcf, 0xc7, 0x35, 0x8f,
	0x83, 0xf6, 0xc7, 0x76, 0x9c, 0xd8, 0xa6, 0x77, 0xbb, 0x33, 0xb3, 0x67, 0xe6, 0x9c, 0x39, 0x36,
	0x74, 0xc3, 0x68, 0xcd, 0xd6, 0xf4, 0x74, 0x8e, 0x2e, 0x8b, 0x23, 0xa4, 0x03, 0x71, 0x27, 0xff,
	0x51, 0x8c, 0xae, 0xfc, 0xe0, 0x9b, 0x75, 0x43, 0xe5, 0x67, 0xeb, 0x08, 0x65, 0xce, 0xfe, 0xa5,
	0x01, 0x8c, 0x65, 0xf9, 0x04, 0x19, 0x39, 0x80, 0x66, 0xec, 0x7b, 0xa6, 0xd6, 0xd7, 0x4e, 0xda,
	0x0e, 0x3f, 0x12, 0x02, 0x7a, 0xe0, 0xae, 0xd0, 0x6c, 0x88, 0x90, 0x38, 0xf3, 0xd8, 0x02, 0x13,
	0x6a, 0x36, 0xfb, 0x4d, 0x1e, 0xe3, 0x67, 0xd2, 0x03, 0x98, 0x45, 0xe8, 0x32, 0xf4, 0xa6, 0x2e,
	0x33, 0xf5, 0xbe, 0x76, 0xd2, 0x72, 0xda, 0x2a, 0x32, 0x62, 0x3c, 0x1d, 0x87, 0x5e, 0x9a, 0x6e,
	0xc9, 0xb4, 0x8a, 0x8c, 0x98, 0xfd, 0x13, 0x0e, 0xf2, 0x29, 0x26, 0xb3, 0xef, 0xb8, 0x72, 0x4b,
	0x66, 0x29, 0xf6, 0x68, 0x6c, 0xf7, 0x78, 0x06, 0xc6, 0xdc, 0xc7, 0xa5, 0x27, 0x07, 0xeb, 0x0c,
	0xef, 0x0d, 0x14, 0xf1, 0xc1, 0x36, 0xf6, 0x98, 0x97, 0x39, 0xaa, 0xda, 0xbe, 0x80, 0x6e, 0x69,
	0x41, 0xc6, 0x5d, 0xdb, 0xe0, 0xde, 0x03, 0xb8, 0x72, 0x97, 0x31, 0x4e, 0x59, 0x12, 0xa6, 0xaa,
	0xb4, 0x45, 0xe4, 0x73, 0x12, 0xa2, 0x6d, 0xc2, 0xe1, 0x07, 0x9f, 0xb2, 0x1c, 0x8f, 0x3a, 0xf8,
	0x23, 0x46, 0xca, 0xec, 0x47, 0xd0, 0x1d, 0xfb, 0x81, 0x97, 0x67, 0x54, 0x62, 0x97, 0xa7, 0x3d,
	0x82, 0xa3, 0x77, 0x82, 0xd5, 0x6e, 0x71, 0xd9, 0x48, 0xe9, 0x3a, 0x1a, 0xf9, 0x3a, 0xec, 0xc7,
	0x70, 0xf4, 0x1e, 0x97, 0x58, 0x06, 0xb1, 0xdb, 0x6f, 0x0c, 0x77, 0x8b, 0x43, 0x4b, 0x11, 0xd2,
	0xd1, 0xc9, 0x31, 0xec, 0x2b, 0x4b, 0x4d, 0x29, 0xb2, 0x69, 0xfe, 0xfa, 0xff, 0x79, 0xf6, 0xe4,
	0x8b, 0xef, 0xd9, 0xa7, 0x70, 0xa7, 0x48, 0x51, 0xe2, 0x54, 0x37, 0xfe, 0xad, 0x41, 0x6f, 0x9b,
	0x69, 0xf1, 0xcd, 0x35, 0x5b, 0x93, 0x0b, 0x30, 0xa8, 0x78, 0x28, 0x54, 0xe8, 0x0c, 0x87, 0xd9,
	0xee, 0x6b, 0xf1, 0x07, 0xf2, 0x76, 0x1e, 0xb0, 0x28, 0x71, 0x14, 0x82, 0xf5, 0x02, 0x3a, 0x1b,
	0x61, 0x3e, 0xf6, 0x02, 0x93, 0x74, 0xec, 0x05, 0x26, 0xe4, 0x16, 0xb4, 0xc4, 0xc6, 0xd5, 0xfa,
	0xe5, 0xe5, 0x65, 0xe3, 0xb9, 0x66, 0x3f, 0x81, 0xde, 0xb6, 0xec, 0xff, 0xd0, 0x60, 0xf8, 0x47,
	0x87, 0x7d, 0x55, 0x4d, 0x27, 0x7c, 0xe6, 0x19, 0x92, 0xd7, 0xa0, 0xf3, 0x85, 0x90, 0xfb, 0x19,
	0x8b, 0x72, 0x53, 0x59, 0x37, 0x4b, 0x2c, 0x7e, 0xa6, 0x91, 0x57, 0xa0, 0xf3, 0x45, 0x90, 0x8d,
	0x2f, 0xa0, 0xcc, 0x7a, 0xa5, 0xcf, 0xc9, 0x1b, 0x30, 0xa4, 0x66, 0xa4, 0x5f, 0x29, 0x62, 0x2d,
	0xc0, 0x5b, 0x30, 0xa4, 0x08, 0x1b, 0x00, 0x15, 0x66, 0xb4, 0x0e, 0xb3, 0x8a, 0xf3, 0x55, 0xc8,
	0x12, 0x07, 0x69, 0xb8, 0x0e, 0x28, 0x92, 0x4b, 0xe8, 0x70, 0xca, 0xca, 0x88, 0xe4, 0x61, 0x85,
	0x10, 0x45, 0xa3, 0x5a, 0xb7, 0x2b, 0xbf, 0xf8, 0x33, 0x8d, 0x7c, 0x04, 0xe0, 0x2a, 0xa8, 0xbf,
	0xcb, 0x83, 0x0a, 0x69, 0x0a, 0xeb, 0xaa, 0x01, 0x24, 0x97, 0xb0, 0x27, 0x55, 0x51, 0xf7, 0xe3,
	0xeb, 0x39, 0xae, 0x0e, 0xf2, 0x13, 0xec, 0x49, 0x9d, 0x76, 0x20, 0x6b, 0x4d, 0x55, 0x25, 0xe2,
	0x57, 0x43, 0xfc, 0xe3, 0x9f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x4c, 0x23, 0x60, 0x18,
	0x06, 0x00, 0x00,
}
