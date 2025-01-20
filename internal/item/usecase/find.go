package usecase

import (
	"errors"
	"go-micro-clean/internal/item/entity"
)

// GetItemByID retrieves an item by its ID with additional error handling.
func (u *itemUsecase) GetItemByID(id uint) (*entity.Item, error) {
	item, err := u.repo.GetByID(id)
	if err != nil || item.DeletedAt != nil {
		return nil, errors.New("item not found")
	}
	return item, nil
}
