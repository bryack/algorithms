package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []bool // ровно по одному элементу на каждую find-операцию, в порядке их вызова
	}{
		{
			name:       "базовый случай из условия LeetCode",
			operations: []string{"add", "add", "add", "find", "find"},
			values:     []int{1, 3, 5, 4, 7},
			expected:   []bool{true, false}, // find(4)=true (1+3), find(7)=false
		},
		{
			name:       "find на пустой структуре",
			operations: []string{"find"},
			values:     []int{5},
			expected:   []bool{false},
		},
		{
			name:       "одно и то же число дважды даёт валидную пару с самим собой",
			operations: []string{"add", "add", "find"},
			values:     []int{3, 3, 6},
			expected:   []bool{true},
		},
		{
			name:       "число добавлено один раз — само с собой пару не образует",
			operations: []string{"add", "find"},
			values:     []int{3, 6},
			expected:   []bool{false},
		},
		{
			name:       "три одинаковых числа — тоже валидно (cnt > 1)",
			operations: []string{"add", "add", "add", "find"},
			values:     []int{5, 5, 5, 10},
			expected:   []bool{true},
		},
		{
			name:       "отрицательные числа",
			operations: []string{"add", "add", "find", "find"},
			values:     []int{-1, -2, -3, 5},
			expected:   []bool{true, false},
		},
		{
			name:       "ноль как одно из слагаемых",
			operations: []string{"add", "add", "find"},
			values:     []int{0, 5, 5},
			expected:   []bool{true},
		},
		{
			name:       "два нуля дают сумму 0",
			operations: []string{"add", "add", "find"},
			values:     []int{0, 0, 0},
			expected:   []bool{true},
		},
		{
			name:       "один ноль не даёт сумму 0 сам с собой",
			operations: []string{"add", "find"},
			values:     []int{0, 0},
			expected:   []bool{false},
		},
		{
			name:       "value недостижим — числа есть, но комбинации нет",
			operations: []string{"add", "add", "add", "find"},
			values:     []int{1, 2, 3, 100},
			expected:   []bool{false},
		},
		{
			name:       "смешанные положительные и отрицательные, сумма в ноль",
			operations: []string{"add", "add", "find"},
			values:     []int{-5, 5, 0},
			expected:   []bool{true},
		},
		{
			name:       "find вызван несколько раз подряд без новых add",
			operations: []string{"add", "add", "find", "find", "find"},
			values:     []int{1, 5, 6, 6, 2},
			expected:   []bool{true, true, false}, // 1+5=6 (найдено), затем ещё раз 6, затем 2 не находится
		},
		{
			name:       "add одного и того же числа много раз подряд",
			operations: []string{"add", "add", "add", "add", "find"},
			values:     []int{2, 2, 2, 2, 4},
			expected:   []bool{true},
		},
		{
			name:       "чередование add и find, влияющих друг на друга",
			operations: []string{"add", "find", "add", "find"},
			values:     []int{1, 2, 1, 2},
			expected:   []bool{false, true}, // до второго add(1) пары 1+1=2 ещё нет
		},
		{
			name:       "большие по модулю числа (граничные значения int)",
			operations: []string{"add", "add", "find"},
			values:     []int{1000000000, 1000000000, 2000000000},
			expected:   []bool{true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ts := Constructor()
			expectedIdx := 0

			for i, op := range tc.operations {
				switch op {
				case "add":
					ts.Add(tc.values[i])
				case "find":
					got := ts.Find(tc.values[i])
					want := tc.expected[expectedIdx]
					if got != want {
						t.Errorf("Find(%d) на шаге %d = %v, ожидалось %v",
							tc.values[i], i, got, want)
					}
					expectedIdx++
				default:
					t.Fatalf("неизвестная операция: %s", op)
				}
			}
		})
	}
}

// Отдельный тест на независимость нескольких экземпляров структуры друг от друга
func TestTwoSum_MultipleInstances(t *testing.T) {
	ts1 := Constructor()
	ts2 := Constructor()

	ts1.Add(1)
	ts1.Add(2)

	ts2.Add(10)
	ts2.Add(20)

	if !ts1.Find(3) {
		t.Errorf("ts1.Find(3) = false, ожидалось true")
	}
	if ts1.Find(30) {
		t.Errorf("ts1.Find(30) = true, ожидалось false — данные из ts2 не должны влиять на ts1")
	}
	if !ts2.Find(30) {
		t.Errorf("ts2.Find(30) = false, ожидалось true")
	}
}

// Тест на порядок вызовов: find до первого add должен корректно вернуть false,
// а не паниковать на nil map
func TestTwoSum_FindBeforeAnyAdd(t *testing.T) {
	ts := Constructor()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Find на пустой структуре вызвал панику: %v", r)
		}
	}()

	if ts.Find(0) {
		t.Errorf("Find(0) на пустой структуре = true, ожидалось false")
	}
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "test1", nums: []int{1, 0, 1, 1, 0}, want: 4},
		{name: "test2", nums: []int{1, 0, 1, 1, 0, 1}, want: 4},
		// single element
		{name: "single one", nums: []int{1}, want: 1},
		{name: "single zero", nums: []int{0}, want: 1},

		// all ones
		{name: "all ones", nums: []int{1, 1, 1, 1}, want: 4},

		// all zeros
		{name: "all zeros", nums: []int{0, 0, 0, 0}, want: 1},

		// zero in the middle
		{name: "bridge two groups", nums: []int{1, 1, 0, 1, 1}, want: 5},

		// zero at the beginning
		{name: "zero at beginning", nums: []int{0, 1, 1, 1}, want: 4},

		// zero at the end
		{name: "zero at end", nums: []int{1, 1, 1, 0}, want: 4},

		// two zeros
		{name: "two zeros", nums: []int{1, 1, 0, 1, 0, 1, 1}, want: 4},

		// long block
		{name: "long block", nums: []int{1, 1, 1, 0, 1, 1, 1}, want: 7},

		// best window not touching edges
		{name: "middle window", nums: []int{0, 1, 1, 0, 1, 1, 1, 0}, want: 6},

		// alternating
		{name: "alternating", nums: []int{1, 0, 1, 0, 1, 0, 1}, want: 3},

		// many zeros
		{name: "many zeros", nums: []int{0, 1, 0, 0, 1, 0}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxConsecutiveOnes(tt.nums)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "test1", s: "eceba", want: 3},
		{name: "example2", s: "ccaabbb", want: 5},

		// one character
		{name: "single char", s: "a", want: 1},

		// all same
		{name: "all same", s: "aaaaa", want: 5},

		// exactly two distinct
		{name: "two distinct", s: "abababab", want: 8},

		// three distinct
		{name: "three distinct", s: "abc", want: 2},

		// third char at the end
		{name: "third at end", s: "aabbc", want: 4},

		// third char at the beginning
		{name: "third at beginning", s: "caabb", want: 4},

		// best window in the middle
		{name: "middle window", s: "abcbbbbcccbdddadacb", want: 10},

		// alternating two chars
		{name: "alternating", s: "ababababab", want: 10},

		// shrink multiple times
		{name: "multiple shrinks", s: "abaccc", want: 4},

		// long suffix
		{name: "long suffix", s: "aabacbebebe", want: 6},

		// all unique
		{name: "all unique", s: "abcdef", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringTwoDistinct(tt.s)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "test1", s: "eceba", k: 2, want: 3},
		{name: "example2", s: "aa", k: 1, want: 2},

		// edge cases
		{name: "empty k", s: "abc", k: 0, want: 0},
		{name: "single char", s: "a", k: 1, want: 1},
		{name: "single char k zero", s: "a", k: 0, want: 0},

		// all same
		{name: "all same", s: "aaaaaa", k: 1, want: 6},
		{name: "all same large k", s: "aaaaaa", k: 5, want: 6},

		// exactly k distinct
		{name: "exactly k distinct", s: "ababab", k: 2, want: 6},

		// all unique
		{name: "all unique k one", s: "abcdef", k: 1, want: 1},
		{name: "all unique k two", s: "abcdef", k: 2, want: 2},
		{name: "all unique k three", s: "abcdef", k: 3, want: 3},

		// k larger than distinct count
		{name: "k larger than distinct", s: "abc", k: 10, want: 3},

		// shrink window
		{name: "shrink once", s: "abac", k: 2, want: 3},
		{name: "shrink many times", s: "abaccc", k: 2, want: 4},

		// optimal window in middle
		{name: "middle window", s: "aabacbebebe", k: 3, want: 7},

		// frequent character changes
		{name: "alternating", s: "abcabcabc", k: 2, want: 2},

		// long valid suffix
		{name: "long suffix", s: "abcbbbbcccbdddadacb", k: 2, want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringKDistinct(tt.s, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{name: "insert middle", intervals: [][]int{{0, 30}, {5, 10}, {15, 20}}, want: 2},
		{
			name:      "leetcode example 2 - non-overlapping",
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      1,
		},
		// === Граничные случаи ===
		{
			name:      "empty input",
			intervals: [][]int{},
			want:      0,
		},
		{
			name:      "single meeting",
			intervals: [][]int{{0, 5}},
			want:      1,
		},
		{
			name:      "all meetings overlap at same point",
			intervals: [][]int{{1, 10}, {1, 10}, {1, 10}, {1, 10}},
			want:      4,
		},
		{
			name:      "back-to-back meetings - same room reused",
			intervals: [][]int{{0, 5}, {5, 10}, {10, 15}, {15, 20}},
			want:      1,
		},
		{
			name:      "start equals end of another - should reuse",
			intervals: [][]int{{0, 5}, {5, 10}},
			want:      1,
		},
		{
			name:      "multiple chains overlapping",
			intervals: [][]int{{0, 10}, {10, 20}, {0, 20}, {5, 15}},
			want:      3,
		},
		{
			name:      "meetings nested inside one long meeting",
			intervals: [][]int{{0, 100}, {10, 20}, {30, 40}, {50, 60}},
			want:      2,
		},
		{
			name:      "unsorted input - heavy overlap",
			intervals: [][]int{{5, 8}, {0, 30}, {3, 19}, {15, 20}, {25, 35}},
			want:      3,
		},
		{
			name:      "all start at same time different durations",
			intervals: [][]int{{1, 3}, {1, 5}, {1, 7}, {1, 9}},
			want:      4,
		},
		{
			name:      "staggered overlaps - max 2 rooms",
			intervals: [][]int{{0, 5}, {2, 7}, {4, 9}, {8, 10}},
			want:      3,
		},
		{
			name:      "large gap then overlap",
			intervals: [][]int{{0, 1}, {100, 101}, {100, 200}, {150, 250}},
			want:      2,
		},
		{
			name:      "many small meetings in one big window",
			intervals: [][]int{{0, 100}, {10, 20}, {20, 30}, {30, 40}, {40, 50}},
			want:      2,
		},
		// 2. Граничный случай: минимально возможный вход (1 интервал)
		// Проверяет, что алгоритм не падает на массиве длины 1 и не выходит за границы при итерации
		{name: "single interval", intervals: [][]int{{1, 5}}, want: 1},

		// 3. КРИТИЧЕСКИЙ граничный случай: касание по времени (конец одного == начало другого)
		// По условию задачи такие интервалы НЕ пересекаются и могут использовать одну комнату.
		// Многие алгоритмы ошибочно считают это перекрытием (off-by-one error).
		{name: "back-to-back meetings", intervals: [][]int{{1, 5}, {5, 10}, {10, 15}}, want: 1},

		// 4. Полное перекрытие (все встречи начинаются в одно время или накладываются друг на друга)
		// Проверяет, что счетчик комнат корректно растет до N
		{name: "fully overlapping", intervals: [][]int{{1, 5}, {2, 6}, {3, 7}}, want: 3},

		// 5. Вложенные интервалы (один длинный, внутри него несколько коротких)
		// Классическая ловушка: алгоритм должен понять, что короткие встречи могут использовать одну и ту же вторую комнату по очереди
		{name: "nested intervals", intervals: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}, want: 2},

		// 6. Обратная сортировка входных данных
		// Проверяет, что твой алгоритм сам корректно сортирует массив (или начала, или концы), а не полагается на порядок ввода
		{name: "reverse sorted input", intervals: [][]int{{15, 20}, {5, 10}, {0, 30}}, want: 2},

		// 7. Несколько интервалов с абсолютно одинаковым временем начала
		// Проверяет корректность обработки дубликатов в массиве `starts` (если используешь Two Pointers)
		{name: "same start time", intervals: [][]int{{1, 5}, {1, 6}, {1, 7}}, want: 3},

		// 8. Несколько интервалов с абсолютно одинаковым временем окончания
		// Проверяет корректность обработки дубликатов в массиве `ends` и порядок освобождения комнат
		{name: "same end time", intervals: [][]int{{1, 5}, {2, 5}, {3, 5}}, want: 3},

		// 9. "Пила" с освобождением комнат (Chain with partial overlap)
		// Проверяет, что алгоритм корректно "освобождает" комнату (сдвигает указатель `endPtr` или делает `pop` из кучи),
		// когда встреча заканчивается, а не накапливает комнаты бесконечно
		{name: "chain with partial overlap", intervals: [][]int{{1, 4}, {2, 5}, {4, 6}, {5, 7}}, want: 2},

		// 10. Идентичные интервалы
		// Проверяет устойчивость к полным дубликатам в массиве
		{name: "identical intervals", intervals: [][]int{{2, 4}, {2, 4}, {2, 4}}, want: 3},

		// 11. Большие значения времени (на границе ограничений LeetCode)
		// Проверяет отсутствие переполнения (хотя в Go int обычно 64-битный, это хорошая практика)
		// и корректность сортировки больших чисел
		{name: "large time values", intervals: [][]int{{0, 1000000}, {500000, 1000000}, {999999, 1000000}}, want: 3},

		// 12. Чередование: много коротких встреч между началом и концом одной длинной
		// Усложненная версия "вложенных интервалов", часто встречается в скрытых тестах LeetCode
		{name: "complex nested sawtooth", intervals: [][]int{{1, 20}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minMeetingRooms(tt.intervals)
			assert.Equal(t, tt.want, res)
		})
	}
}
