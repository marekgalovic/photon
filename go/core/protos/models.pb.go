// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

package photon

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

type Model struct {
	Id                  int64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                string          `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RunnerType          string          `protobuf:"bytes,3,opt,name=runner_type,json=runnerType" json:"runner_type,omitempty"`
	Replicas            int32           `protobuf:"varint,4,opt,name=replicas" json:"replicas,omitempty"`
	Features            []*ModelFeature `protobuf:"bytes,5,rep,name=features" json:"features,omitempty"`
	PrecomputedFeatures []*ModelFeature `protobuf:"bytes,6,rep,name=precomputed_features,json=precomputedFeatures" json:"precomputed_features,omitempty"`
	CreatedAt           int32           `protobuf:"varint,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt           int32           `protobuf:"varint,8,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Model) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetRunnerType() string {
	if m != nil {
		return m.RunnerType
	}
	return ""
}

func (m *Model) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Model) GetFeatures() []*ModelFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *Model) GetPrecomputedFeatures() []*ModelFeature {
	if m != nil {
		return m.PrecomputedFeatures
	}
	return nil
}

func (m *Model) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Model) GetUpdatedAt() int32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ModelVersion struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ModelId   int64  `protobuf:"varint,2,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	FileName  string `protobuf:"bytes,4,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	IsPrimary bool   `protobuf:"varint,5,opt,name=is_primary,json=isPrimary" json:"is_primary,omitempty"`
	IsShadow  bool   `protobuf:"varint,6,opt,name=is_shadow,json=isShadow" json:"is_shadow,omitempty"`
	CreatedAt int32  `protobuf:"varint,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *ModelVersion) Reset()                    { *m = ModelVersion{} }
func (m *ModelVersion) String() string            { return proto.CompactTextString(m) }
func (*ModelVersion) ProtoMessage()               {}
func (*ModelVersion) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ModelVersion) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModelVersion) GetModelId() int64 {
	if m != nil {
		return m.ModelId
	}
	return 0
}

func (m *ModelVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelVersion) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ModelVersion) GetIsPrimary() bool {
	if m != nil {
		return m.IsPrimary
	}
	return false
}

func (m *ModelVersion) GetIsShadow() bool {
	if m != nil {
		return m.IsShadow
	}
	return false
}

func (m *ModelVersion) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type ModelFeature struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Required bool   `protobuf:"varint,2,opt,name=required" json:"required,omitempty"`
}

func (m *ModelFeature) Reset()                    { *m = ModelFeature{} }
func (m *ModelFeature) String() string            { return proto.CompactTextString(m) }
func (*ModelFeature) ProtoMessage()               {}
func (*ModelFeature) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ModelFeature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelFeature) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

type PrecomputedFeaturesSet struct {
	Features []*ModelFeature `protobuf:"bytes,1,rep,name=features" json:"features,omitempty"`
}

func (m *PrecomputedFeaturesSet) Reset()                    { *m = PrecomputedFeaturesSet{} }
func (m *PrecomputedFeaturesSet) String() string            { return proto.CompactTextString(m) }
func (*PrecomputedFeaturesSet) ProtoMessage()               {}
func (*PrecomputedFeaturesSet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *PrecomputedFeaturesSet) GetFeatures() []*ModelFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

type FindModelRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *FindModelRequest) Reset()                    { *m = FindModelRequest{} }
func (m *FindModelRequest) String() string            { return proto.CompactTextString(m) }
func (*FindModelRequest) ProtoMessage()               {}
func (*FindModelRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *FindModelRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateModelRequest struct {
	Name                string                            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RunnerType          string                            `protobuf:"bytes,2,opt,name=runner_type,json=runnerType" json:"runner_type,omitempty"`
	Replicas            int32                             `protobuf:"varint,3,opt,name=replicas" json:"replicas,omitempty"`
	Features            []*ModelFeature                   `protobuf:"bytes,4,rep,name=features" json:"features,omitempty"`
	PrecomputedFeatures map[int64]*PrecomputedFeaturesSet `protobuf:"bytes,5,rep,name=precomputed_features,json=precomputedFeatures" json:"precomputed_features,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateModelRequest) Reset()                    { *m = CreateModelRequest{} }
func (m *CreateModelRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateModelRequest) ProtoMessage()               {}
func (*CreateModelRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *CreateModelRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateModelRequest) GetRunnerType() string {
	if m != nil {
		return m.RunnerType
	}
	return ""
}

func (m *CreateModelRequest) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *CreateModelRequest) GetFeatures() []*ModelFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *CreateModelRequest) GetPrecomputedFeatures() map[int64]*PrecomputedFeaturesSet {
	if m != nil {
		return m.PrecomputedFeatures
	}
	return nil
}

type CreateModelResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateModelResponse) Reset()                    { *m = CreateModelResponse{} }
func (m *CreateModelResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateModelResponse) ProtoMessage()               {}
func (*CreateModelResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *CreateModelResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteModelRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteModelRequest) Reset()                    { *m = DeleteModelRequest{} }
func (m *DeleteModelRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteModelRequest) ProtoMessage()               {}
func (*DeleteModelRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *DeleteModelRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListVersionsRequest struct {
	ModelId int64 `protobuf:"varint,1,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
}

func (m *ListVersionsRequest) Reset()                    { *m = ListVersionsRequest{} }
func (m *ListVersionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVersionsRequest) ProtoMessage()               {}
func (*ListVersionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ListVersionsRequest) GetModelId() int64 {
	if m != nil {
		return m.ModelId
	}
	return 0
}

type FindVersionRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *FindVersionRequest) Reset()                    { *m = FindVersionRequest{} }
func (m *FindVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*FindVersionRequest) ProtoMessage()               {}
func (*FindVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *FindVersionRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SetPrimaryVersionRequest struct {
	ModelId int64 `protobuf:"varint,1,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	Id      int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *SetPrimaryVersionRequest) Reset()                    { *m = SetPrimaryVersionRequest{} }
func (m *SetPrimaryVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPrimaryVersionRequest) ProtoMessage()               {}
func (*SetPrimaryVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *SetPrimaryVersionRequest) GetModelId() int64 {
	if m != nil {
		return m.ModelId
	}
	return 0
}

func (m *SetPrimaryVersionRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateVersionRequest struct {
	// Types that are valid to be assigned to Value:
	//	*CreateVersionRequest_Version
	//	*CreateVersionRequest_Data
	Value isCreateVersionRequest_Value `protobuf_oneof:"value"`
}

func (m *CreateVersionRequest) Reset()                    { *m = CreateVersionRequest{} }
func (m *CreateVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVersionRequest) ProtoMessage()               {}
func (*CreateVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

type isCreateVersionRequest_Value interface {
	isCreateVersionRequest_Value()
}

type CreateVersionRequest_Version struct {
	Version *ModelVersion `protobuf:"bytes,1,opt,name=version,oneof"`
}
type CreateVersionRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*CreateVersionRequest_Version) isCreateVersionRequest_Value() {}
func (*CreateVersionRequest_Data) isCreateVersionRequest_Value()    {}

func (m *CreateVersionRequest) GetValue() isCreateVersionRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *CreateVersionRequest) GetVersion() *ModelVersion {
	if x, ok := m.GetValue().(*CreateVersionRequest_Version); ok {
		return x.Version
	}
	return nil
}

func (m *CreateVersionRequest) GetData() []byte {
	if x, ok := m.GetValue().(*CreateVersionRequest_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CreateVersionRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CreateVersionRequest_OneofMarshaler, _CreateVersionRequest_OneofUnmarshaler, _CreateVersionRequest_OneofSizer, []interface{}{
		(*CreateVersionRequest_Version)(nil),
		(*CreateVersionRequest_Data)(nil),
	}
}

func _CreateVersionRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CreateVersionRequest)
	// value
	switch x := m.Value.(type) {
	case *CreateVersionRequest_Version:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Version); err != nil {
			return err
		}
	case *CreateVersionRequest_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Data)
	case nil:
	default:
		return fmt.Errorf("CreateVersionRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _CreateVersionRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CreateVersionRequest)
	switch tag {
	case 1: // value.version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ModelVersion)
		err := b.DecodeMessage(msg)
		m.Value = &CreateVersionRequest_Version{msg}
		return true, err
	case 2: // value.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &CreateVersionRequest_Data{x}
		return true, err
	default:
		return false, nil
	}
}

func _CreateVersionRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CreateVersionRequest)
	// value
	switch x := m.Value.(type) {
	case *CreateVersionRequest_Version:
		s := proto.Size(x.Version)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CreateVersionRequest_Data:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CreateVersionResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateVersionResponse) Reset()                    { *m = CreateVersionResponse{} }
func (m *CreateVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateVersionResponse) ProtoMessage()               {}
func (*CreateVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *CreateVersionResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteVersionRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteVersionRequest) Reset()                    { *m = DeleteVersionRequest{} }
func (m *DeleteVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteVersionRequest) ProtoMessage()               {}
func (*DeleteVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *DeleteVersionRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Model)(nil), "photon.Model")
	proto.RegisterType((*ModelVersion)(nil), "photon.ModelVersion")
	proto.RegisterType((*ModelFeature)(nil), "photon.ModelFeature")
	proto.RegisterType((*PrecomputedFeaturesSet)(nil), "photon.PrecomputedFeaturesSet")
	proto.RegisterType((*FindModelRequest)(nil), "photon.FindModelRequest")
	proto.RegisterType((*CreateModelRequest)(nil), "photon.CreateModelRequest")
	proto.RegisterType((*CreateModelResponse)(nil), "photon.CreateModelResponse")
	proto.RegisterType((*DeleteModelRequest)(nil), "photon.DeleteModelRequest")
	proto.RegisterType((*ListVersionsRequest)(nil), "photon.ListVersionsRequest")
	proto.RegisterType((*FindVersionRequest)(nil), "photon.FindVersionRequest")
	proto.RegisterType((*SetPrimaryVersionRequest)(nil), "photon.SetPrimaryVersionRequest")
	proto.RegisterType((*CreateVersionRequest)(nil), "photon.CreateVersionRequest")
	proto.RegisterType((*CreateVersionResponse)(nil), "photon.CreateVersionResponse")
	proto.RegisterType((*DeleteVersionRequest)(nil), "photon.DeleteVersionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ModelsService service

type ModelsServiceClient interface {
	// Models
	List(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (ModelsService_ListClient, error)
	Find(ctx context.Context, in *FindModelRequest, opts ...grpc.CallOption) (*Model, error)
	Create(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error)
	Delete(ctx context.Context, in *DeleteModelRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Versions
	ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (ModelsService_ListVersionsClient, error)
	FindVersion(ctx context.Context, in *FindVersionRequest, opts ...grpc.CallOption) (*ModelVersion, error)
	SetPrimaryVersion(ctx context.Context, in *SetPrimaryVersionRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	CreateVersion(ctx context.Context, opts ...grpc.CallOption) (ModelsService_CreateVersionClient, error)
	DeleteVersion(ctx context.Context, in *DeleteVersionRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type modelsServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelsServiceClient(cc *grpc.ClientConn) ModelsServiceClient {
	return &modelsServiceClient{cc}
}

func (c *modelsServiceClient) List(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (ModelsService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ModelsService_serviceDesc.Streams[0], c.cc, "/photon.ModelsService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelsServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModelsService_ListClient interface {
	Recv() (*Model, error)
	grpc.ClientStream
}

type modelsServiceListClient struct {
	grpc.ClientStream
}

func (x *modelsServiceListClient) Recv() (*Model, error) {
	m := new(Model)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelsServiceClient) Find(ctx context.Context, in *FindModelRequest, opts ...grpc.CallOption) (*Model, error) {
	out := new(Model)
	err := grpc.Invoke(ctx, "/photon.ModelsService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelsServiceClient) Create(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error) {
	out := new(CreateModelResponse)
	err := grpc.Invoke(ctx, "/photon.ModelsService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelsServiceClient) Delete(ctx context.Context, in *DeleteModelRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/photon.ModelsService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelsServiceClient) ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (ModelsService_ListVersionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ModelsService_serviceDesc.Streams[1], c.cc, "/photon.ModelsService/ListVersions", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelsServiceListVersionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModelsService_ListVersionsClient interface {
	Recv() (*ModelVersion, error)
	grpc.ClientStream
}

type modelsServiceListVersionsClient struct {
	grpc.ClientStream
}

func (x *modelsServiceListVersionsClient) Recv() (*ModelVersion, error) {
	m := new(ModelVersion)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelsServiceClient) FindVersion(ctx context.Context, in *FindVersionRequest, opts ...grpc.CallOption) (*ModelVersion, error) {
	out := new(ModelVersion)
	err := grpc.Invoke(ctx, "/photon.ModelsService/FindVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelsServiceClient) SetPrimaryVersion(ctx context.Context, in *SetPrimaryVersionRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/photon.ModelsService/SetPrimaryVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelsServiceClient) CreateVersion(ctx context.Context, opts ...grpc.CallOption) (ModelsService_CreateVersionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ModelsService_serviceDesc.Streams[2], c.cc, "/photon.ModelsService/CreateVersion", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelsServiceCreateVersionClient{stream}
	return x, nil
}

type ModelsService_CreateVersionClient interface {
	Send(*CreateVersionRequest) error
	Recv() (*CreateVersionResponse, error)
	grpc.ClientStream
}

type modelsServiceCreateVersionClient struct {
	grpc.ClientStream
}

func (x *modelsServiceCreateVersionClient) Send(m *CreateVersionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *modelsServiceCreateVersionClient) Recv() (*CreateVersionResponse, error) {
	m := new(CreateVersionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelsServiceClient) DeleteVersion(ctx context.Context, in *DeleteVersionRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/photon.ModelsService/DeleteVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ModelsService service

type ModelsServiceServer interface {
	// Models
	List(*EmptyRequest, ModelsService_ListServer) error
	Find(context.Context, *FindModelRequest) (*Model, error)
	Create(context.Context, *CreateModelRequest) (*CreateModelResponse, error)
	Delete(context.Context, *DeleteModelRequest) (*EmptyResponse, error)
	// Versions
	ListVersions(*ListVersionsRequest, ModelsService_ListVersionsServer) error
	FindVersion(context.Context, *FindVersionRequest) (*ModelVersion, error)
	SetPrimaryVersion(context.Context, *SetPrimaryVersionRequest) (*EmptyResponse, error)
	CreateVersion(ModelsService_CreateVersionServer) error
	DeleteVersion(context.Context, *DeleteVersionRequest) (*EmptyResponse, error)
}

func RegisterModelsServiceServer(s *grpc.Server, srv ModelsServiceServer) {
	s.RegisterService(&_ModelsService_serviceDesc, srv)
}

func _ModelsService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelsServiceServer).List(m, &modelsServiceListServer{stream})
}

type ModelsService_ListServer interface {
	Send(*Model) error
	grpc.ServerStream
}

type modelsServiceListServer struct {
	grpc.ServerStream
}

func (x *modelsServiceListServer) Send(m *Model) error {
	return x.ServerStream.SendMsg(m)
}

func _ModelsService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelsServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.ModelsService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelsServiceServer).Find(ctx, req.(*FindModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.ModelsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelsServiceServer).Create(ctx, req.(*CreateModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.ModelsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelsServiceServer).Delete(ctx, req.(*DeleteModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelsService_ListVersions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListVersionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelsServiceServer).ListVersions(m, &modelsServiceListVersionsServer{stream})
}

type ModelsService_ListVersionsServer interface {
	Send(*ModelVersion) error
	grpc.ServerStream
}

type modelsServiceListVersionsServer struct {
	grpc.ServerStream
}

func (x *modelsServiceListVersionsServer) Send(m *ModelVersion) error {
	return x.ServerStream.SendMsg(m)
}

func _ModelsService_FindVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelsServiceServer).FindVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.ModelsService/FindVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelsServiceServer).FindVersion(ctx, req.(*FindVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelsService_SetPrimaryVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPrimaryVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelsServiceServer).SetPrimaryVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.ModelsService/SetPrimaryVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelsServiceServer).SetPrimaryVersion(ctx, req.(*SetPrimaryVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelsService_CreateVersion_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelsServiceServer).CreateVersion(&modelsServiceCreateVersionServer{stream})
}

type ModelsService_CreateVersionServer interface {
	Send(*CreateVersionResponse) error
	Recv() (*CreateVersionRequest, error)
	grpc.ServerStream
}

type modelsServiceCreateVersionServer struct {
	grpc.ServerStream
}

func (x *modelsServiceCreateVersionServer) Send(m *CreateVersionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *modelsServiceCreateVersionServer) Recv() (*CreateVersionRequest, error) {
	m := new(CreateVersionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ModelsService_DeleteVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelsServiceServer).DeleteVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photon.ModelsService/DeleteVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelsServiceServer).DeleteVersion(ctx, req.(*DeleteVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "photon.ModelsService",
	HandlerType: (*ModelsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _ModelsService_Find_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ModelsService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ModelsService_Delete_Handler,
		},
		{
			MethodName: "FindVersion",
			Handler:    _ModelsService_FindVersion_Handler,
		},
		{
			MethodName: "SetPrimaryVersion",
			Handler:    _ModelsService_SetPrimaryVersion_Handler,
		},
		{
			MethodName: "DeleteVersion",
			Handler:    _ModelsService_DeleteVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ModelsService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListVersions",
			Handler:       _ModelsService_ListVersions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateVersion",
			Handler:       _ModelsService_CreateVersion_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "models.proto",
}

func init() { proto.RegisterFile("models.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5d, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0x1d, 0x27, 0x71, 0x6e, 0x92, 0x4f, 0xfd, 0xa6, 0x29, 0x32, 0x2e, 0x85, 0xc8, 0xe2,
	0x27, 0x2f, 0x84, 0xa8, 0xe5, 0x01, 0x81, 0x04, 0x2a, 0xfd, 0xa1, 0xa0, 0x82, 0x2a, 0x07, 0xf1,
	0x1a, 0x99, 0x78, 0x42, 0x47, 0x24, 0xb6, 0x3b, 0x33, 0x29, 0xca, 0x72, 0x58, 0x0c, 0xab, 0x60,
	0x01, 0x6c, 0x03, 0x79, 0x66, 0xec, 0xc4, 0x8e, 0x1d, 0xfa, 0x96, 0x99, 0x73, 0xef, 0xf1, 0xb9,
	0xe7, 0x9e, 0x91, 0x02, 0xad, 0x59, 0xe8, 0xe3, 0x29, 0xeb, 0x47, 0x34, 0xe4, 0x21, 0xaa, 0x45,
	0x57, 0x21, 0x0f, 0x03, 0x1b, 0xc6, 0x21, 0xc5, 0xf2, 0xce, 0xf9, 0xa9, 0x43, 0xf5, 0x63, 0x5c,
	0x84, 0xfe, 0x03, 0x9d, 0xf8, 0x96, 0xd6, 0xd5, 0x7a, 0x15, 0x57, 0x27, 0x3e, 0x42, 0x60, 0x04,
	0xde, 0x0c, 0x5b, 0x7a, 0x57, 0xeb, 0x35, 0x5c, 0xf1, 0x1b, 0x3d, 0x80, 0x26, 0x9d, 0x07, 0x01,
	0xa6, 0x23, 0xbe, 0x88, 0xb0, 0x55, 0x11, 0x10, 0xc8, 0xab, 0xcf, 0x8b, 0x08, 0x23, 0x1b, 0x4c,
	0x8a, 0xa3, 0x29, 0x19, 0x7b, 0xcc, 0x32, 0xba, 0x5a, 0xaf, 0xea, 0xa6, 0x67, 0x34, 0x00, 0x73,
	0x82, 0x3d, 0x3e, 0xa7, 0x98, 0x59, 0xd5, 0x6e, 0xa5, 0xd7, 0x3c, 0xe8, 0xf4, 0xa5, 0xa2, 0xbe,
	0x50, 0x70, 0x26, 0x41, 0x37, 0xad, 0x42, 0xef, 0xa0, 0x13, 0x51, 0x3c, 0x0e, 0x67, 0xd1, 0x9c,
	0x63, 0x7f, 0x94, 0x76, 0xd7, 0x36, 0x74, 0xef, 0xac, 0x74, 0x9c, 0x25, 0x44, 0xfb, 0x00, 0x63,
	0x8a, 0xbd, 0x98, 0xc4, 0xe3, 0x56, 0x5d, 0x08, 0x6b, 0xa8, 0x9b, 0x23, 0x1e, 0xc3, 0xf3, 0xc8,
	0x4f, 0x60, 0x53, 0xc2, 0xea, 0xe6, 0x88, 0x3b, 0xbf, 0x34, 0x68, 0x89, 0x6f, 0x7c, 0xc1, 0x94,
	0x91, 0x30, 0x58, 0xb3, 0xea, 0x2e, 0x98, 0xc2, 0xe8, 0x11, 0xf1, 0x85, 0x5d, 0x15, 0xb7, 0x2e,
	0xce, 0xef, 0x97, 0x2e, 0x56, 0x56, 0x5c, 0xdc, 0x83, 0xc6, 0x84, 0x4c, 0xf1, 0x48, 0x00, 0x86,
	0x00, 0xcc, 0xf8, 0xe2, 0x53, 0x0c, 0xee, 0x03, 0x10, 0x36, 0x8a, 0x28, 0x99, 0x79, 0x74, 0x61,
	0x55, 0xbb, 0x5a, 0xcf, 0x74, 0x1b, 0x84, 0x5d, 0xca, 0x8b, 0xb8, 0x97, 0xb0, 0x11, 0xbb, 0xf2,
	0xfc, 0xf0, 0x87, 0x55, 0x13, 0xa8, 0x49, 0xd8, 0x50, 0x9c, 0xff, 0x31, 0xa6, 0xf3, 0x5a, 0x8d,
	0xa1, 0x6c, 0x49, 0xb5, 0x69, 0x2b, 0xda, 0xc4, 0x02, 0xaf, 0xe7, 0x84, 0x62, 0x39, 0x8a, 0xe9,
	0xa6, 0x67, 0xe7, 0x03, 0xdc, 0xb9, 0x5c, 0x37, 0x77, 0x88, 0x79, 0x66, 0xb5, 0xda, 0x6d, 0x56,
	0xeb, 0x38, 0xb0, 0x7d, 0x46, 0x02, 0x5f, 0xa0, 0x2e, 0xbe, 0x9e, 0x63, 0xc6, 0xf3, 0xb6, 0x3a,
	0x7f, 0x74, 0x40, 0xc7, 0x42, 0x7d, 0xa6, 0xac, 0x48, 0x76, 0x2e, 0x98, 0xfa, 0xc6, 0x60, 0x56,
	0x36, 0x04, 0xd3, 0xb8, 0x55, 0x30, 0x27, 0x25, 0xc1, 0x94, 0xb1, 0x3e, 0x4c, 0xba, 0xd7, 0xc5,
	0xf7, 0x0b, 0x0c, 0x3c, 0x0d, 0x38, 0x5d, 0x14, 0xe6, 0xd6, 0x9e, 0x80, 0x55, 0xd6, 0x80, 0xb6,
	0xa1, 0xf2, 0x1d, 0x2f, 0x94, 0x5d, 0xf1, 0x4f, 0xf4, 0x1c, 0xaa, 0x37, 0xde, 0x74, 0x2e, 0xc7,
	0x6f, 0x1e, 0xdc, 0x4f, 0x64, 0x14, 0x2f, 0xcd, 0x95, 0xc5, 0x2f, 0xf5, 0x17, 0x9a, 0xf3, 0x08,
	0x76, 0x32, 0x5a, 0x59, 0x14, 0x06, 0x0c, 0xaf, 0x2d, 0xe4, 0x21, 0xa0, 0x13, 0x3c, 0xc5, 0xb9,
	0x7d, 0xe4, 0xab, 0x06, 0xb0, 0x73, 0x41, 0x18, 0x57, 0x8f, 0x85, 0x25, 0x65, 0xab, 0x8f, 0x44,
	0xcb, 0x3c, 0x92, 0x98, 0x37, 0x0e, 0x83, 0xea, 0x28, 0xe3, 0x3d, 0x05, 0x6b, 0x88, 0xb9, 0x7a,
	0x08, 0xb9, 0xda, 0x72, 0x72, 0x45, 0xa3, 0xa7, 0x34, 0xdf, 0xa0, 0x23, 0x67, 0xcd, 0x51, 0x0c,
	0xa0, 0x7e, 0x23, 0x6f, 0x04, 0x43, 0x3e, 0x04, 0xaa, 0xfa, 0x7c, 0xcb, 0x4d, 0xca, 0x50, 0x07,
	0x0c, 0xdf, 0xe3, 0x9e, 0xe0, 0x6e, 0x9d, 0x6f, 0xb9, 0xe2, 0xf4, 0xb6, 0xae, 0xb6, 0xe0, 0x3c,
	0x81, 0xdd, 0xdc, 0x87, 0x4a, 0x6c, 0x7d, 0x0c, 0x1d, 0x69, 0xeb, 0x66, 0x03, 0x0e, 0x7e, 0x1b,
	0xd0, 0x16, 0x5a, 0xd8, 0x10, 0xd3, 0x1b, 0x32, 0xc6, 0xe8, 0x29, 0x18, 0xb1, 0xd5, 0x28, 0x95,
	0x7a, 0x3a, 0x8b, 0xf8, 0x42, 0xf5, 0xdb, 0xed, 0xcc, 0x00, 0x03, 0x0d, 0x3d, 0x03, 0x23, 0xf6,
	0x19, 0x59, 0x09, 0x90, 0x7f, 0x82, 0xb9, 0x16, 0x74, 0x04, 0x35, 0x39, 0x02, 0xb2, 0xcb, 0x33,
	0x6d, 0xef, 0x15, 0x62, 0x6a, 0xd8, 0x57, 0x50, 0x93, 0xc3, 0x2d, 0x29, 0xd6, 0x33, 0x64, 0xef,
	0xe6, 0x06, 0x50, 0xcd, 0xc7, 0xd0, 0x5a, 0x8d, 0x12, 0x4a, 0xbf, 0x54, 0x10, 0x30, 0xbb, 0x70,
	0x5f, 0x03, 0x0d, 0xbd, 0x81, 0xe6, 0x4a, 0xba, 0x96, 0x32, 0xd6, 0x23, 0x57, 0x4c, 0x81, 0x2e,
	0xe0, 0xff, 0xb5, 0xe0, 0xa1, 0x6e, 0x52, 0x5a, 0x96, 0xc9, 0xb2, 0x99, 0x2e, 0xa1, 0x9d, 0x89,
	0x05, 0xba, 0x97, 0xb5, 0x2f, 0xc7, 0xb2, 0x5f, 0x82, 0x4a, 0xb6, 0x9e, 0x36, 0xd0, 0xd0, 0x09,
	0xb4, 0x33, 0xf9, 0x59, 0x32, 0x16, 0xc5, 0xaa, 0x44, 0xd7, 0xd7, 0x9a, 0xf8, 0x43, 0x70, 0xf8,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x74, 0xc7, 0x98, 0x3a, 0x34, 0x08, 0x00, 0x00,
}
