package usecase

import (
	"errors"
	"go-micro-clean/internal/item/entity"
	"time"
)

// UpdateItem updates an existing item with validation and consistency checks.
func (u *itemUsecase) UpdateItem(item *entity.Item) error {
	// Get the existing item
	existingItem, err := u.repo.GetByID(item.ID)
	if err != nil {
		return err
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
	item.BeforeUpdate()

	// Fully replace the item in the database with the item being passed through API
	return u.repo.Update(item)
}

// PartiallyUpdateItem updates specific fields of an existing item.
func (u *itemUsecase) PartiallyUpdateItem(id uint, updates entity.ItemUpdate) (*entity.Item, error) {
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
		if *updates.Type != entity.Physical && *updates.Type != entity.Digital {
			return nil, errors.New("invalid item type")
		}
		existingItem.Type = *updates.Type
	}

	if updates.Image != nil {
		existingItem.Image = *updates.Image
	}

	// Set updated timestamp
	now := time.Now()
	existingItem.UpdatedAt = &now

	// Update the existing item in the repository
	if err := u.repo.Update(existingItem); err != nil {
		return nil, err
	}

	return existingItem, nil
}
