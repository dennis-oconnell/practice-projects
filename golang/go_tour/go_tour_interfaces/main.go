package main

import (
	"fmt"
	"math"
)

//9. Interfaces
type Abser interface {
	Abs() float64
}

//10. Interfaces are implemented implicitly
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	//9. Interfaces
	//An Interface type is defined as a set of method signatures
	//A value interface type can hold any value that implements those methods
	var abe Abser
	h := MyFloat(-math.Sqrt2)
	g := Vertex{3,4}

	abe = h		//a MyFloat implements Abser
	abe = &g		//a *Vertex implements Abser

	fmt.Println(abe.Abs())

	//10. Interfaces are implemented implicitly
	//A type implements an interface by implementing its methods
	//There is no explicit declaration of intent (no 'implements' keyword)
	//Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement
	var i I = T{"hello"}
	i.M()

	//11. Interface Values
	//Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
	//An interface value holds a value of a specific underlying concrete type
	//Calling a method on an interface value executes the mehtod on the same name on its underlying type
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}