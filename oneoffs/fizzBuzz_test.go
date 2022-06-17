package oneoffs

import (
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		do   bool
		f    func(*testing.T)
	}{{"naiveTest", true, naiveTest}}

	for _, tt := range tests {
		if tt.do {
			t.Run(tt.name, func(t *testing.T) {
				tt.f(t)
			})
		}
	}
}

func naiveTest(t *testing.T) {
	// t.Errorf("FAIL")
	// return
	naive()
}
