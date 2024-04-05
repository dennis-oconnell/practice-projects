package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type typeCase struct {
		a, b float64
		want float64
	}

	testCases := []typeCase{
		{a: 1, b: 1, want: 2.0},
		{a: 2, b: 1, want: 3.0},
		{a: .75, b: .75, want: 1.5},
	}

	for key, tc := range testCases{

		got := calculator.Add(tc.a, tc.b)

		if testCases[key].want != got{
			t.Errorf("want: %f, got: %f", testCases[key].want, got)
		}
	}

	
}

func TestSubtract(t *testing.T){
	t.Parallel()

	type testCase struct{
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a:1, b:1, want:0},
		{a:2, b:4, want:-2},
		{a:1.5, b:.75, want:.75},
	}

	for key, tc := range testCases{

		got := calculator.Subtract(tc.a, tc.b)

		if testCases[key].want != got{
			t.Errorf("want: %f, got: %f", testCases[key].want, got)
		}
	}
}

func TestMultiply(t *testing.T){
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 1,b:-1, want:-1},
		{a:2, b:4, want: 8},
		{a:1.5, b:.75, want:1.125},
	}

	for key, tc := range testCases{

		got := calculator.Multiply(tc.a,tc.b)

		if testCases[key].want != got{
			t.Errorf("want: %f, got: %f", testCases[key].want, got)
		}
	}
}

func TestDivide(t *testing.T){
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 12,b:3, want:4},
		{a:-20, b:4, want: -5},
		{a:2.5, b:.5, want:5},
	}

	for key, tc := range testCases{

		got, err := calculator.Divide(tc.a,tc.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if testCases[key].want != got{
			t.Errorf("want: %f, got: %f", testCases[key].want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T){
	t.Parallel()

	a:= 1.0
	b:= 0.0

	_, err := calculator.Divide(a, b)

	if(err == nil){
		t.Error("Invalid input, wanted an error to be thrown, but error was not thrown!")
	}

}

func closeEnough (a, b, tolerance float64) bool {
	return Math.Abs()
}