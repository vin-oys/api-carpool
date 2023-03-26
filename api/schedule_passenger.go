package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"net/http"
)

type createSchedulePassengerRequest struct {
	PassengerID int32       `json:"passenger_id" binding:"required"`
	Category    db.Category `json:"category" binding:"required"`
}

type deleteSchedulePassengerRequest struct {
	ID int32 `json:"id" binding:"required"`
}

type getSchedulePassengerRequest struct {
	ID int32 `json:"id" binding:"required"`
}

type listSchedulePassengersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type updatePassengerSchedule struct {
	PassengerID int32 `json:"passenger_id" binding:"required"`
	ScheduleID  int32 `json:"schedule_id" binding:"required"`
}

type updatePassengerSeat struct {
	PassengerID int32 `json:"passenger_id" binding:"required"`
	Seat        int32 `json:"seat" binding:"required"`
}

func (server *Server) createSchedulePassenger(ctx *gin.Context) {
	var req createSchedulePassengerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSchedulePassengerParams{
		PassengerID: req.PassengerID,
		Category:    req.Category,
	}

	schedulePassenger, err := server.store.CreateSchedulePassenger(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulePassenger)
}

func (server *Server) deleteSchedulePassenger(ctx *gin.Context) {
	var req deleteSchedulePassengerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetSchedulePassenger(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteSchedulePassenger(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	deleteMessage := fmt.Sprintf("%v deleted successfully", req.ID)

	ctx.JSON(http.StatusOK, deleteMessage)
}

func (server *Server) getScheduledPassenger(ctx *gin.Context) {
	var req getSchedulePassengerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	schedulePassenger, err := server.store.GetSchedulePassenger(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulePassenger)
}

func (server *Server) listScheduledPassengers(ctx *gin.Context) {
	var req listSchedulePassengersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSchedulePassengersParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	schedulePassengers, err := server.store.ListSchedulePassengers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulePassengers)
}

func (server *Server) updatePassengerSchedule(ctx *gin.Context) {
	var req updatePassengerSchedule
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePassengerScheduleParams{
		PassengerID: req.PassengerID,
		ScheduleID: sql.NullInt32{
			Int32: req.ScheduleID,
			Valid: true,
		},
	}

	schedulePassenger, err := server.store.UpdatePassengerSchedule(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulePassenger)
}

func (server *Server) updatePassengerSeat(ctx *gin.Context) {
	var req updatePassengerSeat
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePassengerSeatParams{
		PassengerID: req.PassengerID,
		Seat: sql.NullInt32{
			Int32: req.Seat,
			Valid: true,
		},
	}

	schedulePassenger, err := server.store.UpdatePassengerSeat(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulePassenger)
}
