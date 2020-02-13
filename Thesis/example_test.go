package example

import (
	"fmt"
	"testing"
)

func TestMult(t *testing.T) {
	tests := []struct {
		frac   Fraction
		n      int
		result MultResult
	}{
		{Fraction{3, 5}, 5, Success(3)},
		{Fraction{3, 8}, 2, Failure{}},
	}

	for _, tt := range tests {
		res := Mult(tt.frac, tt.n)
		if res != tt.result {
			t.Fatalf("results not equal, expected=%d, got=%d", tt.result, res)
		}
	}
}

func TestExec(t *testing.T) {
	testProgram := Program{
		Fraction{91, 33},
		Fraction{11, 13},
		Fraction{1, 11},
		Fraction{399, 34},
		Fraction{17, 19},
		Fraction{1, 17},
		Fraction{2, 7},
		Fraction{187, 5},
		Fraction{1, 3},
	}

	input := 31250
	res := Exec(testProgram, input)
	if res[0] != 8192 {
		t.Errorf("Did not get correct result. expected=%d, got=%d",
			8192,
			res[0])
	}
	fmt.Println(res)
}
