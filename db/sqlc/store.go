package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

type SQLStore struct {
	connPol *pgxpool.Pool
	*Queries
}

func NewStore(connPol *pgxpool.Pool) Store {
	return &SQLStore{
		connPol: connPol,
		Queries: New(connPol),
	}
}
