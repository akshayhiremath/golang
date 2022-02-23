package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Returns a greeting for the named person
func Hello(name string, greetingtype string) (string, error) {
	//Check if name is empty and return error
	if name == "" {
		return "", errors.New("empty name entry")
	}

	var message string
	//If standard greeting is requested
	if greetingtype == "std" {
		//If name has a value return a greeting message
		// message = fmt.Sprintf("Hi, %v. Welcome!", name)
		message = fmt.Sprintf("Hi, %v. Welcome!", "")
	}

	//If random greeting is requested
	if greetingtype == "random" {
		message = fmt.Sprintf(randomFormat(), name)
	}

	return message, nil

}

//Returns a map that associates each of the named people
//with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages
	messages := make(map[string]string) //make map initialization

	// Loop through the received slice of names, calling the Hello
	//function to get a message for each name.
	for _, name := range names {

		message, err := Hello(name, "random")
		if err != nil {
			return nil, err
		}

		//In the map, associate the retrived name with the name
		messages[name] = message
	}

	return messages, nil
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
