package core

import (
	"errors"
	"fmt"
)

func NewError(format string, args ...string) error {
	return errors.New(fmt.Sprintf(format, args))
}