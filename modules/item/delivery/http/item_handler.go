package http

import (
	"project/modules/item/domain/enums"
	"project/modules/item/domain/models"
)

type ItemUsecase interface {
	CreateItem(item *models.ItemCreate) error
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
