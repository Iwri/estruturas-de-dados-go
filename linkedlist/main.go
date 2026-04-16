package main

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// Inserções
func (list *LinkedList) Add(value int) { // ADD final
	newNode := &Node{value: value}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
}

func (list *LinkedList) AddFirst(value int) { // ADD início
	newNode := &Node{value: value}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}

	list.size++
}

func (list *LinkedList) AddAt(index int, value int) { // ADD posição específica
	if index < 0 || index > list.size {
		panic("index out of bounds")
	}

	if index == 0 {
		list.AddFirst(value)
		return
	}

	if index == list.size {
		list.Add(value)
		return
	}

	newNode := &Node{value: value}

	current := list.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode

	list.size++
}

// GET
func (list *LinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}

	current := list.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

// REMOVE
func (list *LinkedList) RemoveAt(index int) int {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}

	var removed *Node

	if index == 0 {
		removed = list.head
		list.head = list.head.next

		if list.head == nil {
			list.tail = nil
		}

	} else {
		current := list.head

		for i := 0; i < index-1; i++ {
			current = current.next
		}

		removed = current.next
		current.next = removed.next

		if removed == list.tail {
			list.tail = current
		}
	}

	list.size--
	return removed.value
}

// FIND
func (list *LinkedList) Find(value int) *Node {
	current := list.head

	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}

	return nil
}

// PRINT
func (list *LinkedList) Print() {
	current := list.head

	for current != nil {
		println(current.value)
		current = current.next
	}
}

// TESTE
func main() {
	list := NewLinkedList()

	list.Add(10)
	list.Add(20)
	list.Add(30)

	list.AddFirst(5)
	list.AddAt(2, 99)

	println("Lista:")
	list.Print()

	println("Elemento índice 2:", list.Get(2))

	list.RemoveAt(2)

	println("Após remoção:")
	list.Print()
}
