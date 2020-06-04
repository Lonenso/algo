package ch10

import (
	"testing"
)

func TestLinkedStackSuite(t *testing.T) {
	q := NewLinkedStack()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	for q.IsEmpty() == false {
		println(q.Pop())
	}
}