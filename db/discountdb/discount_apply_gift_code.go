package discountdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrGiftCodeNotFound       = errors.New("gift code not found")
	ErrGiftCodeNotActive      = errors.New("gift code is not active")
	ErrGiftCodeNotValid       = errors.New("gift code not valid or already fully used")
	ErrGiftCodeAlreadyApplied = errors.New("user has already applied this gift code")
	ErrApplyingGiftCode       = errors.New("error applying gift code")
)

func (ddb *DiscountDB) ApplyGiftCode(ctx context.Context, phone, code string) (string, float64, error) {
	var message string
	var newBalance float64
	query := "SELECT * FROM apply_gift_code($1, $2)"
	err := ddb.postgresConn.QueryRow(ctx, query, phone, code).Scan(&message, &newBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle specific cases based on message content
			switch message {
			case "User not found":
				return message, 0, ErrUserNotFound
			case "Gift code not found":
				return message, 0, ErrGiftCodeNotFound
			case "Gift code is not active":
				return message, 0, ErrGiftCodeNotActive
			case "Gift code not valid or already fully used":
				return message, 0, ErrGiftCodeNotValid
			case "User has already applied this gift code":
				return message, 0, ErrGiftCodeAlreadyApplied
			default:
				return "unknown error", 0, ErrApplyingGiftCode
			}
		}
		return "query error", 0, err
	}

	if message != "Gift code applied successfully" {
		return message, newBalance, fmt.Errorf("failed to apply gift code: %s", message)
	}

	return message, newBalance, nil
}
