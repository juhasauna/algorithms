// See p. 97 or https://en.wikipedia.org/wiki/Knuth%27s_Algorithm_X
package kolman

import (
	"fmt"
	"log"
)

type exactCover struct {
	// This works but the implementation is messy and it doesn't handle elegantly inputs that have no solution.
	sets         map[int][]int
	trySets      map[int][]int
	coverIndexes []int
}

func (x *exactCover) rowLen(m map[int][]int) int {
	return len(m)
}
func (x *exactCover) colLen(m map[int][]int) int {
	colLen := 0
	for _, v := range m {
		if len(v) > colLen {
			colLen = len(v) // Take deleted items into account. Is this required?
		}
	}
	return colLen
}

func (x *exactCover) findFewestOnesColumn() (int, error) {
	colOnesCount := []int{}
	for _, row := range x.trySets {
		for i, v := range row {
			if len(colOnesCount) == i {
				colOnesCount = append(colOnesCount, v)
			} else {
				colOnesCount[i] += v
			}
		}
	}
	var fewestOnesColumn int
	leastOnesSofar := x.rowLen(x.trySets) + 1 // init with big enough number
	for i, v := range colOnesCount {
		if v == 0 {
			return 0, fmt.Errorf("no solution since there is a zero column: %v", x.trySets)
		}
		if v < leastOnesSofar {
			fewestOnesColumn = i
			leastOnesSofar = v
		}
	}

	return fewestOnesColumn, nil
}

func (x *exactCover) initTrySet() {
	fmt.Println("initTrySet()")
	x.trySets = make(map[int][]int)
	for key, v := range x.sets {
		x.trySets[key] = v
	}
}

func (x *exactCover) isSolution() bool {
	sum := 0
	usedIndexes := make(map[int]struct{})
	for _, coverIndex := range x.coverIndexes {
		row, ok := x.sets[coverIndex]
		if !ok {
			log.Fatalf("invalid map key: %d", coverIndex)
		}
		for i, rowItem := range row {
			if rowItem == 1 {
				if _, ok := usedIndexes[i]; ok {
					fmt.Printf("isSolution() error: already used index: %d\n", i)
					return false
				}
				usedIndexes[i] = struct{}{}
			}
			sum += rowItem
		}
	}
	rowLen := x.rowLen(x.sets)
	return sum == rowLen
}

func (x *exactCover) algorithmX() error {
	if x.isSolution() {
		return nil
	}
	if x.trySets == nil {
		x.initTrySet()
	}
	if len(x.trySets) == 0 {
		fmt.Println("return nil")
		return nil
	}
	if len(x.trySets) == 1 {
		for key := range x.trySets {
			x.coverIndexes = append(x.coverIndexes, key)
			return nil
		}
	}
	fewestOnesColumn, err := x.findFewestOnesColumn()
	if err != nil {
		if x.rowLen(x.sets) > 0 && len(x.coverIndexes) > 0 {
			fmt.Printf("backtracking %v, %v\n", x.sets, x.coverIndexes)
			delete(x.sets, x.coverIndexes[0])
			x.trySets = nil
			x.coverIndexes = []int{}
			return x.algorithmX()
		} else {
			return err
		}
	}

	fmt.Println("fewestOnesColumn", fewestOnesColumn)
	// b) Choose a row r with a 1 in column c, and place the number r in coverIndexes.
	chosenRow := 0
	for key, row := range x.trySets {
		if row[fewestOnesColumn] == 1 {
			chosenRow = key
			x.coverIndexes = append(x.coverIndexes, chosenRow)
			break
		}
	}
	fmt.Println("chosenRow", chosenRow)

	// c) Eliminate superfluous rows
	chosenRowValues := []int{}
	for _, v := range x.trySets[chosenRow] {
		chosenRowValues = append(chosenRowValues, v)
	}

	// (e) Eliminate row r
	delete(x.trySets, chosenRow)
	if len(x.trySets) == 0 {
		fmt.Println("sets empty after deleting chosenRow")
		return nil
	}

	rowsToRemove := []int{}
	for key, row := range x.trySets {
		for j, v := range row {
			if v == chosenRowValues[j] && v == 1 {
				rowsToRemove = append(rowsToRemove, key)
				break
			}
		}
	}
	fmt.Printf("rows to remove %v\n", rowsToRemove)
	for _, v := range rowsToRemove {
		delete(x.trySets, v)
	}
	if len(x.trySets) == 0 {
		fmt.Println("sets empty after deleting rowsToRemove")
		return nil
	}
	// d) Eliminate all columns in which chosenRow has a 1.
	for i := len(chosenRowValues) - 1; i >= 0; i-- {
		if chosenRowValues[i] == 1 {
			x.removeColumn(i)
		}
	}

	fmt.Printf("next iteration starting, coverIndexes: %v, sets: %v\n", x.coverIndexes, x.trySets)
	return x.algorithmX()
}

func (x *exactCover) removeColumn(c int) {
	fmt.Printf("removing column: %d, from sets: %v\n", c, x.trySets)
	if c > x.colLen(x.trySets) {
		log.Fatalf("invalid remove column input, column too big: %d", c)
	}
	if c < 0 {
		log.Fatalf("invalid remove column input, column too small: %d", c)
	}
	for key, row := range x.trySets {
		x.trySets[key] = append(row[:c], row[c+1:]...)
	}
}

func removeRow(r int, m [][]int) [][]int {
	if r > len(m)-1 {
		log.Fatalf("invalid remove row input, row too big: %d", r)
	}
	if r < 0 {
		log.Fatalf("invalid remove row input, row too small: %d", r)
	}

	return append(m[:r], m[r+1:]...)
}

func characteristicFunc(universalSet, subset []int) (result []int) {
	for _, v := range subset {
		if sliceIncludes(universalSet, v) {
			result = append(result, 1)
		} else {
			result = append(result, 0)

		}
	}
	return result
}

func sliceIncludes(s []int, element int) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}
	return false
}
