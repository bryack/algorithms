package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestСompress(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  int
	}{
		{name: "classic", chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, want: 6},
		{name: "test2", chars: []byte{'a'}, want: 1},
		{name: "test3", chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, want: 4},
		{name: "no_repeat", chars: []byte{'a', 'b', 'c'}, want: 3},
		{name: "count_9", chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, want: 2},
		{name: "count_10", chars: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, want: 3},
		{name: "mixed", chars: []byte{'a', 'a', 'b', 'b', 'a'}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := compress(tt.chars)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestIsOneEditDistance(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "insert middle", s: "ab", t: "acb", want: true},
		{name: "insert end", s: "ab", t: "abc", want: true},
		{name: "insert start", s: "ab", t: "cab", want: true},
		{name: "delete middle", s: "acb", t: "ab", want: true},
		{name: "delete end", s: "abc", t: "ab", want: true},
		{name: "delete start", s: "cab", t: "ab", want: true},
		{name: "replace", s: "ab", t: "ac", want: true},
		{name: "replace single", s: "1203", t: "1213", want: true},
		{name: "empty to single", s: "", t: "A", want: true},
		{name: "single to empty", s: "A", t: "", want: true},
		{name: "identical", s: "ab", t: "ab", want: false},
		{name: "both empty", s: "", t: "", want: false},
		{name: "diff > 1 length", s: "ab", t: "abcb", want: false},
		{name: "diff > 1 length reverse", s: "abcb", t: "ab", want: false},
		{name: "two replacements", s: "ab", t: "cd", want: false},
		{name: "two inserts needed", s: "a", t: "abc", want: false},
		{name: "two deletes needed", s: "abc", t: "a", want: false},
		{name: "single replace char", s: "a", t: "b", want: true},
		{name: "single replace case", s: "a", t: "A", want: true},
		{name: "transposition", s: "ab", t: "ba", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isOneEditDistance(tt.s, tt.t)
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
