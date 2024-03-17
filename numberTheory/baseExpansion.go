package numberTheory

import (
	"fmt"
	"math"
)

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

// Rosen - Discrete Math p. 275/254 and https://www.youtube.com/watch?v=cbGB__V8MNk
// AKA - SquareMultiplyAlgorithm
func modularExponentiation(base, power, m int) int {
	x := 1
	pow := base % power
	binary := decToBase(power, 2)
	for i, j := 0, len(binary)-1; i < j; i, j = i+1, j-1 {
		binary[i], binary[j] = binary[j], binary[i]
	}
	// fmt.Println(binary)
	for _, v := range binary {
		// fmt.Println("v", v, "pow", pow, "x", x, "m", m)
		if v == 1 {
			x = (x * pow) % m
		}
		pow = (pow * pow) % m
	}
	// fmt.Println("pow", pow, "x", x, "m", m)
	return x
}

func onesComplement(decNumber int) []int {
	if decNumber >= 0 {
		binary := decToBase(decNumber, 2)
		return append([]int{0}, binary...)
	}
	binary := decToBase(-decNumber, 2)
	complement := []int{}
	for _, v := range binary {
		if v == 1 {
			complement = append(complement, 0)
		} else {
			complement = append(complement, 1)
		}
	}
	return append([]int{1}, complement...)
}
func padSignedBinary(bin []int, bits int) []int {
	l := len(bin)
	negative := bin[0] == 1
	if negative {
		bin[0] = 0
	}
	for i := l; i < bits; i++ {
		bin = append([]int{0}, bin...)
	}
	if negative {
		bin[0] = 1
	}
	return bin
}

// We must assume that the sum actually represents a number in the appropriate range. Assume that n bits
// are being used, so that numbers strictly between -2n-l and 2n-l can be represented. The answer is almost,
// but not quite, that to obtain the one's complement representation of the sum of two numbers, we simply add
// the two strings representing these numbers using Algorithm 3. Instead, after performing this operation, there
// may be a carry out of the left-most column; in such a case, we then add 1 more to the answer. For example,
// suppose that n = 4; then numbers from -7 to 7 can be represented. To add -5 and 3, we add 1010 and
// 0011, obtaining 1101; there was no carry out of the left-most column. Since 1101 is the one's complement
// representation of -2, we have the correct answer. On the other hand, to add -4 and -3, we add 1011
// and 1100, obtaining 1 0111. The 1 that was carried out of the left-most column is instead added to 0111,
// yielding 1000, which is the one's complement representation of - 7. A proof that this method works entails
// considering the various cases determined by the signs and magnitudes of the addends.
// a and b are use ones complement representation.
func addOnesComplementBinary(aDec, bDec int) []int {
	// This is not done. I.e. it doesn't work.
	aBin := onesComplement(aDec)
	bBin := onesComplement(bDec)
	l := len(aBin)
	if l < len(bBin) {
		l = len(bBin)
	}
	aBin = padSignedBinary(aBin, l)
	bBin = padSignedBinary(bBin, l)
	fmt.Printf("%v, %v\n", aBin, bBin)
	result := []int{}
	carry := 0
	for i := len(aBin) - 1; i > 0; i-- { // ignore the sign bit for now.
		sum := aBin[i] + bBin[i] + carry
		switch sum {
		case 3:
			result = append(result, 1)
		case 2:
			result = append(result, 0)
			carry = 1
		case 1:
			result = append(result, 1)
			carry = 0
		case 0:
			result = append(result, 0)
		}
	}
	if aBin[0] == 1 {
		result = append(result, 1)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func addTwosComplent(aDec, bDec int) []int {
	return []int{}
}

func addBinary(aDec, bDec int) []int {
	aBin := decToBase(aDec, 2)
	bBin := decToBase(bDec, 2)
	l := len(aBin)
	if l < len(bBin) {
		l = len(bBin)
	}
	aBin = padSignedBinary(aBin, l)
	bBin = padSignedBinary(bBin, l)
	fmt.Printf("%v, %v\n", aBin, bBin)
	result := []int{}
	carry := 0
	for i, v := range aBin {
		sum := v + bBin[i] + carry
		switch sum {
		case 3:
			result = append(result, 1)
		case 2:
			result = append(result, 0)
			carry = 1
		case 1:
			result = append(result, 1)
			carry = 0
		case 0:
			result = append(result, 0)
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func trialDivision(n int) bool {
	for i := 2; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
