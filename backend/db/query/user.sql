-- name: CreateUser :one
INSERT INTO "user" (
    hashed_password,
    email,
    name,
    is_admin
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserById :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user"
WHERE email = $1 LIMIT 1;