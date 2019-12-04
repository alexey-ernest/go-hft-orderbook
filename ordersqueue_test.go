package hftorderbook

import (
	"testing"
)

func TestOrdersQueueEmpty(t *testing.T) {
	q := NewOrdersQueue()
	if !q.IsEmpty() || q.Size() != 0 {
		t.Errorf("a queue should be initialized as empty")
	}
}

func TestOrdersQueueEnqueDequeue(t *testing.T) {
	q := NewOrdersQueue()
	o := &Order{
		Id: 1,
	}
	q.Enqueue(o)
	if q.Size() != 1 || q.IsEmpty() {
		t.Errorf("a queue size should be non-zero")
	}

	r := q.Dequeue()
	if q.Size() != 0 || !q.IsEmpty() {
		t.Errorf("a queue should be empty now")	
	}
	r2 := q.Dequeue()
	if r2 != nil {
		t.Errorf("no elements should remain in the queue")
	}
	if o != r {
		t.Errorf("consistency failed")
	}

	q.Enqueue(o)
	q.Dequeue()
	if q.Size() != 0 || !q.IsEmpty() {
		t.Errorf("a queue should be empty now")	
	}
}

func TestOrdersQueueMultiple(t *testing.T) {
	q := NewOrdersQueue()
	n := 100
	for i := 0; i < n; i += 1 {
		q.Enqueue(&Order{ Id: i })
	}

	if q.Size() != n {
		t.Errorf("queue size should be %d", n)
	}

	for i := 0; i < n; i += 1 {
		o := q.Dequeue()
		if o.Id != i {
			t.Errorf("order invariant failed")
			break
		}
		if q.Size() != n-i-1 {
			t.Errorf("invalid size counter")
			break
		}
	}
}

func TestOrdersQueueMixed(t *testing.T) {
	q := NewOrdersQueue()
	n := 100
	for i := 0; i < n; i += 1 {
		q.Enqueue(&Order{ Id: i })
		q.Dequeue()
	}

	if q.Size() != 0 || !q.IsEmpty() {
		t.Errorf("a queue should be empty now")	
	}
}