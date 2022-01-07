package scope

import "fmt"

// Global declaration of variable Length //
var Length = 10.8  

func Parameter() {
	// Functional declaration of variable width //
	var width = 5.6 

	// As Length is declared as a golbal varible it can be accessed in the function Parameter() //
	fmt.Println("length: ", Length) 

	// width can be accessed as it is declared in the same function //
	fmt.Println("width: ", width) 
}
