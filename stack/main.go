package main

/*
Operações principais:
- push → empilha (insere)
- pop → desempilha (remove)
- peek → vê o topo
- isEmpty
- size
*/

type Node struct {
	value int
	next  *Node
}

type StackLinked struct {
	top  *Node
	size int
}

func NewStackLinked() *StackLinked {
	return &StackLinked{
		top:  nil,
		size: 0,
	}
}

// Push
func (s *StackLinked) Push(value int) {
	newNode := &Node{value: value}

	newNode.next = s.top
	s.top = newNode

	s.size++
}

// Pop
func (s *StackLinked) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	value := s.top.value
	s.top = s.top.next
	s.size--

	return value
}

// Peek
func (s *StackLinked) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.top.value
}

// IsEmpty
func (s *StackLinked) IsEmpty() bool {
	return s.size == 0
}

// Size
func (s *StackLinked) Size() int {
	return s.size
}

// Teste
func main() {
	s := NewStackLinked()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	println(s.Pop())  // 30
	println(s.Peek()) // 20
}
