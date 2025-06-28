package ut

import "testing"

func Test_generic(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"Min", minTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func minTest(t *testing.T) {

	testsInt := []struct {
		a    int
		b    []int
		want int
	}{
		{2, []int{3, 4, 1}, 1},
	}

	for _, tt := range testsInt {
		got := Min(tt.a, tt.b...)
		if got != tt.want {
			t.Errorf("FAIL GOT: %v, want: %v", got, tt.want)
		} else {
			t.Logf("GOT: %v", got)
		}
	}
	testsFloat := []struct {
		a    float64
		b    []float64
		want float64
	}{
		{2, []float64{3, 4, 1, 0.123}, 0.123},
		{.1, []float64{3, 4, 1, 0.123}, 0.1},
	}

	for _, tt := range testsFloat {
		got := Min(tt.a, tt.b...)
		if got != tt.want {
			t.Errorf("FAIL GOT: %v, want: %v", got, tt.want)
		} else {
			t.Logf("GOT: %v", got)
		}
	}
}
