package aalto

import (
	"log"
	"sort"
)

func coin(S []int, t int) int {
	D := make([]int, t+1)
	D[0] = 0
	for i := 1; i <= t; i++ {
		D[i] = 1 << 10
		for _, s := range S {
			if s <= i {
				candidate := D[i-s] + 1
				if D[i] > candidate {
					D[i] = candidate
				}

			}
		}
	}
	return D[t]
}

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
		// if i < j DP[i][i] == 0
		// This  is important for the special case l = 2 && x[i] == x[j].
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

func longestCommonSubsequence(x, y string) int {

	DP := make([][]int, len(x)+1)
	for i := range DP {
		DP[i] = make([]int, len(y)+1)
	}
	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				DP[i][j] = DP[i-1][j-1] + 1
			} else {
				DP[i][j] = max(DP[i-1][j], DP[i][j-1])
			}
		}
	}

	return DP[len(x)][len(y)]
}

func longestIncreasingSubsequence(x []int) int {
	n := len(x)
	L := make([]int, n)
	answer := 0

	// ONLY 2 nested loops required. O(n^2)
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

func matrixChainMultiplication(p []int) int {
	n := len(p) - 1
	DP := make([][]int, n)
	for i := range DP {
		DP[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			DP[i][j] = 1 << 30
			for k := i; k < j; k++ {
				subprobleSum := DP[i][k] + DP[k+1][j]
				multByCurrentSplit := p[i] * p[k+1] * p[j+1]
				candidate := subprobleSum + multByCurrentSplit
				if DP[i][j] > candidate {
					DP[i][j] = candidate
				}
			}
		}
	}
	return DP[0][n-1]
}

// Bottom-up DP approach
func rodCutting(p []int) int {
	n := len(p)
	DP := make([]int, n+1)
	for l := 1; l <= n; l++ {
		q := -1 << 31
		for i := 1; i <= l; i++ {
			counter++
			q = max(q, p[i-1]+DP[l-i])
		}
		DP[l] = q
	}
	return DP[n]
}

// This is pointless. We're simply implementing a loop with recursion.
func rodCuttingRecursiveTabulation(p []int) int {
	n := len(p)
	DP := make([]int, n+1)

	var recurse func(int)
	recurse = func(l int) {

		if l == n+1 {
			return
		}
		MAX := 0
		for i := 1; i <= l; i++ {
			counter++
			candidate := p[i-1] + DP[l-i]
			if candidate > MAX {
				MAX = candidate
			}
		}
		DP[l] = MAX
		recurse(l + 1)
	}
	recurse(1)
	return DP[n]
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

var counter int

func rodCuttingNaive(p []int) int {
	l := len(p)
	return rodCuttingNaiveAux(p, l)
}

func rodCuttingNaiveAux(p []int, l int) int {
	counter++
	if l == 0 {
		return 0
	}
	q := -2
	for i := range l {
		q = max(q, p[i]+rodCuttingNaiveAux(p, l-i-1))
	}
	return q
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
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].finish != jobs[j].finish {
			return jobs[i].finish < jobs[j].finish
		}
		return jobs[i].start < jobs[j].start
	})
	n := len(jobs)
	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		// p := 0
		// for i := 1; i < j; i++ {
		// 	if jobs[i-1].finish > jobs[j-1].start {
		// 		break
		// 	}
		// 	p = i // p = max{ i<j : f_i < s_j }
		// }
		p := binaryJobPredecessorSearch(jobs, j)
		dp[j] = max(dp[j-1], jobs[j-1].weight+dp[p])
	}
	return dp[n]
}

// Returns p(j): the number of earlier jobs compatible with job j.
// j is 1-based with respect to dp; the current job is jobs[j-1].
func binaryJobPredecessorSearch(jobs []job, j int) int {
	start := jobs[j-1].start

	lo, hi := 0, j-1 // Search jobs[0 : j-1].
	for lo < hi {
		mid := lo + (hi-lo)/2

		if jobs[mid].finish <= start {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	// Exactly lo earlier jobs finish no later than the current job starts.
	// Thus dp[lo] is OPT(p(j)).
	return lo
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
