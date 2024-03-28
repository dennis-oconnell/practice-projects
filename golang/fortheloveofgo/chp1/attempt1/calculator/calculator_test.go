package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	want:= []float64{
		2.0,
		4.0,
		1.5,
	}

	got:= []float64{
		calculator.Add(1,1),
		calculator.Add(2,2),
		calculator.Add(.75,.75),
	}

	for key, want := range want{
		if want != got[key]{
			t.Errorf("want: %f, got: %f", want, got[key])
		}
	}

	
}

func TestSubtract(t *testing.T){
	t.Parallel()

	want:= []float64{
		0,
		-2.0,
		.75,
	}

	got:= []float64{
		calculator.Subtract(1,1),
		calculator.Subtract(2,4),
		calculator.Subtract(1.5,.75),
	}
	
	for key, want := range want{
		if want != got[key]{
			t.Errorf("want: %f, got: %f", want, got[key])
		}
	}
}