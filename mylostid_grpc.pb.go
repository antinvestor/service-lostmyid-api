// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mylostidv1

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

// LostMyIdServiceClient is the client API for LostMyIdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LostMyIdServiceClient interface {
	// Log a new Collectible request
	Collectible(ctx context.Context, in *CollectibleRequest, opts ...grpc.CallOption) (*CollectibleResponse, error)
	CollectibleList(ctx context.Context, in *RangeSpanRequest, opts ...grpc.CallOption) (LostMyIdService_CollectibleListClient, error)
	// Log a new search request
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	SearchList(ctx context.Context, in *RangeSpanRequest, opts ...grpc.CallOption) (LostMyIdService_SearchListClient, error)
	Progress(ctx context.Context, in *ProgressRequest, opts ...grpc.CallOption) (*ProgressResponse, error)
	TransactionList(ctx context.Context, in *RangeSpanRequest, opts ...grpc.CallOption) (LostMyIdService_TransactionListClient, error)
}

type lostMyIdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLostMyIdServiceClient(cc grpc.ClientConnInterface) LostMyIdServiceClient {
	return &lostMyIdServiceClient{cc}
}

func (c *lostMyIdServiceClient) Collectible(ctx context.Context, in *CollectibleRequest, opts ...grpc.CallOption) (*CollectibleResponse, error) {
	out := new(CollectibleResponse)
	err := c.cc.Invoke(ctx, "/apis.LostMyIdService/Collectible", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lostMyIdServiceClient) CollectibleList(ctx context.Context, in *RangeSpanRequest, opts ...grpc.CallOption) (LostMyIdService_CollectibleListClient, error) {
	stream, err := c.cc.NewStream(ctx, &LostMyIdService_ServiceDesc.Streams[0], "/apis.LostMyIdService/CollectibleList", opts...)
	if err != nil {
		return nil, err
	}
	x := &lostMyIdServiceCollectibleListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LostMyIdService_CollectibleListClient interface {
	Recv() (*CollectibleResponse, error)
	grpc.ClientStream
}

type lostMyIdServiceCollectibleListClient struct {
	grpc.ClientStream
}

func (x *lostMyIdServiceCollectibleListClient) Recv() (*CollectibleResponse, error) {
	m := new(CollectibleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lostMyIdServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/apis.LostMyIdService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lostMyIdServiceClient) SearchList(ctx context.Context, in *RangeSpanRequest, opts ...grpc.CallOption) (LostMyIdService_SearchListClient, error) {
	stream, err := c.cc.NewStream(ctx, &LostMyIdService_ServiceDesc.Streams[1], "/apis.LostMyIdService/SearchList", opts...)
	if err != nil {
		return nil, err
	}
	x := &lostMyIdServiceSearchListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LostMyIdService_SearchListClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type lostMyIdServiceSearchListClient struct {
	grpc.ClientStream
}

func (x *lostMyIdServiceSearchListClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lostMyIdServiceClient) Progress(ctx context.Context, in *ProgressRequest, opts ...grpc.CallOption) (*ProgressResponse, error) {
	out := new(ProgressResponse)
	err := c.cc.Invoke(ctx, "/apis.LostMyIdService/Progress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lostMyIdServiceClient) TransactionList(ctx context.Context, in *RangeSpanRequest, opts ...grpc.CallOption) (LostMyIdService_TransactionListClient, error) {
	stream, err := c.cc.NewStream(ctx, &LostMyIdService_ServiceDesc.Streams[2], "/apis.LostMyIdService/TransactionList", opts...)
	if err != nil {
		return nil, err
	}
	x := &lostMyIdServiceTransactionListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LostMyIdService_TransactionListClient interface {
	Recv() (*TransactionItem, error)
	grpc.ClientStream
}

type lostMyIdServiceTransactionListClient struct {
	grpc.ClientStream
}

func (x *lostMyIdServiceTransactionListClient) Recv() (*TransactionItem, error) {
	m := new(TransactionItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LostMyIdServiceServer is the server API for LostMyIdService service.
// All implementations must embed UnimplementedLostMyIdServiceServer
// for forward compatibility
type LostMyIdServiceServer interface {
	// Log a new Collectible request
	Collectible(context.Context, *CollectibleRequest) (*CollectibleResponse, error)
	CollectibleList(*RangeSpanRequest, LostMyIdService_CollectibleListServer) error
	// Log a new search request
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	SearchList(*RangeSpanRequest, LostMyIdService_SearchListServer) error
	Progress(context.Context, *ProgressRequest) (*ProgressResponse, error)
	TransactionList(*RangeSpanRequest, LostMyIdService_TransactionListServer) error
	mustEmbedUnimplementedLostMyIdServiceServer()
}

// UnimplementedLostMyIdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLostMyIdServiceServer struct {
}

func (UnimplementedLostMyIdServiceServer) Collectible(context.Context, *CollectibleRequest) (*CollectibleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collectible not implemented")
}
func (UnimplementedLostMyIdServiceServer) CollectibleList(*RangeSpanRequest, LostMyIdService_CollectibleListServer) error {
	return status.Errorf(codes.Unimplemented, "method CollectibleList not implemented")
}
func (UnimplementedLostMyIdServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedLostMyIdServiceServer) SearchList(*RangeSpanRequest, LostMyIdService_SearchListServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchList not implemented")
}
func (UnimplementedLostMyIdServiceServer) Progress(context.Context, *ProgressRequest) (*ProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Progress not implemented")
}
func (UnimplementedLostMyIdServiceServer) TransactionList(*RangeSpanRequest, LostMyIdService_TransactionListServer) error {
	return status.Errorf(codes.Unimplemented, "method TransactionList not implemented")
}
func (UnimplementedLostMyIdServiceServer) mustEmbedUnimplementedLostMyIdServiceServer() {}

// UnsafeLostMyIdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LostMyIdServiceServer will
// result in compilation errors.
type UnsafeLostMyIdServiceServer interface {
	mustEmbedUnimplementedLostMyIdServiceServer()
}

func RegisterLostMyIdServiceServer(s grpc.ServiceRegistrar, srv LostMyIdServiceServer) {
	s.RegisterService(&LostMyIdService_ServiceDesc, srv)
}

func _LostMyIdService_Collectible_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectibleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LostMyIdServiceServer).Collectible(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.LostMyIdService/Collectible",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LostMyIdServiceServer).Collectible(ctx, req.(*CollectibleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LostMyIdService_CollectibleList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeSpanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LostMyIdServiceServer).CollectibleList(m, &lostMyIdServiceCollectibleListServer{stream})
}

type LostMyIdService_CollectibleListServer interface {
	Send(*CollectibleResponse) error
	grpc.ServerStream
}

type lostMyIdServiceCollectibleListServer struct {
	grpc.ServerStream
}

func (x *lostMyIdServiceCollectibleListServer) Send(m *CollectibleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LostMyIdService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LostMyIdServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.LostMyIdService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LostMyIdServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LostMyIdService_SearchList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeSpanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LostMyIdServiceServer).SearchList(m, &lostMyIdServiceSearchListServer{stream})
}

type LostMyIdService_SearchListServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type lostMyIdServiceSearchListServer struct {
	grpc.ServerStream
}

func (x *lostMyIdServiceSearchListServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LostMyIdService_Progress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LostMyIdServiceServer).Progress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.LostMyIdService/Progress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LostMyIdServiceServer).Progress(ctx, req.(*ProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LostMyIdService_TransactionList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeSpanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LostMyIdServiceServer).TransactionList(m, &lostMyIdServiceTransactionListServer{stream})
}

type LostMyIdService_TransactionListServer interface {
	Send(*TransactionItem) error
	grpc.ServerStream
}

type lostMyIdServiceTransactionListServer struct {
	grpc.ServerStream
}

func (x *lostMyIdServiceTransactionListServer) Send(m *TransactionItem) error {
	return x.ServerStream.SendMsg(m)
}

// LostMyIdService_ServiceDesc is the grpc.ServiceDesc for LostMyIdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LostMyIdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.LostMyIdService",
	HandlerType: (*LostMyIdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collectible",
			Handler:    _LostMyIdService_Collectible_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _LostMyIdService_Search_Handler,
		},
		{
			MethodName: "Progress",
			Handler:    _LostMyIdService_Progress_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CollectibleList",
			Handler:       _LostMyIdService_CollectibleList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchList",
			Handler:       _LostMyIdService_SearchList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TransactionList",
			Handler:       _LostMyIdService_TransactionList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mylostid.proto",
}
