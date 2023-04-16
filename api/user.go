package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"github.com/vin-oys/api-carpool/util"
)

type createUserRequest struct {
	Username      string      `json:"username" binding:"required"`
	Password      string      `json:"password" binding:"required"`
	ContactNumber string      `json:"contact_number" binding:"required"`
	Role          db.UserRole `json:"role"`
}

type getUserRequest struct {
	Username string `form:"username" binding:"required"`
}

type updateUserRequest struct {
	Username      string `json:"username" binding:"required"`
	ContactNumber string `json:"contact_number" binding:"required"`
}

type deleteUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type UserCreateResponse struct {
	ID        int32       `json:"id"`
	Username  string      `json:"username"`
	CreatedAt time.Time   `json:"created_at"`
	RoleID    db.UserRole `json:"role_id"`
}

type UserUpdateResponse struct {
	ID        int32        `json:"id"`
	Username  string       `json:"username"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	RoleID    db.UserRole  `json:"role_id"`
}

type UserResponse struct {
	ID            int32          `json:"id"`
	Username      string         `json:"username"`
	Firstname     sql.NullString `json:"firstname"`
	Lastname      sql.NullString `json:"lastname"`
	ContactNumber string         `json:"contact_number"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	RoleID        db.UserRole    `json:"role_id"`
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

	response := UserCreateResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		RoleID:    user.RoleID,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Firstname:     user.Firstname,
		Lastname:      user.Lastname,
		ContactNumber: user.ContactNumber,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		RoleID:        user.RoleID,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getUserList(ctx *gin.Context) {
	var req db.ListUsersParams
	if err := ctx.ShouldBindHeader(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userList, err := server.store.ListUsers(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	responseList := []UserResponse{}

	for _, user := range userList {
		response := UserResponse{
			ID:            user.ID,
			Username:      user.Username,
			Firstname:     user.Firstname,
			Lastname:      user.Lastname,
			ContactNumber: user.ContactNumber,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			RoleID:        user.RoleID,
		}
		responseList = append(responseList, response)
	}

	ctx.JSON(http.StatusOK, responseList)
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

	response := UserUpdateResponse{
		ID:        res.ID,
		Username:  res.Username,
		UpdatedAt: res.UpdatedAt,
		RoleID:    res.RoleID,
	}

	ctx.JSON(http.StatusOK, response)
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
