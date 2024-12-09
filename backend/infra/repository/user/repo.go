package user

import (
	"auth_service/domain/entities/user"
	utility "auth_service/infra/repository/private"
	db "auth_service/infra/repository/private/db/gen"
	repository "auth_service/interactor/ifaces/repos"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgresRepo struct {
	postgresPool *pgxpool.Pool
}

func NewUserPostgresRepo(postgresPool *pgxpool.Pool) *UserPostgresRepo {
	return &UserPostgresRepo{postgresPool}
}

func (r *UserPostgresRepo) CreateUser(ctx context.Context, args repository.CreateUserArgs) (*uuid.UUID, error) {
	queries := db.New(r.postgresPool)

	userID, err := queries.CreateUser(
		ctx,
		args.Email,
	)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	return &userID, nil
}

func (r *UserPostgresRepo) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	queries := db.New(r.postgresPool)

	item, err := queries.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	return &user.User{
		ID:    item.ID,
		Email: item.Email,
	}, nil

}

func (r *UserPostgresRepo) GetUserByID(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	queries := db.New(r.postgresPool)

	item, err := queries.GetUserByID(ctx, userID)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	return &user.User{
		ID:    item.ID,
		Email: item.Email,
	}, nil
}
