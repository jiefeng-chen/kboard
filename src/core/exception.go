package core

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func NewError(format string, args ...string) error {
	return errors.New(fmt.Sprintf(format, args))
}

func CheckError(err error, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		log.Fatal(err.Error())
		if code == 0 {
			os.Exit(code)
		}
	}
}
