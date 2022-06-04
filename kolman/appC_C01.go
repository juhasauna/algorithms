// p. 460
package kolman

var ut utils

func intUnion(a, b []int) []int { // Ex 1
	result := a
	for _, v := range b {
		if !ut.containsInt(v, result) {
			result = append(result, v)
		}
	}
	return ut.distinctInt(result)
}

func intIntersection(a, b []int) []int { // Ex 2
	result := []int{}
	for _, v := range a {
		if ut.containsInt(v, b) {
			result = append(result, v)
		}
	}
	return ut.distinctInt(result)
}

func intDifference(a, b []int) []int { // Ex 3.
	// A - B (or using a different notation) A \ B
	result := a
	for i, v := range result {
		if ut.containsInt(v, b) {
			result = ut.rmElementInt(i, result)
			result = intDifference(result, b)
			break
		}
	}
	return ut.distinctInt(result)
}

func appC_C01_ex4_g(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return -1
	}
	v1 := 3 * appC_C01_ex4_g(n-1)
	v2 := 2 * appC_C01_ex4_g(n-2)
	return v1 - v2
}
func appC_C01_ex4(elements int) []int {
	result := []int{}
	if elements < 0 {
		return result
	}
	for i := 0; i <= elements; i++ {
		result = append(result, appC_C01_ex4_g(i))
	}
	return result
}

func leastCommonIntMult(a, b int) int {
	return 0
}
