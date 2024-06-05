package discountdb

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrGiftCodeNotFound       = errors.New("gift code not found")
	ErrGiftCodeNotActive      = errors.New("gift code is not active")
	ErrGiftCodeNotValid       = errors.New("gift code not valid or already fully used")
	ErrGiftCodeAlreadyApplied = errors.New("user has already applied this gift code")
	ErrApplyingGiftCode       = errors.New("error applying gift code")
)

type ApplyGiftCodeResult struct {
	Message    string
	NewBalance float64
	UserID     uint32
	EventCode  string
	EventTitle string
	EventDesc  string
	GiftAmount float64
}

const applyGiftCardQuery = `
	SELECT message, new_balance, user_id, event_code, event_title, event_desc, gift_amount
        FROM apply_gift_code($1, $2);
`

func (ddb *DiscountDB) ApplyGiftCode(ctx context.Context, phone, code string) (*ApplyGiftCodeResult, error) {
	var result *ApplyGiftCodeResult
	err := ddb.postgresConn.QueryRow(ctx, applyGiftCardQuery, phone, code).Scan(
		&result.Message,
		&result.NewBalance,
		&result.UserID,
		&result.EventCode,
		&result.EventTitle,
		&result.EventDesc,
		&result.GiftAmount,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle specific cases based on message content
			switch result.Message {
			case "User not found":
				return nil, ErrUserNotFound
			case "Gift code not found":
				return nil, ErrGiftCodeNotFound
			case "Gift code is not active":
				return nil, ErrGiftCodeNotActive
			case "Gift code not valid or already fully used":
				return nil, ErrGiftCodeNotValid
			case "User has already applied this gift code":
				return nil, ErrGiftCodeAlreadyApplied
			default:
				return nil, ErrApplyingGiftCode
			}
		}
		return nil, err
	}
	return result, nil
}
