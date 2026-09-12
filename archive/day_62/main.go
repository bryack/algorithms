package main

import (
	"fmt"
	"slices"
)

func longestPalindrome(s string) string {
	maxL, maxR := 0, 0
	n := len(s)

	for center := 0; center < 2*n-1; center++ {
		left := center / 2
		right := left + center%2

		for left >= 0 && right < n && s[left] == s[right] {
			if maxR-maxL < right-left {
				maxL, maxR = left, right
			}
			left--
			right++
		}
	}
	return s[maxL : maxR+1]
}

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	first := make([]int, n)
	second := make([]int, n)

	for i := range intervals {
		first[i] = intervals[i][0]
		second[i] = intervals[i][1]
	}
	slices.Sort(first)
	slices.Sort(second)
	fmt.Println(first)
	fmt.Println(second)

	res := 0
	end := 0
	for start := 0; start < n; start++ {
		fmt.Println("first[start]", first[start])
		fmt.Println("second[end]", second[end])
		fmt.Println("=======================")
		if first[start] < second[end] {
			res++
		} else {
			end++
		}
	}
	return res
}

type Node struct {
	key, value, freq int
	prev, next       *Node
}

type LFUCache struct {
	cap, minFreq int
	keyMap       map[int]*Node
	freqMap      map[int]*List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		keyMap:  make(map[int]*Node),
		freqMap: make(map[int]*List),
	}
}

type List struct {
	head, tail *Node
	size       int
}

func newList() *List {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return &List{
		head: head,
		tail: tail,
	}
}

func (l *List) insertFront(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
	l.size++
}

func (l *List) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (c *LFUCache) increment(node *Node) {
	freq := node.freq

	c.freqMap[freq].remove(node)

	if freq == c.minFreq && c.freqMap[freq].size == 0 {
		c.minFreq++
	}

	node.freq++
	if c.freqMap[node.freq] == nil {
		c.freqMap[node.freq] = newList()
	}
	c.freqMap[node.freq].insertFront(node)
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.keyMap[key]
	if !ok {
		return -1
	}
	c.increment(node)
	return node.value
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap == 0 {
		return
	}

	if node, ok := c.keyMap[key]; ok {
		node.value = value
		c.increment(node)
		return
	}

	if len(c.keyMap) >= c.cap {
		list := c.freqMap[c.minFreq]
		lfu := list.tail.prev
		list.remove(lfu)
		delete(c.keyMap, lfu.key)
	}

	node := &Node{key: key, value: value, freq: 1}
	c.keyMap[key] = node
	if c.freqMap[1] == nil {
		c.freqMap[1] = newList()
	}
	c.freqMap[1].insertFront(node)
	c.minFreq = 1
}
