package ch10

type LinkedQueue struct {
	head *Node
	tail *Node
	n    int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{n: 0, head: nil, tail: nil}
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *LinkedQueue) Size() int {
	return q.n
}

func (q *LinkedQueue) Enqueue(x int) {
	tmp := q.tail
	q.tail = &Node{data: x, next: nil}
	if q.IsEmpty() {
		q.head = q.tail
	} else {
		tmp.next = q.tail
	}
	q.n++
}

func (q *LinkedQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("underflow")
	}
	x := q.head.data
	q.head = q.head.next
	q.n--
	if q.IsEmpty() {
		q.tail = nil
	}
	return x
}
