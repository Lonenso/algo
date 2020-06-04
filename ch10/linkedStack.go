package ch10

type Node struct {
	data int
	next *Node
}

type LinkedStack struct {
	top *Node
	n   int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		top: nil,
		n:   0,
	}
}

func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}

func (s *LinkedStack) Size() int {
	return s.n
}

func (s *LinkedStack) Push(x int) {
	s.top = &Node{data: x, next: s.top}
	s.n++
}

func (s *LinkedStack) Pop() int {
	if s.IsEmpty() {
		panic("underflow")
	}
	n := s.top.data
	s.top = s.top.next
	s.n--
	return n
}
