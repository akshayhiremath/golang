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

	//Send empty value
	message, err := greetings.Hello("")

	//Log error
	if err != nil {
		log.Fatal(err)
	}

	//If non empty name is sent to greetings provider. Print a greeting with the name to console
	fmt.Println(message)
}
