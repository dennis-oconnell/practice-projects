package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	} 

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add (%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 0, b: 1, want: -1},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract (%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 0, b: 1, want: 0},
		{a: 5, b: 5, want: 25},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply (%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

//When comparing floating point values, remember that they have limited precision
//So, we need a looser comparison
func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 1, b: 3, want: .33333},
		{a: -4, b: -2, want: 2},
		{a: 25, b: 5, want: 5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
				//We use Fatalf here rather than Errorf
				//Errorf marks the test as failed, but continues executing the rest of the test
				//Fatalf exits the test immediatley
				//Fatalf says that things are so broken they are no longer worth continuing
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, .0001) {
			t.Errorf("Divide (%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalidInput(t *testing.T) {
	t.Parallel()
	
	type testCase struct {
		a, b float64
	}

	testCases := []testCase{
		{a: 2, b: 0,},
		{a: -4, b: 0,},
		{a: 25, b: 0,},
	}

	for _, tc := range testCases {
		_, err := calculator.Divide(tc.a, tc.b)
		if err == nil {
			t.Errorf("Want error for invalid input, got nil, i.e. no error")
		}
	}
}

func TestSqrt(t *testing.T){
	t.Parallel()

	type testCase struct {
		x float64
		want float64
	}

	testCases := []testCase{
		{x: 4, want:2,},
		{x: 8, want:2.828,},
		{x: 16, want:4,},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.x)

		if err!= nil {
			t.Fatalf("Error!")
		}

		if !closeEnough(tc.want, got, .001){
			t.Errorf("Sqrt of %f: want %f, got %f", tc.x, tc.want, got)
		}
	}
}

func TestSqrtInvalidInput(t *testing.T) {
	t.Parallel()

	_, err := calculator.Sqrt(-1)

	if err == nil {
		t.Errorf("Want error for invalid input, got nil, i.e. no error")
	}
}

/*
	THINGS TO KNOW ABOUT TESTS IN GO!
		1. Each test in Go is a function
		2. Tests must be in a file whose name ends in "_test.go"
		3. The name of a test function must begin with the word "Test"
		4. The test function must accept a parameter whose type is "*testing.T"
		5. Test funcs don't return anything

		6. t.Parallel() is a standard prelude to tests; it tells Go to run this test concurrently with other tests, which saves time
*/

/* NOTES ON FOR LOOPS
	You can use range to iterate over an array, a slice, a string, a map, or a channel.
    range returns one (channel) or two values (array, slice, string and map).

		for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
			// for each pair in the map, print key and value
			fmt.Printf("key=%s, value=%d\n", key, value)
		}

	If you only need the value, use the underscore as the key

		for _, name := range []string{"Bob", "Bill", "Joe"} {
			fmt.Printf("Hello, %s\n", name)
		}
*/