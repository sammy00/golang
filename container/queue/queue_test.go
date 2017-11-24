package queue_test

import (
	"testing"

	"github.com/sammy00/golang/container/queue"
)

func TestStack(t *testing.T) {
	queue := queue.New()

	queue.Push(10)
	queue.Push(20)
	queue.Push(30)
	queue.Push(40)

	frontWant := []int{10, 20, 30, 40}
	lenWant := []int{4, 3, 2, 1}

	for i := range frontWant {
		if lenWant[i] != queue.Len() {
			t.Fatalf("Len(): want %v, got %v", lenWant[i], queue.Len())
		}

		if x, ok := queue.Peek().(int); ok {
			if frontWant[i] != x {
				t.Fatalf("Peek(): want %v, got %v", frontWant[i], x)
			}
		} else {
			t.Fatal("unexpected error due to type assertion")
		}

		queue.Pop()
	}

	if 0 != queue.Len() {
		t.Fatalf("Len(): want 0, got %v\n", queue.Len())
	}

	if !queue.Empty() {
		t.Fatal("Empty(): want true, got false")
	}
}
