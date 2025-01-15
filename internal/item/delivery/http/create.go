package http

import (
	"net/http"
	"project/common"
	"project/internal/item/entity"

	"github.com/gin-gonic/gin"
)

// CreateItem godoc
// @Summary      Create a new item
// @Description  Create a new item with name, type, and image data
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        item  body      entity.ItemCreate  true  "Item data"
// @Success      201   {object}  entity.ItemIdRead
// @Failure      400   {object}  common.AppError
// @Failure      500   {object}  common.AppError
// @Router       /items [post]
func (h *itemHandler) CreateItem(c *gin.Context) {
	var item entity.ItemCreate
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	if err := h.usecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrCannotCreateEntity(item.TableName(), err))
		return
	}

	c.JSON(http.StatusCreated, entity.ItemIdRead{
		ID: item.ID,
	})
}
