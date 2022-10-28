package hftorderbook

import (
	"fmt"
	"sync"
)

// maximum limits per orderbook side to pre-allocate memory
const MaxLimitsNum int = 10000

type Orderbook struct {
	Bids *redBlackBST
	Asks *redBlackBST

	bidLimitsCache map[float64]*LimitOrder
	askLimitsCache map[float64]*LimitOrder
	pool *sync.Pool
}

func NewOrderbook() Orderbook {
	bids := NewRedBlackBST()
	asks := NewRedBlackBST()
	return Orderbook{
		Bids: &bids,
		Asks: &asks,

		bidLimitsCache: make(map[float64]*LimitOrder, MaxLimitsNum),
		askLimitsCache: make(map[float64]*LimitOrder, MaxLimitsNum),
		pool: &sync.Pool {
			New: func()interface{} {
				limit := NewLimitOrder(0.0)
				return &limit
			},
		},
	}
}

func (this *Orderbook) Add(price float64, o *Order) {
	var limit *LimitOrder

	if o.BidOrAsk {
		limit = this.bidLimitsCache[price]
	} else {
		limit = this.askLimitsCache[price]
	}

	if limit == nil {
		// getting a new limit from pool
		limit = this.pool.Get().(*LimitOrder)
		limit.Price = price

		// insert into the corresponding BST and cache
		if o.BidOrAsk {
			this.Bids.Put(price, limit)
			this.bidLimitsCache[price] = limit
		} else {
			this.Asks.Put(price, limit)
			this.askLimitsCache[price] = limit
		}
	}

	// add order to the limit
	limit.Enqueue(o)
}

func (this *Orderbook) Cancel(o *Order) {
	limit := o.Limit
	limit.Delete(o)
	
	if limit.Size() == 0 {
		// remove the limit if there are no orders
		if o.BidOrAsk {
			this.Bids.Delete(limit.Price)
			delete(this.bidLimitsCache, limit.Price)
		} else {
			this.Asks.Delete(limit.Price)
			delete(this.askLimitsCache, limit.Price)
		}

		// put it back to the pool
		this.pool.Put(limit)
	}
}

func (this *Orderbook) ClearBidLimit(price float64) {
	this.clearLimit(price, true)
}

func (this *Orderbook) ClearAskLimit(price float64) {
	this.clearLimit(price, false)
}

func (this *Orderbook) clearLimit(price float64, bidOrAsk bool) {
	var limit *LimitOrder
	if bidOrAsk {
		limit = this.bidLimitsCache[price]
	} else {
		limit = this.askLimitsCache[price]
	}
	
	if limit == nil {
		panic(fmt.Sprintf("there is no such price limit %0.8f", price))
	}

	limit.Clear()
}

func (this *Orderbook) DeleteBidLimit(price float64) {
	limit := this.bidLimitsCache[price]
	if limit == nil {
		return
	}

	this.deleteLimit(price, true)
	delete(this.bidLimitsCache, price)

	// put limit back to the pool
	limit.Clear()
	this.pool.Put(limit)

}

func (this *Orderbook) DeleteAskLimit(price float64) {
	limit := this.askLimitsCache[price]
	if limit == nil {
		return
	}

	this.deleteLimit(price, false)
	delete(this.askLimitsCache, price)

	// put limit back to the pool
	limit.Clear()
	this.pool.Put(limit)
}

func (this *Orderbook) deleteLimit(price float64, bidOrAsk bool) {
	if bidOrAsk {
		this.Bids.Delete(price)
	} else {
		this.Asks.Delete(price)
	}
}

func (this *Orderbook) GetVolumeAtBidLimit(price float64) float64 {
	limit := this.bidLimitsCache[price]
	if limit == nil {
		return 0
	}
	return limit.TotalVolume()
}

func (this *Orderbook) GetVolumeAtAskLimit(price float64) float64 {
	limit := this.askLimitsCache[price]
	if limit == nil {
		return 0
	}
	return limit.TotalVolume()
}

func (this *Orderbook) GetBestBid() float64 {
	return this.Bids.Max()
}

func (this *Orderbook) GetBestOffer() float64 {
	return this.Asks.Min()
}

func (this *Orderbook) BLength() int {
	return len(this.bidLimitsCache)
}

func (this *Orderbook) ALength() int {
	return len(this.askLimitsCache)
}
