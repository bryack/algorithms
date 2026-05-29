package main

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func maxArea(height []int) int {
	maxA := 0
	i, j := 0, len(height)-1

	for i < j {
		curA := 0
		if height[i] < height[j] {
			curA = height[i] * (j - i)
			i++
		} else {
			curA = height[j] * (j - i)
			j--
		}
		maxA = max(maxA, curA)
	}
	return maxA
}

func main() {
	s := strings.Repeat("1", os.Getpagesize())
	fmt.Println(s)
	buf := *(*[]byte)(unsafe.Pointer(&s))

	buf[0] = 2
	fmt.Println(buf[0])
	fmt.Println(s)
}
