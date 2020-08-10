// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingServiceClient interface {
	Pong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error)
}

type pingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingServiceClient(cc grpc.ClientConnInterface) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) Pong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error) {
	out := new(Ping)
	err := c.cc.Invoke(ctx, "/ping.PingService/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
// All implementations must embed UnimplementedPingServiceServer
// for forward compatibility
type PingServiceServer interface {
	Pong(context.Context, *Ping) (*Ping, error)
	mustEmbedUnimplementedPingServiceServer()
}

// UnimplementedPingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (*UnimplementedPingServiceServer) Pong(context.Context, *Ping) (*Ping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (*UnimplementedPingServiceServer) mustEmbedUnimplementedPingServiceServer() {}

func RegisterPingServiceServer(s *grpc.Server, srv PingServiceServer) {
	s.RegisterService(&_PingService_serviceDesc, srv)
}

func _PingService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ping.PingService/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).Pong(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ping.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pong",
			Handler:    _PingService_Pong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}