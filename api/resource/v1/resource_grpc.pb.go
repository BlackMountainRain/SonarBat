// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: resource/v1/resource.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Resource_HealthCheck_FullMethodName   = "/api.resource.v1.Resource/HealthCheck"
	Resource_CreateHost_FullMethodName    = "/api.resource.v1.Resource/CreateHost"
	Resource_UpdateHost_FullMethodName    = "/api.resource.v1.Resource/UpdateHost"
	Resource_OverwriteHost_FullMethodName = "/api.resource.v1.Resource/OverwriteHost"
	Resource_DeleteHost_FullMethodName    = "/api.resource.v1.Resource/DeleteHost"
	Resource_GetHost_FullMethodName       = "/api.resource.v1.Resource/GetHost"
	Resource_GetHosts_FullMethodName      = "/api.resource.v1.Resource/GetHosts"
)

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceClient interface {
	HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error)
	CreateHost(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostReply, error)
	UpdateHost(ctx context.Context, in *UpdateHostRequest, opts ...grpc.CallOption) (*UpdateHostReply, error)
	OverwriteHost(ctx context.Context, in *OverwriteHostRequest, opts ...grpc.CallOption) (*OverwriteHostReply, error)
	DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostReply, error)
	GetHost(ctx context.Context, in *GetHostRequest, opts ...grpc.CallOption) (*GetHostReply, error)
	GetHosts(ctx context.Context, in *GetHostsRequest, opts ...grpc.CallOption) (*GetHostsReply, error)
}

type resourceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceClient(cc grpc.ClientConnInterface) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthReply)
	err := c.cc.Invoke(ctx, Resource_HealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) CreateHost(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateHostReply)
	err := c.cc.Invoke(ctx, Resource_CreateHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) UpdateHost(ctx context.Context, in *UpdateHostRequest, opts ...grpc.CallOption) (*UpdateHostReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateHostReply)
	err := c.cc.Invoke(ctx, Resource_UpdateHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) OverwriteHost(ctx context.Context, in *OverwriteHostRequest, opts ...grpc.CallOption) (*OverwriteHostReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OverwriteHostReply)
	err := c.cc.Invoke(ctx, Resource_OverwriteHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteHostReply)
	err := c.cc.Invoke(ctx, Resource_DeleteHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) GetHost(ctx context.Context, in *GetHostRequest, opts ...grpc.CallOption) (*GetHostReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHostReply)
	err := c.cc.Invoke(ctx, Resource_GetHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) GetHosts(ctx context.Context, in *GetHostsRequest, opts ...grpc.CallOption) (*GetHostsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHostsReply)
	err := c.cc.Invoke(ctx, Resource_GetHosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServer is the server API for Resource service.
// All implementations must embed UnimplementedResourceServer
// for forward compatibility
type ResourceServer interface {
	HealthCheck(context.Context, *HealthRequest) (*HealthReply, error)
	CreateHost(context.Context, *CreateHostRequest) (*CreateHostReply, error)
	UpdateHost(context.Context, *UpdateHostRequest) (*UpdateHostReply, error)
	OverwriteHost(context.Context, *OverwriteHostRequest) (*OverwriteHostReply, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*DeleteHostReply, error)
	GetHost(context.Context, *GetHostRequest) (*GetHostReply, error)
	GetHosts(context.Context, *GetHostsRequest) (*GetHostsReply, error)
	mustEmbedUnimplementedResourceServer()
}

// UnimplementedResourceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (UnimplementedResourceServer) HealthCheck(context.Context, *HealthRequest) (*HealthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedResourceServer) CreateHost(context.Context, *CreateHostRequest) (*CreateHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHost not implemented")
}
func (UnimplementedResourceServer) UpdateHost(context.Context, *UpdateHostRequest) (*UpdateHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHost not implemented")
}
func (UnimplementedResourceServer) OverwriteHost(context.Context, *OverwriteHostRequest) (*OverwriteHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverwriteHost not implemented")
}
func (UnimplementedResourceServer) DeleteHost(context.Context, *DeleteHostRequest) (*DeleteHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHost not implemented")
}
func (UnimplementedResourceServer) GetHost(context.Context, *GetHostRequest) (*GetHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHost not implemented")
}
func (UnimplementedResourceServer) GetHosts(context.Context, *GetHostsRequest) (*GetHostsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHosts not implemented")
}
func (UnimplementedResourceServer) mustEmbedUnimplementedResourceServer() {}

// UnsafeResourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServer will
// result in compilation errors.
type UnsafeResourceServer interface {
	mustEmbedUnimplementedResourceServer()
}

func RegisterResourceServer(s grpc.ServiceRegistrar, srv ResourceServer) {
	s.RegisterService(&Resource_ServiceDesc, srv)
}

func _Resource_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).HealthCheck(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_CreateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).CreateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_CreateHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).CreateHost(ctx, req.(*CreateHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_UpdateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).UpdateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_UpdateHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).UpdateHost(ctx, req.(*UpdateHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_OverwriteHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverwriteHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).OverwriteHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_OverwriteHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).OverwriteHost(ctx, req.(*OverwriteHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_DeleteHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteHost(ctx, req.(*DeleteHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_GetHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).GetHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_GetHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).GetHost(ctx, req.(*GetHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_GetHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).GetHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_GetHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).GetHosts(ctx, req.(*GetHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resource_ServiceDesc is the grpc.ServiceDesc for Resource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.resource.v1.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Resource_HealthCheck_Handler,
		},
		{
			MethodName: "CreateHost",
			Handler:    _Resource_CreateHost_Handler,
		},
		{
			MethodName: "UpdateHost",
			Handler:    _Resource_UpdateHost_Handler,
		},
		{
			MethodName: "OverwriteHost",
			Handler:    _Resource_OverwriteHost_Handler,
		},
		{
			MethodName: "DeleteHost",
			Handler:    _Resource_DeleteHost_Handler,
		},
		{
			MethodName: "GetHost",
			Handler:    _Resource_GetHost_Handler,
		},
		{
			MethodName: "GetHosts",
			Handler:    _Resource_GetHosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource/v1/resource.proto",
}
