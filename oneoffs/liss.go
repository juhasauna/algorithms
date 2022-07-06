// Longest increasing sub-sequence
package oneoffs

func liss(x []int) (result []int) {
	lst := listOfIss(x)
	for i := len(lst) - 1; i >= 0; i-- {
		current := lst[i]
		tail := []iss{}
		for _, v := range lst[i+1:] {
			if v.startValue > current.startValue {
				tail = append(tail, v)
			}
		}
		for _, v := range tail {
			if len(v.lst) >= len(current.lst) {
				updateCurrenLst := append([]int{current.startValue}, v.lst...)
				current = iss{startValue: current.startValue, lst: updateCurrenLst}
				lst[i] = current
			}
		}
	}
	for _, v := range lst {
		if len(v.lst) > len(result) {
			result = v.lst
		}
	}
	return result
}

type iss struct {
	lst        []int
	startValue int
}

func listOfIss(data []int) (result []iss) {
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, iss{lst: firstIss(data[i:]), startValue: data[i]})
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func firstIss(x []int) (result []int) {
	result = append(result, x[0])
	for _, v := range x {
		if v > result[len(result)-1] {
			result = append(result, v)
		}
	}
	return result
}
