package ch06

import (
	"fmt"
)

type MinHeap struct {
	data []int
	size int
}

func (mh *MinHeap) Pop(i int) {
	if mh.data[i] > mh.data[mh.size] {
		mh.data[i] = mh.data[mh.size]
		mh.Heapify(i)
	} else {
		mh.Modify(i, mh.data[mh.size])
	}
	mh.size--
}

func (mh *MinHeap) Heapify(elem int) {
	l := left(elem)
	r := right(elem)
	var smallest = elem

	if l <= mh.size && mh.data[l] < mh.data[elem] {
		smallest = l
	}
	if r <= mh.size && mh.data[r] < mh.data[smallest] {
		smallest = r
	}
	if smallest != elem {
		mh.data[elem], mh.data[smallest] = mh.data[smallest], mh.data[elem]
		mh.Heapify(smallest)
	}
}

func (mh *MinHeap) Size() int {
	return mh.size
}

func (mh *MinHeap) Insert(i int) {
	mh.size++
	if len(mh.data)-1 < mh.size {
		mh.data = append(mh.data, 2147483647)
	} else {
		mh.data[mh.size] = 2147483647
	}
	mh.Modify(mh.size, i)
}

func (mh *MinHeap) Extract() int {
	if mh.Size() < 1 {
		panic("heap underflow")
	}
	top := mh.data[1]
	mh.data[1] = mh.data[mh.size]
	mh.size--
	mh.Heapify(1)
	return top
}

func (mh *MinHeap) Modify(elem int, value int) {
	if elem == 0 {
		panic("start with 1")
	}
	if value > mh.data[elem] {
		return
	}
	for elem > 1 && mh.data[parent(elem)] > value {
		mh.data[elem] = mh.data[parent(elem)]
		elem = parent(elem)
	}
	mh.data[elem] = value
}

func (mh *MinHeap) Top() int {
	return mh.data[1]
}

func (mh *MinHeap) String() string {
	return fmt.Sprintf("%v", mh.data[1:mh.size+1])
}
