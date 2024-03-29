// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: protos/enrollmentpb/enrollment.proto

package enrollmentpb

import (
	context "context"
	studentpb "github.com/nicrodriguezval/grpc/protos/studentpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EnrollmentService_Enroll_FullMethodName             = "/enrollment.EnrollmentService/Enroll"
	EnrollmentService_GetStudentsPerTest_FullMethodName = "/enrollment.EnrollmentService/GetStudentsPerTest"
)

// EnrollmentServiceClient is the client API for EnrollmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnrollmentServiceClient interface {
	Enroll(ctx context.Context, opts ...grpc.CallOption) (EnrollmentService_EnrollClient, error)
	GetStudentsPerTest(ctx context.Context, in *GetStudentsPerTestRequest, opts ...grpc.CallOption) (EnrollmentService_GetStudentsPerTestClient, error)
}

type enrollmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnrollmentServiceClient(cc grpc.ClientConnInterface) EnrollmentServiceClient {
	return &enrollmentServiceClient{cc}
}

func (c *enrollmentServiceClient) Enroll(ctx context.Context, opts ...grpc.CallOption) (EnrollmentService_EnrollClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnrollmentService_ServiceDesc.Streams[0], EnrollmentService_Enroll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &enrollmentServiceEnrollClient{stream}
	return x, nil
}

type EnrollmentService_EnrollClient interface {
	Send(*EnrollmentRequest) error
	CloseAndRecv() (*EnrollmentResponse, error)
	grpc.ClientStream
}

type enrollmentServiceEnrollClient struct {
	grpc.ClientStream
}

func (x *enrollmentServiceEnrollClient) Send(m *EnrollmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *enrollmentServiceEnrollClient) CloseAndRecv() (*EnrollmentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EnrollmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *enrollmentServiceClient) GetStudentsPerTest(ctx context.Context, in *GetStudentsPerTestRequest, opts ...grpc.CallOption) (EnrollmentService_GetStudentsPerTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnrollmentService_ServiceDesc.Streams[1], EnrollmentService_GetStudentsPerTest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &enrollmentServiceGetStudentsPerTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnrollmentService_GetStudentsPerTestClient interface {
	Recv() (*studentpb.Student, error)
	grpc.ClientStream
}

type enrollmentServiceGetStudentsPerTestClient struct {
	grpc.ClientStream
}

func (x *enrollmentServiceGetStudentsPerTestClient) Recv() (*studentpb.Student, error) {
	m := new(studentpb.Student)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EnrollmentServiceServer is the server API for EnrollmentService service.
// All implementations must embed UnimplementedEnrollmentServiceServer
// for forward compatibility
type EnrollmentServiceServer interface {
	Enroll(EnrollmentService_EnrollServer) error
	GetStudentsPerTest(*GetStudentsPerTestRequest, EnrollmentService_GetStudentsPerTestServer) error
	mustEmbedUnimplementedEnrollmentServiceServer()
}

// UnimplementedEnrollmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnrollmentServiceServer struct {
}

func (UnimplementedEnrollmentServiceServer) Enroll(EnrollmentService_EnrollServer) error {
	return status.Errorf(codes.Unimplemented, "method Enroll not implemented")
}
func (UnimplementedEnrollmentServiceServer) GetStudentsPerTest(*GetStudentsPerTestRequest, EnrollmentService_GetStudentsPerTestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentsPerTest not implemented")
}
func (UnimplementedEnrollmentServiceServer) mustEmbedUnimplementedEnrollmentServiceServer() {}

// UnsafeEnrollmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnrollmentServiceServer will
// result in compilation errors.
type UnsafeEnrollmentServiceServer interface {
	mustEmbedUnimplementedEnrollmentServiceServer()
}

func RegisterEnrollmentServiceServer(s grpc.ServiceRegistrar, srv EnrollmentServiceServer) {
	s.RegisterService(&EnrollmentService_ServiceDesc, srv)
}

func _EnrollmentService_Enroll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EnrollmentServiceServer).Enroll(&enrollmentServiceEnrollServer{stream})
}

type EnrollmentService_EnrollServer interface {
	SendAndClose(*EnrollmentResponse) error
	Recv() (*EnrollmentRequest, error)
	grpc.ServerStream
}

type enrollmentServiceEnrollServer struct {
	grpc.ServerStream
}

func (x *enrollmentServiceEnrollServer) SendAndClose(m *EnrollmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *enrollmentServiceEnrollServer) Recv() (*EnrollmentRequest, error) {
	m := new(EnrollmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EnrollmentService_GetStudentsPerTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStudentsPerTestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnrollmentServiceServer).GetStudentsPerTest(m, &enrollmentServiceGetStudentsPerTestServer{stream})
}

type EnrollmentService_GetStudentsPerTestServer interface {
	Send(*studentpb.Student) error
	grpc.ServerStream
}

type enrollmentServiceGetStudentsPerTestServer struct {
	grpc.ServerStream
}

func (x *enrollmentServiceGetStudentsPerTestServer) Send(m *studentpb.Student) error {
	return x.ServerStream.SendMsg(m)
}

// EnrollmentService_ServiceDesc is the grpc.ServiceDesc for EnrollmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnrollmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enrollment.EnrollmentService",
	HandlerType: (*EnrollmentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Enroll",
			Handler:       _EnrollmentService_Enroll_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetStudentsPerTest",
			Handler:       _EnrollmentService_GetStudentsPerTest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/enrollmentpb/enrollment.proto",
}
