package activity

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const DB_URL = "postgresql://temporal:temporal@127.0.0.1:5432/Payment"
const INSERT_CASH_FLOW = `
INSERT INTO "CashFlow" (from_account, to_account, amount) 
VALUES ($1,$2,$3)
`
const UPDATE_FROM_ACCOUNT = `
UPDATE "Account"
SET balance = balance - $1
WHERE id = $2
`
const UPDATE_TO_ACCOUNT = `
UPDATE "Account"
SET balance = balance + $1
WHERE id = $2
`

type TransferDetails struct {
	Amount      float32
	FromAccount int
	ToAccount   int
	ReferenceID string
}

func TransferMoneyActivity(ctx context.Context, transferDetails TransferDetails) error {
	conn, err := pgx.Connect(context.Background(), DB_URL)
	if err != nil {
		panic(err)
	}

	// start a transaction
	err = conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err = conn.Exec(ctx, INSERT_CASH_FLOW,
			transferDetails.FromAccount,
			transferDetails.ToAccount,
			transferDetails.Amount,
		); err != nil {
			return err
		}
		if _, err = conn.Exec(ctx, UPDATE_FROM_ACCOUNT,
			transferDetails.Amount,
			transferDetails.FromAccount,
		); err != nil {
			return err
		}

		if _, err = conn.Exec(ctx, UPDATE_TO_ACCOUNT,
			transferDetails.Amount,
			transferDetails.ToAccount,
		); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
