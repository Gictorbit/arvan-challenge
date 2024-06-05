package walletdb

import "context"

const addTransactionQuery = `
	INSERT INTO transactions (user_id, amount, description, timestamp)VALUES ($1, $2, $3, NOW())
	RETURNING transaction_id;

`

// AddTransactions adds new transaction
func (wdb *WalletDB) AddTransactions(ctx context.Context, userID uint32, amount float64, description string) (uint32, error) {
	var transactionID uint32
	err := wdb.GetPgConn().QueryRow(ctx, addTransactionQuery, userID, amount, description).Scan(&transactionID)
	if err != nil {
		return 0, err
	}
	return transactionID, nil
}
