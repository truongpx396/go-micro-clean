package usecase

import (
	"errors"
	"project/modules/item/entity"
)

// CreateItem creates a new item with validation and uniqueness check.
func (u *itemUsecase) CreateItem(item *entity.ItemCreate) error {
	// Validate item fields
	if err := item.Validate(); err != nil {
		return err
	}

	// Check if item name already exists (if uniqueness is required)
	existingItem, _ := u.repo.GetByName(item.Name)
	if existingItem != nil {
		return errors.New("item with the same name already exists")
	}

	// Set created and updated timestamps
	item.PrepareForInsert()

	// Create the item
	return u.repo.Create(item)
}
