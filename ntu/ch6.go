package ntu

// Chapter 6 From Manber file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/Me/Books/CS/Manber%20-%20Introduction%20to%20Algorithms_%20A%20Creative%20Approach%201989_v2.pdf
// And file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch6_notes_a.pdf
// And Home work 6 (for 2024 file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw6.pdf)

import (
	"algorithms/ut"
	"math"
	"slices"
	"testing"
)

// FindMinAndMax() find min and max in [1.5*n-2]-comparisons. Using a the naive approach takes 2*n-3 comparisons.
func (x *CH6) FindMinAndMax(seq []int) (int, int) {
	n := len(seq)
	switch n {
	case 0:
		return 0, 0
	case 1:
		return seq[0], seq[0]
	}
	max, min := 0, 0
	x.iters++
	if seq[0] > seq[1] {
		max, min = seq[0], seq[1]
	} else {
		max, min = seq[1], seq[0]
	}
	if n == 2 {
		return max, min
	}
	for i := 2; i < n-1; i += 2 {
		a_, b_ := seq[i], seq[i+1]
		x.iters++
		x.iters++
		x.iters++
		if b_ > a_ {
			max = ut.Max(b_, max)
			min = ut.Min(a_, min)
		} else {
			max = ut.Max(a_, max)
			min = ut.Min(b_, min)
		}
	}
	if n%2 != 0 {
		min, max = ut.NewMinAndMax(min, max, seq[n-1])
	}
	return min, max
}

func (x *CH6) KthSmallestElement(seq []int, k int) int {
	// O(n)
	// input gets mutated
	// BONUS EX. Figure out how to also return the index of the smalles element (of the original input sequence). Gemini says that there is no simple way to do this.
	x.Logf("init\t\t\t%d\t\t%v", k, seq)
	var selectk func(left, right, kth int, msg string) int
	selectk = func(left, right, kth int, msg string) int {
		x.iters++
		if x.iters > 10 {
			x.t.Fatal("overflow")
		}
		if left == right {
			return left
		}
		mid := x.Partition(seq, left, right)

		x.Logf("%s: %d %d %d\t\t%d\t\t%v", msg, left, mid, right, kth, seq)
		if mid-left+1 >= kth {
			return selectk(left, mid, kth, " left")
		}
		return selectk(mid+1, right, kth-(mid-left+1), "right")
	}

	n := len(seq)
	if k < 0 || k > n {
		return -1
	}
	index := selectk(0, n-1, k, "start")
	result := seq[index]
	return result
}

// 1. Select pivot.
// 2. Move elements to the left and right of the pivot based on the size of the elements relative to the pivot.
// 3. Return the position of the pivot.
func (x CH6) Partition(seq []int, left, right int) int {
	pivot := seq[left] // Pivot can be any value from the input sequence.
	// pivot := seq[right]
	l, r := left, right
	for l < r {
		for seq[l] <= pivot && l < r {
			l++
		}
		for seq[r] > pivot && r >= l {
			r--
		}
		if l < r {
			seq[l], seq[r] = seq[r], seq[l]
		}
	}
	mid := r
	seq[left], seq[mid] = seq[mid], seq[left]
	return mid
}

// Implementation of pseudocode: file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch6_notes_a.pdf
func (x *CH6) mergeSortPseudo(seq []int, n int) []int {
	var m_sort func(int, int)
	m_sort = func(left, right int) {
		x.iters++
		if right-left == 1 {
			if seq[left] > seq[right] {
				seq[left], seq[right] = seq[right], seq[left]
				x.t.Logf("SW: %v\t", seq)
			}
		} else if seq[left] != seq[right] {
			middle := (left + right + 1) / 2

			ut.AssertEqual(middle, int(math.Ceil(float64(left+right)/2.0)))

			x.t.Logf("LEFT %d, M%d\n", x.iters, middle)
			m_sort(left, middle-1)
			x.t.Logf("RIGHT %d, M%d\n", x.iters, middle)
			m_sort(middle, right)
			x.t.Logf("\n---START MERGE iter:%d, L%d, M%d, R%d\n", x.iters, left, middle, right)

			// This is a PARTICULARLY DIFFICULT to read implementation of the merge part. See sorter.merge() for a better version.
			temp := make([]int, n) // Alternative: temp := make(map[int]int)
			i, j, k := left, middle, 0
			for i < middle && j <= right {
				if seq[i] <= seq[j] {
					x.t.Logf("IF i%d, j%d, k%d\t", i, j, k)
					temp[k] = seq[i]
					i++
				} else {
					x.t.Logf("ELSE i%d, j%d, k%d\t", i, j, k)
					temp[k] = seq[j]
					j++
				}
				x.t.Logf("\nTEMP %v\n", temp)
				k++
			}
			if j > right {
				for t := 0; t < (middle - i); t++ {
					seq[right-t] = seq[middle-1-t]
				}
				x.t.Logf("M: %d, i: %d, X: %v\t", i, middle, seq)
			}
			for t := 0; t < k; t++ {
				seq[left+t] = temp[t]
			}
			x.t.Logf("L: %d, k: %d, X: %v\t", left, middle, seq)
			x.t.Logf("\n---END MERGE\n")
		}

	}
	m_sort(0, n-1)
	return seq
}

// WIP: Improved version of NTU pseudocode translation
func (x *CH6) mergeSortImproved(seq []int) []int {
	merge := func(left, mid, right int) {
		temp := make([]int, right)
		i, j, k := left, mid, 0
		for i < mid && j <= right {
			if seq[i] <= seq[j] {
				x.t.Logf("IF i%d, j%d, k%d\t", i, j, k)
				temp[k] = seq[i]
				i++
			} else {
				x.t.Logf("ELSE i%d, j%d, k%d\t", i, j, k)
				temp[k] = seq[j]
				j++
			}
			x.t.Logf("\nTEMP %v\n", temp)
			k++
		}
		if j > right {
			for t := 0; t < (mid - i); t++ {
				seq[right-t] = seq[mid-1-t]
			}
			x.t.Logf("M: %d, i: %d, X: %v\t", i, mid, seq)
		}
		for t := 0; t < k; t++ {
			seq[left+t] = temp[t]
		}
		x.t.Logf("L: %d, k: %d, X: %v\t", left, mid, seq)
	}
	var sort func(int, int, string)
	sort = func(left, right int, msg string) {
		x.iters++
		if right-left == 1 {
			if seq[left] > seq[right] {
				seq[left], seq[right] = seq[right], seq[left]
				x.t.Logf("SW: %v\t", seq)
			}
		} else if seq[left] != seq[right] {
			mid := (left + right + 1) / 2
			x.t.Logf("%s, iters: %d, mid: %d\n", msg, x.iters, mid)
			sort(left, mid-1, "left")
			sort(mid, right, "right")
			x.t.Logf("\n---START MERGE iter:%d, L%d, M%d, R%d\n", x.iters, left, mid, right)
			merge(left, mid, right)
			x.t.Logf("\n---END MERGE\n")
		}

	}
	sort(0, len(seq)-1, "start")
	return seq
}

func (x *CH6) straightRadixSort(seq []int) []int {
	getDigit := func(num, place int) int {
		return (num / place) % 10
	}
	if len(seq) < 2 {
		return seq
	}
	digits := func() int {
		maxVal := slices.Max(seq)
		return int(math.Floor(math.Log10(float64(maxVal))) + 1)
	}()
	place := 1
	for i := 0; i < digits; i++ {
		queues := make([][]int, 10)
		for j := 0; j < 10; j++ {
			queues[j] = []int{}
		}
		for _, v := range seq {
			d := getDigit(v, place)
			queues[d] = append(queues[d], v)
		}
		seq = []int{}
		for j := 0; j < 10; j++ {
			seq = append(seq, queues[j]...)
		}
		place *= 10
	}
	return seq
}

func (x *CH6) interpolationSearch(sortedSeq []int, target int) int {
	n := len(sortedSeq)
	if n == 0 {
		x.t.Fatalf("invalid input, n: %d", n)
	}
	if sortedSeq[0] > target || sortedSeq[n-1] < target {
		return -1
	}
	var find func(left, right int, msg string) int
	find = func(left int, right int, msg string) int {
		x.iters++
		if x.iters > 5 {
			x.t.Fatalf("too many iters: %d", x.iters)
		}
		if sortedSeq[left] == target {
			return left
		} else if left == right {
			return -1
		}
		var guess int = (target - sortedSeq[left]) * (right - left)
		x.t.Logf("%d", guess)
		var tempGuess float64 = float64(guess) / float64((sortedSeq[right] - sortedSeq[left]))
		x.t.Logf("%.2f\t", tempGuess)

		if guess == 0 {
			guess = left
		} else {
			guess = int(math.Ceil(tempGuess)) + left
		}
		x.t.Log(guess, "\t")
		x.t.Logf("L: %d G: %d, R: %d, msg: %s\n", left, guess, right, msg)
		if sortedSeq[guess] > target {
			return find(left, guess-1, "left")
		}
		return find(guess, right, "right")
	}
	return find(0, n-1, "begin")
}

func (x *CH6) canMakeStutter(text, pattern string, k int) bool {
	if k == 0 {
		return true
	}
	if len(pattern)*k > len(text) {
		return false
	}
	textIndex := 0
	for _, charToFind := range pattern {
		count := 0
		for textIndex < len(text) && count < k {
			if rune(text[textIndex]) == charToFind {
				count++
			}
			textIndex++
		}
		if count < k {
			return false
		}
	}
	return true
}

// StutteringSubsequence finds the maximum k such that pattern_k is a subsequence of text.
func (x *CH6) stutteringSubsequence(text, pattern string) int {
	if len(pattern) == 0 {
		return 1e9
	}
	low := 0
	high := len(text) / len(pattern)
	maxK := 0
	for low <= high {
		mid := low + (high-low)/2
		if mid == 0 {
			low = mid + 1
			continue
		}
		if x.canMakeStutter(text, pattern, mid) {
			maxK = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return maxK
}

// Given a sorted sequence of distinct integers determine if there exists an index i such that a_i = i.
func (x *CH6) specialBinarySearch(sortedSeq []int) int {
	var find func(left, right int, msg string) int
	find = func(left int, right int, msg string) int {
		x.Log(left, right, msg)
		x.iters++
		if x.iters > 5 {
			x.t.Fatalf("too many iters. left: %d, right:%d", left, right)
		}
		if left == right {
			if sortedSeq[left] == left {
				return left
			}
			return -1
		}
		mid := (left + right) / 2
		if sortedSeq[mid] < mid {
			return find(mid+1, right, "right")
		}
		return find(left, mid, "left")
	}
	return find(0, len(sortedSeq)-1, "begin")
}

// cyclicBinarySearch find the minimal index from a cyclically sorted seqeuence.
func (x *CH6) cyclicBinarySearch(sortedSeq []int) int {
	var find func(l, r int) int
	find = func(l, r int) int {
		if x.iters > 10 {
			x.t.Fatal("too many iters")
		}
		x.iters++
		if l == r {
			return l
		}
		mid := (l + r) / 2
		x.t.Log(l, mid, r)
		if sortedSeq[mid] < sortedSeq[r] {
			return find(l, mid)
		}
		return find(mid+1, r)
	}
	return find(0, len(sortedSeq)-1)
}

func (x *CH6) quickSort(seq []int) []int {
	partition := func(left, right int) int {
		pivot := seq[left]
		l, r := left, right
		for l < r {
			for seq[l] <= pivot && l < r {
				l++
			}
			for seq[r] > pivot && r >= l {
				r--
			}
			if l < r {
				seq[l], seq[r] = seq[r], seq[l]
			}
		}
		mid := r
		seq[left], seq[mid] = seq[mid], seq[left]
		return mid
	}
	var sortition func(left, right int, msg string)
	sortition = func(left, right int, msg string) {
		x.iters++
		if x.iters > 100 {
			x.t.Fatalf("too many iters")
		}
		if left < right {
			mid := 0
			mid = partition(left, right)
			x.t.Logf("%d, %d, %d, %v, msg: %s\n", left, mid, right, seq, msg)
			sortition(left, mid-1, "left")
			sortition(mid+1, right, "right")
		}
	}
	sortition(0, len(seq)-1, "begin")
	return seq
}

type CH6 struct {
	verbose bool
	iters   int64
	t       *testing.T
}

func (x CH6) Logf(format string, a ...any) {
	if x.verbose {
		x.t.Logf(format, a...)
	}
}
func (x CH6) Log(a ...any) {
	if x.verbose {
		x.t.Log(a...)
	}
}
