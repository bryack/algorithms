package main

import (
	"fmt"
	"io"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func findNearestNeighbours(r io.Reader, w io.Writer) {
	var size int
	fmt.Fscan(r, &size)

	prev := 0
	fmt.Fscan(r, &prev)
	minDiff := math.MaxInt
	x, y := 0, 0
	current := 0

	for i := 1; i < size; i++ {
		fmt.Fscan(r, &current)
		diff := current - prev
		if diff < minDiff {
			minDiff = diff
			x, y = prev, current
		}
		prev = current
	}

	fmt.Fprintf(w, "%d %d", x, y)
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	current := dummy

	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}

	head = dummy.Next

	return head
}
