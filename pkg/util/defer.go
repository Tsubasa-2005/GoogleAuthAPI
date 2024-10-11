package util

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/taxio/errors"
)

func DeferRollback(ctx context.Context, tx pgx.Tx) {
	if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
		log.Fatal(ctx, err)
	}
}
