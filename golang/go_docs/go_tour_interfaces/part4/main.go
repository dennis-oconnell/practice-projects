package main

import "fmt"

//	12/26 INTERFACE VALUES WITH NIL UNDERLYING VALUES
/*
	If the concrete value inside the interface itself is nil,
	the method will be called with a nil receiver.

	In some other programming languages, this would trigger a null pointer exception.
	But, in Go, it is common to write methods that gracefully handle being called with a nil receiver.

	Note: An interface value that holds a nil concrete value is itself non-nil.
*/

//	I is an interface with a method M.
// Reminder: An interface type is defined as a set of method signatures.
type I interface {
	M()
}

// T is a struct with a string value called S.
type T struct {
	S string
}

// T implements M
func (t *T) M(){
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	// 'i' is an instance of the interface I
	var i I

	// 't' is a pointer to a simple struct with a string value
	var t *T

	// A value of interface type can hold any value that implements those methods.
	i = t
	describe(i)

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
