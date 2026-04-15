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
