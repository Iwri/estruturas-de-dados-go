package main

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type Deque struct {
	head *Node
	tail *Node
	size int
}

func NewDeque() *Deque {
	return &Deque{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

// ADDfront
func (d *Deque) AddFront(value int) {
	newNode := &Node{value: value}

	if d.IsEmpty() {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

// ADDrear
func (d *Deque) AddRear(value int) {
	newNode := &Node{value: value}

	if d.IsEmpty() {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.size++
}

// Removefront
func (d *Deque) RemoveFront() int {
	if d.IsEmpty() {
		panic("deque is empty")
	}

	value := d.head.value
	d.head = d.head.next

	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}

	d.size--
	return value
}

// Removerear
func (d *Deque) RemoveRear() int {
	if d.IsEmpty() {
		panic("deque is empty")
	}

	value := d.tail.value
	d.tail = d.tail.prev

	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}

	d.size--
	return value
}

// Peek
func (d *Deque) PeekFront() int {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	return d.head.value
}

func (d *Deque) PeekRear() int {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	return d.tail.value
}

// Size
func (d *Deque) Size() int {
	return d.size
}

// Teste
func main() {
	d := NewDeque()

	d.AddFront(10)
	d.AddFront(5)
	d.AddRear(20)
	d.AddRear(30)

	println("Front:", d.PeekFront()) // 5
	println("Rear:", d.PeekRear())   // 30

	println("RemoveFront:", d.RemoveFront()) // 5
	println("RemoveRear:", d.RemoveRear())   // 30

	println("Novo Front:", d.PeekFront()) // 10
}
