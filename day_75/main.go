package main

type TwoSum struct {
	data map[int]int
}

func Constructor() TwoSum {
	return TwoSum{data: make(map[int]int)}
}

func (t *TwoSum) Add(number int) {
	t.data[number]++
}

func (t *TwoSum) Find(value int) bool {
	for key, val := range t.data {
		c := value - key

		if c == key {
			if val > 1 {
				return true
			}
		} else if _, ok := t.data[c]; ok {
			return true
		}
	}
	return false
}

// считаем, что есть хотя бы 1 нода в каждом листе
// 1-1-1, 2-2-2 --> 3-3-3
// 1, 2-3-4 --> 4-3-3
// 5-4-3, 0-7-2 --> 6-1-5

type node struct {
	val  int
	next *node
}

func AddNumbers(l1, l2 *node) *node {
	carry := 0
	first, second := l1, l2
	dummy := &node{}
	cur := dummy

	for first != nil || second != nil || carry > 0 {
		sum := carry

		if first != nil {
			sum += first.val
			first = first.next
		}
		if second != nil {
			sum += second.val
			second = second.next
		}
		carry = sum / 10
		cur.next = &node{val: sum % 10}
		cur = cur.next
	}
	return dummy.next
}
