package apps

import (
	"context"
	"fmt"

	db "github.com/ij4l/foodCatalog/database/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

type AppRepository interface {
	db.Querier
}

type SQLAppRepository struct {
	*db.Queries
	db *pgx.Conn
}

func NewRepository(dbms *pgx.Conn) AppRepository {
	return &SQLAppRepository{
		db:      dbms,
		Queries: db.New(dbms),
	}
}

func (SQLCatalog *SQLAppRepository) ExecTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := SQLCatalog.db.Begin(ctx)
	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(q)
	if err != nil {

		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
