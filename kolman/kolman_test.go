package kolman

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		f    func(string, *testing.T)
		do   bool
	}{
		{"intUnion", intUnionTest, false},
		{"intIntersection", intIntersectionTest, false},
		{"intDifference", intDifferenceTest, false},
		{"appC_C01_ex4", appC_C01_ex4Test, false},
		{"greatesCommonDivisor", greatesCommonDivisorTest, false},
		{"leastCommonMult", leastCommonMultTest, false},
		{"doTruthTable", doTruthTableTest, false},
		{"equivalent", equivalentTest, false},
		{"getlogicalExpressionTypeTest", getlogicalExpressionTypeTest, false},
		{"permutations", permutationsTest, false},
		{"permutationsUpToTest", permutationsUpToTest, false},
		{"combinationsTest", combinationsTest, false},
		{"combinationsUpToTest", combinationsTest, false},
		{"fibonacciUpTo", fibonacciUpToTest, false},
		{"fibonacciIterative", fibonacciIterativeTest, false},
		{"isReflexive", isReflexiveTest, true},
	}
	for _, tt := range tests {
		if tt.do {
			tt.f(tt.name, t)
		}
	}
}

func isReflexiveTest(name string, t *testing.T) {
	tests := []struct {
		relation []relation
		expect   bool
	}{
		{[]relation{{"a", "b"}}, false},
		{[]relation{{"a", "b"}, {"b", "a"}}, true},
		{[]relation{{"a", "b"}, {"b", "a"}, {"c", "a"}}, false},
		{[]relation{{"a", "b"}, {"b", "a"}, {"c", "a"}, {"a", "c"}, {"a", "a"}}, true},
	}
	for i, tt := range tests {
		got := isReflexive(tt.relation)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v, expected %t; got %v", name, tt.relation, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
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
func intDifferenceTest(name string, t *testing.T) {
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
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
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
func greatesCommonDivisorTest(name string, t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{1, 1, 1},
		{190, 34, 2},
		{540, 504, 36},
	}
	for i, tt := range tests {
		got := greatestCommonDivisor(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %d, %d, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
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
