package main

import "fmt"

func LPS(p string) []int {
	lps := make([]int, len(p))
	i, j := 1, 0

	for i < len(p) {
		if p[i] == p[j] {
			lps[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				lps[i] = 0
				i++
			} else {
				j = lps[j-1]
			}
		}
	}
	return lps
}

func main() {
	fmt.Println(LPS("lilila"))
}

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

func commonChars(words []string) []string {
	freq := [26]int{}
	for i := 0; i < len(words[0]); i++ {
		freq[words[0][i]-'a']++
	}

	for i := 1; i < len(words); i++ {
		var temp [26]int
		for j := 0; j < len(words[i]); j++ {
			temp[words[i][j]-'a']++
		}
		for c := 0; c < len(freq); c++ {
			freq[c] = min(freq[c], temp[c])
		}
	}
	res := []string{}
	for i := 0; i < len(freq); i++ {
		for freq[i] > 0 {
			res = append(res, string(rune(i+'a')))
			freq[i]--
		}
	}
	return res
}
