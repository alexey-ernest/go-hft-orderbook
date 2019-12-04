package hftorderbook

import (
	"testing"
	"math/rand"
	//"fmt"
)

func TestMinPQOne(t *testing.T) {
	minpq := NewMinPQ(10)
	minpq.Insert(5.0)
	res := minpq.Top()

	if res != 5.0 {
		t.Errorf("actual %+v != expected %+v", res, 5)
	}
}

func TestMinPQTwo(t *testing.T) {
	minpq := NewMinPQ(10)
	minpq.Insert(6.0)
	minpq.Insert(5.0)
	
	res := [2]float64{}
	res[0] = minpq.Top()
	minpq.DelTop()
	res[1] = minpq.Top()

	exp := [2]float64{5.0, 6.0}
	if res != exp {
		t.Errorf("actual %+v != expected %+v", res, exp)
	}
}

func TestMinPQThree(t *testing.T) {
	minpq := NewMinPQ(10)
	minpq.Insert(6.0)
	minpq.Insert(5.0)
	minpq.Insert(4.0)
	
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

func TestMinPQRandom(t *testing.T) {
	minpq := NewMinPQ(100)
	for i := 0; i < 1000; i += 1 {
		if minpq.Size() == 100 {
			minpq.DelTop()
		}
		minpq.Insert(float64(rand.Intn(100)))
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
			break
		}
	}
}



func benchmarkMinPQLimitedRandomInsertWithCaching(n int, b *testing.B) {
	pq := NewMinPQ(n)

	// maximum number of levels in average is ~10k
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
			pq.Insert(price)
		}
	}
}

func BenchmarkMinPQ5kLevelsRandomInsertWithCaching(b *testing.B) {
	benchmarkMinPQLimitedRandomInsertWithCaching(5000, b)
}

func BenchmarkMinPQ10kLevelsRandomInsertWithCaching(b *testing.B) {
	benchmarkMinPQLimitedRandomInsertWithCaching(10000, b)
}

func BenchmarkMinPQ20kLevelsRandomInsertWithCaching(b *testing.B) {
	benchmarkMinPQLimitedRandomInsertWithCaching(20000, b)
}
