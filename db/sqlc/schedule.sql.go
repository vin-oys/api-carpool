// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: schedule.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO "schedule" (departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type CreateScheduleParams struct {
	DepartureDate  time.Time       `json:"departure_date"`
	DepartureTime  time.Time       `json:"departure_time"`
	Pickup         json.RawMessage `json:"pickup"`
	DropOff        json.RawMessage `json:"drop_off"`
	PickupCountry  Country         `json:"pickup_country"`
	DropOffCountry Country         `json:"drop_off_country"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule,
		arg.DepartureDate,
		arg.DepartureTime,
		arg.Pickup,
		arg.DropOff,
		arg.PickupCountry,
		arg.DropOffCountry,
	)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSchedule = `-- name: DeleteSchedule :exec
DELETE
FROM "schedule"
WHERE id = $1
`

func (q *Queries) DeleteSchedule(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSchedule, id)
	return err
}

const getSchedule = `-- name: GetSchedule :one
SELECT id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
FROM "schedule"
WHERE id = $1
`

func (q *Queries) GetSchedule(ctx context.Context, id int32) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, getSchedule, id)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSchedules = `-- name: ListSchedules :many
SELECT id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
FROM "schedule"
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListSchedulesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSchedules(ctx context.Context, arg ListSchedulesParams) ([]Schedule, error) {
	rows, err := q.db.QueryContext(ctx, listSchedules, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Schedule{}
	for rows.Next() {
		var i Schedule
		if err := rows.Scan(
			&i.ID,
			&i.DepartureDate,
			&i.DepartureTime,
			&i.Pickup,
			&i.DropOff,
			&i.PickupCountry,
			&i.DropOffCountry,
			&i.DriverID,
			&i.PlateID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateScheduleDepartureDate = `-- name: UpdateScheduleDepartureDate :one
UPDATE "schedule"
SET departure_date = $2
WHERE id = $1
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type UpdateScheduleDepartureDateParams struct {
	ID            int32     `json:"id"`
	DepartureDate time.Time `json:"departure_date"`
}

func (q *Queries) UpdateScheduleDepartureDate(ctx context.Context, arg UpdateScheduleDepartureDateParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateScheduleDepartureDate, arg.ID, arg.DepartureDate)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateScheduleDepartureTime = `-- name: UpdateScheduleDepartureTime :one
UPDATE "schedule"
SET departure_time = $2
WHERE id = $1
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type UpdateScheduleDepartureTimeParams struct {
	ID            int32     `json:"id"`
	DepartureTime time.Time `json:"departure_time"`
}

func (q *Queries) UpdateScheduleDepartureTime(ctx context.Context, arg UpdateScheduleDepartureTimeParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateScheduleDepartureTime, arg.ID, arg.DepartureTime)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateScheduleDriverId = `-- name: UpdateScheduleDriverId :one
UPDATE "schedule"
SET driver_id = $2
WHERE id = $1
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type UpdateScheduleDriverIdParams struct {
	ID       int32         `json:"id"`
	DriverID sql.NullInt32 `json:"driver_id"`
}

func (q *Queries) UpdateScheduleDriverId(ctx context.Context, arg UpdateScheduleDriverIdParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateScheduleDriverId, arg.ID, arg.DriverID)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateScheduleDropOff = `-- name: UpdateScheduleDropOff :one
UPDATE "schedule"
SET drop_off = $2
WHERE id = $1
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type UpdateScheduleDropOffParams struct {
	ID      int32           `json:"id"`
	DropOff json.RawMessage `json:"drop_off"`
}

func (q *Queries) UpdateScheduleDropOff(ctx context.Context, arg UpdateScheduleDropOffParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateScheduleDropOff, arg.ID, arg.DropOff)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSchedulePickup = `-- name: UpdateSchedulePickup :one
UPDATE "schedule"
SET pickup = $2
WHERE id = $1
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type UpdateSchedulePickupParams struct {
	ID     int32           `json:"id"`
	Pickup json.RawMessage `json:"pickup"`
}

func (q *Queries) UpdateSchedulePickup(ctx context.Context, arg UpdateSchedulePickupParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateSchedulePickup, arg.ID, arg.Pickup)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSchedulePlateId = `-- name: UpdateSchedulePlateId :one
UPDATE "schedule"
SET plate_id = $2
WHERE id = $1
RETURNING id, departure_date, departure_time, pickup, drop_off, pickup_country, drop_off_country, driver_id, plate_id, created_at, updated_at
`

type UpdateSchedulePlateIdParams struct {
	ID      int32          `json:"id"`
	PlateID sql.NullString `json:"plate_id"`
}

func (q *Queries) UpdateSchedulePlateId(ctx context.Context, arg UpdateSchedulePlateIdParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateSchedulePlateId, arg.ID, arg.PlateID)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.DepartureDate,
		&i.DepartureTime,
		&i.Pickup,
		&i.DropOff,
		&i.PickupCountry,
		&i.DropOffCountry,
		&i.DriverID,
		&i.PlateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
