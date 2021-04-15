// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package message

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

// GetimgClient is the client API for Getimg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetimgClient interface {
	Call(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Test(ctx context.Context, in *MTest, opts ...grpc.CallOption) (*MTest, error)
}

type getimgClient struct {
	cc grpc.ClientConnInterface
}

func NewGetimgClient(cc grpc.ClientConnInterface) GetimgClient {
	return &getimgClient{cc}
}

func (c *getimgClient) Call(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/message.Getimg/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getimgClient) Test(ctx context.Context, in *MTest, opts ...grpc.CallOption) (*MTest, error) {
	out := new(MTest)
	err := c.cc.Invoke(ctx, "/message.Getimg/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetimgServer is the server API for Getimg service.
// All implementations must embed UnimplementedGetimgServer
// for forward compatibility
type GetimgServer interface {
	Call(context.Context, *Message) (*Message, error)
	Test(context.Context, *MTest) (*MTest, error)
	mustEmbedUnimplementedGetimgServer()
}

// UnimplementedGetimgServer must be embedded to have forward compatible implementations.
type UnimplementedGetimgServer struct {
}

func (UnimplementedGetimgServer) Call(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedGetimgServer) Test(context.Context, *MTest) (*MTest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedGetimgServer) mustEmbedUnimplementedGetimgServer() {}

// UnsafeGetimgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetimgServer will
// result in compilation errors.
type UnsafeGetimgServer interface {
	mustEmbedUnimplementedGetimgServer()
}

func RegisterGetimgServer(s grpc.ServiceRegistrar, srv GetimgServer) {
	s.RegisterService(&Getimg_ServiceDesc, srv)
}

func _Getimg_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetimgServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Getimg/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetimgServer).Call(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Getimg_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MTest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetimgServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Getimg/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetimgServer).Test(ctx, req.(*MTest))
	}
	return interceptor(ctx, in, info, handler)
}

// Getimg_ServiceDesc is the grpc.ServiceDesc for Getimg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Getimg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Getimg",
	HandlerType: (*GetimgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Getimg_Call_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _Getimg_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// NewUsrClient is the client API for NewUsr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewUsrClient interface {
	Call(ctx context.Context, in *Usr, opts ...grpc.CallOption) (*Usr, error)
}

type newUsrClient struct {
	cc grpc.ClientConnInterface
}

func NewNewUsrClient(cc grpc.ClientConnInterface) NewUsrClient {
	return &newUsrClient{cc}
}

func (c *newUsrClient) Call(ctx context.Context, in *Usr, opts ...grpc.CallOption) (*Usr, error) {
	out := new(Usr)
	err := c.cc.Invoke(ctx, "/message.NewUsr/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewUsrServer is the server API for NewUsr service.
// All implementations must embed UnimplementedNewUsrServer
// for forward compatibility
type NewUsrServer interface {
	Call(context.Context, *Usr) (*Usr, error)
	mustEmbedUnimplementedNewUsrServer()
}

// UnimplementedNewUsrServer must be embedded to have forward compatible implementations.
type UnimplementedNewUsrServer struct {
}

func (UnimplementedNewUsrServer) Call(context.Context, *Usr) (*Usr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedNewUsrServer) mustEmbedUnimplementedNewUsrServer() {}

// UnsafeNewUsrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewUsrServer will
// result in compilation errors.
type UnsafeNewUsrServer interface {
	mustEmbedUnimplementedNewUsrServer()
}

func RegisterNewUsrServer(s grpc.ServiceRegistrar, srv NewUsrServer) {
	s.RegisterService(&NewUsr_ServiceDesc, srv)
}

func _NewUsr_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Usr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewUsrServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.NewUsr/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewUsrServer).Call(ctx, req.(*Usr))
	}
	return interceptor(ctx, in, info, handler)
}

// NewUsr_ServiceDesc is the grpc.ServiceDesc for NewUsr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewUsr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.NewUsr",
	HandlerType: (*NewUsrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _NewUsr_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
