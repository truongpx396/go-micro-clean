package usecase

import (
	"project/modules/item/domain/enums"
	"project/modules/item/domain/models"

	"gorm.io/gorm"
)

// ListItems lists items with pagination, filtering by type, and sorting options.
func (u *itemUsecase) ListItems(pagination *models.Pagination, itemType *enums.ItemType, sortBy string) ([]models.Item, error) {
	// Apply filters
	var filters []func(*gorm.DB) *gorm.DB

	if itemType != nil {
		filters = append(filters, func(db *gorm.DB) *gorm.DB {
			return db.Where("type = ?", *itemType)
		})
	}

	// Apply sorting (default to sorting by name if no sort field provided)
	if sortBy == "" {
		sortBy = "name"
	}

	// Fetch items with pagination
	items, err := u.repo.List(pagination, func(db *gorm.DB) *gorm.DB {
		for _, filter := range filters {
			db = filter(db)
		}
		return db.Order(sortBy)
	})
	if err != nil {
		return nil, err
	}

	return items, nil
}
