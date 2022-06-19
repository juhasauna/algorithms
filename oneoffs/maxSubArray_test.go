package oneoffs

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"maxSubArray_find", maxSubArray_findTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}

}
func maxSubArray_findTest(t *testing.T) {
	verbose := false
	tests := []struct {
		wantLow  int
		wantHigh int
		data     []int
	}{
		{1, 2, []int{10, 3, 8, 1, 3}},
		{7, 11, []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}},
	}

	for _, tt := range tests {
		doer := maxSubArray{data: tt.data, verbose: verbose}
		gotLow, gotHigh, profit := doer.divAndConquer()
		if gotLow != tt.wantLow || gotHigh != tt.wantHigh {
			t.Errorf("FAIL divAndConq: iters: %d, gotLow/want: %d/%d, gotHigh/want: %d/%d", doer.iters, gotLow, tt.wantLow, gotHigh, tt.wantHigh)
		} else {
			fmt.Printf("SUCCESS divAndConq: iters: %d, purchase/sell price: %d/%d, profit %d\n", doer.iters, tt.data[gotLow], tt.data[gotHigh], profit)
		}

		gotLow, gotHigh, profit = doer.brute()
		if gotLow != tt.wantLow || gotHigh != tt.wantHigh {
			t.Errorf("FAIL brute: iters: %d, gotLow/want: %d/%d, gotHigh/want: %d/%d", doer.iters, gotLow, tt.wantLow, gotHigh, tt.wantHigh)
		} else {
			fmt.Printf("SUCCESS brute: iters: %d, purchase/sell price: %d/%d, profit %d\n", doer.iters, tt.data[gotLow], tt.data[gotHigh], profit)
		}
	}
}
