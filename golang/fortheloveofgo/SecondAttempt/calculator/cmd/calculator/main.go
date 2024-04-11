package main

import (
	"calculator"
	"fmt"
)

func main() {
	outputs:= []float64{
		calculator.Add(2,2),
		
	}
	
	for _, o := range outputs{
		fmt.Println(o)
	}
	
}