// Package stack implements a stack.
package stack

import (
	"container/list"
)

// Stack represents a stack
// The zero value for Stack is an unintialized stack
// ought to initialize by invoking Init()
type Stack struct {
	list *list.List // data store for stack element
}

// New returns an initialized stack
func New() *Stack {
	list := list.New()
	return &Stack{list}
}

// Init initializes or clears the stack
func (s *Stack) Init() *Stack {
	s.list.Init()
	return s
}

// Push adds an element to the stack
func (stack *Stack) Push(x interface{}) {
	stack.list.PushBack(x)
}

// Pop out the top element from the stack
func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if nil != e {
		stack.list.Remove(e)
		return e.Value
	}

	return nil
}

// Peek returns the element in the top of the stack
func (stack *Stack) Peek() interface{} {
	e := stack.list.Back()
	if nil != e {
		return e.Value
	}

	return nil
}

// Len returns the size of the stack
func (s *Stack) Len() int {
	return s.list.Len()
}

// Empty returns true if the stack is empty
func (s *Stack) Empty() bool {
	if (nil == s.list) || (0 == s.list.Len()) {
		return true
	}
	return false
}
