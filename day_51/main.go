package main

func maxArea(height []int) int {
	maxA := 0
	i, j := 0, len(height)-1

	for i < j {
		curA := 0
		if height[i] < height[j] {
			curA = height[i] * (j - i)
			i++
		} else {
			curA = height[j] * (j - i)
			j--
		}
		maxA = max(maxA, curA)
	}
	return maxA
}
