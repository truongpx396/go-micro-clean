package http

import (
	"net/http"
	"project/modules/item/domain/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateItem godoc
// @Summary      Update an item by ID
// @Description  Replace the entire item resource with new data
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "Item ID"
// @Param        item  body      models.Item  true  "Updated item data"
// @Success      200   {object}  models.APIResponse
// @Failure      400   {object}  models.APIResponse
// @Failure      404   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /items/{id} [put]
func (h *itemHandler) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid item ID",
		})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid request payload",
		})
		return
	}

	item.ID = uint(id)
	if err := h.usecase.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: "Failed to update item",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data:    item,
		Message: "Item updated successfully",
	})
}

// PatchItem godoc
// @Summary      Partially update an item
// @Description  Update specific fields of an item resource
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Item ID"
// @Param        item  body      models.ItemUpdate  true  "Fields to update"
// @Success      200   {object}  models.APIResponse
// @Failure      400   {object}  models.APIResponse
// @Failure      404   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /items/{id} [patch]
func (h *itemHandler) PatchItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid item ID",
		})
		return
	}

	var updates models.ItemUpdate
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid request payload",
		})
		return
	}

	item, err := h.usecase.PartiallyUpdateItem(uint(id), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: "Failed to update item",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data:    item,
		Message: "Item updated successfully",
	})
}
