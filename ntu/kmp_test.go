package ntu

import (
	"slices"
	"strings"
	"testing"
)

func Test_kmp(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"LPS", computeNextTest},
		// {"KMP", stringMatchKMPTest},
		// {"FindFibonacciWordSequence", FindFibonacciWordSequenceTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func stringMatchKMPTest(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
	}{
		{"juha_minimal_optimization_example", "ABABABABABABABABABABABAA", "AA"},
		{"juha_minimal_optimization_example", "AABAABAABAABAABAABAABAABAABAABAABAAA", "AAA"},
		{"juha_minimal_optimization_example", "AAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAA", "AAAAAAAAAAAAAAAAAAAAA"},
		{"juha_optimization_example_38%", "AAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAA", "AAAAAAAAAAAAAAAAAAAAA"},
		{"juha_optimization_example_75%", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
		{"juha_long2", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAxxxxxxxxxxxxxxxyxxyxyxyyxyxyxyyxyxyxxxyxyyxyxyxxxyxyyxyxyxxAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxy", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxy"},
		{"juha_optimization_demo about 30% improvement", "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxAyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxy", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxyxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxy"},
		{"Manber exp.", "xyxxyxyxyyxyxyxyyxyxyxx", "xyxyyxyxyxx"},
		{"ByteQuest VOD exp.", "BABABABABCABABCABAB", "ABABCABAB"},
		{"ByteQuest VOD exp. simplified", "BABABABABCABABCAB", "ABABCABAB"},
		{"ByteQuest VOD exp. simplified2", "BABABABABCABABCAB", "ABABCAB"},
		{"Lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lobortis rhoncus blandit. Aenean sagittis enim nec nulla fringilla, quis posuere odio cursus. Curabitur vel mattis neque, ac imperdiet metus. Aliquam semper ullamcorper sodales. Maecenas placerat elit felis, ut dapibus sapien venenatis a. Sed condimentum mattis urna sed dictum. Vivamus maximus felis eget libero auctor, in consequat massa scelerisque. Donec lacinia neque tincidunt feugiat rhoncus. Donec vel scelerisque ante. Nulla blandit, sem eget ultrices faucibus, justo risus condimentum ex, nec lobortis lorem erat ac sapien.", "Curabitur vel mattis neque, ac imperdiet metu"},
		{"1-9", "1987677354208080808080843831211112333341233334123422221111123441233412341234123462505791244444444411112333341233334123422221111123444444675708080808", "1111233334123333412342222111112344123"},
		{"simple match", "AAAAA", "AAAA"},
		{"simple match", "AB", "A"},
		{"no match", "AAABBB", "AABA"},
		{"no match2", "AABA", "AAABBB"},
		{"no match3", "A", "AA"},
	}
	for _, tt := range tests {
		want := strings.Index(tt.haystack, tt.needle) + 1
		x := KMP{optimize: false}
		got := x.stringMatchKMP(tt.haystack, tt.needle)
		if got != want {
			t.Errorf("FAIL (%s) got: %d, want: %d", tt.name, got, want)
			continue
		}
		x2 := KMP{optimize: true}
		gotOpt := x2.stringMatchKMP(tt.haystack, tt.needle)
		if gotOpt != want {
			t.Errorf("FAIL (opt %s) got: %d, want: %d", tt.name, gotOpt, want)
			continue
		}
		ratio := float64(x2.iters) / float64(x.iters)
		t.Logf("SUCCESS (%s) got: %v, iters: %d/%d (%.3f)", tt.name, got, x.iters, x2.iters, ratio)

	}
}
func computeNextTest(t *testing.T) {
	tests := []struct {
		name     string
		needle   string
		optimize bool
		want     []int
	}{
		// {"alg2024mid_s", "bbbabbaba", false, []int{-1, 0, 1, 2, 0, 1, 2, 0, 1}},
		{"alg2024mid_s", "bbbabbaba", true, []int{-1, -1, -1, 2, -1, -1, 2, -1, 1}},
		// {"alg2024mid_s", "bbbabbbaa", false, []int{-1, 0, 1, 2, 0, 1, 2, 3, 4}},
		// {"alg2024mid_s_opt.", "bbbabbbaa", true, []int{-1, -1, -1, 2, -1, -1, -1, 2, 4}},
		// {"alg2022mid_s", "abbaabbaa", false, []int{-1, 0, 0, 0, 1, 1, 2, 3, 4}},
		// {"alg2022mid_s_opt", "abbaabbaa", true, []int{-1, 0, 0, -1, 1, 0, 0, -1, 1}},
		// {"AAACAAAA", "AAACAAAA", false, []int{-1, 0, 1, 2, 0, 1, 2, 3}},
		// {"alg2018mid_s", "abaabaaba", false, []int{-1, 0, 0, 1, 1, 2, 3, 4, 5}},
		// {"alg2018mid_s", "abaabaabb", false, []int{-1, 0, 0, 1, 1, 2, 3, 4, 5}},
		// {"24/hw6/5", "abaababaa", false, []int{-1, 0, 0, 1, 1, 2, 3, 2, 3}},
		// {"24/hw6/5 optimized", "abaababaa", true, []int{-1, 0, -1, 1, 0, -1, 3, -1, 1}},
		// {"ByteQuest VOD exp.", "ABABCABAB", false, []int{-1, 0, 0, 1, 2, 0, 1, 2, 3}},
		// {"ByteQuest VOD exp.", "ABABCABAB", true, []int{-1, 0, -1, 0, 2, -1, 0, -1, 0}},
		// {"Manber exp.", "xyxyyxyxyxx", false, []int{-1, 0, 0, 1, 2, 0, 1, 2, 3, 4, 3}},
		// {"Manber exp. optimized", "xyxyyxyxyxx", true, []int{-1, 0, -1, 0, 2, -1, 0, -1, 0, 4, 3}},

		// {"", "AAAA", false, []int{-1, 0, 1, 2}},
		// {"", "AAAAA", false, []int{-1, 0, 1, 2, 3}},
		// {"", "AAAAA", true, []int{-1, -1, -1, -1, -1}},
		// {"", "AAAABA", false, []int{-1, 0, 1, 2, 3, 0}},
		// {"", "AAAABA", true, []int{-1, -1, -1, -1, 3, -1}},
		// {"", "xyxyy", false, []int{-1, 0, 0, 1, 2}},
		// {"", "xyxyyx", false, []int{-1, 0, 0, 1, 2, 0}},

		// {"1-4", "1111233334123333412342222111112344123", false, []int{-1, 0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 4, 5, 6, 0, 0, 1, 0}},

		// {"USING THE HAYSTACK Manber exp.", "xyxxyxyxyyxyxyxyyxyxyxx", false, []int{-1, 0, 0, 1, 1, 2, 3, 2, 3, 2, 0, 1, 2, 3, 2, 3, 2, 0, 1, 2, 3, 2, 3}},
		// {"USING THE HAYSTACK ByteQuest VOD exp.", "BABABABABCABABCABAB", false, []int{-1, 0, 0, 1, 2, 3, 4, 5, 6, 7, 0, 0, 1, 2, 3, 0, 0, 1, 2}},
	}
	for _, tt := range tests {
		x := KMP{optimize: tt.optimize}
		got := x.computeKMPNext(tt.needle)
		if !slices.Equal(got, tt.want) {
			t.Errorf("FAIL (%s) got: %v, want: %v", tt.name, got, tt.want)
		} else {
			t.Logf("SUCCESS %v", got)
		}

	}
}

func FindFibonacciWordSequenceTest(t *testing.T) {
	tests := []struct {
		name       string
		bitPattern string
		length     int
	}{
		{"", "1", 123},
		{"", "1101", 5},
	}
	x := KMP{}
	for _, tt := range tests {
		got := x.FindFibonacciWordSequence(tt.bitPattern, tt.length)
		t.Log(got)
	}
}
