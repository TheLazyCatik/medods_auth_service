package login

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginServer struct {
	ctx          context.Context
	postgresPool *pgxpool.Pool
}

func New(ctx context.Context, postgresPool *pgxpool.Pool) *LoginServer {
	return &LoginServer{ctx, postgresPool}
}
