package main

import "fmt"

func reverseWords(s string) string {
	buf := []byte(s)

	reverse := func(i, j int) {
		fmt.Println("i, j", i, j)
		for i < j {
			buf[i], buf[j] = buf[j], buf[i]
			i++
			j--
		}
	}

	for end := 0; end < len(s); end++ {
		start := end

		for end < len(buf) && buf[end] != ' ' {
			end++
		}
		reverse(start, end-1)
	}
	return string(buf)
}
