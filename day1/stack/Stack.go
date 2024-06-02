package stack

type Node[T any] struct {
	value T
	// pointer-value of nextNode, zerothValue is nil
	prev *Node[T]
}

type Stack[T any] struct {
	Length int
	// reference to pointer-value of headNode, zerothValue is nil
	head *Node[T]
}

type Methods[T any] interface {
Push(item T) 
Pop() T
Peek() T

}

func(s *Stack[T]) Push(item T)  {}
func(s *Stack[T]) Pop() T {}
func(s *Stack[T]) Peek() T {}

