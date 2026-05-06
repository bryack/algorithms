package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	var arr [26]int

	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i]-'a']++
	}

	for i := 0; i < len(magazine); i++ {
		if arr[magazine[i]-'a'] > 0 {
			arr[magazine[i]-'a']--
		}
	}
	return arr == [26]int{}
}

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}
	if k == 0 || len(nums) == 1 {
		return
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
