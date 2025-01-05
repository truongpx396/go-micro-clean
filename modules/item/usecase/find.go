package usecase

import (
	"errors"
	"project/modules/item/domain/models"
)

// GetItemByID retrieves an item by its ID with additional error handling.
func (u *itemUsecase) GetItemByID(id uint) (*models.Item, error) {
	item, err := u.repo.GetByID(id)
	if err != nil || item.DeletedAt != nil {
		return nil, errors.New("item not found")
	}
	return item, nil
}
