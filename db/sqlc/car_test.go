package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/vin-oys/api-carpool/util"
	"testing"
	"time"
)

func createRandomCar(t *testing.T) Car {
	arg := CreateCarParams{
		PlateID: util.RandomCarPlate(),
		Pax:     8,
	}
	car, err := testQueries.CreateCar(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, arg.PlateID, car.PlateID)
	require.Equal(t, arg.Pax, car.Pax)

	require.NotZero(t, car.CreatedAt)

	return car
}

func TestCreateCar(t *testing.T) {
	createRandomCar(t)
}

func TestGetCar(t *testing.T) {
	car1 := createRandomCar(t)
	car2, err := testQueries.GetCar(context.Background(), car1.PlateID)

	require.NoError(t, err)
	require.NotEmpty(t, car2)

	require.Equal(t, car1.PlateID, car2.PlateID)
	require.Equal(t, car1.Pax, car2.Pax)

	require.WithinDuration(t, car1.CreatedAt, car2.CreatedAt, time.Second)
}

func TestUpdateCarPax(t *testing.T) {
	car1 := createRandomCar(t)
	arg := UpdateCarPaxParams{
		PlateID: car1.PlateID,
		Pax:     4,
	}
	car2, err := testQueries.UpdateCarPax(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, car2)

	require.Equal(t, car1.PlateID, car2.PlateID)
	require.Equal(t, arg.Pax, car2.Pax)

	require.WithinDuration(t, car1.CreatedAt, car2.CreatedAt, time.Second)
}

func TestDeleteCar(t *testing.T) {
	car1 := createRandomCar(t)
	err := testQueries.DeleteCar(context.Background(), car1.PlateID)

	require.NoError(t, err)

	car2, err := testQueries.GetCar(context.Background(), car1.PlateID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, car2)
}

func TestListCars(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCar(t)
	}

	arg := ListCarsParams{
		Limit:  5,
		Offset: 5,
	}

	cars, err := testQueries.ListCars(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, cars, 5)

	for _, car := range cars {
		require.NotEmpty(t, car)
	}
}
