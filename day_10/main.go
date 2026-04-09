package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func canJump(nums []int) bool {
	if nums[0] == 0 && len(nums) > 1 {
		return false
	}
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if maxReach >= len(nums)-1 {
			return true
		}
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}

func cheat(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	var size int
	fmt.Fscan(reader, &size)
	if size == 0 {
		return
	}
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	d := 0
	for i := 0; i < size; i++ {
		if arr[i] != 0 {
			arr[d] = arr[i]
			d++
		}
	}

	for i := d; i < size; i++ {
		arr[i] = 0
	}

	for i := 0; i < size; i++ {
		if i == size-1 {
			fmt.Fprintf(w, "%d", arr[i])
		} else {
			fmt.Fprintf(w, "%d ", arr[i])
		}
	}
}

func clean(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	var size int
	fmt.Fscan(reader, &size)

	if size == 0 {
		return
	}

	arr := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	count := 0
	for i := 0; i < size+1; i++ {
		if arr[i] != arr[size] {
			if count > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, strconv.Itoa(arr[i]))
			count++
		}
	}
}
