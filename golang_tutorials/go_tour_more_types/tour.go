package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	//1. The Wonderful World Of Pointers
	//Pointers hold the memory address of a value
	//The type '*T' is a pointer to a 'T' value. It's zero value is nil
	//The & operator generates a pointer to its operand.
	//The * operator denotes the pointer's underlying value.
	//This is known as "dereferencing" or "indirecting". Unlike C, Go has no pointer arithmetic.

	i, j := 42, 2701
	fmt.Println(i, j)

	p:= &i 				//point to i
	fmt.Println(p)		//the memory address
	fmt.Println(*p)		//the value of i

	*p = 21				//change the value of i through the pointer
	fmt.Println(*p, i)	//the value of *p and i are the same and both modified now

	p = &j				//now we point to j
	fmt.Println(*p)		//print the value of j
	*p = *p / 37		//divide j and modify through the pointer
	fmt.Println(j)		//the new value of j

	//2. Structs
	//A struct is a collection of fields
	fmt.Println(Vertex{1,2})

	//3. Struct Fields
	//Struct fields are accessed using a dot
	v := Vertex{1,2}	//declare an instance of the Vertex Struct
	v.X = 4				//Access X field of Vertex Struct
	fmt.Println(v.X, v.Y)

	//4. Pointers to Structs
	//Struct fields can also be accessed through a struct pointer
	//To access the X field, when we have a pointer 'p' we could write (*p).X
	//However, Go allows us to just write p.X without the explicit dereference
	//Go allows this in order to make the code more readable and less cumbersome

	g := &v			//declare a pointer g that points to v, an instantiation of our Vertex struct
	g.X = 1e9		//modify the value of X field of v through our pointer 'g'
	fmt.Println(v)

	//5. Struct Literals
	//A struct literal indicates a new struct declared with values listed in its fields
	//You can list a subset of the fields in a subset by using the 'Name:' syntax
	//With this syntax, the order of the named fields is irrelevant
	//The remaining fields will default to their nil values
	//The special prefix & returns a pointer to the struct value.
	var(
		v1 = Vertex{2,4} 	//has type Vertex
		v2 = Vertex{X: 1}	//Y:0 is implicit	
		v3 = Vertex{}		//X:0 and Y:0
		e = &Vertex{1,3} 	//has type *Vertex
	)

	fmt.Println(v1, e, v2, v3)

}