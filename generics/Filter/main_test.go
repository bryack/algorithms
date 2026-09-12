package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8
}

func isEven[T Number](n T) bool {
	return n%2 == 0
}

func TestFilter_IsEven(t *testing.T) {
	t.Run("int_slice", func(t *testing.T) {
		tests := []struct {
			name string
			slc  []int
			want []int
		}{
			{name: "empty", slc: []int{}, want: []int{}},
			{name: "only even", slc: []int{2, 4, 6, 8}, want: []int{2, 4, 6, 8}},
			{name: "mixed", slc: []int{2, 5, 6, 7}, want: []int{2, 6}},
			{name: "only odd", slc: []int{5, 7, 1, -1}, want: []int{}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.slc, isEven)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("uint_slice", func(t *testing.T) {
		tests := []struct {
			name string
			slc  []uint
			want []uint
		}{
			{name: "empty", slc: []uint{}, want: []uint{}},
			{name: "only even", slc: []uint{2, 4, 6, 8}, want: []uint{2, 4, 6, 8}},
			{name: "mixed", slc: []uint{2, 5, 6, 7}, want: []uint{2, 6}},
			{name: "only odd", slc: []uint{5, 7, 1, 23}, want: []uint{}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.slc, isEven)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
