-- name: CreateUser :one
INSERT INTO users (username, password, contact_number, role_id)
VALUES ($1, $2, $3, $4) RETURNING *;