package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{name: "test 1", input: []int{1, 2, 3, 1}, want: true},
		{name: "test 2", input: []int{1, 2, 3, 4}, want: false},
		{name: "test 3", input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name      string
		firstStr  string
		secondStr string
		want      bool
	}{
		{name: "test 1", firstStr: "anagram", secondStr: "nagaram", want: true},
		{name: "test 2", firstStr: "rat", secondStr: "car", want: false},
		{name: "test 3", firstStr: "anagra", secondStr: "nagaram", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.firstStr, tt.secondStr)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   []int
	}{
		{name: "test 1", arr: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "test 2", arr: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "test 3", arr: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.arr, tt.target)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindTwiceSymbols(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{name: "test 1", input: "programming", want: []rune{'r', 'm', 'g'}},
		{name: "test 2", input: "leecode", want: []rune{}},
		{name: "test 3", input: "cat", want: []rune{}},
		{name: "test 4", input: "ppp", want: []rune{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findTwiceSymbols(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "test 1", input: "leetcode", want: 0},
		{name: "test 2", input: "loveleetcode", want: 2},
		{name: "test 3", input: "aabb", want: -1},
		{name: "test 3", input: "a", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := firstUniqChar(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestRepeatedCharacter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  byte
	}{
		{name: "test 1", input: "abccbaacz", want: byte('c')},
		{name: "test 2", input: "abcdd", want: byte('d')},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repeatedCharacter(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{name: "test 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}},
		},
		{name: "test 2",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{name: "test 3",
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)
			assert.ElementsMatch(t, tt.want, result)
		})
	}
}
