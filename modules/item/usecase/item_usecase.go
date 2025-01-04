package usecase

import (
	"errors"
	"project/modules/item/domain/models"
	"time"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.Item) error
	GetByID(id uint) (*models.Item, error)
	GetByName(name string) (*models.Item, error)
	Update(item *models.Item) error
	Delete(id uint) error
	List(pagination *models.Pagination, filters ...func(*gorm.DB) *gorm.DB) ([]models.Item, error)
}

type itemUsecase struct {
	repo ItemRepository
}

func NewItemUsecase(repo ItemRepository) *itemUsecase {
	return &itemUsecase{repo: repo}
}

// CreateItem creates a new item with validation and uniqueness check.
func (u *itemUsecase) CreateItem(item *models.Item) error {
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
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	// Create the item
	return u.repo.Create(item)
}

// GetItemByID retrieves an item by its ID with additional error handling.
func (u *itemUsecase) GetItemByID(id uint) (*models.Item, error) {
	item, err := u.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("item not found")
	}
	return item, nil
}

// DeleteItem marks an item as deleted (soft delete).
func (u *itemUsecase) DeleteItem(id uint) error {
	// Check if the item exists
	item, err := u.repo.GetByID(id)
	if err != nil {
		return errors.New("item not found")
	}

	// Set the item as deleted (soft delete)
	now := time.Now()
	item.DeletedAt = &now

	// Update the item in the repository (instead of hard deleting)
	return u.repo.Update(item)
}
