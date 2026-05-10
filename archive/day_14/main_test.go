package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "three elements",
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
		{
			name:  "empty list (nil)",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "two elements",
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "all duplicates",
			input: []int{5, 5, 5},
			want:  []int{5, 5, 5},
		},
		{
			name:  "already descending",
			input: []int{9, 5, 2, 1},
			want:  []int{1, 2, 5, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sliceToList(tt.input)
			res := reverseList(list)
			result := listToSlice(res)
			assert.Equal(t, tt.want, result)
		})
	}
}

func listToSlice(list *ListNode) []int {
	s := []int{}
	current := list

	for current != nil {
		s = append(s, current.Val)
		current = current.Next
	}
	return s
}

func sliceToList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	list := &ListNode{Val: s[0]}
	current := list

	for i := 1; i < len(s); i++ {
		node := &ListNode{Val: s[i]}
		current.Next = node
		current = current.Next
	}
	return list
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		cycleToIndex int
		want         bool
	}{
		{
			name:         "cycle at middle",
			input:        []int{1, 2, 3},
			cycleToIndex: 1,
			want:         true,
		},
		{
			name:         "no cycle (nil tail)",
			input:        []int{1, 2, 3},
			cycleToIndex: -1,
			want:         false,
		},
		{
			name:         "empty list",
			input:        []int{},
			cycleToIndex: -1,
			want:         false,
		},
		{
			name:         "cycle at head (last points to first)",
			input:        []int{1, 2, 3},
			cycleToIndex: 0,
			want:         true,
		},
		{
			name:         "single element self-cycle",
			input:        []int{5},
			cycleToIndex: 0,
			want:         true,
		},
		{
			name:         "single element no cycle",
			input:        []int{5},
			cycleToIndex: -1,
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sliceToListWithCycle(tt.input, tt.cycleToIndex)

			result := hasCycle(list)
			assert.Equal(t, tt.want, result)
		})
	}
}

func sliceToListWithCycle(values []int, cycleToIndex int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	var cycleNode *ListNode
	list := &ListNode{Val: values[0]}
	if cycleToIndex == 0 {
		cycleNode = list
	}
	current := list

	for i := 1; i < len(values); i++ {
		node := &ListNode{Val: values[i]}
		if i == cycleToIndex {
			cycleNode = node
		}
		current.Next = node

		current = current.Next
	}
	current.Next = cycleNode
	return list
}

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "odd length",
			input: []int{1, 2, 3, 4, 5},
			want:  3,
		},
		{
			name:  "even length — second middle",
			input: []int{1, 2, 3, 4},
			want:  3,
		},
		{
			name:  "single element",
			input: []int{42},
			want:  42,
		},
		{
			name:  "empty list",
			input: []int{},
			want:  0,
		},
		{
			name:  "two elements",
			input: []int{1, 2},
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sliceToList(tt.input)
			res := middleNode(list)
			if res == nil {
				assert.Equal(t, tt.want, 0)
			} else {
				result := listToSlice(res)
				assert.Equal(t, tt.want, result[0])
			}
		})
	}
}

func TestFindNearestNeighbours(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "basic example",
			input: `4
1 3 4 11`,
			want: "3 4",
		},
		{
			name: "two elements",
			input: `2
		10 20`,
			want: "10 20",
		},
		{
			name: "all equal elements",
			input: `5
		7 7 7 7 7`,
			want: "7 7",
		},
		{
			name: "uniform spacing",
			input: `6
		2 4 6 8 10 12`,
			want: "2 4",
		},
		{
			name: "min diff at the end",
			input: `5
		1 10 20 25 26`,
			want: "25 26",
		},
		{
			name: "large values",
			input: `3
		100000 200000 200001`,
			want: "200000 200001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			w := bytes.Buffer{}
			findNearestNeighbours(r, &w)
			assert.Equal(t, tt.want, w.String())

		})
	}
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		val   int
		want  []int
	}{
		{
			name:  "remove from middle",
			input: []int{1, 2, 6, 3, 4, 5, 6},
			val:   6,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "empty list",
			input: []int{},
			val:   1,
			want:  []int{},
		},
		{
			name:  "all elements match",
			input: []int{7, 7, 7, 7},
			val:   7,
			want:  []int{},
		},
		{
			name:  "remove from head",
			input: []int{1, 1, 1, 2, 3, 4, 5},
			val:   1,
			want:  []int{2, 3, 4, 5},
		},
		{
			name:  "single element matches",
			input: []int{6},
			val:   6,
			want:  []int{},
		},
		{
			name:  "consecutive duplicates at start",
			input: []int{6, 6, 1, 2, 3},
			val:   6,
			want:  []int{1, 2, 3},
		},
		{
			name:  "no elements match",
			input: []int{1, 2, 3, 4, 5},
			val:   9,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "consecutive in middle",
			input: []int{1, 2, 6, 6, 6, 3},
			val:   6,
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sliceToList(tt.input)
			res := removeElements(list, tt.val)
			result := listToSlice(res)
			assert.Equal(t, tt.want, result)
		})
	}
}
