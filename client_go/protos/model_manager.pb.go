// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/model_manager.proto

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

type ListModelsRequest struct {
}

func (m *ListModelsRequest) Reset()                    { *m = ListModelsRequest{} }
func (m *ListModelsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListModelsRequest) ProtoMessage()               {}
func (*ListModelsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*ListModelsRequest)(nil), "serving.ListModelsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ModelManagerService service

type ModelManagerServiceClient interface {
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (ModelManagerService_ListModelsClient, error)
}

type modelManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelManagerServiceClient(cc *grpc.ClientConn) ModelManagerServiceClient {
	return &modelManagerServiceClient{cc}
}

func (c *modelManagerServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (ModelManagerService_ListModelsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ModelManagerService_serviceDesc.Streams[0], c.cc, "/serving.ModelManagerService/ListModels", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelManagerServiceListModelsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ModelManagerService_ListModelsClient interface {
	Recv() (*Model, error)
	grpc.ClientStream
}

type modelManagerServiceListModelsClient struct {
	grpc.ClientStream
}

func (x *modelManagerServiceListModelsClient) Recv() (*Model, error) {
	m := new(Model)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ModelManagerService service

type ModelManagerServiceServer interface {
	ListModels(*ListModelsRequest, ModelManagerService_ListModelsServer) error
}

func RegisterModelManagerServiceServer(s *grpc.Server, srv ModelManagerServiceServer) {
	s.RegisterService(&_ModelManagerService_serviceDesc, srv)
}

func _ModelManagerService_ListModels_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListModelsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ModelManagerServiceServer).ListModels(m, &modelManagerServiceListModelsServer{stream})
}

type ModelManagerService_ListModelsServer interface {
	Send(*Model) error
	grpc.ServerStream
}

type modelManagerServiceListModelsServer struct {
	grpc.ServerStream
}

func (x *modelManagerServiceListModelsServer) Send(m *Model) error {
	return x.ServerStream.SendMsg(m)
}

var _ModelManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serving.ModelManagerService",
	HandlerType: (*ModelManagerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListModels",
			Handler:       _ModelManagerService_ListModels_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/model_manager.proto",
}

func init() { proto.RegisterFile("protos/model_manager.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x89, 0xcf, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d,
	0xd2, 0x03, 0x0b, 0x0a, 0xb1, 0x17, 0xa7, 0x16, 0x95, 0x65, 0xe6, 0xa5, 0x4b, 0x09, 0x42, 0x15,
	0x25, 0xe7, 0x17, 0xa5, 0x42, 0xe4, 0x94, 0x84, 0xb9, 0x04, 0x7d, 0x32, 0x8b, 0x4b, 0x7c, 0x41,
	0xda, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x8c, 0x02, 0xb9, 0x84, 0xc1, 0x02, 0xbe,
	0x10, 0x63, 0x82, 0x41, 0xda, 0x93, 0x53, 0x85, 0xac, 0xb8, 0xb8, 0x10, 0x6a, 0x85, 0xa4, 0xf4,
	0xa0, 0xc6, 0xea, 0x61, 0x18, 0x20, 0xc5, 0x07, 0x97, 0x03, 0x8b, 0x1b, 0x30, 0x26, 0xb1, 0x81,
	0xad, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x16, 0x30, 0x22, 0x93, 0xa8, 0x00, 0x00, 0x00,
}