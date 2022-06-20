package oneoffs

import (
	"fmt"
	"math"
)

type maxSubArray struct {
	data    []int
	change  []int
	high    int
	verbose bool
	iters   int
}

func (x *maxSubArray) brute() (low, high, profit int) {
	x.high = len(x.data) - 1
	low = 0
	high = x.high
	for i := range x.data {
		tempLow := i
		for j := i; j <= x.high; j++ {
			tempHigh := j
			tryProfit := x.data[tempHigh] - x.data[tempLow]
			x.iters++
			if tryProfit > profit {
				profit = tryProfit
				low = tempLow
				high = tempHigh
			}
		}
	}
	return low, high, profit
}

func (x *maxSubArray) divAndConquer() (int, int, int) {
	x.high = len(x.data) - 1
	x.change = x.getChange()
	x.log("change: %v", x.change)
	low, high, sum := x.divAndConquer_inner(0, x.high, "start")
	x.log("final: low: %d,high: %d,sum: %d", low, high, sum)
	return low, high, sum
}
func (x *maxSubArray) log(format string, msg ...interface{}) {
	if x.verbose {
		fmt.Printf(format, msg...)
	}
}
func (x *maxSubArray) divAndConquer_inner(low, high int, msg ...string) (int, int, int) {
	x.iters++
	x.log("%s: low: %d, high: %d\n", msg, low, high)
	mid := (low + high) / 2
	if low == high || mid == high {
		return low, high, x.change[low]
	}
	leftLow, leftHigh, leftSum := x.divAndConquer_inner(low, mid, "left")
	rightLow, rightHigh, rightSum := x.divAndConquer_inner(mid+1, high, "right")
	crossLow, crossHigh, crossSum := x.maxCrossing(low, mid, high)

	x.log("left: %d,%d,%d, cross: %d,%d,%d, right: %d,%d,%d\n", leftLow, leftHigh, leftSum, rightLow, rightHigh, rightSum, crossLow, crossHigh, crossSum)

	if leftSum > rightSum && leftSum > crossSum {
		// return low, mid, leftSum
		return leftLow, leftHigh, leftSum
	}
	if rightSum > crossSum {
		return rightLow, rightHigh, rightSum
		// return mid + 1, high, rightSum
	}
	return crossLow, crossHigh, crossSum
}

func (x *maxSubArray) maxCrossing(low, mid, high int) (int, int, int) {
	x.iters++
	if low == high {
		return low, high, 0
	}
	sum := 0
	lowResult := low

	leftSum := 0
	for i := mid; i > low; i-- {
		x.iters++
		sum += x.change[i]
		if sum > leftSum {
			leftSum = sum
			lowResult = i - 1
		}
	}

	rightSum := 0
	sum = 0
	highResult := mid
	for i := mid + 1; i < high; i++ {
		x.iters++
		sum += x.change[i]
		if sum > rightSum {
			rightSum = sum
			highResult = i
		}
	}
	return lowResult, highResult, (leftSum + rightSum)
}

func (x *maxSubArray) getChange() (result []int) {
	prev := x.data[0]
	for _, v := range x.data {
		chg := v - prev
		result = append(result, chg)
		prev = v
	}
	return result
}

func (x *maxSubArray) linearTime() (low, high, profit int) {
	low = math.MaxInt
	chg := x.getChange()
	for i := len(x.data) - 1; i > 0; i-- {
		if chg[i] < low {
			low = chg[i]
		}

	}
}
