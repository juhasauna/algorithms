package numberTheory

import (
	"fmt"
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
		{"modularInverse", modularInverseTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}

}

func modularInverseTest(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{

		{89, 232, 73},
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
		a int
		b int
		s int
		t int
	}{

		{252, 198, 4, -5},
		// {662, 414, []euclidsAlg{{662, 414}, {414, 248}, {248, 166}, {166, 82}, {82, 2}}},
	}
	for _, tt := range tests {
		x := euclidsAlg{tt.a, tt.b}
		if coef1, coef2 := x.bezoutCoefficients(); coef1 != tt.s || coef2 != tt.t {
			t.Errorf("FAIL: got(s=%d, t=%d), want: (s=%d, t=%d)", coef1, coef2, tt.s, tt.t)
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

		{22220, []int{43, 47}},
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

func greatesCommonDivisorTest(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{414, 662, 1},
		// {1, 1, 1},
		// {12, 30, 6},
		// {150, 70, 10},
		// {32, 27, 1},
		// {190, 34, 2},
		// {17, 95, 1},
		// {273, 98, 7},
		// {540, 504, 36},
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
