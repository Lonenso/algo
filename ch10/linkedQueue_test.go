package ch10

import (
	"testing"
)

func TestLinkedQueueSuite(t *testing.T) {
	q := NewLinkedQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	for q.IsEmpty() == false {
		println(q.Dequeue())
	}
}
