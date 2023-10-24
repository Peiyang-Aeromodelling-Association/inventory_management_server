package api

import (
	"fmt"

	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/token"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/util"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/docs"
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
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/users/create", server.createUser)
	authRoutes.POST("/users/update", server.updateUser)
	authRoutes.GET("/users/list", server.listUsers)

	authRoutes.POST("/items/create", server.createItem)
	authRoutes.POST("/items/update", server.updateItem)
	authRoutes.POST("/items/delete", server.deleteItem)
	authRoutes.GET("/items/count", server.itemsCount)
	authRoutes.GET("/items/list", server.listItems)
	authRoutes.GET("/items/identifier/:identifier", server.getItemByIdentifier)
	authRoutes.GET("/items/checkin/:identifier", server.checkInItem)
	authRoutes.GET("/items/checkout/:identifier", server.checkOutItem)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
