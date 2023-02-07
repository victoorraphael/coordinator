package errs

import (
	"errors"
	"fmt"
)

var (
	ErrFieldViolation = errors.New("field violation")
	ErrInternalError  = errors.New("internal error")
)

func WrapError(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}
