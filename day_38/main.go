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
