package main

import "fmt"

// 13/26 NIL INTERFACE VALUES
/*
	A nil interface value holds neither value nor concrete type

	Calling a method on a nil interface is a run-time error
	This is bc there is no type inside the interface tuple to indicate which concrete method to call
*/

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
