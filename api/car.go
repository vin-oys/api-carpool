package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"net/http"
)

type createCarRequest struct {
	PlateID string `json:"plate_id" binding:"required"`
	Pax     int32  `json:"pax" binding:"required"`
}

type getCarRequest struct {
	PlateID string `json:"plate_id" binding:"required"`
}

type updateCarRequest struct {
	PlateID string `json:"plate_id" binding:"required"`
	Pax     int32  `json:"pax" binding:"required"`
}

type deleteCarRequest struct {
	PlateId string `json:"plate_id" binding:"required"`
}

func (server *Server) createCar(ctx *gin.Context) {
	var req createCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCarParams{
		PlateID: req.PlateID,
		Pax:     req.Pax,
	}

	car, err := server.store.CreateCar(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func (server *Server) getCar(ctx *gin.Context) {
	var req getCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	car, err := server.store.GetCar(ctx, req.PlateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func (server *Server) updateCar(ctx *gin.Context) {
	var req updateCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCarPaxParams{
		PlateID: req.PlateID,
		Pax:     req.Pax,
	}

	car, err := server.store.UpdateCarPax(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func (server *Server) deleteCar(ctx *gin.Context) {
	var req deleteCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCar(ctx, req.PlateId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req.PlateId+" deleted successfully")
}
