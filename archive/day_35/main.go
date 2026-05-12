package main

import (
	"cmp"
	"slices"
)

func frequencySort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	var arr [201]int
	for _, n := range nums {
		arr[n+100]++
	}

	bucket := make([][]int, len(nums)+1)
	for i, n := range arr {
		if arr[i] == 0 {
			continue
		}
		bucket[n] = append(bucket[n], i-100)
	}

	res := make([]int, 0, len(nums))
	for i, b := range bucket {
		slices.SortFunc(b, func(x, y int) int {
			return cmp.Compare(y, x)
		})
		for _, n := range b {
			for range i {
				res = append(res, n)
			}
		}
	}
	return res
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i <= j {
		sum := numbers[i] + numbers[j]
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
