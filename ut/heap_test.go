package ut

import (
	"slices"
	"testing"
)

var td TestData

func Test_heap(t *testing.T) {
	td.Init()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"insert", insertTest},
		{"heapSort", heapSortTest},
		// {"print", printTest},
		// {"heapify", heapifyTest},
		// {"IsMaxHeap", IsMaxHeapTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func IsMaxHeapTest(t *testing.T) {
	tests := []struct {
		name string
		seq  []int
	}{
		{"TDManberHeapIsMaxHeap", TDManberHeapIsMaxHeap},
	}

	for _, tt := range tests {
		h := Heap{Imp: tt.seq, t: t}
		if !h.IsMaxHeap() {
			t.Fatalf("NOT A MAX HEAP")
		}
	}
}
func heapifyTest(t *testing.T) {
	tests := []struct {
		name string
		seq  []int
	}{
		// {"hello", td.Test11},
		{"TDManberHeap", TDManberHeap},
	}

	for _, tt := range tests {
		h := Heap{Imp: tt.seq, t: t}
		h.Heapify()
		if !h.IsMaxHeap() {
			t.Fatalf("NOT A MAX HEAP")
		}
	}
}

func heapSortTest(t *testing.T) {

	tests := []struct {
		name       string
		seq        []int
		descending bool
	}{
		{"hello", td.Test11, true},
		{"hello", td.Test11, false},
		{"TDManberHeap", slices.Clone(TDManberHeap), true},
		{"TDManberHeap", slices.Clone(TDManberHeap), false},
		{"TestDdata4813", TestDdata4813, true},
		{"TestDdata4813", TestDdata4813, false},
	}

	for _, tt := range tests {
		wantSlice := slices.Clone(tt.seq)
		wantSlice = GetSorted(wantSlice)
		h := NewHeap(tt.seq)
		h.t = t
		if !h.IsMaxHeap() {
			t.Fatalf("NOT A MAX HEAP")
		}
		sortedSeq := h.GetSortedValues(tt.descending)

		t.Logf("want 1: %v", wantSlice)
		if tt.descending {
			slices.Reverse(wantSlice)
			t.Logf("want 2: %v", wantSlice)
		}
		if !slices.Equal(sortedSeq, wantSlice) {
			t.Error("FAIL heapSort")
		}
		n := len(tt.seq)
		ratio := float64(h.iters) / float64(n)
		t.Logf("n: %d, iters: %d, ratio: %.2f", n, h.iters, ratio)
		if n < 20 {
			t.Logf("\nsortedSeq: %v\norigSeq: %v\nh.Imp: %v", sortedSeq, tt.seq, h.Imp)
		}
	}
}

func insertTest(t *testing.T) {
	tests := []struct {
		name string
		seq  []int
	}{
		{"hello", td.Test11},
		// {"hello", td.Test12},
	}

	for _, tt := range tests {
		h := NewHeap(tt.seq)
		h.t = t
		h.PrintTree()
		b := h.IsMaxHeap()
		t.Logf("%+v, IsMaxHeap: %t", h.Imp, b)
	}
	return
	// h := NewHeap([]int{5, 4, 3, 2, 1})
	// h := NewHeap([]int{})
	// h.t = t
	// h.Insert(6)
	// h.Insert(8)
	// h.Insert(10)
	// h.Insert(5)
	// h.PrintTree()
	// b := h.IsMaxHeap()
	// t.Logf("%+v, %t", h.Imp, b)
}

func printTest(t *testing.T) {
	// Example 1: A complete integer heap.
	// The structure should be:
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	h := Heap{Imp: []int{1, 2, 3, 4, 5, 6, 7}, Name: "Complete", t: t}
	h.PrintTree()

	// Example 2: A heap that is not perfectly balanced.
	h.Imp = []int{10, 25, 15, 30}
	h.Name = "Unbalanced Heap"
	h.PrintTree()
}
