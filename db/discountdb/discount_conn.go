package discountdb

import (
	"context"
	sqlmaker "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DiscountConn interface {
	GetPgConn() *pgxpool.Pool
	GetSQLBuilder() sqlmaker.StatementBuilderType
}

var _ DiscountConn = &DiscountDB{}

type DiscountDB struct {
	postgresConn  *pgxpool.Pool
	selectBuilder sqlmaker.StatementBuilderType
}

func (ddb *DiscountDB) GetPgConn() *pgxpool.Pool {
	return ddb.postgresConn
}

func (ddb *DiscountDB) GetSQLBuilder() sqlmaker.StatementBuilderType {
	ddb.selectBuilder = sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar)
	return ddb.selectBuilder
}

func NewDiscountDB(rawConn *pgxpool.Pool) *DiscountDB {
	return &DiscountDB{
		selectBuilder: sqlmaker.StatementBuilder.PlaceholderFormat(sqlmaker.Dollar),
		postgresConn:  rawConn,
	}
}

func ConnectToDiscountDB(databaseURL string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rawConn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return rawConn, nil
}
