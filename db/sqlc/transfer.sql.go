// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: transfer.sql

package db

import (
	"context"
	"time"
)

const createTransfer = `-- name: CreateTransfer :exec
insert into transfers
(to_account_id, from_account_id, amount, created_at)
value (?,?,?,?)
`

type CreateTransferParams struct {
	ToAccountID   int64     `json:"to_account_id"`
	FromAccountID int64     `json:"from_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg *CreateTransferParams) error {
	_, err := q.db.ExecContext(ctx, createTransfer,
		arg.ToAccountID,
		arg.FromAccountID,
		arg.Amount,
		arg.CreatedAt,
	)
	return err
}

const getTransfersByAccountId = `-- name: GetTransfersByAccountId :many
select id, to_account_id, from_account_id, amount, created_at from transfers
where (to_account_id or from_account_id) = ?
`

func (q *Queries) GetTransfersByAccountId(ctx context.Context, toAccountID int64) ([]*Transfer, error) {
	rows, err := q.db.QueryContext(ctx, getTransfersByAccountId, toAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.ToAccountID,
			&i.FromAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransferByAccountId = `-- name: UpdateTransferByAccountId :exec
update transfers
set amount = ?
where to_account_id = ?
`

type UpdateTransferByAccountIdParams struct {
	Amount      int64 `json:"amount"`
	ToAccountID int64 `json:"to_account_id"`
}

func (q *Queries) UpdateTransferByAccountId(ctx context.Context, arg *UpdateTransferByAccountIdParams) error {
	_, err := q.db.ExecContext(ctx, updateTransferByAccountId, arg.Amount, arg.ToAccountID)
	return err
}

const deleteTransferByTime = `-- name: deleteTransferByTime :exec
delete from transfers
where created_at < ?
`

func (q *Queries) deleteTransferByTime(ctx context.Context, createdAt time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteTransferByTime, createdAt)
	return err
}
