package models

// Pagination metadata to include in API responses.
type Pagination struct {
	CurrentCursor *uint `json:"current_cursor" form:"cursor"`
	NextCursor    *uint `json:"next_cursor"`
	Limit         int   `json:"limit" form:"limit"`
	TotalItems    uint  `json:"total_items"`
}
