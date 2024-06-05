package discountapi

import (
	"github.com/gictorbit/arvan-challenge/db"
	"github.com/gictorbit/arvan-challenge/envconfig"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"go.uber.org/zap"
)

type DiscountService struct {
	dispb.UnimplementedDiscountServiceServer
	logger       *zap.Logger
	arvanDB      db.ArvanDBConn
	env          *envconfig.DiscountEnvConfig
	walletClient wlpb.WalletServiceClient
}

func NewDiscountService(
	logger *zap.Logger,
	dbConn db.ArvanDBConn,
	env *envconfig.DiscountEnvConfig,
	wcli wlpb.WalletServiceClient) *DiscountService {
	return &DiscountService{
		logger:       logger,
		arvanDB:      dbConn,
		env:          env,
		walletClient: wcli,
	}
}
