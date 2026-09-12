package main

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	var arr [26]int
	for _, char := range s {
		arr[char-'a']++
	}

	for i, char := range s {
		if arr[char-'a'] == 1 {
			return i
		}
	}
	return -1
}

func isHappy(n int) bool {
	m := map[int]struct{}{}

	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		} else {
			m[n] = struct{}{}
		}
		squareSum := 0
		for n > 0 {
			squareSum += (n % 10) * (n % 10)
			n /= 10
		}
		n = squareSum
	}
	return true
}

func getNext(n int) int {
	squareSum := 0
	for n > 0 {
		squareSum += (n % 10) * (n % 10)
		n /= 10
	}
	return squareSum
}

func isHappyFloyd(n int) bool {
	slow := n
	fast := getNext(n)

	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

func isValidSudoku(board [][]byte) bool {
	var rowArr [9][9]bool
	var colArr [9][9]bool
	var boxArr [9][9]bool

	for r, row := range board {
		for c, num := range row {
			if num == '.' {
				continue
			}

			if rowArr[r][num-'1'] == true || colArr[c][num-'1'] == true || boxArr[(r/3)*3+(c/3)][num-'1'] == true {
				return false
			}

			rowArr[r][num-'1'] = true
			colArr[c][num-'1'] = true
			boxArr[(r/3)*3+(c/3)][num-'1'] = true
		}
	}
	return true
}

func findPairs(nums []int, k int) int {
	m := make(map[int]int, len(nums))

	for _, n := range nums {
		m[n]++
	}

	if k == 0 {
		count := 0
		for _, value := range m {
			if value > 1 {
				count++
			}
		}
		return count
	}

	if k > 0 {
		count := 0
		for key := range m {
			if _, ok := m[key+k]; ok {
				count++
			}
		}
		return count
	}
	return 0
}
