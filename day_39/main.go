package main

import "unicode"

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
