package hftorderbook

import (
	"github.com/alexey-ernest/go-hft-orderbook/types"
)

func NewOrderbook() OrderbookAPI {
	b := types.NewOrderbook()
	return &b
}