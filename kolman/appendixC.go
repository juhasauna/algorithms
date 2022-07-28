// p. 460
package kolman

import (
	"fmt"
	"log"
	"sort"
)

var ut utils

func intUnion(a, b []int) []int { // Ex 1
	result := a
	for _, v := range b {
		if !ut.containsInt(v, result) {
			result = append(result, v)
		}
	}
	return ut.distinctInt(result)
}

func intIntersection(a, b []int) []int { // Ex 2
	result := []int{}
	for _, v := range a {
		if ut.containsInt(v, b) {
			result = append(result, v)
		}
	}
	return ut.distinctInt(result)
}

func intDifference(a, b []int) []int { // Ex 3.
	// A - B (or using a different notation) A \ B
	result := a
	for i, v := range result {
		if ut.containsInt(v, b) {
			result = ut.rmElementInt(i, result)
			result = intDifference(result, b)
			break
		}
	}
	return ut.distinctInt(result)
}

func appC_C01_ex4_g(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return -1
	}
	v1 := 3 * appC_C01_ex4_g(n-1)
	v2 := 2 * appC_C01_ex4_g(n-2)
	return v1 - v2
}
func appC_C01_ex4(elements int) []int {
	result := []int{}
	if elements < 0 {
		return result
	}
	for i := 0; i <= elements; i++ {
		result = append(result, appC_C01_ex4_g(i))
	}
	return result
}


func (x utils) not(p bool) bool {
	return !p
}

func (x utils) and(p ...bool) bool {
	for _, v := range p {
		if !v {
			return false
		}
	}
	return true
}
func (x utils) or(p ...bool) bool {
	for _, v := range p {
		if v {
			return true
		}
	}
	return false
}

func doTruthTable(f func(...bool) bool, pqr []bool) []bool {
	switch len(pqr) {
	case 0, 1:
		return []bool{}
	case 2:
		return []bool{
			f(pqr[0], pqr[1]),
			f(pqr[0], !pqr[1]),
			f(!pqr[0], pqr[1]),
			f(!pqr[0], !pqr[1]),
		}
	case 3:
		return []bool{
			f(pqr[0], pqr[1], pqr[2]),
			f(pqr[0], pqr[1], !pqr[2]),
			f(pqr[0], !pqr[1], pqr[2]),
			f(pqr[0], !pqr[1], !pqr[2]),
			f(!pqr[0], pqr[1], pqr[2]),
			f(!pqr[0], pqr[1], !pqr[2]),
			f(!pqr[0], !pqr[1], pqr[2]),
			f(!pqr[0], !pqr[1], !pqr[2]),
		}
	}
	return []bool{}
}

func equivalent(p bool, pf func(bool) bool, q bool, qf func(bool) bool) bool {
	result := pf(p) == qf(q)
	return result
}

type logicalExprType int

const (
	Tautology logicalExprType = iota
	Absurdity
	Contingency
)

func (x logicalExprType) String() string {
	return [...]string{
		"Tautology",
		"Absurdity",
		"Contingency",
	}[x]
}

func getlogicalExpressionType(p, q bool, f func(bool, bool) bool) logicalExprType {
	isTautology := f(p, q) && f(p, !q) && f(!p, q) && f(!p, !q)
	if isTautology {
		return Tautology
	}
	isAbsurdity := !f(p, q) && !f(p, !q) && !f(!p, q) && !f(!p, !q)
	if isAbsurdity {
		return Absurdity
	}
	return Contingency
}

func (x utils) fact(n int) int {
	if n < 2 {
		return 1
	}
	return n * x.fact(n-1)
}
func (x utils) combinations(n, r int) int {
	if n < r {
		return 0
	}
	return x.fact(n) / (x.fact(r) * x.fact(n-r))
}
func (x utils) permutations(n, r int) int {
	if n < r {
		return 0
	}
	return x.fact(n) / x.fact(n-r)
}
func (x utils) combinationsUpTo(n, r int) []int {
	result := []int{}
	if r > n {
		return result
	}
	for i := 1; i <= r && i <= n; i++ {
		result = append(result, x.combinations(n, i))
	}
	return result
}
func (x utils) permutationsUpTo(n, r int) []int {
	result := []int{}
	if r > n {
		return result
	}
	for i := 1; i <= r && i <= n; i++ {
		result = append(result, x.permutations(n, i))
	}
	return result
}

func fibonacciIterative(k int) int {
	results := []int{0}
	for i := 1; i <= k; i++ {
		fmt.Printf("%v\n", results)
		if i < 3 {
			results = append(results, 1)
		} else {
			v := results[i-2] + results[i-3]
			results = append(results, v)
		}
	}
	return results[k]
}
func fibonacci(k int) int {
	if k < 1 {
		return 0
	}
	if k == 1 {
		return 1
	}
	return fibonacci(k-1) + fibonacci(k-2)
}

func fibonacciUpTo(k int) []int {
	result := []int{}
	for i := 0; i <= k; i++ {
		result = append(result, fibonacci(i))
	}
	return result
}

type point struct {
	x int
	y int
}

func crossProduct(m, n int) []point {
	result := []point{}
	if m < 0 || n < n {
		fmt.Println("invalid input, expected positive integers")
		return result
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			result = append(result, point{i, j})
		}
	}
	return result
}

type permutationFunc struct {
	f    map[int]int
	keys []int
}

func (x *permutationFunc) disjointCycleProduct() (result [][]int) {
	// App. C Cp.5: 3. Write a program that writes a given permutation as a product of disjoint cycles.
	keys := []int{}
	for key := range x.f {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	x.keys = keys
	for len(x.keys) != 0 {
		key := x.keys[0]
		cycle := x.disjointCycle([]int{}, key)
		x.removeKeys(cycle)
		result = append([][]int{cycle}, result...)
	}
	return result
}

func (x *permutationFunc) disjointCycle(cycle_ []int, v int) (cycle []int) {
	cycle = append(cycle_, v)
	cycleValue, ok := x.f[v]
	if !ok {
		log.Fatalf("invalid permutationFunc, key: %d", v)
	}
	if cycleValue == cycle[0] {
		return cycle
	}
	return x.disjointCycle(cycle, cycleValue)
}

func (x *permutationFunc) removeKeys(cycle []int) {
	for _, v := range cycle {
		for i := len(x.keys) - 1; i >= 0; i-- {
			if v == x.keys[i] {
				x.keys = append(x.keys[:i], x.keys[i+1:]...)
			}
		}
	}
}

func (x *permutationFunc) toTranspositionProduct() (result [][]int) {
	cycles := x.disjointCycleProduct()
	// App. C Cp.5: 4. Write a program that writes a given permutation as a product of transpositions.
	// A cycle of length 2 is called a transposition.
	// Every cycle can be written as a product of transpositions.
	// Every permutation of a finite set with at least two elements can be written as a product of transpositions. The transpositions need not be disjoint.
	for _, cycle := range cycles {
		for j := len(cycle) - 1; j > 0; j-- {
			result = append(result, []int{cycle[0], cycle[j]})
		}
	}
	return result
}
