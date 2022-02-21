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
	message, err = greetings.Hello("Akshay", "std")

	//Log error and exit
	if err != nil {
		log.Fatal(err)
	}

	//Print a greeting message with the name to console
	fmt.Println(message)

	//Request a random greeting
	message, err = greetings.Hello("Akshay", "random")

	//Log error and exit
	if err != nil {
		log.Fatal(err)
	}

	//Print a greeting mesage with the name to console
	fmt.Println(message)

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
