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
	} else if f.Denominator%other.Denominator == 0 {
		mult := f.Denominator / other.Denominator
		return reduce(other.Numerator*mult+f.Numerator, f.Denominator)
	}

	gcd := findGreatestCommonDivisor(f.Denominator, other.Denominator)
	multiple := f.Denominator * other.Denominator / gcd
	f.Numerator = multiple / f.Denominator * f.Numerator
	other.Numerator = multiple / other.Denominator * other.Numerator
	return reduce(f.Numerator+other.Numerator, multiple)
}

func (f Fraction) Subtract(other Fraction) Fraction {
	if f.Denominator == other.Denominator {
		return reduce(f.Numerator-other.Numerator, f.Denominator)
	}

	if other.Denominator%f.Denominator == 0 {
		mult := other.Denominator / f.Denominator
		return reduce(f.Numerator*mult-other.Numerator, other.Denominator)
	} else if f.Denominator%other.Denominator == 0 {
		mult := f.Denominator / other.Denominator
		return reduce(f.Numerator-other.Numerator*mult, f.Denominator)
	}

	gcd := findGreatestCommonDivisor(f.Denominator, other.Denominator)
	multiple := f.Denominator * other.Denominator / gcd
	f.Numerator = multiple / f.Denominator * f.Numerator
	other.Numerator = multiple / other.Denominator * other.Numerator

	return reduce(f.Numerator-other.Numerator, multiple)
}

func (f Fraction) Multiply(other Fraction) Fraction {
	return reduce(f.Numerator*other.Numerator, f.Denominator*other.Denominator)
}

func (f Fraction) Divide(other Fraction) Fraction {
	return reduce(f.Numerator*other.Denominator, f.Denominator*other.Numerator)
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

func gcdOfStrings(s, t string) string {
	if s+t == t+s {
		gcd := findGreatestCommonDivisor(len(s), len(t))
		return s[:gcd]
	}
	return ""
}
