package walletapi

import (
	"context"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ws *WalletService) AddTransaction(ctx context.Context, req *wlpb.AddTransactionRequest) (*wlpb.AddTransactionResponse, error) {
	if e := req.ValidateAll(); e != nil {
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}
	tid, err := ws.arvanDB.AddTransactions(ctx, req.GetUserId(), req.GetAmount(), req.GetDescription())
	if err != nil {
		ws.logger.Error("failed to add transaction",
			zap.Error(err),
			zap.Uint32("userID", req.GetUserId()),
		)
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &wlpb.AddTransactionResponse{
		TransactionId: tid,
	}, nil
}
