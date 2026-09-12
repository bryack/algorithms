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

func TestDivideArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "example 1 - can form pairs",
			nums:     []int{3, 2, 3, 2, 1, 1},
			expected: true,
		},
		{
			name:     "example 1 - can form pairs",
			nums:     []int{3, 2, 3, 2, 2, 2},
			expected: true,
		},
		{
			name:     "example 2 - cannot form pairs",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "two identical elements",
			nums:     []int{5, 5},
			expected: true,
		},
		{
			name:     "two different elements",
			nums:     []int{1, 2},
			expected: false,
		},
		{
			name:     "all same element even count",
			nums:     []int{7, 7, 7, 7},
			expected: true,
		},
		{
			name:     "all same element odd count",
			nums:     []int{7, 7, 7},
			expected: false,
		},
		{
			name:     "multiple pairs different values",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: true,
		},
		{
			name:     "one odd frequency breaks pairs",
			nums:     []int{1, 1, 2, 2, 3, 3, 4},
			expected: false,
		},
		{
			name:     "single element odd frequency among evens",
			nums:     []int{1, 1, 2, 2, 3, 3, 5},
			expected: false,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := divideArray(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinMaxGame(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single element",
			nums:     []int{3},
			expected: 3,
		},
		{
			name:     "example from problem",
			nums:     []int{1, 3, 5, 2, 4, 8, 2, 2},
			expected: 1,
		},
		{
			name:     "two elements - min",
			nums:     []int{5, 3},
			expected: 3,
		},
		{
			name:     "four elements",
			nums:     []int{10, 5, 8, 2},
			expected: 5,
		},
		{
			name:     "all same elements",
			nums:     []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "alternating min max",
			nums:     []int{1, 100, 2, 99, 3, 98, 4, 97},
			expected: 1,
		},
		{
			name:     "descending order",
			nums:     []int{8, 7, 6, 5, 4, 3, 2, 1},
			expected: 3,
		},
		{
			name:     "ascending order",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: 1,
		},
		{
			name:     "sixteen elements",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			expected: 1,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -1, -10, -2, -3, -8, -6, -4},
			expected: -5,
		},
		{
			name:     "two elements max wins",
			nums:     []int{2, 5},
			expected: 2,
		},
		{
			name:     "eight elements with duplicates",
			nums:     []int{3, 3, 3, 3, 3, 3, 3, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMaxGame(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInsertAtHead(t *testing.T) {
	tests := []struct {
		name string
		list []int
		val  int
		want []int
	}{
		{
			name: "insert into non-empty list",
			list: []int{1, 2, 3},
			val:  0,
			want: []int{0, 1, 2, 3},
		},
		{
			name: "insert into empty list",
			list: []int{},
			val:  0,
			want: []int{0},
		},
		{
			name: "insert into single element list",
			list: []int{5},
			val:  10,
			want: []int{10, 5},
		},
		{
			name: "insert negative number",
			list: []int{-1, -2},
			val:  -3,
			want: []int{-3, -1, -2},
		},
		{
			name: "insert duplicate value",
			list: []int{1, 1, 1},
			val:  1,
			want: []int{1, 1, 1, 1},
		},
		{
			name: "insert large number",
			list: []int{1000000},
			val:  999999,
			want: []int{999999, 1000000},
		},
		{
			name: "insert zero into positive list",
			list: []int{1, 2, 3},
			val:  0,
			want: []int{0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.list)
			resList := insertAtHead(list1, tt.val)

			result := listToSlice(resList)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestInsertAtTail(t *testing.T) {
	tests := []struct {
		name string
		list []int
		val  int
		want []int
	}{
		{
			name: "insert into non-empty list",
			list: []int{1, 2, 3},
			val:  0,
			want: []int{1, 2, 3, 0},
		},
		{
			name: "insert into empty list",
			list: []int{},
			val:  0,
			want: []int{0},
		},
		{
			name: "insert into single element list",
			list: []int{5},
			val:  10,
			want: []int{5, 10},
		},
		{
			name: "insert negative number",
			list: []int{-1, -2},
			val:  -3,
			want: []int{-1, -2, -3},
		},
		{
			name: "insert duplicate value",
			list: []int{1, 1, 1},
			val:  1,
			want: []int{1, 1, 1, 1},
		},
		{
			name: "insert large number",
			list: []int{1000000},
			val:  999999,
			want: []int{1000000, 999999},
		},
		{
			name: "insert zero into list",
			list: []int{1, 2, 3},
			val:  0,
			want: []int{1, 2, 3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.list)
			resList := insertAtTail(list1, tt.val)

			result := listToSlice(resList)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestInsertAtIndex(t *testing.T) {
	tests := []struct {
		name  string
		list  []int
		val   int
		index int
		want  []int
	}{
		{
			name:  "insert into middle of list",
			list:  []int{1, 2, 3},
			val:   99,
			index: 1,
			want:  []int{1, 99, 2, 3},
		},
		{
			name:  "insert at index 0 (head)",
			list:  []int{1, 2, 3},
			val:   0,
			index: 0,
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "insert at end of list",
			list:  []int{1, 2, 3},
			val:   4,
			index: 3,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "insert into single element list",
			list:  []int{5},
			val:   10,
			index: 1,
			want:  []int{5, 10},
		},
		{
			name:  "insert into empty list at index 0",
			list:  []int{},
			val:   42,
			index: 0,
			want:  []int{42},
		},
		{
			name:  "insert with negative value",
			list:  []int{-1, -2, -3},
			val:   -99,
			index: 2,
			want:  []int{-1, -2, -99, -3},
		},
		{
			name:  "insert duplicate value",
			list:  []int{1, 2, 3},
			val:   2,
			index: 2,
			want:  []int{1, 2, 2, 3},
		},
		{
			name:  "test 6",
			list:  []int{1, 2, 3},
			val:   2,
			index: 5,
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.list)
			resList := insertAtIndex(list1, tt.val, tt.index)

			result := listToSlice(resList)
			assert.Equal(t, tt.want, result)
		})
	}
}
