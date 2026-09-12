package main

import (
	"fmt"
	"slices"
)

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

func findMaxConsecutiveOnes(nums []int) int {
	wState, begin := 0, 0
	res := 0

	for end := range nums {
		if nums[end] == 0 {
			wState++
		}

		for wState > 1 {
			if nums[begin] == 0 {
				wState--
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := make(map[byte]int)
	res, begin := 0, 0

	for end := range s {
		m[s[end]]++

		for len(m) > 2 {
			m[s[begin]]--
			if m[s[begin]] == 0 {
				delete(m, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func f(n int) {
	fmt.Printf("Начало f(%d)\n", n)

	if n == 0 {
		fmt.Println("return")
		return
	}

	f(n - 1)

	fmt.Printf("Конец f(%d)\n", n)
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := make(map[byte]int)
	begin, res := 0, 0

	for end := range s {
		m[s[end]]++

		for len(m) > k {
			m[s[begin]]--
			if m[s[begin]] == 0 {
				delete(m, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func fact(n int) int {
	fmt.Println("вошли:", n)

	if n <= 1 {
		fmt.Println("возвращаем:", 1)
		return 1
	}

	result := n * fact(n-1)

	fmt.Println("для", n, "получили", result)

	return result
}

func rcs3(x int) {
	fmt.Println("down: ", x)
	fmt.Println("up: ", x)
}

func rcs2(x int) {
	fmt.Println("down: ", x)
	rcs3(x - 1)
	fmt.Println("up: ", x)
}

func rcs1(x int) {
	fmt.Println("down: ", x)
	rcs2(x - 1)
	fmt.Println("up: ", x)
}

func rcs(x int) {
	fmt.Println("down: ", x)
	if x > 1 {
		rcs(x - 1)
	}
	fmt.Println("up: ", x)
}

func main() {
	fmt.Println(factorial(6))
}

func factorial(n int) int {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

func minMeetingRooms(intervals [][]int) int {
	start, end := make([]int, len(intervals)), make([]int, len(intervals))

	for i := range intervals {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}

	slices.Sort(start)
	slices.Sort(end)

	j := 0
	res := 0

	for i := range start {
		if start[i] < end[j] {
			res++
		} else {
			j++
		}
	}
	return res
}
