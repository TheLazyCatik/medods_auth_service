package signup

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SignupServer struct {
	ctx          context.Context
	postgresPool *pgxpool.Pool
}

func New(ctx context.Context, postgresPool *pgxpool.Pool) *SignupServer {
	return &SignupServer{ctx, postgresPool}
}
