package main

import "slices"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	s := make([]int, 0, len(nums1))

	for _, n := range nums1 {
		m[n]++
	}

	for _, n := range nums2 {
		if m[n] > 0 {
			s = append(s, n)
			m[n]--
		}
	}
	return s
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0, len(nums)/2)
	if len(nums) < 4 {
		return res
	}
	slices.Sort(nums)
	a := 0
	for a < len(nums)-3 {
		if a > 0 && nums[a-1] == nums[a] {
			a++
			continue
		}
		if nums[a]+nums[a+1]+nums[a+2]+nums[a+3] > target {
			break
		}
		if nums[a]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			a++
			continue
		}
		b := a + 1
		for b < len(nums)-2 {
			if b > a+1 && nums[b-1] == nums[b] {
				b++
				continue
			}

			if nums[a]+nums[b]+nums[b+1]+nums[b+2] > target {
				break
			}
			if nums[a]+nums[b]+nums[len(nums)-2]+nums[len(nums)-1] < target {
				b++
				continue
			}

			c, d := b+1, len(nums)-1

			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
					for c < d && nums[c-1] == nums[c] {
						c++
					}
					for c < d && nums[d+1] == nums[d] {
						d--
					}
				} else if sum < target {
					c++
				} else if sum > target {
					d--
				}
			}
			b++
		}
		a++
	}
	return res
}
