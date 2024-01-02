package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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

	//6. Arrays 
	//The Type [n]T is an array of n values of type T
	//An array's length is part of its type, so arrays can't be resized
	
	var a [10]int		//declared a variable as an array of ten integers with default values	
	fmt.Println(a)

	var b [2]string
	b[0] = "Hi"
	b[1] = "Planet"
	fmt.Println(b[0],b[1])
	fmt.Println(b)

	sixPrimes := [6]int{2,3,5,7,11,13}
	fmt.Println(sixPrimes)

	//7. Slices
	//An array is of a fixed size, but slices are dynamically-sized!
	//A slice is a dynamically-sized, flexible view into the elements of an array.
	//In practice, slices are much more common than arrays.
	//The type []T is slice with elements of type T.

	//A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low : high]
	//This selects a half-open range which includes the first element but excludes the last one
	//This example creates a slice that includes elements 1 through 3 of a: 'a[1:4]'

	var pSlice []int = sixPrimes[1:4] 
	fmt.Println(pSlice)

	//8. Slices Are Like References To Arrays
	//A slice alone doesn't store any data, it just describes a section of an underlying array
	//Changing the elements of a slice modifies the corresponding elements of the underlying array
	//Other slices that share the same underlying array will also see these changes

	beatles:=[4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(beatles)

	aBeatles := beatles[0:2]
	bBeatles := beatles[2:4]

	fmt.Println(aBeatles,bBeatles)

	bBeatles[0] = "XXX" 	//modify the element of one slice and you also modify the element of the underlying array

	fmt.Println(aBeatles,bBeatles)
	fmt.Println(beatles)

	//9. Slice literals 
	//A slice literal is like an array literal without the length
	//Here is an array literal: '[3]bool{true,false,false}'
	//This creates the same array as above, then builds a slice that references it: '[]bool{true, true, false}'
	qNums := []int{2,3,5,7,11,13}
	fmt.Println(qNums)

	rBoos:=[]bool{true,false,true,false,true,false}
	fmt.Println(rBoos)
	
	sStru:=[]struct {
		i int
		b bool
	} {
		{2,true},
		{3,false},
		{5,true},
	}

	fmt.Println(sStru)

	//10. Slice Defaults
	//When slicing, you do not need to include high or low bounds, instead you can use their defaults
	//The default is zero for the low bound and the length of the slice for the high bound
	//For the array: 'var a [10]int' the expressions below are all equivalent:
	//a[0,10]
	//a[:10]
	//a[0:]
	//a[:]

	fmt.Println(sixPrimes[1:4])
	fmt.Println(sixPrimes[:2])
	fmt.Println(sixPrimes[1:])

	//11. Slice Length and Capacity
	//A slice has both length and capacity
	//The length of the slice is the number of elements that it contains
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice
	//The length and capacity of a slice 'a' can be obtained with the expressions len(s) and cap(s)
	//You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	
	printSlice(sixPrimes[:])
	printSlice(sixPrimes[3:4])
	printSlice(sixPrimes[2:])
	printSlice(sixPrimes[:0])

	//12. The nil value of a slice is 'nil'
	//A nil slice has a length and capacity of zero and has no underlying array
	var emp []int
	fmt.Println(emp,len(emp), cap(emp))
	if(emp == nil){
		fmt.Println("the 'emp' array is empty, its a nil slice! It has a length of 0 and a capacity of 0 and no underlying array!")
	}

	//13. Creating a Slice With Make
	// Slices can be made with built in function 'make' function
	//This is how we can make dynamically sized arrays in Go
	//The 'make' function allocates a zeroed array and returns a slice that refers to that array 
	// a:= make([]int, 5)         //len(a) = 5
	//You can specify a capacity, if you pass a third argument to make
	// b:=make([]int, 5) 		//len(b)=0, cap(b)=5
	// b=b[:cap(b)]				//len(b)=5, cap(b)=5
	// b=b[1:]					//len(b)=4, cap(b)=4

	mA:=make([]int,5)
	printSlice(mA)
	
	mB:=make([]int,0,5)
	printSlice(mB)

	mC:=mB[:2]
	printSlice(mC)

	mD:=mC[2:5]
	printSlice(mD)
}