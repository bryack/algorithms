package main

import "slices"

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	var arr [26]int

	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i]-'a']++
	}

	for i := 0; i < len(magazine); i++ {
		if arr[magazine[i]-'a'] > 0 {
			arr[magazine[i]-'a']--
		}
	}
	return arr == [26]int{}
}

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}
	if k == 0 || len(nums) == 1 {
		return
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func maxArea(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1

	for i < j {
		if height[i] <= height[j] {
			maxArea = max(maxArea, (j-i)*height[i])
			i++
		} else {
			maxArea = max(maxArea, (j-i)*height[j])
			j--
		}
	}
	return maxArea
}

func numRescueBoats(people []int, limit int) int {
	boats := 0

	slices.Sort(people)
	i, j := 0, len(people)-1
	for i <= j {
		sum := people[i] + people[j]
		if sum <= limit {
			boats++
			i++
			j--
		} else {
			boats++
			j--
		}
	}
	return boats
}
