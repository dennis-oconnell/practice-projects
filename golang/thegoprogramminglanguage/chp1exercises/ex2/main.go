// TITLE: GOPL Chapter 1 Exercise 2
// DATE: 03/04/2024
// START:8:56
// END:9:01
// MINUTES:5
// GRADE: Pass

// PROBLEM: Exercise 1.2: Modif y the echo prog ram to print the index and value of each of its arguments, one per line.
// SOLUTION: Used formatted printing and for loop

package main

import (
	"fmt"
	"os"
)

func main(){
	for i, arg := range os.Args[1:] {
		fmt.Printf("Index: %d Value:%s\n", i, arg)
	}
}