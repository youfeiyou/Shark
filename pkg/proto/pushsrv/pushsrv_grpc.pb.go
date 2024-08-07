// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: pushsrv.proto

package pushsrv

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

// PushSrvServiceClient is the client API for PushSrvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushSrvServiceClient interface {
	Push(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error)
}

type pushSrvServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushSrvServiceClient(cc grpc.ClientConnInterface) PushSrvServiceClient {
	return &pushSrvServiceClient{cc}
}

func (c *pushSrvServiceClient) Push(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/pushsrv.PushSrvService/push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushSrvServiceServer is the server API for PushSrvService service.
// All implementations must embed UnimplementedPushSrvServiceServer
// for forward compatibility
type PushSrvServiceServer interface {
	Push(context.Context, *Req) (*Rsp, error)
	mustEmbedUnimplementedPushSrvServiceServer()
}

// UnimplementedPushSrvServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushSrvServiceServer struct {
}

func (UnimplementedPushSrvServiceServer) Push(context.Context, *Req) (*Rsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedPushSrvServiceServer) mustEmbedUnimplementedPushSrvServiceServer() {}

// UnsafePushSrvServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushSrvServiceServer will
// result in compilation errors.
type UnsafePushSrvServiceServer interface {
	mustEmbedUnimplementedPushSrvServiceServer()
}

func RegisterPushSrvServiceServer(s grpc.ServiceRegistrar, srv PushSrvServiceServer) {
	s.RegisterService(&PushSrvService_ServiceDesc, srv)
}

func _PushSrvService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushSrvServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pushsrv.PushSrvService/push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushSrvServiceServer).Push(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// PushSrvService_ServiceDesc is the grpc.ServiceDesc for PushSrvService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushSrvService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pushsrv.PushSrvService",
	HandlerType: (*PushSrvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "push",
			Handler:    _PushSrvService_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pushsrv.proto",
}
