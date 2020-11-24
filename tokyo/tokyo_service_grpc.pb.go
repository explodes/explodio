// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tokyo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TokyoClient is the client API for Tokyo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokyoClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type tokyoClient struct {
	cc grpc.ClientConnInterface
}

func NewTokyoClient(cc grpc.ClientConnInterface) TokyoClient {
	return &tokyoClient{cc}
}

func (c *tokyoClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/tokyo.Tokyo/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokyoServer is the server API for Tokyo service.
// All implementations must embed UnimplementedTokyoServer
// for forward compatibility
type TokyoServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedTokyoServer()
}

// UnimplementedTokyoServer must be embedded to have forward compatible implementations.
type UnimplementedTokyoServer struct {
}

func (UnimplementedTokyoServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTokyoServer) mustEmbedUnimplementedTokyoServer() {}

// UnsafeTokyoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokyoServer will
// result in compilation errors.
type UnsafeTokyoServer interface {
	mustEmbedUnimplementedTokyoServer()
}

func RegisterTokyoServer(s grpc.ServiceRegistrar, srv TokyoServer) {
	s.RegisterService(&_Tokyo_serviceDesc, srv)
}

func _Tokyo_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokyoServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokyo.Tokyo/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokyoServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokyo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokyo.Tokyo",
	HandlerType: (*TokyoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Tokyo_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokyo_service.proto",
}