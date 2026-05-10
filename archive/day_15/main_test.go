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

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "even length palindrome",
			input: []int{1, 2, 2, 1},
			want:  true,
		},
		{
			name:  "even length not palindrome",
			input: []int{1, 2},
			want:  false,
		},
		{
			name:  "empty list",
			input: []int{},
			want:  true,
		},
		{
			name:  "single element",
			input: []int{1},
			want:  true,
		},
		{
			name:  "odd length palindrome",
			input: []int{1, 2, 1},
			want:  true,
		},
		{
			name:  "all same elements",
			input: []int{5, 5, 5},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := sliceToList(tt.input)
			result := isPalindrome(list)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name         string
		firstList    []int
		secondList   []int
		intersectVal int
		want         int
	}{
		{
			name:         "intersection in middle",
			firstList:    []int{1, 9, 1, 2, 4},
			secondList:   []int{3, 2, 4},
			intersectVal: 2,
			want:         2,
		},
		{
			name:         "intersection at tail",
			firstList:    []int{1, 2, 3},
			secondList:   []int{4, 5, 3},
			intersectVal: 3,
			want:         3,
		},
		{
			name:         "intersection at head",
			firstList:    []int{8, 4, 5},
			secondList:   []int{1, 8, 4, 5},
			intersectVal: 8,
			want:         8,
		},
		{
			name:         "no intersection",
			firstList:    []int{1, 2, 3},
			secondList:   []int{4, 5, 6},
			intersectVal: -1,
			want:         0, // res will be nil, check separately
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1, list2 := createIntersectingLists(tt.firstList, tt.secondList, tt.intersectVal)
			res := getIntersectionNode(list1, list2)

			if tt.intersectVal == -1 {
				assert.Nil(t, res)
			} else {
				assert.NotNil(t, res)
				assert.Equal(t, tt.want, res.Val)
			}
		})
	}
}

func createIntersectingLists(listA, listB []int, intersectVal int) (*ListNode, *ListNode) {
	if len(listA) == 0 && len(listB) == 0 {
		return nil, nil
	}
	common := make([]int, 0, len(listA)+len(listB))
	for i := 0; i < len(listA); i++ {
		if listA[i] == intersectVal {
			common = append(common, listA[i:]...)
			break
		}
	}

	var commonList *ListNode
	if len(common) != 0 {
		commonList = &ListNode{Val: common[0]}
		current := commonList
		for i := 1; i < len(common); i++ {
			node := &ListNode{Val: common[i]}
			current.Next = node
			current = current.Next
		}
	}

	var headA *ListNode
	if listA[0] == intersectVal {
		headA = commonList
	} else {
		headA = &ListNode{Val: listA[0]}
		current1 := headA
		for i := 1; i < len(listA); i++ {
			if listA[i] == intersectVal {
				current1.Next = commonList
				break
			}
			node := &ListNode{Val: listA[i]}
			current1.Next = node
			current1 = current1.Next
		}
	}

	var headB *ListNode
	if listB[0] == intersectVal {
		headB = commonList
	} else {
		headB = &ListNode{Val: listB[0]}
		current2 := headB
		for i := 1; i < len(listB); i++ {
			if listB[i] == intersectVal {
				current2.Next = commonList
				break
			}
			node := &ListNode{Val: listB[i]}
			current2.Next = node
			current2 = current2.Next
		}
	}

	return headA, headB
}

func TestCreateIntersectingLists_HeadReference(t *testing.T) {
	// Этот тест проверяет, что при intersectVal на нулевом индексе
	// head'ы списков указывают на один и тот же узел в памяти
	res1, res2 := createIntersectingLists([]int{8, 4, 5}, []int{8, 4, 5}, 8)

	// Проверяем, что значения правильные
	assert.Equal(t, []int{8, 4, 5}, listToSlice(res1))
	assert.Equal(t, []int{8, 4, 5}, listToSlice(res2))

	// Проверяем, что это действительно один и тот же узел в памяти (по ссылке)
	// Этот тест должен упасть с текущей реализацией, показывая баг
	if res1 != res2 {
		t.Errorf("Head A и Head B должны быть одним и тем же узлом в памяти, но они разные: res1=%p, res2=%p", res1, res2)
	}
}

func TestCreateIntersectingLists(t *testing.T) {
	tests := []struct {
		name           string
		firstList      []int
		secondList     []int
		intersectVal   int
		wantFirstList  []int
		wantSecondList []int
	}{
		{
			name:           "intersection in middle",
			firstList:      []int{1, 9, 1, 2, 4},
			secondList:     []int{3, 2, 4},
			intersectVal:   2,
			wantFirstList:  []int{1, 9, 1, 2, 4},
			wantSecondList: []int{3, 2, 4},
		},
		{
			name:           "no intersection",
			firstList:      []int{1, 2, 3},
			secondList:     []int{4, 5, 6},
			intersectVal:   -1,
			wantFirstList:  []int{1, 2, 3},
			wantSecondList: []int{4, 5, 6},
		},
		{
			name:           "intersection at head of both",
			firstList:      []int{8, 4, 5},
			secondList:     []int{8, 4, 5},
			intersectVal:   8,
			wantFirstList:  []int{8, 4, 5},
			wantSecondList: []int{8, 4, 5},
		},
		{
			name:           "intersection at tail only",
			firstList:      []int{1, 2, 3},
			secondList:     []int{4, 5, 3},
			intersectVal:   3,
			wantFirstList:  []int{1, 2, 3},
			wantSecondList: []int{4, 5, 3},
		},
		{
			name:           "second list is entirely intersection",
			firstList:      []int{1, 2, 3},
			secondList:     []int{3},
			intersectVal:   3,
			wantFirstList:  []int{1, 2, 3},
			wantSecondList: []int{3},
		},
		{
			name:           "intersection at index 0 for listA only",
			firstList:      []int{8, 4, 5},
			secondList:     []int{1, 8, 4, 5},
			intersectVal:   8,
			wantFirstList:  []int{8, 4, 5},
			wantSecondList: []int{1, 8, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res1, res2 := createIntersectingLists(tt.firstList, tt.secondList, tt.intersectVal)
			result1 := listToSlice(res1)
			result2 := listToSlice(res2)
			assert.Equal(t, tt.wantFirstList, result1)
			assert.Equal(t, tt.wantSecondList, result2)
		})
	}
}
