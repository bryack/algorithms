package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name      string
		firstStr  string
		secondStr string
		want      bool
	}{
		{name: "test 1", firstStr: "egg", secondStr: "add", want: true},
		{name: "test 2", firstStr: "foo", secondStr: "bar", want: false},
		{name: "test 3", firstStr: "paper", secondStr: "title", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isIsomorphic(tt.firstStr, tt.secondStr)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{name: "test 1", nums: []int{1, 2, 3, 1}, k: 3, want: true},
		{name: "test 2", nums: []int{1, 0, 1, 1}, k: 1, want: true},
		{name: "test 1", nums: []int{1, 2, 3, 1, 2, 3}, k: 2, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsNearbyDuplicate(tt.nums, tt.k)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{name: "test 1", words: []string{"bella", "label", "roller"}, want: []string{"e", "l", "l"}},
		{name: "test 1", words: []string{"cool", "lock", "cook"}, want: []string{"c", "o"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := commonChars(tt.words)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name        string
		firstSlice  []int
		secondSlice []int
		want        []int
	}{
		{name: "test 1", firstSlice: []int{1, 2, 2, 1}, secondSlice: []int{2, 2}, want: []int{2, 2}},
		{name: "test 2", firstSlice: []int{4, 9, 5}, secondSlice: []int{9, 4, 9, 8, 4}, want: []int{9, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersect(tt.firstSlice, tt.secondSlice)
			assert.Equal(t, tt.want, result)
		})
	}
}
