package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{name: "classic", nums: []int{0, 1, 2, 4, 5, 7}, want: []string{"0->2", "4->5", "7"}},
		{name: "test2", nums: []int{0, 2, 3, 4, 6, 8, 9}, want: []string{"0", "2->4", "6", "8->9"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := summaryRanges(tt.nums)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestСompress(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  int
	}{
		{name: "classic", chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, want: 6},
		{name: "test2", chars: []byte{'a'}, want: 1},
		{name: "test3", chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, want: 4},
		{name: "no_repeat", chars: []byte{'a', 'b', 'c'}, want: 3},
		{name: "count_9", chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, want: 2},
		{name: "count_10", chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, want: 3},
		{name: "mixed", chars: []byte{'a', 'a', 'b', 'b', 'a'}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := compress(tt.chars)
			assert.Equal(t, tt.want, res)
		})
	}
}
