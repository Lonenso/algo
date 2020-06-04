package ch10

type Deque struct {
	data []int
	head int
	tail int
	n    int
}

func NewDeque() *Deque {
	return &Deque{
		data: make([]int, 80),
		head: 79,
		tail: 1,
		n:    0,
	}
}

func (d *Deque) AddHead(x int) {
	if d.IsFull() {
		panic("overflow")
	}
	if d.head == 1 {
		d.head = len(d.data)
	} else {
		d.head--
	}
	d.data[d.head] = x
}

func (d *Deque) AddTail(x int) {
	if d.IsFull() {
		panic("overflow")
	}
	d.data[d.tail] = x
	if d.tail == len(d.data) {
		d.tail = 1
	} else {
		d.tail++
	}
}

func (d *Deque) DelHead() int {
	if d.IsEmpty() {
		panic("underflow")
	}
	x := d.data[d.head]
	if d.head == len(d.data) - 1 {
		d.head = 1
	} else {
		d.head++
	}
	return x
}

func (d *Deque) DelTail() int {
	if d.IsEmpty() {
		panic("underflow")
	}
	if d.tail == 1 {
		d.tail = len(d.data) - 1
	} else {
		d.tail--
	}
	x := d.data[d.tail]
	return x
}

func (d *Deque) IsFull() bool {
	if d.head == d.tail+1 || (d.head == 1 && d.tail == len(d.data)) {
		return true
	}
	return false
}

func (d *Deque) IsEmpty() bool {
	if d.head == d.tail {
		return true
	} else {
		return false
	}
}
