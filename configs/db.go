package configs

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mdayat/fullstack2024-test/go/repository"
)

type Db struct {
	Conn    *pgxpool.Pool
	Queries *repository.Queries
}

func NewDb(ctx context.Context, dbURL string) (Db, error) {
	conn, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return Db{}, err
	}

	db := Db{
		Conn:    conn,
		Queries: repository.New(conn),
	}

	return db, err
}
