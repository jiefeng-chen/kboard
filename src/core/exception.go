package core

import (
	"errors"
	"fmt"
	"os"
)

func NewError(format string, args ...string) error {
	return errors.New(fmt.Sprintf(format, args))
}

func CheckError(err error, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(code)
	}
}