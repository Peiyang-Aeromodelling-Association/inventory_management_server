package api

import (
	"fmt"

	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/token"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config      util.Config
	transaction *db.Transaction
	tokenMaker  token.Maker
	router      *gin.Engine
}

func NewServer(config util.Config, transaction *db.Transaction) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:      config,
		transaction: transaction,
		tokenMaker:  tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/create-user", server.createUser)
	router.GET("/list-users", server.listUsers)
	router.POST("/login", server.loginUser)
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
