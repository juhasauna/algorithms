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
		{"insert", insertTest},
		// {"heapSort", heapSortTest},
		// {"print", printTest},
		{"heapify", heapifyTest},
		// {"IsMaxHeap", IsMaxHeapTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func makeWorstBestCaseInputArray(count int) []int {
	result := []int{}
	for i := range count {
		result = append(result, i)
	}
	return result
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
		name  string
		seq   []int
		isMin bool
	}{
		// {"Test11", td.Test11, false},
		// {"Test11", td.Test11, true},
		// {"TDManberHeap", TDManberHeap, false},
		// {"TDManberHeap", TDManberHeap, true},
		// {"JuhaHeapBigger", JuhaHeapBigger, true},
		// {"TestDdata4813", TestDdata4813, true},
		// {"worstCase", makeWorstBestCaseInputArray(100), false},
		// {"worstCase", makeWorstBestCaseInputArray(100), false},
		{"NTU2024mid7", NTU2024mid7, false},
	}

	for _, tt := range tests {
		h := Heap{Imp: tt.seq, t: t, IsMin: tt.isMin, countSwaps: true}
		h.Heapify()
		t.Logf("Heapify iters: %d", h.iters)
		if h.IsMin && !h.IsMinHeap() {
			t.Fatalf("NOT MIN HEAP")
		} else if !h.IsMin && !h.IsMaxHeap() {
			t.Fatalf("NOT MAX HEAP")
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
		if tt.descending {
			slices.Reverse(wantSlice)
		}
		h := NewHeap(tt.seq, true)
		h.t = t
		h.VerifyHeap()
		sortedSeq := h.GetSortedValues(tt.descending)

		if !slices.Equal(sortedSeq, wantSlice) {
			t.Errorf("FAIL h.GetSortedValues, got: %v !=  want: %v", wantSlice, sortedSeq)
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
		name  string
		seq   []int
		isMin bool
	}{
		// {"hello", td.Test11, false},
		// {"hello", td.Test11, true},
		// {"TDManberHeap", TDManberHeap, true},
		// {"TDManberHeap", TDManberHeap, false},
		// {"JuhaHeapBigger", JuhaHeapBigger, true},
		// {"TestDdata4813", TestDdata4813, true},
		{"worstCase", makeWorstBestCaseInputArray(100), false},
	}

	for _, tt := range tests {
		h := Heap{IsMin: tt.isMin, t: t, countSwaps: true}
		h.InsertAll(tt.seq)
		t.Logf("InsertAll iters: %d", h.iters)
		// h.PrintTree()
		ok := false
		if tt.isMin {
			ok = h.IsMinHeap()
		} else {
			ok = h.IsMaxHeap()
		}
		if !ok {
			t.Errorf("FAIL %s", tt.name)
		}
		// t.Logf("%+v, IsMaxHeap: %t, IsMinHeap: %t", h.Imp)
	}
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
