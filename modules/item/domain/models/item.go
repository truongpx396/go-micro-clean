package models

import (
	"errors"
	"project/common"
	"project/modules/item/domain/enums"
	"strings"
)

// Item represents the Item model.
type Item struct {
	common.SQLModel
	Name  string         `json:"name"`
	Type  enums.ItemType `json:"type"`
	Image Image          `json:"image" gorm:"type:jsonb"`
}

// ItemCreate represents the request payload for creating an item.
type ItemCreate struct {
	common.SQLModelCreate
	Name  string         `json:"name"`
	Type  enums.ItemType `json:"type"`
	Image Image          `json:"image"`
}

func (ItemCreate) TableName() string {
	return "items"
}

// Validate validates the Item fields.
func (item *Item) Validate() error {
	return validateItemFields(item.Name, item.Type)
}

// Validate validates the ItemCreate fields.
func (item *ItemCreate) Validate() error {
	return validateItemFields(item.Name, item.Type)
}

// validateItemFields is a shared function to validate item fields.
func validateItemFields(name string, itemType enums.ItemType) error {
	// Check for empty name
	if strings.TrimSpace(name) == "" {
		return errors.New("item name cannot be empty")
	}

	// Ensure item type is valid (only Physical or Digital)
	if itemType != enums.Physical && itemType != enums.Digital {
		return errors.New("invalid item type")
	}

	return nil
}

// ItemResponse represents the response created/updated item.
type ItemIdRead struct {
	ID uint `json:"id"`
}

// ItemUpdate represents the fields that can be updated for an item.
type ItemUpdate struct {
	Name  *string         `json:"name,omitempty"`
	Type  *enums.ItemType `json:"type,omitempty"`
	Image *Image          `json:"image,omitempty"`
}

type ItemDeleteResult struct {
	Success bool `json:"successful_deleted"`
}
