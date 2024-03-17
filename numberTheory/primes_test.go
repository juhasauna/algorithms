package numberTheory

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func Test_primes(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"primeFactorization", primeFactorizationTest},
		// {"leastCommonMult", leastCommonMultTest},
		// {"greatesCommonDivisor", greatesCommonDivisorTest},
		// {"euclideanAlgorithmSteps", euclideanAlgorithmStepsTest},
		// {"modularInverse", modularInverseTest},
		// {"mod", divModTest},
		// {"bezoutCoefficients", bezoutCoefficientsTest},
		{"eulersTotient", eulersTotientTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func eulersTotientTest(t *testing.T) {
	tests := []struct {
		a    int
		want int
	}{
		{10, 4},
		{21, 12},
	}
	for _, tt := range tests {
		got := eulersTotient(tt.a)
		if got != tt.want {
			t.Errorf("got = %d != %d = want", got, tt.want)
		}
	}
}

func divModTest(t *testing.T) {
	tests := []struct {
		a       int
		b       int
		wantDiv int
		wantMod int
	}{
		{-17, 2, -9, -1},
		{144, 7, 20, 4},
		{-101, 13, -8, 3},
		{199, 19, 10, 9},
	}
	for _, tt := range tests {
		d, m := divMod(tt.a, tt.b)
		fmt.Println(d, m)
	}
}

func modularInverseTest(t *testing.T) {
	// Computing the inverse of a modulo b. The inverse is the bezout coeffiecient of a (call it s).
	// Also any integer congruent to s mod b is also an inverse. That is, s + k * b, where k is any integer.
	// See Rosen 7e 4.4 Theorem 1 p.296
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{

		{"", 89, 232, 73},
		{"Rosen Cp4.4 Ex1", 7, 26, -11},
		{"Rosen Cp4.4 Exp2", 4620, 101, 1601},
		{"Rosen Cp4.4 Ex2", 2436, 13, 937},
		{"Rosen Cp4.4 Ex3", 9, 4, -2},
		{"Rosen Cp4.4 Ex4", 17, 2, -8},
		{"Rosen Cp4.4 Ex5a", 4, 9, -2},
	}
	for _, tt := range tests {
		x := euclidsAlg{tt.a, tt.b}
		if got := x.modularInverse(); got != tt.want {
			t.Errorf("FAIL: got: %d, tt: %+v", got, tt)
		}
	}
}
func bezoutCoefficientsTest(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		s    int
		t    int
	}{
		// {"Rosen Cp4.4 Ex1", 26, 7, 3, -11},
		// {"Rosen Cp4.4 Ex2", 2436, 13, -5, 937},
		// {"Rosen Cp4.4 Exp1", 7, 3, 1, -2},
		// {"Rosen Cp4.4 Exp2", 4620, 101, -35, 1601},
		// {"Rosen Cp4.3 Exp17", 252, 198, 4, -5},
		// {"Rosen Cp4.3 Ex44. TODO: Check with pen&paper", 1001, 100001, 10, -999},
		// {"Kohonen2023 slides syt(162, 114)", 162, 114, -7, 10},
		// {"Diophantine eq: 27x + 11y = 1", 27, 11, -2, 5},
		// {"Diophantine eq: 514x + 387y = 1", 514, 387, 64, -85},
		{"Ragnar lecture exp no-sol-diophantine-eq", 112, 49, -3, 7},
		// {"Kohonen2023 ExSet6 P3", 2331, 2037, 7, -8},
		// {"",662, 414, []euclidsAlg{{662, 414}, {414, 248}, {248, 166}, {166, 82}, {82, 2}}},
	}
	for _, tt := range tests {
		if tt.a < tt.b {
			t.Error("first arg should be bigger than the first")
		}
		x := euclidsAlg{tt.a, tt.b}
		if coef1, coef2 := x.bezoutCoefficients(); coef1 != tt.s || coef2 != tt.t {
			t.Errorf("FAIL: got(s=%d, t=%d), want: (s=%d, t=%d)", coef1, coef2, tt.s, tt.t)
		} else {
			gcd := tt.a*coef1 + tt.b*coef2
			t.Logf("SUCCESS: gdc(%d, %d) = %d * %d + (%d * %d) = %d", tt.a, tt.b, coef1, tt.a, coef2, tt.b, gcd)
		}
	}
}
func euclideanAlgorithmStepsTest(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want []euclidsAlg
	}{

		{662, 414, []euclidsAlg{{662, 414}, {414, 248}, {248, 166}, {166, 82}, {82, 2}}},
	}
	for _, tt := range tests {
		x := euclidsAlg{tt.a, tt.b}
		if got := x.getSteps(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL: %+v, want %+v", got, tt.want)
		} else {
			fmt.Println(got)
		}
	}
}
func trialDivisionTest(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{

		// {5, true},
		{6, false},
		// {11, true},
		// {13, true},
		{14, false},
		// {173, true},
	}
	for _, tt := range tests {
		if got := trialDivision(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("trialDivisionTest(%d) = %t, want %t", tt.n, got, tt.want)
		} else {
			fmt.Println(got)
		}
	}
}
func primeFactorizationTest(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{

		// {22220, []int{43, 47}},
		// {2021, []int{43, 47}},
		// {101, []int{101}},
		// {635, []int{5, 127}},
		// {121, []int{11,11}},
		// {5, []int{5}},
		// {7007, []int{7, 7, 11, 13}},
		// {92928, []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 11, 11}},
		// {123552, []int{2, 2, 2, 2, 2, 3, 3, 3, 11, 13}},
	}
	for _, tt := range tests {
		if got := primeFactorization(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("primeFactorization(%d) = %v, want %v", tt.n, got, tt.want)
		} else {
			fmt.Println(got)
		}
	}
}
func P(a, b int) int {
	x := math.Pow(float64(a), float64(b))
	return int(x)
}

func greatesCommonDivisorTest(t *testing.T) {
	tests := []struct {
		name   string
		a      int
		b      int
		expect int
	}{
		{"Rosen Cp4.3 Exp17", 252, 198, 18},
		// {"Kohonen2023 ExSet6 Pre3", 2331, 2037, 21},
		// {"Kohonen2023 slides syt(162, 114)", 162, 114, 6},
		// {"Rosen Cp4.3 Ex32 a", 5, 1, 1},
		// {"Rosen Cp4.3 Ex32 b", 100, 101, 1},
		// {"Rosen Cp4.3 Ex32 c", 123, 277, 1},
		// {"Rosen Cp4.3 Ex32 d", 1529, 14039, 139},
		// {"Rosen Cp4.3 Ex32 e", 1529, 14038, 1},
		// {"Rosen Cp4.3 Ex32 f", 111111, 11111, 1},
		// {"Rosen Cp4.3 Ex24 a", P(2, 2) * P(3, 3) * P(5, 5), P(2, 5) * P(3, 3) * P(5, 2), 2700},
		// {"Rosen Cp4.3 Ex24 d", 2 * 2 * 7, 5 * 5 * 5 * 13, 1},
		// {"Rosen Cp4.3 Ex24 e", 0, 5, 5},
		// {"Rosen Cp4.3 Ex24 f", 2 * 3 * 5 * 7, 2 * 3 * 5 * 7, 2 * 3 * 5 * 7},
		// {"Rosen Cp4.3 Ex24 b", 111111, 11111, 1},
		// {"Rosen Cp4.4 Exp2", 101, 4620, 1},
		{"Rosen Cp4.4 Ex1", 26, 7, 1},

		// {"", 414, 662, 2},
		// {"", 1, 1, 1},
		// {"", 12, 30, 6},
		// {"", 150, 70, 10},
		// {"", 32, 27, 1},
		// {"", 190, 34, 2},
		// {"", 17, 95, 1},
		// {"", 273, 98, 7},
		// {"", 540, 504, 36},
		// {"Cp4.3 Ex24 b", 2 * 3 * 5 * 7 * 11 * 13, P(2, 11) * P(3, 4) * 11 * 17, 2 * 3 * 11},
		// {"TOO SLOW DONT RUN THIS Rosen Cp4.3 Ex24 b", 2 * 3 * 5 * 7 * 11 * 13, P(2, 11) * P(3, 9) * 11 * P(17, 14), 2 * 3 * 11},
	}
	for i, tt := range tests {
		got := greatestCommonDivisor(tt.a, tt.b)
		// got2 := got
		got2 := euclideanAlgorithm(tt.a, tt.b, true)
		if !(got == tt.expect && got2 == tt.expect) {
			t.Errorf("FAIL: inputs %d, %d, expected %d; got %d; got2 %d", tt.a, tt.b, tt.expect, got, got2)
		} else {
			fmt.Printf("\tSUCCESS: %d, inputs (%d,\t%d),\tgot %v\n", i, tt.a, tt.b, got)
		}
	}
}
func leastCommonMultTest(t *testing.T) {
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
			t.Errorf("FAIL: inputs %d, %d, expected %v; got %v", tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}
