package api

import (
	"database/sql"
	"net/http"

	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
	"github.com/Peiyang-Aeromodelling-Association/inventory_management_server/token"
	"github.com/gin-gonic/gin"
)

type listItemRequest struct {
	Limit  int32 `json:"limit" binding:"min=1,max=100"`
	Offset int32 `json:"offset" binding:"min=0"`
}

// listItem
// @Summary List items
// @Description List items
// @Tags items
// @Accept json
// @Produce json
// @Param request body listItemRequest true "list item request"
// @Success 200 {array} db.Item "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/list [post]
func (server *Server) listItems(ctx *gin.Context) {
	var req listItemRequest

	// check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListItemParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	// get the user from the database
	items, err := server.transaction.ListItem(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return OK
	ctx.JSON(http.StatusOK, items)
}

// itemsCount
// @Summary Count items
// @Description Count items
// @Tags items
// @Success 200 {integer} integer "OK"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/count [get]
func (server *Server) itemsCount(ctx *gin.Context) {
	count, err := server.transaction.CountItems(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, count)
}

type createItemRequest struct {
	IdentifierCode string `json:"identifier_code" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Holder         int32  `json:"holder" binding:"required"`
	Description    string `json:"description" binding:"omitempty"`
}

// createItem
// @Summary Create item
// @Description Create item
// @Tags items
// @Accept json
// @Produce json
// @Param request body createItemRequest true "create item request"
// @Success 200 {object} db.Item "OK"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/create [post]
func (server *Server) createItem(ctx *gin.Context) {
	var req createItemRequest

	// check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// transaction
	var result db.Item
	err := server.transaction.ExecTx(ctx, func(q *db.Queries) error {
		// 1. get the user from the database
		user, err := q.GetUserByUsername(ctx, authPayload.Username)
		if err != nil {
			return err
		}

		// 2. create the item with modifier uid
		arg := db.CreateItemParams{
			IdentifierCode: req.IdentifierCode,
			Name:           req.Name,
			Holder:         req.Holder,
			Modifier:       user.Uid,
			Description:    sql.NullString{String: req.Description, Valid: true},
		}

		result, err = server.transaction.CreateItemTx(ctx, arg)

		return err
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// return OK
	ctx.JSON(http.StatusOK, result)
}
