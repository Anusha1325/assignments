package scope

import "fmt"

func Area() {
	 
	/* width can't be accessed here as it is declared in another function even though under the same package scope */
	// var area = length * width

	w := 5.6
	var area = Length * w
	fmt.Println("area: ", area)
}