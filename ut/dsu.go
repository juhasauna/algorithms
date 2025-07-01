package ut

import (
	"fmt"
	"log"
	"slices"
)

// Disjoint Set Union (DSU), is another name for the Union-Find data structure.
// See file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch4_notes.pdf

type DSUNode struct {
	Ancestor int // Some node in the path: parent ---> root.
	Size     int
}

type DSU struct {
	Nodes              []DSUNode
	Balanced           bool
	UsePathCompression bool // Whether to apply path compression in Find
}

func (x DSU) Print() {
	counter := 0
	for i, v := range x.Nodes {
		if v.Ancestor != -2 && v.Size != 0 {
			fmt.Printf("%d: %v\t", i, v)
			if (counter+1)%5 == 0 {
				fmt.Println()
			}
			counter++
		}
	}
	if (counter+1)%5 != 0 {
		fmt.Println()
	}
}

const dsuUnused = -2
const dsuNoParent = -1

func (x *DSU) DSUInit(seq []int) {
	for v := range slices.Max(seq) + 1 {
		if v < 0 {
			log.Fatal("DSUInit: negative sequence values not supported")
		}
		item := DSUNode{Size: 1, Ancestor: dsuNoParent}
		if !slices.Contains(seq, v) {
			item.Size = 0
			item.Ancestor = dsuUnused
		}
		x.Nodes = append(x.Nodes, item)
	}
}

func (x *DSU) Find(a int) int {
	if a < 0 || a >= len(x.Nodes) {
		log.Fatalf("DSU.Find: index %d not in %v", a, x.Nodes)
	}
	if x.Nodes[a].Ancestor == dsuUnused {
		log.Fatalf("DSU.Find: element %d not initialized", a)
	}
	if x.Nodes[a].Ancestor == -1 {
		return a
	}
	parent := x.Nodes[a].Ancestor
	root := x.Find(parent) // Recursive call to find the root.

	if x.UsePathCompression { // We update the ancestor of the incident node to shorten the path to its root. Otherwise traversing the path to discover the "Unioned-set" is slower.
		x.Nodes[a].Ancestor = root
	}
	return root
}
func (x *DSU) Union(a, b int) {
	aRoot := x.Find(a)
	bRoot := x.Find(b)
	if aRoot == bRoot {
		// This is a normal case in DSU usage. You often call Union(a, b) many times in an algorithm (e.g. Kruskal’s MST or grouping), and some of those times, a and b may already be connected.
		fmt.Printf("Redundant Union(%d, %d). a & b have the same parent: %d\n", aRoot, a, b)
		return
	}
	if x.Balanced {
		// In the case of combining two groups of the same size, please always point the second group to the first.
		if x.Nodes[aRoot].Size >= x.Nodes[bRoot].Size {
			x.Nodes[bRoot].Ancestor = aRoot
			x.Nodes[aRoot].Size += x.Nodes[bRoot].Size
			x.Nodes[bRoot].Size = 0 // The node size becomes irrelevant at this point.
		} else {
			x.Nodes[aRoot].Ancestor = bRoot
			x.Nodes[bRoot].Size += x.Nodes[aRoot].Size
			x.Nodes[aRoot].Size = 0
		}
	} else {
		// Point the second group to the firsts ancestor
		x.Nodes[bRoot].Ancestor = aRoot
		x.Nodes[aRoot].Size += x.Nodes[bRoot].Size
		x.Nodes[bRoot].Size = 0
	}
	x.isValid()
}

func (x *DSU) isValid() {
	sizeSum, nodeSum := 0, 0
	for _, v := range x.Nodes {
		if v.Ancestor != dsuUnused {
			sizeSum += v.Size
			nodeSum++
		}
	}
	if nodeSum != sizeSum {
		log.Fatalf("invalid size: %d, total number of nodes is %d", sizeSum, nodeSum)
	}
}

func (dsu *DSU) PrintTree(name string) {
	// Step 1: Build an adjacency list representation of the trees.
	children := make([][]int, len(dsu.Nodes))
	roots := []int{}
	for i, node := range dsu.Nodes {
		if node.Ancestor == dsuUnused {
			continue
		}
		if node.Ancestor == -1 {
			roots = append(roots, i)
		} else {
			children[node.Ancestor] = append(children[node.Ancestor], i)
		}
	}
	fmt.Printf("Tree of %s. ", name)
	// Step 2: Print each tree starting from its root.
	for _, root := range roots {
		fmt.Printf("Set with root %d (Size: %d):\n", root, dsu.Nodes[root].Size)
		dsu.printNode(root, "", true, children)
		fmt.Println()
	}
}

// printNode is a helper function to recursively print the tree structure.
func (dsu *DSU) printNode(nodeID int, prefix string, isLast bool, children [][]int) {
	// Print the current node.
	fmt.Print(prefix)
	if isLast {
		fmt.Print("└── ")
		prefix += "    "
	} else {
		fmt.Print("├── ")
		prefix += "│   "
	}
	fmt.Printf("%d\n", nodeID)

	// Recursively print the children.
	nodeChildren := children[nodeID]
	for i, childID := range nodeChildren {
		isLastChild := (i == len(nodeChildren)-1)
		dsu.printNode(childID, prefix, isLastChild, children)
	}
}
