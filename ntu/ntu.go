package ntu

import (
	"algorithms/ut"
	"fmt"
	"log"
	"slices"
	"testing"
)

// Ex.1 from file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw4.pdf
// Given a sorted array A of n integers and an integer x, determine whether A contains two integers whose sum is exactly x.
func (x NTU) hw24_04_01_polyTime(arr []int, targetSum int) bool {
	// O(n/2)
	slices.Sort(arr)

	n := len(arr)
	if n < 2 {
		log.Fatal("input must have more than 1 element.")
	}
	for i, v := range arr {
		for j := i + 1; j < n; j++ {
			x.iters++
			if v+arr[j] == targetSum {
				return true
			}
		}
	}
	timeRatio := float64(x.iters) / float64(n)
	x.log("Time taken: %d for %d len arr. Ratio: %.2f", x.iters, n, timeRatio)
	return false
}
func (x NTU) hw24_04_01_LinTime(arr []int, targetSum int) bool {
	// O(n) - Two pointer technique
	slices.Sort(arr)

	n := len(arr)
	if n < 2 {
		log.Fatal("input must have more than 1 element.")
	}
	i := 0
	j := n - 1
	for i < j {
		x.iters++
		attempt := arr[i] + arr[j]
		if attempt > targetSum {
			j--
		} else if attempt < targetSum {
			i++
		} else {
			return true
		}
	}
	timeRatio := float64(x.iters) / float64(n)
	x.log("Time taken (worst case): %d for %d len arr. Ratio: %.2f", x.iters, n, timeRatio)
	return false
}

// Question 3 // file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw4_s.pdf
// Manber 5.20
type TwoPartition struct {
	p1 []int
	p2 []int
}

func (x *NTU) partitionToEqualSums(arr []int) *TwoPartition {
	// Time complexity: O(nk) + O(n) = O(n*(k+1))
	totalSum := ut.Sum(arr)
	if totalSum%2 == 1 { // Odd sum cannot be partitioned.
		return nil
	}
	target := totalSum / 2
	n := len(arr) - 1

	p := x.knapsackGemini(arr, target)
	if !p[n][target].exist {
		return nil
	}
	parts := TwoPartition{}
	for n >= 0 {
		x.iters++
		if p[n+1][target].belong {
			parts.p1 = append(parts.p1, arr[n])
			target -= arr[n]
		} else {
			parts.p2 = append(parts.p2, arr[n])
		}
		n--
	}
	return &parts
}

// non-recursive Towers of Hanoi
func (x NTU) hw24_04_05_fromPseudo(n, source, target, auxiliary int) {
	x.iters++
	type hanoi struct {
		n int
		s int
		t int
		a int
	}
	// Stack: pop and push elements off and on the top.
	stack := []hanoi{{n, source, target, auxiliary}}
	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		if p.n == 1 {
			x.log("%d: %d -> %d\n", p.n, p.s, p.t)
		} else {
			stack = append([]hanoi{{p.n - 1, p.a, p.t, p.s}}, stack...)
			stack = append([]hanoi{{1, p.s, p.t, p.a}}, stack...)
			stack = append([]hanoi{{p.n - 1, p.s, p.a, p.t}}, stack...)
		}
	}
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

// Cp5/3 One-to-One Mapping
// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch5_notes.pdf
func (x *NTU) cp5MappingBijective(f map[int]int) []int {
	// f is a mapping of a given set to itself
	set := []int{}
	c := make(map[int]int)
	for i, v := range f {
		set = append(set, i)
		if _, ok := c[i]; !ok {
			c[i] = 0
		}
		c[v]++
	}
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

// Cp5/5 Celebrity
func (x *NTU) celebrity()

// Cp5/6 The Skyline Problem
func (x *NTU) mergeSkylines(A, m, B, n int)

// hw24_04_05 recursive NTU pseudocode version
func (x *NTU) towersOfHanoiFromPseudo(n, source, target, auxiliary int) {
	x.iters++
	if n == 1 {
		x.log("%d: %d -> %d\n", n, source, target)
		return
	}
	x.towersOfHanoiFromPseudo(n-1, source, auxiliary, target)
	x.towersOfHanoiFromPseudo(1, source, target, auxiliary)
	x.towersOfHanoiFromPseudo(n-1, auxiliary, target, source)
}

// hw24_04_05 recursive internet version
func (x *NTU) towersOfHanoi(n, start, end int) {
	x.iters++
	if n == 1 {
		x.log("%d: %d -> %d\n", n, start, end)
	} else {
		other := 6 - start - end
		x.towersOfHanoi(n-1, start, other)
		x.log("%d: %d -> %d\n", n, start, end)
		x.towersOfHanoi(n-1, other, end)
	}
}
func (x NTU) hw24_04_04_from_pseudocode(arr []int, target int) []int {
	if !x.subsetSumBlackBox(arr, target) {
		return nil
	}

	n := len(arr)
	for i := 0; i < n; i++ {
		temp := arr[0]
		arr = arr[1:]
		if !x.subsetSumBlackBox(arr, target) {
			arr = append(arr, temp) // Note that we're pushing 'temp' to the END of the array -> we don't need the index 'i' for anything.
		}
		x.log("%d %d %v\n", i, temp, arr)
	}
	return arr
}

func (x *NTU) subsetSumBlackBox(arr []int, target int) bool {
	verbose := x.verbose
	x.verbose = false
	defer func(v bool) { x.verbose = v }(verbose)
	p := x.knapsackExact(arr, target)
	for _, v := range p {
		if v[target].exist {
			return true
		}
	}
	return false
}

// Algorithm SCE, file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch2_notes.pdf
func smallestCommonElement(s1, s2 []int) int {
	if len(s1) == 0 || len(s2) == 0 {
		return -1
	}
	v1 := s1[0]
	v2 := s2[0]
	if v1 == v2 {
		return v1
	}
	if v1 < v2 {
		return smallestCommonElement(s1[1:], s2)
	}
	if v1 > v2 {
		return smallestCommonElement(s1, s2[1:])
	}
	return -1
}

type NTU struct {
	verbose bool
	iters   int64
	t       *testing.T
}

func (x NTU) log(format string, a ...interface{}) {
	if x.verbose {
		fmt.Printf(format, a...)
	}
}

// Adaptation of NTU pseudocode.
// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw4_s.pdf
// s set of sizes, k size capacity to match exactly.
type ksP struct {
	exist  bool
	belong bool
	set    []int
}

func (x NTU) myKnapsack(S []int, K int) bool {
	// TODO: analyze time complexity
	sets := powerSetIterative(S)
	for _, v := range sets {
		if ut.Sum(v) == K {
			return true
		}
	}
	return false
}

// This is the best knapsack version. Adapted from Manber.
func (x NTU) knapsackExact(S []int, targetWeight int) map[int]map[int]ksP {

	p := make(map[int]map[int]ksP) // First key the row index. Second column is the weight (ranges from 0 to K).
	if len(S) == 0 {
		return p
		// if targetWeight == 0 {
		// 	p[0] = make(map[int]ksP)
		// 	p[0][0] = ksP{true, false, []int{}}
		// }
	}
	for i := range S {
		p[i] = make(map[int]ksP)
		p[i][0] = ksP{true, false, []int{}} // weight 0 => solution is the empty set.
	}

	p[0][S[0]] = ksP{true, true, []int{S[0]}} // This is needed for a correct results e.g. {1,2,3}, k = 6.
	for i := 1; i < len(S); i++ {
		for k := 0; k <= targetWeight; k++ {
			itemWeight := S[i]
			weightMinusItem := k - itemWeight
			if p[i-1][k].exist {
				x.log("if: i: %d, k: %d, itemWeight: %d\n", i, k, itemWeight)
				p[i][k] = ksP{true, false, p[i-1][k].set} // Solution exists for k. But the solution does not include i.
			} else if weightMinusItem >= 0 {
				if p[i-1][weightMinusItem].exist {
					x.log("else if if: i: %d, k: %d, itemWeight: %d, weightMinusItem: %d\n", i, k, itemWeight, weightMinusItem)
					tempSet := append(p[i-1][weightMinusItem].set, itemWeight)
					p[i][k] = ksP{true, true, tempSet}
				}
			}
		}
	}

	return p
}

// Supplementary Ex. 41, from file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/Me/Books/Rosen%20-%20Discrete%20Mathematics%20and%20Its%20Applications%207th.ed.%20(2012).pdf
// Returns the largest possible total weight of a subset of items that do not exceed capacity.
func rosenKnapsack(capacity float64, items []float64) []int {
	return []int{}
}

// By Gemini
func powerSetIterative(set []int) [][]int {
	n := len(set)
	powerSetSize := 1 << n                     // Equivalent to 2^n
	powerSet := make([][]int, 0, powerSetSize) // Pre-allocate capacity for efficiency
	for i := 0; i < powerSetSize; i++ {
		subset := make([]int, 0)
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				subset = append(subset, set[j])
			}
		}
		powerSet = append(powerSet, subset)
	}
	return powerSet
}

type Cell struct {
	exist  bool
	belong bool
}

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/Me/Books/CS/Manber%20-%20Introduction%20to%20Algorithms_%20A%20Creative%20Approach%201989_v2.pdf
func (x *NTU) knapsackGemini(S []int, K int) [][]Cell {
	// Runs in pseudopolynomial time O(nK)
	n := len(S)
	P := make([][]Cell, n+1)
	for r := 0; r <= n; r++ {
		P[r] = make([]Cell, K+1)
	}
	P[0][0].exist = true
	for k := 1; k <= K; k++ {
		P[0][k].exist = false
	}
	for i := 1; i <= n; i++ {
		itemWeight := S[i-1]
		for k := 0; k <= K; k++ {
			x.iters++
			P[i][k].exist = false
			if P[i-1][k].exist {
				P[i][k].exist = true
				P[i][k].belong = false
			} else {
				if k-itemWeight >= 0 {
					if P[i-1][k-itemWeight].exist {
						P[i][k].exist = true
						P[i][k].belong = true
					}
				}
			}
		}
	}
	return P
}
