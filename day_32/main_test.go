package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	res := maxNumberOfBalloons("nllbblooon")
	assert.Equal(t, 0, res)
}
