package api

import (
	"database/sql"
	"net/http"

	db "github.com/Peiyang-Aeromodelling-Association/inventory_management_server/db/sqlc"
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
// @Router /list-item [post]
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
