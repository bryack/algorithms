package main

import (
	"fmt"
	"slices"
)

func mergeAlternately(word1 string, word2 string) string {
	buf := make([]byte, len(word1)+len(word2))

	i, j := 0, 0
	idx := 0
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

func isSubsequence(sub string, s string) bool {
	j := 0

	for i := 0; i < len(s); i++ {
		if j < len(sub) && s[i] == sub[j] {
			j++
		}
	}
	return len(sub) == j
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

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0, len(nums)/2)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

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
		}
	}
	return res
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) == 1 {
		return
	}

	swap(nums)
	swap(nums[:k])
	swap(nums[k:len(nums)])

}

func swap(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

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

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	i, j := 0, len(people)-1
	res := 0

	for i <= j {
		if people[j]+people[i] > limit {
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
	res := 0
	i, j := 0, len(height)-1
	maxLeft, maxRight := 0, 0

	for i < j {
		if height[i] < maxLeft {
			res += maxLeft - height[i]
		}
		if height[j] < maxRight {
			res += maxRight - height[j]
		}
		maxLeft = max(maxLeft, height[i])
		maxRight = max(maxRight, height[j])
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}

func reverseOnlyLetters(s string) string {
	buf := []byte(s)

	i, j := 0, len(s)-1
	for i <= j {
		if !isLetter(buf[i]) {
			i++
			continue
		}
		if !isLetter(buf[j]) {
			j--
			continue
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func isLetter(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func longestPalindrome(s string) string {
	maxL, maxR := 0, 0

	expand := func(i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			if j-i > maxR-maxL {
				maxR, maxL = j, i
			}
			i--
			j++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i+1)
		expand(i, i)
	}
	return s[maxL : maxR+1]
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
			i--
		} else if count > 0 {
			count--
			i--
		} else {
			return i
		}
	}
	return -1
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0, len(nums))
	i, j := 0, 0

	for i < len(nums) {
		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			i++
		}

		if nums[j] != nums[i] {
			res = append(res, fmt.Sprintf("%d->%d", nums[j], nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[j]))
		}
		i++
		j = i
	}
	return res
}
