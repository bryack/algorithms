package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry[T any] struct {
	key string
	val T
}

func TestCache(t *testing.T) {
	t.Run("int cache", func(t *testing.T) {
		tests := []struct {
			name   string
			putSlc []entry[int]
			want   []int
		}{
			{name: "1_test", putSlc: []entry[int]{{"1", 1}, {"2", 2}, {"3", 3}}, want: []int{1, 2, 3}},
			{name: "2_test", putSlc: []entry[int]{{"put", -400}, {"utp", 0}, {"tup", 30}}, want: []int{-400, 0, 30}},
			{name: "empty", putSlc: []entry[int]{}, want: []int{}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cache := NewCache[int]()
				for _, p := range tt.putSlc {
					cache.Put(p.key, p.val)
				}
				get := make([]int, 0, len(cache.data))
				for key := range cache.data {
					k, _ := cache.Get(key)
					get = append(get, k)
				}
				assert.ElementsMatch(t, tt.want, get)
			})
		}
	})

	t.Run("float64 cache", func(t *testing.T) {
		tests := []struct {
			name   string
			putSlc []entry[float64]
			want   []float64
		}{
			{name: "1_test", putSlc: []entry[float64]{{"1", 1.5}, {"2", 2.1}, {"3", 3.3}}, want: []float64{1.5, 2.1, 3.3}},
			{name: "2_test", putSlc: []entry[float64]{{"put", -400}, {"utp", 0}, {"tup", 30}}, want: []float64{-400, 0, 30}},
			{name: "empty", putSlc: []entry[float64]{}, want: []float64{}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cache := NewCache[float64]()
				for _, p := range tt.putSlc {
					cache.Put(p.key, p.val)
				}
				get := make([]float64, 0, len(cache.data))
				for key := range cache.data {
					k, _ := cache.Get(key)
					get = append(get, k)
				}
				assert.ElementsMatch(t, tt.want, get)
			})
		}
	})

	t.Run("rewrite string value", func(t *testing.T) {
		cache := NewCache[string]()
		cache.Put("rewrite", "first")
		cache.Put("rewrite", "second")
		get, ok := cache.Get("rewrite")
		assert.True(t, ok)
		assert.Equal(t, "second", get)
	})
	t.Run("not exist", func(t *testing.T) {
		cache := NewCache[string]()
		cache.Put("put", "first")
		cache.Put("utp", "second")
		get, ok := cache.Get("not")
		assert.False(t, ok)
		assert.Empty(t, get)
	})
}
