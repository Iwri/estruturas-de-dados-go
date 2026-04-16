package main

type Node struct {
	value int
	next  *Node
}

type QueueLinked struct {
	head *Node
	tail *Node
	size int
}

func NewQueueLinked() *QueueLinked {
	return &QueueLinked{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Enqueue
func (q *QueueLinked) Enqueue(value int) {
	newNode := &Node{value: value}

	if q.size == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

// Dequeue
func (q *QueueLinked) Dequeue() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	value := q.head.value
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return value
}

// Peek
func (q *QueueLinked) Peek() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.head.value
}

// IsEmpty
func (q *QueueLinked) IsEmpty() bool {
	return q.size == 0
}

// Size
func (q *QueueLinked) Size() int {
	return q.size
}

// Teste
func main() {
	q := NewQueueLinked()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	println(q.Dequeue()) // 10
	println(q.Peek())    // 20
}
