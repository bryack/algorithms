package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{}
}

func findMaxLength(nums []int) int {
	m := make(map[int]int, len(nums))
	m[0] = -1
	maxLen := 0
	currentLen := 0
	prefixSum := 0

	for i, num := range nums {
		if num == 0 {
			prefixSum--
		} else {
			prefixSum++
		}
		if v, ok := m[prefixSum]; ok {
			currentLen = i - v
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			m[prefixSum] = i
		}
	}
	return maxLen
}
