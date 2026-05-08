package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOnlyLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "test1", input: "ab-cd", want: "dc-ba"},
		{name: "test2", input: "a-bC-dEf-ghIj", want: "j-Ih-gfE-dCba"},
		{name: "test3", input: "Test1ng-Leet=code-Q!", want: "Qedo1ct-eeLg=ntse-T!"},
		{name: "test4", input: "a-", want: "a-"},
		{name: "test5", input: "----", want: "----"},
		{name: "test6", input: "asdfsf", want: "fsfdsa"},
	}

	for _, tt := range tests {
		res := reverseOnlyLetters(tt.input)
		assert.Equal(t, tt.want, res)
	}
}
