// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: wallet/v1/wallet.proto

package walletv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/arvan-challenge/protos/gen/wallet/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// WalletServiceName is the fully-qualified name of the WalletService service.
	WalletServiceName = "wallet.v1.WalletService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WalletServiceMyBalanceProcedure is the fully-qualified name of the WalletService's MyBalance RPC.
	WalletServiceMyBalanceProcedure = "/wallet.v1.WalletService/MyBalance"
	// WalletServiceTransactionsProcedure is the fully-qualified name of the WalletService's
	// Transactions RPC.
	WalletServiceTransactionsProcedure = "/wallet.v1.WalletService/Transactions"
)

// WalletServiceClient is a client for the wallet.v1.WalletService service.
type WalletServiceClient interface {
	MyBalance(context.Context, *connect.Request[v1.MyBalanceRequest]) (*connect.Response[v1.MyBalanceResponse], error)
	Transactions(context.Context, *connect.Request[v1.TransactionsRequest]) (*connect.Response[v1.TransactionsResponse], error)
}

// NewWalletServiceClient constructs a client for the wallet.v1.WalletService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWalletServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WalletServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &walletServiceClient{
		myBalance: connect.NewClient[v1.MyBalanceRequest, v1.MyBalanceResponse](
			httpClient,
			baseURL+WalletServiceMyBalanceProcedure,
			opts...,
		),
		transactions: connect.NewClient[v1.TransactionsRequest, v1.TransactionsResponse](
			httpClient,
			baseURL+WalletServiceTransactionsProcedure,
			opts...,
		),
	}
}

// walletServiceClient implements WalletServiceClient.
type walletServiceClient struct {
	myBalance    *connect.Client[v1.MyBalanceRequest, v1.MyBalanceResponse]
	transactions *connect.Client[v1.TransactionsRequest, v1.TransactionsResponse]
}

// MyBalance calls wallet.v1.WalletService.MyBalance.
func (c *walletServiceClient) MyBalance(ctx context.Context, req *connect.Request[v1.MyBalanceRequest]) (*connect.Response[v1.MyBalanceResponse], error) {
	return c.myBalance.CallUnary(ctx, req)
}

// Transactions calls wallet.v1.WalletService.Transactions.
func (c *walletServiceClient) Transactions(ctx context.Context, req *connect.Request[v1.TransactionsRequest]) (*connect.Response[v1.TransactionsResponse], error) {
	return c.transactions.CallUnary(ctx, req)
}

// WalletServiceHandler is an implementation of the wallet.v1.WalletService service.
type WalletServiceHandler interface {
	MyBalance(context.Context, *connect.Request[v1.MyBalanceRequest]) (*connect.Response[v1.MyBalanceResponse], error)
	Transactions(context.Context, *connect.Request[v1.TransactionsRequest]) (*connect.Response[v1.TransactionsResponse], error)
}

// NewWalletServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWalletServiceHandler(svc WalletServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	walletServiceMyBalanceHandler := connect.NewUnaryHandler(
		WalletServiceMyBalanceProcedure,
		svc.MyBalance,
		opts...,
	)
	walletServiceTransactionsHandler := connect.NewUnaryHandler(
		WalletServiceTransactionsProcedure,
		svc.Transactions,
		opts...,
	)
	return "/wallet.v1.WalletService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WalletServiceMyBalanceProcedure:
			walletServiceMyBalanceHandler.ServeHTTP(w, r)
		case WalletServiceTransactionsProcedure:
			walletServiceTransactionsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWalletServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWalletServiceHandler struct{}

func (UnimplementedWalletServiceHandler) MyBalance(context.Context, *connect.Request[v1.MyBalanceRequest]) (*connect.Response[v1.MyBalanceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wallet.v1.WalletService.MyBalance is not implemented"))
}

func (UnimplementedWalletServiceHandler) Transactions(context.Context, *connect.Request[v1.TransactionsRequest]) (*connect.Response[v1.TransactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wallet.v1.WalletService.Transactions is not implemented"))
}
