package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name      string
		newVal    int
		items     []int
		head      int
		tail      int
		size      int
		wantQueue *RingQueue
		wantOk    bool
	}{
		{
			name:   "push element into free cell",
			newVal: 4,
			items:  []int{1, 2, 3, 0},
			head:   0,
			tail:   3,
			size:   3,
			wantQueue: &RingQueue{
				items: []int{1, 2, 3, 4},
				head:  0,
				tail:  0,
				size:  4,
			},
			wantOk: true,
		},
		{
			name:   "buffer overflow - full buffer",
			newVal: 5,
			items:  []int{1, 2, 3, 4},
			head:   0,
			tail:   0,
			size:   4,
			wantQueue: &RingQueue{
				items: []int{1, 2, 3, 4},
				head:  0,
				tail:  0,
				size:  4,
			},
			wantOk: false,
		},
		{
			name:   "tail wrap around",
			newVal: 50,
			items:  []int{0, 20, 30, 40},
			head:   1,
			tail:   0,
			size:   3,
			wantQueue: &RingQueue{
				items: []int{50, 20, 30, 40},
				head:  1,
				tail:  1,
				size:  4,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RingQueue{
				items: tt.items,
				head:  tt.head,
				tail:  tt.tail,
				size:  tt.size,
			}

			ok := q.Enqueue(tt.newVal)

			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantQueue, q)
		})
	}
}

func TestRingQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		head      int
		tail      int
		size      int
		want      int
		wantQueue *RingQueue
		wantOk    bool
	}{
		{
			name:  "pop element",
			items: []int{1, 2, 3, 0},
			head:  0,
			tail:  3,
			size:  3,
			want:  1,
			wantQueue: &RingQueue{
				items: []int{0, 2, 3, 0},
				head:  1,
				tail:  3,
				size:  2,
			},
			wantOk: true,
		},
		{
			name:  "dequeue from empty queue",
			items: []int{0, 0, 0, 0},
			head:  0,
			tail:  0,
			size:  0,
			want:  0,
			wantQueue: &RingQueue{
				items: []int{0, 0, 0, 0},
				head:  0,
				tail:  0,
				size:  0,
			},
			wantOk: false,
		},
		{
			name:  "dequeue last element",
			items: []int{0, 0, 0, 42},
			head:  3,
			tail:  0,
			size:  1,
			want:  42,
			wantQueue: &RingQueue{
				items: []int{0, 0, 0, 0},
				head:  0,
				tail:  0,
				size:  0,
			},
			wantOk: true,
		},
		{
			name:  "head wrap around",
			items: []int{10, 20, 30, 40},
			head:  3,
			tail:  1,
			size:  2,
			want:  40,
			wantQueue: &RingQueue{
				items: []int{10, 20, 30, 0},
				head:  0,
				tail:  1,
				size:  1,
			},
			wantOk: true,
		},
		{
			name:  "dequeue after wrap around",
			items: []int{50, 20, 30, 40},
			head:  1,
			tail:  1,
			size:  4,
			want:  20,
			wantQueue: &RingQueue{
				items: []int{50, 0, 30, 40},
				head:  2,
				tail:  1,
				size:  3,
			},
			wantOk: true,
		},
		{
			name:  "sequential dequeue",
			items: []int{1, 2, 3, 4},
			head:  0,
			tail:  0,
			size:  4,
			want:  1,
			wantQueue: &RingQueue{
				items: []int{0, 2, 3, 4},
				head:  1,
				tail:  0,
				size:  3,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RingQueue{
				items: tt.items,
				head:  tt.head,
				tail:  tt.tail,
				size:  tt.size,
			}

			res, ok := q.Dequeue()

			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantQueue, q)
		})
	}
}

func TestRingQueue_Front(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		head      int
		tail      int
		size      int
		want      int
		wantQueue *RingQueue
		wantOk    bool
	}{
		{
			name:  "test 1",
			items: []int{1, 2, 3, 0},
			head:  0,
			tail:  3,
			size:  3,
			want:  1,
			wantQueue: &RingQueue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RingQueue{
				items: tt.items,
				head:  tt.head,
				tail:  tt.tail,
				size:  tt.size,
			}

			res, ok := q.Front()

			assert.Equal(t, tt.want, res)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantQueue, q)
		})
	}
}

func TestRingQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		head      int
		tail      int
		size      int
		wantQueue *RingQueue
		wantOk    bool
	}{
		{
			name:  "not empty",
			items: []int{1, 2, 3, 0},
			head:  0,
			tail:  3,
			size:  3,
			wantQueue: &RingQueue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			wantOk: false,
		},
		{
			name:  "empty",
			items: []int{},
			head:  0,
			tail:  0,
			size:  0,
			wantQueue: &RingQueue{
				items: []int{},
				head:  0,
				tail:  0,
				size:  0,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RingQueue{
				items: tt.items,
				head:  tt.head,
				tail:  tt.tail,
				size:  tt.size,
			}

			ok := q.IsEmpty()

			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantQueue, q)
		})
	}
}

func TestRingQueue_IsFull(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		head      int
		tail      int
		size      int
		wantQueue *RingQueue
		wantOk    bool
	}{
		{
			name:  "not full",
			items: []int{1, 2, 3, 0},
			head:  0,
			tail:  3,
			size:  3,
			wantQueue: &RingQueue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			wantOk: false,
		},
		{
			name:  "full",
			items: []int{1, 2},
			head:  0,
			tail:  0,
			size:  2,
			wantQueue: &RingQueue{
				items: []int{1, 2},
				head:  0,
				tail:  0,
				size:  2,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RingQueue{
				items: tt.items,
				head:  tt.head,
				tail:  tt.tail,
				size:  tt.size,
			}

			ok := q.IsFull()

			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.wantQueue, q)
		})
	}

}
