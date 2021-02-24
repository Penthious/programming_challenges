package main

import (
	"container/list"
	"fmt"
	"sync"
)

type stack struct {
	List  *list.List
	Mutex sync.Mutex
}

func NewStack() *stack {
	return &stack{List: list.New()}
}

func (s *stack) isEmpty() bool {
	return s.List.Len() == 0
}

func (s *stack) push(item interface{}) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.List.PushBack(item)

}

func (s *stack) pop() (interface{}, bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.isEmpty() {
		return "", false
	}
	tail := s.List.Back()
	value := tail.Value
	s.List.Remove(tail)
	return value, true
}

func main() {

	stack := NewStack()
	stack.push("Hello")
	stack.push("my")
	stack.push("name")
	stack.push("is")
	stack.push("bob")
	stack.push(1)

	item, _ := stack.pop()
	fmt.Println(item)
}
