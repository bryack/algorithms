package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "1 - I", input: 1, want: "I"},
		{name: "2 - II", input: 2, want: "II"},
		{name: "3 - III", input: 3, want: "III"},
		{name: "58 - LVIII", input: 58, want: "LVIII"},
		{name: "63 - LXIII", input: 63, want: "LXIII"},
		{name: "22 - XXII", input: 22, want: "XXII"},
		{name: "33 - XXXIII", input: 33, want: "XXXIII"},
		{name: "44 - XLIV", input: 44, want: "XLIV"},
		{name: "1984 - MCMXCIV", input: 1984, want: "MCMLXXXIV"},
		{name: "1994 - MCMXCIV", input: 1994, want: "MCMXCIV"},
		{name: "3599 - MMMDXCIX", input: 3599, want: "MMMDXCIX"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intToRoman(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "increasing sequence", input: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "example 2", input: []int{7, 1, 5, 3, 6, 4}, want: 7},
		{name: "example 3", input: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "single element", input: []int{5}, want: 0},                  // Нельзя продать
		{name: "two elements - profit", input: []int{1, 5}, want: 4},        // Одна сделка
		{name: "two elements - no profit", input: []int{5, 1}, want: 0},     // Убыток, не продаём
		{name: "all same", input: []int{5, 5, 5, 5}, want: 0},               // Плоская линия
		{name: "strictly decreasing", input: []int{5, 4, 3, 2, 1}, want: 0}, // Только падение
		{name: "strictly increasing", input: []int{1, 2, 3, 4, 5}, want: 4}, // Все сделки
		{name: "zigzag", input: []int{1, 2, 1, 2, 1, 2}, want: 3},           // Множество мелких сделок
		{name: "original example", input: []int{7, 1, 5, 3, 6, 4}, want: 7}, // (5-1) + (6-3) = 7
		{name: "peak at start", input: []int{5, 4, 3, 2, 3, 4, 5}, want: 3}, // Спад, потом рост
		{name: "empty array", input: []int{}, want: 0},                      // Пустой массив

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{name: "example 1 - success", input: []int{5, 5, 5, 10, 20}, want: true},
		{name: "example 2 - fail", input: []int{5, 5, 10, 10, 20}, want: false},
		{name: "single 5", input: []int{5}, want: true},                                // Нет сдачи, OK
		{name: "single 10 - no change", input: []int{10}, want: false},                 // Нужна сдача $5, нет
		{name: "single 20 - no change", input: []int{20}, want: false},                 // Нужна сдача $15, нет
		{name: "only fives", input: []int{5, 5, 5, 5}, want: true},                     // Все платят точно
		{name: "5 then 10 - success", input: []int{5, 10}, want: true},                 // Есть $5 для сдачи
		{name: "alternating 5 and 10", input: []int{5, 10, 5, 10}, want: true},         // Хватает пятёрок
		{name: "greedy choice matters", input: []int{5, 5, 10, 20}, want: true},        // 20 → 10+5 (правильно)
		{name: "greedy trap - three 5s better", input: []int{5, 5, 5, 20}, want: true}, // 20 → 5+5+5 (тоже ок)
		{name: "not enough 5s for 20", input: []int{5, 5, 20}, want: false},            // Нужно $15, есть только две $5
		{name: "20 when only 10s left", input: []int{5, 10, 20, 20}, want: false},      // После второго 20 не хватит
		{name: "empty", input: []int{}, want: true},                                    // Никого нет — успех
		{name: "many 20s fail", input: []int{5, 5, 5, 5, 20, 20, 20}, want: false},     // Три 20 потребуют много сдачи
		{name: "10 without prior 5", input: []int{10, 5}, want: false},                 // Порядок важен!

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lemonadeChange(tt.input)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		name  string
		greed []int
		size  []int
		want  int
	}{
		{name: "example 1", greed: []int{1, 2, 3}, size: []int{1, 1}, want: 1},
		{name: "example 2", greed: []int{1, 2}, size: []int{1, 2, 3}, want: 2},
		{name: "cookie larger than greed", greed: []int{1, 2}, size: []int{2, 2}, want: 2},
		// Твой код: 0 (нет точных совпадений), но правильный ответ: 2 (оба ребёнка могут получить печеньку 2)
		{name: "need greedy choice", greed: []int{1, 2, 3}, size: []int{1, 2}, want: 2},
		// Твой код: 2 (ok), но если бы проверял в другом порядке — мог бы сломаться
		{name: "unsorted arrays", greed: []int{3, 1, 2}, size: []int{1, 2}, want: 2},
		// Твой код: зависит от порядка, может дать неправильный ответ
		{name: "no cookies", greed: []int{1, 2, 3}, size: []int{}, want: 0},
		// Твой код: 0 (ok)
		{name: "all cookies too small", greed: []int{5, 6, 7}, size: []int{1, 2, 3}, want: 0},
		// Твой код: 0 (ok)
		{name: "one big cookie for all", greed: []int{1, 2, 3}, size: []int{10}, want: 1},
		// Твой код: 0 (нет точного совпадения 10), правильный: 1
		{name: "optimal assignment", greed: []int{1, 2, 3}, size: []int{1, 2, 3}, want: 3},
		// Твой код: 3 (ok, точные совпадения)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findContentChildren(tt.greed, tt.size)
			assert.Equal(t, tt.want, result)
		})
	}
}
