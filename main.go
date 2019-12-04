package hftorderbook

import (
	"types"
)

func NewOrderbook() OrderbookAPI {
	bids := types.NewRedBlackBST()
	asks := types.NewRedBlackBST()

	return types.Orderbook{
		Bids: &bids,
		Asks: &asks,

		bidLimitsCache: make(map[float64]*types.LimitOrder, types.MaxLimitsNum),
		askLimitsCache: make(map[float64]*types.LimitOrder, types.MaxLimitsNum),
	}
}