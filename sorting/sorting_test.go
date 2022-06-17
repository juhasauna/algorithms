package sorting

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var td testData

func Test_sorting(t *testing.T) {
	td.init()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"binarySearch", binarySearchTest},
		// {"insertionSort", insertionSortTest},
		// {"selectionSort", selectionSortTest},
		// {"merge", mergeTest},
		{"mergeSort", mergeSortTest},
		// {"mergeSortInplace", mergeSortInplaceTest},
		{"mergeSortPointer", mergeSortPointerTest},
		// {"mergePointer", mergePointerTest},
		// {"bubble", bubbleTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func mergeTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		{td.test10, []int{1, 7, 3, 5, 9, 2, 4, 12, 15, 6}},
	}
	for i, tt := range tests {
		x := sorter{}
		l := len(tt.values) / 2
		a := tt.values[:l]
		b := tt.values[l:]
		got := x.merge(a, b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: inputs a=%v, b=%v, expected %v, got %v", i, a, b, tt.expect, got)
		} else {
			fmt.Printf("got: %v\n", got)
		}
	}
}
func mergePointerTest(t *testing.T) {
	getExpect := func(values []int) []int {
		x := sorter{}
		l := len(values) / 2
		a := values[:l]
		b := values[l:]
		return x.merge(a, b)
	}
	tests := []struct {
		values []int
		expect []int
	}{
		{td.test10, getExpect(td.test10)},
		{td.test5, getExpect(td.test5)},
		{td.test3, getExpect(td.test3)},
	}
	for i, tt := range tests {
		x := sorter{}
		got := tt.values
		r := len(got)
		p := 0
		q := r / 2
		x.mergePointer(&got, 0, r/2, r)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: inputs values=%v, p=%d, q=%d,r=%d, expected %v, got %v", i, tt.values, p, q, r, tt.expect, got)
		} else {
			fmt.Printf("got: %v\n", got)
		}
	}
}
func mergeSortPointerTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		// {td.test2, td.getSorted(td.test2)},
		// {td.test3, td.getSorted(td.test3)},
		// {td.test4, td.getSorted(td.test4)},
		// {td.test5, td.getSorted(td.test5)},
		// {td.test7, td.getSorted(td.test7)},
		// {td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		x := sorter{}
		got := make([]int, len(tt.values))
		copy(got, tt.values)
		x.mergeSortPointer(&got, 1, len(tt.values))
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: inputs %v, expected %v, got %v", i, tt.values, tt.expect, got)
		}
	}
}
func mergeSortInplaceTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		// {td.test2, td.getSorted(td.test2)},
		// {td.test3, td.getSorted(td.test3)},
		// {td.test4, td.getSorted(td.test4)},
		// {td.test5, td.getSorted(td.test5)},
		// {td.test7, td.getSorted(td.test7)},
		// {td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values}
		x.mergeSortInplace(1, len(tt.values))
		if !reflect.DeepEqual(x.values, tt.expect) {
			t.Errorf("FAIL %d: inputs %v, expected %v, got %v", i, tt.values, tt.expect, x.values)
		}
	}
}
func bubbleTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		{td.test5, td.getSorted(td.test5)},
		// {td.test10, td.getSorted(td.test10)},
		// {td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		valuesCopy := make([]int, len(tt.values))
		copy(valuesCopy, tt.values)
		x := sorter{values: valuesCopy}
		fmt.Println(x.values, td.test5)
		x.bubble()
		fmt.Println(x.values, td.test5)
		if !reflect.DeepEqual(x.values, tt.expect) {
			t.Errorf("FAIL %d: inputs %v, expected %v, got %v", i, tt.values, tt.expect, x.values)
		}
	}
}
func mergeSortTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		// {td.test5, td.getSorted(td.test5)},
		// {td.test10, td.getSorted(td.test10)},
		// {td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values}
		got := x.mergeSort(x.values)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: inputs %v, expected %v, got %v", i, tt.values, tt.expect, got)
		}
	}
}
func selectionSortTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		// {td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values}
		x.selectionSort()
		if !reflect.DeepEqual(x.values, tt.expect) {
			t.Errorf("FAIL %d: inputs %v, expected %v, got %v", i, tt.values, tt.expect, x.values)
		}
	}
}
func insertionSortTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		// {td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		// fmt.Println(len(tt.values), len(tt.expect))
		// return
		x := sorter{values: tt.values}
		x.insertionSort()
		if !reflect.DeepEqual(x.values, tt.expect) {
			t.Errorf("FAIL %d: inputs %v, expected %v, got %v", i, tt.values, tt.expect, x.values)
		}
	}
}
func binarySearchTest(t *testing.T) {
	tests := []struct {
		values []int
		target int
		expect bool
	}{
		{td.getSorted(td.test10), 7, true},
		{td.getSorted(td.test10), 5, true},
		{td.getSorted(td.test10), 8, false},
	}
	for i, tt := range tests {
		x := binarySearch{values: tt.values, target: tt.target}
		found, iterations := x.search(0, len(tt.values)-1, 0)
		if found != tt.expect {
			t.Errorf("FAIL %d: inputs %v, %d, expected %t, got %t, iterations: %d", i, tt.values, tt.target, tt.expect, found, iterations)
		} else {
			fmt.Printf("SUCCESS %d: iterations: %d\n", i, iterations)
		}
	}
}

type testData struct {
	test2         []int
	test3         []int
	test4         []int
	test5         []int
	test7         []int
	test10        []int
	testDdata4813 []int
	// sorted []int
}

func (x *testData) init() {
	x.test2 = []int{2, 1}
	x.test3 = []int{3, 2, 1}
	x.test4 = []int{4, 3, 2, 1}
	x.test5 = []int{5, 4, 3, 2, 1}
	x.test7 = []int{7, 6, 4, 3, 2, 1}
	x.test10 = []int{7, 3, 5, 15, 6, 1, 9, 2, 4, 12}
	x.testDdata4813 = testDdata4813
}
func (x *testData) getSorted(data []int) []int {
	sortMe := make([]int, len(data))
	copy(sortMe, data)
	sort.Slice(sortMe, func(i, j int) bool {
		return sortMe[i] < sortMe[j]
	})
	return sortMe
}
