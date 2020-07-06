package models

import (
	"errors"
	"fmt"
)

var ErrItemNotFound = errors.New("item not found")

var ErrItemAlreadyExists = errors.New("item already exists")

var ErrWriteFailed = errors.New("write failed")

func ErrWriteFailedWithCause(cause error) error {
	if cause != nil {
		return fmt.Errorf("%w: %s", ErrWriteFailed, cause)
	}
	return ErrWriteFailed
}
