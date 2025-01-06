package postgre

import (
	"project/modules/item/entity"
)

func (r *itemRepository) GetByID(id uint) (*entity.Item, error) {
	var item entity.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) GetByName(name string) (*entity.Item, error) {
	var item entity.Item
	if err := r.db.Where("name = ?", name).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
