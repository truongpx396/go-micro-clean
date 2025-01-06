package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Image represents a custom image structure.
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Scan implements the sql.Scanner interface for Image.
func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan Image")
	}
	return json.Unmarshal(bytes, i)
}

// Value implements the driver.Valuer interface for Image.
func (i Image) Value() (driver.Value, error) {
	return json.Marshal(i)
}
