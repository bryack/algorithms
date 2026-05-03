package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums)/2)

	for i, n := range nums {
		if v, ok := m[n]; ok {
			if i-v <= k {
				return true
			}
			m[n] = i
		} else {
			m[n] = i
		}
	}
	return false
}
