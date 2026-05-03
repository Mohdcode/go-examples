package main

import (
	"fmt"
)

func variabl() {

	// Language basics: variables, constants, types, conversions
	var a int
	var d, e int = 4, 5
	a = 1
	var b = 2
	var k, j = 4, 5
	c, f, n, g := 2, 3, 4, 56
	n = a + b
	i := n

	fmt.Printf("printing the result of type var decalartion, %d,%d,%d,%d,%d,%d,%d,%d,%v\n", d, e, k, j, c, f, n, g, i)

	type m int
	var o m = 8

	const pil = "string"

	fmt.Println("starting off the new shitt", pil)
	fmt.Printf("type of variable is %T\n", o)

	q := int(o)
	fmt.Printf("type of var after type conversion is %T\n", q)
	var myInterface interface{} = "dada"

	value, ok := myInterface.(string) // type assertion to check the type of the interface as interface can keep any type or even functions
	if !ok {
		fmt.Printf("var is not a string %v", value)
	}
	fmt.Printf("var is  a string %v", value)

	//Basic Types: Includes bool, numeric types like int, float64, uint8 (alias byte), int32 (alias rune), and string.
	// Composite Types: Includes array, slice, map, struct, pointer, function, interface, and channel.

	//no type casting in golang natively
	
	
	

}
