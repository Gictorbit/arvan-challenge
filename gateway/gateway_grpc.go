package gateway

import (
	"context"
	"fmt"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func (gws *GatewayService) GrpcGatewayServer(ctx context.Context) (*http.Server, error) {
	gwMux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(gws.GetForwardRespHeaders()),
	)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Register TrackingService gRPC service endpoint
	if gws.env.WalletEndPoint != "" {
		if e := wlpb.RegisterWalletServiceHandlerFromEndpoint(ctx, gwMux, gws.env.WalletEndPoint, opts); e != nil {
			return nil, fmt.Errorf("failed to register user: %v", e.Error())
		}
	}
	if gws.env.DiscountEndPoint != "" {
		if e := dispb.RegisterDiscountServiceHandlerFromEndpoint(ctx, gwMux, gws.env.DiscountEndPoint, opts); e != nil {
			return nil, fmt.Errorf("failed to register tracking: %v", e.Error())
		}
	}

	httpServer := &http.Server{
		Addr: gws.env.Address,
		Handler: wsproxy.WebsocketProxy(gwMux, wsproxy.WithForwardedHeaders(func(header string) bool {
			return true
		})),
	}

	return httpServer, nil
}

func (gws *GatewayService) GetForwardRespHeaders() func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	return func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
		if gws.env.DebugMode {
			gws.logger.Info("response details",
				zap.Any("response headers", w.Header()),
			)
		}
		md, ok := runtime.ServerMetadataFromContext(ctx)
		if !ok {
			return nil
		}
		if cookies, foundCookie := md.HeaderMD["set-cookie"]; foundCookie {
			for _, cookie := range cookies {
				w.Header().Add("Set-Cookie", cookie)
			}
		}
		if location, foundRedirect := md.HeaderMD["location"]; foundRedirect {
			w.Header().Set("Location", location[0])
			w.WriteHeader(http.StatusTemporaryRedirect)
		}
		return nil
	}
}
