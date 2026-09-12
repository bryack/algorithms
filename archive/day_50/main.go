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
