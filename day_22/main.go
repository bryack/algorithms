package main

import (
	"cmp"
	"slices"
)

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr1 [256]int
	var arr2 [256]int

	for i := 0; i < len(s); i++ {
		if arr1[s[i]] != arr2[t[i]] {
			return false
		}
		arr1[s[i]] = i + 1
		arr2[t[i]] = i + 1
	}
	return true
}

func frequencySort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	var arr [201]int
	for i := 0; i < len(nums); i++ {
		arr[nums[i]+100]++
	}

	bucket := make([][]int, len(nums)+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]] = append(bucket[arr[i]], i-100)
	}

	res := make([]int, 0, len(nums))
	for i := 0; i < len(bucket); i++ {
		if len(bucket[i]) == 0 {
			continue
		}

		slices.SortFunc(bucket[i], func(a, b int) int {
			return cmp.Compare(b, a)
		})

		for _, b := range bucket[i] {
			for j := 0; j < i; j++ {
				res = append(res, b)
			}
		}
	}
	return res
}

func topKFrequent(words []string, k int) []string {
	if len(words) == 1 {
		return words
	}

	m := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	bucket := make([][]string, len(words)+1)
	for key, value := range m {
		bucket[value] = append(bucket[value], key)
	}

	res := make([]string, 0, k)

	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		slices.Sort(bucket[i])
		for _, w := range bucket[i] {
			if len(res) >= k {
				return res
			}
			res = append(res, w)
		}
	}

	return res
}

func productExceptSelf(nums []int) []int {
	p := 1
	s := 1
	res := make([]int, 0, len(nums))

	for i := 1; i < len(nums)+1; i++ {
		res = append(res, p)
		p *= nums[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= s
		s *= nums[i]
	}
	return res
}

type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	prefixSums := make([]int, len(nums)+1)
	prefixSums[0] = 0
	for i := 1; i <= len(nums); i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i-1]
	}
	return NumArray{prefixSums: prefixSums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSums[right+1] - this.prefixSums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
