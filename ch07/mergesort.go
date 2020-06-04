package ch07

import "fmt"

func mergesort(a []int, p int, r int) {
	if p < r {
		q := p + (r-p)/2
		mergesort(a, p, q)
		mergesort(a, q+1, r)
		merge(a, p, q, r)
	}
}

func merge(a []int, p int, q int, r int) {
	left := make([]int, q-p+1)
	right := make([]int, r-q)
	// fmt.Printf("p: %v, q: %v, r: %v\n", p, q, r)
	for i := 0; i < len(left); i++ {
		left[i] = a[i+p]
	}
	for i := 0; i < len(right); i++ {
		right[i] = a[q+1+i]
	}
	for i, j, k := 0, 0, p; k < r+1; k++ {
		// fmt.Printf("i: %v, j: %v, k:%v\n", i, j, k)
		if !(i < q-p+1) {
			a[k] = right[j]
			j++
		} else if !(j < r-q) {
			a[k] = left[i]
			i++
		} else if left[i] < right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
	}
}

type Node struct {
	val  int
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.val)
}

func (n *Node) Output() {
	t := n
	for t != nil {
		fmt.Printf("%v\n", t)
		t = t.next
	}
}

func mergesortLinkedList(a *Node) *Node {
	if a == nil || a.next == nil {
		return a
	}
	slow := a
	fast := a.next.next
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	left := a
	right := slow.next
	slow.next = nil
	left = mergesortLinkedList(left)
	right = mergesortLinkedList(right)
	return mergeLinkedList(left, right)
}

func mergeLinkedList(left *Node, right *Node) *Node {
	head := &Node{}
	res := head
	for left != nil && right != nil {
		if left.val < right.val {
			res.next = left
			left = left.next
		} else {
			res.next = right
			right = right.next
		}
		res = res.next
	}
	if left != nil {
		res.next = left
	}
	if right != nil {
		res.next = right
	}
	return head.next
}
