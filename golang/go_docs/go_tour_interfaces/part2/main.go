package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This means type T implements the interface I
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// 10/26 INTERFACES ARE IMPLEMENTED IMPLICITLY
/*
	A type implements an interface by implementing its methods
		A type implements an interface by implementing its methods
			A type implements an interface by implementing its methods

	There is no explicit declaration of intent, no "implements" keyword
	
	Implicit interfaces "decouple" the definition of an interface from its implementation,
	Which could then appear in any package without prearrangement
*/