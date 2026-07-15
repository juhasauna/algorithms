package aalto

import "reflect"

// Show that MODEXP is in P. Note that the naive algorithm is exponential time.
func modexp0713(a, b, c, p int64) bool {
	var base int64 = a % p
	var result int64 = 1
	for b > 0 { // O(log n)
		if b%2 == 1 {
			result = (result * base) % p
		}

		base = (base * base) % p
		b /= 2
	}

	return result == c
}

func permPower0714Obvious(p map[int]int, q map[int]int, t int) bool {
	if t < 1 {
		return false
	}
	r := make(map[int]int)
	for key, v := range q {
		r[key] = v
	}
	if t == 1 {
		return reflect.DeepEqual(r, p)
	} else {
		t--
	}

	for range t - 1 {
		for _, v := range q {
			r[v] = v
		}
	}

	return reflect.DeepEqual(r, p)
}
func permPower0714RepeatedSquaring(p, q map[int]int, t int) bool {
	// TODO: implement a version that exploits cycle length LCM - Least common multiple
	if t < 0 {
		return false
	}

	r := make(map[int]int)
	for key := range q {
		r[key] = key
	}

	base := make(map[int]int)
	for key, v := range q {
		base[key] = v
	}

	for t > 0 {
		if t%2 == 1 {
			r = composePerm(r, base)
		}
		base = composePerm(base, base) // Square
		t /= 2
	}

	return reflect.DeepEqual(r, p)
	// Running time: O(k log t)
}

func composePerm(a map[int]int, b map[int]int) map[int]int {
	out := make(map[int]int)
	for key, v := range b {
		out[key] = a[v]
	}
	return out
}
