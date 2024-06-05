package discount

import (
	"context"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ds *DiscountService) EventUsers(ctx context.Context, req *dispb.EventUsersRequest) (*dispb.EventUsersResponse, error) {
	if e := req.ValidateAll(); e != nil {
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}
	eventUsers, err := ds.arvanDB.EventUsers(ctx, req.GetEventCode())
	if err != nil {
		ds.logger.Error("failed to get event Users", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &dispb.EventUsersResponse{
		UserUsages: eventUsers,
	}, nil
}
