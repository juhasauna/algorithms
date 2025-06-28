package ntu

// Chapter 6 From Manber file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/Me/Books/CS/Manber%20-%20Introduction%20to%20Algorithms_%20A%20Creative%20Approach%201989_v2.pdf
// And file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch6_notes_a.pdf
// And Home work 6 (for 2024 file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw6.pdf)

import (
	"algorithms/ut"
	"fmt"
	"log"
	"math"
	"slices"
	"testing"
)

func (x *CH6) scheduleRoundRobin(playerCount int) {
	if !ut.IsPowerOfTwo(playerCount) {
		log.Fatalf("only powers of two supported: got %d", playerCount)
	}

	// Create player IDs starting from 1
	players := make([]int, playerCount)
	for i := range players {
		players[i] = i + 1
	}

	x.recursiveRoundRobin(players, 1)
}

func (x *CH6) recursiveRoundRobin(players []int, startRound int) {
	n := len(players)
	if n == 2 {
		// Base case: one match
		fmt.Printf("round: %d\t%d vs %d\n", startRound, players[0], players[1])
		return
	}

	half := n / 2
	left := players[:half]
	right := players[half:]

	// Recursive calls for intra-group matches
	x.recursiveRoundRobin(left, startRound)
	x.recursiveRoundRobin(right, startRound)

	// Cross matches between left and right halves
	for i := range half {
		for r := range half {
			round := startRound + r + half
			p1 := left[(i+r)%half]
			p2 := right[i]
			fmt.Printf("round: %d\t%d vs %d\n", round, p1, p2)
		}
	}
}

func (x *CH6) RoundRobinCGPTJuha(players []int) [][][2]int {
	n := len(players)
	if !ut.IsPowerOfTwo(n) {
		log.Fatal("only powers of two for player count")
	}
	if n == 2 {
		return [][][2]int{{{players[0], players[1]}}}
	}
	A := players[:n/2]
	B := players[n/2:]
	schedA := x.RoundRobinCGPT(A)
	schedB := x.RoundRobinCGPT(B)
	// Combine internal matches
	result := make([][][2]int, 0, n-1)
	for i := 0; i < n/2-1; i++ {
		round := append(schedA[i], schedB[i]...)
		result = append(result, round)
	}
	// Cross matches (A vs B)
	for r := 0; r < n/2; r++ {
		round := make([][2]int, 0, n/2)
		for i := 0; i < n/2; i++ {
			round = append(round, [2]int{A[i], B[(i+r)%(n/2)]})
		}
		result = append(result, round)
	}
	return result
}
func (x *CH6) RoundRobinCGPT(players []int) [][][2]int {
	n := len(players)
	if !ut.IsPowerOfTwo(n) {
		log.Fatal("only powers of two for player count")
	}
	if n == 2 {
		return [][][2]int{{{players[0], players[1]}}}
	}
	A := players[:n/2]
	B := players[n/2:]
	schedA := x.RoundRobinCGPT(A)
	schedB := x.RoundRobinCGPT(B)
	// Combine internal matches
	result := make([][][2]int, 0, n-1)
	for i := 0; i < n/2-1; i++ {
		round := append(schedA[i], schedB[i]...)
		result = append(result, round)
	}
	// Cross matches (A vs B)
	for r := 0; r < n/2; r++ {
		round := make([][2]int, 0, n/2)
		for i := 0; i < n/2; i++ {
			round = append(round, [2]int{A[i], B[(i+r)%(n/2)]})
		}
		result = append(result, round)
	}
	return result
}

// Simplified Gemini version
func (x *CH6) RoundRobinScheduleGeminiJuha(playerCount int) {
	// Create a slice to represent the players, numbered 1 to playerCount.
	players := make([]int, playerCount)
	for i := range playerCount {
		players[i] = i + 1
	}

	if !ut.IsPowerOfTwo(playerCount) {
		log.Fatal("only powers of two for player count")
	}

	numRounds := playerCount - 1
	matchesPerRound := playerCount / 2

	for r := range numRounds {
		fmt.Printf("Round %d:\t", r+1)

		// Create the pairings for the current round.
		for i := range matchesPerRound {
			p1 := players[i]
			p2 := players[playerCount-1-i]

			// Check if a player has a bye in this round.
			// if hasBye && p1 == 0 {
			// 	fmt.Printf("%d has a bye\t", p2)
			// } else if hasBye && p2 == 0 {
			// 	fmt.Printf("%d has a bye\t", p1)
			// } else {
			fmt.Printf("%d vs %d\t", p1, p2)
			// }
		}
		fmt.Println()

		// Rotate the players for the next round. Player at index 0 is fixed.
		if playerCount > 2 {
			// Save the last player in the list.
			lastPlayer := players[playerCount-1]

			// Shift all players from index 1 to the end one position to the right.
			copy(players[2:], players[1:playerCount-1])

			// Move the saved last player to the second position in the list.
			players[1] = lastPlayer
		}
	}
}

// Here "bye" handles odd numbered groups.
func (x *CH6) RoundRobinScheduleGeminiBye(playerCount int) {
	// Create a slice to represent the players, numbered 1 to playerCount.
	players := make([]int, playerCount)
	for i := 0; i < playerCount; i++ {
		players[i] = i + 1
	}

	// If the number of players is odd, add a dummy player (0) to represent a "bye".
	// The player matched with the "bye" player in a round does not play.
	hasBye := false
	if playerCount%2 != 0 {
		players = append(players, 0) // 0 is the placeholder for a bye
		playerCount++
		hasBye = true
	}

	numRounds := playerCount - 1
	matchesPerRound := playerCount / 2

	for r := 0; r < numRounds; r++ {
		fmt.Printf("Round %d:\t", r+1)

		// Create the pairings for the current round.
		for i := range matchesPerRound {
			p1 := players[i]
			p2 := players[playerCount-1-i]

			// Check if a player has a bye in this round.
			if hasBye && p1 == 0 {
				fmt.Printf("%d has a bye\t", p2)
			} else if hasBye && p2 == 0 {
				fmt.Printf("%d has a bye\t", p1)
			} else {
				fmt.Printf("%d vs %d\t", p1, p2)
			}
		}
		fmt.Println()

		// Rotate the players for the next round. Player at index 0 is fixed.
		if playerCount > 2 {
			// Save the last player in the list.
			lastPlayer := players[playerCount-1]

			// Shift all players from index 1 to the end one position to the right.
			copy(players[2:], players[1:playerCount-1])

			// Move the saved last player to the second position in the list.
			players[1] = lastPlayer
		}
	}
}

// Translated from pseudocode: file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw6_s.pdf
func (x *CH6) RoundRobinSchedule(left, right, playerCount int, msg string) {
	// This does not quite work perfectly. It assigns multiple matches for the same player on the same round.
	nextCount := playerCount / 2
	med := (left + right) / 2
	// fmt.Printf("%s\t %d\t%d\t%d \tplayerCount: %d\n", msg, left, med, right, playerCount)
	if right-left == 1 {
		fmt.Printf("Round 1: %d vs %d\t", left, right)
		return
	} else {
		x.RoundRobinSchedule(left, med, nextCount, "1.rec.")
		x.RoundRobinSchedule(med+1, right, nextCount, "2.rec.")
	}
	fmt.Println()
	for r := nextCount; r < playerCount; r++ {
		for shift := range med - left + 1 {
			p1 := left + shift
			p2 := med + 1 + (r % nextCount)
			fmt.Printf("Round %d: %d vs %d\t", r, p1, p2)
		}
		fmt.Println()
	}
}

func (x *CH6) RoundRobinScheduleGrok(left, right, playerCount int) {
	nextCount := playerCount / 2
	med := (left + right) / 2
	if right-left == 1 {
		fmt.Printf("Round 1: %d vs %d\t", left, right)
		return
	} else {
		x.RoundRobinScheduleGrok(left, med, nextCount)
		x.RoundRobinScheduleGrok(med+1, right, nextCount)
	}
	fmt.Println()
	for r := nextCount; r < playerCount; r++ {
		for shift := range med - left + 1 {
			p1 := left + shift
			p2 := med + 1 + (shift+r)%nextCount
			fmt.Printf("Round %d: %d vs %d\t", r, p1, p2)
		}
		fmt.Println()
	}
}

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw6_s.pdf
// FindSubset - Subsetsum
func (x *CH6) HW6_24_03_pseudo(set []int) []int {
	sum := 0
	length := len(set)
	remSet := make([]int, length) // Elements are remainders. The possible remainders are [0, n). The zero index is not used!
	for k, v := range set {
		sum += v
		rem := sum % length
		x.iters++
		if rem == 0 {
			return set[:k+1]
		}
		fmt.Printf("subSet[rem]: %d, rem: %d, k: %d\t\t%v\n", remSet[rem], rem, k, remSet)
		if remSet[rem] != 0 { // Remainders of S_i and S_k are the same => take S_i \ S_k = Y. Sum of Y ≡ 0 (mod n)
			i := remSet[rem] + 1
			set_k_substract_set_i := set[i : k+1]
			return set_k_substract_set_i
		}
		remSet[rem] = k
	}
	temp := []int{}
	for i, v := range remSet {
		if v != 0 {
			temp = append(temp, set[i])
		}
	}
	return temp
}
func (x *CH6) HW6_24_03_bruteforce_WRONG(set []int) []int {
	length := len(set)
	for i, v := range set {
		w := v % length
		if w == 0 {
			return []int{v}
		}
		for j := i + 1; j < length; j++ {
			if (w+set[j])%length == 0 {
				return []int{v, set[j]}
			}
		}
	}
	log.Fatalf("HW6_24_03_bruteforce is wrong (or input %v is wrong)", set)
	return []int{}
}

// Find the majority in input if there is one. θ(n) time. Worst case 2n-2 comparisons.
func (x *CH6) Majority(seq []int) int {

	length := len(seq)
	if length < 1 {
		log.Fatalln("Majority: invalid input")
	}
	for _, v := range seq {
		if v < 0 {
			log.Fatalln("Majority: invalid input")
		}
	}

	candidate := seq[0]
	mult := 1
	for i := 1; i < length; i++ {
		if mult == 0 {
			candidate = seq[i]
			mult = 1
		} else if candidate == seq[i] {
			mult++
		} else {
			mult--
		}
	}
	count := 0
	if mult == 0 {
		return -1
	}
	for i := range length {
		if seq[i] == candidate {
			count++
		}
	}
	if count > length/2 {
		return candidate
	}
	return -1
}

func (x *CH6) minimumEditDistance(text1, text2 string) [][]int {
	rows, cols := len(text1)+1, len(text2)+1
	costMatrix := make([][]int, rows)
	for i := range costMatrix {
		costMatrix[i] = make([]int, cols)
	}
	for i := range rows {
		costMatrix[i][0] = i
	}
	for j := 1; j < cols; j++ {
		costMatrix[0][j] = j
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			x := costMatrix[i-1][j] + 1
			y := costMatrix[i][j-1] + 1
			z := costMatrix[i-1][j-1]
			if text1[i-1] != text2[j-1] {
				z++
			}
			costMatrix[i][j] = ut.Min(x, y, z)
		}
	}

	return costMatrix
}

func (x CH6) stringMatchKMP(text1, text2 string) int {
	// The book uses 1 based index which makes this much harder than it needs to be.
	next := x.computeKMPNext(text2)
	len1, len2 := len(text1), len(text2)
	result := 0
	i, j := 0, 1
	for result == 0 && i < len1 {
		if text2[j] == text1[i] {
			j++
			i++
		} else {
			j = next[j-1] + 1
			if j == 0 {
				j = 1
				i++
			}
		}
		if j == len2 {
			result = i - (len2 - 1)
		}
	}
	return result
}

func (x CH6) computeKMPNext(text string) []int {
	length := len(text)
	fmt.Println("len: ", length, "text: ", text)
	next := []int{-1, 0}
	for i := 2; i < length; i++ {
		j := next[i-1] + 1
		fmt.Println(i, j, next, string(text[i-1]), string(text[j]))
		for text[i-1] != text[j-1] {
			fmt.Printf("j=%d\ttext[i]=%c text[j]=%c\n", j, text[i-1], text[j])
			j = next[j-1] + 1
			if j == 0 {
				break
			}
		}
		next = append(next, j)
	}
	return next
}

func (x CH6) computeKMPNextGemini(pattern string) []int {
	length := len(pattern)
	if length == 0 {
		return []int{}
	}

	// lps is the array that will hold the Longest Proper Prefix suffix values for the pattern.
	lps := make([]int, length)
	lps[0] = 0 // lps[0] is always 0.

	// 'currentLPSLength' keeps track of the length of the previous longest prefix suffix.
	currentLPSLength := 0

	// The loop calculates lps[i] for i from 1 to length-1.
	i := 1
	for i < length {
		// Case 1: The characters match.
		// We found a longer prefix that is also a suffix.
		if pattern[i] == pattern[currentLPSLength] {
			currentLPSLength++
			lps[i] = currentLPSLength
			i++
		} else { // Case 2: The characters do not match.
			if currentLPSLength != 0 {
				// This is the key step of the KMP algorithm. We do not restart from scratch.
				// We fall back to the LPS value of the previous state, effectively trying a shorter prefix.
				// We do NOT increment 'i' here, as we need to re-evaluate pattern[i] against the new, shorter prefix.
				currentLPSLength = lps[currentLPSLength-1]
			} else {
				// If currentLPSLength is 0, it means we couldn't find a shorter prefix to match.
				// We set lps[i] to 0 and move to the next character in the pattern.
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// FindMinAndMax() find min and max in [1.5*n-2]-comparisons. Using a the naive approach takes 2*n-3 comparisons.
func (x *CH6) FindMinAndMax(seq []int) (int, int) {
	n := len(seq)
	switch n {
	case 0:
		return 0, 0
	case 1:
		return seq[0], seq[0]
	}
	max, min := 0, 0
	x.iters++
	if seq[0] > seq[1] {
		max, min = seq[0], seq[1]
	} else {
		max, min = seq[1], seq[0]
	}
	if n == 2 {
		return max, min
	}
	for i := 2; i < n-1; i += 2 {
		a_, b_ := seq[i], seq[i+1]
		x.iters++
		x.iters++
		x.iters++
		if b_ > a_ {
			max = ut.Max(b_, max)
			min = ut.Min(a_, min)
		} else {
			max = ut.Max(a_, max)
			min = ut.Min(b_, min)
		}
	}
	if n%2 != 0 {
		min, max = ut.NewMinAndMax(min, max, seq[n-1])
	}
	return min, max
}

func (x *CH6) KthSmallestElement(seq []int, k int) int {
	// O(n)
	// input gets mutated
	// BONUS EX. Figure out how to also return the index of the smalles element (of the original input sequence). Gemini says that there is no simple way to do this.
	x.Logf("init\t\t\t%d\t\t%v", k, seq)
	var selectk func(left, right, kth int, msg string) int
	selectk = func(left, right, kth int, msg string) int {
		x.iters++
		if x.iters > 10 {
			x.t.Fatal("overflow")
		}
		if left == right {
			return left
		}
		mid := x.Partition(seq, left, right)

		x.Logf("%s: %d %d %d\t\t%d\t\t%v", msg, left, mid, right, kth, seq)
		if mid-left+1 >= kth {
			return selectk(left, mid, kth, " left")
		}
		return selectk(mid+1, right, kth-(mid-left+1), "right")
	}

	n := len(seq)
	if k < 0 || k > n {
		return -1
	}
	index := selectk(0, n-1, k, "start")
	result := seq[index]
	return result
}

// 1. Select pivot.
// 2. Move elements to the left and right of the pivot based on the size of the elements relative to the pivot.
// 3. Return the position of the pivot.
func (x CH6) Partition(seq []int, left, right int) int {
	pivot := seq[left] // Pivot can be any value from the input sequence.
	// pivot := seq[right]
	l, r := left, right
	for l < r {
		for seq[l] <= pivot && l < r {
			l++
		}
		for seq[r] > pivot && r >= l {
			r--
		}
		if l < r {
			seq[l], seq[r] = seq[r], seq[l]
		}
	}
	mid := r
	seq[left], seq[mid] = seq[mid], seq[left]
	return mid
}

// Implementation of pseudocode: file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch6_notes_a.pdf
func (x *CH6) mergeSortPseudo(seq []int, n int) []int {
	var m_sort func(int, int)
	m_sort = func(left, right int) {
		x.iters++
		if right-left == 1 {
			if seq[left] > seq[right] {
				seq[left], seq[right] = seq[right], seq[left]
				x.t.Logf("SW: %v\t", seq)
			}
		} else if seq[left] != seq[right] {
			middle := (left + right + 1) / 2

			ut.AssertEqual(middle, int(math.Ceil(float64(left+right)/2.0)))

			x.t.Logf("LEFT %d, M%d\n", x.iters, middle)
			m_sort(left, middle-1)
			x.t.Logf("RIGHT %d, M%d\n", x.iters, middle)
			m_sort(middle, right)
			x.t.Logf("\n---START MERGE iter:%d, L%d, M%d, R%d\n", x.iters, left, middle, right)

			// This is a PARTICULARLY DIFFICULT to read implementation of the merge part. See sorter.merge() for a better version.
			temp := make([]int, n) // Alternative: temp := make(map[int]int)
			i, j, k := left, middle, 0
			for i < middle && j <= right {
				if seq[i] <= seq[j] {
					x.t.Logf("IF i%d, j%d, k%d\t", i, j, k)
					temp[k] = seq[i]
					i++
				} else {
					x.t.Logf("ELSE i%d, j%d, k%d\t", i, j, k)
					temp[k] = seq[j]
					j++
				}
				x.t.Logf("\nTEMP %v\n", temp)
				k++
			}
			if j > right {
				for t := 0; t < (middle - i); t++ {
					seq[right-t] = seq[middle-1-t]
				}
				x.t.Logf("M: %d, i: %d, X: %v\t", i, middle, seq)
			}
			for t := 0; t < k; t++ {
				seq[left+t] = temp[t]
			}
			x.t.Logf("L: %d, k: %d, X: %v\t", left, middle, seq)
			x.t.Logf("\n---END MERGE\n")
		}

	}
	m_sort(0, n-1)
	return seq
}

// WIP: Improved version of NTU pseudocode translation
func (x *CH6) mergeSortImproved(seq []int) []int {
	merge := func(left, mid, right int) {
		temp := make([]int, right)
		i, j, k := left, mid, 0
		for i < mid && j <= right {
			if seq[i] <= seq[j] {
				x.t.Logf("IF i%d, j%d, k%d\t", i, j, k)
				temp[k] = seq[i]
				i++
			} else {
				x.t.Logf("ELSE i%d, j%d, k%d\t", i, j, k)
				temp[k] = seq[j]
				j++
			}
			x.t.Logf("\nTEMP %v\n", temp)
			k++
		}
		if j > right {
			for t := 0; t < (mid - i); t++ {
				seq[right-t] = seq[mid-1-t]
			}
			x.t.Logf("M: %d, i: %d, X: %v\t", i, mid, seq)
		}
		for t := 0; t < k; t++ {
			seq[left+t] = temp[t]
		}
		x.t.Logf("L: %d, k: %d, X: %v\t", left, mid, seq)
	}
	var sort func(int, int, string)
	sort = func(left, right int, msg string) {
		x.iters++
		if right-left == 1 {
			if seq[left] > seq[right] {
				seq[left], seq[right] = seq[right], seq[left]
				x.t.Logf("SW: %v\t", seq)
			}
		} else if seq[left] != seq[right] {
			mid := (left + right + 1) / 2
			x.t.Logf("%s, iters: %d, mid: %d\n", msg, x.iters, mid)
			sort(left, mid-1, "left")
			sort(mid, right, "right")
			x.t.Logf("\n---START MERGE iter:%d, L%d, M%d, R%d\n", x.iters, left, mid, right)
			merge(left, mid, right)
			x.t.Logf("\n---END MERGE\n")
		}

	}
	sort(0, len(seq)-1, "start")
	return seq
}

func (x *CH6) straightRadixSort(seq []int) []int {
	getDigit := func(num, place int) int {
		return (num / place) % 10
	}
	if len(seq) < 2 {
		return seq
	}
	digits := func() int {
		maxVal := slices.Max(seq)
		return int(math.Floor(math.Log10(float64(maxVal))) + 1)
	}()
	place := 1
	for i := 0; i < digits; i++ {
		queues := make([][]int, 10)
		for j := 0; j < 10; j++ {
			queues[j] = []int{}
		}
		for _, v := range seq {
			d := getDigit(v, place)
			queues[d] = append(queues[d], v)
		}
		seq = []int{}
		for j := 0; j < 10; j++ {
			seq = append(seq, queues[j]...)
		}
		place *= 10
	}
	return seq
}

func (x *CH6) interpolationSearch(sortedSeq []int, target int) int {
	n := len(sortedSeq)
	if n == 0 {
		x.t.Fatalf("invalid input, n: %d", n)
	}
	if sortedSeq[0] > target || sortedSeq[n-1] < target {
		return -1
	}
	var find func(left, right int, msg string) int
	find = func(left int, right int, msg string) int {
		x.iters++
		if x.iters > 5 {
			x.t.Fatalf("too many iters: %d", x.iters)
		}
		if sortedSeq[left] == target {
			return left
		} else if left == right {
			return -1
		}
		var guess int = (target - sortedSeq[left]) * (right - left)
		x.t.Logf("%d", guess)
		var tempGuess float64 = float64(guess) / float64((sortedSeq[right] - sortedSeq[left]))
		x.t.Logf("%.2f\t", tempGuess)

		if guess == 0 {
			guess = left
		} else {
			guess = int(math.Ceil(tempGuess)) + left
		}
		x.t.Log(guess, "\t")
		x.t.Logf("L: %d G: %d, R: %d, msg: %s\n", left, guess, right, msg)
		if sortedSeq[guess] > target {
			return find(left, guess-1, "left")
		}
		return find(guess, right, "right")
	}
	return find(0, n-1, "begin")
}

func (x *CH6) canMakeStutter(text, pattern string, k int) bool {
	if k == 0 {
		return true
	}
	if len(pattern)*k > len(text) {
		return false
	}
	textIndex := 0
	for _, charToFind := range pattern {
		count := 0
		for textIndex < len(text) && count < k {
			if rune(text[textIndex]) == charToFind {
				count++
			}
			textIndex++
		}
		if count < k {
			return false
		}
	}
	return true
}

// StutteringSubsequence finds the maximum k such that pattern_k is a subsequence of text.
func (x *CH6) stutteringSubsequence(text, pattern string) int {
	if len(pattern) == 0 {
		return 1e9
	}
	low := 0
	high := len(text) / len(pattern)
	maxK := 0
	for low <= high {
		mid := low + (high-low)/2
		if mid == 0 {
			low = mid + 1
			continue
		}
		if x.canMakeStutter(text, pattern, mid) {
			maxK = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return maxK
}

// Given a sorted sequence of distinct integers determine if there exists an index i such that a_i = i.
func (x *CH6) specialBinarySearch(sortedSeq []int) int {
	var find func(left, right int, msg string) int
	find = func(left int, right int, msg string) int {
		x.Log(left, right, msg)
		x.iters++
		if x.iters > 5 {
			x.t.Fatalf("too many iters. left: %d, right:%d", left, right)
		}
		if left == right {
			if sortedSeq[left] == left {
				return left
			}
			return -1
		}
		mid := (left + right) / 2
		if sortedSeq[mid] < mid {
			return find(mid+1, right, "right")
		}
		return find(left, mid, "left")
	}
	return find(0, len(sortedSeq)-1, "begin")
}

// cyclicBinarySearch find the minimal index from a cyclically sorted seqeuence.
func (x *CH6) cyclicBinarySearch(sortedSeq []int) int {
	var find func(l, r int) int
	find = func(l, r int) int {
		if x.iters > 10 {
			x.t.Fatal("too many iters")
		}
		x.iters++
		if l == r {
			return l
		}
		mid := (l + r) / 2
		x.t.Log(l, mid, r)
		if sortedSeq[mid] < sortedSeq[r] {
			return find(l, mid)
		}
		return find(mid+1, r)
	}
	return find(0, len(sortedSeq)-1)
}

func (x *CH6) quickSort(seq []int) []int {
	partition := func(left, right int) int {
		pivot := seq[left]
		l, r := left, right
		for l < r {
			for seq[l] <= pivot && l < r {
				l++
			}
			for seq[r] > pivot && r >= l {
				r--
			}
			if l < r {
				seq[l], seq[r] = seq[r], seq[l]
			}
		}
		mid := r
		seq[left], seq[mid] = seq[mid], seq[left]
		return mid
	}
	var sortition func(left, right int, msg string)
	sortition = func(left, right int, msg string) {
		x.iters++
		if x.iters > 100 {
			x.t.Fatalf("too many iters")
		}
		if left < right {
			mid := 0
			mid = partition(left, right)
			x.t.Logf("%d, %d, %d, %v, msg: %s\n", left, mid, right, seq, msg)
			sortition(left, mid-1, "left")
			sortition(mid+1, right, "right")
		}
	}
	sortition(0, len(seq)-1, "begin")
	return seq
}

type CH6 struct {
	verbose bool
	iters   int64
	t       *testing.T
}

func (x CH6) Logf(format string, a ...any) {
	if x.verbose {
		x.t.Logf(format, a...)
	}
}
func (x CH6) Log(a ...any) {
	if x.verbose {
		x.t.Log(a...)
	}
}
