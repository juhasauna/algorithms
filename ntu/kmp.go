package ntu

import (
	"strconv"
)

type KMP struct {
	iters    int64
	optimize bool
}

func (x *KMP) stringMatchKMP(haystack, needle string) int {
	// fmt.Printf("stringMatchKMP: from/search %s/%s", from, search)
	if len(haystack) < len(needle) {
		return 0
	} else if len(haystack) == len(needle) {
		if haystack == needle {
			return 1
		}
		return 0
	}
	if needle == "" {
		return 1
	}
	if len(needle) == 1 {
		if string([]rune(haystack)[0]) == needle {
			return 1
		}
	}
	// The book uses 1 based indexing.
	next := x.computeKMPNext(needle)
	m := len(needle)
	i, j := 0, 1
	for i < len(haystack) && j < m+1 {
		x.iters++
		if needle[j-1] == haystack[i] {
			j++
			i++
		} else {
			x.iters++
			j = next[j-1] + 1
			if j == 0 {
				j = 1
				i++
			}
		}
	}
	if j == m+1 {
		return i - (m - 1)
	}
	return 0
}

// next AKA LongestPrefixSuffix (LPS) := longest proper prefix which is also a suffix. A proper prefix is a prefix that doesn’t include whole string.
func (x *KMP) computeKMPNext(needle string) []int {
	length := len(needle)
	// fmt.Println("len: ", length, "text: ", text)
	next := make([]int, length)
	next[0], next[1] = -1, 0
	for i := 2; i < length; i++ {

		j := next[i-1]
		if needle[i-1] == needle[j] {
			next[i] = j + 1
			continue
		}

		j++
		// j := next[i-1] + 1
		for needle[i-1] != needle[j-1] {
			x.iters++
			j = next[j-1] + 1
			if j == 0 {
				break
			}
		}
		next[i] = j
	}
	if x.optimize {
		x.optimizeKMPNextResult(needle, next)
	}
	return next
}

// See HW 2024/6 Ex.5 alg2024hw6_s.pdf
func (x *KMP) optimizeKMPNextResult(needle string, next []int) {
	for i := 1; i < len(needle); i++ {
		j := next[i] + 1
		x.iters++
		if needle[i] == needle[j-1] {
			next[i] = next[j-1]
		}
	}
}

func (x *KMP) computeKMPNextGemini(pattern string) []int {
	length := len(pattern)
	if length == 0 {
		return []int{}
	}

	// lps is the array that will hold the Longest Proper Prefix suffix values for the pattern.
	lps := make([]int, length)
	lps[0] = 0 // lps[0] is always 0.

	// 'currentLPSLength' keeps track of the length of the previous longest prefix suffix.
	currentLPSLength := 0

	// The loop calculates lps[i] for i from 1 to length-1.
	i := 1
	for i < length {
		// Case 1: The characters match.
		// We found a longer prefix that is also a suffix.
		if pattern[i] == pattern[currentLPSLength] {
			currentLPSLength++
			lps[i] = currentLPSLength
			i++
		} else { // Case 2: The characters do not match.
			if currentLPSLength != 0 {
				// This is the key step of the KMP algorithm. We do not restart from scratch.
				// We fall back to the LPS value of the previous state, effectively trying a shorter prefix.
				// We do NOT increment 'i' here, as we need to re-evaluate pattern[i] against the new, shorter prefix.
				currentLPSLength = lps[currentLPSLength-1]
			} else {
				// If currentLPSLength is 0, it means we couldn't find a shorter prefix to match.
				// We set lps[i] to 0 and move to the next character in the pattern.
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// returns the smallest m F(m) where the bitPattern appears. If it does not appear for n then returns -1
func (x *KMP) FindFibonacciWordSequence(bitPattern string, n int) int {
	if bitPattern == "0" && n >= 0 {
		return 0
	}
	if bitPattern == "1" && n >= 1 {
		return 1
	}
	if !(bitPattern != "" && n >= 2) {
		return -1
	}
	p_len := len(bitPattern)
	m := 2
	f, f1, f2 := "10", "1", "0"
	l, l1, l2 := 2, 1, 1
	for l < p_len {
		m++
		f, f1, f2 = f1+f2, f, f1
		l, l1, l2 = l1+l2, l, l1
	}
	bitPattern = ""
	for _, v := range x.computeKMPNext(bitPattern)[1:] {
		bitPattern += strconv.Itoa(v)
	}
	found, i := false, 0
	for !found && i < 4 {
		kmpResult := x.stringMatchKMP(f, bitPattern)
		if kmpResult != -1 {
			found = true
		} else {
			i++
			f, f1, f2 = f1+f2, f, f1
		}
	}
	if found {
		return m + i
	}
	return -1
}
