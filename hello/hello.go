package main

import (
	"fmt"

	greetings "github.com/akshayhiremath/golang/greetings"
)

func main() {

	//fmt.Println("Hello, World!")
	message := greetings.Hello("Akshay")
	fmt.Println(message)
}
