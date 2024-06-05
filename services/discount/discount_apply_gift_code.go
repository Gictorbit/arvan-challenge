package discount

import (
	"context"
	"errors"
	disdb "github.com/gictorbit/arvan-challenge/db/discountdb"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ds *DiscountService) ApplyGiftCode(ctx context.Context, req *dispb.ApplyGiftCodeRequest) (*dispb.ApplyGiftCodeResponse, error) {
	if e := req.ValidateAll(); e != nil {
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}
	message, newBalance, err := ds.arvanDB.ApplyGiftCode(ctx, req.Phone, req.Code)
	if err != nil {
		if errors.Is(err, disdb.ErrUserNotFound) || errors.Is(err, disdb.ErrGiftCodeNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if errors.Is(err, disdb.ErrGiftCodeNotActive) || errors.Is(err, disdb.ErrGiftCodeNotActive) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if errors.Is(err, disdb.ErrGiftCodeAlreadyApplied) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		ds.logger.Error("failed to apply gift code",
			zap.String("event", req.Code),
			zap.String("phone", req.Phone),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &dispb.ApplyGiftCodeResponse{
		Message:    message,
		NewBalance: newBalance,
	}, err
}
