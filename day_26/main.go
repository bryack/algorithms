package main

import (
	"math/rand"
	"unicode"
)

func repeatedCharacter(s string) string {
	var arr [26]bool

	for _, char := range s {
		if arr[char-'a'] {
			return string(char)
		}
		arr[char-'a'] = true
	}
	return ""
}

type RandomizedSet struct {
	nums []int
	m    map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		m:    map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}

	idx := this.m[val]
	this.nums[idx] = this.nums[len(this.nums)-1]
	this.m[this.nums[idx]] = idx

	this.nums = this.nums[:len(this.nums)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if !unicode.IsDigit(rune(s[left])) && !unicode.IsLetter(rune(s[left])) {
			left++
			continue
		}

		if !unicode.IsDigit(rune(s[right])) && !unicode.IsLetter(rune(s[right])) {
			right--
			continue
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}
	return true
}
