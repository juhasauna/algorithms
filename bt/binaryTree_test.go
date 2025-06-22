package bt

import (
	"testing"
)

func Test_binaryTree(t *testing.T) {
	InitTestData()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"computeHeights", computeHeightsTest},
		// {"print", printTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func printTest(t *testing.T) {
	tests := []struct {
		name              string
		wantHeight        int
		wantBalanceFactor int
	}{
		// {"h2", 2, -1},
		// {"h4", 4, -3},
		// {"h7", 7, 6},
		// {"justRoot", 0, 0},
	}
	for _, tt := range tests {
		node, ok := TestNodes[tt.name]
		if !ok {
			t.Errorf("FAIL test node '%s' does not exits", tt.name)
			return
		}
		node.ComputeBalanceFactors()
		node.Print(node.Height)

	}

}
func computeHeightsTest(t *testing.T) {
	tests := []struct {
		name              string
		wantHeight        int
		wantBalanceFactor int
	}{
		// {"h2", 2, -1},
		// {"h4", 4, -3},
		{"h7", 7, 6},
		// {"justRoot", 0, 0},
	}
	for _, tt := range tests {
		node, ok := TestNodes[tt.name]
		if !ok {
			t.Errorf("FAIL test node '%s' does not exits", tt.name)
			return
		}
		node.ComputeHeights()
		if node.Height != tt.wantHeight {
			t.Errorf("FAIL gotHeight: %d, wantHeight: %d", node.Height, tt.wantHeight)
		}
		node.ComputeBalanceFactors()
		if node.BalanceFactor != tt.wantBalanceFactor {
			t.Errorf("FAIL gotBalanceFactor: %d, wantBalanceFactor: %d", node.BalanceFactor, tt.wantBalanceFactor)

		}
		// t.Logf("%+v", node)
		node.Print(node.Height)
	}
}
