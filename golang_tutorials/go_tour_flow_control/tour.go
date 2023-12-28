package main

import "fmt"

func main() {

	//Go has only one looping construct, the for loop
	//the braces { } are always required.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i)
	}
	fmt.Println("sum is equal to ", sum)
}
