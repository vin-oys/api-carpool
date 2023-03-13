package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"github.com/vin-oys/api-carpool/util"
	"net/http"
)

type createUserRequest struct {
	Username      string      `json:"username" binding:"required"`
	Password      string      `json:"password" binding:"required"`
	ContactNumber string      `json:"contact_number" binding:"required"`
	Role          db.UserRole `json:"role"`
}

type getUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type updateUserRequest struct {
	Username      string `json:"username" binding:"required"`
	ContactNumber string `json:"contact_number" binding:"required"`
}

type deleteUserRequest struct {
	Username string `json:"username" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	arg := db.CreateUserParams{
		Username:      req.Username,
		Password:      hashedPassword,
		ContactNumber: req.ContactNumber,
		RoleID:        req.Role,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		Username:      req.Username,
		ContactNumber: req.ContactNumber,
	}

	res, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req.Username+" deleted successfully")
}
