package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneEditDistance(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "insert middle", s: "ab", t: "acb", want: true},
		{name: "insert end", s: "ab", t: "abc", want: true},
		{name: "insert start", s: "ab", t: "cab", want: true},
		{name: "delete middle", s: "acb", t: "ab", want: true},
		{name: "delete end", s: "abc", t: "ab", want: true},
		{name: "delete start", s: "cab", t: "ab", want: true},
		{name: "replace", s: "ab", t: "ac", want: true},
		{name: "replace single", s: "1203", t: "1213", want: true},
		{name: "empty to single", s: "", t: "A", want: true},
		{name: "single to empty", s: "A", t: "", want: true},
		{name: "identical", s: "ab", t: "ab", want: false},
		{name: "both empty", s: "", t: "", want: false},
		{name: "diff > 1 length", s: "ab", t: "abcb", want: false},
		{name: "diff > 1 length reverse", s: "abcb", t: "ab", want: false},
		{name: "two replacements", s: "ab", t: "cd", want: false},
		{name: "two inserts needed", s: "a", t: "abc", want: false},
		{name: "two deletes needed", s: "abc", t: "a", want: false},
		{name: "single replace char", s: "a", t: "b", want: true},
		{name: "single replace case", s: "a", t: "A", want: true},
		{name: "transposition", s: "ab", t: "ba", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isOneEditDistance(tt.s, tt.t)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{name: "test1", nums: []int{3, 2, 2, 3}, val: 3, want: 2},
		{name: "test2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5},
		{name: "empty", nums: []int{}, val: 1, want: 0},
		{name: "no match", nums: []int{1, 2, 3}, val: 4, want: 3},
		{name: "all match", nums: []int{1, 1, 1}, val: 1, want: 0},
		{name: "single match", nums: []int{1}, val: 1, want: 0},
		{name: "single no match", nums: []int{1}, val: 2, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeElement(tt.nums, tt.val)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{name: "test1", nums1: []int{3, 2, 2, 3}, nums2: []int{2, 2}, want: []int{2}},
		{name: "test2", nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{4, 9}},
		{name: "empty1", nums1: []int{}, nums2: []int{1, 2}, want: []int{}},
		{name: "empty2", nums1: []int{1, 2}, nums2: []int{}, want: []int{}},
		{name: "no intersection", nums1: []int{1, 2}, nums2: []int{3, 4}, want: []int{}},
		{name: "all same", nums1: []int{1, 1, 1}, nums2: []int{1, 1}, want: []int{1}},
		{name: "single elements", nums1: []int{1}, nums2: []int{1}, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersection(tt.nums1, tt.nums2)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "test1", nums: []int{1, 1, 1, 2, 2, 3}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeDuplicates(tt.nums)
			assert.Equal(t, tt.want, res)
		})
	}
}
