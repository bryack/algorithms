package main

import (
	"slices"
	"unicode"
)

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

func isSubsequence(sub string, s string) bool {
	j := 0

	for i := 0; i < len(s); i++ {
		if j < len(sub) && s[i] == sub[j] {
			j++
		}
	}
	return len(sub) == j
}

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		i = nextValid(s, i)
		j = nextValid(t, j)

		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 {
			return false
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func nextValid(s string, i int) int {
	count := 0
	for i >= 0 {
		if s[i] == '#' {
			count++
		} else if count > 0 {
			count--
		} else {
			return i
		}
		i--
	}
	return -1
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	idx := len(nums1) - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
		idx--
	}

	for j >= 0 {
		nums1[idx] = nums2[j]
		j--
		idx--
	}
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func mergeAlternately(word1 string, word2 string) string {
	buf := make([]byte, len(word1)+len(word2))
	idx := 0
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		buf[idx] = word1[i]
		i++
		idx++
		buf[idx] = word2[j]
		j++
		idx++
	}
	for i < len(word1) {
		buf[idx] = word1[i]
		i++
		idx++
	}
	for j < len(word2) {
		buf[idx] = word2[j]
		j++
		idx++
	}

	return string(buf)
}

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0, len(nums)/2)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		for j < len(nums)-1 {
			x, y := j+1, len(nums)-1
			for x < y {
				sum := nums[i] + nums[j] + nums[x] + nums[y]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[x], nums[y]})
					x++
					y--
					for x < y && nums[x] == nums[x-1] {
						x++
					}
					for x < y && nums[y] == nums[y+1] {
						y--
					}
				} else if sum > target {
					y--
				} else {
					x++
				}
			}
			j++
			for j < len(nums)-1 && nums[j] == nums[j-1] {
				j++
			}

		}
	}
	return res
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func numRescueBoats(people []int, limit int) int {
	res := 0
	slices.Sort(people)

	i, j := 0, len(people)-1
	for i <= j {
		sum := people[i] + people[j]
		if sum > limit {
			j--
		} else {
			i++
			j--
		}
		res++
	}
	return res
}

func trap(height []int) int {
	i, j := 0, len(height)-1
	maxRight := height[j]
	maxLeft := height[i]
	res := 0

	for i < j {
		if maxRight > height[j] {
			res += maxRight - height[j]
		} else {
			maxRight = height[j]
		}
		if maxLeft > height[i] {
			res += maxLeft - height[i]
		} else {
			maxLeft = height[i]
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func reverseOnlyLetters(s string) string {
	buf := []rune(s)

	i, j := 0, len(s)-1
	for i <= j {
		if !unicode.IsLetter(buf[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(buf[j]) {
			j--
			continue
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func longestPalindrome(s string) string {
	maxL, maxR := 0, 0

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left > maxR-maxL {
				maxL = left
				maxR = right
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[maxL : maxR+1]
}
