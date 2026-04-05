package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFraction_Add(t *testing.T) {
	tests := []struct {
		name           string
		firstFraction  Fraction
		secondFraction Fraction
		want           Fraction
	}{
		{
			name:           "1/3 + 2/3 = 3/3",
			firstFraction:  Fraction{1, 3},
			secondFraction: Fraction{2, 3},
			want:           Fraction{1, 1},
		},
		{
			name:           "1/4 + 1/4 = 1/2",
			firstFraction:  Fraction{1, 4},
			secondFraction: Fraction{1, 4},
			want:           Fraction{1, 2},
		},
		{
			name:           "1/2 + 1/4 = 3/4",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{1, 4},
			want:           Fraction{3, 4},
		},
		{
			name:           "2/6 + 1/2 = 5/6",
			firstFraction:  Fraction{2, 6},
			secondFraction: Fraction{1, 2},
			want:           Fraction{5, 6},
		},
		{
			name:           "1/2 + 2/12 = 2/3",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{2, 12},
			want:           Fraction{2, 3},
		},
		{
			name:           "1/2 + 1/3 = 5/6",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{1, 3},
			want:           Fraction{5, 6},
		},
		{
			name:           "1/10 + 1/7 = 17/70",
			firstFraction:  Fraction{1, 10},
			secondFraction: Fraction{1, 7},
			want:           Fraction{17, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.firstFraction.Add(tt.secondFraction)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindGreatestCommonDivisor(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		a := 48
		b := 18
		res := findGreatestCommonDivisor(a, b)
		assert.Equal(t, 6, res)
	})

	t.Run("coprime numbers", func(t *testing.T) {
		a := 3
		b := 2
		res := findGreatestCommonDivisor(a, b)
		assert.Equal(t, 1, res)
	})
	t.Run("stack overflow", func(t *testing.T) {
		a := 1134903170
		b := 1836311903
		res := findGreatestCommonDivisor(a, b)
		assert.Equal(t, 1, res)
	})
}

func TestFraction_Reduce(t *testing.T) {
	fraction := Fraction{24, 48}
	res := reduce(fraction.Numerator, fraction.Denominator)
	want := Fraction{1, 2}
	assert.Equal(t, want, res)
}

func TestFraction_Subtract(t *testing.T) {
	tests := []struct {
		name           string
		firstFraction  Fraction
		secondFraction Fraction
		want           Fraction
	}{
		{
			name:           "2/3 - 1/3 = 1/3",
			firstFraction:  Fraction{2, 3},
			secondFraction: Fraction{1, 3},
			want:           Fraction{1, 3},
		},
		{
			name:           "10/15 - 7/15 = 1/5",
			firstFraction:  Fraction{10, 15},
			secondFraction: Fraction{7, 15},
			want:           Fraction{1, 5},
		},
		{
			name:           "1/2 - 2/6 = 1/6",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{2, 6},
			want:           Fraction{1, 6},
		},
		{
			name:           "9/10 - 1/5 = 7/10",
			firstFraction:  Fraction{9, 10},
			secondFraction: Fraction{1, 5},
			want:           Fraction{7, 10},
		},
		{
			name:           "1/3 - 2/10 = 2/15",
			firstFraction:  Fraction{1, 3},
			secondFraction: Fraction{2, 10},
			want:           Fraction{2, 15},
		},
		{
			name:           "7/10 - 1/7 = 39/70",
			firstFraction:  Fraction{7, 10},
			secondFraction: Fraction{1, 7},
			want:           Fraction{39, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.firstFraction.Subtract(tt.secondFraction)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFraction_Multiply(t *testing.T) {
	tests := []struct {
		name           string
		firstFraction  Fraction
		secondFraction Fraction
		want           Fraction
	}{
		{
			name:           "2/3 * 1/3 = 2/9",
			firstFraction:  Fraction{2, 3},
			secondFraction: Fraction{1, 3},
			want:           Fraction{2, 9},
		},
		{
			name:           "1/2 * 2/6 = 1/6",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{2, 6},
			want:           Fraction{1, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.firstFraction.Multiply(tt.secondFraction)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFraction_Divide(t *testing.T) {
	tests := []struct {
		name           string
		firstFraction  Fraction
		secondFraction Fraction
		want           Fraction
	}{
		{
			name:           "2/3 / 1/3 = 2/1",
			firstFraction:  Fraction{2, 3},
			secondFraction: Fraction{1, 3},
			want:           Fraction{2, 1},
		},
		{
			name:           "1/2 / 2/7 = 7/4",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{2, 7},
			want:           Fraction{7, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.firstFraction.Divide(tt.secondFraction)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFraction_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name         string
		firstString  string
		secondString string
		want         string
	}{
		{
			name:         "1 example",
			firstString:  "ABCAB",
			secondString: "ABC",
			want:         "",
		},
		{
			name:         "2 example",
			firstString:  "ABCABC",
			secondString: "ABC",
			want:         "ABC",
		},
		{
			name:         "3 example",
			firstString:  "ABCABCABC",
			secondString: "ABC",
			want:         "ABC",
		},
		{
			name:         "4 example",
			firstString:  "ABCACABC",
			secondString: "ABC",
			want:         "",
		},
		{
			name:         "4 example",
			firstString:  "ABABAB",
			secondString: "ABAB",
			want:         "AB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcdOfStrings(tt.firstString, tt.secondString)
			assert.Equal(t, tt.want, result)
		})
	}
}

// func TestFraction_gcd(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input string
// 		want  string
// 	}{
// 		{
// 			name:  "simple",
// 			input: "ABAB",
// 			want:  "AB",
// 		},
// 		{
// 			name:  "simple",
// 			input: "ABABAB",
// 			want:  "AB",
// 		},
// 		{
// 			name:  "simple",
// 			input: "ABC",
// 			want:  "ABC",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := gcd(tt.input)
// 			assert.Equal(t, tt.want, result)
// 		})
// 	}
// }
