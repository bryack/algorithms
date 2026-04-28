package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums)/2)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}

		m[n] = struct{}{}
	}
	return false
}

func subarraySum(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	prefixSum := 0
	count := 0
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		if _, ok := m[prefixSum-k]; ok {
			count += m[prefixSum-k]
		}
		m[prefixSum]++
	}
	return count
}
