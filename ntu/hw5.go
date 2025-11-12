package ntu

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Cat struct {
	target int
	likes  []int
}

func herdSetsOfCats(fileName string) []bool {
	// Open the input file for scanning. Later the file is scanned one line at a time.
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()                                 // Scan the first line, which should be the number of tests.
	testCount, err := strconv.Atoi(scanner.Text()) // Convert string to int.
	if err != nil {
		log.Fatal(err)
	}
	solution := []bool{}
	for range testCount {
		catsCount, potsCount := getNumberOfCatsAndPots(scanner)
		cats := []Cat{}

		for range catsCount {
			target, likes := getCatTargetAndLikes(scanner)
			cats = append(cats, Cat{target: target, likes: likes})
		}
		solution = append(solution, herdCats(cats, potsCount))
	}
	return solution
}

func herdCats(cats []Cat, potCount int) bool {
	for target := 1; target <= potCount; target++ {
		deferredLikes := make(map[int]struct{}) // This is a counter for the likes that must be placed at the end.

		validLikes := make(map[int]int)
		targetCatCounter := 0
		for _, cat := range cats {
			if cat.target == target {
				targetCatCounter++
				for _, likedPlant := range cat.likes {
					deferredLikes[likedPlant] = struct{}{} // In Go using "struct{}{}" in a map is the most memory efficient way to indicate that a key value exists.
					if validLikes[likedPlant] != -1 {
						validLikes[likedPlant]++ // increment existing entry
					}
				}
			} else {
				for _, aLike := range cat.likes {
					validLikes[aLike] = -1 // Indicate that pot-number cannot be used, since it is liked by cat with different target AND the cat has not been removed from the cats-list yet.
				}
			}
		}
		if targetCatCounter == 0 {
			continue // No cat has the target for this iteration.
		}

		{ // If a cat likes a plant but that plant is not used to place the cat at its target, THEN that plant must be placed AFTER the cat has been placed at its target. There must be enough room at the end to place the 'unused' plant.

			// Simple test case to illustrate this constraint:
			// 1
			// 2 3
			// 2 2 1 2
			// 2 2 2 3
			// The answer is 'no', because both 1,3 cannot be placed at the end.

			numberOfPlantsToPlaceAtTheEnd := len(deferredLikes) - 1 // '-1' is to account for the entry for the target.
			spaceLeftAtEnd := potCount - target
			isEnoughRoomAtEnd := numberOfPlantsToPlaceAtTheEnd <= spaceLeftAtEnd
			if !isEnoughRoomAtEnd {
				return false
			}
		}

		for _, likes := range validLikes {
			if likes == targetCatCounter { // Happy path
				// fmt.Printf("found VALID selection for position %d\n", target)
				// Remove cats that have been placed in a suitable position in this iteration.
				cats = slices.DeleteFunc(cats, func(el Cat) bool { return el.target == target })
				break // no need to keep looping since we already found a suitable plant.
			}
		}
	}
	return len(cats) == 0 // If all cats have been removed, then the goal has been achieved and we should return TRUE.
}

func getNumberOfCatsAndPots(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	if len(parts) != 2 {
		log.Fatal("should have two parts: one for cats and one for pots\n")
	}
	catsCount, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	potsCount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return catsCount, potsCount
}
func getCatTargetAndLikes(scanner *bufio.Scanner) (int, []int) {
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	partsCount := len(parts)
	if partsCount < 2 {
		log.Fatal("Each test line for a cat should have at least 2 parts.")
	}
	target, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("Expected int for the pot where cat should stop.")
	}
	likesCount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Expected int for the # of plants cat likes.")
	}
	likes := []int{}
	for _, s := range parts[2:] {
		like, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		likes = append(likes, like)
	}
	if likesCount != len(likes) {
		log.Fatalf("Should have %d cat liked plants but have %d.", likesCount, len(likes))
	}
	return target, likes
}
