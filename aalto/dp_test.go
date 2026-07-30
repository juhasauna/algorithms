package aalto

import (
	"testing"
)

func Test_dp(t *testing.T) {

	// activitySelectionTest(t)
	// checkerboardTest(t)
	// coinTest(t)
	// cuttingOrderTest(t)
	// fibonacciTest(t)
	// knapsackTest(t)
	// longestCommonSubsequenceTest(t)
	// longestIncreasingSubsequenceTest(t)
	// longestPalindromeTest(t)
	optimalBSTTest(t)
	// longestPalindromicSubsequenceTest(t)
	// matrixChainMultiplicationTest(t)
	// rodCuttingTest(t)
	// scrambledStringsTest(t)
	// tournamentProbabilityTest(t)
	// weightedIntervalSchedulingTest(t)
}

func optimalBSTTest(t *testing.T) {
	tests := []struct {
		name string
		f    []int
		want int
	}{
		{name: "2", f: []int{10, 20}, want: 40},
		{name: "3", f: []int{34, 8, 50}, want: 142},
		{name: "", f: []int{6, 3, 1}, want: 15},
		{name: "4", f: []int{4, 2, 6, 3}, want: 26},
		{name: "4", f: []int{213, 20, 547, 100, 120}, want: 1573},
		{name: "7", f: []int{4, 2, 6, 3, 5, 1, 7}, want: 63},
		{name: "", f: []int{12, 3, 18, 7, 25, 4, 10, 15, 2, 20, 6, 14, 9, 5, 16}, want: 480},
	}
	got := 0
	for _, tt := range tests {
		got = optimalBST(tt.f)
		t.Logf("got: %d, counter: %d\n", got, counter)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got, counter)
		}

	}
}
func wordSeparationTest(t *testing.T) {
	tests := []struct {
		name  string
		dict  []string
		story string
		want  int
	}{
		{"", []string{"once", "up", "on", "at", "i", "me"}, "onceuponatime", 6},
	}
	for _, tt := range tests {
		got := 0
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got, counter)
		}

	}
}

func knapsackTest(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		values  []int
		cap     int
		want    int
	}{
		{"", []int{1, 2}, []int{10, 10}, 1, 10},
		{"", []int{1, 2}, []int{10, 10}, 2, 10},
		{"", []int{1, 2}, []int{10, 10}, 3, 20},
		{"", []int{10, 15, 20}, []int{4, 20, 30}, 29, 30},
		{"", []int{1, 2, 3, 4, 5}, []int{1, 9, 10, 11, 12}, 10, 31},
		{"", []int{1, 2, 3, 4, 5}, []int{1, 9, 10, 11, 15}, 10, 34},
		{"", []int{10, 20, 30}, []int{60, 100, 120}, 50, 220},
		{"", []int{1, 2, 3, 3, 4}, []int{2, 3, 4, 3, 5}, 7, 10},
	}
	for _, tt := range tests {
		// got := knapsack(tt.weights, tt.values, tt.cap)
		got := knapsackDP(tt.weights, tt.values, tt.cap)
		// got := ks(tt.weights, tt.values, tt.cap)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got)
		}
	}
}
func cuttingOrderTest(t *testing.T) {
	tests := []struct {
		name string
		l    []int
		n    int
		want int
	}{
		{"", []int{1}, 10, 10},
		{"", []int{2}, 10, 10},
		{"", []int{5}, 10, 10},
		{"", []int{3, 9}, 10, 17},
		{"", []int{2, 8, 10}, 20, 38},
	}
	for _, tt := range tests {
		// got := cuttingOrder(tt.l, tt.n)
		got := cuttingOrderBottomUp(tt.l, tt.n)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got)
		}
	}
}
func fibonacciTest(t *testing.T) {
	tests := []struct {
		name string
		k    int
		want int
	}{
		// {"", 3, 2},
		// {"", 4, 3},
		// {"", 5, 5},
		// {"", 6, 8},
		// {"", 7, 13},
		// {"", 8, 21},
		// {"", 9, 34},
		// {"", 10, 55},
		// {"", 11, 89},
		{"", 12, 144},
	}
	for _, tt := range tests {
		// got := fibonacciBottomUpTabulation(tt.k)
		// got := fibonacciBottomUp(tt.k)
		got := fibonacci(tt.k)
		t.Log(got, counter)
		counter = 0
		got = fibonacciMemoization(tt.k)
		t.Log(got, counter)
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got)
		}
	}
}
func scrambledStringsTest(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		// {"", "a", "b", false},
		// {"", "a", "a", true},
		// {"", "aa", "aa", true},
		// {"", "ab", "ab", true},
		// {"", "ab", "ba", true},
		// {"", "abcd", "abcd", true},
		// {"", "abcd", "cdab", true},
		// {"", "abcd", "dcab", true},
		// {"", "abcd", "dcba", true},
		// {"", "abcd", "bacd", true},
		// {"", "abcd", "badc", true},
		// {"", "abcd", "cadb", false},
		// {"", "abcd", "bdac", false},
		// {"", "ABCDE", "ABCDE", true},
		// {"", "abcdefgh", "efghcdab", true},
		{"", "abcdefgh", "hgfedcba", true},
		// {"", "abcdefgh", "hedfcbag", true},
		// {"", "abcdefgh", "acdehbfg", false},
	}
	for _, tt := range tests {
		got := scrambledStrings(tt.a, tt.b)
		t.Log(counter)
		counter = 0
		got = scrambledStringsMemoization(tt.a, tt.b)
		t.Log(counter)
		if got != tt.want {
			t.Errorf("got/want (%t/%t)", got, tt.want)
		} else {
			// t.Log(got)
		}
	}
}
func longestPalindromeTest(t *testing.T) {
	tests := []struct {
		name    string
		x       string
		wantLen int
		want    string
	}{
		// {"", "a", 1, "a"},
		// {"", "aa", 2, "aa"},
		// {"", "aaa", 3, "aaa"},
		// {"", "aaaa", 4, "aaaa"},
		// {"", "abc", 1, "a"},
		// {"", "abca", 1, "a"},
		// {"", "abccba", 6, "abccba"},
		// {"", "abccbaaa", 6, "abccba"},

		// {"", "abcba", 5, "abcba"},
		// {"", "abcbaaa", 5, "abcba"},
		// {"", "bbbabcbaaa", 5, "abcba"},

		// {"", "forgeeksskeegfor", 10, "geeksskeeg"},
		{"", "1212saippuakauppias8989", 15, "saippuakauppias"},
		{"", "1234saippuaXXXkauppias6789", 3, "XXX"},
	}
	gotStr := ""
	for _, tt := range tests {
		// gotStr = longestPalindrome(tt.x)
		// t.Logf("lonPal1:\t%s, counter: %d, \t%s\n", tt.want, counter, gotStr)
		// counter = 0
		gotStr = longestPalindrome2(tt.x)
		t.Logf("lonPal2:\t%s, counter: %d, \t%s\n", tt.want, counter, gotStr)
		counter = 0
		gotStr = longestPalindromeTopDown(tt.x)
		t.Logf("lonPalMemo:\t%s, counter: %d, \t%s\n", tt.want, counter, gotStr)
		counter = 0
		// gotStr = lonPalNaive(tt.x)
		// t.Logf("naive:\t%s, counter: %d, \t%s\n", tt.want, counter, gotStr)
		counter = 0
		if gotStr != tt.want || len(gotStr) != tt.wantLen {
			t.Errorf("FAIL: got/want (%s/%s)", gotStr, tt.want)
		} else {
			// t.Log(gotStr)
		}
		// if got != tt.wantLen {
		// 	t.Errorf("got/want (%d/%d)", got, tt.wantLen)
		// } else {
		// 	t.Log(got)
		// }
	}
}

func weightedIntervalSchedulingTest(t *testing.T) {
	tests := jobSelectionTests(true)
	got := 0
	counter = 0
	for _, tt := range tests {
		got = weightedIntervalScheduling(tt.jobs)
		t.Logf("%d, counter: %d, \t%d\n", tt.want, counter, got)
		counter = 0
		got = wisProfBryce(tt.jobs)
		t.Logf("%d, counter: %d, \t%d\n", tt.want, counter, got)
		counter = 0
		got = wis(tt.jobs)
		t.Logf("%d, counter: %d, \t%d\n", tt.want, counter, got)
		counter = 0
		if got != tt.want {
			t.Errorf("got/want (%d/%d)", got, tt.want)
		} else {
			// t.Log(got, counter)
		}
		t.Logf("\n")
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
		{"", []int{1, 3}, 3},
		{"", []int{1, 3, 10}, 10},
		{"", []int{1, 6, 7, 8}, 12},
		{"", []int{2, 3, 4, 5}, 8},
		{"", []int{1, 4, 9, 9}, 10},
		{"", []int{1, 7, 8, 9}, 14},
		{"", []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 30},
		{"", []int{1, 5, 8, 15, 16, 17, 17, 20, 24, 30}, 35},
		// {"naive: 70s, iters: 34 359 738 368", []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 20, 23, 29, 30, 36, 42, 44, 49, 55, 56, 61, 66, 69, 70, 75, 78, 80, 82, 87, 91, 91, 91, 92, 94, 95, 99}, 104},
	}
	got := 0
	for _, tt := range tests {

		got = rodCutting(tt.x)
		t.Log("rodCutting", counter, got)
		counter = 0
		got = rodCuttingMemoized(tt.x)
		t.Log("memo", counter, got)
		counter = 0
		got = rodCuttingNaive(tt.x)
		t.Log("naive", counter, got)
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
		// {"", []int{2, 2, 2}, 8},
		// {"", []int{2, 3, 1}, 6},
		{"", []int{2, 3, 5}, 30},
		{"", []int{10, 30, 5, 60}, 4500},
		{"", []int{30, 35, 15, 5, 10, 20, 25}, 15125},
		{"", []int{5, 10, 3, 12, 5, 50, 6}, 2010},
	}

	for _, tt := range tests {
		got := matrixChainMultiplication(tt.x)
		// got := mcm(tt.x)
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
		{"", "cab", "abc", 2},
		{"", "ABCD", "ACBD", 3},
		{"alg2021_final", "abcabcabc", "aaabbbccc", 5},
		{"", "ABCBDAB", "BDCABAA", 4},
		{"allow for diff. lenghts", "abcabcabc", "abccc", 5},
		{"allow for diff. lenghts", "abcabcabc", "abbcc", 5},
	}
	for _, tt := range tests {
		// got := longestCommonSubsequence(tt.x, tt.y)
		// got := lcsSlow(tt.x, tt.y)
		got := lcsMemoization(tt.x, tt.y)

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
		{"", "forgeeksskeegforracecarxyz", 12},
		{"", "qqforgeeksskeegforracecarxyz", 12},
	}
	for _, tt := range tests {
		// got := longestPalindromicSubsequence(tt.x)
		got := lps(tt.x)
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
		name   string
		S      []int
		target int
		want   int
	}{
		{"", set, 1, 1},
		{"", set, 2, 2},
		{"", coinSet(set, 2), 2, 1},
		{"", coinSet(set, 3, 5), 16, 4},
		{"", coinSet(set, 3, 5), 17, 5},
		{"", coinSet(set, 2, 5), 43, 10},
		{"", coinSet(set, 2, 3, 5, 7, 11), 43, 5},
	}
	got := 0
	for _, tt := range tests {
		// got = coinRec(tt.S, tt.target)
		// t.Logf("coinRec %d\t%d", got, counter)
		// counter = 0
		got = coinMemo(tt.S, tt.target)
		t.Logf("coinMemo\t%d\tcounter: %d", got, counter)
		counter = 0
		got = coinDP(tt.S, tt.target)
		t.Logf("coinDP\t%d\tcounter: %d", got, counter)
		counter = 0
		if got != tt.want {
			t.Errorf("got != tt.want %d %d", got, tt.want)
		} else {
			// t.Log(got)
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
			name: "ProfBryce",
			jobs: []job{
				{start: -1, finish: 1, weight: 2},
				{start: -2, finish: 2, weight: 2},
				{start: 0, finish: 3, weight: 3},
				{start: 1, finish: 4, weight: 2},
				{start: 2, finish: 5, weight: 1},
				{start: 3, finish: 6, weight: 1},
				{start: 5, finish: 7, weight: 3},
				{start: 6, finish: 8, weight: 1},
				{start: 7, finish: 9, weight: 4},
				{start: 4, finish: 10, weight: 5},
				{start: 8, finish: 11, weight: 3},
			},
			wantWeight: 11, wantActivities: -1,
		},
		{
			name: "1",
			jobs: []job{
				{start: 1, finish: 3, weight: 5},
				{start: 2, finish: 5, weight: 6},
				{start: 4, finish: 6, weight: 5},
				{start: 6, finish: 7, weight: 4},
				{start: 5, finish: 8, weight: 11},
			},
			wantWeight: 17, wantActivities: -1,
		},
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
