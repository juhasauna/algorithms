package kolman

import "testing"

func Test_nxUsingAddition(t *testing.T) {
	//Added "go.testFlags": ["-v"] to settings.json for verbosity
	tests := []struct {
		name string
		n    int
		x    int
		want int
	}{
		{"a", 1, 1, 1},
		{"b", 1, 10, 10},
		{"c", 5, 2, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nxUsingAddition(tt.n, tt.x); got != tt.want {
				t.Errorf("nxUsingAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
