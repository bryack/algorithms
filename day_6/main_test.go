package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusFahrenheit(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{
			name:  "0 + 32",
			input: 0,
			want:  32,
		},
		{
			name:  "+7 +44.6",
			input: 7,
			want:  44.6,
		},
		{
			name:  "-40 -40",
			input: -40,
			want:  -40,
		},
		{
			name:  "-1 +30.2",
			input: -1,
			want:  30.2,
		},
		{
			name:  "+1000000000 +1800000032",
			input: 1000000000,
			want:  1800000032,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusFahrenheit(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFahrenheitCelsius(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{
			name:  "0 + 32",
			input: 32,
			want:  0,
		},
		{
			name:  "+7 +44.6",
			input: 44.6,
			want:  7,
		},
		{
			name:  "-40 -40",
			input: -40,
			want:  -40,
		},
		{
			name:  "-1 +30.2",
			input: 30.2,
			want:  -1,
		},
		{
			name:  "+1000000000 +1800000032",
			input: 1800000032,
			want:  1000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FahrenheitCelsius(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestRUBUSD(t *testing.T) {
	tests := []struct {
		name   string
		rub    int
		kopeck int
		rate   float64
		wantD  int
		wantC  int
	}{
		{
			name:   "100,0 - 1,2661",
			rub:    100,
			kopeck: 0,
			rate:   78.9801,
			wantD:  1,
			wantC:  27,
		},
		{
			name:   "100,50 - 1,2725", // 100.50 рублей / 78.9801 ≈ 1.2725
			rub:    100,
			kopeck: 50,
			rate:   78.9801,
			wantD:  1,
			wantC:  27,
		},
		{
			name:   "20,856 - 0,2641", // 100.50 рублей / 78.9801 ≈ 1.2725
			rub:    20,
			kopeck: 85,
			rate:   78.9801,
			wantD:  0,
			wantC:  26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, c := RUBUSD(tt.rub, tt.kopeck, tt.rate)
			assert.Equal(t, tt.wantD, d)
			assert.Equal(t, tt.wantC, c)
		})
	}
}

func TestUSDRUB(t *testing.T) {
	tests := []struct {
		name  string
		dol   int
		cen   int
		rate  float64
		wantR int
		wantK int
	}{
		{
			name:  "100,0 - 1,2661",
			dol:   1,
			cen:   27,
			rate:  0.0127,
			wantR: 100,
			wantK: 0,
		},
		{
			name:  "0,26 - 20,47",
			dol:   0,
			cen:   26,
			rate:  0.0127,
			wantR: 20,
			wantK: 47,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, c := USDRUB(tt.dol, tt.cen, tt.rate)
			assert.Equal(t, tt.wantR, d)
			assert.Equal(t, tt.wantK, c)
		})
	}
}

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "A", input: 1, want: "A"},
		{name: "Z", input: 26, want: "Z"},
		{name: "AA", input: 27, want: "AA"},
		{name: "AB", input: 28, want: "AB"},
		{name: "BB", input: 54, want: "BB"},
		{name: "ZY", input: 701, want: "ZY"},
		{name: "AZ", input: 52, want: "AZ"},
		{name: "DD", input: 108, want: "DD"},
		{name: "AAA", input: 703, want: "AAA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToTitle(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
