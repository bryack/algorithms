package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	res := frequencySort(nums)
	assert.Equal(t, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}, res)
}
