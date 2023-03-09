-- name: CreateSchedulePassenger :one
INSERT INTO "schedule_passenger" (passenger_id, category)
VALUES ($1, $2)
RETURNING *;
-- name: ListSchedulePassengers :many
SELECT *
FROM "schedule_passenger"
ORDER BY schedule_id;
-- name: UpdatePassengerSchedule :one
UPDATE "schedule_passenger"
SET schedule_id = $2
WHERE passenger_id = $1
RETURNING *;
-- name: UpdatePassengerSeat :one
UPDATE "schedule_passenger"
SET seat = $2
WHERE passenger_id = $1
RETURNING *;
-- name: DeleteSchedulePassenger :exec
DELETE
FROM "schedule_passenger"
WHERE passenger_id = $1;