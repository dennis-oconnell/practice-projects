package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("1. My favorite number is", rand.Intn(10))

	//Printf is print formatted
	//%g is a format specifier used for floating point numbers
	fmt.Printf("2. Now you have %g problems. \n", math.Sqrt(7))

	//In Go, a name is exported if it begins with a capital letter. 
	//Pi must be written with a capital p for the program to function and for Pi to exported from the math package.
	fmt.Println("3. Pi can be written as", math.Pi)

	//A function can take zero or more arguments.
	//Notice that the type comes after the variable name.
	fmt.Println("4. Two plus four is equal to", add(2,4))
}