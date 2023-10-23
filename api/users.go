package api

import (
	"database/sql"
	"net/http"

	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Struct to handle create user request. Activated field will be defaulted to true.
type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	Description string `json:"description" binding:"omitempty"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	// check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// create a new user in the database via transaction
	arg := db.CreateUserParams{
		Username:    req.Username,
		Password:    req.Password,
		Description: sql.NullString{String: req.Description, Valid: true},
		Activated:   true,
	}

	user, err := server.transaction.CreateUserTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return OK
	ctx.JSON(http.StatusOK, user)
}
