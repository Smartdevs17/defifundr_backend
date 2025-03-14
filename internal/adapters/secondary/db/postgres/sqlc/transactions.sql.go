// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: transactions.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
  user_id,
  tx_hash,
  amount,
  currency,
  type,
  status
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, user_id, tx_hash, amount, currency, type, status, created_at
`

type CreateTransactionParams struct {
	UserID   uuid.UUID       `json:"user_id"`
	TxHash   string          `json:"tx_hash"`
	Amount   decimal.Decimal `json:"amount"`
	Currency string          `json:"currency"`
	Type     string          `json:"type"`
	Status   string          `json:"status"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transactions, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.UserID,
		arg.TxHash,
		arg.Amount,
		arg.Currency,
		arg.Type,
		arg.Status,
	)
	var i Transactions
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TxHash,
		&i.Amount,
		&i.Currency,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, user_id, tx_hash, amount, currency, type, status, created_at FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id uuid.UUID) (Transactions, error) {
	row := q.db.QueryRow(ctx, getTransaction, id)
	var i Transactions
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TxHash,
		&i.Amount,
		&i.Currency,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getTransactionByHash = `-- name: GetTransactionByHash :one
SELECT id, user_id, tx_hash, amount, currency, type, status, created_at FROM transactions
WHERE tx_hash = $1 LIMIT 1
`

func (q *Queries) GetTransactionByHash(ctx context.Context, txHash string) (Transactions, error) {
	row := q.db.QueryRow(ctx, getTransactionByHash, txHash)
	var i Transactions
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TxHash,
		&i.Amount,
		&i.Currency,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listUserTransactions = `-- name: ListUserTransactions :many
SELECT id, user_id, tx_hash, amount, currency, type, status, created_at FROM transactions
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type ListUserTransactionsParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListUserTransactions(ctx context.Context, arg ListUserTransactionsParams) ([]Transactions, error) {
	rows, err := q.db.Query(ctx, listUserTransactions, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transactions{}
	for rows.Next() {
		var i Transactions
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TxHash,
			&i.Amount,
			&i.Currency,
			&i.Type,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransactionStatus = `-- name: UpdateTransactionStatus :one
UPDATE transactions
SET status = $2
WHERE id = $1
RETURNING id, user_id, tx_hash, amount, currency, type, status, created_at
`

type UpdateTransactionStatusParams struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (q *Queries) UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) (Transactions, error) {
	row := q.db.QueryRow(ctx, updateTransactionStatus, arg.ID, arg.Status)
	var i Transactions
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TxHash,
		&i.Amount,
		&i.Currency,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
