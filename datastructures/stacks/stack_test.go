package stacks

import "testing"

func TestStack_Pop(t *testing.T) {
	Stack := &Stack{}
	Stack.Push(1)
	Stack.Push(2)
	Stack.Push(3)
	val := Stack.Pop()
	if val != 3 {
		t.Errorf("Expected 3 but got %d", val)
	}
}

func TestStack_Push(t *testing.T) {
	Stack := &Stack{}
	Stack.Push(1)
	Stack.Push(2)
	Stack.Push(3)
	if len(Stack.items) != 3 {
		t.Errorf("Expected stack length of 3 but got %d", len(Stack.items))
	}
}
