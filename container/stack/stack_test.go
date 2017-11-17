package stack_test

import "testing"
import "github.com/sammy00/golang.lib/container/stack"

func TestStack(t *testing.T) {
	stack := stack.New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	len := stack.Len()
	if 4 != len {
		t.Errorf("stack.Len(): want %v, got %v\n", 4, len)
	}

	x := stack.Peek().(int)
	if 40 != x {
		t.Errorf("stack.Peek(): want %v, got %v\n", 40, x)
	}

	x = stack.Pop().(int)
	if 40 != x {
		t.Errorf("stack.Peek(): want %v, got %v\n", 40, x)
	}

	x = stack.Peek().(int)
	if 30 != x {
		t.Errorf("stack.Peek(): want %v, got %v\n", 30, x)
	}

	len = stack.Len()
	if 3 != len {
		t.Errorf("stack.Len(): want %v, got %v\n", 3, len)
	}

	x = stack.Pop().(int)
	if 30 != x {
		t.Errorf("stack.Peek(): want %v, got %v\n", 30, x)
	}

	x = stack.Pop().(int)
	if 20 != x {
		t.Errorf("stack.Peek(): want %v, got %v\n", 20, x)
	}

	x = stack.Pop().(int)
	if 10 != x {
		t.Errorf("stack.Peek(): want %v, got %v\n", 10, x)
	}

	y := stack.Peek()
	if nil != y {
		t.Errorf("stack.Peek(): want nil, got %v\n", y)
	}

	y = stack.Pop()
	if nil != y {
		t.Errorf("stack.Peek(): want nil, got %v\n", y)
	}

	len = stack.Len()
	if 0 != len {
		t.Errorf("stack.Len(): want %v, got %v\n", 0, len)
	}

	if !stack.Empty() {
		t.Errorf("stack.Empty(): want true, got false\n")
	}
}
