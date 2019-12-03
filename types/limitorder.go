package types

// Limit price orders combined as a FIFO queue
type LimitOrder struct {
	Price float64
	
	orders *ordersQueue
	totalVolume float64
}

func NewLimitOrder(price float64) LimitOrder {
	q := NewOrdersQueue()
	return LimitOrder{
		Price: price,
		orders: &q,
	}
}

func (this *LimitOrder) TotalVolume() float64 {
	return this.totalVolume
}

func (this *LimitOrder) Size() int {
	return this.orders.Size()
}

func (this *LimitOrder) Enqueue(o *Order) {
	this.orders.Enqueue(o)
	o.Limit = this
	this.totalVolume += o.Volume
}

func (this *LimitOrder) Dequeue() *Order{
	if this.orders.IsEmpty() {
		return nil
	}

	o := this.orders.Dequeue()
	this.totalVolume -= o.Volume
	return o
}