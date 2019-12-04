package hftorderbook

import (
	"github.com/alexey-ernest/go-hft-orderbook/types"
)

type OrderbookAPI interface {
	
	// Add an order to a price level: O(LogM) for a new price level, O(1) for existing
	Add(float64, *types.Order)

	// Cancel an order: O(1) if limit still has some orders, O(logM) otherwise
	Cancel(*types.Order)

	// Clear a limit: O(1)
	ClearBidLimit(float64)
	ClearAskLimit(float64)

	// Delete a limit from orderbook eagerly: O(logM)
	DeleteBidLimit(float64)
	DeleteAskLimit(float64)

	// Get a volume at a specific limit price: O(1)
	GetVolumeAtBidLimit(float64) float64
	GetVolumeAtAskLimit(float64) float64

	// O(1)
	GetBestBid() float64

	// O(1)
	GetBestOffer() float64

	BLength() int
	ALength() int
}