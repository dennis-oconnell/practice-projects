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

	for x := 0; x < len(input); x++ {
		if input[x] % 2 == 0 {
			output = append(output, input[x])
		}
    }

	return output
}

func main() {
	testnums := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(IsEven(testnums))

	testmor := []int{1,6,23,2,1,64,-1,2,4,2}

	fmt.Println(IsEven(testmor))
}