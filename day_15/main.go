package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}

	return head
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var reverse *ListNode
	current := slow
	for current != nil {
		next := current.Next
		current.Next = reverse
		reverse = current
		current = next
	}

	c := reverse
	d := head
	for c != nil {
		if d.Val != c.Val {
			return false
		}
		c = c.Next
		d = d.Next
	}

	return true
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
