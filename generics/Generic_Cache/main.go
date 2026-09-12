package main

type Cache[T any] struct {
	data map[string]T
}

func NewCache[T any]() *Cache[T] {
	return &Cache[T]{data: make(map[string]T)}
}

func (c *Cache[T]) Put(key string, value T) {
	c.data[key] = value
}

func (c *Cache[T]) Get(key string) (T, bool) {
	if val, ok := c.data[key]; ok {
		return val, true
	}
	var zero T
	return zero, false
}
