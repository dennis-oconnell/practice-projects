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

//4. Pointer Receivers
//The pointer receiver is necessary here to make permanent changes to the Vertex that we are trying to Scale. 
//If change the receiver to the value form 'v Vertex' instead of 'v *Vertex', Scale doesn't change when we apply Scale to it
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) AlwaysBePos() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//5. Pointers and Functions 
	//Here we've re-written 'Abs' and 'Scale' to be regular functions 'AbsRegular' and 'ScaleRegular'
	//This is to evaluate the behavior of these functions when we play with pointer values and references
	//When we do this we see how to behavior of a normal function and a method are different from one another

func AbsRegular(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleRegular(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f
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

	//4. Pointer Receivers
	//You can declare methods with pointer receivers
	//This means that the receiver type has the literal syntax '*T' for some type 'T' (Also, 'T' itself cannot be a pointer such as '*int')

	//Methods with pointer receivers can modify the value to which the receiver points 
	//Since methods often need to modify their receiver, pointer receivers are more common than value receivers
	v = Vertex{3,4}
	v.Scale(2)
	fmt.Println(v)
	fmt.Println(v.Abs())

	//5. Pointers and Functions 
	//Here we've re-written 'Abs' and 'Scale' to be regular functions 'AbsRegular' and 'ScaleRegular'
	//This is to evaluate the behavior of these functions when we play with pointer values and references
	//When we do this we see how to behavior of a normal function and a method are different from one another
	v = Vertex{9,40}
	ScaleRegular(&v, 10)
	fmt.Println(v)
	fmt.Println(AbsRegular(v))
	//here, we could remove the * operator in the ScaleRegular func declaration, but then we would also have to remove the & operator here in the function call for the program to compile.
	//When we do remove the & and * then the func will run but not make lasting changes to the thing which you are calling the function on
}