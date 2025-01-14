package day3

type multiplication struct{ a, b int }

type program struct {
	muls []multiplication
}

func (p program) calculateMultiplications() (result int) {
	for _, mul := range p.muls {
		result += mul.a * mul.b
	}
	return result
}
