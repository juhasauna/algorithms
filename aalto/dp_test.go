package aalto

import (
	"testing"
)

func Test_dp(t *testing.T) {

	// coinTest(t)
	// checkerboardTest(t)
	// tournamentProbabilityTest(t)
	// longestCommonSubsequenceTest(t)
	// longestIncreasingSubsequenceTest(t)
	// matrixChainMultiplicationTest(t)
	// rodCuttingTest(t)
	weightedIntervalSchedulingTest(t)
	// activitySelectionTest(t)
	// longestPalindromicSubsequenceTest(t)

}

func weightedIntervalSchedulingTest(t *testing.T) {
	tests := jobSelectionTests(true)
	for _, tt := range tests {
		got := weightedIntervalScheduling(tt.jobs)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			t.Log(got, counter)
		}

	}
}
func activitySelectionTest(t *testing.T) {
	tests := jobSelectionTests(false)
	for _, tt := range tests {
		// got := activitySelectionRecursive(tt.jobs)
		// got := activitySelectionIterative(tt.jobs)
		jobs, got := activitySelectionDP(tt.jobs)
		t.Log(jobs)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got, counter)
		}

	}
}
func rodCuttingTest(t *testing.T) {
	tests := []struct {
		name string
		x    []int
		want int
	}{
		// {"", []int{1, 3}, 3},
		// {"", []int{1, 3, 10}, 10},
		// {"", []int{1, 6, 7, 8}, 12},
		// {"", []int{2, 3, 4, 5}, 8},
		// {"", []int{1, 5, 8, 9}, 10},
		// {"", []int{1, 7, 8, 9}, 14},
		// {"", []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 30},
		// {"", []int{1, 5, 8, 15, 16, 17, 17, 20, 24, 30}, 35},
		{"naive: 70s, iters: 34 359 738 368", []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 20, 23, 29, 30, 36, 42, 44, 49, 55, 56, 61, 66, 69, 70, 75, 78, 80, 82, 87, 91, 91, 91, 92, 94, 95, 99}, 35},
	}

	for _, tt := range tests {
		got := rodCutting(tt.x)
		t.Log("rodCutting", counter)
		counter = 0
		got = rodCuttingRecursiveTabulation(tt.x)
		t.Log("tabulation", counter)
		counter = 0
		got = rodCuttingMemoized(tt.x)
		t.Log("memo", counter)
		counter = 0
		got = rodCuttingNaive(tt.x)
		t.Log("naive", counter)
		counter = 0
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got, counter)
		}
	}
}
func matrixChainMultiplicationTest(t *testing.T) {
	tests := []struct {
		name string
		x    []int
		want int
	}{
		{"", []int{2, 2, 2}, 8},
		{"", []int{10, 30, 5, 60}, 4500},
		{"", []int{30, 35, 15, 5, 10, 20, 25}, 15125},
		// {"", []int{1, 2}, 2},
	}
	for _, tt := range tests {
		got := matrixChainMultiplication(tt.x)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}
func longestIncreasingSubsequenceTest(t *testing.T) {
	tests := []struct {
		name string
		x    []int
		want int
	}{
		{"", []int{1}, 1},
		{"", []int{1, 2}, 2},
		{"", []int{1, 2, 3}, 3},
		{"", []int{1, 2, 3, 4}, 4},
		{"", []int{1, 2, 2}, 2},
		{"", []int{1, 2, 1, 3}, 3},
		{"", []int{1, 3, 11, 5, 12, 14, 7, 9, 15}, 6},
	}
	for _, tt := range tests {
		got := longestIncreasingSubsequence(tt.x)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}
func longestCommonSubsequenceTest(t *testing.T) {
	tests := []struct {
		name string
		x    string
		y    string
		want int
	}{
		{"", "ABCD", "ACBD", 3},
		{"alg2021_final", "abcabcabc", "aaabbbccc", 5},
	}
	for _, tt := range tests {
		got := longestCommonSubsequence(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}
func tournamentProbabilityTest(t *testing.T) {
	tests := []struct {
		name string
		n    int
		i    int
		j    int
		want float32
	}{
		{"", 1, 0, 0, .5},
		{"", 2, 0, 0, .5},
		{"", 10, 5, 5, .5},
		{"", 10, 8, 9, .25},
		{"", 10, 9, 8, .75},
		{"", 5, 1, 4, 1.0 / 16.0},
		{"", 5, 0, 4, 1.0 / 32.0},
	}
	for _, tt := range tests {
		got := tournamentProbability(tt.n, tt.i, tt.j)
		if got != tt.want {
			t.Errorf("got/want (%.4f/%.4f)", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}
func longestPalindromicSubsequenceTest(t *testing.T) {
	tests := []struct {
		name string
		x    string
		want int
	}{
		{"", "a", 1},
		{"", "aa", 2},
		{"", "aaa", 3},
		{"", "aaaa", 4},
		{"aalto_exam", "ACGTGTCAAAATCG", 8},
	}
	for _, tt := range tests {
		got := longestPalindromicSubsequence(tt.x)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}
func checkerboardTest(t *testing.T) {
	var checkerboards = makeCheckerboards()
	tests := []struct {
		name string
		want int
	}{
		// {"1", 31},
		{"2", 22},
	}
	for _, tt := range tests {
		got := checkerboard(checkerboards[tt.name])
		if got != tt.want {
			t.Errorf("got != tt.want %d %d", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}

func makeCheckerboards() map[string][][]int {
	m := make(map[string][][]int)
	m["1"] = [][]int{
		{3, 7, 1, 9},
		{8, 2, 5, 4},
		{6, 1, 9, 3},
		{2, 8, 4, 7},
	}
	m["2"] = [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1, 5, 7, 1},
		{9, 1, 1, 9},
	}
	return m
}
func coinTest(t *testing.T) {
	set := []int{1}
	tests := []struct {
		name string
		S    []int
		t    int
		want int
	}{
		{"", set, 1, 1},
		{"", set, 2, 2},
		{"", coinSet(set, 2), 2, 1},
		{"", coinSet(set, 3, 5), 16, 4},
		{"", coinSet(set, 3, 5), 17, 5},
	}
	for _, tt := range tests {
		got := coin(tt.S, tt.t)
		if got != tt.want {
			t.Errorf("got != tt.want %d %d", got, tt.want)
		} else {
			t.Log(got)
		}
	}
}

func coinSet[T comparable](s []T, a ...T) []T {
	seen := make(map[T]struct{}, len(s))

	for _, x := range s {
		seen[x] = struct{}{}
	}

	for _, x := range a {
		if _, ok := seen[x]; !ok {
			seen[x] = struct{}{}
			s = append(s, x)
		}
	}

	return s
}

type jstest struct {
	name           string
	jobs           []job
	wantWeight     int
	wantActivities int
	want           int
}

func jobSelectionTests(weighted bool) []jstest {

	tests := []jstest{
		{
			name: "weighted interval scheduling",
			jobs: []job{
				{start: 1, finish: 3, weight: 5}, // optimal activity
				{start: 2, finish: 5, weight: 6},
				{start: 4, finish: 6, weight: 5}, // optimal activity
				{start: 6, finish: 7, weight: 4}, // optimal activity
				{start: 5, finish: 8, weight: 11},
				{start: 7, finish: 9, weight: 2}, // optimal activity
			},
			wantWeight: 17, wantActivities: 4,
		},
		{
			name: "CLRS activity set with weights",
			jobs: []job{
				{start: 1, finish: 4, weight: 5}, // optimal activity
				{start: 3, finish: 5, weight: 6},
				{start: 0, finish: 6, weight: 5},
				{start: 5, finish: 7, weight: 4}, // optimal activity
				{start: 3, finish: 9, weight: 11},
				{start: 5, finish: 9, weight: 2},
				{start: 6, finish: 10, weight: 7},
				{start: 7, finish: 11, weight: 8}, // optimal activity
				{start: 8, finish: 12, weight: 6},
				{start: 2, finish: 14, weight: 13},
				{start: 12, finish: 16, weight: 9}, // optimal activity
			},
			wantWeight: 27, wantActivities: 4,
		},
	}
	for i, v := range tests {
		if weighted {
			v.want = v.wantWeight
		} else {
			v.want = v.wantActivities
		}
		tests[i] = v
	}
	return tests
}
