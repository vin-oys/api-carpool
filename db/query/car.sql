-- name: CreateCar :one
INSERT INTO "car" (plate_id, pax)
VALUES ($1, $2)
RETURNING *;
-- name: ListCars :many
SELECT *
FROM "car"
ORDER BY plate_id
LIMIT $1 OFFSET $2;
-- name: UpdateCar :one
UPDATE "car"
SET pax = $2
WHERE plate_id = $1
RETURNING *;
-- name: DeleteCar :exec
DELETE
FROM "car"
WHERE plate_id = $1;