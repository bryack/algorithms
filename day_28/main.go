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

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			if isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1) {
				return true
			} else {
				return false
			}
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
