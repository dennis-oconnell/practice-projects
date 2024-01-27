package main

import "fmt"

func fib(j int) {
	alpha := 0
	beta := 1

	fmt.Println(alpha)
	fmt.Println(beta)

	i := 0
	for i < j {
		fmt.Println(alpha + beta)
		temp := alpha
		alpha = beta
		beta = temp + beta
		i++
	}

}

func main() {
	fib(20)
}
