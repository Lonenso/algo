package ch06

import (
	"fmt"
)

type MaxHeap struct {
	data []int
	size int
}

// 如果元素的值较大, 得上浮(Modify), 较小得下沉(Heapify)
func (mh *MaxHeap) Pop(i int) {
	if mh.data[i] < mh.data[mh.size] {
		mh.data[i] = mh.data[mh.size]
		mh.Heapify(i)
	} else {
		mh.Modify(i, mh.data[mh.size])
	}
	mh.size--
}

func (mh *MaxHeap) Heapify(elem int) {
	l := left(elem)
	r := right(elem)
	var largest = elem

	if l <= mh.Size() && mh.data[l] > mh.data[elem] {
		largest = l
	}

	if r <= mh.Size() && mh.data[r] > mh.data[largest] {
		largest = r
	}
	if largest != elem {
		mh.data[elem], mh.data[largest] = mh.data[largest], mh.data[elem]
		mh.Heapify(largest)
	}
}

func (mh *MaxHeap) Size() int {
	return mh.size
}

func (mh *MaxHeap) Insert(i int) {
	mh.size++
	if len(mh.data)-1 < mh.size {
		mh.data = append(mh.data, -2147483648)
	} else {
		mh.data[mh.size] = -2147483648
	}
	mh.Modify(mh.size, i)
}

func (mh *MaxHeap) Extract() int {
	if mh.Size() < 1 {
		panic("heap underflow")
	}
	top := mh.data[1]
	mh.data[1] = mh.data[mh.size]
	mh.size--
	mh.Heapify(1)
	return top
}

func (mh *MaxHeap) Modify(elem int, value int) {
	if elem == 0 {
		panic("start with 1")
	}
	if value < mh.data[elem] {
		return
	}
	for elem > 1 && mh.data[parent(elem)] < value {
		mh.data[elem] = mh.data[parent(elem)]
		elem = parent(elem)
	}
	mh.data[elem] = value
}

func (mh *MaxHeap) Top() int {
	return mh.data[1]
}

func (mh *MaxHeap) String() string {
	return fmt.Sprintf("%v", mh.data[1:mh.size+1])
}
