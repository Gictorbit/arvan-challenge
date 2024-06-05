package walletdb

import (
	"context"
	"database/sql"
	wlpb "github.com/gictorbit/arvan-challenge/protos/gen/wallet/v1"
)

const myWalletInfoQuery = `
	SELECT 
		u.balance,
		t.transaction_id,
		t.amount,
		t.description,
		EXTRACT(EPOCH FROM t.timestamp)::BIGINT AS timestamp
	FROM 
		users u
	LEFT JOIN 
		transactions t ON u.id = t.user_id
	WHERE 
		u.id = $1;
`

func (wdb *WalletDB) MyWallet(ctx context.Context, userID uint32) (*wlpb.Wallet, error) {
	rows, err := wdb.GetPgConn().Query(ctx, myWalletInfoQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wallet := &wlpb.Wallet{}
	for rows.Next() {
		var (
			balance       sql.NullFloat64
			transactionID sql.NullString
			amount        sql.NullFloat64
			description   sql.NullString
			timestamp     sql.NullInt64
		)

		if err := rows.Scan(&balance, &transactionID, &amount, &description, &timestamp); err != nil {
			return nil, err
		}

		if balance.Valid {
			wallet.Balance = balance.Float64
		}

		if transactionID.Valid {
			transaction := &wlpb.Transaction{
				Id:          transactionID.String,
				Amount:      amount.Float64,
				Description: description.String,
				Timestamp:   uint64(timestamp.Int64),
			}
			wallet.Transactions = append(wallet.Transactions, transaction)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return wallet, nil
}
