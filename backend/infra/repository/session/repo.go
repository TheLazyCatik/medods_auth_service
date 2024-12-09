package session

import (
	"auth_service/domain/entities/session"
	utility "auth_service/infra/repository/private"
	db "auth_service/infra/repository/private/db/gen"
	repository "auth_service/interactor/ifaces/repos"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionPostgresRepo struct {
	postgresPool *pgxpool.Pool
}

func NewSessionPostgresRepo(postgresPool *pgxpool.Pool) *SessionPostgresRepo {
	return &SessionPostgresRepo{postgresPool}
}

func (r *SessionPostgresRepo) CreateSession(ctx context.Context, args repository.CreateSessionArgs) (*uuid.UUID, error) {
	queries := db.New(r.postgresPool)

	var expiresAtToDB pgtype.Timestamp

	err := expiresAtToDB.Scan(args.ExpiresAt.UTC())

	if err != nil {
		return nil, utility.TransformError(err)
	}

	sessionID, err := queries.CreateSession(ctx, db.CreateSessionParams{
		UserID:        args.UserID,
		HashToken:     args.HashToken,
		AccessTokenID: args.AccessTokenID,
		ExpiresAt:     expiresAtToDB,
	})

	if err != nil {
		return nil, utility.TransformError(err)
	}

	return &sessionID, nil
}

func (r *SessionPostgresRepo) GetSessionByUserID(ctx context.Context, userID uuid.UUID) (*session.Session, error) {
	queries := db.New(r.postgresPool)

	item, err := queries.GetSessionByUserID(ctx, userID)

	if err != nil {
		return nil, utility.TransformError(err)
	}

	return &session.Session{
		ID:            item.ID,
		UserID:        item.UserID,
		AccessTokenID: item.AccessTokenID,
		HashToken:     item.HashToken,
		ExpiresAt:     item.ExpiresAt.Time,
	}, nil
}

func (r *SessionPostgresRepo) UpdateSessionByID(ctx context.Context, id uuid.UUID, args repository.UpdateSessionArgs) error {
	queries := db.New(r.postgresPool)

	var expiresAtToDB pgtype.Timestamp

	err := expiresAtToDB.Scan(args.ExpiresAt)

	if err != nil {
		return utility.TransformError(err)
	}

	err = queries.UpdateSession(ctx, db.UpdateSessionParams{
		ID:            id,
		AccessTokenID: args.AccessTokenID,
	})

	return err
}

func (r *SessionPostgresRepo) DeleteSessionByID(ctx context.Context, id uuid.UUID) error {
	queries := db.New(r.postgresPool)

	err := queries.DeleteSessionByID(ctx, id)

	if err != nil {
		return utility.TransformError(err)
	}

	return nil
}
