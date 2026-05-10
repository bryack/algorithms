package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push_Pop(t *testing.T) {
	tests := []struct {
		name   string
		list   *ListNode
		input  int
		want   int
		wantOk bool
	}{
		{
			name:   "not empty stack",
			list:   &ListNode{Val: 2},
			input:  10,
			want:   10,
			wantOk: true,
		},
		{
			name:   "empty stack",
			list:   nil,
			input:  10,
			want:   10,
			wantOk: true,
		},
		{
			name:   "push to existing chain returns pushed value",
			list:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			input:  3,
			want:   3,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := Stack{head: tt.list}
			stack.Push(tt.input)
			v, b := stack.Pop()
			assert.Equal(t, tt.want, v)
			assert.Equal(t, tt.wantOk, b)
		})
	}
}

func TestStack_Top(t *testing.T) {
	tests := []struct {
		name   string
		list   *ListNode
		want   int
		wantOk bool
	}{
		{
			name:   "not empty stack",
			list:   &ListNode{Val: 2},
			want:   2,
			wantOk: true,
		},
		{
			name:   "empty stack",
			list:   nil,
			want:   0,
			wantOk: false,
		},
		{
			name:   "push to existing chain returns pushed value",
			list:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want:   1,
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := Stack{head: tt.list}
			v, b := stack.Top()
			assert.Equal(t, tt.want, v)
			assert.Equal(t, tt.wantOk, b)
		})
	}
}

func TestSliceStack_Push(t *testing.T) {
	tests := []struct {
		name  string
		input *SliceStack
		val   int
		want  []int
	}{
		{
			name:  "not empty stack",
			input: &SliceStack{items: []int{1, 2, 3}},
			val:   5,
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "empty stack",
			input: &SliceStack{},
			val:   5,
			want:  []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			s.Push(tt.val)
			assert.Equal(t, tt.want, s.items)
		})
	}
}

func TestSliceStack_Pop(t *testing.T) {
	tests := []struct {
		name      string
		input     *SliceStack
		want      int
		wantOk    bool
		wantSlice []int
	}{
		{
			name:      "not empty stack",
			input:     &SliceStack{items: []int{1, 2, 3}},
			want:      3,
			wantOk:    true,
			wantSlice: []int{1, 2},
		},
		{
			name:      "empty stack",
			input:     &SliceStack{},
			want:      0,
			wantOk:    false,
			wantSlice: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.input
			v, ok := s.Pop()
			assert.Equal(t, tt.want, v)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantSlice, s.items)
		})
	}
}
