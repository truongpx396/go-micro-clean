package http

import (
	"net/http"
	"project/common"
	"project/internal/item/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetItemByID godoc
// @Summary      Retrieve an item by ID
// @Description  Get details of an item by its unique ID
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "Item ID"
// @Success      200   {object}  entity.Item
// @Failure      400   {object}  common.AppError
// @Failure      404   {object}  common.AppError
// @Failure      500   {object}  common.AppError
// @Router       /items/{id} [get]
func (h *itemHandler) GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	item, err := h.usecase.GetItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrEntityNotFound(entity.Item{}.TableName(), err))
		return
	}

	c.JSON(http.StatusOK,
		item,
	)
}
