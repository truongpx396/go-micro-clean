package http

import (
	"net/http"
	"project/common"
	"project/modules/item/domain/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteItem godoc
// @Summary      Delete an item
// @Description  Delete an item resource by its ID
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "Item ID"
// @Success      204   {object}  nil
// @Failure      400   {object}  common.AppError
// @Failure      404   {object}  common.AppError
// @Failure      500   {object}  common.AppError
// @Router       /items/{id} [delete]
func (h *itemHandler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	if _, err := h.usecase.GetItemByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, common.ErrEntityNotFound(models.Item{}.TableName(), err))
		return
	}

	if err := h.usecase.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrCannotDeleteEntity(models.Item{}.TableName(), err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
