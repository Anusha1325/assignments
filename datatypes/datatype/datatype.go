package datatype

import (
	"fmt"
	"unsafe"
)

func Booldatatype() {
	y := true
	x := false
	fmt.Println("y: , x:", y, x)
	a := y || x
	b := y && x
	fmt.Println("a: , b: ", a, b)
}

func Stringdatatype() {
	var name, month string = "Anusha", "August"
	fmt.Println("\nname: , month:", name, month)
}

func Signedintdatatype() {
	var num int = 15
	var num1 int8 = 120
	// var nnum1 int8 = 2500 --> can't be assigned
	var num2 int16 = -30000
	// var nnum2 int16 = -2500738477 --> can't be assigned
	var num3 int32 = 2000120000
	// var nnum3 int32 = 73874827385726374 --> can't be assigned
	var num4 int64 = 8124758732097929473
	// var nnum4 int64 = -93747367428637573845 --> can't be assigned

	/* Printing the value of the integers along with type and size (returned in bytes) */
	fmt.Printf("\nvalue of num: %v, type of num: %T, size of num: %d", num, num, unsafe.Sizeof(num))
	fmt.Printf("\nvalue of num1: %v, type of num1: %T, size of num1: %d", num1, num1, unsafe.Sizeof(num1))
	fmt.Printf("\nvalue of num2: %v, type of num2: %T, size of num2: %d", num2, num2, unsafe.Sizeof(num2))
	fmt.Printf("\nvalue of num3: %v, type of num3: %T, size of num3: %d", num3, num3, unsafe.Sizeof(num3))
	fmt.Printf("\nvalue of num4: %v, type of num4: %T, size of num4: %d", num4, num4, unsafe.Sizeof(num4))
}

func Unsignedintdatatype() {
	var n uint = 15
	var n1 uint8 = 250
	// var nn1 uint8 = 2500 --> can't be assigned
	var n2 uint16 = 62345
	// var nn2 uint16 = -6583763 --> can't be assigned
	var n3 uint32 = 3343126347
	// var nn3 uint32 = 33431263476354 --> can't be assigned
	var n4 uint64 = 18124758732097929473
	// var nn4 uint64 = -181247587320979294737873 --> can't be assigned

	/* Printing the value of the integers along with type and size (returned in bytes) */
	fmt.Printf("\n\nvalue of n: %v, type of n: %T, size of n: %d", n, n, unsafe.Sizeof(n))
	fmt.Printf("\nvalue of n1: %v, type of n1: %T, size of n1: %d", n1, n1, unsafe.Sizeof(n1))
	fmt.Printf("\nvalue of n2: %v, type of n2: %T, size of n2: %d", n2, n2, unsafe.Sizeof(n2))
	fmt.Printf("\nvalue of n3: %v, type of n3: %T, size of n3: %d", n3, n3, unsafe.Sizeof(n3))
	fmt.Printf("\nvalue of n4: %v, type of n4: %T, size of n4: %d", n4, n4, unsafe.Sizeof(n4))
}

func Signedtyperange() {

	// int8 --> range: -128 to 127
	var sn1 int8 = 127 // initializing with the maximum allowed positive number
	fmt.Printf("\n\nmax value of int8: %d and type is: %T\n", sn1, sn1); // printing the value and data type
	sn1 = sn1+1 // making the value out of range by incrementing by 1
	fmt.Printf("min value of int8: %d and type is: %T\n", sn1, sn1); // printing out new value and type

	//int16 --> range: -32768 to 32767
	var sn2 int16 = 32767 
	fmt.Printf("\nmax value of int16: %d and type is: %T\n", sn2, sn2); 
	sn2 = sn2+1 
	fmt.Printf("min value of int16: %d and type is: %T\n", sn2, sn2); 

	//int32 --> range: -2147483648 to 2147483647
	var sn3 int32 = 2147483647 
	fmt.Printf("\nmax value of int32: %d and type is: %T\n", sn3, sn3); 
	sn3 = sn3+1 
	fmt.Printf("min value of int32: %d and type is: %T\n", sn3, sn3);

	//int32 --> range: -9223372036854775808 to 9223372036854775807
	var sn4 int64 = 9223372036854775807 
	fmt.Printf("\nmax value of int64: %d and type is: %T\n", sn4, sn4); 
	sn4 = sn4+1 
	fmt.Printf("min value of int64: %d and type is: %T\n", sn4, sn4);

}
