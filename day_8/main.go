package main

import (
	"strconv"
	"strings"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	negative := num < 0
	if negative {
		num = -num
	}

	s := []rune{}
	for {
		d := num / 7
		r := num % 7
		s = append(s, rune(r))
		if d == 0 {
			break
		}
		num = d
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	var builder strings.Builder
	for _, v := range s {
		builder.WriteString(strconv.Itoa(int(v)))
	}

	if negative {
		return "-" + builder.String()
	}

	return builder.String()
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	s := []rune{}
	n := uint32(num)
	for {
		d := n / 16
		r := n % 16
		if r < 10 {
			s = append(s, rune('0'+r))
		} else {
			s = append(s, rune('a'+r%10))
		}
		if d == 0 {
			break
		}
		n = d
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
