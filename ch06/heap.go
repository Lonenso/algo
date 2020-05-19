package ch06

type Heap interface {
	Heapify(elem int)
	Size() int
	Insert(int)
	Extract() int
	Modify(elem int, value int)
	Top() int
	String() string
	Pop(int)
}

func left(elem int) int {
	return elem << 1
}

func right(elem int) int {
	return elem<<1 + 1
}

func parent(elem int) int {
	return elem >> 1
}

func MaxBuild(ints []int) Heap {
	data := make([]int, len(ints)+1)
	for i := 0; i < len(ints); i++ {
		data[i+1] = ints[i]
	}
	h := MaxHeap{data: data, size: len(ints)}

	for i := h.Size() >> 1; i > 0; i-- {
		h.Heapify(i)
	}
	return &h
}

func MaxSort(ints []int) []int {
	h := MaxBuild(ints)
	for i := h.Size(); i > 1; i-- {
		h.(*MaxHeap).data[1], h.(*MaxHeap).data[i] = h.(*MaxHeap).data[i], h.(*MaxHeap).data[1]
		h.(*MaxHeap).size--
		h.Heapify(1)
	}
	return h.(*MaxHeap).data[1:]
}

func MinBuild(ints []int) Heap {
	data := make([]int, len(ints)+1)
	for i := 0; i < len(ints); i++ {
		data[i+1] = ints[i]
	}
	h := MinHeap{data: data, size: len(ints)}

	for i := h.Size() >> 1; i > 0; i-- {
		h.Heapify(i)
	}
	return &h
}

func MinSort(ints []int) []int {
	h := MinBuild(ints)
	for i := h.Size(); i > 1; i-- {
		h.(*MinHeap).data[1], h.(*MinHeap).data[i] = h.(*MinHeap).data[i], h.(*MinHeap).data[1]
		h.(*MinHeap).size--
		h.Heapify(1)
	}
	return h.(*MinHeap).data[1:]
}
