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
	assert.Equal(t, 1, random)
	res = set.Remove(1)
	assert.True(t, res)
	assert.Equal(t, []int{2}, set.nums)
	assert.Equal(t, map[int]int{2: 0}, set.m)
	res = set.Insert(2)
	assert.False(t, res)
	random = set.GetRandom()
	assert.Equal(t, 2, random)
}
