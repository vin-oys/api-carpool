package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomSchedulePassenger(t *testing.T) SchedulePassenger {
	user := CreateRandomUser(t)
	arg := CreateSchedulePassengerParams{
		PassengerID: user.ID,
		Category:    CategoryAdult,
	}
	passenger, err := testQueries.CreateSchedulePassenger(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, passenger)

	require.Equal(t, arg.PassengerID, passenger.PassengerID)
	require.Equal(t, arg.Category, passenger.Category)

	require.NotZero(t, passenger.ID)
	require.NotZero(t, passenger.CreatedAt)

	return passenger
}

func TestCreateSchedulePassenger(t *testing.T) {
	createRandomSchedulePassenger(t)
}

func TestGetSchedulePassenger(t *testing.T) {
	passenger1 := createRandomSchedulePassenger(t)
	passenger2, err := testQueries.GetSchedulePassenger(context.Background(), passenger1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, passenger2)

	require.Equal(t, passenger1.PassengerID, passenger2.PassengerID)
	require.Equal(t, passenger1.Category, passenger2.Category)

	require.WithinDuration(t, passenger1.CreatedAt, passenger2.CreatedAt, time.Second)
}

func TestUpdatePassengerSchedule(t *testing.T) {
	passenger1 := createRandomSchedulePassenger(t)
	schedule := CreateRandomSchedule(t)
	arg := UpdatePassengerScheduleParams{
		PassengerID: passenger1.PassengerID,
		ScheduleID: sql.NullInt32{
			Int32: schedule.ID,
			Valid: true,
		},
	}
	passenger2, err := testQueries.UpdatePassengerSchedule(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, passenger2)

	require.Equal(t, passenger1.PassengerID, passenger2.PassengerID)
	require.Equal(t, passenger1.Category, passenger2.Category)
	require.Equal(t, arg.ScheduleID, passenger2.ScheduleID)

	require.WithinDuration(t, passenger1.CreatedAt, passenger2.CreatedAt, time.Second)
}

func TestUpdatePassengerSeat(t *testing.T) {
	passenger1 := createRandomSchedulePassenger(t)
	arg := UpdatePassengerSeatParams{
		PassengerID: passenger1.PassengerID,
		Seat: sql.NullInt32{
			Int32: 2,
			Valid: true,
		},
	}
	passenger2, err := testQueries.UpdatePassengerSeat(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, passenger2)

	require.Equal(t, passenger1.PassengerID, passenger2.PassengerID)
	require.Equal(t, passenger1.Category, passenger2.Category)
	require.Equal(t, arg.Seat, passenger2.Seat)

	require.WithinDuration(t, passenger1.CreatedAt, passenger2.CreatedAt, time.Second)
}

func TestDeleteSchedulePassenger(t *testing.T) {
	passenger1 := createRandomSchedulePassenger(t)
	err := testQueries.DeleteSchedulePassenger(context.Background(), passenger1.ID)

	require.NoError(t, err)

	passenger2, err := testQueries.GetSchedulePassenger(context.Background(), passenger1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, passenger2)
}
