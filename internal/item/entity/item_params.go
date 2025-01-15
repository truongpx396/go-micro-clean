package entity

import "project/common"

// ItemCreate represents the request payload for creating an item.
type ItemCreate struct {
	common.SQLModelCreate
	Name  string       `json:"name"`
	Type  ItemType     `json:"type"`
	Image common.Image `json:"image"`
}

func (ItemCreate) TableName() string {
	return Item{}.TableName()
}

// Validate validates the ItemCreate fields.
func (item *ItemCreate) Validate() error {
	return validateItemFields(item.Name, item.Type)
}

// ItemResponse represents the response created/updated item.
type ItemIdRead struct {
	ID uint `json:"id"`
}

// ItemUpdate represents the fields that can be updated for an item.
type ItemUpdate struct {
	Name  *string       `json:"name,omitempty"`
	Type  *ItemType     `json:"type,omitempty"`
	Image *common.Image `json:"image,omitempty"`
}

func (ItemUpdate) TableName() string {
	return Item{}.TableName()
}

type ItemListResponse struct {
	Data   []Item            `json:"data"`   // The main response data (can be list or single resource).
	Paging common.Pagination `json:"paging"` // Pagination details (optional, for paginated responses).
}
