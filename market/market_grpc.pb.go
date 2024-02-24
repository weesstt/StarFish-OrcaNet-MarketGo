// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: market.proto

package starfish_market

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

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	// A simple RPC method to query the server for the current producer asks for specific data
	// This method takes a MarketQueryArgs message which will contain the hash/URL of the data being queried.
	// Returns message MarketQueries which contains an array of MarketAsk messages that represent current producer asks.
	MarketQuery(ctx context.Context, in *MarketQueryArgs, opts ...grpc.CallOption) (*MarketQueries, error)
	// A simple RPC method to register a producers ask for specific data retrieval.
	// Returns a MarketAsk message to confirm details of producers ask transaction.
	RegisterMarketAsk(ctx context.Context, in *MarketAskArgs, opts ...grpc.CallOption) (*MarketAsk, error)
	// A simple RPC method for a consumer to initiate a transaction with a producer
	// This method takes a MarketAsk message from the client to figure out which producer holds data.
	// The client will then hang until the transaction is accepted by a producer and a MarketDataTransfer message
	// is returned or an error is caused by a timeout. This MarketAsk message must contain the consumer's public IP.
	// Returns a MarketDataTransfer message which will contain the URL of the web server exposed on producer side so consumer can download data.
	InitiateMarketTransaction(ctx context.Context, in *MarketAsk, opts ...grpc.CallOption) (*MarketDataTransfer, error)
	// A simple RPC method for a producer to see which consumers want to receive certain data.
	// This method takes a MarketQueryArgs message to indicate which data they want to see consumer requests for.
	// Returns a MarketQueries message.
	ProducerMarketQuery(ctx context.Context, in *MarketQueryArgs, opts ...grpc.CallOption) (*MarketQueries, error)
	// A simple RPC method to accept an incoming request for specific data.
	// This method takes a MarketAsk message from the producer to identify the correct transaction.
	// A web server will be exposed by the producer to server data. This method will hang until the server
	// receives a FinalizeMarketTransaction gRPC request by the consumer which will then be forwarded to the producer.
	// or an error will occur because of a timeout.
	// Returns a Receipt message with transaction ID from the consumer.
	ProducerAcceptTransaction(ctx context.Context, in *MarketAsk, opts ...grpc.CallOption) (*Receipt, error)
	// A simple RPC method to finalize a transaction between a producer and a consumer.
	// TODO: Connect whenever OrcaNet blockchain running
	// This method takes a MarketAsk message pretaining to the transaction to finalize, then will eventually
	// initiate the transaction on the blockchain when it is running.
	// Returns a Receipt message which will contain the hash of the transaction ID on the blockchain.
	FinalizeMarketTransaction(ctx context.Context, in *MarketAsk, opts ...grpc.CallOption) (*Receipt, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) MarketQuery(ctx context.Context, in *MarketQueryArgs, opts ...grpc.CallOption) (*MarketQueries, error) {
	out := new(MarketQueries)
	err := c.cc.Invoke(ctx, "/Market/MarketQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) RegisterMarketAsk(ctx context.Context, in *MarketAskArgs, opts ...grpc.CallOption) (*MarketAsk, error) {
	out := new(MarketAsk)
	err := c.cc.Invoke(ctx, "/Market/RegisterMarketAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) InitiateMarketTransaction(ctx context.Context, in *MarketAsk, opts ...grpc.CallOption) (*MarketDataTransfer, error) {
	out := new(MarketDataTransfer)
	err := c.cc.Invoke(ctx, "/Market/InitiateMarketTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) ProducerMarketQuery(ctx context.Context, in *MarketQueryArgs, opts ...grpc.CallOption) (*MarketQueries, error) {
	out := new(MarketQueries)
	err := c.cc.Invoke(ctx, "/Market/ProducerMarketQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) ProducerAcceptTransaction(ctx context.Context, in *MarketAsk, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/Market/ProducerAcceptTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) FinalizeMarketTransaction(ctx context.Context, in *MarketAsk, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/Market/FinalizeMarketTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	// A simple RPC method to query the server for the current producer asks for specific data
	// This method takes a MarketQueryArgs message which will contain the hash/URL of the data being queried.
	// Returns message MarketQueries which contains an array of MarketAsk messages that represent current producer asks.
	MarketQuery(context.Context, *MarketQueryArgs) (*MarketQueries, error)
	// A simple RPC method to register a producers ask for specific data retrieval.
	// Returns a MarketAsk message to confirm details of producers ask transaction.
	RegisterMarketAsk(context.Context, *MarketAskArgs) (*MarketAsk, error)
	// A simple RPC method for a consumer to initiate a transaction with a producer
	// This method takes a MarketAsk message from the client to figure out which producer holds data.
	// The client will then hang until the transaction is accepted by a producer and a MarketDataTransfer message
	// is returned or an error is caused by a timeout. This MarketAsk message must contain the consumer's public IP.
	// Returns a MarketDataTransfer message which will contain the URL of the web server exposed on producer side so consumer can download data.
	InitiateMarketTransaction(context.Context, *MarketAsk) (*MarketDataTransfer, error)
	// A simple RPC method for a producer to see which consumers want to receive certain data.
	// This method takes a MarketQueryArgs message to indicate which data they want to see consumer requests for.
	// Returns a MarketQueries message.
	ProducerMarketQuery(context.Context, *MarketQueryArgs) (*MarketQueries, error)
	// A simple RPC method to accept an incoming request for specific data.
	// This method takes a MarketAsk message from the producer to identify the correct transaction.
	// A web server will be exposed by the producer to server data. This method will hang until the server
	// receives a FinalizeMarketTransaction gRPC request by the consumer which will then be forwarded to the producer.
	// or an error will occur because of a timeout.
	// Returns a Receipt message with transaction ID from the consumer.
	ProducerAcceptTransaction(context.Context, *MarketAsk) (*Receipt, error)
	// A simple RPC method to finalize a transaction between a producer and a consumer.
	// TODO: Connect whenever OrcaNet blockchain running
	// This method takes a MarketAsk message pretaining to the transaction to finalize, then will eventually
	// initiate the transaction on the blockchain when it is running.
	// Returns a Receipt message which will contain the hash of the transaction ID on the blockchain.
	FinalizeMarketTransaction(context.Context, *MarketAsk) (*Receipt, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) MarketQuery(context.Context, *MarketQueryArgs) (*MarketQueries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketQuery not implemented")
}
func (UnimplementedMarketServer) RegisterMarketAsk(context.Context, *MarketAskArgs) (*MarketAsk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMarketAsk not implemented")
}
func (UnimplementedMarketServer) InitiateMarketTransaction(context.Context, *MarketAsk) (*MarketDataTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateMarketTransaction not implemented")
}
func (UnimplementedMarketServer) ProducerMarketQuery(context.Context, *MarketQueryArgs) (*MarketQueries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProducerMarketQuery not implemented")
}
func (UnimplementedMarketServer) ProducerAcceptTransaction(context.Context, *MarketAsk) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProducerAcceptTransaction not implemented")
}
func (UnimplementedMarketServer) FinalizeMarketTransaction(context.Context, *MarketAsk) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeMarketTransaction not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_MarketQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketQueryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).MarketQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Market/MarketQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).MarketQuery(ctx, req.(*MarketQueryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_RegisterMarketAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketAskArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).RegisterMarketAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Market/RegisterMarketAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).RegisterMarketAsk(ctx, req.(*MarketAskArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_InitiateMarketTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketAsk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).InitiateMarketTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Market/InitiateMarketTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).InitiateMarketTransaction(ctx, req.(*MarketAsk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_ProducerMarketQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketQueryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).ProducerMarketQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Market/ProducerMarketQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).ProducerMarketQuery(ctx, req.(*MarketQueryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_ProducerAcceptTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketAsk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).ProducerAcceptTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Market/ProducerAcceptTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).ProducerAcceptTransaction(ctx, req.(*MarketAsk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_FinalizeMarketTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketAsk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FinalizeMarketTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Market/FinalizeMarketTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FinalizeMarketTransaction(ctx, req.(*MarketAsk))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarketQuery",
			Handler:    _Market_MarketQuery_Handler,
		},
		{
			MethodName: "RegisterMarketAsk",
			Handler:    _Market_RegisterMarketAsk_Handler,
		},
		{
			MethodName: "InitiateMarketTransaction",
			Handler:    _Market_InitiateMarketTransaction_Handler,
		},
		{
			MethodName: "ProducerMarketQuery",
			Handler:    _Market_ProducerMarketQuery_Handler,
		},
		{
			MethodName: "ProducerAcceptTransaction",
			Handler:    _Market_ProducerAcceptTransaction_Handler,
		},
		{
			MethodName: "FinalizeMarketTransaction",
			Handler:    _Market_FinalizeMarketTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market.proto",
}
