package main

import (
	"math"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	count := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[count] = nums2[j]
			j--
		} else {
			nums1[count] = nums1[i]
			i--
		}
		count--
	}
	copy(nums1, nums2[:j+1])
}

func maxProduct(nums []int) int {
	if len(nums) < 3 {
		return (nums[0] - 1) * (nums[1] - 1)
	}
	max1 := max(nums[0], nums[1])
	max2 := min(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	return (max1 - 1) * (max2 - 1)
}

func thirdMax(nums []int) int {
	max1 := math.MinInt
	max2 := math.MinInt
	max3 := math.MinInt
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == max1 || nums[i] == max2 {
			continue
		}
		if nums[i] > max1 {
			max3 = max2
			max2 = max1
			max1 = nums[i]
			count++
		} else if nums[i] > max2 {
			max3 = max2
			max2 = nums[i]
			count++
		} else if nums[i] > max3 {
			max3 = nums[i]
			count++
		}
	}
	if count < 3 {
		max3 = max1
	}
	return max3
}

func maximumProduct(nums []int) int {
	max1 := math.MinInt
	max2 := math.MinInt
	max3 := math.MinInt
	min1 := math.MaxInt
	min2 := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] < min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
		if nums[i] > max1 {
			max3 = max2
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max3 = max2
			max2 = nums[i]
		} else if nums[i] > max3 {
			max3 = nums[i]
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		if max < count {
			max = count
		}
	}
	return max
}

func shuffle(nums []int, n int) []int {
	s := make([]int, len(nums))
	for i := 0; i < n; i++ {
		s[2*i] = nums[i]
		s[2*i+1] = nums[i+n]
	}

	return s
}
