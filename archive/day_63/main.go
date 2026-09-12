package main

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		i = getNextValid(s, i)
		j = getNextValid(t, j)

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

func getNextValid(s string, idx int) int {
	counter := 0

	for idx >= 0 {
		if s[idx] == '#' {
			counter++
			idx--
		} else if counter > 0 {
			counter--
			idx--
		} else {
			return idx
		}
	}
	return -1
}

func findAnagrams(s string, sub string) []int {
	freq1, freq2 := [26]int{}, [26]int{}

	for i := range sub {
		freq1[sub[i]-'a']++
		freq2[s[i]-'a']++
	}

	begin := 0
	res := []int{}

	for end := len(sub); end < len(s); end++ {
		if freq1 == freq2 {
			res = append(res, begin)
		}

		freq2[s[begin]-'a']--
		freq2[s[end]-'a']++
		begin++
	}
	if freq1 == freq2 {
		res = append(res, begin)
	}
	return res
}

func findMaxConsecutiveOnes(nums []int) int {
	begin := 0
	wState := 0
	res := 0

	for end := range nums {
		if nums[end] == 0 {
			wState++
		}

		for wState > 1 {
			if nums[begin] == 0 {
				wState--
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	begin := 0
	res := 0
	m := map[byte]int{}

	for end := range s {
		m[s[end]]++

		for len(m) > 2 {
			m[s[begin]]--
			if m[s[begin]] == 0 {
				delete(m, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}
