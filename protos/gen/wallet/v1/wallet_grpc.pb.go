// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: wallet/v1/wallet.proto

package walletv1

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
	WalletService_MyBalance_FullMethodName    = "/wallet.v1.WalletService/MyBalance"
	WalletService_Transactions_FullMethodName = "/wallet.v1.WalletService/Transactions"
)

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	MyBalance(ctx context.Context, in *MyBalanceRequest, opts ...grpc.CallOption) (*MyBalanceResponse, error)
	Transactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) MyBalance(ctx context.Context, in *MyBalanceRequest, opts ...grpc.CallOption) (*MyBalanceResponse, error) {
	out := new(MyBalanceResponse)
	err := c.cc.Invoke(ctx, WalletService_MyBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) Transactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error) {
	out := new(TransactionsResponse)
	err := c.cc.Invoke(ctx, WalletService_Transactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	MyBalance(context.Context, *MyBalanceRequest) (*MyBalanceResponse, error)
	Transactions(context.Context, *TransactionsRequest) (*TransactionsResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) MyBalance(context.Context, *MyBalanceRequest) (*MyBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyBalance not implemented")
}
func (UnimplementedWalletServiceServer) Transactions(context.Context, *TransactionsRequest) (*TransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transactions not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_MyBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).MyBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_MyBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).MyBalance(ctx, req.(*MyBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_Transactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).Transactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_Transactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).Transactions(ctx, req.(*TransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.v1.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MyBalance",
			Handler:    _WalletService_MyBalance_Handler,
		},
		{
			MethodName: "Transactions",
			Handler:    _WalletService_Transactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/v1/wallet.proto",
}
