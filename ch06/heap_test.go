package ch06

import (
	"fmt"
	"testing"
)

var a = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

func assert(i, j string) {
	if i != j {
		fmt.Printf("%v, %v\n", i, j)
		panic("failed")
	}
}

func TestMaxBuild(t *testing.T) {
	h := MaxBuild(a)
	assert(h.String(), "[16 14 10 8 7 9 3 2 4 1]")
}

func TestMinBuild(t *testing.T) {
	h := MinBuild(a)
	assert(h.String(), "[1 2 3 4 7 9 10 14 8 16]")
}

func TestMaxSort(t *testing.T) {
	b := MaxSort(a)
	assert(fmt.Sprint(b), "[1 2 3 4 7 8 9 10 14 16]")
}

func TestMinSort(t *testing.T) {
	b := MinSort(a)
	assert(fmt.Sprint(b), "[16 14 10 9 8 7 4 3 2 1]")
}

func TestMaxHeap_Modify(t *testing.T) {
	h := MaxBuild(a)
	h.Modify(9, 15)
	assert(h.String(), "[16 15 10 14 7 9 3 2 8 1]")
	h.Modify(2, 13)
	assert(h.String(), "[16 15 10 14 7 9 3 2 8 1]")
}

func TestMaxHeap_Insert(t *testing.T) {
	h := MaxBuild(a)
	h.Modify(9, 15)
	assert(h.String(), "[16 15 10 14 7 9 3 2 8 1]")
	h.Insert(42)
	assert(h.String(), "[42 16 10 14 15 9 3 2 8 1 7]")
}

func TestMinHeap_Merge(t *testing.T) {
	m := []int{1, 3, 5, 7}
	n := []int{2, 4, 6, 8}
	h := MinBuild(nil)
	for _, x := range m {
		h.Insert(x)
	}
	for _, x := range n {
		h.Insert(x)
	}
	for i := h.Size(); i > 0; i-- {
		fmt.Printf("%v", h.Extract())
	}
}
