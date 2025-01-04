package enums

import "errors"

// ItemType represents a custom enum using iota.
type ItemType int

const (
	Physical ItemType = iota
	Digital
)

func (it ItemType) String() string {
	return [...]string{"Physical", "Digital"}[it]
}

// ParseItemType converts a string to the corresponding ItemType.
func ParseItemType(s string) (ItemType, error) {
	switch s {
	case "Physical":
		return Physical, nil
	case "Digital":
		return Digital, nil
	default:
		return -1, errors.New("invalid item type")
	}
}
