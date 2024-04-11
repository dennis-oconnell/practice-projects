package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("you can't divide by zero! sorry")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error){
	if a < 0 {
		return 0, errors.New("(no real number squared equals a negative number, so taking the square root of a negative number is not a valid operation")
	}
	return math.Sqrt(a), nil
}