package http

import (
	"net/http"
	"project/common"
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
// @Success      200   {object}  models.Item
// @Failure      400   {object}  common.AppError
// @Failure      404   {object}  common.AppError
// @Failure      500   {object}  common.AppError
// @Router       /items/{id} [put]
func (h *itemHandler) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	if _, err := h.usecase.GetItemByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, common.ErrEntityNotFound(models.Item{}.TableName(), err))
		return
	}

	item.ID = uint(id)
	if err := h.usecase.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrCannotUpdateEntity(item.TableName(), err))
		return
	}

	c.JSON(http.StatusOK, item)
}

// PatchItem godoc
// @Summary      Partially update an item
// @Description  Update specific fields of an item resource
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Item ID"
// @Param        item  body      models.ItemUpdate  true  "Fields to update"
// @Success      200   {object}  models.Item
// @Failure      400   {object}  common.AppError
// @Failure      404   {object}  common.AppError
// @Failure      500   {object}  common.AppError
// @Router       /items/{id} [patch]
func (h *itemHandler) PatchItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequestWithMsg(err, "Invalid item ID"))
		return
	}

	var updates models.ItemUpdate
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequestWithMsg(err, "Invalid request payload"))
		return
	}

	if _, err := h.usecase.GetItemByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, common.ErrEntityNotFound(models.Item{}.TableName(), err))
		return
	}

	item, err := h.usecase.PartiallyUpdateItem(uint(id), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrCannotUpdateEntity(models.Item{}.TableName(), err))
		return
	}

	c.JSON(http.StatusOK, item)
}
