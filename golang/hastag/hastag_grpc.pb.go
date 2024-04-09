// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: hastag.proto

package hastag

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

const (
	HastagService_HastagCreate_FullMethodName = "/cti.hastag.v1.HastagService/HastagCreate"
)

// HastagServiceClient is the client API for HastagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HastagServiceClient interface {
	HastagCreate(ctx context.Context, in *HastagCreateRequest, opts ...grpc.CallOption) (*HastagCreateRespone, error)
}

type hastagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHastagServiceClient(cc grpc.ClientConnInterface) HastagServiceClient {
	return &hastagServiceClient{cc}
}

func (c *hastagServiceClient) HastagCreate(ctx context.Context, in *HastagCreateRequest, opts ...grpc.CallOption) (*HastagCreateRespone, error) {
	out := new(HastagCreateRespone)
	err := c.cc.Invoke(ctx, HastagService_HastagCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HastagServiceServer is the server API for HastagService service.
// All implementations must embed UnimplementedHastagServiceServer
// for forward compatibility
type HastagServiceServer interface {
	HastagCreate(context.Context, *HastagCreateRequest) (*HastagCreateRespone, error)
	mustEmbedUnimplementedHastagServiceServer()
}

// UnimplementedHastagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHastagServiceServer struct {
}

func (UnimplementedHastagServiceServer) HastagCreate(context.Context, *HastagCreateRequest) (*HastagCreateRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HastagCreate not implemented")
}
func (UnimplementedHastagServiceServer) mustEmbedUnimplementedHastagServiceServer() {}

// UnsafeHastagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HastagServiceServer will
// result in compilation errors.
type UnsafeHastagServiceServer interface {
	mustEmbedUnimplementedHastagServiceServer()
}

func RegisterHastagServiceServer(s grpc.ServiceRegistrar, srv HastagServiceServer) {
	s.RegisterService(&HastagService_ServiceDesc, srv)
}

func _HastagService_HastagCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HastagCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HastagServiceServer).HastagCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HastagService_HastagCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HastagServiceServer).HastagCreate(ctx, req.(*HastagCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HastagService_ServiceDesc is the grpc.ServiceDesc for HastagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HastagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cti.hastag.v1.HastagService",
	HandlerType: (*HastagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HastagCreate",
			Handler:    _HastagService_HastagCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hastag.proto",
}
