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

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	idx := len(nums) - 1
	i, j := 0, len(nums)-1

	for i <= j {
		squareI := nums[i] * nums[i]
		squareJ := nums[j] * nums[j]
		if squareI > squareJ {
			res[idx] = squareI
			i++
		} else {
			res[idx] = squareJ
			j--
		}
		idx--
	}
	return res
}

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	j := 0

	for i := 0; i < len(t); i++ {
		if j < len(s) && s[j] == t[i] {
			j++
		}
	}
	return len(s) == j
}
