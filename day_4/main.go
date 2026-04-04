package main

type Fraction struct {
	Numerator   int
	Denominator int
}

func (f Fraction) Add(other Fraction) Fraction {
	if f.Denominator == other.Denominator {
		return reduce(f.Numerator+other.Numerator, f.Denominator)
	}

	if other.Denominator%f.Denominator == 0 {
		mult := other.Denominator / f.Denominator
		return reduce(f.Numerator*mult+other.Numerator, other.Denominator)
	}

	gcd := findGreatestCommonDivisor(f.Denominator, other.Denominator)
	multiple := f.Denominator * other.Denominator / gcd
	f.Numerator = multiple / f.Denominator * f.Numerator
	other.Numerator = multiple / other.Denominator * other.Numerator
	return reduce(f.Numerator+other.Numerator, multiple)
}

func findGreatestCommonDivisor(a, b int) int {
	a, b = max(a, b), min(a, b)
	for {
		divisor := a % b
		a = b
		b = divisor
		if b == 0 {
			return a
		}
	}
}

func reduce(a, b int) Fraction {
	multiple := findGreatestCommonDivisor(a, b)
	return Fraction{a / multiple, b / multiple}
}
