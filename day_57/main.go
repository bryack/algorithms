package main

import (
	"fmt"
	"math"
	"strconv"
)

func summaryRanges(nums []int) []string {
	i, j := 0, 0
	res := []string{}

	for i < len(nums) {

		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}

		if nums[i] == nums[j] {
			res = append(res, fmt.Sprintf("%d", nums[j]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[j], nums[i]))
		}

		i++
		j = i
	}
	return res
}

func compress(chars []byte) int {
	end := 0
	idx := 0

	for end < len(chars) {
		start := end

		for end < len(chars) && chars[start] == chars[end] {
			end++
		}

		chars[idx] = chars[start]
		idx++

		count := end - start
		if count > 1 {
			num := strconv.Itoa(count)
			for _, n := range num {
				chars[idx] = byte(n)
				idx++
			}
		}
	}
	return idx
}

func isOneEditDistance(s string, t string) bool {
	m, n := len(s), len(t)

	if m > n {
		s, t = t, s
		m, n = n, m
	}

	if n-m > 1 {
		return false
	}

	for i := 0; i < m; i++ {
		if s[i] != t[i] {
			if m == n {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}
	return n-m == 1
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		li := max(firstList[i][0], secondList[j][0])
		ri := min(firstList[i][1], secondList[j][1])

		if li <= ri {
			res = append(res, []int{li, ri})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func reverseWords(s string) string {
	buf := []byte(s)
	end := 0

	for end < len(buf) {
		start := end

		for end < len(buf) && buf[end] != ' ' {
			end++
		}

		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			buf[i], buf[j] = buf[j], buf[i]
		}

		end++

	}
	return string(buf)
}

func addStrings(num1 string, num2 string) string {
	res := []byte{}
	carry := 0
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		carry = sum / 10
		res = append(res, byte(sum%10)+'0')
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func countSubstrings(s string) int {
	res := 0
	n := len(s)

	for center := 0; center < 2*n+1; center++ {
		left := center / 2
		right := left + center%2

		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	return res
}

func findMaxAverage(nums []int, k int) float64 {
	begin := 0
	res := math.MinInt
	wState := 0

	for end := range nums {
		wState += nums[end]

		if end-begin+1 == k {
			res = max(res, wState)
			wState -= nums[begin]
			begin++
		}
	}
	return float64(res) / float64(k)
}

func numOfSubarrays(arr []int, k int, threshold int) int {
	begin := 0
	res := 0
	wState := 0

	for end := range arr {
		wState += arr[end]

		if end-begin+1 == k {
			if wState/k >= threshold {
				res++
			}
			wState -= arr[begin]
			begin++
		}
	}
	return res
}

func countGoodSubstrings(s string) int {
	begin := 0
	res := 0

	for end := range s {

		if end-begin+1 == 3 {
			if s[end] != s[begin] && s[begin+1] != s[end] && s[begin] != s[begin+1] {
				res++
			}
			begin++
		}
	}
	return res
}

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))

	for i := range res {
		res[i] = -1
	}

	begin := 0
	wState := 0

	for end := range nums {
		wState += nums[end]

		if end-begin == 2*k {
			res[begin+k] = wState / (end - begin + 1)
			wState -= nums[begin]
			begin++
		}
	}
	return res
}

func maximumSubarraySum(nums []int, k int) int64 {
	begin := 0
	m := map[int]int{}
	res := 0
	wState := 0

	for end := range nums {
		wState += nums[end]
		m[nums[end]]++

		if end-begin+1 == k {
			if len(m) == k {
				res = max(res, wState)
			}

			m[nums[begin]]--
			if m[nums[begin]] <= 0 {
				delete(m, nums[begin])
			}

			wState -= nums[begin]
			begin++
		}
	}
	return int64(res)
}

func maxVowels(s string, k int) int {
	begin := 0
	res := 0
	wState := 0

	for end := range s {
		if s[end] == 'a' || s[end] == 'e' || s[end] == 'i' || s[end] == 'o' || s[end] == 'u' {
			wState++
		}

		if end-begin+1 == k {
			res = max(res, wState)

			if s[begin] == 'a' || s[begin] == 'e' || s[begin] == 'i' || s[begin] == 'o' || s[begin] == 'u' {
				wState--
			}
			begin++
		}
	}
	return res
}

func getSubarrayBeauty(nums []int, k int, x int) []int {
	begin := 0
	res := []int{}
	var arr [101]int

	for end := range nums {
		arr[nums[end]+50]++

		if end-begin+1 == k {
			count := 0
			found := false
			for i := 0; i < 50; i++ {
				count += arr[i]
				if count >= x {
					res = append(res, i-50)
					found = true
					break
				}
			}

			if !found {
				res = append(res, 0)
			}

			arr[nums[begin]+50]--
			begin++
		}
	}
	return res
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	wState := 0
	end := n - k

	for i := end; i < n; i++ {
		wState += cardPoints[i]
	}

	maxP := wState
	begin := 0

	for end < n {
		wState = wState - cardPoints[end] + cardPoints[begin]
		maxP = max(maxP, wState)
		end++
		begin++
	}
	return maxP
}

func minSubArrayLen(target int, nums []int) int {
	begin := 0
	res := len(nums) + 1
	wState := 0

	for end := range nums {
		wState += nums[end]

		for wState >= target {
			res = min(res, end-begin+1)
			wState -= nums[begin]
			begin++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func longestOnes(nums []int, k int) int {
	begin := 0
	res := 0
	wState := 0

	for end := range nums {
		if nums[end] == 0 {
			wState++
		}

		for wState > k {
			if nums[begin] == 0 {
				wState--
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func longestSubarray(nums []int) int {
	begin := 0
	res := 0
	wState := 0

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
	return res - 1
}

func totalFruit(fruits []int) int {
	begin := 0
	m := map[int]int{}
	res := 0

	for end := range fruits {
		m[fruits[end]]++

		for len(m) > 2 {
			m[fruits[begin]]--
			if m[fruits[begin]] == 0 {
				delete(m, fruits[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	begin := 0
	wState := map[byte]int{}
	res := 0

	for end := range s {
		wState[s[end]]++

		for wState[s[end]] > 1 {
			wState[s[begin]]--
			if wState[s[begin]] == 0 {
				delete(wState, s[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func characterReplacement(s string, k int) int {
	begin := 0
	freq := [26]int{}
	maxFreq := 0
	res := 0

	for end := range s {
		freq[s[end]-'A']++
		maxFreq = max(maxFreq, freq[s[end]-'A'])

		for (end-begin+1)-maxFreq > k {
			freq[s[begin]-'A']--
			begin++
		}

		res = max(res, end-begin+1)

	}
	return res
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count, s2Count := [26]int{}, [26]int{}

	for i := range s1 {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	for end := len(s1); end < len(s2); end++ {
		if s1Count == s2Count {
			return true
		}
		s2Count[s2[end]-'a']++
		s2Count[s2[end-len(s1)]-'a']--
	}
	return s1Count == s2Count
}

func findAnagrams(s string, sub string) []int {
	if len(sub) > len(s) {
		return []int{}
	}
	res := []int{}

	freqS, freqSub := [26]int{}, [26]int{}
	for i := range sub {
		freqS[s[i]-'a']++
		freqSub[sub[i]-'a']++
	}

	end := len(sub)
	begin := end - len(sub)
	for ; end < len(s); end++ {
		if freqS == freqSub {
			res = append(res, begin)
		}
		freqS[s[end]-'a']++
		freqS[s[begin]-'a']--
		begin++
	}
	if freqS == freqSub {
		res = append(res, begin)
	}
	return res
}
