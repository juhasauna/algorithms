package kolman

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		f    func(string, *testing.T)
		do   bool
	}{
		{"intUnion", intUnionTest, true},
		{"intIntersection", intIntersectionTest, true},
		{"intDifference", intDifferenceTest, true},
		{"appC_C01_ex4", appC_C01_ex4Test, true},
		{"leastCommonIntMult", leastCommonIntMultTest, true},
	}
	for _, tt := range tests {
		if tt.do {
			tt.f(tt.name, t)
		}
	}
}

func intUnionTest(name string, t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{-1, -3, -3}, []int{-3, -4, -5}, []int{-1, -3, -4, -5}},
	}
	for i, tt := range tests {
		got := intUnion(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func intIntersectionTest(name string, t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{3}},
		{[]int{-1, -3, -3, -4}, []int{-3, -4, -5}, []int{-3, -4}},
	}
	for i, tt := range tests {
		got := intIntersection(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func intDifferenceTest(name string, t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{-1, -3, -3, -4}, []int{-3, -4, -5}, []int{-1}},
		{[]int{-1, -3, -3, -4}, []int{-12, -12, -5}, []int{-1, -3, -4}},
	}
	for i, tt := range tests {
		got := intDifference(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %v %v, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func appC_C01_ex4Test(name string, t *testing.T) {
	tests := []struct {
		n      int
		expect []int
	}{
		{0, []int{1}},
		{1, []int{1, -1}},
		{2, []int{1, -1, -3 - 2}},
		{3, []int{1, -1, -5, -13}},
		{4, []int{1, -1, -5, -13, -39 + 10}},
	}
	for i, tt := range tests {
		got := appC_C01_ex4(tt.n)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with input %d, expected %v; got %v", name, tt.n, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
func leastCommonIntMultTest(name string, t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{0, 2, 4},
		{540, 504, 7560},
	}
	for i, tt := range tests {
		got := leastCommonIntMult(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("%s fail with inputs %d, %d, expected %v; got %v", name, tt.a, tt.b, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %s %d\n", name, i)
		}
	}
}
