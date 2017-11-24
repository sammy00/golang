// Package queue implements a queue based on single linked list.
package queue

// Element is a node to build a single linked list
type Element struct {
	Value interface{}
	next  *Element
}

// Next accesses the next element of current one
func (e *Element) Next() *Element {
	return e.next
}

// Queue represents a queue
type Queue struct {
	front *Element
	tail  *Element
	size  int
}

// New returns an initialized queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Push adds an element to the queue
func (q *Queue) Push(x interface{}) {
	if nil == q.front {
		q.front = &Element{x, nil}
		q.tail = q.front
	} else {
		q.tail.next = &Element{x, nil}
		q.tail = q.tail.next
	}

	q.size++
}

// Pop out the front element from the queue
func (q *Queue) Pop() interface{} {
	var x interface{}

	if nil != q.front {
		x, q.front = q.front.Value, q.front.next
		q.size--
	}

	return x
}

// Peek returns the element in the front of the s
func (s *Queue) Peek() interface{} {
	if nil != s.front {
		return s.front.Value
	}

	return nil
}

// Len returns the size of the s
func (s *Queue) Len() int {
	return s.size
}

// Empty returns true if the queue is empty
func (s *Queue) Empty() bool {
	return (0 == s.size)
}
