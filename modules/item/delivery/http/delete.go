package http

import (
	"net/http"
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
// @Failure      400   {object}  models.APIResponse
// @Failure      404   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /items/{id} [delete]
func (h *itemHandler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid item ID",
		})
		return
	}

	if err := h.usecase.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: "Failed to delete item",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Message: "Item deleted successfully",
	})
}
