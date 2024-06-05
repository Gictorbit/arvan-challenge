package discount

import (
	"context"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Events returns published events
func (ds *DiscountService) Events(ctx context.Context, req *dispb.EventsRequest) (*dispb.EventsResponse, error) {
	if e := req.ValidateAll(); e != nil {
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}
	events, err := ds.arvanDB.GetPublishedEvents(ctx)
	if err != nil {
		ds.logger.Error("failed to get published errors",
			zap.Error(err),
			zap.Uint32("userID", req.UserId))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &dispb.EventsResponse{
		Events: events,
	}, nil
}
