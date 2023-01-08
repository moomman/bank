package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TXer
}

type TXer interface {
	TransferTo(ctx context.Context, params *TransferToParams) error
}

// 使用组合来增强功能

type SqlStore struct {
	*Queries
	db *sql.DB
}

var Dao Store

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		Queries: New(db),
		db:      db,
	}
}

func (s *SqlStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 4,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	q := s.WithTx(tx)

	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%v,rb err:%v", err, rbErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
