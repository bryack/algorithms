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

func frequencySort(s string) string {
	if len(s) <= 2 {
		return s
	}

	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	bucket := make([][]byte, len(s)+1)
	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]byte, 0, len(s))
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}
		for _, b := range bucket[i] {
			for j := 0; j < i; j++ {
				res = append(res, b)
			}
		}
	}

	return string(res)
}
