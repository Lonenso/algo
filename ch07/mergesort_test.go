package ch07

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	a := generate(20)
	mergesort(a, 0, len(a)-1)
	fmt.Printf("%v\n", a)
}

func TestMergeLinkedList(t *testing.T) {
	a := generateLinkedList(20)
	c := mergesortLinkedList(a)
	c.Output()
}

func generateLinkedList(N int) *Node {
	head := &Node{}
	res := head
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		res.next = &Node{val: rand.Intn(100)}
		res = res.next
	}
	return head.next
}
