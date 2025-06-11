package sorting

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var td testData

func Test_sorting(t *testing.T) {
	td.init()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"separateIntThousands", separateIntThousandsTest},
		// {"binarySearch", binarySearchTest},
		// {"insertionSort", insertionSortTest},
		// {"selectionSort", selectionSortTest},
		// {"mergeSort", mergeSortTest},
		// {"mergeSortNTU", mergeSortNTUTest},
		// {"bubble", bubbleTest},
		// {"merge", mergeTest},
		{"ternarySearch", ternarySearchTest},
		// {"mergeSortInplace", mergeSortInplaceTest},
		// {"mergeSortPointer", mergeSortPointerTest},
		// {"mergePointer", mergePointerTest},
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
		// {td.testDdata4813, td.getSorted(td.testDdata4813)},
		{[]int{8, 2, 4, 6, 9, 7, 10, 1, 5, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values, verbose: true}
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
		// {td.test5, td.getSorted(td.test5)},
		// {td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		valuesCopy := make([]int, len(tt.values))
		copy(valuesCopy, tt.values)
		x := sorter{values: valuesCopy}
		x.bubble()
		if !reflect.DeepEqual(x.values, tt.expect) {
			t.Errorf("FAIL %d: iters: %s, inputs %v, expected %v, got %v", i, x.itersFormat(), tt.values, tt.expect, x.values)
		} else {
			fmt.Printf("SUCCESS bubble iters: %s\n", x.itersFormat())
		}
	}
}
func mergeSortNTUTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		// {td.test3, td.getSorted(td.test3)},
		// {td.test4, td.getSorted(td.test4)},
		// {td.test6, td.getSorted(td.test6)},
		{td.testNTUalg2022mid_6, td.getSorted(td.testNTUalg2022mid_6)},
		// {td.test7, td.getSorted(td.test7)},
		// {td.test8, td.getSorted(td.test8)},
		// {td.test10, td.getSorted(td.test10)},
		// {td.testDdata4813, td.getSorted(td.testDdata4813)},
		// {[]int{8, 2, 4, 6, 9, 7, 10, 1, 5, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		// {[]int{8, 3, 2, 6, 5, 9, 10, 7, 1, 12, 4, 11}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values, verbose: true}
		got := x.mergeSortNTU(x.values, len(x.values))
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: iters: %s, inputs %v, expected %v, got %v", i, x.itersFormat(), tt.values, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS merge iters: %s\n", x.itersFormat())
		}
	}
}
func mergeSortTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		{td.test5, td.getSorted(td.test5)},
		{td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
		{[]int{8, 2, 4, 6, 9, 7, 10, 1, 5, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values, verbose: false}
		got := x.mergeSort(x.values)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: iters: %s, inputs %v, expected %v, got %v", i, x.itersFormat(), tt.values, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS merge iters: %s\n", x.itersFormat())
		}
	}
}
func selectionSortTest(t *testing.T) {
	tests := []struct {
		values []int
		expect []int
	}{
		{td.test10, td.getSorted(td.test10)},
		{td.testDdata4813, td.getSorted(td.testDdata4813)},
	}
	for i, tt := range tests {
		x := sorter{values: tt.values}
		x.selectionSort()
		if !reflect.DeepEqual(x.values, tt.expect) {
			t.Errorf("FAIL %d: iters: %s, inputs %v, expected %v, got %v", i, x.itersFormat(), tt.values, tt.expect, x.values)
		} else {
			fmt.Printf("SUCCESS selection iters: %s\n", x.itersFormat())
		}
	}
}
func separateIntThousandsTest(t *testing.T) {
	fmt.Println(separateIntThousands(123456789))
	fmt.Println(separateIntThousands(12345678))
	fmt.Println(separateIntThousands(1234567))
}

func separateIntThousands_my_feeble_attempt(i int) string {
	slice := strings.Split(strconv.Itoa(i), "")
	s := ""
	j := len(slice)
	prev := j
	for {
		j -= 3
		s = strings.Join(slice[j:prev], "") + " " + s
		prev = j
		if prev-3 < 0 {
			break
		}
	}
	leftOver := len(slice) % 3
	if leftOver > 0 {
		s += strings.Join(slice[:leftOver], "")
	}
	return s
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
			t.Errorf("FAIL %d: iters: %s, inputs %v, expected %v, got %v", i, x.itersFormat(), tt.values, tt.expect, x.values)
		} else {
			fmt.Printf("SUCCESS insertion iters: %s\n", x.itersFormat())
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
		x := searcher{values: tt.values, target: tt.target}
		found, iterations := x.binarySearch(0, len(tt.values)-1, 0)
		if found != tt.expect {
			t.Errorf("FAIL %d: inputs %v, %d, expected %t, got %t, iterations: %d", i, tt.values, tt.target, tt.expect, found, iterations)
		} else {
			fmt.Printf("SUCCESS %d: iterations: %d\n", i, iterations)
		}
	}
}
func ternarySearchTest(t *testing.T) {
	tests := []struct {
		values []int
		target int
		expect int
	}{
		{[]int{10, 20, 30}, 1, -1},
		{[]int{10, 20, 30}, 10, 0},
		{[]int{10, 20, 30}, 20, 1},
		{[]int{10, 20, 30}, 30, 2},
		{[]int{10, 20, 30}, 31, -1},
		{[]int{10, 20, 30, 40}, 40, 3},
		{[]int{10, 20, 30, 40, 50}, 50, 4},
		{[]int{10, 20, 30, 40, 50, 60}, 10, 0},
		{[]int{11, 22, 33, 44, 55, 66}, 44, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 6, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, 7},
		{td.getSorted(td.test10), 1, 0},
		{td.getSorted(td.test10), 2, 1},
		{td.getSorted(td.test10), 3, 2},
		{td.getSorted(td.test10), 4, 3},
		{td.getSorted(td.test10), 5, 4},
		{td.getSorted(td.test10), 6, 5},
		{td.getSorted(td.test10), 7, 6},
		{td.getSorted(td.test10), 9, 7},
		{td.getSorted(td.test10), 15, 9},
		{td.getSorted(td.test9), 5, 4},
	}
	for i, tt := range tests {
		x := searcher{values: tt.values, target: tt.target}
		found, iterations := x.ternarySearch(tt.target, tt.values)
		if found != tt.expect {
			t.Errorf("FAIL %d: inputs %v, %d, expected %d, got %d, iterations: %d", i, tt.values, tt.target, tt.expect, found, iterations)
		} else {
			fmt.Printf("SUCCESS %d: iterations: %d\n", i, iterations)
		}
	}
}

type testData struct {
	test2               []int
	test3               []int
	test4               []int
	test5               []int
	test6               []int
	test7               []int
	test8               []int
	test9               []int
	test10              []int
	testDdata4813       []int
	testNTUalg2022mid_6 []int
	// sorted []int
}

func (x *testData) init() {
	x.test2 = []int{2, 1}
	x.test3 = []int{3, 2, 1}
	x.test4 = []int{4, 3, 2, 1}
	x.test5 = []int{5, 4, 3, 2, 1}
	x.test6 = []int{6, 5, 4, 3, 2, 1}
	x.test7 = []int{7, 6, 4, 3, 2, 1}
	x.test8 = []int{5, 4, 3, 2, 1, 0}
	x.test9 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x.test10 = []int{7, 3, 5, 15, 6, 1, 9, 2, 4, 12}
	x.testDdata4813 = TestDdata4813
	x.testNTUalg2022mid_6 = []int{8, 3, 2, 6, 5, 9, 10, 7, 1, 12, 4, 11}
}
func (x *testData) getSorted(data []int) []int {
	sortMe := make([]int, len(data))
	copy(sortMe, data)
	sort.Slice(sortMe, func(i, j int) bool {
		return sortMe[i] < sortMe[j]
	})
	return sortMe
}
