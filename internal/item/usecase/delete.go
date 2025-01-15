package usecase

import (
	"time"
)

// DeleteItem marks an item as deleted (soft delete).
func (u *itemUsecase) DeleteItem(id uint) error {
	// Get existing item
	item, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Set the item as deleted (soft delete) if not marked as deleted
	if item.DeletedAt != nil {
		now := time.Now()
		item.DeletedAt = &now
		return nil
	}

	// Update the item in the repository (instead of hard deleting)
	return u.repo.Update(item)
}
