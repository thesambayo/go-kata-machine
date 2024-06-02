package queue

type Node[T any] struct {
	value T
	// pointer-value of nextNode, zerothValue is nil
	prev *Node[T]
}

type Queue[T any] struct {
	Length int
	// reference to pointer-value of headNode, zerothValue is nil
	head *Node[T]
}

type Methods[T any] interface {
Enqueue(item T) 
Deque() T
Peek() T

}

func(q *Queue[T]) Enqueue(item T)  {}
func(q *Queue[T]) Deque() T {}
func(q *Queue[T]) Peek() T {}

