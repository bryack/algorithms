package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkSum[K comparable, V Number](t *testing.T, m map[K]V, want V) {
	t.Helper()
	got := SumMapValues(m)
	assert.Equal(t, want, got)
}

func TestSumMapValues(t *testing.T) {
	t.Run("string int", func(t *testing.T) {
		tests := []struct {
			name string
			m    map[string]int
			want int
		}{
			{"empty", map[string]int{}, 0},
			{"single", map[string]int{"a": 5}, 5},
			{"multiple", map[string]int{"a": 1, "b": 2, "c": 3}, 6},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkSum(t, tt.m, tt.want)
			})
		}
	})

	t.Run("int_float64", func(t *testing.T) {
		tests := []struct {
			name string
			m    map[int]float64
			want float64
		}{
			{"empty", map[int]float64{}, 0},
			{"single", map[int]float64{1: 1.5}, 1.5},
			{"multiple", map[int]float64{1: 1.5, 2: 2.5}, 4.0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkSum(t, tt.m, tt.want)
			})
		}
	})

	t.Run("uint_int32", func(t *testing.T) {
		tests := []struct {
			name string
			m    map[uint]int32
			want int32
		}{
			{"empty", map[uint]int32{}, 0},
			{"single", map[uint]int32{1: 1}, 1},
			{"multiple", map[uint]int32{1: 1, 2: 2}, 3},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkSum(t, tt.m, tt.want)
			})
		}
	})
}
