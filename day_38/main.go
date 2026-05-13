package main

func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	slow := n
	fast := getNext(n)

	for fast != 1 && fast != slow {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

func pivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	left := 0

	for i, n := range nums {
		total -= n
		if total == left {
			return i
		}
		left += n
	}
	return -1
}

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		m[n] = struct{}{}
	}

	maxLen := 0
	for k := range m {
		if _, ok := m[k-1]; !ok {
			currentNum := k
			currentLen := 1
			for {
				if _, ok := m[currentNum+1]; !ok {
					break
				}
				currentNum++
				currentLen++
			}
			maxLen = max(maxLen, currentLen)
		}
	}
	return maxLen
}

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		i = findNextValidIndex(s, i)
		j = findNextValidIndex(t, j)

		if i < 0 && j < 0 {
			return true
		}

		if i < 0 || j < 0 {
			return false
		}

		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func findNextValidIndex(str string, i int) int {
	bs := 0
	for i >= 0 {
		if str[i] == '#' {
			bs++
			i--
		} else if bs > 0 {
			bs--
			i--
		} else {
			return i
		}
	}
	return -1
}
