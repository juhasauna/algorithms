package aalto

import "testing"

func Test_dp(t *testing.T) {

	// coinTest(t)
	// checkerboardTest(t)
	// tournamentProbabilityTest(t)
	longestCommonSubsequenceTest(t)
	// longestPalindromicSubsequenceTest(t)

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
