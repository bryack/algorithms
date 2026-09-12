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

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected []string
	}{
		{
			name:     "базовый случай из условия LeetCode",
			words:    []string{"bella", "label", "roller"},
			expected: []string{"e", "l", "l"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := commonChars(tc.words)
			assert.Equal(t, tc.expected, res)
		})
	}
}
