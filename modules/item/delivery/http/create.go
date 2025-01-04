package http

import (
	"fmt"
	"net/http"
	"project/modules/item/domain/models"

	"github.com/gin-gonic/gin"
)

// CreateItem godoc
// @Summary      Create a new item
// @Description  Create a new item with name, type, and image data
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        item  body      models.ItemCreate  true  "Item data"
// @Success      201   {object}  models.APIResponse
// @Failure      400   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /items [post]
func (h *itemHandler) CreateItem(c *gin.Context) {
	var item models.ItemCreate
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: fmt.Sprintf("Invalid request payload: %v\n", err),
		})
		return
	}

	if err := h.usecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: fmt.Sprintf("Failed to create item %v\n", err),
		})
		return
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Data: models.ItemIdRead{
			ID: item.Id,
		},
		Message: "Item created successfully",
	})
}
