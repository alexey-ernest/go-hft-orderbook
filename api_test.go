package hftorderbook

import "testing"

func TestOrderbookAPI(t *testing.T) {
	b := NewOrderbook()
	if b.BLength() != 0 {
		t.Errorf("book should be empty")
	}
	if b.ALength() != 0 {
		t.Errorf("book should be empty")
	}
}