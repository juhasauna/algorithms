package ut

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_dsu(t *testing.T) {
	initDSUTests()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"DSU", DSUTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

type Tuple struct {
	a int
	b int
}

func DSUTest(t *testing.T) {
	// tests := []dsuTest {
	// 	name  string
	// 	seq   []int
	// 	union []Tuple
	// }{
	// {"23HW03", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []Tuple{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {1, 4}, {6, 7}, {4, 10}, {10, 8}}},
	// }
	tests := []string{
		// "1",
		// "1_balanced",
		// "1_reverse",
		// "1_balanced_reverse",
		// "2",
		// "2_balanced",
		// "2_path_compression",
		// "2_balanced_&_path_compression",
		"23HW03",
		// "23HW03_balanced",
		// "23HW03_mod_balanced",
		// "23HW03_mod_balanced_path_compression",
	}
	for _, name := range tests {
		tt, ok := dsuTests[name]
		if !ok {
			t.Errorf("test %s does not exist", name)
			continue
		}
		x := DSU{Balanced: tt.balanced, UsePathCompression: tt.usePathCompression}
		x.DSUInit(tt.seq)

		for i, v := range tt.union {
			if v.a > 64 {
				t.Logf("Union: {%c %c}", v.a, v.b)
			} else {
				t.Logf("Union: %v", v)
			}
			x.Union(v.a, v.b)
			usedNodes := []DSUNode{}
			for _, v := range x.Nodes {
				if v.Ancestor != dsuUnused {
					usedNodes = append(usedNodes, v)
				}
			}
			if !slices.Equal(usedNodes, tt.wantNodes[i]) {
				t.Errorf("%s❌ at %d\n%v != \n%v", name, i, usedNodes, tt.wantNodes[i])
				return
			}
		}
		x.PrintTree(name)
		// 	✅
	}
}

type dsuTest struct {
	seq                []int
	union              []Tuple
	wantNodes          [][]DSUNode
	balanced           bool
	usePathCompression bool
}

var dsuTests map[string]dsuTest

func initDSUTests() {
	m := make(map[string]dsuTest)
	seq := []int{1, 2}
	union := []Tuple{{1, 2}}
	unionLen := len(union)
	wantNodes := make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{-1, 2}, {1, 0}},
	}
	m["1"] = dsuTest{seq: seq, union: union, balanced: false, usePathCompression: false, wantNodes: wantNodes}
	// 1_reverse
	seq = []int{1, 2}
	union = []Tuple{{2, 1}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{2, 0}, {-1, 2}},
	}
	m["1_reverse"] = dsuTest{seq: seq, union: union, balanced: false, usePathCompression: false, wantNodes: wantNodes}

	// 1 balanced
	seq = []int{1, 2}
	union = []Tuple{{1, 2}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{-1, 2}, {1, 0}},
	}
	m["1_balanced"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: false, wantNodes: wantNodes}
	// 1_balanced_reverse
	seq = []int{1, 2}
	union = []Tuple{{2, 1}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{2, 0}, {-1, 2}},
	}
	m["1_balanced_reverse"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: false, wantNodes: wantNodes}

	// 2
	seq = []int{1, 2, 3, 4, 5}
	union = []Tuple{{1, 2}, {3, 4}, {4, 1}, {5, 2}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{-1, 2}, {1, 0}, {-1, 1}, {-1, 1}, {-1, 1}},
		{{-1, 2}, {1, 0}, {-1, 2}, {3, 0}, {-1, 1}},
		{{3, 0}, {1, 0}, {-1, 4}, {3, 0}, {-1, 1}},
		{{3, 0}, {1, 0}, {5, 0}, {3, 0}, {-1, 5}},
	}
	m["2"] = dsuTest{seq: seq, union: union, balanced: false, usePathCompression: false, wantNodes: wantNodes}

	// 2_balanced
	seq = []int{1, 2, 3, 4, 5}
	union = []Tuple{{1, 2}, {3, 4}, {4, 1}, {5, 2}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{-1, 2}, {1, 0}, {-1, 1}, {-1, 1}, {-1, 1}},
		{{-1, 2}, {1, 0}, {-1, 2}, {3, 0}, {-1, 1}},
		{{3, 0}, {1, 0}, {-1, 4}, {3, 0}, {-1, 1}},
		{{3, 0}, {1, 0}, {-1, 5}, {3, 0}, {3, 0}},
	}
	m["2_balanced"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: false, wantNodes: wantNodes}

	// 2_path_compression
	seq = []int{1, 2, 3, 4, 5}
	union = []Tuple{{1, 2}, {3, 4}, {4, 1}, {5, 2}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{-1, 2}, {1, 0}, {-1, 1}, {-1, 1}, {-1, 1}},
		{{-1, 2}, {1, 0}, {-1, 2}, {3, 0}, {-1, 1}},
		{{3, 0}, {1, 0}, {-1, 4}, {3, 0}, {-1, 1}},
		{{3, 0}, {3, 0}, {5, 0}, {3, 0}, {-1, 5}},
	}
	m["2_path_compression"] = dsuTest{seq: seq, union: union, balanced: false, usePathCompression: true, wantNodes: wantNodes}

	// 2_balanced_&_path_compression
	seq = []int{1, 2, 3, 4, 5}
	union = []Tuple{{1, 2}, {3, 4}, {4, 1}, {5, 2}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		{{-1, 2}, {1, 0}, {-1, 1}, {-1, 1}, {-1, 1}},
		{{-1, 2}, {1, 0}, {-1, 2}, {3, 0}, {-1, 1}},
		{{3, 0}, {1, 0}, {-1, 4}, {3, 0}, {-1, 1}},
		{{3, 0}, {3, 0}, {-1, 5}, {3, 0}, {3, 0}},
	}
	m["2_balanced_&_path_compression"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: true, wantNodes: wantNodes}

	// 23HW03
	seq = []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	union = []Tuple{{'A', 'B'}, {'C', 'D'}, {'E', 'F'}, {'G', 'H'}, {'I', 'J'}, {'A', 'D'}, {'F', 'G'}, {'D', 'J'}, {'J', 'H'}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		// A		B			C		D		E		F		G			H		I		J
		{{-1, 2}, {'A', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},          // 0
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},         // 1
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},        // 2
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 1}, {-1, 1}},       // 3
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},      // 4 {'I', 'J'}
		{{-1, 4}, {'A', 0}, {'A', 0}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},     // 5 {'A', 'D'}
		{{-1, 4}, {'A', 0}, {'A', 0}, {'C', 0}, {-1, 4}, {'E', 0}, {'E', 0}, {'G', 0}, {-1, 2}, {'I', 0}},    // 6 {'F', 'G'}
		{{-1, 6}, {'A', 0}, {'A', 0}, {'C', 0}, {-1, 4}, {'E', 0}, {'E', 0}, {'G', 0}, {'A', 0}, {'I', 0}},   // 7 {'D', 'J'}
		{{-1, 10}, {'A', 0}, {'A', 0}, {'C', 0}, {'A', 0}, {'E', 0}, {'E', 0}, {'G', 0}, {'A', 0}, {'I', 0}}, // 8 {'J', 'H'}
	}
	m["23HW03"] = dsuTest{seq: seq, union: union, balanced: false, usePathCompression: false, wantNodes: wantNodes}

	// 23HW03_balanced
	seq = []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	union = []Tuple{{'A', 'B'}, {'C', 'D'}, {'E', 'F'}, {'G', 'H'}, {'I', 'J'}, {'A', 'D'}, {'F', 'G'}, {'D', 'J'}, {'J', 'H'}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		// A		B			C		D		E		F		G			H		I		J
		{{-1, 2}, {'A', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},          // 0
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},         // 1
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},        // 2
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 1}, {-1, 1}},       // 3
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},      // 4 {'I', 'J'}
		{{-1, 4}, {'A', 0}, {'A', 0}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},     // 5 {'A', 'D'}
		{{-1, 4}, {'A', 0}, {'A', 0}, {'C', 0}, {-1, 4}, {'E', 0}, {'E', 0}, {'G', 0}, {-1, 2}, {'I', 0}},    // 6 {'F', 'G'}
		{{-1, 6}, {'A', 0}, {'A', 0}, {'C', 0}, {-1, 4}, {'E', 0}, {'E', 0}, {'G', 0}, {'A', 0}, {'I', 0}},   // 7 {'D', 'J'}
		{{-1, 10}, {'A', 0}, {'A', 0}, {'C', 0}, {'A', 0}, {'E', 0}, {'E', 0}, {'G', 0}, {'A', 0}, {'I', 0}}, // 8 {'J', 'H'}
	}
	m["23HW03_balanced"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: false, wantNodes: wantNodes}

	// 23HW03_mod_balanced
	seq = []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	union = []Tuple{{'A', 'B'}, {'C', 'D'}, {'E', 'F'}, {'G', 'H'}, {'I', 'J'}, {'C', 'F'}, {'F', 'G'}, {'D', 'J'}, {'J', 'H'}, {'A', 'C'}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		// A		B			C		D		E		F		G			H		I		J
		{{-1, 2}, {'A', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},          // 0
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},         // 1
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},        // 2
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 1}, {-1, 1}},       // 3
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},      // 4 {'I', 'J'}
		{{-1, 2}, {'A', 0}, {-1, 4}, {'C', 0}, {'C', 0}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},     // 5 {'C', 'F'} <--- this union modified
		{{-1, 2}, {'A', 0}, {-1, 6}, {'C', 0}, {'C', 0}, {'E', 0}, {'C', 0}, {'G', 0}, {-1, 2}, {'I', 0}},    // 6 {'F', 'G'} root1 = fRoot = C, root2 = gRoot = G
		{{-1, 2}, {'A', 0}, {-1, 8}, {'C', 0}, {'C', 0}, {'E', 0}, {'C', 0}, {'G', 0}, {'C', 0}, {'I', 0}},   // 7 {'D', 'J'} root1 = dRoot = C, root2 = jRoot = I
		{{-1, 2}, {'A', 0}, {-1, 8}, {'C', 0}, {'C', 0}, {'E', 0}, {'C', 0}, {'G', 0}, {'C', 0}, {'I', 0}},   // 8 {'J', 'H'} jRoot = I, hRoot = 'G'. This is a noop
		{{'C', 0}, {'A', 0}, {-1, 10}, {'C', 0}, {'C', 0}, {'E', 0}, {'C', 0}, {'G', 0}, {'C', 0}, {'I', 0}}, // 9 {'A', 'C'} <--- 2. modification (new union to test balancing else branch )
	}
	m["23HW03_mod_balanced"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: false, wantNodes: wantNodes}

	// 23HW03_mod_balanced_path_compression
	seq = []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	union = []Tuple{{'A', 'B'}, {'C', 'D'}, {'E', 'F'}, {'G', 'H'}, {'I', 'J'}, {'C', 'F'}, {'F', 'G'}, {'D', 'J'}, {'J', 'H'}, {'A', 'C'}}
	unionLen = len(union)
	wantNodes = make([][]DSUNode, unionLen)
	for i := range unionLen {
		wantNodes[i] = make([]DSUNode, len(seq))
	}
	wantNodes = [][]DSUNode{
		// A		B			C		D		E		F		G			H		I		J
		{{-1, 2}, {'A', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},          // 0
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},         // 1
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 1}, {-1, 1}, {-1, 1}, {-1, 1}},        // 2
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 1}, {-1, 1}},       // 3
		{{-1, 2}, {'A', 0}, {-1, 2}, {'C', 0}, {-1, 2}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},      // 4 {'I', 'J'}
		{{-1, 2}, {'A', 0}, {-1, 4}, {'C', 0}, {'C', 0}, {'E', 0}, {-1, 2}, {'G', 0}, {-1, 2}, {'I', 0}},     // 5 {'C', 'F'} <--- this union modified. cRoot = C, fRoot = E. No path compression activated yet.
		{{-1, 2}, {'A', 0}, {-1, 6}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'G', 0}, {-1, 2}, {'I', 0}},    // 6 {'F', 'G'} fRoot = C, gRoot = G. ACTIVATED path comperession for F.
		{{-1, 2}, {'A', 0}, {-1, 8}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'G', 0}, {'C', 0}, {'I', 0}},   // 7 {'D', 'J'} dRoot = C, jRoot = I
		{{-1, 2}, {'A', 0}, {-1, 8}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}},   // 8 {'J', 'H'} jRoot = I, hRoot = 'G'. ONLY path compression happens. Nice!
		{{'C', 0}, {'A', 0}, {-1, 10}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}, {'C', 0}}, // 9 {'A', 'C'} <--- 2. modification (new union to test balancing else branch )
	}
	m["23HW03_mod_balanced_path_compression"] = dsuTest{seq: seq, union: union, balanced: true, usePathCompression: true, wantNodes: wantNodes}

	dsuTests = m
}
