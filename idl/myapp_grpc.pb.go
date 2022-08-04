// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: myapp.proto

package idl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MyappClient is the client API for Myapp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyappClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (Myapp_GetImageClient, error)
}

type myappClient struct {
	cc grpc.ClientConnInterface
}

func NewMyappClient(cc grpc.ClientConnInterface) MyappClient {
	return &myappClient{cc}
}

func (c *myappClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/myapp.Myapp/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myappClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (Myapp_GetImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Myapp_ServiceDesc.Streams[0], "/myapp.Myapp/GetImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &myappGetImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Myapp_GetImageClient interface {
	Recv() (*GetImageResponse, error)
	grpc.ClientStream
}

type myappGetImageClient struct {
	grpc.ClientStream
}

func (x *myappGetImageClient) Recv() (*GetImageResponse, error) {
	m := new(GetImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyappServer is the server API for Myapp service.
// All implementations must embed UnimplementedMyappServer
// for forward compatibility
type MyappServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	GetImage(*GetImageRequest, Myapp_GetImageServer) error
	mustEmbedUnimplementedMyappServer()
}

// UnimplementedMyappServer must be embedded to have forward compatible implementations.
type UnimplementedMyappServer struct {
}

func (UnimplementedMyappServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedMyappServer) GetImage(*GetImageRequest, Myapp_GetImageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedMyappServer) mustEmbedUnimplementedMyappServer() {}

// UnsafeMyappServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyappServer will
// result in compilation errors.
type UnsafeMyappServer interface {
	mustEmbedUnimplementedMyappServer()
}

func RegisterMyappServer(s grpc.ServiceRegistrar, srv MyappServer) {
	s.RegisterService(&Myapp_ServiceDesc, srv)
}

func _Myapp_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyappServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myapp.Myapp/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyappServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Myapp_GetImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyappServer).GetImage(m, &myappGetImageServer{stream})
}

type Myapp_GetImageServer interface {
	Send(*GetImageResponse) error
	grpc.ServerStream
}

type myappGetImageServer struct {
	grpc.ServerStream
}

func (x *myappGetImageServer) Send(m *GetImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Myapp_ServiceDesc is the grpc.ServiceDesc for Myapp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Myapp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myapp.Myapp",
	HandlerType: (*MyappServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Myapp_HealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetImage",
			Handler:       _Myapp_GetImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "myapp.proto",
}
