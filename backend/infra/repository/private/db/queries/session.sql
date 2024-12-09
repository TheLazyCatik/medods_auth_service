-- name: GetSessionByUserID :one
SELECT *
FROM "sessions"
WHERE "user_id" = $1;

-- name: CreateSession :one
INSERT INTO "sessions" ("user_id", "hash_token", "access_token_id", "expires_at")
VALUES ($1, $2, $3, $4)
RETURNING "id";

-- name: UpdateSession :exec
UPDATE "sessions"
SET "access_token_id" = $2
WHERE "id" = $1;

-- name: DeleteSessionByID :exec 
DELETE FROM "sessions"
WHERE "id" = $1;