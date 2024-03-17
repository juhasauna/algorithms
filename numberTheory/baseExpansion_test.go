package numberTheory

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_baseExpansion(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"decToBase", decToBaseTest},
		{"modularExponentiation", modularExponentiationTest},
		// {"onesComplement", onesComplementTest},
		// {"addBinary", addBinaryTest},
		// {"trialDivision", trialDivisionTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}

}

func modularExponentiationTest(t *testing.T) {
	tests := []struct {
		name  string
		base  int
		power int
		m     int
		want  int
	}{
		// {"Rosen Cp4.2 Exp12", 3, 644, 645, 36},
		// {"ExamQ", 7, 16, 100, 1},
		{"tentti_2023-02-24_dunno 4 a)", 2, 32, 11, 4},
		{"tentti_2023-02-24_dunno 4 b)", 3, 35, 17, 10},
		{"tentti_2023-02-24_dunno 4 c)", 9, 100, 21, 9},
	}
	for _, tt := range tests {
		if got := modularExponentiation(tt.base, tt.power, tt.m); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: modularExponentiation() = %v, want %v", tt.name, got, tt.want)
		} else {
			fmt.Println(got)
		}
	}
}
func decToBaseTest(t *testing.T) {
	tests := []struct {
		base int
		dec  int
		want []int
	}{
		{4, 158, []int{2, 1, 3, 2}},
		{2, 6, []int{1, 1, 0}},
		{16, 6, []int{6}},
		{16, 15, []int{15}},
		{16, 16, []int{1, 0}},
		{16, 31, []int{1, 15}},
		{16, 32, []int{2, 0}},
		{2, 231, []int{1, 1, 1, 0, 0, 1, 1, 1}},
	}
	for _, tt := range tests {
		if got := decToBase(tt.dec, tt.base); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("decToBase() = %v, want %v", got, tt.want)
		} else {
			fmt.Println(got)

		}
	}
}

// One’s complement representations of integers are used to
// simplify computer arithmetic. To represent positive and negative integers with absolute value less than 2n−1, a total of n bits
// is used. The leftmost bit is used to represent the sign. A 0 bit
// in this position is used for positive integers, and a 1 bit in this
// position is used for negative integers. For positive integers,
// the remaining bits are identical to the binary expansion of the
// integer. For negative integers, the remaining bits are obtained
// by first finding the binary expansion of the absolute value of
// the integer, and then taking the complement of each of these
// bits, where the complement of a 1 is a 0 and the complement
// of a 0 is a 1.
func onesComplementTest(t *testing.T) {
	tests := []struct {
		dec  int
		want []int
	}{
		{-6, []int{1, 0, 0, 1}},
		{6, []int{0, 1, 1, 0}},
		{-127, []int{1, 0, 0, 0, 0, 0, 0, 0}},
		{231, []int{0, 1, 1, 1, 0, 0, 1, 1, 1}},
		{104, []int{0, 1, 1, 0, 0, 1, 1, 1}},
	}
	for _, tt := range tests {
		if got := onesComplement(tt.dec); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("decToBase() = %v, want %v", got, tt.want)
		} else {
			fmt.Println(got)

		}
	}
}

func addTwosComplentTest(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want []int
	}{
		// {-5, 3, []int{1, 1, 0, 1}},
		// {-4,-3,[]int{1, 0, 0, 0}},
		{6, -6, []int{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		if got := addTwosComplent(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("decToBase() = %v, want %v", got, tt.want)
		} else {
			fmt.Println(got)

		}
	}
}
func addBinaryTest(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want []int
	}{

		{1, 5, []int{1, 1, 0}},
	}
	for _, tt := range tests {
		if got := addBinary(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("decToBase() = %v, want %v", got, tt.want)
		} else {
			fmt.Println(got)

		}
	}
}
