package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x 
	return 
}

func main() {
	fmt.Println("1. My favorite number is", rand.Intn(10))

	
	//Printf stands for print formatted. %g is a format specifier used for floating point numbers
	fmt.Printf("2. Now you have %g problems. \n", math.Sqrt(7))

	//In Go, a name is exported if it begins with a capital letter. 
	//Pi must be written with a capital p for the program to function and for Pi to exported from the math package.
	fmt.Println("3. Pi can be written as", math.Pi)

	//A function can take zero or more arguments.
	//Notice that the type comes after the variable name.
	fmt.Println("4. Two plus four is equal to", add(2,4))

	//A function can return any number of results. The swap function returns two strings.
	a, b := swap("hello", "world")
	fmt.Println("5. If I swap the words 'hello' and 'world' I get the following:",a, b)

	//Step 6 was a matter of reformatting the add function
	fmt.Println("6. Step 6 was a matter of reformatting previous work")

	//Go's return values may be named. If so, they are treated as variables defined at the top of the function
	//A return statement without arguments returns the named return values. This is known as a "naked" return.
	//Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
	fmt.Print("7. If I wanted to split a number in two such as the number 17, it would look like:")
	fmt.Print(split(17))

}