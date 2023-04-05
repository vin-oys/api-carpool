package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"net/http"
	"time"
)

type createScheduleRequest struct {
	DepartureDate  time.Time       `json:"departure_date" binding:"required"`
	DepartureTime  time.Time       `json:"departure_time" binding:"required"`
	PickUp         json.RawMessage `json:"pickup" binding:"required"`
	DropOff        json.RawMessage `json:"drop_off" binding:"required"`
	PickUpCountry  db.Country      `json:"pick_up_country" binding:"required"`
	DropOffCountry db.Country      `json:"pick_up_country" binding:"required"`
}

type getScheduleRequest struct {
	ID int32 `json:"id" binding:"required"`
}

type updateScheduleRequest struct {
	ID            int32           `json:"id" binding:"required"`
	DepartureDate time.Time       `json:"departure_date"`
	DepartureTime time.Time       `json:"departure_time" `
	PickUp        json.RawMessage `json:"pickup"`
	DropOff       json.RawMessage `json:"drop_off"`
	DriverID      int32           `json:"driver_id"`
	PlateID       string          `json:"plate_id"`
}

type deleteScheduleRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) createSchedule(ctx *gin.Context) {
	var req createScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateScheduleParams{
		DepartureDate:  req.DepartureDate,
		DepartureTime:  req.DepartureTime,
		Pickup:         req.PickUp,
		DropOff:        req.DropOff,
		PickupCountry:  req.PickUpCountry,
		DropOffCountry: req.DropOffCountry,
	}

	schedule, err := server.store.CreateSchedule(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, schedule)
}

func (server *Server) getSchedule(ctx *gin.Context) {
	var req getScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	schedule, err := server.store.GetSchedule(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedule)
}

func (server *Server) updateScheduleDepartureDate(ctx *gin.Context) {
	var req updateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateScheduleDepartureDateParams{
		ID:            req.ID,
		DepartureDate: req.DepartureDate,
	}

	res, err := server.store.UpdateScheduleDepartureDate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateScheduleDepartureTime(ctx *gin.Context) {
	var req updateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateScheduleDepartureTimeParams{
		ID:            req.ID,
		DepartureTime: req.DepartureTime,
	}

	res, err := server.store.UpdateScheduleDepartureTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateScheduleDriverId(ctx *gin.Context) {
	var req updateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateScheduleDriverIdParams{
		ID: req.ID,
		DriverID: sql.NullInt32{
			req.DriverID,
			true,
		},
	}

	res, err := server.store.UpdateScheduleDriverId(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateScheduleDropOff(ctx *gin.Context) {
	var req updateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateScheduleDropOffParams{
		ID:      req.ID,
		DropOff: req.DropOff,
	}

	res, err := server.store.UpdateScheduleDropOff(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateSchedulePickup(ctx *gin.Context) {
	var req updateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSchedulePickupParams{
		ID:     req.ID,
		Pickup: req.PickUp,
	}

	res, err := server.store.UpdateSchedulePickup(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateSchedulePlateId(ctx *gin.Context) {
	var req updateScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSchedulePlateIdParams{
		ID: req.ID,
		PlateID: sql.NullString{
			req.PlateID,
			true,
		},
	}

	res, err := server.store.UpdateSchedulePlateId(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) deleteSchedule(ctx *gin.Context) {
	var req deleteScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetSchedule(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteSchedule(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, string(req.ID)+" deleted successfully")
}
