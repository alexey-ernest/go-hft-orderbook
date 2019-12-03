package hftorderbook

import (
	"types"
)

func NewOrderbook() OrderbookAPI {
	return types.Orderbook{
		bidLimitsCache: make(map[float64]*LimitOrder),
		askLimitsCache: make(map[float64]*LimitOrder),
		ordersCache: make(map[int]*Order),
	}
}