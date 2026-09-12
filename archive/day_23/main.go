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

func runningSum(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	prefixSum := 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		nums[i] = prefixSum
	}
	return nums
}

func largestAltitude(gain []int) int {
	prefixSum := 0
	count := 0
	for i := 0; i < len(gain); i++ {
		prefixSum += gain[i]
		count = max(count, prefixSum)
	}
	return count
}

func pivotIndex(nums []int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum := totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
