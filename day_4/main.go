package main

type Fraction struct {
	Numerator   int
	Denominator int
}

func (f Fraction) Add(other Fraction) Fraction {
	if f.Denominator == other.Denominator {
		return Fraction{f.Numerator + other.Numerator, f.Denominator}
	}

	if other.Denominator%f.Denominator == 0 {
		mult := other.Denominator / f.Denominator
		return Fraction{f.Numerator*mult + other.Numerator, other.Denominator}
	}

	a, b := max(f.Denominator, other.Denominator), min(f.Denominator, other.Denominator)
	gcd := findGreatestCommonDivisor(a, b)
	multiple := a * b / gcd
	f.Numerator = multiple / f.Denominator * f.Numerator
	other.Numerator = multiple / other.Denominator * other.Numerator
	return Fraction{f.Numerator + other.Numerator, multiple}
}

func findGreatestCommonDivisor(a, b int) int {
	for {
		divisor := a % b
		a = b
		b = divisor
		if b == 0 {
			return a
		}
	}
}
