package bt

import (
	"fmt"
	"strings"
)

type Node struct { // Binary tree Node
	Value         int
	Left          *Node
	Right         *Node
	Height        int
	BalanceFactor int // '-1' means Right sub tree is 1 taller than the Left one.
}

func (x *Node) Print(height int) {
	padding := strings.Repeat("   ", height-x.Height)
	fmt.Printf("%s%+v\n", padding, x)
	if x.Left != nil {
		x.Left.Print(height)
	}
	if x.Right != nil {
		x.Right.Print(height)
	}
}
func (x *Node) ComputeHeights() {
	lh := 0
	if x.Left != nil {
		x.Left.ComputeHeights()
		lh = x.Left.Height + 1
	}
	rh := 0
	if x.Right != nil {
		x.Right.ComputeHeights()
		rh = x.Right.Height + 1
	}
	if lh > rh {
		x.Height = lh
		return
	}
	x.Height = rh
}

func (x *Node) ComputeBalanceFactors() {
	x.ComputeHeights()
	lh := 0
	if x.Left != nil {
		lh = x.Left.Height
		x.Left.ComputeBalanceFactors()
	}
	rh := 0
	if x.Right != nil {
		rh = x.Right.Height
		x.Right.ComputeBalanceFactors()
	}
	x.BalanceFactor = lh - rh
}

var TestNodes map[string]Node

func InitTestData() {
	m := make(map[string]Node)
	m["h2"] = Node{Left: &Node{}, Right: &Node{Right: &Node{}}}
	m["h4"] = Node{Left: &Node{}, Right: &Node{Right: &Node{Right: &Node{Right: &Node{}}}}}
	m["h7"] = Node{Left: &Node{Value: 90, Left: &Node{Value: 68, Right: &Node{Value: 83, Left: &Node{Value: 56, Left: &Node{Value: 31, Left: &Node{Value: 5, Left: &Node{Value: 82}, Right: &Node{Value: 98}}, Right: &Node{Value: 2, Left: &Node{Value: 90}}}}}}}}

	m["justRoot"] = Node{}
	TestNodes = m
}
