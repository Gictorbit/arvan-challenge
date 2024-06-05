package discountapi

import (
	"context"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ds *DiscountService) PublishEvent(ctx context.Context, req *dispb.PublishEventRequest) (*dispb.PublishEventResponse, error) {
	if e := req.ValidateAll(); e != nil {
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}
	err := ds.arvanDB.PublishEvent(ctx, req.GetEventId())
	if err != nil {
		ds.logger.Error("failed to publish event", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &dispb.PublishEventResponse{}, nil
}
