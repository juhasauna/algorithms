package aalto

import (
	"log"
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
	for i := range len(x) + 1 {
		DP[i] = make([]int, len(y)+1)
	}
	// fmt.Printf("%v\n", DP)
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
