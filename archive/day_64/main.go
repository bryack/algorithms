package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)

	for end := 0; end < n; end++ {
		start := end

		for end < n-1 && nums[end]+1 == nums[end+1] {
			end++
		}

		if nums[end] == nums[start] {
			res = append(res, fmt.Sprintf("%d", nums[end]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		}
	}
	return res
}

func findMaxConsecutiveOnes(nums []int) int {
	begin := 0
	wState := 0
	res := 0

	for end := range nums {
		if nums[end] == 0 {
			wState++
		}

		for wState > 1 {
			if nums[begin] == 0 {
				wState--
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	begin := 0
	res := 0
	m := map[byte]int{}

	for end := range s {
		m[s[end]]++

		for len(m) > 2 {
			m[s[begin]]--
			if m[s[begin]] == 0 {
				delete(m, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)-k

	for lo < hi {
		mid := lo + (hi-lo)/2

		if x-arr[mid] > arr[mid+k]-x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[lo : lo+k]
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	begin := 0
	res := 0
	m := map[byte]int{}

	for end := range s {
		m[s[end]]++

		for len(m) > k {
			m[s[begin]]--
			if m[s[begin]] == 0 {
				delete(m, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 {
		return false
	}

	buckets := make(map[int]int)
	bucketSize := valueDiff + 1

	for i, num := range nums {
		bucketID := getBucketID(num, bucketSize)

		if _, ok := buckets[bucketID]; ok {
			return true
		}

		if val, ok := buckets[bucketID+1]; ok && abs(num-val) < bucketSize {
			return true
		}

		if val, ok := buckets[bucketID-1]; ok && abs(num-val) < bucketSize {
			return true
		}

		buckets[bucketID] = num

		if i >= indexDiff {
			delete(buckets, getBucketID(nums[i-indexDiff], bucketSize))
		}
	}
	return false
}

func getBucketID(num, bucketSize int) int {
	if num >= 0 {
		return num / bucketSize
	}
	return (num+1)/bucketSize - 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
