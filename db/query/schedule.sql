-- name: CreateSchedule :one
INSERT INTO "schedule" (departure_date, departure_time, pickup, drop_off)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: ListSchedules :many
SELECT *
FROM "schedule"
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateScheduleDepartureDate :one
UPDATE "schedule"
SET departure_date = $2
WHERE id = $1
RETURNING *;
-- name: UpdateScheduleDepartureTime :one
UPDATE "schedule"
SET departure_time = $2
WHERE id = $1
RETURNING *;
-- name: UpdateSchedulePickup :one
UPDATE "schedule"
SET pickup = $2
WHERE id = $1
RETURNING *;
-- name: UpdateScheduleDropOff :one
UPDATE "schedule"
SET drop_off = $2
WHERE id = $1
RETURNING *;
-- name: UpdateScheduleDriverId :one
UPDATE "schedule"
SET driver_id = $2
WHERE id = $1
RETURNING *;
-- name: UpdateSchedulePlateId :one
UPDATE "schedule"
SET plate_id = $2
WHERE id = $1
RETURNING *;
-- name: DeleteSchedule :exec
DELETE
FROM "schedule"
WHERE id = $1;