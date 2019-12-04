package hftorderbook

import (
	"math/rand"
	"testing"
)

func TestLimitOrderEmpty(t *testing.T) {
	price := 3.141593
	l := NewLimitOrder(price)
	if l.Price != price || l.TotalVolume() != 0.0 {
		t.Errorf("limit order init error")
	}
}

func TestLimitOrderAddOrder(t *testing.T) {
	price := 3.141593
	volume := 25.0
	l := NewLimitOrder(price)
	o := &Order{ Volume: volume }
	l.Enqueue(o)
	
	if l.TotalVolume() != volume {
		t.Errorf("total volume counted incorrectly")
	}
	if l.Size() != 1 {
		t.Errorf("it should have size = 1")
	}
	if o.Limit != &l {
		t.Errorf("Parent Limit link should be set for an order")
	}
}

func TestLimitOrderAddMultipleOrders(t *testing.T) {
	price := 3.141593
	volume := 0.0
	l := NewLimitOrder(price)
	n := 100
	for i := 0; i < n; i += 1 {
		o := &Order{ Id: i, Volume: rand.Float64() }
		volume += o.Volume
		l.Enqueue(o)
	}

	if volume != l.TotalVolume() {
		t.Errorf("total volume calculated incorrectly")
	}

	if l.Size() != n {
		t.Errorf("total count calculated incorrectly")
	}

	o := l.Dequeue()
	if l.TotalVolume() != volume - o.Volume {
		t.Errorf("total volume calculated incorrectly")
	}
	if l.Size() != n - 1 {
		t.Errorf("total count calculated incorrectly")
	}
}
