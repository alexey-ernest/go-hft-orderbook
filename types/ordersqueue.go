package types

// Doubly linked orders queue
// TODO: this should be compared with ring buffer queue performance
type ordersQueue struct {
	head *Order
	tail *Order
	size int
}

func NewOrdersQueue() ordersQueue {
	return ordersQueue{}
}

func (this *ordersQueue) Size() int {
	return this.size
}

func (this *ordersQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *ordersQueue) Enqueue(o *Order) {
	tail := this.tail
	this.tail = o
	if tail != nil {
		tail.Next = o
	}
	if this.head == nil {
		this.head = o
	}
	this.size++
}

func (this *ordersQueue) Dequeue() *Order {
	if this.size == 0 {
		return nil
	}

	head := this.head
	if this.tail == this.head {
		this.tail = nil
	}

	this.head = this.head.Next
	this.size--
	return head
}
