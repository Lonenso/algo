package ch10

import (
	"testing"
)

func TestDequeSuite(t *testing.T) {
	d := NewDeque()
	d.AddHead(1)
	d.DelTail()
	d.AddTail(1)
	d.DelHead()
	d.AddHead(1)
	d.AddTail(1)
	d.DelTail()
}