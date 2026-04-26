package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name      string
		firstStr  string
		secondStr string
		want      bool
	}{
		{name: "test 1", firstStr: "a", secondStr: "b", want: false},
		{name: "test 2", firstStr: "aa", secondStr: "aab", want: true},
		{name: "test 3", firstStr: "aa", secondStr: "ab", want: false},
		{name: "test 4", firstStr: "aa", secondStr: "aba", want: true},
		{name: "test 5", firstStr: "aaaa", secondStr: "aab", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canConstruct(tt.firstStr, tt.secondStr)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "test 1", text: "nlaebolko", want: 1},
		{name: "test 2", text: "loonbalxballpoon", want: 2},
		{name: "test 3", text: "leetcode", want: 0},
		{name: "test 3", text: "loonbalxballpoonloonbalxballpoonnlaebolko", want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxNumberOfBalloons(tt.text)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "test 1", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "test 2", nums: []int{1}, k: 1, want: []int{1}},
		{name: "test 3", nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, k: 2, want: []int{2, 1}},
		{name: "test 1", nums: []int{1, 1, 1, 2, 2, 3, 3, 3, 3}, k: 2, want: []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)
			assert.ElementsMatch(t, tt.want, result)
		})
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "test 1", nums: []int{1, 2, 3, 4, 5, 6}, target: 2, want: 1},
		{name: "test 1", nums: []int{1, 2, 3, 4, 5, 6}, target: 6, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binarySearch(tt.nums, tt.target)
			assert.Equal(t, tt.want, result)
		})
	}
}
