package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
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
	list := &ListNode{}
	list.Val = s[0]
	current := list

	for i := 1; i < len(s); i++ {
		node := &ListNode{Val: s[i]}
		current.Next = node
		current = current.Next
	}
	return list
}

func divideArray(nums []int) bool {
	var arr [501]int

	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	for _, v := range arr {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	i := 0
	for i < len(nums)/2 {
		if i%2 == 0 {
			nums[i] = min(nums[2*i], nums[2*i+1])
		} else {
			nums[i] = max(nums[2*i], nums[2*i+1])
		}

		if i == len(nums)/2-1 {
			nums = nums[:len(nums)/2]
			i = 0
		} else {
			i++
		}
	}
	return nums[0]
}

func insertAtHead(head *ListNode, val int) *ListNode {
	return &ListNode{Val: val, Next: head}
}

func insertAtTail(head *ListNode, val int) *ListNode {
	if head == nil {
		return &ListNode{Val: val, Next: nil}
	}

	tail := head
	for tail != nil {
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}
	tail.Next = &ListNode{Val: val, Next: nil}
	return head
}

func insertAtIndex(head *ListNode, val int, index int) *ListNode {
	if index < 0 {
		return head
	}
	if index == 0 {
		return &ListNode{Val: val, Next: head}
	}

	current := head
	for i := 0; i != index-1; i++ {
		if current == nil {
			return head
		}
		current = current.Next
	}

	current.Next = &ListNode{Val: val, Next: current.Next}

	return head
}
