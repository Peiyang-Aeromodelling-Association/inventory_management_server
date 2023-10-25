package api

import (
	"database/sql"
	"net/http"

	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/util"
	"github.com/gin-gonic/gin"
)

// Struct to handle create user request. Activated field will be defaulted to true.
type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	Description string `json:"description" binding:"omitempty"`
}

// createUser
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param request body createUserRequest true "create user request"
// @Success 200 {object} db.User "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /users/create [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	// check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// create a new user in the database via transaction
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:    req.Username,
		Password:    hashedPassword,
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

// get from query params
type listUsersRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1,max=100"`
	Offset int32 `form:"offset" binding:"min=0"`
}

// listUsers
// @Summary List users
// @Description List users
// @Tags users
// @Produce json
// @Param limit query int true "limit"
// @Param offset query int false "offset"
// @Success 200 {array} db.User "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Failure 500 {object} error "Internal Server Error"
// @Router /users/list [get]
func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest

	// check if the request query is valid
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	users, err := server.transaction.ListUsersTx(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err)) // return 404 if user not found
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return OK
	ctx.JSON(http.StatusOK, users)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

// loginUser
// @Summary Login a user
// @Description Login a user
// @Tags users
// @Accept json
// @Produce json
// @Param request body loginUserRequest true "login user request"
// @Success 200 {object} loginUserResponse "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Failure 500 {object} error "Internal Server Error"
// @Router /users/login [post]
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	// check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get the user from the database
	user, err := server.transaction.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err)) // return 404 if user not found
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check if the password is correct
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// generate access token
	accessToken, _, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return OK
	rsp := loginUserResponse{
		AccessToken: accessToken,
		Username:    user.Username,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type updateUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=6"`
	Description string `json:"description" binding:"omitempty"`
}

// updateUser
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param request body updateUserRequest true "update user request"
// @Success 200 {object} db.User
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Failure 500 {object} error "Internal Server Error"
// @Router /users/update [post]
func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest

	// check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// update the user in the database via transaction
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UpdateUserByUsernameParams{
		UpdateUserParams: db.UpdateUserParams{
			Username:    req.Username,
			Password:    hashedPassword,
			Description: sql.NullString{String: req.Description, Valid: true},
			Activated:   true,
		},
		QueryUsername: req.Username,
	}

	user, err := server.transaction.UpdateUserByUsernameTx(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err)) // return 404 if user not found
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return OK
	ctx.JSON(http.StatusOK, user)
}
