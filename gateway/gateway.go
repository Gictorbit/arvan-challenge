package gateway

import (
	"context"
	"errors"
	"github.com/gictorbit/arvan-challenge/envconfig"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type GatewayService struct {
	env    *envconfig.GatewayEnvConfig
	logger *zap.Logger
}

func NewGatewayService(env *envconfig.GatewayEnvConfig, logger *zap.Logger) *GatewayService {
	return &GatewayService{
		env:    env,
		logger: logger,
	}
}

func (gws *GatewayService) Run(ctx context.Context) error {
	signalCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	grpcGatewayServer, err := gws.GrpcGatewayServer(signalCtx)
	if err != nil {
		return err
	}

	// Start gRPC Gateway server in a goroutine
	go func() {
		gws.logger.Info("start grpc gateway server", zap.String("address", gws.env.Address))
		if err := grpcGatewayServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			gws.logger.Error("gRPC Gateway server failed to start", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the servers
	<-signalCtx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown gRPC Gateway server
	if err := grpcGatewayServer.Shutdown(shutdownCtx); err != nil {
		gws.logger.Warn("gRPC Gateway server shutdown error", zap.Error(err))
	}

	gws.logger.Info("gateway servers stopped")
	return nil
}
