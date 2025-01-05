package http

import (
	"net/http"
	"project/common"
	"project/modules/item/domain/enums"
	"project/modules/item/domain/models"

	"github.com/gin-gonic/gin"
)

// ListItems godoc
// @Summary      List items with cursor-based pagination
// @Description  Retrieve a paginated list of items
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        cursor  query     int  false  "Cursor for pagination"
// @Param        limit   query     int  false  "Number of items to retrieve"
// @Param        type    query     string  false  "Item type"
// @Param        sortBy  query     string  false  "Sort by field"
// @Success      200     {object}  models.ItemListResponse
// @Failure      400     {object}  common.AppError
// @Failure      500     {object}  common.AppError
// @Router       /items [get]
func (h *itemHandler) ListItems(c *gin.Context) {
	var pagination models.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequestWithMsg(err, "Invalid pagination parameters"))
		return
	}

	itemType := c.Query("type")
	sortBy := c.DefaultQuery("sortBy", "name")

	var itemTypeEnum *enums.ItemType
	if itemType != "" {
		enumValue, err := enums.ParseItemType(itemType)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequestWithMsg(err, "Invalid item type"))
			return
		}
		itemTypeEnum = &enumValue
	}

	items, err := h.usecase.ListItems(&pagination, itemTypeEnum, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrCannotListEntity(models.Item{}.TableName(), err))
		return
	}

	response := models.ItemListResponse{
		Data: items,
		Paging: models.Pagination{
			CurrentCursor: pagination.CurrentCursor,
			NextCursor:    pagination.NextCursor,
			Limit:         pagination.Limit,
			TotalItems:    pagination.TotalItems,
		},
	}

	c.JSON(http.StatusOK, response)
}
