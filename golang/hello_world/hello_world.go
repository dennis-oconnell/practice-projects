// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"fmt"

	"rsc.io/quote"
)

//Implement a main function to print a message to the console. A main function executes by default when you run the main package.g
func main() {
	fmt.Println("This is my Go Hello World Program; it's so Hello Worldy! On the next line, look for a pithy quote!")

	fmt.Println(quote.Go())
	fmt.Println("And on the next few lines, we'll call some other functions from the quote package for more interesting and useful quotes!")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}