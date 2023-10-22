package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccoubtID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer      int64 `json:"transfer"`
	FromAccoubtID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	FromEntry     int64 `json:"from_entry"`
	ToEntry       int64 `json:"to_entry"`
}

// func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
// 	var result TransferTxResult

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error
// 		result.Transfer, err := q.CreateTrasnfer(ctx, CreateTransferParams{
// 			FromAccountID: arg.FromAccoubtID,
// 			ToAccountID:   arg.ToAccountID,
// 			Amount:        arg.Amount,
// 		},
// 		)
// 	})
// }
