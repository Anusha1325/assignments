package datatype

import "fmt"

func Booldatatype() {

	m := true
	n := false
	fmt.Println("m: ", m, "n; ", n)
	x := m || n
	y := m && n
	fmt.Println("x: ", x, "y; ", y)
}