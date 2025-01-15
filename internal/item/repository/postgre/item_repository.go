package postgre

import (
	"project/internal/item/entity"

	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(item *entity.ItemCreate) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) Update(item *entity.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Item{}, id).Error
}
