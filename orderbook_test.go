package hftorderbook

import (
	"testing"
	"math"
	"math/rand"
	//"fmt"
)

func TestOrderbookEmpty(t *testing.T) {
	b := NewOrderbook()
	if b.BLength() != 0 {
		t.Errorf("book should be empty")
	}
	if b.ALength() != 0 {
		t.Errorf("book should be empty")
	}
}

func TestOrderbookAddOne(t *testing.T) {
	b := NewOrderbook()
	bid := &Order{
		BidOrAsk: true,
	}
	ask := &Order{
		BidOrAsk: false,
	}
	b.Add(1.0, bid)
	b.Add(2.0, ask)
	if b.BLength() != 1 {
		t.Errorf("book should have 1 bid")
	}
	if b.ALength() != 1 {
		t.Errorf("book should have 1 ask")
	}
}

func TestOrderbookAddMultiple(t *testing.T) {
	b := NewOrderbook()
	for i := 0; i < 100; i += 1 {
		bid := &Order{
			BidOrAsk: true,
		}
		b.Add(float64(i), bid)
	}

	for i := 100; i < 200; i += 1 {
		bid := &Order{
			BidOrAsk: false,
		}
		b.Add(float64(i), bid)
	}

	if b.BLength() != 100 {
		t.Errorf("book should have 100 bids")
	}
	if b.ALength() != 100 {
		t.Errorf("book should have 100 asks")
	}

	if b.GetBestBid() != 99.0 {
		t.Errorf("best bid should be 99.0")
	}

	if b.GetBestOffer() != 100.0 {
		t.Errorf("best offer should be 100.0")
	}
}

func TestOrderbookAddAndCancel(t *testing.T) {
	b := NewOrderbook()
	bid1 := &Order{
		Id: 1,
		BidOrAsk: true,
	}
	bid2 := &Order{
		Id: 2,
		BidOrAsk: true,
	}
	b.Add(1.0, bid1)
	b.Add(2.0, bid2)
	if b.GetBestBid() != 2.0 {
		t.Errorf("best bid should be 2.0")
	}
	b.Cancel(bid2)
	if b.GetBestBid() != 1.0 {
		t.Errorf("best bid should be 1.0 now")
	}
}

func TestGetVolumeAtLimit(t *testing.T) {
	b := NewOrderbook()
	bid1 := &Order{
		Id: 1,
		BidOrAsk: true,
		Volume: 0.1,
	}
	bid2 := &Order{
		Id: 2,
		BidOrAsk: true,
		Volume: 0.2,
	}
	b.Add(1.0, bid1)
	b.Add(1.0, bid2)
	if math.Abs(b.GetVolumeAtBidLimit(1.0) - 0.3) > 0.0000001 {
		t.Errorf("invalid volume at limit: %0.8f", b.GetVolumeAtBidLimit(1.0))
	}
}

func benchmarkOrderbookLimitedRandomInsert(n int, b *testing.B) {
	book := NewOrderbook()

	// maximum number of levels in average is 10k
	limitslist := make([]float64, n)
	for i := range limitslist {
		limitslist[i] = rand.Float64()
	}
	
	// preallocate empty orders
	orders := make([]*Order, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		orders = append(orders, &Order{})
	}

	// measure insertion time
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		price := limitslist[rand.Intn(len(limitslist))]

		// create a new order
		o := orders[i]
		o.Id = i
		o.Volume = rand.Float64()
		o.BidOrAsk = price < 0.5

		// add to the book
		book.Add(price, o)
	}

	//fmt.Printf("bid size %d, ask size %d\n", book.BLength(), book.ALength())
}

func BenchmarkOrderbook5kLevelsRandomInsert(b *testing.B) {
	benchmarkOrderbookLimitedRandomInsert(5000, b)
}

func BenchmarkOrderbook10kLevelsRandomInsert(b *testing.B) {
	benchmarkOrderbookLimitedRandomInsert(10000, b)
}

func BenchmarkOrderbook20kLevelsRandomInsert(b *testing.B) {
	benchmarkOrderbookLimitedRandomInsert(20000, b)
}
