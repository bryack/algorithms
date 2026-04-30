package main

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	var arr [26]int
	for _, char := range s {
		arr[char-'a']++
	}

	for i, char := range s {
		if arr[char-'a'] == 1 {
			return i
		}
	}
	return -1
}
