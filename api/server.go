package api

import (
	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	transaction *db.Transaction
	router      *gin.Engine
}

func NewServer(transaction *db.Transaction) *Server {
	server := &Server{transaction: transaction}

	server.setupRouter()

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/create-user", server.createUser)
	// router.POST("/items", server.createItem)
	// router.GET("/items", server.listItems)
	// router.GET("/items/:id", server.getItem)
	// router.PUT("/items/:id", server.updateItem)
	// router.DELETE("/items/:id", server.deleteItem)
	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
