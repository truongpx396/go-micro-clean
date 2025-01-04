package postgre

import (
	"project/modules/item/domain/models"

	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(item *models.ItemCreate) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) Update(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&models.Item{}, id).Error
}
