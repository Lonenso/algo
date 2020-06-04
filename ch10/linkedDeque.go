package ch10

type LinkedDeque struct {
	head *Node
	tail *Node
	n    int
}

func NewLinkedDeque() *LinkedDeque {
	return &LinkedDeque{
		head: nil,
		tail: nil,
		n:    0,
	}
}

func (d *LinkedDeque) AddFirst(x int) {
	node := &Node{data: x, next: d.head}
	d.head = node
	if d.tail == nil {
		d.tail = d.head
	}
	d.n++
}

func (d *LinkedDeque) AddLast(x int) {
	node := &Node{data: x, next: nil}
	if d.tail == nil {
		d.tail = node
	} else {
		d.tail.next = node
		d.tail = node
	}
	if d.head == nil {
		d.head = d.tail
	}
	d.n++
}

func (d *LinkedDeque) RemoveFirst() int {
	if d.IsEmpty() {
		panic("underflow")
	}
	x := d.head.data
	d.head = d.head.next
	d.n--
	return x
}

func (d *LinkedDeque) RemoveLast() int {
	if d.IsEmpty(){
		panic("underflow")
	}
	x := d.tail.data
	d.
}

func (d *LinkedDeque) Size() int {
	return d.n
}

func (d *LinkedDeque) IsEmpty() bool {
	return d.n == 0
}
