package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet(t *testing.T) {
	set := Constructor()

	res := set.Insert(1)
	assert.True(t, res)
	res = set.Remove(2)
	assert.False(t, res)
	res = set.Insert(2)
	assert.True(t, res)
	assert.Equal(t, []int{1, 2}, set.nums)
	assert.Equal(t, map[int]int{1: 0, 2: 1}, set.m)
	random := set.GetRandom()
	res = set.Remove(1)
	assert.True(t, res)
	assert.Equal(t, []int{2}, set.nums)
	assert.Equal(t, map[int]int{2: 0}, set.m)
	res = set.Insert(2)
	assert.False(t, res)
	random = set.GetRandom()
	assert.Equal(t, 2, random)
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "test 1",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "test 2",
			input: "race a car",
			want:  false,
		},
		{
			name:  "test 3",
			input: " ",
			want:  true,
		},
		{
			name:  "test 4",
			input: "a1,ma:  m1a",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
