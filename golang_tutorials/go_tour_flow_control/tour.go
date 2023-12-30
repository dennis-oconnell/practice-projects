package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {

	//Go has only one looping construct, the for loop
	//the braces { } are always required.
	//the init statement: executed before the first iteration --> 'i:=0'
	//the condition expression: evaluated before every iteration --> 'i<10'
	//the post statement: executed at the end of every iteration --> 'i++'
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println("1. This sum is equal to ", sum)

	//the init and post statements are optional in Go! What?!
	sum = 1
	for ; sum <1000; {
		//fmt.Println(sum)
		sum += sum

	}
	fmt.Println("2. This sum is equal to ", sum)

	//For is Go's while loop
	//and we can drop the semicolons C's while is spelled for in Go.
	sum = 1
	for sum < 1000 {
		sum+= sum
	}
	fmt.Println("3. This sum is equal to ",sum)

	//If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	/*
		for {
			fmt.Println("4. This is a forever loop! Oh no!")
		}
	*/

	//Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	fmt.Println("5. See the sqrt function for a nice use of the Golang 'if' statement:", sqrt(2), " ", sqrt(-4))

	//If with a short statement
	//Like for, the if statement can start with a short statement to execute before the condition.
	//Variables declared by the statement are only in scope until the end of the if.
	//(If we try using v in the last return statement of the pow function above, the program does not work since v is outside the scope there.)
	fmt.Println(
		"6. If statements can start with a short statement to execute before the condition. This can be a good place for variable declarations",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//If and Else
	//Variables declared inside an if short statement are also available inside any of the else blocks.
	fmt.Println(
		"7. Here, both calls to pow2 return their results before the call to fmt.Println in main begins:",
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)


}
