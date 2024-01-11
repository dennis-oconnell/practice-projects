package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func SqrtGuessMethod(x float64) float64 {
	e:= 0
	z:= 1.0
	for e < 10 {
		z -= (z*z - x) / (2*z)	
		e += 1
		fmt.Println(z)
	}
	return z
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

	//When we run the function SqrtGuessMethod we utilize a for loop as if it is a while loop
	//We also play around with some other conditionals
	//Again, the call to this function returns its results before the print call in the main method
	fmt.Println("8. Watch what happens when we run the function SqrtGuessMethod: ", SqrtGuessMethod(51))

	//Switch Statements
	//A switch statement is a shorter way of writing a sequence of if - else statements
	//It runs the first case whose value is equal to the condition expression
	//Go only runs the selected case, not all the cases that follow.
	//Go's switch cases need not be constants, and the values involved need not be integers.
	//When we run this code in the go tour online, it returns Linux, how interesting!
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	//Switch Evaluation Order
	//Switch cases evaluate from the top down, stopping when a case succeeds
	/*
		the code below does not call if i==0
		switch j {
		case 0:
		case f():
		}
	*/
	fmt.Println("10. When is Saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1: 
		fmt.Println("Tomorrow.")
	case today + 2: 
		fmt.Println("In just two days.")
	default:
		fmt.Println("Too far away...")
	}

	//11.Switch With No Condition
	//This can be a clean way to write long if-then-else chains
	fmt.Println("11. A switch with no condition is the same as 'switch true'")
	t:= time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!.")
	default:
		fmt.Println("Good Evening!")
	}

	//12. Defer
	//A defer statement defers the execution of a function until the surrounding function returns
	//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Println("wait")

	//13. Defer Stacking
	//Defer statements can be stacked, when the function returns, the defer statements are executed in a last in first out manner
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done!")
}
