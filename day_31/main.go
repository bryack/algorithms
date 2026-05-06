package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	var arr [26]int

	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i]-'a']++
	}

	for i := 0; i < len(magazine); i++ {
		if arr[magazine[i]-'a'] > 0 {
			arr[magazine[i]-'a']--
		}
	}
	return arr == [26]int{}
}
