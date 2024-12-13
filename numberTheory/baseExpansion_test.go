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
	// Use to check:
	//	Fermat's little theorem,
	//	Eulers's theorem (uses the totient function),
	//	Repeated squaring,
	//	And repeating pattern problems
	tests := []struct {
		name  string
		base  int
		power int
		m     int
		want  int
	}{
		// {"Rosen Cp4.2 Exp12", 3, 644, 645, 36},
		// {"ExamQ", 7, 16, 100, 1},
		// {"tentti_2023-02-24_dunno 4 a)", 2, 32, 11, 4},     // Same as tentti_2022-12-08_Ragnar 4 a)
		// {"tentti_2023-02-24_dunno 4 b)"honen, 3, 35, 17, 10},    // Same as tentti_2022-12-08_Ragnar 4 b)
		// {"tentti_2023-02-24_dunno 4 c)", 9, 100, 21, 9},    // Same as tentti_2022-12-08_Ragnar 4 c). Use repeated squaring
		// {"tentti_2023-12-14_Kohonen2 5 b)", 7, 16, 100, 1}, // Use repeated squaring
		{"tentti_2023-04-19_Kohonen", 9, 33, 11, 3}, // Use repeated squaring
		// {"tentti_2023-12-14_Kohonen2 5 c)", 10, 1200002, 13, 9},
		// {"tentti_2024_04_19_Kohonen 5 b)", 5, 1000, 1000, 625}, // Solve by noticing the repeating pattern for even and odd powers.
		// {"Kohonen slides exp(p304)", 3, 13, 100, 23},
		// {"Kohonen2023 Set6 Lisä 5 a", 2, 16, 11, 9}, // Solve with BOTH Fermat's LT and repeated squaring.
		// {"Kohonen2023 Set6 Lisä 5 b", 3, 19, 13, 3},
		// {"Kohonen2023 Set6 Lisä 5 c", 4, 9999, 120, 64},
		// {"Michael Penn NT12 exp1", 7, 950, 100, 49},
		// {"2018-05-25_Kangaslampi", 4, 2018, 7, 2},
		// {"Kohonen2023 Set6 Lisä 3", 27, 1932, 2021, 1},
		// {"Kohonen2023 Set6 Lisä 3", 27, 205, 2021, 2}, // Find k, j where 11k = 1+j*φ(2021); j, k ∈ integers
		// {"Kohonen2023 Set6 Lisä 3 & tentti_2021-04-15_dunno ex5", 27, 527, 2021, 2},
		// {"Rosen Cp4.4 ex33", 7, 121, 13, 7},
		// {"Rosen Cp4.4 ex34", 23, 1002, 41, 37},
		// {"Rosen Cp4.4 ex36", 5, 40, 41, 1},
		// {"2023 Radnell Set6 HW4", 2023, 2023, 100, 67},
		// {"2023 Radnell Set6 HW4", 23, 2023, 100, 67},
		// {"2023 Radnell Set6 HW4 simplified w/ totient", 23, 23, 100, 67},
		// {"2023 Radnell Set6 AA2 a)", 3, 19, 13, 3},
		// {"2023 Radnell Set6 AA2 b)", 4, 12, 27, 10},
		// {"2023 Radnell Set6 AA2 c)", 12, 27, 15, 3},
		// {"2023 Radnell Set6 AA2 d)", 146, 2, 21, 1},
		// {"Theoretically 2 www.youtube.com/watch?v=r5_65lLtHkA", 3, 547, 825, 537},
		// {"2017-04-25_Kangaslampi 3 a)", 4, 119, 7, 2}, // Reduced power with Eulers formula.
		// {"Theoretically 2 www.youtube.com/watch?v=r5_65lLtHkA", 3, 147, 825, 537}, // Reduced power with Eulers formula.
		// {"", 27, 205, 2021, 2},
		// {"2020-10-20_dunno 4d", 2, 341, 31, 2},
		// {"2020-10-20_dunno 4d", 2, 30, 31, 1},
		// {"2020-10-20_dunno 4d", 2, 69, 31, 2},
		// {"2024_Kohonen_ex_6B6", 2055, 2055, 1000, 375},

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
