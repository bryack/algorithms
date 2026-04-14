package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "duplicates at start",
			input: []int{1, 1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "duplicates at end",
			input: []int{1, 1, 2, 3, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty list",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "all same elements",
			input: []int{5, 5, 5, 5},
			want:  []int{5},
		},
		{
			name:  "no duplicates",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sliceToList(tt.input)
			res := deleteDuplicates(list)
			result := listToSlice(res)
			assert.Equal(t, tt.want, result)
		})
	}
}

func listToSlice(list *ListNode) []int {
	current := list
	s := []int{}
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
