package ntu

import (
	"algorithms/ut"
	"fmt"
	"log"
	"slices"
	"testing"
)

// Disjoint Set Union (DSU), is another name for the Union-Find data structure.
type DSU struct {
	Parent int
	Size   int
}

type UnionFind struct {
	// See file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch4_notes.pdf
	sets []DSU
}

func (x *UnionFind) UnionFindInit(seq []int) {
	for _, v := range seq {
		if v < 0 {
			log.Fatal("UnionFindInit: negative sequence values not supported")
		}
		x.sets = append(x.sets, DSU{Size: 1, Parent: -1})
	}
}

func (x *UnionFind) Find(a int) int {
	if len(x.sets) <= a {
		log.Fatalf("UnionFind.Find: index %d not in %v", a, x.sets)
	}
	if x.sets[a].Parent != -1 {
		x.sets[a].Parent = x.Find(x.sets[a].Parent)
		return x.sets[a].Parent
	}
	return a
}
func (x *UnionFind) Union(a, b int) {
	a = x.Find(a)
	b = x.Find(b)
	if a != b {
		if x.sets[a].Size > x.sets[b].Size {
			x.sets[b].Parent = a
			x.sets[a].Size += x.sets[b].Size
		} else {
			x.sets[a].Parent = b
			x.sets[b].Size += x.sets[a].Size
		}
	}
}

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
func (x NTU) hw24_04_05_fromPseudo(n, source, destination, auxiliary int) {
	x.iters++
	type hanoi struct {
		n int
		s int
		t int
		a int
	}
	// Stack: pop and push elements off and on the top.
	stack := []hanoi{{n, source, destination, auxiliary}}
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
func (x NTU) nonRecursiveHanoiJuha(n int) {
	type hanoi struct {
		n    int
		peg1 string
		peg2 string
		peg3 string
	}
	// Stack: pop and push elements off and on the top.
	stack := []hanoi{{n: n, peg1: "A", peg2: "B", peg3: "C"}}
	for len(stack) != 0 {
		x.iters++
		pop := stack[0]
		stack = stack[1:]
		if pop.n == 1 {
			x.log("%s -> %s\n", pop.peg1, pop.peg2)
		} else {
			stack = append([]hanoi{{pop.n - 1, pop.peg3, pop.peg2, pop.peg1}}, stack...)
			stack = append([]hanoi{{1, pop.peg1, pop.peg2, pop.peg3}}, stack...)
			stack = append([]hanoi{{pop.n - 1, pop.peg1, pop.peg3, pop.peg2}}, stack...)
		}
	}
}

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
	p := x.knapsackExact01(arr, target)
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
	value  int // Only relevant for teh 'valued' version.
	set    []int
}
type ksPU struct { // KnapScakPUnlimited. AKA UnboundedKnapsack.
	exist  bool
	belong int // Records how many times the same item is included in the knapsack.
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

// This is the best knapsack version. Adapted from Manber. 0/1 Means that we either take the item once or we dont.
func (x *NTU) knapsackExact01(S []int, targetWeight int) map[int]map[int]ksP {
	p := make(map[int]map[int]ksP) // First key the row index. Second key is the weight (column) (ranges from 0 to K).
	if len(S) == 0 {
		return p
	}
	for i := range S {
		p[i] = make(map[int]ksP)
		p[i][0] = ksP{true, false, 0, []int{}} // weight 0 => solution is the empty set.
	}
	// fmt.Printf("%+v\n", p)
	p[0][S[0]] = ksP{true, true, 0, []int{S[0]}} // This is needed for a correct result e.g. {1,2,3}, k = 6.
	for i := 1; i < len(S); i++ {
		itemWeight := S[i]
		for k := 0; k <= targetWeight; k++ {
			weightMinusItem := k - itemWeight
			if p[i-1][k].exist {
				x.log("if: i: %d, k: %d, itemWeight: %d\n", i, k, itemWeight)
				p[i][k] = ksP{true, false, 0, p[i-1][k].set} // Solution exists for k. But the solution does not include i.
			} else if weightMinusItem >= 0 {
				if p[i-1][weightMinusItem].exist {
					x.log("else if if: i: %d, k: %d, itemWeight: %d, weightMinusItem: %d\n", i, k, itemWeight, weightMinusItem)
					tempSet := append(p[i-1][weightMinusItem].set, itemWeight)
					p[i][k] = ksP{true, true, 0, tempSet}
				}
			}
		}
	}
	return p
}
func (x *NTU) knapsackExact01Manber5_8(S []int, targetWeight int) map[int]map[int]ksP {
	p := make(map[int]map[int]ksP) // First key the row index. Second key is the weight (column) (ranges from 0 to K).
	if len(S) == 0 {
		return p
	}
	for i := range S {
		p[i] = make(map[int]ksP)
		p[i][0] = ksP{true, false, 0, []int{}} // weight 0 => solution is the empty set.
	}
	p[0][S[0]] = ksP{true, true, 0, []int{S[0]}}
	for i := 1; i < len(S); i++ {
		itemWeight := S[i]
		for k := 0; k <= targetWeight; k++ {
			weightMinusItem := k - itemWeight
			if weightMinusItem >= 0 && p[i-1][weightMinusItem].exist {
				tempSet := append(p[i-1][weightMinusItem].set, itemWeight)
				p[i][k] = ksP{true, true, 0, tempSet}
			} else if p[i-1][k].exist {
				p[i][k] = ksP{true, false, 0, p[i-1][k].set}
			}
		}
	}
	return p
}

func (x *NTU) knapsackRecover(items []int, target int, solutions map[int]map[int]ksP) []int {
	solution := []int{}

	i := len(items) - 1
	_, ok := solutions[i][target]
	if !ok {
		return nil
	}
	// fmt.Printf("%+v\n", solutions)
	for target > 0 && i >= 0 {
		// fmt.Println(target, i)
		if solutions[i][target].belong {
			solution = append(solution, items[i])
			target -= items[i]
		}
		i--
	}

	return solution
}
func (x *NTU) knapsackExactUnlimited(S []int, targetWeight int) map[int]map[int]ksPU {
	p := make(map[int]map[int]ksPU) // First key the row index. Second column is the weight (ranges from 0 to K).
	if len(S) == 0 {
		return p
	}
	for i := range S {
		p[i] = make(map[int]ksPU)
		p[i][0] = ksPU{true, 0, []int{}} // weight 0 => solution is the empty set.
		p[i][S[i]] = ksPU{true, 1, []int{S[i]}}
	}

	{
		// This is for handling the first item in the list.
		weight := S[0]
		for {
			belong := p[0][weight].belong + 1
			tempSet := append(p[0][weight].set, S[0])
			weight += S[0]
			if weight > targetWeight {
				break
			}
			p[0][weight] = ksPU{true, belong, tempSet}
		}
	}
	for i := 1; i < len(S); i++ {
		for k := 0; k <= targetWeight; k++ {
			weightMinusItem := k - S[i]
			if p[i-1][k].exist {
				p[i][k] = ksPU{true, 0, p[i-1][k].set} // Solution exists for k. But the solution does not include i.
			} else if weightMinusItem >= 0 {
				if p[i][weightMinusItem].exist {
					// make a copy to avoid a classic Go gotcha.
					// tempSet := append(int{S[i]}, p[i][j].set...) // This is O(n)
					tempSlice := ut.CopyAndAppendSlice(p[i][weightMinusItem].set, S[i])
					belong := p[i][weightMinusItem].belong + 1
					p[i][k] = ksPU{true, belong, tempSlice}
				}
			}
		}
	}
	return p
}

func (x *NTU) knapsackExactValued(S []int, v []int, targetWeight int) map[int]map[int]ksP {
	p := make(map[int]map[int]ksP) // First key the row index. Second column is the weight (ranges from 0 to K).
	if len(S) == 0 {
		return p
	}
	for i := range S {
		p[i] = make(map[int]ksP)
		p[i][0] = ksP{true, false, 0, []int{}} // weight 0 => solution is the empty set.
	}

	p[0][S[0]] = ksP{true, true, v[0], []int{S[0]}} // This is needed for a correct result e.g. {1,2,3}, k = 6.
	for i := 1; i < len(S); i++ {
		itemWeight := S[i]
		for k := 0; k <= targetWeight; k++ {
			if p[i-1][k].exist {
				x.log("if: i: %d, k: %d, itemWeight: %d\n", i, k, itemWeight)
				p[i][k] = ksP{true, false, p[i-1][k].value, p[i-1][k].set}
			}
			weightMinusItem := k - itemWeight
			if weightMinusItem >= 0 {
				valueCandidate := p[i-1][weightMinusItem].value + v[i]
				// fmt.Printf("%d\t", valueCandidate)
				if valueCandidate > p[i][k].value {
					// fmt.Printf("else if if: i: %d, k: %d, itemWeight: %d, weightMinusItem: %d, valueCandidate: %d\n", i, k, itemWeight, weightMinusItem, valueCandidate)
					tempSet := append(p[i-1][weightMinusItem].set, itemWeight)
					p[i][k] = ksP{true, true, valueCandidate, tempSet}
				}
			}
		}
	}
	return p
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

// This does not give us the actual subsequence. Just the maximum length.
func (x *NTU) longestIncreasingSubsequence(seq []int) int {
	// fmt.Println(seq)
	seqLen := len(seq)
	max := 1
	lengths := []int{}
	for range seqLen {
		lengths = append(lengths, 1) // Init with 1 since that is the shortest possible LISS.
	}
	for i := 1; i < seqLen; i++ {
		for j := range i {
			// fmt.Print(seq[i], seq[j], " - ", result[j]+1, result[i])
			if seq[i] > seq[j] && lengths[j]+1 > lengths[i] {
				lengths[i] = lengths[j] + 1
			}
			// fmt.Println("\t", result)
		}
		if lengths[i] > max {
			// ss = append(ss, seq[i])
			max = lengths[i]
		}
	}
	fmt.Println(lengths)
	// max := slices.Max(result)
	return max
}

func (x *NTU) binaryToGray(bin string) string {
	if bin == "" {
		return ""
	}
	gray := string(bin[0])
	for i := 1; i < len(bin); i++ {
		if bin[i-1] != bin[i] {
			gray += "1"
		} else {
			gray += "0"
		}
	}

	return gray
}
