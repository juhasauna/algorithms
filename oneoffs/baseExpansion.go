package oneoffs

func decToBase(dec, base int) []int {
	result := []int{}
	rem := dec
	for {
		if r := rem / base; r == 0 {
			result = append([]int{rem}, result...)
			break
		}
		mod := rem % base
		result = append([]int{mod}, result...)
		rem = rem / base
	}
	return result
}

// func intToAlphanum(nums []int) string {
// fmt.Println(byte())
// }
