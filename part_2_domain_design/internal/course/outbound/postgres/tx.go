package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
)

type txKey struct{}

func injectTx(ctx context.Context, tx bun.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) (bun.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(bun.Tx)

	return tx, ok
}

type TransactionManager struct {
	conn bun.IDB
}

func NewTransactionManager(db bun.IDB) *TransactionManager {
	return &TransactionManager{
		conn: db,
	}
}

type unitFunc func(ctx context.Context, tx bun.Tx) error

func (tm *TransactionManager) InTransaction(ctx context.Context, fn unitFunc, opts *sql.TxOptions) error {
	tx, exists := extractTx(ctx)
	if exists {
		return fn(ctx, tx)
	}

	return tm.withinTransaction(ctx, opts, fn)
}

func (tm *TransactionManager) withinTransaction(ctx context.Context, opts *sql.TxOptions, fn unitFunc) (err error) {
	tx, err := tm.conn.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			err = errors.Join(err, tx.Rollback())
		}
	}()

	if err = fn(injectTx(ctx, tx), tx); err != nil {
		return fmt.Errorf("transaction: %w", err)
	}

	if err = tx.Commit(); err != nil && !errors.Is(err, sql.ErrTxDone) {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (tm *TransactionManager) DB(ctx context.Context) bun.IDB {
	tx, exists := extractTx(ctx)
	if exists {
		return tx
	}

	return tm.conn
}
