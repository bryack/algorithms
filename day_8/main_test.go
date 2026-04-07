package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "100 - 202", input: 100, want: "202"},
		{name: "15 - 21", input: 15, want: "21"},
		{name: "1 - 1", input: 1, want: "1"},
		{name: "-7 - -10", input: -7, want: "-10"},
		{name: "-15 - -21", input: -15, want: "-21"},
		{name: "0 - 0", input: 0, want: "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToBase7(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
