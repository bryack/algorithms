package asm

import (
    "fmt"
    "testing"
)

func TestSum(t *testing.T) {
    tests := []struct {
        name string
        s    []int32
        want int64
    }{
        {"empty", []int32{}, 0},
        {"single", []int32{5}, 5},
        {"multiple", []int32{1, 2, 3, 4}, 10},
        {"negative", []int32{-1, -2, 3}, 0},
        {"large", []int32{1000000, 2000000}, 3000000},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Sum(tt.s)
            if got != tt.want {
                t.Errorf("Sum(%v) = %d, want %d", tt.s, got, tt.want)
            }
        })
    }
}

func TestFibonacci(t *testing.T) {
    tests := []struct {
        n    uint64
        want uint64
    }{
        {0, 0},
        {1, 1},
        {2, 1},
        {3, 2},
        {4, 3},
        {5, 5},
        {10, 55},
    }
    for _, tt := range tests {
        t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
            got := Fibonacci(tt.n)
            if got != tt.want {
                t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
            }
        })
    }
}
