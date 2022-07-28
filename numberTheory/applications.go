package numberTheory

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Rosen: p306, ex.25. Write out in pseudocode an algorithm for solving a simultaneous system of linear congruences based on the
// construction in the proof of the Chinese remainder theorem.
func solveCongruences(mods, rems []int) (result int) {
	m := mods[0]
	for i := 1; i < len(mods); i++ {
		m *= mods[i]
	}
	for i, remainder := range rems {
		M := m / mods[i]
		euclidsTools := euclidsAlg{mods[i], M}
		inverse := euclidsTools.modularInverse()
		println(M, inverse)
		result += remainder * M * inverse
	}

	return result
}

type rng struct {
	mod        int
	mult       int
	increment  int
	seed       int
	iterations int
}

func (x *rng) linearCongruential() (result int) {
	result = x.seed
	for i := 0; i < x.iterations; i++ {
		result = (x.mult*result + x.increment) % x.mod
	}
	return result
}
func (x *rng) pureMultiplicative() (result int) {
	x.increment = 0
	return x.linearCongruential()
}

// The middle-square method for generating pseudorandom
// numbers begins with an n-digit integer. This number is
// squared, initial zeros are appended to ensure that the result
// has 2n digits, and its middle n digits are used to form the next
// number in the sequence. This process is repeated to generate
// additional terms.
func (x *rng) middleSquare() (result int) {
	n := countDigits(x.seed)
	result = x.seed
	for i := 0; i < x.iterations; i++ {
		temp := result
		result = 0
		temp = temp * temp
		digitsSlice := strings.Split(strconv.Itoa(temp), "")
		for len(digitsSlice) < 2*n {
			digitsSlice = append([]string{"0"}, digitsSlice...)
		}
		start := 2 * n / 4
		for j := 0; j < start*2; j++ {
			digit, err := strconv.Atoi(digitsSlice[j+start])
			if err != nil {
				log.Fatal(err)
			}
			result += digit * (int(math.Pow(10, float64(n-j-1))))
		}
		fmt.Printf("result: %d, n: %d, n2: %d, digitsSlice: %v\n", result, n, 2*n, digitsSlice)
	}

	return result
}

func countDigits(n int) int {
	digits := 1
	for n > 10 {
		n = n / 10
		digits++
	}
	return digits
}

// The power generator is a method for generating pseudorandom numbers.
// To use the power generator, parameters p and d are specified,
// where p is a prime, d is a positive integer such that p does not divide d, and a seed x_0 is specified.
// The pseudorandom numbers x_1, x_2,... are generated
// using the recursive definition x_n+1 = (x_n)^d mod p.
// https://en.wikipedia.org/wiki/Blum_Blum_Shub
// Is it any good? https://crypto.stackexchange.com/questions/3454/blum-blum-shub-vs-aes-ctr-or-other-csprngs#3456
func (x *rng) blumBlumShub() (result int) {
	result = x.seed
	for i := 0; i < x.iterations; i++ {
		result = int(math.Pow(float64(result), float64(x.mult))) % x.mod
		println(result)
	}
	return result
}

type isbn struct {
	num string
}

func (x *isbn) isValid() bool {
	sum := 0
	num := strings.Split(x.num, "")
	for i := 0; i < len(x.num)-1; i++ {
		getInt, err := strconv.Atoi(num[i])
		if err != nil {
			log.Fatal(err)
		}
		sum += (i + 1) * getInt
	}
	mod11 := sum % 11
	println(mod11, sum)
	var lastDigit string = string(x.num[len(num)-1])
	if mod11 == 10 {
		if lastDigit == "X" {
			return true
		}
	}

	if strconv.Itoa(mod11) == lastDigit {
		return true
	}
	return false
}

type hash struct {
	mod       int
	values    []int
	locations map[int]int
}

func (x *hash) get() {
	for _, v := range x.values {
		key := v % x.mod
		if _, ok := x.locations[key]; ok {
			i := key
			done := false
			for !done {
				if i == v-1 {
					log.Fatal("all hash values used\n")
				}
				if _, ok := x.locations[i]; !ok {
					x.locations[i] = v
					println("break")
					done = true
				}
				i++
				i = i % x.mod
			}
		} else {
			x.locations[v%x.mod] = v
		}
	}
	println("exit")
}
