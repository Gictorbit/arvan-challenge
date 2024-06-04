package db

import (
	sqlmaker "github.com/Masterminds/squirrel"
	"github.com/gictorbit/arvan-challenge/db/discountdb"
	"github.com/gictorbit/arvan-challenge/db/walletdb"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate mockgen -source=$GOFILE -destination=mock_db/conn.go -package=$GOPACKAG
type ArvanDBConn interface {
	GetPgConn() *pgxpool.Pool
	GetSQLBuilder() sqlmaker.StatementBuilderType
	walletdb.WalletConn
	discountdb.DiscountConn
}

var _ ArvanDBConn = &ArvanDatabase{}

type ArvanDatabase struct {
	pgConn *pgxpool.Pool
	*walletdb.WalletDB
	*discountdb.DiscountDB
	selectBuilder sqlmaker.StatementBuilderType
}

func (tdb *ArvanDatabase) GetPgConn() *pgxpool.Pool {
	return tdb.pgConn
}

func (tdb *ArvanDatabase) GetSQLBuilder() sqlmaker.StatementBuilderType {
	tdb.selectBuilder = sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar)
	return tdb.selectBuilder
}

func NewArvanDB(walletDbURL, discountDbURL string) (*ArvanDatabase, error) {
	walletDbConn, err := walletdb.ConnectToWalletDB(walletDbURL)
	if err != nil {
		return nil, err
	}
	discountDbConn, err := discountdb.ConnectToDiscountDB(discountDbURL)
	if err != nil {
		return nil, err
	}
	return &ArvanDatabase{
		WalletDB:      walletdb.NewWalletDB(walletDbConn),
		DiscountDB:    discountdb.NewDiscountDB(discountDbConn),
		selectBuilder: sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar),
	}, nil
}
