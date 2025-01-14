package day3

type program struct {
	muls []multiplication
}

func (p program) addMultiplications() (result int) {
	for _, mul := range p.muls {
		result += mul.a * mul.b
	}
	return result
}

type multiplication struct{ a, b int }
