package aalto

import (
	"math/bits"
)

func nLogSquaredN(n int) int {
	if n <= 1 {
		return 0
	}

	counter := 0

	counter += nLogSquaredN(n / 2)
	counter += nLogSquaredN(n / 2)

	log2 := bits.Len(uint(n)) - 1 // assumes n is a power of 2

	for range n {
		for range log2 {
			counter++
		}
	}

	return counter
}
