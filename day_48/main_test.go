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
