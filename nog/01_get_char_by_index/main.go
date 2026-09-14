package main

import "unicode/utf8"

// GetCharByIndex returns the i-th character from the given string.
func GetCharByIndex(str string, idx int) rune {
	var counter, offset int

	for offset < len(str) {
		r, size := utf8.DecodeRuneInString(str[offset:])
		if counter == idx {
			return r
		}
		offset += size
		counter++
	}
	panic("index out of bounds")
}
