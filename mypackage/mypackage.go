package mypackage

import (
	"errors"
	"fmt"
)

func PrintHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}

	message := fmt.Sprintf("Hello, %s! This is mypackage speaking!", name)

	return message, nil
}
