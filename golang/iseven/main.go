package main

import "fmt"

/*
Problem:
Write a function in Go
called IsEven that takes
in a list of integers and returns
the all even numbers in the list.
*/

func IsEven(input []int) []int{
	output:= []int{}

	for x := 0; x < len(input); x++ { // ++ is a statement.
        fmt.Println(input[x])
		if input[x] % 2 == 0 {
			output[x] = input[x]
		}
    }

	return output
}

func main() {
	testnums := []int{1,2,3,4,5,6,7,8,9,10}

	for x := 0; x < len(testnums); x++ { // ++ is a statement.
        if testnums[x] % 2 == 0 {
			fmt.Println(testnums[x])
		}
		
    }

	IsEven(testnums)
}