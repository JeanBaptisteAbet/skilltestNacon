// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0
// source: liveevent.proto

package __skilltestnacon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LiveEventsService_CreateEvent_FullMethodName = "/liveevents.LiveEventsService/CreateEvent"
	LiveEventsService_UpdateEvent_FullMethodName = "/liveevents.LiveEventsService/UpdateEvent"
	LiveEventsService_DeleteEvent_FullMethodName = "/liveevents.LiveEventsService/DeleteEvent"
	LiveEventsService_ListEvents_FullMethodName  = "/liveevents.LiveEventsService/ListEvents"
)

// LiveEventsServiceClient is the client API for LiveEventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveEventsServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
}

type liveEventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveEventsServiceClient(cc grpc.ClientConnInterface) LiveEventsServiceClient {
	return &liveEventsServiceClient{cc}
}

func (c *liveEventsServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, LiveEventsService_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveEventsServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, LiveEventsService_UpdateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveEventsServiceClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, LiveEventsService_DeleteEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveEventsServiceClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, LiveEventsService_ListEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveEventsServiceServer is the server API for LiveEventsService service.
// All implementations must embed UnimplementedLiveEventsServiceServer
// for forward compatibility.
type LiveEventsServiceServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	mustEmbedUnimplementedLiveEventsServiceServer()
}

// UnimplementedLiveEventsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveEventsServiceServer struct{}

func (UnimplementedLiveEventsServiceServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedLiveEventsServiceServer) UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedLiveEventsServiceServer) DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedLiveEventsServiceServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedLiveEventsServiceServer) mustEmbedUnimplementedLiveEventsServiceServer() {}
func (UnimplementedLiveEventsServiceServer) testEmbeddedByValue()                           {}

// UnsafeLiveEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveEventsServiceServer will
// result in compilation errors.
type UnsafeLiveEventsServiceServer interface {
	mustEmbedUnimplementedLiveEventsServiceServer()
}

func RegisterLiveEventsServiceServer(s grpc.ServiceRegistrar, srv LiveEventsServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiveEventsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiveEventsService_ServiceDesc, srv)
}

func _LiveEventsService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveEventsServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveEventsService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveEventsServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveEventsService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveEventsServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveEventsService_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveEventsServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveEventsService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveEventsServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveEventsService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveEventsServiceServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveEventsService_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveEventsServiceServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveEventsService_ListEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveEventsServiceServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveEventsService_ServiceDesc is the grpc.ServiceDesc for LiveEventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveEventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "liveevents.LiveEventsService",
	HandlerType: (*LiveEventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _LiveEventsService_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _LiveEventsService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _LiveEventsService_DeleteEvent_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _LiveEventsService_ListEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liveevent.proto",
}
