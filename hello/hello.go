package main

import (
	"fmt"
	"log"

	greetings "github.com/akshayhiremath/golang/greetings"
)

func main() {

	//Add logger with initial config to set prefix and disable timestamp, source file and line number
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	var message string
	var err error
	//Request a standard greeting
	message, err = greetings.Hello("Akshay", "std")

	//Print a greeting message with the name to console
	fmt.Println(message)

	//Request a random greeting
	message, err = greetings.Hello("Akshay", "random")

	//Print a greeting mesage with the name to console
	fmt.Println(message)

	//Log error
	if err != nil {
		log.Fatal(err)
	}

}
