-- name: GetUserByEmail :one
SELECT id, email, created_at, updated_at
FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (email, password)
VALUES ($1, $2);