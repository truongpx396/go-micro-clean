package entity

import (
	"go-micro-clean/common"
)

// Item represents the Item model.
type Item struct {
	common.SQLModel
	Name  string       `json:"name"`
	Type  ItemType     `json:"type"`
	Image common.Image `json:"image" gorm:"type:jsonb"`
}

func (Item) TableName() string {
	return "items"
}

// Validate validates the Item fields.
func (item *Item) Validate() error {
	return validateItemFields(item.Name, item.Type)
}
