package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	//Check if name is empty and return error
	if name == "" {
		return "", errors.New("empty name entry")
	}

	//If name has a value return a greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
