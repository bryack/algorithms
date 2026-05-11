package main

func largestAltitude(gain []int) int {
	counter, inter := 0, 0

	for i := 0; i < len(gain); i++ {
		inter += gain[i]
		counter = max(counter, inter)
	}
	return counter
}

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1

	for i <= j {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
