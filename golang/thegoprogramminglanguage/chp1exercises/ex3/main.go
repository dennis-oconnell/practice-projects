// TITLE: GOPL Chapter 1 Exercise 3
// DATE:03/04/2024
// START:9:04
// END:
// MINUTES:
// GRADE:

// PROBLEM:Exercis e 1.3: Experiment to measure the dif ference in running time bet ween our potential ly
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates par t of the
// 	time package, and Sec tion 11.4 shows how to write benchmark tests for systematic performance
// 	evaluation.)
// SOLUTION:

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