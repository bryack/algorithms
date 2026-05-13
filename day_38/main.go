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
