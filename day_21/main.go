package main

import (
	"math"
)

func isValid(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	return arr == [26]int{}
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	if len(ransomNote) == 1 && len(magazine) == 1 && ransomNote != magazine {
		return false
	}

	var arr [26]int
	for i := 0; i < len(magazine); i++ {
		arr[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i]-'a']--
		if arr[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func maxNumberOfBalloons(text string) int {
	if len(text) < 7 {
		return 0
	}

	var arrText [26]int

	for i := 0; i < len(text); i++ {
		arrText[text[i]-'a']++
	}

	b := "balloon"
	var arrBallon [26]int
	for i := 0; i < len(b); i++ {
		arrBallon[b[i]-'a']++
	}

	count := math.MaxInt
	for i := 0; i < len(arrBallon); i++ {
		if arrBallon[i] == 0 {
			continue
		}
		c := arrText[i] / arrBallon[i]
		count = min(count, c)
	}
	return count
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, value := range m {
		bucket[value] = append(bucket[value], key)
	}

	res := make([]int, 0, k)
	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		for _, v := range bucket[i] {
			if len(res) == k {
				return res
			}
			res = append(res, v)
		}
	}

	return res
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
