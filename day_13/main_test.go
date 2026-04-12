package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name       string
		firstList  []int
		secondList []int
		want       []int
	}{
		{
			name:       "interleaved values",
			firstList:  []int{1, 2, 4},
			secondList: []int{1, 3, 4},
			want:       []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:       "both empty",
			firstList:  []int{},
			secondList: []int{},
			want:       []int{},
		},
		{
			name:       "first empty, second single element",
			firstList:  []int{},
			secondList: []int{0},
			want:       []int{0},
		},
		{
			name:       "first single element, second empty",
			firstList:  []int{0},
			secondList: []int{},
			want:       []int{0},
		},
		{
			name:       "single element each",
			firstList:  []int{1},
			secondList: []int{2},
			want:       []int{1, 2},
		},
		{
			name:       "all first smaller than second",
			firstList:  []int{1, 2, 3},
			secondList: []int{4, 5, 6},
			want:       []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:       "all second smaller than first",
			firstList:  []int{4, 5, 6},
			secondList: []int{1, 2, 3},
			want:       []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:       "duplicates across both lists",
			firstList:  []int{1, 2, 2, 3},
			secondList: []int{2, 2, 4},
			want:       []int{1, 2, 2, 2, 2, 3, 4},
		},
		{
			name:       "one element vs many",
			firstList:  []int{5},
			secondList: []int{1, 2, 3, 4, 6},
			want:       []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:       "negative numbers",
			firstList:  []int{-5, -3, 0},
			secondList: []int{-4, -2, 1},
			want:       []int{-5, -4, -3, -2, 0, 1},
		},
		{
			name:       "identical lists",
			firstList:  []int{1, 2, 3},
			secondList: []int{1, 2, 3},
			want:       []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.firstList)
			list2 := sliceToList(tt.secondList)
			resList := mergeTwoLists(list1, list2)

			result := listToSlice(resList)
			assert.Equal(t, tt.want, result)
		})
	}
}
