package http

import (
	"project/modules/item/domain/enums"
	"project/modules/item/domain/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemUsecase interface {
	CreateItem(item *models.Item) error
	GetItemByID(id uint) (*models.Item, error)
	UpdateItem(item *models.Item) error
	DeleteItem(id uint) error
	ListItems(pagination *models.Pagination, itemType *enums.ItemType, sortBy string) ([]models.Item, error)
	PartiallyUpdateItem(id uint, updates models.ItemUpdate) (*models.Item, error)
}

type itemHandler struct {
	usecase ItemUsecase
}

func NewItemHandler(usecase ItemUsecase) *itemHandler {
	return &itemHandler{usecase: usecase}
}

// CreateItem godoc
// @Summary      Create a new item
// @Description  Create a new item with name, type, and image data
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        item  body      models.Item  true  "Item data"
// @Success      201   {object}  models.APIResponse
// @Failure      400   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /items [post]
func (h *itemHandler) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid request payload",
		})
		return
	}

	if err := h.usecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Error: "Failed to create item",
		})
		return
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Data:    item,
		Message: "Item created successfully",
	})
}

// GetItemByID godoc
// @Summary      Retrieve an item by ID
// @Description  Get details of an item by its unique ID
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "Item ID"
// @Success      200   {object}  models.APIResponse
// @Failure      404   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /items/{id} [get]
func (h *itemHandler) GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Error: "Invalid item ID",
		})
		return
	}

	item, err := h.usecase.GetItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Error: "Item not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Data: item,
	})
}

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
