package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var c, python, java bool

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x 
	return 
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func main() {
	fmt.Println("1. My favorite number is", rand.Intn(10))

	
	//Printf stands for print formatted. %g is a format specifier used for floating point numbers
	fmt.Printf("2. Now you have %g problems. \n", math.Sqrt(7))

	//In Go, a name is exported if it begins with a capital letter. 
	//Pi must be written with a capital p for the program to function and for Pi to exported from the math package.
	fmt.Println("3. Pi can be written as", math.Pi)

	//A function can take zero or more arguments.
	//Notice that the type comes after the variable name.
	fmt.Println("4. Two plus four is equal to", add(2,4))

	//A function can return any number of results. The swap function returns two strings.
	a, b := swap("hello", "world")
	fmt.Println("5. If I swap the words 'hello' and 'world' I get the following:",a, b)

	//Step 6 was a matter of reformatting the add function
	fmt.Println("6. Step 6 was a matter of reformatting previous work")

	//Go's return values may be named. If so, they are treated as variables defined at the top of the function
	//A return statement without arguments returns the named return values. This is known as a "naked" return.
	//Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
	fmt.Print("7. If I wanted to split a number in two such as the number 17, it would look like:")
	fmt.Print(split(17))
	fmt.Print("\n")

	//The var statement declares a list of variables; as in function argument lists, the type is last.
	//A var statement can be at package or function level. 
	//Var c, python, java are booleans declared at the package level.
	//Var i is at the function level
	var i int
	fmt.Println("8. Here are the values of some variables that I declared: ",i, c, python, java)

	//Variables with Initializers
	//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var j, k int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println("9. Here are some variables declared with initializers:", j, k, c, python, java)

	//Short Variable Declarations can only be made inside a function
	// ':=' can be used in place of var declarations with implicit type
	l:= 3;
	fmt.Println("10. This variable was declared with the short assingment:", l)

	//the following are Go's basic types
	//bool, string
	//int, int8, int16, int32,  int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte // alias for uint8
	//rune // alias for int32
	// represents a Unicode code point
	//float32 float64
	//complex64 complex128
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("11. This step involves exploring Go's basic types-> Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("11. This step involves exploring Go's basic types-> Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("11. This step involves exploring Go's basic types-> Type: %T Value: %v\n", z, z)

	//Variables declared without an explicit initial value are given their zero value.
	// The zero value is: 0 for numeric types, false for the boolean type, and "" (the empty string) for strings.
	var aZero int
	var bZero float64
	var cZero bool
	var dZero string
	fmt.Printf("12. The Zero values for int, float64, bool, and string are the following: %v %v %v %q\n", aZero, bZero, cZero, dZero)

	//Type Conversions: The expression T(v) converts the value v to the type T.
	//Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	iCon := 42
	fCon := float64(iCon)
	uCon := uint(fCon)
	fmt.Println("13. Type Conversions Part 1: ", iCon, fCon, uCon)
	var xCon, yCon = 3, 4
	var gCon float64 = math.Sqrt(float64(xCon*xCon + yCon*yCon))
	var zCon uint = uint(gCon)
	fmt.Println("13. Type Conversions Part 2: ", xCon, yCon, gCon, zCon)

	//Type Inference
	//When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	var jInf int
	iInf := jInf // iInf is also an int
	fmt.Println("14. Type Inference Part 1: ",iInf)

	//when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant	
	eInf := 42           // int
	fInf := 3.142        // float64
	gInf := 0.867 + 0.5i // complex128
	fmt.Println("14. Type Inference Part 2: ", eInf, fInf, gInf)

	//when we change the initial value of vChange below, the type is affected
	v := 42i // change me!
	fmt.Printf("14. Type Inference Part 3: v is of type %T\n", v)

	//Constants are declared like variables but with the const keyword. They cannot be declared using ':=' syntax and they can be character, string, boolean, or numeric values.
	const World = "世界"
	const Pi = 3.14	
	const Truth = true
	fmt.Println("15. Hello", World)
	fmt.Println("15. Happy", Pi, "Day")
	fmt.Println("15. Go rules?", Truth)

	//Numeric constants are high-precision values. An untyped constant takes the type needed by its context.
	//(An int can store at maximum a 64-bit integer, and sometimes less.)
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)
	fmt.Println("16. ", needInt(Small))
	fmt.Println("16. ", needFloat(Small))
	fmt.Println("16. ", needFloat(Big))
}