package main

func rotate(nums []int, k int) {
	k = k % len(nums)

	swap := func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	swap(0, len(nums)-1)
	swap(0, k-1)
	swap(k, len(nums)-1)
}

func maxArea(height []int) int {
	maxArea := 0

	i, j := 0, len(height)-1

	for i < j {
		if height[i] < height[j] {
			cArea := height[i] * (j - i)
			maxArea = max(maxArea, cArea)
			i++
		} else {
			cArea := height[j] * (j - i)
			maxArea = max(maxArea, cArea)
			j--
		}
	}
	return maxArea
}
