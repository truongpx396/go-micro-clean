package postgre

import (
	"project/modules/item/domain/models"

	"gorm.io/gorm"
)

func (r *itemRepository) List(pagination *models.Pagination, filters ...func(*gorm.DB) *gorm.DB) ([]models.Item, error) {
	var items []models.Item
	query := r.db

	// Apply filters
	for _, filter := range filters {
		query = filter(query)
	}

	// Count total items
	var totalItems int64
	if err := query.Model(&models.Item{}).Count(&totalItems).Error; err != nil {
		return nil, err
	}
	pagination.TotalItems = uint(totalItems)

	// Apply cursor and limit
	query = query.Order("id").Limit(pagination.Limit)
	if pagination.CurrentCursor != nil && *pagination.CurrentCursor > 0 {
		query = query.Where("id > ?", *pagination.CurrentCursor)
	}

	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}

	if len(items) > 0 {
		nextCursor := items[len(items)-1].Id
		pagination.NextCursor = &nextCursor
	}

	return items, nil
}
