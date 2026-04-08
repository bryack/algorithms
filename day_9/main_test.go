package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "1 - I", input: 1, want: "I"},
		{name: "2 - II", input: 2, want: "II"},
		{name: "3 - III", input: 3, want: "III"},
		{name: "58 - LVIII", input: 58, want: "LVIII"},
		{name: "63 - LXIII", input: 63, want: "LXIII"},
		{name: "22 - XXII", input: 22, want: "XXII"},
		{name: "33 - XXXIII", input: 33, want: "XXXIII"},
		{name: "44 - XLIV", input: 44, want: "XLIV"},
		{name: "1984 - MCMXCIV", input: 1984, want: "MCMLXXXIV"},
		{name: "1994 - MCMXCIV", input: 1994, want: "MCMXCIV"},
		{name: "3599 - MMMDXCIX", input: 3599, want: "MMMDXCIX"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intToRoman(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
