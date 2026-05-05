package main

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
