package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var lastSeenS [256]int
	var lastSeenT [256]int

	for i := 0; i < len(s); i++ {
		if lastSeenS[s[i]] == lastSeenT[t[i]] {
			lastSeenS[s[i]] = i + 1
			lastSeenT[t[i]] = i + 1
		} else {
			return false
		}
	}
	return true
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 1 {
		return false
	}
	m := make(map[int]int, k)

	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			if i-v <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

func commonChars(words []string) []string {
	var arr [26]int

	for _, v := range words[0] {
		arr[v-'a']++
	}

	for i := 1; i < len(words); i++ {
		var tempArr [26]int
		for _, ch := range words[i] {
			tempArr[ch-'a']++
		}

		for i := 0; i < len(arr); i++ {
			arr[i] = min(arr[i], tempArr[i])
		}
	}

	res := make([]string, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		count := arr[i]
		for j := 0; j < count; j++ {
			res = append(res, string(rune(i+'a')))
		}
	}

	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	var counts [1001]int
	res := make([]int, 0, min(len(nums1), len(nums2)))

	for _, num := range nums1 {
		counts[num]++
	}

	for _, num := range nums2 {
		if counts[num] > 0 {
			res = append(res, num)
			counts[num]--
		}
	}
	return res
}
