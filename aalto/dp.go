package aalto

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

var counter int

func checkerboard(board [][]int) int {
	n := len(board)
	if n != len(board[0]) {
		log.Fatal("invalid input\n")
	}
	DP := make([][]int, n)
	for i := range n {
		DP[i] = make([]int, n)
		DP[0][i] = board[0][i]
	}

	for i := 1; i < n; i++ {
		for j := range n {
			prevRow := DP[i-1]
			a, b, c := 0, prevRow[j], 0
			if j > 0 {
				a = prevRow[j-1]
			}
			if j < n-1 {
				c = prevRow[j+1]
			}
			candidate := max(a, b, c) + board[i][j]
			if candidate > DP[i][j] {
				DP[i][j] = candidate
			}
		}
	}

	MAX := 0
	for j := range n {
		if DP[n-1][j] > MAX {
			MAX = DP[n-1][j]
		}
	}
	return MAX
}

func longestPalindromicSubsequence(x string) int {
	n := len(x)
	DP := make([][]int, n)
	for i := range DP {
		DP[i] = make([]int, n)
		DP[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			if x[i] == x[j] {
				DP[i][j] = DP[i+1][j-1] + 2
			} else {
				DP[i][j] = max(DP[i+1][j], DP[i][j-1])
			}
		}
	}
	return DP[0][n-1]
}
func lps(x string) int {
	n := len(x)
	if n < 2 {
		return n
	}
	if x[0] == x[n-1] {
		return lps(x[1:n-1]) + 2
	}
	return max(lps(x[1:]), lps(x[:n-1]))
}

func tournamentProbability(n, start_i, start_j int) float32 {
	if start_i >= n || start_j >= n || n <= 0 {
		log.Fatal("invalid arguments", n, start_i, start_j)
	}
	DP := make([][]float32, n+1)
	for i := range DP {
		DP[i] = make([]float32, n+1)
	}

	for i := 0; i <= n; i++ {
		DP[i][n] = 0.0
		DP[n][i] = 1.0
	}
	DP[n][n] = -1.0 // Impossible state

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			DP[i][j] = (DP[i+1][j] + DP[i][j+1]) / 2
		}
	}

	return DP[start_i][start_j]
}

// LongestCommonSubsequence
func lcs(x, y string) int {
	n, m := len(x), len(y)
	if n == 0 || m == 0 {
		return 0
	}
	if x[0] == y[0] {
		return lcs(x[1:], y[1:]) + 1
	}
	return max(lcs(x[1:], y), lcs(x, y[1:]))
}

func lcsMemo(x, y string) int {
	n, m := len(x), len(y)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var recurse func(i, j int) int
	recurse = func(i, j int) int {
		if i == n || j == m {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if x[i] == y[j] {
			memo[i][j] = 1 + recurse(i+1, j+1)
		} else {
			memo[i][j] = max(recurse(i+1, j), recurse(i, j+1))
		}
		return memo[i][j]
	}
	return recurse(0, 0)
}

func lcsBottomUp(x, y string) int {
	n, m := len(x), len(y)
	DP := make([][]int, n+1)
	for i := range DP {
		DP[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if x[i-1] == y[j-1] {
				DP[i][j] = DP[i-1][j-1] + 1
			} else {
				DP[i][j] = max(DP[i-1][j], DP[i][j-1])
			}
		}
	}
	return DP[n][m]
}

func longestIncreasingSubsequence(x []int) int {
	n := len(x)
	L := make([]int, n)
	answer := 0

	for i := range n {
		L[i] = 1
		best := 0
		for j := range i {
			if x[i] > x[j] {
				best = max(L[j], best)
			}
		}
		L[i] = best + 1
		answer = max(answer, L[i])
	}
	return answer
}
func mcm(p []int) int {
	n := len(p)
	if n < 3 {
		return 0
	}

	best := 1<<32 - 1
	for k := 1; k < n-1; k++ {
		counter++
		left, right := p[:k+1], p[k:n]
		m := p[0] * p[k] * p[n-1]
		best = min(best, mcm(left)+m+mcm(right))
	}
	return best
}

func mcmMemo(p []int) int {
	n := len(p)
	memo := make([][]int, n)
	for i := range n {
		memo[i] = make([]int, n)
	}
	var recurse func(i, j int) int
	recurse = func(i, j int) int {
		if i >= j {
			return 0
		}
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		best := 1<<32 - 1
		for k := i + 1; k <= j; k++ {
			counter++
			m := p[i] * p[k] * p[j+1]
			best = min(best, recurse(i, k-1)+m+recurse(k, j))
		}
		memo[i][j] = best
		return best
	}
	return recurse(0, n-2)
}

func matrixChainMultiplication(p []int) int {
	n := len(p) - 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			dp[i][j] = 1 << 30
			for k := i; k < j; k++ {
				counter++
				subprobleSum := dp[i][k] + dp[k+1][j]
				multByCurrentSplit := p[i] * p[k+1] * p[j+1]
				candidate := subprobleSum + multByCurrentSplit
				if dp[i][j] > candidate {
					dp[i][j] = candidate
				}
			}
		}
	}
	return dp[0][n-1]
}

// Bottom-up DP approach
func rodCutting(p []int) int {
	n := len(p)
	dp := make([]int, n+1)
	for l := 1; l <= n; l++ {
		best := 0
		for i := 1; i <= l; i++ {
			counter++
			best = max(best, p[i-1]+dp[l-i])
		}
		dp[l] = best
	}
	return dp[n]
}

// Memoization solves subproblems top-down and caches their results, whereas tabulation computes subproblems bottom-up.
func rodCuttingMemoized(p []int) int {
	n := len(p)
	DP := make([]int, n+1)
	for i := 1; i <= n; i++ {
		DP[i] = -1
	}
	var recurse func(int) int
	recurse = func(l int) int {
		counter++
		if DP[l] != -1 {
			return DP[l]
		}
		maxRevenue := -2
		for i := 1; i <= l; i++ {

			maxRevenue = max(maxRevenue, p[i-1]+recurse(l-i))
		}
		DP[l] = maxRevenue
		// fmt.Printf("%v\n", DP)
		return maxRevenue
	}
	return recurse(n)
}

func rodCuttingNaive(p []int) int {
	counter++
	l := len(p)
	if l == 0 {
		return 0
	}
	best := p[l-1]
	for i := 1; i < l; i++ {
		candidate := p[i-1] + rodCuttingNaive(p[:l-i])
		best = max(best, candidate)
	}
	return best
}

// CLRS: 15.1 An activity-selection problem
type job struct {
	start  int
	finish int
	weight int // Weighted Interval Scheduling
}

func activitySelectionRecursive(jobs []job) int {
	// Weights are ignored
	// We are only interested in how many activities we can accomplish.
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].finish != jobs[j].finish {
			return jobs[i].finish < jobs[j].finish
		}
		return jobs[i].start < jobs[j].start
	})
	n := len(jobs)

	var recurse func(k int) int
	recurse = func(k int) int {
		m := k + 1
		for m < n && jobs[m].start < jobs[k].finish {
			m++
		}
		if m == n {
			return 1
		}
		return 1 + recurse(m)
	}
	return recurse(0)
}
func activitySelectionIterative(jobs []job) int {
	// Weights are ignored
	// We are only interested in how many activities we can accomplish.
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].finish != jobs[j].finish {
			return jobs[i].finish < jobs[j].finish
		}
		return jobs[i].start < jobs[j].start
	})
	n := len(jobs)

	selection := []job{jobs[0]}

	k := 0
	for m := 2; m < n; m++ {
		if jobs[m].start >= jobs[k].finish {
			selection = append(selection, jobs[m])
			k = m
		}
	}
	return len(selection)

}

// Weighted Interval Scheduling
func weightedIntervalScheduling(jobs []job) int {
	n := len(jobs)
	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		p := binaryJobPredecessorSearch(jobs, j)
		dp[j] = max(dp[j-1], jobs[j-1].weight+dp[p])
	}
	return dp[n]
}

func binaryJobPredecessorSearch(jobs []job, j int) int {
	lo, hi, start := 0, j-1, jobs[j-1].start
	for lo < hi {
		counter++
		mid := lo + (hi-lo)/2
		if jobs[mid].finish <= start {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
func lastJobTofinishBefore(jobs []job, start int) int {
	for i := len(jobs) - 1; i >= 0; i-- {
		counter++
		if jobs[i].finish <= start {
			return i
		}
	}
	return -1
}
func wis(jobs []job) int {
	counter++
	n := len(jobs)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return jobs[0].weight
	}

	j := jobs[n-1]
	i := n - 2
	for i >= 0 && jobs[i].finish > j.start {
		i--
	}

	include := wis(jobs[:i+1]) + j.weight
	exclude := wis(jobs[:n-1])

	return max(include, exclude)
}
func wisProfBryce(jobs []job) int {
	// sort by increasing finish time
	n := len(jobs)
	opt := make([]int, n)

	opt[0] = jobs[0].weight
	for i := 1; i < n; i++ {
		j := lastJobTofinishBefore(jobs, jobs[i].start)
		v := 0
		if j != -1 {
			v = opt[j]
		}
		if candidate := v + jobs[i].weight; candidate >= opt[i-1] {
			opt[i] = candidate
		} else {
			opt[i] = opt[i-1]
		}
	}
	return opt[n-1]
}

func weightedIntervalSchedulingAbomination(jobs []job) int {
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].finish != jobs[j].finish {
			return jobs[i].finish < jobs[j].finish
		}
		return jobs[i].start < jobs[j].start
	})
	n := len(jobs)
	all := make([]job, n+2)
	all[0] = job{
		start:  -1 << 30,
		finish: -1 << 30,
		weight: 0,
	}
	copy(all[1:n+1], jobs)
	all[n+1] = job{
		start:  1 << 30,
		finish: 1 << 30,
		weight: 0,
	}
	DP := make([][]int, n+2)
	for i := range DP {
		DP[i] = make([]int, n+2)
	}
	for length := 2; length <= n+1; length++ {
		for i := 0; i+length <= n+1; i++ {
			j := i + length
			best := 0
			for k := i + 1; k < j; k++ {
				if all[i].finish <= all[k].start && all[k].finish <= all[j].start {
					candidate := DP[i][k] + all[k].weight + DP[k][j]
					if candidate > best {
						best = candidate
					}
				}
			}
			DP[i][j] = best
		}
	}

	return DP[0][n+1]
}

func activitySelectionDP(jobs []job) ([]job, int) {
	n_ := len(jobs)
	n := n_ + 2
	DP := make([][]int, n)
	choice := make([][]int, n)

	for i := range n {
		DP[i] = make([]int, n)
		choice[i] = make([]int, n)

		for j := range n {
			choice[i][j] = -1
		}
	}

	for l := 2; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			for k := i + 1; k < j; k++ {
				current := jobs[k-1]
				compatibleWithLeft := i == 0 || jobs[i-1].finish <= current.start
				compatibleWithRight := j == n-1 || current.finish <= jobs[j-1].start
				if compatibleWithLeft && compatibleWithRight {
					candidate := DP[i][k] + 1 + DP[k][j]
					if candidate > DP[i][j] {
						DP[i][j] = candidate
						choice[i][j] = k
					}
				}
			}
		}
	}
	selection := make([]job, 0, DP[0][n-1])

	var reconstruct func(int, int)
	reconstruct = func(i, j int) {
		k := choice[i][j]
		if k == -1 {
			return
		}

		reconstruct(i, k)
		selection = append(selection, jobs[k-1])
		reconstruct(k, j)
	}

	reconstruct(0, n-1)
	return selection, DP[0][n-1]
}

func scrambledStrings(a, b string) bool {
	n := len(a)
	if n != len(b) {
		log.Fatal("input lengths must match")
	}
	dp := make([][][]bool, n)
	for i := range n {
		dp[i] = make([][]bool, n)
		for j := range n {
			dp[i][j] = make([]bool, n+1)
			dp[i][j][1] = a[i] == b[j]
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			for j := 0; j+l <= n; j++ {
				for s := 1; s < l; s++ {
					counter++
					sameOrder := dp[i][j][s] && dp[i+s][j+s][l-s]
					flipOrder := dp[i][j+l-s][s] && dp[i+s][j][l-s]
					dp[i][j][l] = dp[i][j][l] || sameOrder || flipOrder
				}
			}
		}
	}

	return dp[0][0][n]
}

func scrambledStringsMemoization(a, b string) bool {
	if len(a) != len(b) {
		log.Fatal("input lengths must match")
	}
	type state struct {
		i, j, l int
	}
	memo := make(map[state]bool)
	var scramble func(i, j, l int) bool
	scramble = func(i, j, l int) bool {
		current := state{i, j, l}
		if result, computed := memo[current]; computed {
			return result
		}
		if l == 1 {
			result := a[i] == b[j]
			memo[current] = result
			return result
		}
		for s := 1; s < l; s++ {
			counter++
			sameOrder := scramble(i, j, s) && scramble(i+s, j+s, l-s)
			if sameOrder {
				memo[current] = true
				return true
			}
			flippedOrder := scramble(i, j+l-s, s) && scramble(i+s, j, l-s)
			if flippedOrder {
				memo[current] = true
				return true
			}
		}
		memo[current] = false
		return false
	}
	return scramble(0, 0, len(a))
}

func fibonacciBottomUpTabulation(k int) int {
	DP := make([]int, k+1)
	DP[0] = 0
	DP[1] = 1
	for i := 2; i <= k; i++ {
		counter++
		DP[i] = DP[i-1] + DP[i-2]
	}
	return DP[k]
}
func fibonacciBottomUp(k int) int {
	x, y := 0, 1
	for counter < k-1 {
		counter++
		x, y = y, x+y
	}
	return y
}

func fibonacci(k int) int {
	counter++
	if k < 2 {
		return k
	}
	return fibonacci(k-1) + fibonacci(k-2)
}

func fibonacciMemoization(k int) int { // Memoization (top-down) typically done with recursion.
	memo := make([]int, k+1)
	memo[0] = 0
	memo[1] = 1

	var recurse func(l int) int
	recurse = func(l int) int {
		if l == 0 {
			return 0
		}
		if memo[l] != 0 {
			return memo[l]
		}
		counter++
		memo[l] = recurse(l-1) + recurse(l-2)
		return memo[l]
	}

	return recurse(k)
}

// 14-9 Breaking a string
func cuttingOrderBottomUp(l []int) int {
	l = append([]int{0}, l...)
	n := len(l)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	for w := 2; w < n; w++ {
		for i := 0; i+w < n; i++ {
			j := i + w
			best := 1 << 10
			for k := i + 1; k < j; k++ {
				counter++
				candidate := dp[i][k] + dp[k][j]
				best = min(best, candidate)
			}
			dp[i][j] = l[j] - l[i] + best
		}
	}
	// fmt.Println(dp)
	return dp[0][n-1]
}

func cuttingOrderSUCK(l []int) int {
	counter++
	n := len(l)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 0
	}

	m := l[n-1]
	best := 1 << 10
	for _, v := range l[:n-1] {
		var left, right []int
		for _, w := range l[:n-1] {
			if w < v {
				left = append(left, w)
			} else if w > v {
				right = append(right, w-v)
			}
		}

		left = append(left, v)
		right = append(right, m-v)
		candidate := m + cuttingOrderSUCK(left) + cuttingOrderSUCK(right)
		best = min(best, candidate)
	}
	return best
}

func cuttingOrder(lengths []int) int {
	lengths = append([]int{0}, lengths...)
	n := len(lengths)

	var recurse func(i, j int) int
	recurse = func(i, j int) int {

		l := j - i
		if l <= 1 {
			return 0
		}
		m := lengths[j] - lengths[i]
		if l == 2 {
			return m
		}
		best := 1 << 10
		for k := i + 1; k < j; k++ {
			counter++
			left, right := recurse(i, k), recurse(k, j)
			best = min(best, min(left+right))
		}
		return m + best
	}
	return recurse(0, n-1)
}
func cuttingOrderMemo(lengths []int) int {
	lengths = append([]int{0}, lengths...)
	n := len(lengths)
	memo := make([][]int, n)
	for i := range n {
		memo[i] = make([]int, n)
	}
	var recurse func(i, j int) int
	recurse = func(i, j int) int {
		l := j - i
		if l <= 1 {
			return 0
		}
		m := lengths[j] - lengths[i]
		if l == 2 {
			return m
		}
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		best := 1 << 10
		for k := i + 1; k < j; k++ {
			counter++
			left, right := recurse(i, k), recurse(k, j)
			best = min(best, min(left+right))
		}
		best += m
		memo[i][j] = best
		return best
	}
	return recurse(0, n-1)
}

func knapsack(weights, values []int, cap int) int {
	n := len(weights)
	if n == 0 {
		return 0
	}
	i := n - 1
	best := knapsack(weights[:i], values[:i], cap)

	if weights[i] <= cap {
		candidate := values[i] + knapsack(weights[:i], values[:i], cap-weights[i])
		best = max(best, candidate)
	}
	return best
}

func knapsackDP(weights []int, values []int, cap int) int {
	n := len(weights)

	dp := make([][]int, n+1)
	for i := range len(dp) {
		dp[i] = make([]int, cap+1)
	}

	for i := 1; i <= n; i++ {
		for c := 1; c <= cap; c++ {
			w, v := weights[i-1], values[i-1]
			if c < w {
				dp[i][c] = dp[i-1][c]
			} else {
				candidate := v + dp[i-1][c-w]
				dp[i][c] = max(dp[i-1][c], candidate)
			}
		}
	}
	return dp[n][cap]
}
func ks(weights []int, values []int, cap int) int {
	n := len(weights)
	if n == 0 {
		return 0
	}
	best := 0
	for k := range n {
		leftW := weights[:k]
		leftV := values[:k]

		rightW := []int{}
		rightV := []int{}
		if k < n-1 {
			rightW = weights[k+1:]
			rightV = values[k+1:]
		}
		newC := cap - weights[k]
		if newC < 0 {
			continue
		}
		candidate := values[k] + max(ks(leftW, leftV, newC), ks(rightW, rightV, newC))
		best = max(best, candidate)
	}
	return best
}

func coinMemo(s []int, t int) int {
	const inf = 1 << 20
	memo := make([]int, t+1)
	for i := range memo {
		memo[i] = inf
	}
	memo[0] = 0
	var recurse func(target int) int
	recurse = func(target int) int {
		if target < 0 {
			return inf
		}
		if memo[target] != inf {
			return memo[target]
		}
		best := inf
		for _, coin := range s {
			counter++
			if coin <= target {
				best = min(best, recurse(target-coin)+1)
			}
		}
		memo[target] = best
		return best
	}
	return recurse(t)
}

func coinRec(s []int, t int) int {
	counter++
	best := 1 << 20
	for _, coin := range s {
		next := t - coin
		if next == 0 {
			return 1
		} else if next > 0 {
			candidate := coinRec(s, next) + 1
			best = min(best, candidate)
		}
	}
	return best
}

func coinDP(s []int, t int) int {
	dp := make([]int, t+1)
	const inf = 1 << 10
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= t; i++ {
		dp[i] = inf
		for _, coin := range s {
			counter++
			if x := i - coin; x >= 0 {
				dp[i] = min(dp[i], dp[x]+1)
			}
		}
	}
	return dp[t]
}

func longestPalindrome2(x string) string {
	n := len(x)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		dp[i][i] = 1
		counter++
		if i < n-1 {
			if x[i] == x[i+1] {
				dp[i][i+1] = 2
			}
		}
	}

	for l := 2; l < n; l++ {
		for i := range n - l {
			counter++
			j := i + l
			isPalindrome := x[i] == x[j] && // fringes match
				l-1 == dp[i+1][j-1] // inside is palindrome
			if isPalindrome {
				dp[i][j] = dp[i+1][j-1] + 2
			}
		}
	}
	best, q, r := 0, 0, 0
	for i := range dp {
		for j := i; j < n; j++ {
			counter++
			temp := best
			best = max(best, dp[i][j])
			if temp != best {
				q, r = i, j
			}
		}
	}
	return x[q : r+1]
}

func longestPalindrome(x string) string {
	n := len(x)
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		counter++
		if i < n-1 {
			dp[i][i+1] = x[i] == x[i+1]
		}
	}
	for l := 3; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			counter++
			j := i + l - 1
			isPalindrome := dp[i+1][j-1] && x[i] == x[j]
			dp[i][j] = isPalindrome
		}
	}
	best, q, r := 0, 0, 0
	for i, v := range dp {
		for j := i; j < n; j++ {
			counter++
			if v[j] {
				best = max(best, j-i+1)
				if best == j-i+1 {
					q, r = i, j

				}
			}
		}
	}
	return x[q : r+1]
}

func lonPalNaive(y string) string {
	counter++
	if len(y) == 0 {
		return ""
	} else if len(y) == 1 || (len(y) == 2 && y[0] == y[1]) {
		return y
	} else if len(y) == 2 {
		return y[0:1]
	}
	j := len(y) - 1
	insideIsPalindrome := lonPalNaive(y[1:j]) == y[1:j]
	if y[0] == y[j] && insideIsPalindrome {
		return y
	}
	a, b := lonPalNaive(y[0:j]), lonPalNaive(y[1:j+1])
	if len(a) < len(b) {
		return b
	}
	return a
}
func longestPalindromeTopDown(x string) string {
	n := len(x)
	memo := make(map[string]string)
	memo[""] = ""
	for i := range n {
		counter++
		trivialCase := string(x[i])
		memo[trivialCase] = trivialCase
		if i < n-1 {
			if x[i] == x[i+1] {
				memo[x[i:i+2]] = x[i : i+2]
			} else {
				memo[x[i:i+2]] = trivialCase
			}
		}
	}

	var recurse func(y string) string
	recurse = func(y string) string {
		j := len(y) - 1
		counter++
		if s, ok := memo[y]; ok {
			return s
		}

		if y[0] == y[j] {
			inside := y[1:j]
			if s, ok := memo[inside]; ok {
				if s == inside {
					memo[y] = y
					return y
				}
			} else {
				insidePal := recurse(inside)
				memo[inside] = insidePal
				if insidePal == inside {
					memo[y] = y
					return y
				}
			}
		}
		q, r := y[0:j], y[1:j+1]
		i := ""
		if s, ok := memo[q]; ok {
			i = s
		} else {
			i = recurse(q)
			memo[q] = i
		}
		k := ""
		if s, ok := memo[r]; ok {
			k = s
		} else {
			k = recurse(r)
			memo[r] = k
		}
		if len(i) >= len(k) {
			return i
		}
		return k
	}
	return recurse(x)
}

func optimalBST(f []int) int {
	n := len(f)
	dp := make([][]int, n)
	for i := range dp {
		counter++
		dp[i] = make([]int, n)
		dp[i][i] = f[i]
	}
	const inf = 1 << 10
	for i := range n {
		for j := i + 1; j < n; j++ {
			counter++
			dp[i][j] = inf
		}
	}
	for l := 2; l <= n; l++ {
		for i := range n - l + 1 {
			j := i + l - 1
			freqs := 0
			for k := i; k <= j; k++ {
				counter++
				candidate := 0
				// When both conditionals below are true, k has 2 child nodes,
				// which contain subtrees: left: [i..k) and right: (k..j]
				if k > i {
					candidate += dp[i][k-1]
				}
				if k < j {
					candidate += dp[k+1][j]
				}
				dp[i][j] = min(dp[i][j], candidate)
				freqs += f[k]
			}
			dp[i][j] += freqs
		}
	}
	fmt.Println(dp)
	return dp[0][n-1]
}

// Todo: Draw the tree from the DP-table (wiht pen & paper)
func oBST(f []int) int {
	counter++
	n := len(f)
	if n == 0 {
		return 0
	} else if n == 1 {
		return f[0]
	}
	freqs := 0
	for _, v := range f {
		counter++
		freqs += v
	}
	best := 1 << 10
	for k := range n {
		l, r := 0, 0
		if k > 0 {
			l = oBST(f[0:k])
		}
		if k < n-1 {
			r = oBST(f[k+1 : n])
		}
		best = min(best, l+r)
	}
	return best + freqs
}

func oBSTTopDownMemoized(frequencies []int) int {
	n := len(frequencies)
	memo := make(map[string]int, n)
	for _, v := range frequencies {
		f := []int{v}
		memo[makeOBSTMemoKey(f)] = v
	}
	memo[""] = 0

	var recurse func(f []int) int
	recurse = func(f []int) int {

		key := makeOBSTMemoKey(f)
		if v, ok := memo[key]; ok {
			return v
		}

		freqSum := 0
		for _, v := range f {
			freqSum += v
		}
		m := len(f)
		best := 1 << 10
		for k := range m {
			l, r := 0, 0
			if k > 0 {
				l = recurse(f[0:k])
			}
			if k < m-1 {
				r = recurse(f[k+1 : m])
			}
			best = min(best, l+r)
		}
		best += freqSum
		memo[key] = best
		return best
	}

	return recurse(frequencies)
}

func makeOBSTMemoKey(f []int) string {
	s := ""
	for _, v := range f {
		counter++
		intStr := strconv.Itoa(v)
		s = s + intStr
	}
	return s
}

func optimalBSTDrawIt(f []int) int {
	n := len(f)
	dp := make([][]int, n)
	roots := make([][]int, n)
	for i := range dp {
		counter++
		dp[i] = make([]int, n)
		dp[i][i] = f[i]

		roots[i] = make([]int, n)
		roots[i][i] = i
	}
	const inf = 1 << 10
	for i := range n {
		for j := i + 1; j < n; j++ {
			counter++
			dp[i][j] = inf
		}
	}
	for l := 2; l <= n; l++ {
		for i := range n - l + 1 {
			j := i + l - 1
			freqs := 0
			for k := i; k <= j; k++ {
				counter++
				candidate := 0
				// When both conditionals below are true, k has 2 child nodes,
				// which contain subtrees: left: [i..k) and right: (k..j]
				if k > i {
					candidate += dp[i][k-1]
				}
				if k < j {
					candidate += dp[k+1][j]
				}
				dp[i][j] = min(dp[i][j], candidate)
				if dp[i][j] == candidate {
					roots[i][j] = k
				}
				freqs += f[k]
			}
			dp[i][j] += freqs
		}
	}
	// fmt.Println(roots)
	printOptimalBST(roots, f)
	return dp[0][n-1]
}

func printOptimalBST(root [][]int, frequencies []int) {
	n := len(frequencies)
	if n == 0 {
		fmt.Println("(empty)")
		return
	}

	var draw func(i, j int, prefix string, isLeft bool)

	draw = func(i, j int, prefix string, isLeft bool) {
		if i > j {
			return
		}

		r := root[i][j]

		// Print the right subtree above the current node.
		if r < j {
			nextPrefix := prefix
			if isLeft {
				nextPrefix += "│   "
			} else {
				nextPrefix += "    "
			}
			draw(r+1, j, nextPrefix, false)
		}

		connector := "┌── "
		if isLeft {
			connector = "└── "
		}

		fmt.Printf("%s%sk%d: %d\n",
			prefix, connector, r, frequencies[r])

		// Print the left subtree below the current node.
		if r > i {
			nextPrefix := prefix
			if isLeft {
				nextPrefix += "    "
			} else {
				nextPrefix += "│   "
			}
			draw(i, r-1, nextPrefix, true)
		}
	}

	r := root[0][n-1]

	// Draw the root separately so it has no connector.
	if r < n-1 {
		draw(r+1, n-1, "", false)
	}

	fmt.Printf("k%d: %d\n", r, frequencies[r])

	if r > 0 {
		draw(0, r-1, "", true)
	}
	fmt.Println()
}

func lcsDP(x, y string) int {
	n, m := len(x), len(y)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if x[i-1] == y[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[n][m]
}
