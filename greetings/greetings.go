package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Returns a greeting for the named person
func Hello(name string, gtype string) (string, error) {
	//Check if name is empty and return error
	if name == "" {
		return "", errors.New("empty name entry")
	}

	var message string
	//If standard greeting is requested
	if gtype == "std" {
		//If name has a value return a greeting message
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	}

	//If random greeting is requested
	if gtype == "random" {
		message = fmt.Sprintf(randomFormat(), name)
	}

	return message, nil

}

//init sets intial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of the set greeting messages. The returned
//message is selected at random
func randomFormat() string {
	//A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Nice to see you, %v!",
		"Hail, %v! Well met!",
		"Hiya, How are you %v?",
	}

	//Return a random selected message format by specifying
	//a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
