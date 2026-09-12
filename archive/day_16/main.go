package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	head *ListNode
}

func (s *Stack) Push(val int) {
	next := s.head
	s.head = &ListNode{Val: val, Next: next}
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	p := s.head.Val
	s.head = s.head.Next
	return p, true
}

func (s *Stack) Top() (int, bool) {
	if s.head == nil {
		return 0, false
	}
	return s.head.Val, true
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

type SliceStack struct {
	items []int
}

func (s *SliceStack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *SliceStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *SliceStack) Top() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *SliceStack) IsEmpty() bool {
	return len(s.items) == 0
}
