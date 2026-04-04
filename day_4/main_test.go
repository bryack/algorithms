package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFraction_Add(t *testing.T) {
	firstFraction := Fraction{1, 3}
	secondFraction := Fraction{2, 3}
	result := firstFraction.Add(secondFraction)
	want := Fraction{3, 3}

	assert.Equal(t, want, result)
}
