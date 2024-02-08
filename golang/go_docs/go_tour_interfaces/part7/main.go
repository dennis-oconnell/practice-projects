package main

import "fmt"

// 15/26 TYPE ASSERTIONS
/*
	A type assertion provides access to an interface value's underlying concrete value

		i := i.(T)

	This statement asserts that the interface value 'i' holds the concrete type T
	and assigns the underlying T value to the variable 't'

	If i does not hold a T, this statement will trigger a panic

	To test wether an interface value holds a specific type, a type assertion can return two values:
	the underlying value and a boolean value that reports wether the assertion succeeded

	t, ok := i.(T)

	If i holds a T, then t will be the underlying value and ok will be true

	If not, ok will be false, and t will be the zero value of the type T, and no panic occurs

	Note the similarity of this syntax and that of reading from a map
*/

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s,ok)

	f, ok := i.(float64)
	fmt.Println(f,ok)

	f = i.(float64) // panic
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
