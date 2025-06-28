package ntu

import (
	"fmt"
	"log"
	"testing"
)

type CharFreq struct {
	Char      rune
	Frequency int
}

type HuffmanNode struct {
	Data  CharFreq
	Left  *HuffmanNode
	Right *HuffmanNode
}

type Heap struct { // You could user Go's standard priority queue. We're just learning about heaps at the same time so we made our own 'priority queue'.
	Imp   []HuffmanNode
	t     *testing.T
	iters int
	IsMin bool
}

func (x Heap) IsEmpty() bool {
	return x.Size() == 0
}

func NewHeap(seq []CharFreq, isMin bool) Heap {
	initialNodes := []HuffmanNode{}
	for _, v := range seq {
		initialNodes = append(initialNodes, HuffmanNode{Data: v})
	}
	h := Heap{Imp: initialNodes, IsMin: isMin}
	h.Heapify()
	return h
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
	if x.IsMin && !x.IsMinHeap() {
		panic("Heapify() failed: NOT MIN HEAP")
	}
	if !x.IsMin && !x.IsMaxHeap() {
		panic("Heapify() failed: NOT MAX HEAP")
	}
}

func (x *Heap) siftDown(parent, end int) {
	for parent != -1 {
		parent = x.siftHelper(parent, end)
	}
}

func (x *CH6) HuffmanEncoding(seq []CharFreq) HuffmanNode {
	heap := NewHeap(seq, true)
	for {
		left := heap.Extract()
		if heap.Size() == 0 {
			return left
		}
		right := heap.Extract()
		heap.Insert(HuffmanNode{
			Data:  CharFreq{Frequency: left.Data.Frequency + right.Data.Frequency, Char: 0},
			Left:  &left,
			Right: &right,
		})
	}
}

func (x *Heap) Insert(value HuffmanNode) {
	compare := func(a, b int) bool {
		c := x.Imp[a].Data.Frequency < x.Imp[b].Data.Frequency
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
			x.Imp[pi], x.Imp[ci] = x.Imp[ci], x.Imp[pi]
			child = parent
			parent = parent / 2
		} else {
			break
		}
	}
}

func (x *Heap) Extract() HuffmanNode {
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
		log.Fatal("not MIN heap")
	} else if !x.IsMin && !x.IsMaxHeap() {
		log.Fatal("not MAX heap")
	}
}

func (x *Heap) Size() int {
	return len(x.Imp)
}
func (x *Heap) Swap(i, j int) {
	x.Imp[i], x.Imp[j] = x.Imp[j], x.Imp[i]
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
func (x *Heap) Head() HuffmanNode {
	if x.Size() == 0 {
		log.Fatal("Cannot Head() from empty heap")
	}
	return x.Imp[0]
}

func (x Heap) IsMaxHeap() bool {
	n := x.Size()
	for i := 0; i <= (n/2)-1; i++ {
		leftChild := x.LeftChild(i)
		rightChild := x.RightChild(i)
		if leftChild < n && x.Imp[leftChild].Data.Frequency > x.Imp[i].Data.Frequency {
			return false
		}
		if rightChild < n && x.Imp[rightChild].Data.Frequency > x.Imp[i].Data.Frequency {
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
		if leftChild < n && x.Imp[leftChild].Data.Frequency < x.Imp[i].Data.Frequency {
			return false
		}
		if rightChild < n && x.Imp[rightChild].Data.Frequency < x.Imp[i].Data.Frequency {
			return false
		}
	}
	return true
}

func (x *Heap) siftHelper(parent, end int) int {
	compare := func(a, b int) bool {
		c := x.Imp[a].Data.Frequency < x.Imp[b].Data.Frequency
		if x.IsMin {
			c = !c
		}
		return c
	}
	compareEq := func(a, b int) bool {
		c := x.Imp[a].Data.Frequency <= x.Imp[b].Data.Frequency
		if x.IsMin {
			c = !c
		}
		return c
	}
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

func printHuffmanTree(root *HuffmanNode) {
	if root == nil {
		fmt.Println("Tree is empty.")
		return
	}
	// Start the recursive printing from the root node.
	printTreeRecursive(root, "", true)
}

// printTreeRecursive recursively prints the Huffman tree structure.
// 'prefix' is used to build the indentation and connector lines.
// 'isTail' helps determine whether to use a '└──' or '├──' connector.
func printTreeRecursive(node *HuffmanNode, prefix string, isTail bool) {
	// Base case: do nothing if the node is nil.
	if node == nil {
		return
	}

	// Print the connector for the current node.
	// '└──' is used for the last child of a node (the right child in this case).
	// '├──' is used for other children (the left child).
	if isTail {
		fmt.Printf("%s└── ", prefix)
	} else {
		fmt.Printf("%s├── ", prefix)
	}

	// Print the node's data.
	// If the character is 0, it's an internal node, so we only print frequency.
	// Otherwise, it's a leaf node, and we print the character and frequency.
	if node.Data.Char == 0 {
		fmt.Printf("'' %d\n", node.Data.Frequency)
	} else {
		fmt.Printf("'%c' %d\n", node.Data.Char, node.Data.Frequency)
	}

	// Prepare the prefix for the children nodes.
	// If the current node was the tail, the new prefix should have a space.
	// Otherwise, it needs a vertical bar '|' to connect to the next sibling.
	newPrefix := prefix
	if isTail {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	// Recursively print the left and right children.
	// The right child is printed first to get a more intuitive tree layout in the console.
	// We pass 'false' for the left child because it's not the last one at this level.
	// We pass 'true' for the right child because it is the last one.
	printTreeRecursive(node.Left, newPrefix, false)
	printTreeRecursive(node.Right, newPrefix, true)
}

func isValidHuffmanTree(root *HuffmanNode) bool {
	if root == nil {
		return true
	}
	isValid, _ := validateNode(root)
	return isValid
}

func validateNode(node *HuffmanNode) (bool, int) {
	if node.Left == nil && node.Right == nil {
		if node.Data.Char == 0 {
			fmt.Printf("Validation Error: Leaf node with frequency %d has no character.\n", node.Data.Frequency)
			return false, 0
		}
		return true, node.Data.Frequency
	}

	// An internal node must have exactly two children. Having only one is invalid.
	if (node.Left == nil && node.Right != nil) || (node.Left != nil && node.Right == nil) {
		fmt.Printf("Validation Error: Internal node with frequency %d has only one child.\n", node.Data.Frequency)
		return false, 0
	}

	// This is an internal node; validate its properties.
	// A valid internal node must not have a character.
	if node.Data.Char != 0 {
		fmt.Printf("Validation Error: Internal node has a character '%c'.\n", node.Data.Char)
		return false, 0
	}

	// Recursively validate the left and right subtrees.
	isLeftValid, leftFreq := validateNode(node.Left)
	isRightValid, rightFreq := validateNode(node.Right)

	// Check if both subtrees returned as valid.
	if !isLeftValid || !isRightValid {
		return false, 0
	}

	// Check if the parent's frequency is the sum of its children's frequencies.
	expectedFreq := leftFreq + rightFreq
	if node.Data.Frequency != expectedFreq {
		fmt.Printf("Validation Error: Frequency mismatch at internal node. Parent Freq: %d, Expected Sum: %d (Children: %d, %d)\n", node.Data.Frequency, expectedFreq, leftFreq, rightFreq)
		return false, 0
	}

	return true, node.Data.Frequency
}
