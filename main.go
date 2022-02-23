package main

import (
	"fmt"
	"log"

	greetings "github.com/akshayhiremath/golang/greetings"
	stringutils "github.com/alessiosavi/GoGPUtils/string"
	"rsc.io/quote"
)

func main() {

	//Add logger with initial config to set prefix and disable timestamp, source file and line number
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	var message string
	var err error
	//Request a standard greeting
	message, err = greetings.Hello("John", "std")

	//Log error and exit
	if err != nil {
		log.Fatal(err)
	}

	//Print a greeting message with the name to console
	fmt.Println(message)

	//Request a random greeting
	message, err = greetings.Hello("John", "random")

	//Log error and exit
	if err != nil {
		log.Fatal(err)
	}

	//Print a greeting mesage with the name to console
	fmt.Println(message)

	//A slice of names
	names := []string{"George", "Ricky", "Tony", "Sandy"}

	messages, err := greetings.Hellos(names)

	//Error in name mapping
	if err != nil {
		log.Fatal(err)
	}
	//If no error was returned, print the returned map of
	//messages to the console
	fmt.Println(messages)

}

func stringUtilUse() {
	//Use function from stringutils module
	result := stringutils.CheckPresence([]string{"India", "Canada", "USA"}, "Japan")
	fmt.Println("Result ", result)
}

func getAndPrintQuote() {
	//TODO use the sampler and get greeting in Marathi
	// Print quote using quote module
	fmt.Println(quote.Go())
}
