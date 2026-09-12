package main

type Dequeue struct {
	items []int
	head  int
	tail  int
	size  int
}

func (d *Dequeue) IsEmpty() bool {
	return d.size == 0
}

func (d *Dequeue) Front() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	return d.items[d.head], true
}

func (d *Dequeue) Back() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.items[(d.tail-1+len(d.items))%len(d.items)], true
}

func (d *Dequeue) PushBack(val int) bool {
	if d.size == len(d.items) {
		return false
	}
	d.items[d.tail] = val
	d.tail = (d.tail + 1) % len(d.items)
	d.size++
	return true
}

func (d *Dequeue) PushFront(val int) bool {
	if d.size == len(d.items) {
		return false
	}

	d.head = (d.head - 1 + len(d.items)) % len(d.items)
	d.items[d.head] = val

	d.size++
	return true
}

func (d *Dequeue) PopFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	res := d.items[d.head]
	d.items[d.head] = 0
	d.head = (d.head + 1) % len(d.items)
	d.size--
	return res, true
}

func (d *Dequeue) PopBack() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	tailIndex := (d.tail - 1 + len(d.items)) % len(d.items)
	res := d.items[tailIndex]
	d.items[tailIndex] = 0
	d.tail = tailIndex
	d.size--
	return res, true
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	arr := make([]rune, 0, len(s)/2)
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, r := range s {
		_, ok := pairs[r]
		if !ok {
			arr = append(arr, r)
		} else {
			if len(arr) == 0 {
				return false
			}
			if arr[len(arr)-1] != pairs[r] {
				return false
			}
			arr = arr[:len(arr)-1]
		}
	}

	return len(arr) == 0
}
