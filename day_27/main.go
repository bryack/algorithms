package main

func groupAnagrams(strs []string) [][]string {
	s := make([][]string, len(strs))
	m := make(map[[26]int]int, len(strs))
	idx := 0

	for _, str := range strs {
		var arr [26]int

		for _, char := range str {
			arr[char-'a']++
		}

		if val, ok := m[arr]; ok {
			s[val] = append(s[val], str)
		} else {
			m[arr] = idx
			s[idx] = append(s[idx], str)
			idx++
		}
	}
	return s[:idx]
}

func sortedSquares(nums []int) []int {
	s := make([]int, len(nums))
	idx := len(nums) - 1
	i := 0
	j := len(nums) - 1

	for i <= j {
		first := nums[i] * nums[i]
		second := nums[j] * nums[j]

		if first > second {
			s[idx] = first
			i++
		} else {
			s[idx] = second
			j--
		}
		idx--
	}
	return s
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum > target {
			j--
		}
		if sum < target {
			i++
		}
	}
}
