package fractran

type fraction struct {
	numerator   int
	denominator int
}

type multResult interface {
	multResult()
}

type success int

func (success) multResult() {}

type failure struct{}

func (failure) multResult() {}

func mult(f fraction, n int) multResult {
	switch (n * f.numerator) % f.denominator {
	case 0:
		return success(n * f.numerator / f.denominator)
	default:
		return failure{}
	}
}

type program []fraction

func exec(p program, input int) []int {
	var run func([]fraction, []int) []int

	run = func(f []fraction, inputs []int) []int {
		if len(f) == 0 {
			return inputs
		}
		switch r := mult(f[0], inputs[0]).(type) {
		case success:
			return run(p, prepend(int(r), inputs))
		case failure:
			return run(f[1:], inputs)
		}
		return nil
	}

	return run(p, []int{input})
}
