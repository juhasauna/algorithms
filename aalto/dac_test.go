package aalto

import (
	"math"
	"testing"
)

func Test_dac(t *testing.T) {

	nLogSquaredNTest(t)

}

func nLogSquaredNTest(t *testing.T) {
	tests := []struct {
		name string
		len  int
	}{
		{"", 8},
		{"", 16},
		// {"", 32},
		// {"", 64},
		{"", 128},
	}
	for _, tt := range tests {
		got := nLogSquaredN(tt.len)
		logn := int(math.Log2(float64(tt.len)))
		want := (tt.len / 2) * logn * (logn + 1) // O(n log^2 n)
		if got != want {
			t.Errorf("got/want (%d/%d)", got, want)
		} else {
			t.Log(got)
		}
	}
}
