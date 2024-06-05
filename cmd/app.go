package main

import (
	"fmt"
	arvandb "github.com/gictorbit/arvan-challenge/db"
	"github.com/gictorbit/arvan-challenge/envconfig"
	"github.com/gictorbit/arvan-challenge/gateway"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"github.com/gictorbit/arvan-challenge/services/discountapi"
	"github.com/gictorbit/arvan-challenge/services/walletapi"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	app := &cli.App{
		Name:  "arvan",
		Usage: "arvan challenge",
		Commands: []*cli.Command{
			{
				Name:  "gateway",
				Usage: "starts gateway server",
				Action: func(ctx *cli.Context) error {
					gatewayEnv, err := envconfig.ReadGatewayEnvConfig()
					if err != nil {
						return err
					}
					gw := gateway.NewGatewayService(gatewayEnv, logger)
					return gw.Run(ctx.Context)
				},
			},
			{
				Name:  "wallet",
				Usage: "starts wallet api",
				Action: func(ctx *cli.Context) error {
					walletEnv, err := envconfig.ReadWalletEnvironment()
					if err != nil {
						return err
					}
					loggerConfig := zap.NewProductionConfig()
					if walletEnv.DebugMode {
						loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
					}
					logger, err := loggerConfig.Build()
					if err != nil {
						return err
					}
					server := grpcServer(logger, walletEnv.LogRequests, walletEnv.DebugMode)
					reflection.Register(server)
					lis, err := net.Listen("tcp", walletEnv.Address)
					if err != nil {
						return fmt.Errorf("faild to make listen address:%v", err)
					}

					walletdb, err := arvandb.NewArvanDB(walletEnv.WalletDatabase, walletEnv.WalletDatabase)
					if err != nil {
						return err
					}
					walletSrv := walletapi.NewWalletService(logger, walletdb, walletEnv)
					wlpb.RegisterWalletServiceServer(server, walletSrv)
					go func() {
						logger.Info("Server running ",
							zap.String("host", walletEnv.Host),
							zap.Uint("port", walletEnv.Port),
						)
						if err := server.Serve(lis); err != nil {
							logger.Fatal("Failed to serve",
								zap.Error(err))
							return
						}
					}()

					sigs := make(chan os.Signal, 1)
					signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
					<-sigs
					server.Stop()
					return nil
				},
			},
			{
				Name:  "discount",
				Usage: "starts discount api",
				Action: func(ctx *cli.Context) error {
					dsenv, err := envconfig.ReadDiscountEnvironment()
					if err != nil {
						return err
					}
					loggerConfig := zap.NewProductionConfig()
					if dsenv.DebugMode {
						loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
					}
					logger, err := loggerConfig.Build()
					if err != nil {
						return err
					}
					server := grpcServer(logger, dsenv.LogRequests, dsenv.DebugMode)
					reflection.Register(server)
					lis, err := net.Listen("tcp", dsenv.Address)
					if err != nil {
						return fmt.Errorf("faild to make listen address:%v", err)
					}

					walletdb, err := arvandb.NewArvanDB(dsenv.DiscountDatabase, dsenv.DiscountDatabase)
					if err != nil {
						return err
					}
					discountSrv := discountapi.NewDiscountService(logger, walletdb, dsenv)
					dispb.RegisterDiscountServiceServer(server, discountSrv)
					go func() {
						logger.Info("Server running ",
							zap.String("host", dsenv.Host),
							zap.Uint("port", dsenv.Port),
						)
						if err := server.Serve(lis); err != nil {
							logger.Fatal("Failed to serve",
								zap.Error(err))
							return
						}
					}()

					sigs := make(chan os.Signal, 1)
					signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
					<-sigs
					server.Stop()
					return nil
				},
			},
		},
	}

	if e := app.Run(os.Args); e != nil {
		logger.Error("failed to run app", zap.Error(e))
	}
}
