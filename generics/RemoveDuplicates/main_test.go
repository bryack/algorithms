package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		tests := []struct {
			name string
			slc  []int
			want []int
		}{
			{name: "empty", slc: []int{}, want: []int{}},
			{name: "single", slc: []int{1}, want: []int{1}},
			{name: "have dupl", slc: []int{1, 1, 2, -2, 2, 3, 3, 4}, want: []int{1, 2, -2, 3, 4}},
			{name: "no dupl", slc: []int{3, 5, -1, 8, 10}, want: []int{3, 5, -1, 8, 10}},
			{name: "all same", slc: []int{7, 7, 7, 7}, want: []int{7}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := RemoveDuplicates(tt.slc)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("uint slice", func(t *testing.T) {
		tests := []struct {
			name string
			slc  []uint
			want []uint
		}{
			{name: "empty", slc: []uint{}, want: []uint{}},
			{name: "single", slc: []uint{1}, want: []uint{1}},
			{name: "have dupl", slc: []uint{1, 1, 2, 2, 2, 3, 3, 4}, want: []uint{1, 2, 3, 4}},
			{name: "no dupl", slc: []uint{3, 5, 1, 8, 10}, want: []uint{3, 5, 1, 8, 10}},
			{name: "all same", slc: []uint{7, 7, 7, 7}, want: []uint{7}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := RemoveDuplicates(tt.slc)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		tests := []struct {
			name string
			slc  []float64
			want []float64
		}{
			{name: "empty", slc: []float64{}, want: []float64{}},
			{name: "single", slc: []float64{1}, want: []float64{1}},
			{name: "have dupl", slc: []float64{1, 1, 1.5, 2, 2, 2, 3, 3, 4}, want: []float64{1, 1.5, 2, 3, 4}},
			{name: "no dupl", slc: []float64{3, 5, -1, 8, 10}, want: []float64{3, 5, -1, 8, 10}},
			{name: "all same", slc: []float64{7.5, 7.5, 7.5, 7.5}, want: []float64{7.5}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := RemoveDuplicates(tt.slc)
				assert.Equal(t, tt.want, got)
			})
		}
	})
	t.Run("string slice", func(t *testing.T) {
		tests := []struct {
			name string
			slc  []string
			want []string
		}{
			{name: "empty", slc: []string{}, want: []string{}},
			{name: "single", slc: []string{"one"}, want: []string{"one"}},
			{name: "have dupl", slc: []string{"go", "rust", "go", "c", "rust"}, want: []string{"go", "rust", "c"}},
			{name: "no dupl", slc: []string{"go", "rust", "c"}, want: []string{"go", "rust", "c"}},
			{name: "all same", slc: []string{"go", "go", "go", "go"}, want: []string{"go"}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := RemoveDuplicates(tt.slc)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
