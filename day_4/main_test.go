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
			want:           Fraction{3, 3},
		},
		{
			name:           "1/5 + 3/5 = 4/5",
			firstFraction:  Fraction{1, 5},
			secondFraction: Fraction{3, 5},
			want:           Fraction{4, 5},
		},
		{
			name:           "1/2 + 1/4 = 3/4",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{1, 4},
			want:           Fraction{3, 4},
		},
		{
			name:           "1/2 + 1/12 = 7/12",
			firstFraction:  Fraction{1, 2},
			secondFraction: Fraction{1, 12},
			want:           Fraction{7, 12},
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
		a := 1836311903
		b := 1134903170
		res := findGreatestCommonDivisor(a, b)
		assert.Equal(t, 1, res)
	})
}
