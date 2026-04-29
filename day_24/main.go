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

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	maxLen := 0

	for key := range m {
		if _, ok := m[key-1]; !ok {
			currentLen := 1
			currentNum := key
			for {
				if _, ok := m[currentNum+1]; ok {
					currentLen++
					currentNum++
				} else {
					break
				}
			}
			maxLen = max(currentLen, maxLen)
		}
	}
	return maxLen
}
