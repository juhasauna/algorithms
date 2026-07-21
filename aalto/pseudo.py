Longest_palindromic_subsequence(x[1..n]):
    # 1-based indexing
    DP := table of size n by n.
    for i := 1 to n:
        for j := 1 to n:
            DP[i,j] = 0
            if i == j: DP[i,j] = 1
    
    for l := 2 to n:
        for i := 1 to n - l + 1:
            j := i + l - 1
            if x[i] == x[j]:
                DP[i,j] = DP[i + 1,j - 1] + 2
            else:
                DP[i,j] = max{DP[i + 1,j], DP[i,j - 1]}
    
    return DP[1, n]


Coin(S, t):
    # DP values represent how many coins have been selected at step i.
    DP := array of size t 
    DP[0] := 0
    for i in range t:
        DP[i] := infinity
        for each s in S:
            if s <= i:
                cand := DP[i-s] + 1
                if DP[i] > cand:
                    DP[i] = cand

    return DP[t]


Checkerboard(M, n):
    DP := table of size n by n each cell initialized to zero

    for i range n:
        DP[i, 1] = r(i, j)
    
    for i := 2 to n:
        for j range n:
            a, b, c := 0, DP[i-1, j], 0
            if j < n:
                a = DP[i-1, j + 1]
            if j > 1:
                b = DP[i-1, j - 1]
            
            DP[i, j] = max{a, b, c} + DP[i, j]


TournamentProbability(n, i, j):
    if i >= n or j >= n: return "invalid input"

    start_i, start_j, := i, j
    DP := table of size n by n
    for j range n-1:
        DP[n][j] = 1
    
    for i := n-1; i > 0; i--:
        for j := n-1; j > 0; j--:
            DP[i][j] = (DP[i+1][j] + DP[i][j+1]) / 2


    return DP[start_i][start_j]


LongestCommonSubsequence(x, y):
    DP := table of size x+1 by y+1
    for i := 0 to x: DP[i][0] := 0
    for j := 0 to y: DP[0][j] := 0

    for i := 1 to x:
        for j := 1 to y:
            if x[i] == y[j]:
                DP[i,j] = DP[i-1,j-1] + 1
            else:
                DP[i,j] = max{DP[i-1,j], DP[i,j-1]}

    return DP[x][y]

LongestIncreasingSubsequence(x, n):
    L := array of length n, elements initialized to one
    for i := 1 to n:
        j := 1
        while j < i:
            for k = 1 to j:
                MAX := 0
                if x[i] > x[k]:
                    candidate := L[k]
                    if candidate > MAX:
                        MAX = candidate
                    
            if MAX >= L[i]:
                L[i] = MAX + 1
            j++

    MAX := 0
    for i := range len(L):
        candidate := L[i]
        if candidate > MAX:
            MAX = candidate
    return MAX


# The number of valid multiplication orders is the Catalan number.
MatrixChainMultiplication(p): 
    # p := dimensions from 0 to n
    # p_0 = # rows in the first matrix
    # p_1 = # columns in the first matrix
    # p_n = # columns in the last matrix

    DP := table of size n by n

    for l := 2 to n:
        for i := 1 to n - l + 1:
            j := i + l - 1
            DP[i][j] := BIG_INT
            for k := i to j-1:
                subproblem_sum := DP[i][k] + DP[k+1][j]
                # k denotes an index after which where we are placing ")("
                k_split_multiplications := p[i-1] * p[k] * p[j]
                candidate := subproblem_sum + k_split_multiplications
                if candidate < DP[i][j]:
                    DP[i][j] = candidate

    return DP[1][n]


RodCutting(p, l):
    # p gives the prices for 0-l
    # p[0] = 0
    # p[i] > 0, for 1 < i <= l
    # l is the length of the rod
    DP := array [0:n]
    DP[0] = 0

    for j := 1 to n:
        q := -inf
        for i := 1 to j:
            q := max{q, p[i] + DP[j-i]}
        DP[j] = q
    
    return DP[n]


activitySelectionCLRS_15_1():
	# S given activities sorted by finish time.
    # c[i, j] optimal solution on interval i to j
    n := len(S)
    c := table of integers of size n by n

    choice := optimal subset of S

    for i := 0 to n:
        c[i, i+1] = 0
    
    for l := 2 to n+1:
        for i := 0 to n + 1 - l:
            j := i + l
            c[i, j] = 0

            for k := i + 1 to l:
                candidate := c[i, k] + 1 + c[k, j] 
                if candidate > c[i, j]:
                    c[i, j] = candidate
                    choice[i, j] = k
            
