package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"github.com/vin-oys/api-carpool/token"
	"github.com/vin-oys/api-carpool/util"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.Use(cors.Default())

	userRoutes := router.Group("/user")

	userRoutes.POST("/login", server.loginUser)
	userRoutes.POST("/create", server.createUser)
	userRoutes.GET("/get", server.getUser)
	userRoutes.PUT("/update", server.updateUser)
	userRoutes.DELETE("/delete", server.deleteUser)

	carRoutes := router.Group("/car")

	carRoutes.POST("/create", server.createCar)
	carRoutes.GET("/get", server.getCar)
	carRoutes.PUT("/update", server.updateCar)
	carRoutes.DELETE("/delete", server.deleteCar)

	passengerRoute := router.Group("/passenger")

	passengerRoute.POST("/", server.createSchedulePassenger)
	passengerRoute.GET("/", server.getScheduledPassenger)
	passengerRoute.GET("/list", server.listScheduledPassengers)
	passengerRoute.PUT("/schedule", server.updatePassengerSchedule)
	passengerRoute.PUT("/seat", server.updatePassengerSeat)
	passengerRoute.DELETE("/", server.deleteSchedulePassenger)

	scheduleRoutes := router.Group("/schedule")
	scheduleRoutes.POST("/create", server.createSchedule)
	scheduleRoutes.GET("/get", server.getSchedule)
	scheduleRoutes.GET("/list", server.listSchedule)
	scheduleRoutes.PUT("/update/departureDate", server.updateScheduleDepartureDate)
	scheduleRoutes.PUT("/update/departureTime", server.updateScheduleDepartureTime)
	scheduleRoutes.PUT("/update/driverId", server.updateScheduleDriverId)
	scheduleRoutes.PUT("/update/dropOff", server.updateScheduleDropOff)
	scheduleRoutes.PUT("/update/pickup", server.updateSchedulePickup)
	scheduleRoutes.PUT("/update/plateId", server.updateSchedulePlateId)
	scheduleRoutes.DELETE("/delete", server.deleteSchedule)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
