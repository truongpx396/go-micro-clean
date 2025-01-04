package usecase

import (
	"errors"
	"project/modules/item/domain/enums"
	"project/modules/item/domain/models"
	"time"
)

// UpdateItem updates an existing item with validation and consistency checks.
func (u *itemUsecase) UpdateItem(item *models.Item) error {
	// Check if item exists
	existingItem, err := u.repo.GetByID(item.ID)
	if err != nil {
		return errors.New("item not found")
	}

	// Validate item fields
	if err := item.Validate(); err != nil {
		return err
	}

	// Check if the name is being updated, and ensure uniqueness
	if existingItem.Name != item.Name {
		existingItemWithSameName, _ := u.repo.GetByName(item.Name)
		if existingItemWithSameName != nil {
			return errors.New("item with the same name already exists")
		}
	}

	// Set updated timestamp
	item.UpdatedAt = time.Now()

	// Fully replace the item in the database with the item being passed through API
	return u.repo.Update(item)
}

// PartiallyUpdateItem updates specific fields of an existing item.
func (u *itemUsecase) PartiallyUpdateItem(id uint, updates models.ItemUpdate) (*models.Item, error) {
	// Check if the existingItem exists
	existingItem, err := u.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("item not found")
	}

	// Validate and apply updates
	if updates.Name != nil {
		// Check if the name is being updated, and ensure uniqueness
		existingItemWithSameName, _ := u.repo.GetByName(*updates.Name)
		if existingItemWithSameName != nil && existingItemWithSameName.ID != id {
			return nil, errors.New("item with the same name already exists")
		}
		existingItem.Name = *updates.Name
	}

	if updates.Type != nil {
		if *updates.Type != enums.Physical && *updates.Type != enums.Digital {
			return nil, errors.New("invalid item type")
		}
		existingItem.Type = *updates.Type
	}

	if updates.Image != nil {
		existingItem.Image = *updates.Image
	}

	// Set updated timestamp
	existingItem.UpdatedAt = time.Now()

	// Update the existing item in the repository
	if err := u.repo.Update(existingItem); err != nil {
		return nil, err
	}

	return existingItem, nil
}
