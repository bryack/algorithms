package main

import (
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
