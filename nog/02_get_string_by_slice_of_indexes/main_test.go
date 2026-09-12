package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStringBySliceOfIndexesAlloc(t *testing.T) {
	result := testing.Benchmark(func(b *testing.B) {
		s := "1"
		ind := slices.Repeat([]int{0}, 1_000_000)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			GetStringBySliceOfIndexes(s, ind)
		}
	})

	require.LessOrEqual(t, result.AllocsPerOp(), int64(2))
}

func TestGetStringBySliceOfIndexesBenchmark(t *testing.T) {
	result := testing.Benchmark(func(b *testing.B) {
		b.StopTimer()
		inputString := strings.Repeat("椅子摆得不整重新摆一下儿", 42*42*42*42)
		targetSlice := make([]int, 42*42)

		for i := 0; i < len(targetSlice); i++ {
			targetSlice[i] = 42
		}
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			GetStringBySliceOfIndexes(inputString, targetSlice)
		}
	})

	require.Less(t, result.AllocsPerOp(), int64(127))
}

func TestGetStringBySliceOfIndexesPerformance(t *testing.T) {
	const n = 100_000

	solution := testing.Benchmark(func(b *testing.B) {
		str := strings.Repeat("a", n)
		indexes := make([]int, n)
		for i := 0; i < n; i++ {
			indexes[i] = n - 1
		}

		for b.Loop() {
			GetStringBySliceOfIndexes(str, indexes)
		}
	})

	emulator := testing.Benchmark(func(b *testing.B) {
		str := strings.Repeat("a", n)
		b.ResetTimer()

		var keep []rune

		for b.Loop() {
			runes := []rune(str)

			keep = runes
			_ = keep

			GetCharByIndex(str, n-1)
		}
	})

	require.LessOrEqual(t, float64(solution.NsPerOp())/float64(emulator.NsPerOp()), 300.)
}

func TestGetStringBySliceOfIndexesUnreachable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		inputString string
		indexes     []int
	}{
		{
			name:        "empty string",
			inputString: "",
			indexes:     []int{0},
		},
		{
			name:        "index greater than length",
			inputString: "abcdef",
			indexes:     []int{1, 239, 5},
		},
		{
			name:        "negative index",
			inputString: "abcdef",
			indexes:     []int{1, 3, -3, 3, 4, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Panics(t, func() {
				GetStringBySliceOfIndexes(tc.inputString, tc.indexes)
			})
		})
	}
}

func TestGetStringBySliceOfIndexes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		inputString    string
		indexes        []int
		expectedString string
	}{
		{
			name:           "simple get",
			inputString:    "abcdef",
			indexes:        []int{1, 3, 5},
			expectedString: "bdf",
		},
		{
			name:           "simple get random index order",
			inputString:    "abcdef",
			indexes:        []int{1, 3, 3, 3, 4, 2},
			expectedString: "bdddec",
		},
		{
			name:           "chinese symbols",
			inputString:    "椅子摆得不整重新摆一下儿",
			indexes:        []int{0, 1, 3, 5},
			expectedString: "椅子得整",
		},
		{
			name:           "arabic symbols",
			inputString:    "كورنييف جورج الكسندروفيتش",
			indexes:        []int{0, 8, 15},
			expectedString: "كجك",
		},
		{
			name:           "emoji",
			inputString:    "a🙂🙃🌚😑😐z",
			indexes:        []int{6, 3, 2, 5, 4, 4},
			expectedString: "z🌚🙃😐😑😑",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := GetStringBySliceOfIndexes(tc.inputString, tc.indexes)

			require.Equal(t, tc.expectedString, actual)
		})
	}
}
