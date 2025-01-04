package postgre

import "project/modules/item/domain/models"

func (r *itemRepository) GetByID(id uint) (*models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) GetByName(name string) (*models.Item, error) {
	var item models.Item
	if err := r.db.Where("name = ?", name).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
