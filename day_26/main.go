package main

func repeatedCharacter(s string) string {
	var arr [26]bool

	for _, char := range s {
		if arr[char-'a'] {
			return string(char)
		}
		arr[char-'a'] = true
	}
	return ""
}
