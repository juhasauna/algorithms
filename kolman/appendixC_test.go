package kolman

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_appendixC(t *testing.T) {
	// time.Sleep(2 * time.Second)
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"appC_C01_ex4", appC_C01_ex4Test},
		// {"combinationsTest", combinationsTest},
		// {"combinationsUpToTest", combinationsTest},
		// {"doTruthTable", doTruthTableTest},
		// {"equivalent", equivalentTest},
		// {"fibonacciIterative", fibonacciIterativeTest},
		// {"fibonacciUpTo", fibonacciUpToTest},
		// {"getlogicalExpressionTypeTest", getlogicalExpressionTypeTest},
		{"greatesCommonDivisor", greatesCommonDivisorTest},
		// {"intDifference", intDifferenceTest},
		// {"intIntersection", intIntersectionTest},
		// {"intUnion", intUnionTest},
		// {"leastCommonMult", leastCommonMultTest},
		// {"permutations", permutationsTest},
		// {"permutationsUpToTest", permutationsUpToTest},
		// {"permutationFunc_disjointCycleProduct", permutationFunc_disjointCycleProductTest},
		// {"permutationFunc_toTranspositionProductTest", permutationFunc_toTranspositionProductTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func permutationFunc_disjointCycleProductTest(t *testing.T) {
	x := permutationFunc{}
	tests := []struct {
		permutation map[int]int
		expect      [][]int
	}{
		{map[int]int{1: 3, 2: 4, 3: 6, 4: 5, 5: 2, 6: 1, 7: 8, 8: 7}, [][]int{{7, 8}, {2, 4, 5}, {1, 3, 6}}},
		{map[int]int{1: 2, 2: 4, 3: 5, 4: 7, 5: 6, 6: 3, 7: 1}, [][]int{{3, 5, 6}, {1, 2, 4, 7}}},
		{map[int]int{1: 9, 2: 7, 3: 3, 4: 6, 5: 4, 6: 2, 7: 1, 8: 5, 9: 8}, [][]int{{3}, {1, 9, 8, 5, 4, 6, 2, 7}}},
	}

	for _, tt := range tests {
		x.f = tt.permutation
		got := x.disjointCycleProduct()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL: expected: %v, got: %v", tt.expect, got)
		}
	}
}
func permutationFunc_toTranspositionProductTest(t *testing.T) {
	// Solution to ex. 4,5 for Cp. 5
	// App. C Cp.5: 5. Use the program in Exercise 4 as a subroutine in a program that determines whether a given permutation is even or odd.

	x := permutationFunc{}
	tests := []struct {
		permutation map[int]int
		expect      [][]int
		// even        bool
	}{
		// {map[int]int{1: 3, 2: 4, 3: 6, 4: 5, 5: 2, 6: 1, 7: 8, 8: 7}, [][]int{{7, 8}, {2, 5}, {2, 4}, {1, 6}, {1, 3}}},
		// {map[int]int{1: 2, 2: 4, 3: 5, 4: 7, 5: 6, 6: 3, 7: 1}, [][]int{{3, 6}, {3, 5}, {1, 7}, {1, 4}, {1, 2}}},
		{map[int]int{1: 9, 2: 7, 3: 3, 4: 6, 5: 4, 6: 2, 7: 1, 8: 5, 9: 8}, [][]int{{1, 7}, {1, 2}, {1, 6}, {1, 4}, {1, 5}, {1, 8}, {1, 9}}},
	}

	for i, tt := range tests {
		x.f = tt.permutation
		got := x.toTranspositionProduct()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL: expected: %v, got: %v", tt.expect, got)
		} else {
			parity := "odd"
			if len(got)%2 == 0 {
				parity = "even"
			}
			fmt.Printf("SUCCESS %d: parity: %s\n", i, parity)
		}
	}
}

func fibonacciIterativeTest(name string, t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{6, 8},
		{12, 144},
	}
	for i, tt := range tests {
		got := fibonacciIterative(tt.n)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %d, expected %v; got %v", name, tt.n, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func fibonacciUpToTest(name string, t *testing.T) {
	tests := []struct {
		n      int
		expect []int
	}{
		{0, []int{0}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 1}},
		{3, []int{0, 1, 1, 2}},
		{4, []int{0, 1, 1, 2, 3}},
		{5, []int{0, 1, 1, 2, 3, 5}},
		{6, []int{0, 1, 1, 2, 3, 5, 8}},
		{12, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}},
	}
	for i, tt := range tests {
		got := fibonacciUpTo(tt.n)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %d, expected %v; got %v", name, tt.n, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func combinationsUpToTest(name string, t *testing.T) {
	ut := utils{}
	tests := []struct {
		n      int
		r      int
		expect []int
	}{
		{1, 10, []int{0}},
		{1, 1, []int{1}},
		{2, 1, []int{2, 1}},
		{3, 1, []int{3, 3}},
		{4, 1, []int{4, 6, 3, 1}},
		{4, 2, []int{6, 3, 1}},
		{4, 3, []int{4, 1}},
	}
	for i, tt := range tests {
		got := ut.combinationsUpTo(tt.n, tt.r)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.n, tt.r, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func combinationsTest(name string, t *testing.T) {
	ut := utils{}
	tests := []struct {
		n      int
		r      int
		expect int
	}{
		{1, 10, 0},
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{4, 1, 4},
		{4, 2, 6},
		{4, 3, 4},
	}
	for i, tt := range tests {
		got := ut.combinations(tt.n, tt.r)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.n, tt.r, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func permutationsUpToTest(name string, t *testing.T) {
	ut := utils{}
	tests := []struct {
		n      int
		r      int
		expect []int
	}{
		{1, 10, []int{}},
		{1, 1, []int{1}},
		{4, 3, []int{4, 12, 24}},
	}
	for i, tt := range tests {
		got := ut.permutationsUpTo(tt.n, tt.r)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.n, tt.r, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func permutationsTest(name string, t *testing.T) {
	ut := utils{}
	tests := []struct {
		n      int
		r      int
		expect int
	}{
		{1, 10, 0},
		{1, 1, 1},
		{4, 3, 24},
	}
	for i, tt := range tests {
		got := ut.permutations(tt.n, tt.r)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.n, tt.r, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func intUnionTest(name string, t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{-1, -3, -3}, []int{-3, -4, -5}, []int{-1, -3, -4, -5}},
	}
	for i, tt := range tests {
		got := intUnion(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func intIntersectionTest(name string, t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{3}},
		{[]int{-1, -3, -3, -4}, []int{-3, -4, -5}, []int{-3, -4}},
	}
	for i, tt := range tests {
		got := intIntersection(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func intDifferenceTest(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{-1, -3, -3, -4}, []int{-3, -4, -5}, []int{-1}},
		{[]int{-1, -3, -3, -4}, []int{-12, -12, -5}, []int{-1, -3, -4}},
	}
	for i, tt := range tests {
		got := intDifference(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("\tfail with inputs %v %v, expected %v; got %v", tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("\tSUCCESS: %d\n", i)
		}
	}
}
func appC_C01_ex4Test(name string, t *testing.T) {
	tests := []struct {
		n      int
		expect []int
	}{
		{0, []int{1}},
		{1, []int{1, -1}},
		{2, []int{1, -1, -3 - 2}},
		{3, []int{1, -1, -5, -13}},
		{4, []int{1, -1, -5, -13, -39 + 10}},
	}
	for i, tt := range tests {
		got := appC_C01_ex4(tt.n)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with input %d, expected %v; got %v", name, tt.n, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func greatesCommonDivisorTest(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{1, 1, 1},
		{12, 30, 6},
		{150, 70, 10},
		{32, 27, 1},
		{190, 34, 2},
		{17, 95, 1},
		{273, 98, 7},
		{540, 504, 36},
	}
	for i, tt := range tests {
		got := greatestCommonDivisor(tt.a, tt.b)
		got2 := euclideanAlgorithm(tt.a, tt.b)
		if !(got == tt.expect && got2 == tt.expect) {
			t.Errorf("FAIL: inputs %d, %d, expected %d; got %d; got2 %d", tt.a, tt.b, tt.expect, got, got2)
		} else {
			fmt.Printf("\tSUCCESS: %d, inputs (%d,\t%d),\tgot %v\n", i, tt.a, tt.b, got)
		}
	}
}
func leastCommonMultTest(name string, t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{2, 4, 4},
		{2, 6, 6},
		{6, 12, 12},
		{540, 504, 7560},
	}
	for i, tt := range tests {
		got := leastCommonMult(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %d, %d, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func doTruthTableTest(name string, t *testing.T) {
	tests := []struct {
		pqr    []bool
		f      func(...bool) bool
		expect []bool
	}{
		{[]bool{false, false}, func(pqr ...bool) bool { return pqr[0] && !pqr[1] }, []bool{false, false, true, false}},
		{[]bool{false, false}, func(pqr ...bool) bool { return pqr[0] || pqr[1] }, []bool{false, true, true, true}},
		{[]bool{false, false, false}, func(pqr ...bool) bool { return pqr[0] || pqr[1] == pqr[2] }, []bool{true, false, false, true, true, true, true, true}},
	}
	for i, tt := range tests {
		got := doTruthTable(tt.f, tt.pqr)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v, expected %v; got %v", name, tt.pqr, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func equivalentTest(name string, t *testing.T) {
	tests := []struct {
		p      bool
		q      bool
		pf     func(bool) bool
		qf     func(bool) bool
		expect bool
	}{
		{false, false, func(bool) bool { return true }, func(bool) bool { return true }, true},
		{false, true, func(p bool) bool { return p }, func(q bool) bool { return q }, false},
	}
	for i, tt := range tests {
		got := equivalent(tt.p, tt.pf, tt.q, tt.qf)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %t, %t, expected %v; got %v", name, tt.p, tt.q, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func getlogicalExpressionTypeTest(name string, t *testing.T) {
	tests := []struct {
		p      bool
		q      bool
		f      func(bool, bool) bool
		expect logicalExprType
	}{
		{false, false, func(bool, bool) bool { return true }, Tautology},
		{false, false, func(bool, bool) bool { return false }, Absurdity},
		{false, false, func(p bool, q bool) bool {
			if p {
				return p
			}
			return q
		}, Contingency},
	}
	for i, tt := range tests {

		got := getlogicalExpressionType(tt.p, tt.q, tt.f)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %t, %t, expected %v; got %v", name, tt.p, tt.q, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d, got=%s\n", name, i, got.String())
		}
	}
}
