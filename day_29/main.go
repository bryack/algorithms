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
				s = append(s, string(i+'a'))
			}
		}
	}
	return s
}
