package oneoffs

import (
	"reflect"
	"testing"
)

func Test_baseExpansion(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"decToBase", decToBaseTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
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
	}
	for _, tt := range tests {
		if got := decToBase(tt.dec, tt.base); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("decToBase() = %v, want %v", got, tt.want)
		}
	}
}
