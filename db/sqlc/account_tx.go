package db

import (
	"context"
	"time"
)

type TransferToParams struct {
	From   int64 `json:"from"`
	To     int64 `json:"to"`
	Amount int64 `json:"amount"`
}

func (s *Store) TransferTo(ctx context.Context, params *TransferToParams) error {
	err := s.execTx(ctx, func(q *Queries) error {
		if err := q.CreateTransfer(ctx, &CreateTransferParams{
			ToAccountID:   params.To,
			FromAccountID: params.From,
			Amount:        params.Amount,
			CreatedAt:     time.Now()}); err != nil {
			return err
		}

		if err := q.CreateEntry(ctx, &CreateEntryParams{
			AccountID: params.To,
			Amount:    params.Amount,
			CreatedAt: time.Now()}); err != nil {
			return err
		}

		if err := q.CreateEntry(ctx, &CreateEntryParams{
			AccountID: params.From,
			Amount:    -params.Amount,
			CreatedAt: time.Now()}); err != nil {
			return err
		}

		if err := q.UpdateAccountById(ctx, &UpdateAccountByIdParams{
			Balance: -params.Amount,
			ID:      params.From}); err != nil {
			return err
		}

		if err := q.UpdateAccountById(ctx, &UpdateAccountByIdParams{
			Balance: params.Amount,
			ID:      params.To}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
