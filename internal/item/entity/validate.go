package entity

import (
	"errors"
	"strings"
)

// validateItemFields is a shared function to validate item fields.
func validateItemFields(name string, itemType ItemType) error {
	// Check for empty name
	if strings.TrimSpace(name) == "" {
		return errors.New("item name cannot be empty")
	}

	// Ensure item type is valid (only Physical or Digital)
	if itemType != Physical && itemType != Digital {
		return ErrInvalidItemType
	}

	return nil
}
