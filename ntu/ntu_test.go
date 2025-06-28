package ntu

import (
	"algorithms/sorting"
	"algorithms/ut"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_ntu(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"powerSetIterative", powerSetIterativeTest},
		// {"myKnapsack", myKnapsackTest},
		// {"knapsackGemini", knapsackGeminiTest},
		// {"hw04_01", hw04_01Test},
		// {"partitionToEqualSums", partitionToEqualSumsTest},
		// {"subsetSumBlackBox", subsetSumBlackBoxTest},
		// {"knapsackExact", knapsackExactTest},
		// {"hw24_04_04", hw24_04_04Test},
		// {"towersOfHanoiTest", towersOfHanoiTest},
		// {"cp5MappingBijective", cp5MappingBijectiveTest},
		// {"evaluatePolynomials", evaluatePolynomialsTest},
		// {"maximalInducedSubgraph", maximalInducedSubgraphTest},
		// {"newAdjMatrix", newAdjMatrixTest},
		// {"celebrityBruteForce", celebrityBruteForceTest},
		// {"maximumConsequtiveSubsequence", maximumConsequtiveSubsequenceTest},
		{"UnionFind", UnionFindTest},
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

func UnionFindTest(t *testing.T) {
	tests := []struct {
		seq   []int
		union []Tuple
	}{
		{[]int{1, 2, 3, 4, 5}, []Tuple{{1, 2}, {3, 4}, {0, 1}, {4, 1}}},
	}
	for _, tt := range tests {
		x := UnionFind{}
		x.UnionFindInit(tt.seq)
		t.Logf("%v", x.sets)
		for _, v := range tt.union {
			x.Union(v.a, v.b)
		}
		t.Logf("%v", x.sets)

	}
}
func maximumConsequtiveSubsequenceTest(t *testing.T) {
	tests := []struct {
		arr  []float64
		want float64
	}{
		{[]float64{2, -3, 1.5, -1, 3, -2, -3, 3}, 3.5},
	}
	for i, tt := range tests {
		x := NTU{verbose: true}
		got := x.maximumConsequtiveSubsequence(tt.arr)
		if got != tt.want {
			t.Errorf("FAIL got: %.2f, want: %.2f", got, tt.want)
		} else {
			t.Logf("i: %d SUCCESS got: %v", i, got)
		}
	}
}
func celebrityBruteForceTest(t *testing.T) {
	tests := []struct {
		s         []string
		celebrity int
	}{
		{[]string{
			"101",
			"011",
			"001",
		}, 2},
		{[]string{
			"1001",
			"0101",
			"0011",
			"0001",
		}, 3},
		{[]string{
			"1110",
			"0100",
			"0111",
			"1101",
		}, 1},
		{[]string{
			"1110",
			"0110",
			"0111",
			"1001",
		}, -1},
	}
	for i, tt := range tests {
		x := NTU{verbose: false}
		m := x.newAdjMatrix(tt.s)
		got := x.celebrityBruteForce(m)
		t.Logf("BF iters: %d", x.iters)
		x.iters = 0
		gotCeleb := x.celebrityFromPseudo(m)
		t.Logf("Pseudo iters: %d", x.iters)
		if gotCeleb != tt.celebrity {
			t.Errorf("i: %d FAIL got: %v, want: %d", i, gotCeleb, tt.celebrity)
		} else {
			t.Logf("i: %d SUCCESS got: %d", i, gotCeleb)
		}
		if !got[tt.celebrity] && tt.celebrity != -1 {
			t.Errorf("FAIL got: %v, %d should be true", got, tt.celebrity)
		} else {
			t.Logf("i: %d SUCCESS got: %v", i, got)

		}
	}
}
func newAdjMatrixTest(t *testing.T) {
	tests := []struct {
		s []string
	}{
		{[]string{
			"101",
			"111",
			"000",
		}},
	}
	for _, tt := range tests {
		x := NTU{}
		got := x.newAdjMatrix(tt.s)
		t.Logf("%v", got)
	}
}

func maximalInducedSubgraphTest(t *testing.T) {
	tests := []struct {
		g       map[string][]string
		k       int
		wantLen int
	}{
		{map[string][]string{"a": {"b", "c", "d"}, "b": {"a", "c", "d"}, "c": {"a", "b"}, "d": {"a"}}, 2, 3},
		{map[string][]string{"1": {"2", "3", "7", "8"}, "2": {"1", "3"}, "3": {"1", "2", "4", "5", "6", "7", "8"}, "4": {"3", "5", "6", "7", "8"}, "5": {"3", "4", "6", "7"}, "6": {"3", "4", "5", "7", "8"}, "7": {"1", "4", "5", "6", "8"}, "8": {"1", "3", "4", "6", "7"}}, 4, 6},
	}
	for _, tt := range tests {
		x := NTU{verbose: true}
		// got := x.maximalInducedSubgraphRecursive(tt.g, tt.k)
		got := x.maximalInducedSubgraphIterative(tt.g, tt.k)
		if len(got) != tt.wantLen {
			t.Errorf("FAIL gotlen: %d, wantlen: %d", len(got), tt.wantLen)
		} else {
			for key, v := range got {
				if len(v) < tt.k {
					t.Errorf("FAIL got [%s]%v, but len(v) must be less than k:%d", key, v, tt.k)
				}
			}
			t.Logf("SUCCESS: %v", got)
		}
	}
}
func evaluatePolynomialsTest(t *testing.T) {
	tests := []struct {
		a_i  []float64
		y    float64
		want float64
	}{
		// {[]float64{1}, 1, 1},
		// {[]float64{1, 1}, 1, 2},
		// {[]float64{0, 1, 2}, 1, 3},
		{[]float64{1, 2, 3}, 2, 11},
	}
	for _, tt := range tests {
		x := NTU{verbose: true}
		got := x.hornersRule(tt.a_i, tt.y)
		// got := x.evaluatePolySlow(tt.a_i, tt.y)
		if got != tt.want {
			t.Errorf("FAIL got %f != %f", got, tt.want)
		} else {
			t.Logf("SUCCESS: %f, iters: %d", got, x.iters)
		}
	}
}

func cp5MappingBijectiveTest(t *testing.T) {
	tests := []struct {
		mapping map[int]int
		want    []int
	}{
		// IMPORTANT: The function needs to be defined for all elements in the set! If the set includes 5, then f(5) must be defined.
		{map[int]int{1: 3, 2: 1, 3: 1, 4: 4, 5: 2}, []int{1, 3, 4}},
		// {map[int]int{1: 2, 3: 4}, []int{}}, // Invalid function
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		got := x.cp5MappingBijective(tt.mapping)
		if !reflect.DeepEqual(got, tt.want) && !(len(got) == 0 && len(tt.want) == 0) {
			t.Errorf("FAIL got: %v != %v", got, tt.want)
		} else {
			t.Logf("SUCCESS: %v", got)

		}
	}
}
func towersOfHanoiTest(t *testing.T) {
	tests := []struct {
		n         int
		source    int
		target    int
		auxiliary int
	}{
		{3, 1, 2, 3},
		// {8, 1, 2,3},
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		// x.towersOfHanoiFromPseudo(tt.n, tt.source, tt.target, tt.auxiliary)
		x.hw24_04_05_fromPseudo(tt.n, tt.source, tt.target, tt.auxiliary)
		t.Logf("towers of hanoi moves for n=%d: %d", tt.n, x.iters)
	}
}
func towersOfHanoiTestStartEnd(t *testing.T) {
	tests := []struct {
		n     int
		start int
		end   int
	}{
		// {3, 1, 2},
		{8, 1, 2},
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		x.towersOfHanoi(tt.n, tt.start, tt.end)
		t.Logf("towers of hanoi moves for n=%d: %d", tt.n, x.iters)
	}
}
func hw24_04_04Test(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   bool
	}{
		// {"1", []int{1}, 1, true},
		// {"1", []int{1}, 2, false},
		// {"1", []int{1, 2}, 2, true},
		// {"1", []int{1, 2}, 3, true},
		// {"1", []int{1, 2}, 4, false},
		// {"", []int{2, 3, 5, 6}, 6, true},
		{"", []int{2, 3, 5, 6}, 10, true},
		{"", []int{2, 3, 5, 6}, 8, true},
		{"", []int{2, 3, 5, 6}, 9, true},
		{"", []int{2, 3, 5, 6}, 12, false},
		{"", []int{2, 3, 5, 6}, 16, true},
		{"", []int{12, 23, 34, 45, 56}, 35, true},
		{"", []int{12, 23, 34, 45, 56}, 102, true},
		{"", []int{12, 23, 34, 45, 56}, 170, true},
		{"", []int{12, 23, 34, 45, 56}, 169, false},
		{"", []int{12, 23, 34, 45, 56}, 200, false},
		// {"This is too much29691+41=29732", sorting.TestDdata4813, 29732, true},
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		got := x.hw24_04_04_from_pseudocode(tt.arr, tt.target)
		if (tt.target == ut.Sum(got)) != tt.want {
			t.Errorf("FAIL: want: %t, %v, target: %d", tt.want, got, tt.target)
		} else {
			t.Logf("SUCCESS: want: %t, %v, target: %d", tt.want, got, tt.target)
		}
	}
}
func subsetSumBlackBoxTest(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   bool
	}{
		// {"1", []int{1}, 1, true},
		// {"1", []int{1}, 2, false},
		// {"1", []int{1, 2}, 2, true},
		// {"1", []int{1, 2}, 3, true},
		// {"1", []int{1, 2}, 4, false},
		// {"", []int{2, 3, 5, 6}, 6, true},
		// {"", []int{2, 3, 5, 6}, 10, true},
		// {"", []int{2, 3, 5, 6}, 8, true},
		// {"", []int{2, 3, 5, 6}, 9, true},
		// {"", []int{2, 3, 5, 6}, 12, false},
		// {"", []int{2, 3, 5, 6}, 16, true},
		// {"29691+41=29732", sorting.TestDdata4813, 29732, true},
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		// got := x.hw04_01_polyTime(tt.arr, tt.targetSum)
		if got := x.subsetSumBlackBox(tt.arr, tt.target); got != tt.want {
			t.Errorf("got: %t, want: %t", got, tt.want)
		}

	}
}
func partitionToEqualSumsTest(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		// {"", []int{1, 2, 3}, true},
		// {"", []int{1, 2}, false},
		// {"", []int{33, 22, 55, 44}, true},
		// {"", []int{123, 2, 17, 138}, true},
		{"", []int{300, 2, 200, 3, 100, 1, 400}, true},
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		// got := x.hw04_01_polyTime(tt.arr, tt.targetSum)
		got := x.partitionToEqualSums(tt.arr)
		if (got != nil) != tt.want {
			t.Errorf("got: %v, want: %t", got, tt.want)
		} else if got != nil {
			sum1, sum2 := ut.Sum(got.p1), ut.Sum(got.p2)
			if sum1 != sum2 {
				t.Errorf("Partition sums don't match: %d != %d", sum1, sum2)
			}
			t.Logf("got: %+v, iters: %d, n:%d", got, x.iters, len(tt.arr))
		}
	}
}

func hw04_01Test(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		targetSum int
		want      bool
	}{
		// {"", []int{1, 2, 3}, 5, true},
		// {"", []int{1, 2, 3}, 4, true},
		// {"", []int{1, 2, 3}, 3, true},
		// {"", []int{1, 2, 3}, 2, false},
		// {"", []int{1, 2, 3}, 6, false},
		// {"29691+41=29732", sorting.TestDdata4813, 29732, true},
		{"", sorting.TestDdata4813, 200000, false},
	}

	for _, tt := range tests {
		x := NTU{verbose: true}
		// got := x.hw24_04_01_polyTime(tt.arr, tt.targetSum)
		got := x.hw24_04_01_LinTime(tt.arr, tt.targetSum)
		if got != tt.want {
			t.Errorf("got: %t, want: %t", got, tt.want)
		}
		// t.Logf("got: %+v", got)
	}
}
func knapsackGeminiTest(t *testing.T) {
	tests := []struct {
		name string
		S    []int
		K    int
	}{
		// {"1", []int{1}, 1},
		// {"1", []int{1, 2}, 3},
		{"manber", []int{2, 3, 5, 6}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NTU{verbose: true}
			got := x.knapsackGemini(tt.S, tt.K)
			for i, v := range got {
				for j, w := range v {
					// if j == len(v)-1 {
					if w.belong {
						t.Logf("[%d][%d]\t%t\t%t", i, j, w.exist, w.belong)
					}
					// }
				}

			}
			// t.Logf("got: %+v", got)
		})
	}
}
func myKnapsackTest(t *testing.T) {
	tests := []struct {
		name string
		S    []int
		K    int
		want bool
	}{
		{"1", []int{1}, 1, true},
		{"1", []int{1, 2}, 1, true},
		{"1", []int{1, 2}, 2, true},
		{"1", []int{1, 2}, 3, true},
		{"1", []int{1, 2, 3}, 5, true},
		{"1", []int{2, 4, 6}, 7, false},
		{"2", []int{1, 1, 2}, 2, true},
		{"3", []int{2, 4}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NTU{verbose: true}
			got := x.myKnapsack(tt.S, tt.K)
			if got != tt.want {
				t.Errorf("got: %t, want: %t", got, tt.want)
			}
			t.Logf("got: %+v", got)
		})
	}

}

func (x NTU) checkKnapsackResult(t *testing.T, m map[int]map[int]ksP, K int) bool {
	solutionExists := false
	for _, v := range m {
		if !v[0].exist {
			t.Error("knapsack empty set should exists for all rows")
			return false
		}
		if v[0].belong {
			t.Error("knapsack empty set should not have any belongs")
			return false
		}
		if v[K].exist {
			solutionExists = true
		}
		for weight, w := range v {
			if weight != ut.Sum(w.set) {
				t.Errorf("The elements (%v) do not sum up to the weight (%d)", w.set, weight)
				return false
			}
		}

	}
	for _, v := range m {
		delete(v, 0) // Delete empty sets because they are uninteresting.
		for k, w := range v {
			if k == w.set[0] {
				delete(v, k) // Delete sets with single element because they are uninteresting.
			}
			if !w.belong {
				delete(v, k)
			}
		}
	}

	return solutionExists
}

func knapsackExactTest(t *testing.T) {
	tests := []struct {
		name           string
		S              []int
		K              int
		solutionExists bool
	}{
		// {"True or false?", []int{}, 0, false},
		{"1", []int{2}, 2, true},
		{"1", []int{2}, 1, false},
		{"1", []int{2}, 0, true},
		// {"1", []int{2}, 2, true},
		// {"1", []int{1, 2}, 3, true},
		// {"1", []int{1, 2, 3}, 5, true},
		// {"1", []int{1, 2, 3}, 6, true},
		// {"Manber exp. (p.127)", []int{2, 3, 5, 6}, 16, true},
		// {"1", []int{2, 4, 6}, 7, false},
		// {"2", []int{1, 1, 2}, 2, true},
		// {"3", []int{2, 4}, 3, false},
		// {"", []int{4, 8, 12, 7}, 13, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NTU{verbose: false}
			got := x.knapsackExact(tt.S, tt.K)
			solutionExists := x.checkKnapsackResult(t, got, tt.K)
			if solutionExists != tt.solutionExists {
				t.Errorf("solutionExists: %t, want: %t", solutionExists, tt.solutionExists)
			}
			t.Log("\ti / v / weight & set")
			for i, v := range got {
				trimMap := []string{}
				for j, w := range v {
					trimMap = append(trimMap, fmt.Sprintf("%d,%v", j, w.set))
				}
				t.Logf("\t%d / %d /\t%s", i, tt.S[i], strings.Join(trimMap, "   "))
			}
		})
	}

}

func powerSetIterativeTest(t *testing.T) {
	tests := []struct {
		name string
		set  []int
		want [][]int
	}{
		{"empty", []int{}, [][]int{{}}},
		{"simple", []int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		{"simple", []int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerSetIterative(tt.set)
			if len(got) != len(tt.want) {
				t.Errorf("powerSetIterative() = %v, want %v", got, tt.want)
			}
			t.Logf("%v", got)
		})
	}
}

func smallestCommonElementTest(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want int
	}{
		{"empty", []int{}, []int{}, -1},
		{"simple", []int{1, 2, 3}, []int{2, 3, 4}, 2},
		{"simple", []int{10, 15, 18}, []int{1, 2, 3, 4, 5, 10, 15, 18}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestCommonElement(tt.s1, tt.s2); got != tt.want {
				t.Errorf("smallestCommonElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
