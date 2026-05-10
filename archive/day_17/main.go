package main

type RingQueue struct {
	items []int
	head  int
	tail  int
	size  int
}

func (q *RingQueue) Enqueue(val int) bool {
	if q.IsFull() {
		return false
	}
	q.items[q.tail] = val
	q.tail = (q.tail + 1) % len(q.items)
	q.size++
	return true
}

func (q *RingQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	val := q.items[q.head]
	q.items[q.head] = 0
	q.head = (q.head + 1) % len(q.items)
	q.size--

	return val, true
}

func (q *RingQueue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[q.head], true
}

func (q *RingQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *RingQueue) IsFull() bool {
	return q.size == len(q.items)

}
