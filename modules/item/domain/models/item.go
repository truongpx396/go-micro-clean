package models

import (
	"errors"
	"project/modules/item/domain/enums"
	"strings"
	"time"
)

// Item represents the Item model.
type Item struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Type      enums.ItemType `json:"type"`
	Image     Image          `json:"image" gorm:"type:jsonb"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty" gorm:"index"`
}

// Validate validates the Item fields.
func (item *Item) Validate() error {
	// Check for empty name
	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	// Ensure item type is valid (only Physical or Digital)
	if item.Type != enums.Physical && item.Type != enums.Digital {
		return errors.New("invalid item type")
	}

	return nil
}

// ItemUpdate represents the fields that can be updated for an item.
type ItemUpdate struct {
	Name  *string         `json:"name,omitempty"`
	Type  *enums.ItemType `json:"type,omitempty"`
	Image *Image          `json:"image,omitempty"`
}
