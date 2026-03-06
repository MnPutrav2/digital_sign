package errs

import (
	"errors"
)

var (
	IsErrorDuplicate = errors.New("duplicate")
)
