package main

import (
	"fmt"
	"github.com/Anusha1325/assignments/tree/main/scopes/scope" // importing the package scope
)

func main() {
	scope.Parameter() // Calling function Parameter()
	scope.Area() // Calling function Area()

    /* Calling global variable Length from package scope by importing it */
	fmt.Println("Length: ", scope.Length)
}