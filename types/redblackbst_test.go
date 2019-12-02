package types

import "testing"

func TestRedBlackEmpty(t *testing.T) {
	rb := NewRedBlackBST()
	if rb.Size() != 0 || !rb.IsEmpty() {
		t.Errorf("Red Black BST should be empty")
	}
}

// func TestRedBlackBasic(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + i)
// 		st.Put(k, i)
// 	}

// 	if st.Size() != 10 {
// 		t.Errorf("size should equals 10, got %d", st.Size())
// 	}
// 	if st.IsEmpty() {
// 		t.Errorf("st should not be empty")	
// 	}

// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + i)
// 		if !st.Contains(k) {
// 			t.Errorf("st should contain the key %q", k)
// 		}
// 		if st.Get(k) != i {
// 			t.Errorf("value %s != %s", st.Get(k), i)
// 		}
// 	}
// }

// func TestRedBlackHeight(t *testing.T) {
// 	st := NewRedBlackBST()
// 	n := 1024
// 	for i := 0; i < n; i+=1 {
// 		k := string(i)
// 		st.Put(k, i)
// 	}

// 	if st.Size() != n {
// 		t.Errorf("size should equals %d, got %d", n, st.Size())
// 	}
// 	if st.IsEmpty() {
// 		t.Errorf("st should not be empty")	
// 	}

// 	height := st.Height()
// 	if height < 10 || height > 20 {
// 		t.Errorf("red black bst height should be in range lgN <= height <= 2*lgN, in our case from 10 to 20, but we got %d", height)
// 	}
// }

// func TestRedBlackLeipzig1M(t *testing.T) {

// 	file, err := os.Open("leipzig1M.txt")
//     if err != nil {
//         panic(err)
//     }
//     defer file.Close()

//     scanner := bufio.NewScanner(file)
//     st := NewRedBlackBST()
//     n := 0
//     for scanner.Scan() {
//         line := scanner.Text()
        
//         words := strings.Fields(line)
//         n += len(words)
        
//         for _, w := range words {
//         	if !st.Contains(w) {
//         		st.Put(w, 1)
//         	} else {
//         		st.Put(w, st.Get(w) + 1)
//         	}
//         }
//     }

//     if err := scanner.Err(); err != nil {
//         panic(err)
//     }

//     allwords := 622098
//     if n != allwords {
//     	t.Errorf("number of all words should be %d, got %d", allwords, n)
//     }

//     m := 69321
//     if st.Size() != m {
// 		t.Errorf("size should equals %d, got %d", m, st.Size())
// 	}

// 	height := st.Height()
// 	if height < 16 || height > 32 {
// 		t.Errorf("red black bst height should be in range lgN <= height <= 2*lgN, in our case from 16 to 32, but we got %d", height)
// 	}

// 	if !st.IsRedBlack() {
// 		t.Errorf("certification failed")
// 	}
// }

// func TestRedBlackMinMax(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + 9 - i)
// 		st.Put(k, i)
// 	}

// 	min := string('a')
// 	if st.Min() != min {
// 		t.Errorf("min %q != %q", st.Min(), min)
// 	}

// 	max := string('a' + 9)
// 	if st.Max() != max {
// 		t.Errorf("min %q != %q", st.Max(), max)
// 	}
// }

// func TestRedBlackFloor(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i += 1 {
// 		k := string('a' + 20 - 2*i)
// 		st.Put(k, i)
// 	}

// 	keymiss := string('a' + 3)
// 	flmiss := string('a' + 2)
// 	if st.Floor(keymiss) != flmiss {
// 		t.Errorf("floor != %s", st.Floor(keymiss))
// 	}

// 	keyhit := string('a' + 10)
// 	if st.Floor(keyhit) != keyhit {
// 		t.Errorf("floor != %s", st.Floor(keyhit))
// 	}
// }

// func TestRedBlackCeiling(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i += 1 {
// 		k := string('a' + 20 - 2*i)
// 		st.Put(k, i)
// 	}

// 	keymiss := string('a' + 3)
// 	clmiss := string('a' + 4)
// 	if st.Ceiling(keymiss) != clmiss {
// 		t.Errorf("ceiling != %s", st.Ceiling(keymiss))
// 	}

// 	keyhit := string('a' + 10)
// 	if st.Ceiling(keyhit) != keyhit {
// 		t.Errorf("ceiling != %s", st.Ceiling(keyhit))
// 	}
// }

// func TestRedBlackSelect(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + 10 - i)
// 		st.Put(k, i)
// 	}

// 	key := string('a' + 3)
// 	if st.Select(2) != key {
// 		t.Errorf("element with rank=2 should be %s", key)
// 	}

// 	key = string('a' + 10)
// 	if st.Select(9) != key {
// 		t.Errorf("element with rank=9 should be %s", key)
// 	}
// }

// func TestRedBlackRank(t *testing.T) {
// 	st := NewRedBlackBST()
// 	nodes := []string{"S", "E", "X", "A", "R", "C", "H", "M"}
// 	for i, v := range nodes {
// 		st.Put(v, i)
// 	}

// 	for i := range nodes {
// 		k := st.Select(i)
// 		if st.Rank(k) != i {
// 			t.Errorf("rank of %q != %d", k, i)
// 		}
// 	}

// 	if st.Rank("Y") != len(nodes) {
// 		t.Errorf("rank of new maximum should equal to the number of nodes in the tree")
// 	}

// 	if st.Rank("Y") != st.Rank("Z") {
// 		t.Errorf("rank of new maximum should not depend on the new maximum concrete value")
// 	}
// }

// func TestRedBlackKeys(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + 10 - i)
// 		st.Put(k, i)
// 	}

// 	lo := string('a' + 3)
// 	hi := string('a' + 6)
// 	keys := st.Keys(lo, hi)
// 	if len(keys) != 4 {
// 		t.Errorf("keys len should equals 4, %+v", keys)
// 	}

// 	if keys[0] != lo {
// 		t.Errorf("first key should be %s", lo)
// 	}

// 	if keys[len(keys)-1] != hi {
// 		t.Errorf("last key should be %s", hi)
// 	}

// 	for i := 1; i < len(keys); i += 1 {
// 		if keys[i] < keys[i-1] {
// 			t.Errorf("non-decreasing keys order validation failed")
// 		}
// 	}
// }

// func TestRedBlackDeleteMin(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + 10 - i)
// 		st.Put(k, i)
// 	}

// 	st.DeleteMin()
// 	if st.Size() != 9 {
// 		t.Errorf("tree size should shrink")
// 	}

// 	if st.Contains(string('a' + 1)) {
// 		t.Errorf("minimum element should be removed from the tree")
// 	}

// 	if !st.IsRedBlack() {
// 		t.Errorf("certification failed")
// 	}
// }

// func TestRedBlackDeleteMax(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + i)
// 		st.Put(k, i)
// 	}

// 	st.DeleteMax()
// 	if st.Size() != 9 {
// 		t.Errorf("tree size should shrink")
// 	}

// 	if st.Contains(string('a' + 9)) {
// 		t.Errorf("minimum element should be removed from the tree")
// 	}

// 	if !st.IsRedBlack() {
// 		t.Errorf("certification failed")
// 	}
// }

// func TestRedBlackDelete(t *testing.T) {
// 	st := NewRedBlackBST()
// 	for i := 0; i < 10; i+=1 {
// 		k := string('a' + i)
// 		st.Put(k, i)
// 	}

// 	key := string('a' + 5)
// 	st.Delete(key)
// 	if st.Size() != 9 {
// 		t.Errorf("tree size should shrink")
// 	}

// 	if st.Contains(key) {
// 		t.Errorf("minimum element should be removed from the tree")
// 	}

// 	if !st.IsRedBlack() {
// 		t.Errorf("certification failed")
// 	}
// }