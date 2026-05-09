package main

func longestPalindrome(s string) string {
	length := 0
	var arr [2]int

	for center := 0; center < len(s)-1; center++ {
		i, j := center, center
		for i >= 0 && j < len(s) {
			if s[i] != s[j] {
				break
			}
			if length < j-i+1 {
				arr[0] = i
				arr[1] = j
				length = j - i + 1
			}
			i--
			j++
		}
		i, j = center, center+1
		for i >= 0 && j < len(s) {
			if s[i] != s[j] {
				break
			}
			if length < j-i+1 {
				arr[0] = i
				arr[1] = j
				length = j - i + 1
			}
			i--
			j++
		}
	}

	return s[arr[0] : arr[1]+1]
}
