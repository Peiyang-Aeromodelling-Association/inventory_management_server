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

type updateItemRequest struct {
	QueryIdentifierCode string `json:"query_identifier_code" binding:"required"`
	IdentifierCode      string `json:"identifier_code" binding:"required"`
	Name                string `json:"name" binding:"required"`
	Holder              int32  `json:"holder" binding:"required"`
	Description         string `json:"description" binding:"omitempty"`
}

// updateItem
// @Summary Update item
// @Description Update item
// @Tags items
// @Accept json
// @Produce json
// @Param request body updateItemRequest true "update item request"
// @Success 200 {object} db.Item
// @Failure 404 {object} error "Not Found"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/update [post]
func (server *Server) updateItem(ctx *gin.Context) {
	var req updateItemRequest

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

		// 2. update the item with modifier uid
		arg := db.UpdateItemByIdentifierCodeParams{
			QueryIdentifierCode: req.QueryIdentifierCode,
			UpdateItemParams: db.UpdateItemParams{
				IdentifierCode: req.IdentifierCode,
				Name:           req.Name,
				Holder:         req.Holder,
				Modifier:       user.Uid,
				Description:    sql.NullString{String: req.Description, Valid: true},
			},
		}

		result, err = server.transaction.UpdateItemByIdentifierCodeTx(ctx, arg)

		return err
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// return OK
	ctx.JSON(http.StatusOK, result)
}

type deleteItemRequest struct {
	IdenfifierCode string `json:"identifier_code" binding:"required"`
}

// deleteItem
// @Summary Delete item
// @Description Delete item
// @Tags items
// @Accept json
// @Produce json
// @Param request body deleteItemRequest true "delete item request"
// @Success 200 {object} db.Item
// @Failure 404 {object} error "Not Found"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/delete [post]
func (server *Server) deleteItem(ctx *gin.Context) {
	var req deleteItemRequest

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

		// 2. delete the item with modifier uid
		arg := db.DeleteItemsByIdentifierCodeForUpdateParams{
			IdentifierCode: req.IdenfifierCode,
			Modifier:       user.Uid,
		}

		result, err = server.transaction.DeleteItemByIdentifierCodeTx(ctx, arg)

		return err
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// return OK
	ctx.JSON(http.StatusOK, result)
}

type getItemByIdentifierRequest struct {
	IdentifierCode string `uri:"identifier" binding:"required"`
}

// getItemByIdentifier
// @Summary Get item by identifier
// @Description Get item by identifier
// @Tags items
// @Produce json
// @Param identifier path string true "identifier code"
// @Success 200 {object} db.Item
// @Failure 404 {object} error "Not Found"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/identifier/{identifier} [get]
func (server *Server) getItemByIdentifier(ctx *gin.Context) {
	var req getItemByIdentifierRequest

	// check if the request body is valid
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	item, err := server.transaction.GetItemsByIdentifierCode(ctx, db.GetItemsByIdentifierCodeParams{
		IdentifierCode: req.IdentifierCode,
		Deleted:        false,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err)) // return 404 if item not found
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return OK
	ctx.JSON(http.StatusOK, item)
}

type checkInItemRequest struct {
	IdentifierCode string `uri:"identifier" binding:"required"`
}

// checkInItem
// @Summary Check in item
// @Description Check in item
// @Tags items
// @Produce json
// @Param identifier path string true "identifier code"
// @Success 200 {object} db.Item
// @Failure 404 {object} error "Not Found"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/checkin/{identifier} [get]
func (server *Server) checkInItem(ctx *gin.Context) {
	var req checkInItemRequest

	// check if the request body is valid
	if err := ctx.ShouldBindUri(&req); err != nil {
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

		// 2. get admin user id from database
		adminUser, err := q.GetUserByUsername(ctx, server.config.AdminUsername)
		if err != nil {
			return err
		}

		// 2. check in the item with modifier uid
		arg := db.AlterHolderByIdentifierCodeParams{
			IdentifierCode: req.IdentifierCode,
			Holder:         adminUser.Uid,
			Modifier:       user.Uid,
		}

		result, err = server.transaction.AlterHolderByIdentifierCodeTx(ctx, arg)

		return err
	})

	if err != nil {
		if err == sql.ErrNoRows || err == db.ErrItemDeleted {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// return OK
	ctx.JSON(http.StatusOK, result)
}

type checkOutItemRequest struct {
	IdentifierCode string `uri:"identifier" binding:"required"`
}

// checkOutItem
// @Summary Check out item
// @Description Check out item
// @Tags items
// @Produce json
// @Param identifier path string true "identifier code"
// @Success 200 {object} db.Item
// @Failure 404 {object} error "Not Found"
// @Failure 400 {object} error "Bad Request"
// @Failure 500 {object} error "Internal Server Error"
// @Router /items/checkout/{identifier} [get]
func (server *Server) checkOutItem(ctx *gin.Context) {
	var req checkOutItemRequest

	// check if the request body is valid
	if err := ctx.ShouldBindUri(&req); err != nil {
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

		// 2. check out the item with modifier uid
		arg := db.AlterHolderByIdentifierCodeParams{
			IdentifierCode: req.IdentifierCode,
			Holder:         user.Uid,
			Modifier:       user.Uid,
		}

		result, err = server.transaction.AlterHolderByIdentifierCodeTx(ctx, arg)

		return err
	})

	if err != nil {
		if err == sql.ErrNoRows || err == db.ErrItemDeleted {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// return OK
	ctx.JSON(http.StatusOK, result)
}
