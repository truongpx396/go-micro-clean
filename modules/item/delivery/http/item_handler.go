package http

import (
	"project/common"
	"project/modules/item/entity"
)

type ItemUsecase interface {
	CreateItem(item *entity.ItemCreate) error
	GetItemByID(id uint) (*entity.Item, error)
	UpdateItem(item *entity.Item) error
	DeleteItem(id uint) error
	ListItems(pagination *common.Pagination, itemType *entity.ItemType, sortBy string) ([]entity.Item, error)
	PartiallyUpdateItem(id uint, updates entity.ItemUpdate) (*entity.Item, error)
}

type itemHandler struct {
	usecase ItemUsecase
}

func NewItemHandler(usecase ItemUsecase) *itemHandler {
	return &itemHandler{usecase: usecase}
}
