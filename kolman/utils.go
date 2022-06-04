package kolman

type utils struct {
}

func (x utils) rmElementInt(i int, s []int) []int {
	s[i] = s[len(s)-1] // Replace the element that we want to remove with the last element.
	s = s[:len(s)-1]   // Reduce the size of the slice by one from the end.
	return s
}
func (x utils) containsInt(a int, b []int) bool {
	for _, vb := range b {
		if a == vb {
			return true
		}
	}
	return false
}

func (x utils) distinctInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
