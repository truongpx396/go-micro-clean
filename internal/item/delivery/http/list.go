package http

import (
	"net/http"
	"project/common"
	"project/internal/item/entity"

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
// @Success      200     {object}  entity.ItemListResponse
// @Failure      400     {object}  common.AppError
// @Failure      500     {object}  common.AppError
// @Router       /items [get]
func (h *itemHandler) ListItems(c *gin.Context) {
	var pagination common.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrInvalidRequestWithMsg(err, "Invalid pagination parameters"))
		return
	}

	itemType := c.Query("type")
	sortBy := c.DefaultQuery("sortBy", "name")

	var itemTypeEnum *entity.ItemType
	if itemType != "" {
		enumValue, err := entity.ParseItemType(itemType)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequestWithMsg(err, entity.ErrInvalidItemType.Error()))
			return
		}
		itemTypeEnum = &enumValue
	}

	items, err := h.usecase.ListItems(&pagination, itemTypeEnum, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrCannotListEntity(entity.Item{}.TableName(), err))
		return
	}

	response := entity.ItemListResponse{
		Data: items,
		Paging: common.Pagination{
			CurrentCursor: pagination.CurrentCursor,
			NextCursor:    pagination.NextCursor,
			Limit:         pagination.Limit,
			TotalItems:    pagination.TotalItems,
		},
	}

	c.JSON(http.StatusOK, response)
}
