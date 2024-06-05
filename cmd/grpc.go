package main

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

func grpcServer(logger *zap.Logger, logReqs, debugMode bool) *grpc.Server {
	streamServerOptions := []grpc.StreamServerInterceptor{
		grpcTags.StreamServerInterceptor(grpcTags.WithFieldExtractor(grpcTags.CodeGenRequestFieldExtractor)),
		grpcZap.StreamServerInterceptor(logger),
		grpcZap.PayloadStreamServerInterceptor(logger, func(ctx context.Context, fullMethodName string, servingObject any) bool {
			return logReqs
		}),
		grpcRecovery.StreamServerInterceptor(grpcRecovery.WithRecoveryHandler(func(p any) (err error) {
			logger.Error("stack trace from panic " + string(debug.Stack()))
			return status.Errorf(codes.Internal, "%v", p)
		})),
		func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if debugMode {
				if debugMode {
					md, ok := metadata.FromIncomingContext(stream.Context())
					if ok {
						logger.Info("request details",
							zap.String("method", info.FullMethod),
							zap.Any("headers", md),
						)
					}
				}
			}
			return handler(srv, stream)
		},
	}

	unaryServerOptions := []grpc.UnaryServerInterceptor{
		grpcTags.UnaryServerInterceptor(grpcTags.WithFieldExtractor(grpcTags.CodeGenRequestFieldExtractor)),
		grpcZap.UnaryServerInterceptor(logger),
		grpcZap.PayloadUnaryServerInterceptor(logger, func(ctx context.Context, fullMethodName string, servingObject any) bool {
			return logReqs
		}),
		grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandler(func(p any) (err error) {
			logger.Error("stack trace from panic " + string(debug.Stack()))
			return status.Errorf(codes.Internal, "%v", p)
		})),
		func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			if debugMode {
				md, ok := metadata.FromIncomingContext(ctx)
				if ok {
					logger.Info("request details",
						zap.String("method", info.FullMethod),
						zap.Any("headers", md),
					)
				}
			}
			return handler(ctx, req)
		},
	}
	return grpc.NewServer(
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(streamServerOptions...)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(unaryServerOptions...)),
	)
}
