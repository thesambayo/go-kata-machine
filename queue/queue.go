package queue

// Node note that I used pointers instead of the actual node values
type Node[T any] struct {
	value T
	// pointer-value of nextNode, zerothValue is nil
	next *Node[T]
}
type Queue[T any] struct {
	// number of nodes in queue
	Length int
	// reference to pointer-value of headNode, zerothValue is nil
	head *Node[T]
	// reference to pointer-value of tailNode, zerothValue is nil
	tail *Node[T]
}

type Methods[T any] interface {
	Enqueue(item T)
	Deque() T
	Peek() T
}

func (q *Queue[T]) Enqueue(item T) {
	q.Length++
	newNode := Node[T]{value: item}

	if q.tail == nil {
		q.tail = &newNode
		q.head = &newNode
		return
	}

	q.tail.next = &newNode
	q.tail = &newNode
}

func (q *Queue[T]) Deque() T {
	var result T

	if q.head == nil {
		return result
	}
	q.Length--
	currentHead := q.head
	q.head = q.head.next

	if q.Length == 0 {
		q.tail = nil
	}

	return currentHead.value
}

func (q *Queue[T]) Peek() T {
	var result T // zeroth value of T

	if q.head == nil {
		return result
	}
	return q.head.value
}
