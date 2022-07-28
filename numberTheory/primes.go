package numberTheory

import (
	"math"
)

func primeFactorization(n int) (result []int) {
	i := 2
	for i <= int(math.Ceil(math.Sqrt(float64(n)))) {
		if trialDivision(i) {
			if n%i == 0 {
				n = n / i
				result = append(result, i)
				i = 2
				continue
			}
		}
		i++
	}
	result = append(result, n)
	return result
}

func greatestCommonDivisor(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
func euclideanAlgorithm(a, b int) int {
	if b > a {
		temp := b
		b = a
		a = temp
	}
	for b != 0 {
		temp := a % b
		a = b
		b = temp
	}
	return a
}

type euclidsAlg struct {
	a int
	b int
}

func (x *euclidsAlg) modularInverse() int {
	_, inverse := x.bezoutCoefficients()
	return inverse
}

// Using the extended Euclidian algorithm
func (x *euclidsAlg) bezoutCoefficients() (s int, t int) {
	steps := x.getSteps()
	s0 := 1
	s1 := 0
	t0 := 0
	t1 := 1
	for i := 0; i < len(steps)-1; i++ {
		quotient := steps[i].quotient()
		s = s0 - quotient*s1
		s0 = s1
		s1 = s
		t = t0 - quotient*t1
		t0 = t1
		t1 = t
	}
	return s, t
}
func (x *euclidsAlg) gcd() int {
	steps := x.getSteps()
	return steps[len(steps)].b
}
func (x *euclidsAlg) sort() {
	if x.b > x.a {
		temp := x.b
		x.b = x.a
		x.a = temp
	}
}

func (x euclidsAlg) quotient() int {
	return x.a / x.b
}
func (x euclidsAlg) remainder() int {
	return x.a - x.b*x.quotient()
}

func (x *euclidsAlg) getSteps() (result []euclidsAlg) {
	x.sort()
	for x.b != 0 {
		step := euclidsAlg{a: x.a, b: x.b}
		temp := x.a % x.b
		x.a = x.b
		x.b = temp
		result = append(result, step)
	}
	return result
}

func leastCommonMult(a, b int) int {
	gcd := greatestCommonDivisor(a, b)
	lcm := a * b / gcd
	return lcm
}
