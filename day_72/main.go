package main

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if n < m {
		return -1
	}
	firstChar := needle[0]
	for i := 0; i < n+1-m; i++ {
		if firstChar == haystack[i] {
			if needle == haystack[i:i+m] {
				return i
			}
		}
	}
	return -1
}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		for j < n-1 {
			x := j + 1
			y := n - 1
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
				} else if sum < target {
					x++
				} else {
					y--
				}
			}
			j++
			for j < n-1 && nums[j] == nums[j-1] {
				j++
			}
		}
	}
	return res
}
