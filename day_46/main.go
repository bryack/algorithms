package main

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func isSubsequence(sub string, s string) bool {
	j := 0

	for i := 0; i < len(s); i++ {
		if j < len(sub) && s[i] == sub[j] {
			j++
		}
	}
	return len(sub) == j
}

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		i = nextValid(s, i)
		j = nextValid(t, j)

		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 {
			return false
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func nextValid(s string, i int) int {
	count := 0
	for i >= 0 {
		if s[i] == '#' {
			count++
		} else if count > 0 {
			count--
		} else {
			return i
		}
		i--
	}
	return -1
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

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func mergeAlternately(word1 string, word2 string) string {
	buf := make([]byte, len(word1)+len(word2))
	idx := 0
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		buf[idx] = word1[i]
		i++
		idx++
		buf[idx] = word2[j]
		j++
		idx++
	}
	for i < len(word1) {
		buf[idx] = word1[i]
		i++
		idx++
	}
	for j < len(word2) {
		buf[idx] = word2[j]
		j++
		idx++
	}

	return string(buf)
}
