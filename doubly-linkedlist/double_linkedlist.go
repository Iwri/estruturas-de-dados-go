package main

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Funções Básicas
func (list *DoubleLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *DoubleLinkedList) Size() int {
	return list.size
}

// Inserções

func (list *DoubleLinkedList) Add(value int) { // ADD final
	newNode := &Node{value: value}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
}

func (list *DoubleLinkedList) AddFirst(value int) { // ADD início
	newNode := &Node{value: value}

	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}

	list.size++
}

func (list *DoubleLinkedList) AddAt(index int, value int) { // ADD em uma posição específica
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
	for i := 0; i < index; i++ {
		current = current.next
	}

	prevNode := current.prev

	newNode.prev = prevNode
	newNode.next = current

	prevNode.next = newNode
	current.prev = newNode

	list.size++
}

// Acesso
func (list *DoubleLinkedList) Get(index int) int { // Versão otimizada
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}

	var current *Node

	//Decide de onde começar
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}

	}

	return current.value
}

// Remoção
func (list *DoubleLinkedList) RemoveAt(index int) int {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}

	var current *Node

	if index == 0 {
		current = list.head
		list.head = current.next

		if list.head != nil {
			list.head.prev = nil
		} else {
			list.tail = nil
		}

	} else if index == list.size-1 {
		current = list.tail
		list.tail = current.prev

		if list.tail != nil {
			list.tail.next = nil
		} else {
			list.head = nil
		}

	} else {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}

		current.prev.next = current.next
		current.next.prev = current.prev
	}

	list.size--
	return current.value
}

// Extra - Remoção por nó
func (list *DoubleLinkedList) Remove(node *Node) {
	if node == nil {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}

	list.size--
}

// Busca
func (list *DoubleLinkedList) Find(value int) *Node {
	current := list.head

	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}

	return nil
}

// Percuso
func (list *DoubleLinkedList) PrintForward() {
	current := list.head

	for current != nil {
		println(current.value)
		current = current.next
	}
}

// Reverse
func (list *DoubleLinkedList) PrintBackward() {
	current := list.tail

	for current != nil {
		println(current.value)
		current = current.prev
	}
}

// Teste
func main() {
	list := NewDoubleLinkedList()

	list.Add(10)
	list.Add(20)
	list.Add(30)

	list.AddFirst(5)
	list.AddAt(2, 99)

	println("Forward:")
	list.PrintForward()

	println("Backward:")
	list.PrintBackward()

	println("Elemento no índice 2:", list.Get(2))

	list.RemoveAt(2)

	println("Após remoção:")
	list.PrintForward()
}
