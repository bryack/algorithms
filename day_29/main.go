package main

func commonChars(words []string) []string {
	var arrDefault [26]int
	var s []string

	for _, char := range words[0] {
		arrDefault[char-'a']++
	}

	for i := 1; i < len(words); i++ {
		var arrTemp [26]int
		for _, char := range words[i] {
			arrTemp[char-'a']++
		}

		for j := 0; j < len(arrDefault); j++ {
			arrDefault[j] = min(arrDefault[j], arrTemp[j])
		}
	}

	for i := 0; i < len(arrDefault); i++ {
		if arrDefault[i] > 0 {
			for j := 0; j < arrDefault[i]; j++ {
				s = append(s, string(rune(i+'a')))
			}
		}
	}
	return s
}

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	slow, fast := 0, 1

	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			nums[slow+1] = nums[fast]
			fast++
			slow++
		}
	}
	return slow + 1
}
