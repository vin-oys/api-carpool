-- name: CreateUser :one
INSERT INTO "user" (username, password, contact_number, role_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM "user"
WHERE username = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM "user"
ORDER BY role_id
LIMIT $1 OFFSET $2;
-- name: UpdateUser :one
UPDATE "user"
SET contact_number = $2
WHERE username = $1
RETURNING *;
-- name: UpdateUserRole :one
UPDATE "user"
SET role_id = $2
WHERE username = $1
RETURNING *;
-- name: UpdateUserPassword :one
UPDATE "user"
SET password = $2
WHERE username = $1
RETURNING *;
-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE username = $1;