package utility

import (
	prjerror "auth_service/interactor/error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func TransformError(err error) error {
	if err == pgx.ErrNoRows {
		return &prjerror.NotFoundError{}
	} else {
		return err
	}
}

func TransformPgTypeUUIDToUUID(pgUUID pgtype.UUID) *uuid.UUID {
	var result *uuid.UUID

	if pgUUID.Valid {
		uuidGoogle := uuid.UUID(pgUUID.Bytes)
		result = &uuidGoogle
	}

	return result
}
