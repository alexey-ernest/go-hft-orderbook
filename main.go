package hftorderbook

import (
	"./types"
)

func NewOrderbook() OrderbookAPI {
	b := types.NewOrderbook()
	return &b
}