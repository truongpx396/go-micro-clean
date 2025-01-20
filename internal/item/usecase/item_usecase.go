package usecase

import (
	"go-micro-clean/common"
	"go-micro-clean/internal/item/entity"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *entity.ItemCreate) error
	GetByID(id uint) (*entity.Item, error)
	GetByName(name string) (*entity.Item, error)
	Update(item *entity.Item) error
	Delete(id uint) error
	List(pagination *common.Pagination, filters ...func(*gorm.DB) *gorm.DB) ([]entity.Item, error)
}

type itemUsecase struct {
	repo ItemRepository
}

func NewItemUsecase(repo ItemRepository) *itemUsecase {
	return &itemUsecase{repo: repo}
}
