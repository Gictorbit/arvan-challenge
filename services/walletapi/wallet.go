package walletapi

import (
	"github.com/gictorbit/arvan-challenge/db"
	"github.com/gictorbit/arvan-challenge/envconfig"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"go.uber.org/zap"
)

type WalletService struct {
	wlpb.UnimplementedWalletServiceServer
	logger  *zap.Logger
	arvanDB db.ArvanDBConn
	env     *envconfig.WalletEnvConfig
}

func NewWalletService(
	logger *zap.Logger,
	dbConn db.ArvanDBConn,
	env *envconfig.WalletEnvConfig) *WalletService {
	return &WalletService{
		logger:  logger,
		arvanDB: dbConn,
		env:     env,
	}
}
