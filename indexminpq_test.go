package hftorderbook

import (
	"testing"
	"math/rand"
	//"fmt"
)

func TestIndexMinPQOne(t *testing.T) {
	minpq := NewIndexMinPQ(10)
	minpq.Insert(0, 5.0)
	res := minpq.Top()

	if res != 5.0 {
		t.Errorf("actual %+v != expected %+v", res, 5)
	}
}

func TestIndexMinPQTwo(t *testing.T) {
	minpq := NewIndexMinPQ(10)
	minpq.Insert(0, 6.0)
	minpq.Insert(1, 5.0)
	
	res := [2]float64{}
	res[0] = minpq.Top()
	minpq.DelTop()
	res[1] = minpq.Top()

	exp := [2]float64{5.0, 6.0}
	if res != exp {
		t.Errorf("actual %+v != expected %+v", res, exp)
	}
}

func TestIndexMinPQThree(t *testing.T) {
	minpq := NewIndexMinPQ(10)
	minpq.Insert(0, 6.0)
	minpq.Insert(1, 5.0)
	minpq.Insert(2, 4.0)
	
	res := [3]float64{}
	res[0] = minpq.Top()
	minpq.DelTop()
	res[1] = minpq.Top()
	minpq.DelTop()
	res[2] = minpq.Top()
	minpq.DelTop()

	exp := [3]float64{4.0, 5.0, 6.0}
	if res != exp {
		t.Errorf("actual %+v != expected %+v", res, exp)
	}

	if !minpq.IsEmpty() {
		t.Errorf("pq should be empty")
	}
}

func TestIndexMinPQRandom(t *testing.T) {
	minpq := NewIndexMinPQ(100)
	emptyindex := 0
	for i := 0; i < 1000; i += 1 {
		emptyindex = i
		if minpq.Size() == 100 {
			emptyindex = minpq.DelTop()
		}
		minpq.Insert(emptyindex, float64(rand.Intn(100)))
	}

	res := [100]float64{}
	for i := range res {
		res[i] = minpq.Top()
		minpq.DelTop()
	}

	if !minpq.IsEmpty() {
		t.Errorf("pq should be empty after all")
	}

	for i := 1; i < 100; i += 1 {
		if res[i] < res[i-1] {
			t.Errorf("invalid order")
		}
	}
}

func BenchmarkIndexMinPQLimitedRandomInsertWithCaching(b *testing.B) {
	pq := NewIndexMinPQ(10000)

	// maximum number of levels in average is 10k
	limitslist := make([]float64, 10000)
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

	limitscache := make(map[float64]*LimitOrder)
	for i := 0; i < b.N; i += 1 {
		// create a new order
		o := orders[i]
		o.Id = i
		o.Volume = rand.Float64()
		// o := &Order{
		// 	Id: i,
		// 	Volume: rand.Float64(),
		// }

		// set the price
		price := limitslist[rand.Intn(len(limitslist))]

		// append order to the limit price
		if limitscache[price] != nil {
			// append to the existing limit in cache
			limitscache[price].Enqueue(o)
		} else {
			// new limit
			l := NewLimitOrder(price)
			l.Enqueue(o)

			// caching limit
			limitscache[price] = &l
			
			// inserting into heap
			pq.Insert(len(limitscache)-1, price)
		}
	}
}