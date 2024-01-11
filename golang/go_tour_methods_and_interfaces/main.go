package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//1. Methods
//This is a method defined on a type vertex. A method has a special receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

//2. Methods are just funcs
//This is a regular function that does the same thing as Abs()
//There is no change in functionality, only a difference in construction
func Hypotenuse(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//3. Methods Cont.
//Here is a numeric type 'MyFloat' with an accompanying method called 'AlwaysBePos'
type MyFloat float64

func (f MyFloat) AlwaysBePos() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	//1. Methods
	//Go has no classes! Java programmers be warned!
	//However, you can define methods on types.
	//A method is a function with a special receiver argument.
	//The receiver appears in its own argument list between the func keyword and the method name
	//Below, the 'Abs' method has a receiver type 'Vertex' named 'v'

	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//2. Methods are functions
	//A method is just a function with a receiver argument
	//Hypotenuse is the same as Abs, just written as a regular function
	fmt.Println(Hypotenuse(v))

	//3. Methods continued
	//You can declare a method on non struct types too
	//You can only declare a method with a receiver whose type is defined in the same package as the mehtod.
	//You cannot declare a method with a receiver whose type is defined in another package (which includes the built in types such as int)
	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.AlwaysBePos())
}