// package calculator performs simple arithmetic
package calculator

import (
	"errors"
	"math"
)

//Add takes two numbers and returns the result of adding them together
func Add(a,b float64) float64 {
	return a + b
}

//Subtract takes two numbers a and b
//returns the result of subtracting b from a 
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b 
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b , nil
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("sqrt of a negative number is not allowed")
	}
	return math.Sqrt(x), nil
}
/* AN OUTLINE OF THE DEVELOPMENT PROCESS
	I.E. The "red, green, refactor" cycle
	1. Start with the test for the function (before the function itself exists)
	2. See the test fail with a compilation error bc the func doesn't exist yet
	3. Write minimum code necessary to make the test compile and fail
	4. Ensure that the test fails for the reason we expect (look for accuracy and informative)
	5. Write minimum code necessary to make the test pass
	6. Tweek and improve code if necessary 
*/