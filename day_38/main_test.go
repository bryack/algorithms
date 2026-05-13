package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "1", s: "ab#c", t: "ad#c", want: true},
		{name: "2", s: "ab##", t: "c#d#", want: true},
		{name: "3", s: "a#c", t: "b", want: false},
		{name: "4", s: "xywrrmp", t: "xywrrmu#p", want: true},
		{name: "5", s: "a", t: "aa#a", want: false},
		{name: "empty both", s: "", t: "", want: true},
		{name: "only backspaces", s: "####", t: "##a#b##", want: true},
		{name: "multiple consecutive backspaces", s: "abc#d##e", t: "ae", want: true},
		{name: "leading backspaces", s: "##a", t: "a", want: true},
		{name: "backspaces exceed chars", s: "bxj##tw", t: "bxj###tw", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := backspaceCompare(tt.s, tt.t)
			assert.Equal(t, tt.want, res)
		})
	}
}
