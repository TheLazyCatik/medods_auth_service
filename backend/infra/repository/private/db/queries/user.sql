-- name: CreateUser :one
INSERT INTO "users" ("email")
VALUES ($1)
RETURNING "id";


-- name: GetUserByEmail :one
SELECT
    * 
FROM
    "users"
WHERE
    "email" = $1;

-- name: GetUserByID :one
SELECT
    * 
FROM
    "users"
WHERE
    "id" = $1;