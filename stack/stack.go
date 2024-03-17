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

func (s *Stack[T]) Push(item T) {
	s.Length++
	newNode := Node[T]{value: item}

	if s.head == nil {
		s.head = &newNode
	}

	newNode.prev = s.head
	s.head = &newNode
}

func (s *Stack[T]) Pop() T {
	var result T

	if s.head == nil {
		return result
	}
	s.Length--

	currentHead := s.head
	s.head = s.head.prev

	// very important to set head to nil when length is zero
	if s.Length == 0 {
		s.head = nil
	}
	return currentHead.value
}

func (s *Stack[T]) Peek() T {
	var result T
	if s.head == nil {
		return result
	}
	return s.head.value
}
