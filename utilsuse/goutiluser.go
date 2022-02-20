package main

import (
	"fmt"

	stringutils "github.com/alessiosavi/GoGPUtils/string"
)

func main() {
	result := stringutils.CheckPresence([]string{"India", "Canada", "USA"}, "Japan")
	fmt.Println("Result ", result)
}
