// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apricotpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ApricotClient is the client API for Apricot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApricotClient interface {
	NewRunNumber(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RunNumberResponse, error)
	GetDefaults(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringMap, error)
	GetVars(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringMap, error)
	GetComponentConfiguration(ctx context.Context, in *ComponentRequest, opts ...grpc.CallOption) (*ComponentResponse, error)
	RawGetRecursive(ctx context.Context, in *RawGetRecursiveRequest, opts ...grpc.CallOption) (*ComponentResponse, error)
	GetRuntimeEntry(ctx context.Context, in *GetRuntimeEntryRequest, opts ...grpc.CallOption) (*ComponentResponse, error)
	SetRuntimeEntry(ctx context.Context, in *SetRuntimeEntryRequest, opts ...grpc.CallOption) (*Empty, error)
}

type apricotClient struct {
	cc grpc.ClientConnInterface
}

func NewApricotClient(cc grpc.ClientConnInterface) ApricotClient {
	return &apricotClient{cc}
}

func (c *apricotClient) NewRunNumber(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RunNumberResponse, error) {
	out := new(RunNumberResponse)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/NewRunNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apricotClient) GetDefaults(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/GetDefaults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apricotClient) GetVars(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/GetVars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apricotClient) GetComponentConfiguration(ctx context.Context, in *ComponentRequest, opts ...grpc.CallOption) (*ComponentResponse, error) {
	out := new(ComponentResponse)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/GetComponentConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apricotClient) RawGetRecursive(ctx context.Context, in *RawGetRecursiveRequest, opts ...grpc.CallOption) (*ComponentResponse, error) {
	out := new(ComponentResponse)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/RawGetRecursive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apricotClient) GetRuntimeEntry(ctx context.Context, in *GetRuntimeEntryRequest, opts ...grpc.CallOption) (*ComponentResponse, error) {
	out := new(ComponentResponse)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/GetRuntimeEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apricotClient) SetRuntimeEntry(ctx context.Context, in *SetRuntimeEntryRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/apricot.Apricot/SetRuntimeEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApricotServer is the server API for Apricot service.
// All implementations should embed UnimplementedApricotServer
// for forward compatibility
type ApricotServer interface {
	NewRunNumber(context.Context, *Empty) (*RunNumberResponse, error)
	GetDefaults(context.Context, *Empty) (*StringMap, error)
	GetVars(context.Context, *Empty) (*StringMap, error)
	GetComponentConfiguration(context.Context, *ComponentRequest) (*ComponentResponse, error)
	RawGetRecursive(context.Context, *RawGetRecursiveRequest) (*ComponentResponse, error)
	GetRuntimeEntry(context.Context, *GetRuntimeEntryRequest) (*ComponentResponse, error)
	SetRuntimeEntry(context.Context, *SetRuntimeEntryRequest) (*Empty, error)
}

// UnimplementedApricotServer should be embedded to have forward compatible implementations.
type UnimplementedApricotServer struct {
}

func (UnimplementedApricotServer) NewRunNumber(context.Context, *Empty) (*RunNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRunNumber not implemented")
}
func (UnimplementedApricotServer) GetDefaults(context.Context, *Empty) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaults not implemented")
}
func (UnimplementedApricotServer) GetVars(context.Context, *Empty) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVars not implemented")
}
func (UnimplementedApricotServer) GetComponentConfiguration(context.Context, *ComponentRequest) (*ComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentConfiguration not implemented")
}
func (UnimplementedApricotServer) RawGetRecursive(context.Context, *RawGetRecursiveRequest) (*ComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawGetRecursive not implemented")
}
func (UnimplementedApricotServer) GetRuntimeEntry(context.Context, *GetRuntimeEntryRequest) (*ComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRuntimeEntry not implemented")
}
func (UnimplementedApricotServer) SetRuntimeEntry(context.Context, *SetRuntimeEntryRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRuntimeEntry not implemented")
}

// UnsafeApricotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApricotServer will
// result in compilation errors.
type UnsafeApricotServer interface {
	mustEmbedUnimplementedApricotServer()
}

func RegisterApricotServer(s grpc.ServiceRegistrar, srv ApricotServer) {
	s.RegisterService(&_Apricot_serviceDesc, srv)
}

func _Apricot_NewRunNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).NewRunNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/NewRunNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).NewRunNumber(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apricot_GetDefaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).GetDefaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/GetDefaults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).GetDefaults(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apricot_GetVars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).GetVars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/GetVars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).GetVars(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apricot_GetComponentConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).GetComponentConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/GetComponentConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).GetComponentConfiguration(ctx, req.(*ComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apricot_RawGetRecursive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawGetRecursiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).RawGetRecursive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/RawGetRecursive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).RawGetRecursive(ctx, req.(*RawGetRecursiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apricot_GetRuntimeEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuntimeEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).GetRuntimeEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/GetRuntimeEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).GetRuntimeEntry(ctx, req.(*GetRuntimeEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apricot_SetRuntimeEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRuntimeEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApricotServer).SetRuntimeEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apricot.Apricot/SetRuntimeEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApricotServer).SetRuntimeEntry(ctx, req.(*SetRuntimeEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Apricot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apricot.Apricot",
	HandlerType: (*ApricotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewRunNumber",
			Handler:    _Apricot_NewRunNumber_Handler,
		},
		{
			MethodName: "GetDefaults",
			Handler:    _Apricot_GetDefaults_Handler,
		},
		{
			MethodName: "GetVars",
			Handler:    _Apricot_GetVars_Handler,
		},
		{
			MethodName: "GetComponentConfiguration",
			Handler:    _Apricot_GetComponentConfiguration_Handler,
		},
		{
			MethodName: "RawGetRecursive",
			Handler:    _Apricot_RawGetRecursive_Handler,
		},
		{
			MethodName: "GetRuntimeEntry",
			Handler:    _Apricot_GetRuntimeEntry_Handler,
		},
		{
			MethodName: "SetRuntimeEntry",
			Handler:    _Apricot_SetRuntimeEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/apricot.proto",
}