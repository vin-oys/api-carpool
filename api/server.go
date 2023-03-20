package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	userRoutes := router.Group("/user")

	userRoutes.POST("/create", server.createUser)
	userRoutes.GET("/get", server.getUser)
	userRoutes.PUT("/update", server.updateUser)
	userRoutes.DELETE("/delete", server.deleteUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
