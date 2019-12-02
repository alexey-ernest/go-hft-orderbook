package types

// Limit price orders combined as a FIFO queue
type LimitOrder struct {
	Price float64
	Orders *ordersQueue

	totalVolume float64
}

func NewLimitOrder(price float64) LimitOrder {
	q := NewOrdersQueue()
	return LimitOrder{
		Price: price,
		Orders: &q,
	}
}

func (this *LimitOrder) TotalVolume() float64 {
	return this.totalVolume
}

func (this *LimitOrder) Size() int {
	return this.Orders.Size()
}

func (this *LimitOrder) Enqueue(o *Order) {
	this.Orders.Enqueue(o)
	o.Limit = this
	this.totalVolume += o.Volume
}

func (this *LimitOrder) Dequeue() *Order{
	if this.Orders.IsEmpty() {
		return nil
	}

	o := this.Orders.Dequeue()
	this.totalVolume -= o.Volume
	return o
}