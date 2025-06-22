package ut

import (
	"fmt"
	"slices"
	"testing"
)

type Heap struct {
	Imp   []int // Implicit representation.
	Name  string
	t     *testing.T
	iters int
}

func (x Heap) Logf(format string, a ...any) {
	if x.t != nil {
		x.t.Logf(format, a...)
	}
}
func (x *Heap) IncrementIters() {
	if x.t != nil {
		x.iters++
	}
}
func (x Heap) Fatalf(format string, a ...any) {
	if x.t != nil {
		x.t.Fatalf(format, a...)
	}
}
func (x Heap) Fatal(format string) {
	if x.t != nil {
		x.t.Fatalf(format)
	}
}
func (x Heap) Log(a ...any) {
	if x.t != nil {
		x.t.Log(a...)
	}
}

func NewHeap(seq []int) Heap {
	h := Heap{Imp: seq}
	h.Heapify()
	return h
}

func (x *Heap) Head() int {
	if x.Size() == 0 {
		x.Fatal("Cannot Head() from empty heap")
		return -1
	}
	return x.Imp[0]
}
func (x *Heap) Size() int {
	return len(x.Imp)
}
func (x *Heap) Swap(i, j int) {
	x.Imp[i], x.Imp[j] = x.Imp[j], x.Imp[i]
}

func (x *Heap) GetSortedValues(descending bool) []int {
	{
		// Let's keep the original heap intact.
		originalImp := slices.Clone(x.Imp)
		defer func(orig []int) { x.Imp = orig }(originalImp)
	}
	if descending {
		sortedSeq := []int{}
		for x.Size() > 0 {
			sortedSeq = append(sortedSeq, x.Extract())
		}
		return sortedSeq
	}
	return x.SortInplace()
}
func HeapSort(seq []int, descending bool) []int {
	{
		// TODO:Let's keep the original seq intact.
		// originalSeq := slices.Clone(seq)
		// defer func(orig []int) { seq = orig }(originalSeq)
	}
	x := Heap{Imp: seq}
	x.Heapify()
	if descending {
		sortedSeq := []int{}
		for x.Size() > 0 {
			max := x.Extract()
			sortedSeq = append(sortedSeq, max)
		}
		return sortedSeq
	}
	return x.SortInplace()
}

func (x Heap) SortInplace() []int {
	return x.sortHelper(x.Size() - 1)
}

func (x Heap) sortHelper(lastIndex int) []int {
	if lastIndex < 1 {
		return x.Imp
	}

	x.Swap(0, lastIndex)
	lastIndex--
	parent := 0
	for {
		child := x.LeftChild(parent)
		if child > lastIndex {
			break
		}
		if r := x.RightChild(parent); r <= lastIndex {
			if x.Imp[child] < x.Imp[r] {
				child = r
			}
		}
		if x.Imp[child] <= x.Imp[parent] {
			break
		}
		x.Swap(parent, child)
		parent = child
	}
	return x.sortHelper(lastIndex)
}

// Top-down insertion
func (x *Heap) InsertAll(seq []int) {
	for _, v := range seq {
		x.Insert(v)
	}
}

// (x *Heap) Extract() modified from pseudocode.
func (x *Heap) Extract() int {
	if x.Size() == 0 {
		x.t.Fatal("cannot Extract() from empty heap")
	}
	head := x.Head()

	lastIndex := x.Size() - 1
	x.Imp[0] = x.Imp[lastIndex]
	x.Imp = x.Imp[:lastIndex]
	lastIndex--
	// x.Logf("start Extract: head: %d, Imp: %v\n", head, x.Imp)

	parent := 0
	for {
		child := x.LeftChild(parent)
		if child > lastIndex {
			break
		}
		if r := x.RightChild(parent); r <= lastIndex {
			if x.Imp[child] < x.Imp[r] {
				child = r
			}
		}
		if x.Imp[child] <= x.Imp[parent] {
			break
		}
		x.Swap(parent, child)
		parent = child
	}

	if !x.IsMaxHeap() {
		x.Fatal("quit")
	}
	return head
}

func (x *Heap) Insert(value int) {
	child := x.Size() + 1
	x.Imp = append(x.Imp, value)
	parent := child / 2
	for parent > 0 {
		pi := parent - 1
		ci := child - 1
		if x.Imp[pi] < x.Imp[ci] {
			x.Imp[pi], x.Imp[ci] = x.Imp[ci], x.Imp[pi]
			child = parent
			parent = parent / 2
		} else {
			break
		}
	}
}

func (x Heap) printNode(index int, prefix string, isLeft bool) {
	if index >= x.Size() {
		return
	}

	rightChildIndex := x.RightChild(index)
	if rightChildIndex < len(x.Imp) {
		newPrefix := prefix
		if isLeft {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		x.printNode(rightChildIndex, newPrefix, false)
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}
	fmt.Println(x.Imp[index])

	leftChildIndex := x.LeftChild(index)
	if leftChildIndex < len(x.Imp) {
		newPrefix := prefix
		if isLeft {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		x.printNode(leftChildIndex, newPrefix, true)
	}
}
func (x Heap) PrintTree() {
	fmt.Println(x.Name)
	if x.Size() == 0 {
		x.Log("(empty heap)")
		return
	}
	rightChildIndex := x.RightChild(0)
	if rightChildIndex < x.Size() {
		x.printNode(rightChildIndex, "", false)
	}
	fmt.Println(x.Imp[0])
	leftChildIndex := x.LeftChild(0)
	if leftChildIndex < x.Size() {
		x.printNode(leftChildIndex, "", true)
	}
}

func (x Heap) NodeParent(i int) int {
	return (i - 1) / 2
}
func (x Heap) LeftChild(i int) int {
	return 2*i + 1
}
func (x Heap) RightChild(i int) int {
	return 2*i + 2
}
func (x Heap) IsMaxHeap() bool {
	n := x.Size()
	for i := 0; i <= (n/2)-1; i++ {
		leftChild := x.LeftChild(i)
		rightChild := x.RightChild(i)
		if leftChild < n && x.Imp[leftChild] > x.Imp[i] {
			x.Logf("Violation: parent x.Imp[%d](%d) < left child x.Imp[%d](%d)\n", i, x.Imp[i], leftChild, x.Imp[leftChild])
			return false
		}
		if rightChild < n && x.Imp[rightChild] > x.Imp[i] {
			x.Logf("Violation: parent x.Imp[%d](%d) < right child x.Imp[%d](%d)\n", i, x.Imp[i], rightChild, x.Imp[rightChild])
			return false
		}
	}
	return true
}

func (x *Heap) siftDown(parent, end int) {
	leftChild := x.LeftChild(parent)
	for leftChild <= end {
		x.IncrementIters()
		leftChild = x.LeftChild(parent)
		if leftChild > end {
			break
		}
		newParent := leftChild
		rightChild := leftChild + 1
		if rightChild <= end {
			if x.Imp[leftChild] < x.Imp[rightChild] {
				newParent = rightChild
			}
		}
		if x.Imp[newParent] <= x.Imp[parent] {
			break
		}
		x.Swap(parent, newParent)
		x.Logf("parent: %d/%d, swap: %d/%d, heap: %v", parent, x.Imp[parent], newParent, x.Imp[newParent], x.Imp)
		parent = newParent
	}
}

func (x *Heap) Heapify() {
	// O(n) time complexity
	n := x.Size()
	if n <= 1 {
		return
	}

	lastParentNode := (n / 2) - 1
	for i := lastParentNode; i >= 0; i-- {
		x.siftDown(i, n-1)
	}
	if !x.IsMaxHeap() {
		panic("Heapify() failed: NOT MAX HEAP")
	}
}
