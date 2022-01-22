package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/drkennetz/codingchallenge/db/sqlc"
)

// Server serves HTTP requests for our coding challenge site
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	// need to add update and delete account routes

	server.router = router
	return server

}

// Start runs the HTTP server on a specific address to start listening for http request
func (server *Server) Start(address string) error {
	// router is private, so it cannot be accessed outside of this API package
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}