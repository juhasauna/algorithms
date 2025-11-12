package ntu

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch5_notes.pdf
import (
	"algorithms/bt"
	"algorithms/ut"
	"slices"
	"strconv"
)

// Cp5/2 Evaluating polynomials
func (x *NTU) evaluatePolySlow(a_i []float64, y float64) float64 {
	x.iters++
	if len(a_i) == 1 {
		return a_i[0]
	}
	var t float64 = 1
	for i := 1; i < len(a_i); i++ {
		t = t * y
	}
	v := a_i[0]
	x.log("%.2f\n", v)
	return v*t + x.evaluatePolySlow(a_i[1:], y)
}
func (x *NTU) hornersRule(a_i []float64, y float64) float64 {
	n := len(a_i)
	p := a_i[0]
	x.log("%.2f\n", p)
	for i := 1; i < n; i++ {
		x.iters++
		p = y*p + a_i[i]
		// p = y*p + a_i[n-i]
		x.log("%.2f\n", p)
	}
	return p
}

// Cp5/4 One-to-One Mapping. (One-to-one = INJECTIVE)
// 2023 HW4/1 Is it possible that the set S will become empty at the end of the algorithm? Show an example, or prove that it cannot happen.
func (x *NTU) cp5MappingBijective(f map[int]int) []int {
	if valid, example := ut.IsFunctionFullyDefined(f); !valid {
		x.t.Errorf("FUNCTION NOT FULLY DEFINED example: %v", *example)
		return nil
	}
	// f is a mapping of a given set to itself
	set := []int{} // Subset of given mapping set, such that (when we're finished) the subset is bijective in given mapping.
	c := make(map[int]int)
	for i, v := range f {
		set = append(set, i)
		set = append(set, v)
		if _, ok := c[i]; !ok {
			c[i] = 0
		}
		c[v]++
	}
	set = ut.Uniquify(set) // Here we should have the original set inherent in the given mapping.
	x.log("c: %v, set: %v\n", c, set)
	queue := []int{}
	for i, v := range c {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		set = ut.RemoveSliceValue(set, head)
		c[f[head]]--
		if c[f[head]] == 0 {
			queue = append(queue, f[head])
		}
	}
	slices.Sort(set)
	return set
}

func (x NTU) degreesGreaterThan(m map[string][]string, k int) (bool, string) {
	for key, v := range m {
		if len(v) < k {
			return false, key
		}
	}
	return true, ""
}

// Cp5 3 Maximal Induced Subgraph (Algorithm Max Ind Subgraph (G, k);)
// Induced means we're removing a node(s) and all the edges to/from those nodes. But we're not allowed to remove other edges.
func (x *NTU) maximalInducedSubgraphRecursive(m map[string][]string, k int) map[string][]string {
	b, key := x.degreesGreaterThan(m, k)
	if b {
		return m
	}
	delete(m, key)
	for key2, v := range m {
		w := ut.RemoveSliceValue(v, key)
		m[key2] = w
	}
	return x.maximalInducedSubgraphRecursive(m, k)
}
func (x *NTU) maximalInducedSubgraphIterative(m map[string][]string, k int) map[string][]string {
	b, key := x.degreesGreaterThan(m, k)
	for !b {
		delete(m, key)
		for key2, v := range m {
			w := ut.RemoveSliceValue(v, key)
			m[key2] = w
		}
		b, key = x.degreesGreaterThan(m, k)
	}
	return m
}

type adjMatrix [][]int

func (x adjMatrix) rowLen() int {
	return len(x)
}
func (x adjMatrix) colLen() int {
	return len(x[0])
}

func (x NTU) newAdjMatrix(s []string) adjMatrix {
	// ss := []string{
	// 	"101",
	// 	"111",
	// 	"000",
	// }
	m := make(adjMatrix, len(s))
	for i, row := range s {
		m[i] = []int{}
		for j, w := range row {
			wInt, err := strconv.Atoi(string(w))
			if err != nil {
				x.t.Fatalf("Invalid character '%c' in input at (%d,%d)", w, i, j)
			}
			m[i] = append(m[i], wInt)
		}
	}
	if !m.validate() {
		x.t.Fatalf("invalid adjMatrix")
	}
	return m

}
func (x adjMatrix) validate() bool {
	if len(x) != len(x[0]) {
		return false
	}
	for _, v := range x {
		for _, w := range v {
			if w != 0 && w != 1 {
				return false
			}
		}
	}
	return true
}

// Cp5/5 Celebrity // O(n)
func (x *NTU) celebrityFromPseudo(m adjMatrix) int {
	n, i, j := m.rowLen(), 0, 1
	for next := 2; next <= n; next++ {
		if m[i][j] == 1 {
			i = next
		} else {
			j = next
		}
	}
	candidate := i
	if i == n {
		candidate = j
	}

	m[candidate][candidate] = 0
	for k := range n {
		if m[candidate][k] == 1 {
			return -1
		}
		if m[k][candidate] == 0 {
			if candidate != k {
				return -1
			}
		}
	}
	return candidate
}

func (x *NTU) celebrityBruteForce(m adjMatrix) map[int]bool {
	// O(n^2)
	n := m.rowLen()
	celebrityMap := make(map[int]bool)
	countMap := make(map[int]int)
	for i := 0; i < n; i++ {
		countMap[i] = 0
		celebrityMap[i] = false
	}
	for _, row := range m {
		for j, cell := range row {
			x.iters++
			countMap[j] += cell
			if countMap[j] == n {
				celebrityMap[j] = true
			}
		}
	}
	// x.log("celeb: %v\n", celebrityMap)
	for i, row := range m {
		iKnows := 0
		for _, cell := range row {
			x.iters++
			iKnows += cell
		}
		if iKnows > 1 {
			celebrityMap[i] = false
		}
	}
	return celebrityMap
}

// Cp5/6 The Skyline Problem
// sk = skyline is an array of alternating x and y coordinates = []sk1{1,2,4,4} = []Coordinate{{1,2}, {4,4}}
func (x *NTU) mergeSkylines(sk1, sk2 []int) {
	n1, n2 := len(sk1), len(sk2)
	if n1 < 2 || n2 < 2 {
		x.t.Fatalf("invalid input")
	}
	if sk1[0] < sk2[0] {
		x.log("%d %d", sk1[0], sk1[1])
		x.mergeSk1(sk1[1], 0, sk1[2:], sk2)
		return
	}
	if sk1[0] > sk2[0] {
		x.log("%d %d", sk2[0], sk2[1])
		x.mergeSk2(0, sk2[1], sk1, sk1[2:])
		return
	}
	if sk1[1] < sk2[0] {
		x.log("%d %d", sk2[0], sk2[1])
		x.mergeSk2(sk1[0], sk2[1], sk1[2:], sk1[2:])
		return
	}

	x.log("%d %d", sk1[0], sk1[1])
	x.mergeSk1(sk1[1], sk2[1], sk1[2:], sk1[2:])
}

func (x *NTU) mergeSk1(ya, yb int, sk1, sk2 []int) {
	n1, n2 := len(sk1), len(sk2)
	if n1 == 1 && n2 == 1 {
		if sk1[0] < sk2[0] {
			x.log("%d, %d,%d", sk1[0], yb, sk2[0])
			return
		}
		x.log("%d", sk1[0])
	}
	if n1 == 1 {
		if sk1[0] < sk2[0] {
			temp := append(sk1, yb)
			temp = append(temp, sk2...)
			x.log("%v", temp)
		}
		return
	}

	if ya >= sk2[1] {
		x.mergeSk1(ya, sk2[1], sk1, sk2[2:])
		return
	} else {
		x.mergeSk2(ya, sk2[1], sk1, sk2[2:])
	}
	// Got bored translating the pseudocode so this is not done
}
func (x *NTU) mergeSk2(ya, yb int, sk1, sk2 []int) {}

// Cp5/7 Balance Factors in Binary Trees
func (x *NTU) computeBalanceFactors(n *bt.Node) {
	if n == nil {
		bt.InitTestData()
		for _, v := range bt.TestNodes {
			n = &v
			break
		}
	}
	n.ComputeBalanceFactors()
}

// Cp5/8 Maximum Consecutive Subsequence // Kadane's algorithm is O(n). Brute force in O(n^3)
func (x *NTU) maximumConsequtiveSubsequence(seq []float64) float64 {
	var globalMax, suffixMax float64
	subSeq, maxSubSeq := []float64{}, []float64{}
	for i, v := range seq {
		suffixMax += v
		if suffixMax > globalMax {
			globalMax = suffixMax
			maxSubSeq = append(maxSubSeq, subSeq...)
			maxSubSeq = append(maxSubSeq, v)
			subSeq = []float64{}
		} else if suffixMax > 0 {
			subSeq = append(subSeq, v)
		} else {
			suffixMax = 0
			subSeq = []float64{}
		}
		x.log("%d, suf: %.2f, glob: %.2f, set: %v\n", i, suffixMax, globalMax, maxSubSeq)
	}
	return globalMax
}
