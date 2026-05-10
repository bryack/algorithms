package main

import (
	"fmt"
	"io"
)

func concatTwoArrays(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	i := 0
	j := 0
	count := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res[count] = a[i]
			i++
			count++
		} else {
			res[count] = b[j]
			j++
			count++
		}
	}

	if i < len(a) {
		copy(res[count:], a[i:])
	}

	if j < len(b) {
		copy(res[count:], b[j:])
	}

	return res
}

func findHighScore(r io.Reader, w io.Writer) {
	var size int
	fmt.Fscan(r, &size)

	if size == 0 {
		return
	}

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Fscan(r, &arr[i])
	}

	c := arr[0]

	for i := 0; i < size; i++ {
		if arr[i] > c {
			c = arr[i]
		}
	}

	fmt.Fprint(w, c)
}
