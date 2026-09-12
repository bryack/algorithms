package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCharByIndexUnreachable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		inputString string
		index       int
	}{
		{
			name:        "empty string",
			inputString: "",
			index:       0,
		},
		{
			name:        "index greater than length",
			inputString: "123",
			index:       10_000,
		},
		{
			name:        "negative index",
			inputString: "椅子",
			index:       -42,
		},
		{
			name:        "index out of bounds",
			inputString: "椅1",
			index:       3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Panics(t, func() {
				GetCharByIndex(tc.inputString, tc.index)
			})
		})
	}
}

func TestGetCharByIndex(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		inputString string
		index       int
		expected    rune
	}{
		{
			name:        "simple get",
			inputString: "abcdef",
			index:       4,
			expected:    'e',
		},
		{
			name:        "chinese symbols",
			inputString: "椅子摆得不整重新摆一下儿",
			index:       4,
			expected:    '不',
		},
		{
			name:        "arabic symbols",
			inputString: "كورنييف جورج الكسندروفيتش",
			index:       8,
			expected:    'ج',
		},
		{
			name:        "emoji",
			inputString: "a🙂🙃🌚😑😐z",
			index:       3,
			expected:    '🌚',
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := GetCharByIndex(tc.inputString, tc.index)

			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetCharByIndexAll(t *testing.T) {
	t.Parallel()
	s := "e\u0301👨‍👩‍👧‍👦"
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		require.Equal(t, runes[i], GetCharByIndex(s, i))
	}

	require.Panics(t, func() { GetCharByIndex(s, len(runes)) })
}

func TestCharByIndexCopy(t *testing.T) {
	result := testing.Benchmark(func(b *testing.B) {
		s := strings.Repeat("🙃", 10_000)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			GetCharByIndex(s, i%10_000)
		}
	})

	require.EqualValues(t, 0, result.AllocsPerOp())
}
