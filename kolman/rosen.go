package kolman

func nxUsingAddition(n, x int) int {
	if n == 0 {
		return 0
	}
	return x + nxUsingAddition(n-1, x)
}
