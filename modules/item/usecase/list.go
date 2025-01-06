package usecase

import (
	"project/common"
	"project/modules/item/entity"

	"gorm.io/gorm"
)

// ListItems lists items with pagination, filtering by type, and sorting options.
func (u *itemUsecase) ListItems(pagination *common.Pagination, itemType *entity.ItemType, sortBy string) ([]entity.Item, error) {
	// Apply filters
	var filters []func(*gorm.DB) *gorm.DB

	// Exclude soft-deleted items
	filters = append(filters, func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at IS NULL")
	})

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
		// return db.Order(sortBy)
		return db
	})
	if err != nil {
		return nil, err
	}

	return items, nil
}
