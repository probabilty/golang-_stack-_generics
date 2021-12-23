package gstack
import ("sync")
type stack[T any] struct{
	mutex sync.Mutex
	name string
	stack []T
}
// var refrence *stack

func New[T any](item T) *stack[T]{
	return &(stack[T]{
		stack:[]T{item},
	})
}

func (s *stack[T]) Push(item T) {
	s.mutex.Lock()
	(*s).stack = append((*s).stack, item)
	s.mutex.Unlock()
}

func (s *stack[T]) Pop() (item T) {
	if len((*s).stack) == 0 {
		return
	}
	s.mutex.Lock()
	(*s).stack, item = (*s).stack[:len((*s).stack)-1], (*s).stack[len((*s).stack)-1]
	s.mutex.Unlock()
	return item
}
func (s *stack[T]) Peak() (item T) {
	if len((*s).stack) == 0 {
		return
	}
	s.mutex.Lock()
	item = (*s).stack[len((*s).stack)-1]
	s.mutex.Unlock()
	return item
}