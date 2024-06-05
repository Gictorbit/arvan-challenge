package discountapi

import (
	"github.com/gictorbit/arvan-challenge/db"
	"github.com/gictorbit/arvan-challenge/envconfig"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	"go.uber.org/zap"
)

type DiscountService struct {
	dispb.UnimplementedDiscountServiceServer
	logger  *zap.Logger
	arvanDB db.ArvanDBConn
	env     *envconfig.DiscountEnvConfig
}

func NewDiscountService(
	logger *zap.Logger,
	dbConn db.ArvanDBConn,
	env *envconfig.DiscountEnvConfig) *DiscountService {
	return &DiscountService{
		logger:  logger,
		arvanDB: dbConn,
		env:     env,
	}
}
