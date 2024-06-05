package walletdb

import (
	"context"
	sqlmaker "github.com/Masterminds/squirrel"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type WalletConn interface {
	GetPgConn() *pgxpool.Pool
	GetSQLBuilder() sqlmaker.StatementBuilderType
	MyWallet(ctx context.Context, userID uint32) (*wlpb.Wallet, error)
	AddTransactions(ctx context.Context, userID uint32, amount float64, description string) (uint32, error)
}

var _ WalletConn = &WalletDB{}

type WalletDB struct {
	postgresConn  *pgxpool.Pool
	selectBuilder sqlmaker.StatementBuilderType
}

func (wdb *WalletDB) GetPgConn() *pgxpool.Pool {
	return wdb.postgresConn
}

func (wdb *WalletDB) GetSQLBuilder() sqlmaker.StatementBuilderType {
	wdb.selectBuilder = sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar)
	return wdb.selectBuilder
}

func NewWalletDB(rawConn *pgxpool.Pool) *WalletDB {
	return &WalletDB{
		selectBuilder: sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar),
		postgresConn:  rawConn,
	}
}

func ConnectToWalletDB(databaseURL string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rawConn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return rawConn, nil
}
