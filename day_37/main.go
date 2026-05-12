package main

func largestAltitude(gain []int) int {
	sum, count := 0, 0

	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		count = max(count, sum)
	}
	return count
}

func pivotIndex(nums []int) int {
	prefixSum := 0
	for _, n := range nums {
		prefixSum += n
	}

	tempSum := 0
	for i, n := range nums {
		prefixSum = prefixSum - n
		if prefixSum == tempSum {
			return i
		}
		tempSum += n
	}
	return -1
}
