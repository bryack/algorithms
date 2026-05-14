package main

import (
	"slices"
	"unicode"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		for i < len(s) && !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
		}
		for j >= 0 && !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
		}

		if i < len(s) && j >= 0 && unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
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

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[j] + nums[k]
			if sum+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum+nums[i] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
