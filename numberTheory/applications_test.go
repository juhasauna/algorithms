package numberTheory

import (
	"reflect"
	"testing"
)

func Test_applications(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"isbn", isbnIsValidTest},
		// {"hash", hashTest},
		// {"linearCongruential", linearCongruentialTest},
		// {"countDigits", countDigitsTest},
		// {"middleSquare", middleSquareTest},
		// {"blumBlumShub", blumBlumShubTest},
		{"solveCongruences", solveCongruencesTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
func solveCongruencesTest(t *testing.T) {

	tests := []struct {
		mods []int // If relatively prime positive integers > 1, then has unique solution.
		rems []int
		want int
	}{
		{[]int{3, 5, 7}, []int{2, 3, 2}, 23},
		// {[]int{5, 6, 7}, []int{1, 2, 3}, 206},
	}
	for i, tt := range tests {
		got := solveCongruences(tt.mods, tt.rems)
		if got != tt.want {
			t.Errorf("FAIL %d: input: %+v, got: %d", i, tt, got)
		}
	}
}
func countDigitsTest(t *testing.T) {

	// println(countDigits(123))
}

func blumBlumShubTest(t *testing.T) {

	tests := []struct {
		mod        int // a prime
		mult       int // "d" such that p does not divide d
		increment  int // not used
		seed       int
		iterations int
		want       int
	}{
		// {7, 3, 0, 2, 10, 1},
		{11, 3, 0, 2, 10, 6},
	}
	for i, tt := range tests {
		x := rng{tt.mod, tt.mult, tt.increment, tt.seed, tt.iterations}
		got := x.blumBlumShub()
		if got != tt.want {
			t.Errorf("FAIL %d: input: %+v, got: %d, want: %d", i, x, got, tt.want)
		}
	}
}
func middleSquareTest(t *testing.T) {
	tests := []struct {
		mod        int
		mult       int
		increment  int
		seed       int
		iterations int
		want       int
	}{
		// {0, 0, 0, 2357, 10, 0},
		// {0, 0, 0, 3792, 10, 0}, // bad choice #1
		{0, 0, 0, 2916, 10, 0}, // bad choice #2
	}
	for i, tt := range tests {
		x := rng{tt.mod, tt.mult, tt.increment, tt.seed, tt.iterations}
		got := x.middleSquare()
		if got != tt.want {
			t.Errorf("FAIL %d: input: %+v, got: %d, want: %d", i, x, got, tt.want)
		}
	}
}
func linearCongruentialTest(t *testing.T) {
	tests := []struct {
		mod        int
		mult       int
		increment  int
		seed       int
		iterations int
		want       int
	}{
		{9, 7, 4, 3, 0, 3},
		{9, 7, 4, 3, 1, 7},
		{9, 7, 4, 3, 2, 8},
		{9, 7, 4, 3, 9, 3},
	}
	for i, tt := range tests {
		x := rng{tt.mod, tt.mult, tt.increment, tt.seed, tt.iterations}
		got := x.linearCongruential()
		if got != tt.want {
			t.Errorf("FAIL %d: input: %+v, got: %d, want: %d", i, x, got, tt.want)
		}
	}
}
func isbnIsValidTest(t *testing.T) {
	tests := []struct {
		num  string
		want bool
	}{
		{"0072880082", true},
		{"084930149X", false},
		{"984930149X", true},
	}
	for _, tt := range tests {
		x := isbn{tt.num}
		got := x.isValid()
		if got != tt.want {
			t.Errorf("FAIL: input: %v, got: %t, want: %t", tt.num, got, tt.want)
		}
	}
}
func hashTest(t *testing.T) {
	tests := []struct {
		mod    int
		values []int
		want   map[int]int
	}{
		{31, []int{317, 918, 007, 100, 111, 310}, map[int]int{0: 310, 7: 317, 8: 7, 9: 100, 18: 111, 19: 918}},
	}
	for _, tt := range tests {
		x := hash{tt.mod, tt.values, make(map[int]int)}
		x.get()
		if !reflect.DeepEqual(x.locations, tt.want) {
			t.Errorf("FAIL: mod: %d, got: %v, want: %v", tt.mod, x.locations, tt.want)
		}
	}
}
