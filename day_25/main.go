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
