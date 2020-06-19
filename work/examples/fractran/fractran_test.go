package fractran

import (
	"testing"
)

func TestFractran(t *testing.T) {
	p := program{
		{91, 33},
		{11, 13},
		{1, 11},
		{399, 34},
		{17, 19},
		{1, 17},
		{2, 7},
		{187, 5},
		{1, 3},
	}

	input := 31250

	res := exec(p, input)
	if res[0] != 8192 {
		t.Errorf("did not get correct result. expected=%v, got=%v", 8192, res[0])
	}
}
