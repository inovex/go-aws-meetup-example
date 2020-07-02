package models

import (
	"errors"
	"fmt"
)

var ErrItemAlreadyExists = errors.New("item already exists")

var ErrWriteFailed = errors.New("write failed")

func ErrWriteFailedWithCause(cause error) error {
	if cause != nil {
		return fmt.Errorf("%w: %s", ErrWriteFailed, cause)
	}
	return ErrWriteFailed
}
