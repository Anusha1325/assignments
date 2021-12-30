package variables

import "fmt"

// Single variable declaration with an initial value
func Variables1() { 
	var name string = "Anusha"
	var place string = "Texas"
	var age int = 24
	var height int = 63
	fmt.Println("Hi! I am ", name, "and my age is", age, "\nI'm about", height, "inches and I live in", place)
}

/* Multiple variable declaration by defining same variables 
i.e., name, age, height with the type inference*/
func Variables2() {
	var ( 
		name, age, height, place = "Anusha", 24, 63, "Texas"
	)
	fmt.Println("Name:", name, "\nAge:", age, "\nHeight:", height, "\nPlace", place)
}

// Short-handed multiple variable decalration with new Variables
func Variables3() { 
	Name, Age, Height, Place := "Anusha", 24, 63, "Texas" 
	fmt.Println("Hi! I am", Name, "My age is", Age, "I'm about", Height, "inches and I live in", Place)
}