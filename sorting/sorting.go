package sorting

import (
	"fmt"
	"math"
	"strconv"
)

type binarySearch struct {
	target int
	values []int
}

func (x binarySearch) search(lo, hi, counter int) (bool, int) {
	if lo <= hi {
		mid := lo + ((hi - lo) / 2)
		cmp := x.target - x.values[mid]
		if cmp == 0 {
			return true, counter + 1
		} else if cmp < 0 {
			return x.search(lo, mid-1, counter+1)
		} else {
			return x.search(mid+1, hi, counter+1)
		}
	}
	return false, counter
}

type sorter struct {
	values  []int
	verbose bool
	iters   int64
}

func (x *sorter) itersFormat() string {
	return separateIntThousands(x.iters)
}

func (x sorter) log(format string, a ...interface{}) {
	if x.verbose {
		fmt.Printf(format, a...)
	}
}

func (x *sorter) merge(a, b []int) []int {
	result := []int{}
	for len(a) > 0 && len(b) > 0 {
		x.iters++
		if a[0] < b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}
	result = append(result, append(a, b...)...)
	return result
}

func (x *sorter) mergeSort(values []int) []int {
	x.iters++
	l := len(values) / 2
	a := values[:l]
	b := values[l:]
	if l < 1 {
		return x.merge(a, b)
	}
	return x.merge(x.mergeSort(a), x.mergeSort(b))
}

func (x *sorter) mergePointer(A *[]int, p, q, r int) {
	tempA := *A
	left := make([]int, len(tempA[p:q]))
	copy(left, tempA[p:q])
	left = append(left, math.MaxInt)
	right := make([]int, len(tempA[q:r]))
	copy(right, tempA[q:r])
	right = append(right, math.MaxInt)
	x.log("merge tempA=%v, left=%v, right=%v, p=%d, q=%d, r=%d\n", tempA, left, right, p, q, r)

	i := 0
	j := 0
	for k := p; k < r; k++ {
		x.log("i=%d, j=%d, k=%d, A=%v, l=%d, r=%d\n", i, j, k, tempA, left[i], right[j])
		if left[i] <= right[j] {
			tempA[k] = left[i]
			i++
		} else {
			tempA[k] = right[j]
			j++
		}
	}
	A = &tempA
}

func (x *sorter) mergeSortPointer(A *[]int, p, r int) {
	// Adapted from the book Introcution to algorithms.
	x.log("start A: %v, p=%d, r=%d\n", *A, p, r)
	if p < r {
		q := (p + r) / 2
		x.mergeSortPointer(A, p, q)
		x.mergeSortPointer(A, q+1, r)
		x.mergePointer(A, p-1, q, r)
	}
	x.log("end A: %v\n", A)
}
func (x *sorter) mergeInplace(p, q, r int, msg ...string) {
	lLen := q - p + 1
	left := make([]int, lLen)
	copy(left, x.values[p:q])
	left[lLen-1] = math.MaxInt
	rLen := r - q + 1
	right := make([]int, rLen)
	copy(right, x.values[q:r])
	right[rLen-1] = math.MaxInt
	x.log("merge %s, A=%v, left=%v, right=%v, p=%d, q=%d, r=%d, lLen=%d, rLen=%d\n", msg, x.values, left, right, p, q, r, lLen, rLen)

	i := 0
	j := 0
	for k := p; k < r; k++ {
		x.log("i=%d, j=%d, k=%d, A=%v, l=%d, r=%d\n", i, j, k, x.values, left[i], right[j])
		if left[i] <= right[j] {
			x.values[k] = left[i]
			i++
		} else {
			x.values[k] = right[j]
			j++
		}
	}
}

func (x *sorter) mergeSortInplace(p, r int, msg ...string) {
	// Revised version from the book Introcution to algorithms.
	if p < r {
		q := (p + r) / 2
		x.log("start %s A: %v, p=%d, q=%d, r=%d\n", msg, x.values, p, q, r)
		x.mergeSortInplace(p, q, "left")
		x.mergeSortInplace(q+1, r, "right")
		x.mergeInplace(p-1, q, r, msg...)
		x.log("end A: %v\n", x.values)
	}
}
func (x *sorter) insertionSort() {
	for j := 1; j < len(x.values); j++ {
		key := x.values[j]
		i := j - 1
		for i >= 0 && x.values[i] > key {
			x.iters++
			x.values[i+1] = x.values[i]
			i--
		}
		x.values[i+1] = key
	}
}
func (x *sorter) selectionSort() {
	for i, v := range x.values {
		smallestIndex := i
		for j := i; j < len(x.values); j++ {
			x.iters++
			if v > x.values[j] {
				v = x.values[j]
				smallestIndex = j
			}
		}
		temp := x.values[i]
		x.values[i] = v
		x.values[smallestIndex] = temp
	}
}
func (x *sorter) bubble() {
	for i := 0; i < len(x.values); i++ {
		for j := len(x.values) - 1; j > i; j-- {
			x.iters++
			if x.values[j-1] > x.values[j] {
				temp := x.values[j-1]
				x.values[j-1] = x.values[j]
				x.values[j] = temp
			}
		}
	}
}

func separateIntThousands(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ' '
		}
	}
}
