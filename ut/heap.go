package ut

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

type Heap struct {
	Imp   []int // Implicit representation.
	Name  string
	t     *testing.T
	iters int
	IsMin bool
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

func NewHeap(seq []int, isMin bool) Heap {
	h := Heap{Imp: seq, IsMin: isMin}
	fmt.Printf("New Heap %v,%t\n", seq, isMin)
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
	sorted := x.SortInplace()
	if x.IsMin && !descending {
		slices.Reverse(sorted)
		return sorted
	} else if !x.IsMin && descending {
		slices.Reverse(sorted)
		return sorted
	}
	return sorted
}
func HeapSort(seq []int, descending bool) []int {
	{
		// TODO:Let's keep the original seq intact.
		// originalSeq := slices.Clone(seq)
		// defer func(orig []int) { seq = orig }(originalSeq)
	}
	x := Heap{Imp: seq}
	return x.GetSortedValues(descending)
}

func (x Heap) SortInplace() []int {
	return x.sortHelper(x.Size() - 1)
}

func (x Heap) sortHelper(end int) []int {
	if end < 1 {
		return x.Imp
	}

	x.Swap(0, end)
	end--
	parent := 0
	for parent != -1 {
		parent = x.siftHelper(parent, end)
	}
	return x.sortHelper(end)
}

// Top-down insertion
func (x *Heap) InsertAll(seq []int) {
	for _, v := range seq {
		x.Insert(v)
	}
}

func (x *Heap) Extract() int {
	if x.Size() == 0 {
		x.t.Fatal("cannot Extract() from empty heap")
	}
	head := x.Head()

	end := x.Size() - 1
	x.Imp[0] = x.Imp[end]
	x.Imp = x.Imp[:end]
	end--
	// x.Logf("start Extract: head: %d, Imp: %v\n", head, x.Imp)
	parent := 0

	for parent != -1 {
		parent = x.siftHelper(parent, end)
	}
	x.VerifyHeap()
	return head
}

func (x *Heap) VerifyHeap() {
	if x.IsMin && !x.IsMinHeap() {
		x.PrintTree()
		x.Fatal("not MIN heap")
	} else if !x.IsMin && !x.IsMaxHeap() {
		x.PrintTree()
		x.Fatal("not MAX heap")
	}
}
func (x *Heap) Insert_maxOnly(value int) {
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

func (x *Heap) Insert(value int) {
	compare := func(a, b int) bool {
		c := x.Imp[a] < x.Imp[b]
		if x.IsMin {
			c = !c
		}
		return c
	}
	child := x.Size() + 1
	x.Imp = append(x.Imp, value)
	parent := child / 2
	for parent > 0 {
		pi := parent - 1
		ci := child - 1
		if compare(pi, ci) {
			x.Swap(pi, ci)
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
func (x Heap) IsMinHeap() bool {
	n := x.Size()
	for i := 0; i <= (n/2)-1; i++ {
		leftChild := x.LeftChild(i)
		rightChild := x.RightChild(i)
		if leftChild < n && x.Imp[leftChild] < x.Imp[i] {
			x.Logf("Violation: parent x.Imp[%d](%d) > left child x.Imp[%d](%d)\n", i, x.Imp[i], leftChild, x.Imp[leftChild])
			return false
		}
		if rightChild < n && x.Imp[rightChild] < x.Imp[i] {
			x.Logf("Violation: parent x.Imp[%d](%d) > right child x.Imp[%d](%d)\n", i, x.Imp[i], rightChild, x.Imp[rightChild])
			return false
		}
	}
	return true
}

func (x *Heap) siftHelper(parent, end int) int {
	compare := func(a, b int) bool {
		c := x.Imp[a] < x.Imp[b]
		if x.IsMin {
			c = !c
		}
		return c
	}
	compareEq := func(a, b int) bool {
		c := x.Imp[a] <= x.Imp[b]
		if x.IsMin {
			c = !c
		}
		return c
	}
	x.IncrementIters()
	leftChild := x.LeftChild(parent)
	if leftChild > end {
		return -1
	}
	newParent := leftChild
	rightChild := leftChild + 1
	if rightChild <= end {
		if compare(leftChild, rightChild) {
			newParent = rightChild
		}
	}
	if compareEq(newParent, parent) {
		return -1
	}
	x.Swap(parent, newParent)
	// x.Logf("parent: %d/%d, swap: %d/%d, heap: %v", parent, x.Imp[parent], newParent, x.Imp[newParent], x.Imp)
	return newParent
}
func (x *Heap) siftDown(parent, end int) {
	for parent != -1 {
		parent = x.siftHelper(parent, end)
	}
}

// Bottom-up insertion
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
	if x.IsMin && !x.IsMinHeap() {
		panic("Heapify() failed: NOT MIN HEAP")
	}
	if !x.IsMin && !x.IsMaxHeap() {
		x.PrintTree()
		panic("Heapify() failed: NOT MAX HEAP")
	}
}

type DijkstraNode struct {
	Name   string
	Weight int
	Edge   UndirectedEdge // Using this to keep track of Prim's edges.
}

type DijkstraHeap struct {
	Imp         []*DijkstraNode
	iters       int
	nodeIndexes map[string]int
}

func (x *DijkstraHeap) Heapify() {
	// O(n) time complexity
	n := x.Size()
	if n <= 1 {
		return
	}
	lastParentNode := (n / 2) - 1
	for i := lastParentNode; i >= 0; i-- {
		// x.siftDown(i, n-1)
		x.heapifyDown(i)
	}
	if !x.IsMinHeap() {
		panic("Heapify() failed: NOT MIN HEAP")
	}

}

// func (x *DijkstraHeap) siftHelper(parent, end int) int {
// 	leftChild := x.LeftChild(parent)
// 	if leftChild > end {
// 		return -1
// 	}
// 	newParent := leftChild
// 	rightChild := leftChild + 1
// 	if rightChild <= end {
// 		if x.Imp[leftChild].Weight >= x.Imp[rightChild].Weight {
// 			newParent = rightChild
// 		}
// 	}
// 	if x.Imp[parent].Weight > x.Imp[newParent].Weight {
// 		x.Swap(parent, newParent)
// 	}

//		return newParent
//	}
//
// func (h *DijkstraHeap) siftDown(parent, end int) { // AKA bubble-down, heapifyDown
//
//		for parent != -1 && !h.IsMinHeap() {
//			parent = h.siftHelper(parent, end)
//		}
//	}
func (h *DijkstraHeap) ImpString() string {
	s := ""
	for _, v := range h.Imp {
		s += fmt.Sprintf("%s/%d _ ", v.Name, v.Weight)
	}
	return s
}
func (h *DijkstraHeap) Swap(i, j int) {
	h.nodeIndexes[h.Imp[i].Name], h.nodeIndexes[h.Imp[j].Name] = j, i
	// fmt.Printf("SWAP: %d %d, %v, nodeIndexes: %v\n", i, j, h.ImpString(), h.nodeIndexes)
	h.Imp[i], h.Imp[j] = h.Imp[j], h.Imp[i]
}
func (x *DijkstraHeap) Size() int {
	return len(x.Imp)
}
func (x *DijkstraHeap) IsEmpty() bool {
	return len(x.Imp) == 0
}

func (x DijkstraHeap) IsMinHeap() bool {
	n := x.Size()
	for i := 0; i <= (n/2)-1; i++ {
		leftChild := x.LeftChild(i)
		rightChild := x.RightChild(i)
		if leftChild < n && x.Imp[leftChild].Weight < x.Imp[i].Weight {
			fmt.Printf("Violation: parent x.Imp[%d](%+v) > left child x.Imp[%d](%+v)\n", i, x.Imp[i], leftChild, x.Imp[leftChild])
			return false
		}
		if rightChild < n && x.Imp[rightChild].Weight < x.Imp[i].Weight {
			fmt.Printf("Violation: parent x.Imp[%d](%+v) > right child x.Imp[%d](%+v)\n", i, x.Imp[i], rightChild, x.Imp[rightChild])
			return false
		}
	}
	return true
}

func (x DijkstraHeap) LeftChild(i int) int {
	return 2*i + 1
}
func (x DijkstraHeap) RightChild(i int) int {
	return 2*i + 2
}

func (h *DijkstraHeap) Extract() *DijkstraNode {
	head := h.Head()
	// fmt.Printf("BEF Exctract: %s, head: %s/%d\n", h.ImpString(), head.Name, head.Weight)

	end := h.Size() - 1
	h.Swap(0, end)
	delete(h.nodeIndexes, head.Name)
	h.Imp = h.Imp[:end]
	end--

	h.heapifyDown(0)
	// parent := 0
	// for parent != -1 {
	// 	parent = h.siftHelper(parent, end)
	// }
	// fmt.Printf("AFT Exctract: %s, nodeIndexes: %+v\n", h.ImpString(), h.nodeIndexes)
	if !h.IsMinHeap() {
		panic("DijkstraHeap: not a min heap.")
	}
	return head
}

func (h *DijkstraHeap) Head() *DijkstraNode {
	if h.Size() == 0 {
		log.Fatal("Cannot Head() from empty heap")
		return nil
	}
	return h.Imp[0]
}
func NewDijkstraHeap(seq []*DijkstraNode) DijkstraHeap {
	nodeIndexes := make(map[string]int)
	for i, node := range seq {
		nodeIndexes[node.Name] = i
	}
	h := DijkstraHeap{Imp: seq, nodeIndexes: nodeIndexes}
	// fmt.Printf("BEF Heapify: %v\n", h.nodeIndexes)
	h.Heapify()
	// fmt.Printf("AFT Heapify: %v\n", h.nodeIndexes)
	if !h.IsMinHeap() {
		panic("NewDijkstraHeap not a min heap")
	}
	return h
}

func (h *DijkstraHeap) GetNode(name string) *DijkstraNode {
	if index, ok := h.nodeIndexes[name]; ok {
		return h.Imp[index]
	}
	return nil
}
func (h *DijkstraHeap) DecreaseKey(node *DijkstraNode) {
	index := h.nodeIndexes[node.Name]
	h.Imp[index] = node
	h.heapifyUp(index)
}

func (h *DijkstraHeap) UpdateNode(node *DijkstraNode) {
	if _, ok := h.nodeIndexes[node.Name]; !ok {
		h.Insert(node)
		return
	}
	h.DecreaseKey(node)
}
func (h *DijkstraHeap) Insert(node *DijkstraNode) {
	index := h.Size()
	h.Imp = append(h.Imp, node)
	h.nodeIndexes[node.Name] = index

	// fmt.Printf("Insert: Imp: %s, node: %+v, index: %d\n", h.ImpString(), *node, index)
	h.heapifyUp(index)
}

func (h *DijkstraHeap) heapifyUp(child int) {
	for child > 0 {
		parent := (child - 1) / 2
		if h.Imp[parent].Weight <= h.Imp[child].Weight {
			return
		}
		h.Swap(parent, child)
		child = parent
	}
}

func (h *DijkstraHeap) heapifyDown(parent int) { // AKA bubbleDown, siftDown
	lastIndex := h.Size() - 1
	for {
		leftChild := h.LeftChild(parent)
		if leftChild > lastIndex {
			return
		}
		newParent := leftChild
		rightChild := leftChild + 1
		if rightChild <= lastIndex {
			if h.Imp[leftChild].Weight >= h.Imp[rightChild].Weight {
				newParent = rightChild // Preferring rightChild on equal weights. This should be consistent with heapifyUp.
			}
		}
		if h.Imp[parent].Weight <= h.Imp[newParent].Weight {
			return
		}
		h.Swap(parent, newParent)
		parent = newParent
		// if h.Imp[parent].Weight > h.Imp[newParent].Weight {
		// 	h.Swap(parent, newParent)
		// } else {
		// 	return // If the weights are the same we're done.
		// }
		// parent = newParent
	}
}

func (x DijkstraHeap) PrintTree() {
	if x.Size() == 0 {
		fmt.Println("(empty heap)")
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

func (x DijkstraHeap) printNode(index int, prefix string, isLeft bool) {
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
