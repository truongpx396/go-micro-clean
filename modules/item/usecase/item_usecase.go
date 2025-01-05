package usecase

import (
	"project/modules/item/domain/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.ItemCreate) error
	GetByID(id uint) (*models.Item, error)
	GetByName(name string) (*models.Item, error)
	Update(item *models.Item) error
	Delete(id uint) error
	List(pagination *models.Pagination, filters ...func(*gorm.DB) *gorm.DB) ([]models.Item, error)
}

type itemUsecase struct {
	repo ItemRepository
}

func NewItemUsecase(repo ItemRepository) *itemUsecase {
	return &itemUsecase{repo: repo}
}
