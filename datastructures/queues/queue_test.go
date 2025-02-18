package queues

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	Queue := &Queue{}
	Queue.Enqueue(1)
	Queue.Enqueue(2)
	Queue.Enqueue(3)
	if len(Queue.items) != 3 {
		t.Errorf("Expected queue length of 3 but got %d", len(Queue.items))
	}
}

func TestQueue_Dequeue(t *testing.T) {
	Queue := &Queue{}
	val := Queue.Dequeue()
	if val != -1 {
		t.Errorf("Expected -1 but got %d", val)
	}
	Queue.Enqueue(1)
	Queue.Enqueue(2)
	Queue.Enqueue(3)
	val = Queue.Dequeue()
	if val != 1 {
		t.Errorf("Expected 1 but got %d", val)
	}
	val = Queue.Dequeue()
	if val != 2 {
		t.Errorf("Expected 2 but got %d", val)
	}
	val = Queue.Dequeue()
	if val != 3 {
		t.Errorf("Expected 3 but got %d", val)
	}
	val = Queue.Dequeue()
	if val != -1 {
		t.Errorf("Expected -1 but got %d", val)
	}
}
