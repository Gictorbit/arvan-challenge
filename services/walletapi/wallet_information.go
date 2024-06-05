package walletapi

import (
	"context"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ws *WalletService) MyWallet(ctx context.Context, req *wlpb.MyWalletRequest) (*wlpb.MyWalletResponse, error) {
	if e := req.ValidateAll(); e != nil {
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}
	wallet, err := ws.arvanDB.MyWallet(ctx, req.GetUserId())
	if err != nil {
		ws.logger.Error("get wallet info failed",
			zap.Error(err),
			zap.Uint32("userID", req.GetUserId()),
		)
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &wlpb.MyWalletResponse{
		Wallet: wallet,
	}, nil
}
