package main

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	return count == [26]int{}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func findTwiceSymbols(s string) []rune {
	if len(s) <= 1 {
		return nil
	}
	var arr [26]int
	res := make([]rune, 0, len(s)/2)

	for _, r := range s {
		arr[r-'a']++
	}

	var indexCounter [26]int
	for _, r := range s {
		indexCounter[r-'a']++
		if indexCounter[r-'a'] == 2 && indexCounter[r-'a'] == arr[r-'a'] {
			res = append(res, r)
		}
	}

	// for i, v := range arr {
	// 	if v == 2 {
	// 		res = append(res, rune(i+'a'))
	// 	}
	// }

	return res
}

func firstUniqChar(s string) int {
	var arr [26]int

	for _, r := range s {
		arr[r-'a']++
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func repeatedCharacter(s string) byte {
	var arr [26]bool

	for _, r := range s {
		if arr[r-'a'] == true {
			return byte(r)
		}
		arr[r-'a'] = true
	}
	return 0
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	groups := make([][]string, 0, len(strs))
	var arr [26]int
	m := make(map[[26]int][]string)

	for _, word := range strs {
		for i := 0; i < len(word); i++ {
			arr[word[i]-'a']++
		}
		m[arr] = append(m[arr], word)
		arr = [26]int{}
	}

	for _, v := range m {
		groups = append(groups, v)
	}

	return groups
}
