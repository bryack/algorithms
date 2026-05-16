package main

func maxArea(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1

	for i < j {
		if height[i] < height[j] {
			currentArea := height[i] * (j - i)
			maxArea = max(maxArea, currentArea)
			i++
		} else {
			currentArea := height[j] * (j - i)
			maxArea = max(maxArea, currentArea)
			j--
		}
	}
	return maxArea
}
