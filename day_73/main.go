package main

type TwoSum struct {
	data map[int]int
}

func Constructor() TwoSum {
	return TwoSum{data: make(map[int]int)}
}

func (t *TwoSum) Add(number int) {
	t.data[number]++
}

func (t *TwoSum) Find(value int) bool {
	for key, val := range t.data {
		c := value - key

		if c == key {
			if val > 1 {
				return true
			}
		} else if _, ok := t.data[c]; ok {
			return true
		}
	}
	return false
}

func isOneEditDistance(s string, t string) bool {
	n, m := len(s), len(t)
	if n > m {
		s, t = t, s
		n, m = m, n
	}
	for i := range n {
		if s[i] != t[i] {
			if m == n {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}
	return m-n == 1
}
