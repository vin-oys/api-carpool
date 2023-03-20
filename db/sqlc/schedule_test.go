package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func MockJSON() json.RawMessage {
	randomData := map[string]any{
		"location": map[string]string{
			"pickup":  "test1",
			"dropoff": "test2",
		},
		"wait": true,
	}

	data, _ := json.Marshal(randomData)
	return data
}

func MockTime() time.Time {
	return time.Date(2023, 3, 14, 14, 30, 45, 12345, time.UTC)
}

func CreateRandomSchedule(t *testing.T) Schedule {
	departureDate, _ := time.Parse(time.DateOnly, MockTime().Format(time.DateOnly))
	departureTime, _ := time.Parse(time.TimeOnly, MockTime().Format(time.TimeOnly))

	arg := CreateScheduleParams{
		DepartureDate: departureDate,
		DepartureTime: departureTime,
		Pickup:        MockJSON(),
		DropOff:       MockJSON(),
	}

	schedule, err := testQueries.CreateSchedule(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, schedule)

	require.WithinDuration(t, arg.DepartureDate, schedule.DepartureDate, time.Second)
	require.WithinDuration(t, arg.DepartureTime, schedule.DepartureTime, time.Second)
	require.JSONEq(t, string(arg.Pickup), string(schedule.Pickup))
	require.JSONEq(t, string(arg.DropOff), string(schedule.DropOff))

	require.NotZero(t, schedule.CreatedAt)

	return schedule
}

func TestCreateSchedule(t *testing.T) {
	CreateRandomSchedule(t)
}

func TestGetSchedule(t *testing.T) {
	schedule1 := CreateRandomSchedule(t)
	schedule2, err := testQueries.GetSchedule(context.Background(), schedule1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, schedule2)

	require.Equal(t, schedule1.ID, schedule2.ID)
	require.WithinDuration(t, schedule1.DepartureDate, schedule2.DepartureDate, time.Second)
	require.WithinDuration(t, schedule1.DepartureTime, schedule2.DepartureTime, time.Second)
	require.JSONEq(t, string(schedule1.Pickup), string(schedule2.Pickup))
	require.JSONEq(t, string(schedule1.DropOff), string(schedule2.DropOff))
}

func TestDeleteSchedule(t *testing.T) {
	schedule1 := CreateRandomSchedule(t)
	err := testQueries.DeleteSchedule(context.Background(), schedule1.ID)

	require.NoError(t, err)

	schedule2, err := testQueries.GetSchedule(context.Background(), schedule1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, schedule2)
}

func TestListSchedules(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomSchedule(t)
	}

	arg := ListSchedulesParams{
		Limit:  5,
		Offset: 5,
	}

	schedules, err := testQueries.ListSchedules(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, schedules, 5)

	for _, schedule := range schedules {
		require.NotEmpty(t, schedule)
	}
}
