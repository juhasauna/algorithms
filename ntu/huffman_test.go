package ntu

import (
	"fmt"
	"testing"
)

func Test_huffman(t *testing.T) {
	tests := []struct {
		name string
		seq  []CharFreq
	}{
		// {"hello", []CharFreq{{'D', 3}, {'A', 10}, {'B', 5}, {'E', 2}, {'C', 4}}},
		{"alg2020mid.pdf", []CharFreq{{'A', 18}, {'B', 10}, {'C', 3}, {'D', 8}, {'E', 24}, {'F', 4}, {'G', 10}}},
	}
	for _, tt := range tests {
		x := CH6{}
		got := x.HuffmanEncoding(tt.seq)

		if !isValidHuffmanTree(&got) {
			t.Errorf("invalid Huffman tree")
		}
		// t.Logf("%+v", got)
		printHuffmanTree(&got)
	}
}

func printInOrder(n *HuffmanNode) {
	if n == nil {
		return
	}
	printInOrder(n.Left)
	fmt.Printf("%c\t\t%d\n", n.Data.Char, n.Data.Frequency)
	printInOrder(n.Right)
}
func PrintBreadthFirst(root *HuffmanNode) {
	if root == nil {
		return
	}

	queue := []*HuffmanNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%c\t%d\n", current.Data.Char, current.Data.Frequency) // print rune as character

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
