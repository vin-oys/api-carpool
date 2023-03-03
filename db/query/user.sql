-- name: CreateUser :one
INSERT INTO users (username, password, contact_number, role_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: SelectUser :one
SELECT * FROM users
WHERE username = $1;