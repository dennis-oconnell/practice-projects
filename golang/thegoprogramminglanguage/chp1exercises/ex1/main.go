// TITLE: GOPL Chapter 1 Exercise 1
// DATE: 03/04/2024
// START: 8:49
// END: 8:51
// MINUTES: 2
// GRADE: A

// PROBLEM: Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
// SOLUTION: When os.Args[0] is printed the result is the name of the command that invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[0:], " "))	
}