package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	swap := func(i, j int) {
		for i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	swap(0, len(nums)-1)
	swap(0, k-1)
	swap(k, len(nums)-1)
}
