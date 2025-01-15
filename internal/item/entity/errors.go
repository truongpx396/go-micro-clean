package entity

type customError struct {
	ErrMsg string
}

func (e *customError) Error() string {
	return e.ErrMsg
}

func (e *customError) String() string {
	return e.ErrMsg
}

func NewCustomError(msg string) error {
	return &customError{ErrMsg: msg}
}

var ErrInvalidItemType = NewCustomError("invalid item type")
