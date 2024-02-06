package main

import (
	"fmt"
	"math"
)

// 11/26 INTERFACE VALUES
/*
	Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
		(value, type)

	An interface value holds a specific underlying concrete type
	Calling a method on an interface value executes the method on the same name of its underlying type
*/

// I is an interface that specifies a single method method M()
// All implementing types should provide their own implementation of M(). This is not optional.
type I interface {
	M()
}

// T is a simple struct with a string "field" S
type T struct {
	S string
}

// F is a simple type alias for the float64 type.
// It allows you to create a new type name for float64 for improved code readability and semantic meaning.
type F float64

// M is a method that is associated with a type T
// It prints the value of the "S" field of the T instance
// This methods operates on a pointer to T, allowing it to modify the underlying T instance
func(t *T) M() {
	fmt.Println(t.S)
}

// M is a method that is associated with type F
// It prints the value of F which is of the type foat64
func(f F) M(){
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}