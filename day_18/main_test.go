package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeue_PushBack(t *testing.T) {
	tests := []struct {
		name        string
		dequeue     *Dequeue
		val         int
		wantDequeue *Dequeue
		wantOk      bool
	}{
		{
			name: "test 1",
			dequeue: &Dequeue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			val: 4,
			wantDequeue: &Dequeue{
				items: []int{1, 2, 3, 4},
				head:  0,
				tail:  0,
				size:  4,
			},
			wantOk: true,
		},
		{
			name: "empty dequeue",
			dequeue: &Dequeue{
				items: make([]int, 3),
				head:  0,
				tail:  0,
				size:  0,
			},
			val: 4,
			wantDequeue: &Dequeue{
				items: []int{4, 0, 0},
				head:  0,
				tail:  1,
				size:  1,
			},
			wantOk: true,
		},
		{
			name: "full dequeue",
			dequeue: &Dequeue{
				items: []int{1, 2, 3},
				head:  0,
				tail:  0,
				size:  3,
			},
			val: 4,
			wantDequeue: &Dequeue{
				items: []int{1, 2, 3},
				head:  0,
				tail:  0,
				size:  3,
			},
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.dequeue
			ok := d.PushBack(tt.val)
			assert.Equal(t, tt.wantDequeue, d)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}

func TestDequeue_PushFront(t *testing.T) {
	tests := []struct {
		name        string
		dequeue     *Dequeue
		val         int
		wantDequeue *Dequeue
		wantOk      bool
	}{
		{
			name: "test 1",
			dequeue: &Dequeue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			val: 4,
			wantDequeue: &Dequeue{
				items: []int{1, 2, 3, 4},
				head:  3,
				tail:  3,
				size:  4,
			},
			wantOk: true,
		},
		{
			name: "test 2",
			dequeue: &Dequeue{
				items: []int{1, 2, 3, 0, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			val: 4,
			wantDequeue: &Dequeue{
				items: []int{1, 2, 3, 0, 4},
				head:  4,
				tail:  3,
				size:  4,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.dequeue
			ok := d.PushFront(tt.val)
			assert.Equal(t, tt.wantDequeue, d)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}

func TestDequeue_PopFront(t *testing.T) {
	tests := []struct {
		name        string
		dequeue     *Dequeue
		wantVal     int
		wantDequeue *Dequeue
		wantOk      bool
	}{
		{
			name: "test 1",
			dequeue: &Dequeue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			wantVal: 1,
			wantDequeue: &Dequeue{
				items: []int{0, 2, 3, 0},
				head:  1,
				tail:  3,
				size:  2,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.dequeue
			res, ok := d.PopFront()

			assert.Equal(t, tt.wantVal, res)
			assert.Equal(t, tt.wantDequeue, d)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}

func TestDequeue_PopBack(t *testing.T) {
	tests := []struct {
		name        string
		dequeue     *Dequeue
		wantVal     int
		wantDequeue *Dequeue
		wantOk      bool
	}{
		{
			name: "test 1",
			dequeue: &Dequeue{
				items: []int{1, 2, 3, 0},
				head:  0,
				tail:  3,
				size:  3,
			},
			wantVal: 3,
			wantDequeue: &Dequeue{
				items: []int{1, 2, 0, 0},
				head:  0,
				tail:  2,
				size:  2,
			},
			wantOk: true,
		},
		{
			name: "tail = 0",
			dequeue: &Dequeue{
				items: []int{1, 2, 3, 4},
				head:  0,
				tail:  0,
				size:  4,
			},
			wantVal: 4,
			wantDequeue: &Dequeue{
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
			d := tt.dequeue
			res, ok := d.PopBack()

			assert.Equal(t, tt.wantVal, res)
			assert.Equal(t, tt.wantDequeue, d)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}
