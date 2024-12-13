package toc

import (
	"fmt"
	"log"
	"regexp"
)

// type FSM struct {
// 	// Q []int // Set of states. ()'Formally' this should include the final states and the starting state but we're doing it differently.)
// 	//Σ   []string   // The Alphabet: set of charcaters. We're not using this. Instead we deduce the alphabet from the input?
// 	//s_0 int        // The starting state. We're using 0 for the starting state
// 	delta  map[int]map[string]int // state-transition function δ:Q x Σ -> Q.
// 	F      []int                  // Set of final states. ('Formally' should be aSubset of Q but we're doing it differently).
// 	regexp string
// }
type FSM struct {
	name       string
	delta      map[int]map[string]int // state-transition function δ:Q x Σ -> Q.
	F          []int                  // Set of final states. ('Formally' should be aSubset of Q but we're doing it differently).
	regexp     string
	regexpFunc func(string) bool
}

func fsmCEtumerkittömätReaaliluvut() FSM {
	return FSM{delta: cEtumerkittömätReaaliluvut_d(), F: []int{2, 3, 6}}
}

func cEtumerkittömätReaaliluvut_d() map[int]map[string]int {
	m := map[int]map[string]int{
		0: {".": 7, "digit": 1},
		1: {"digit": 1, ".": 2, "esp": 4},
		2: {"digit": 3, "exp": 4},
		3: {"digit": 3, "exp": 4},
		4: {"digit": 6, "+": 5, "-": 5},
		5: {"digit": 6},
		6: {"digit": 6},
		7: {"digit": 3},
	}
	return m
}
func (x FSM) machine(state int, language []string) bool {
	if len(language) == 0 {
		for _, finalState := range x.F {
			if finalState == state {
				return true
			}
		}
		return false
	}
	head, tail := language[0], language[1:]
	nextState := x.delta[state][head]
	return x.machine(nextState, tail)
}

func fsm0011() FSM {
	return FSM{name: "0011", delta: m0011_d(), F: []int{4}, regexp: "0011", regexpFunc: regexpMatch("0011")}
}
func fsm_not_0011() FSM {
	// regexp that fails with all inputs "^$a"
	// regexp that succeeds with all inputs "'*"
	return FSM{name: "not 0011",
		delta:      m0011_d(),
		F:          []int{0, 1, 2, 3},
		regexpFunc: regexpMatch("0011", true),
	}
	//regexp: "^(0|1|01|10|11|00(?!11))*$"} // This kind of regex is not supported in Go.
}

func m2022_H22_ii_d() map[int]map[string]int {
	m := map[int]map[string]int{
		0: {"0": 1, "1": 0},
		1: {"0": 1, "1": 2},
		2: {"0": 3, "1": 2},
		3: {"0": 3, "1": 0},
	}
	return m
}

func fsm_2022_H22_ii() FSM {
	return FSM{
		name:       "fsm_2022_H22_ii",
		delta:      m2022_H22_ii_d(),
		F:          []int{2, 3},
		regexpFunc: regexpMatch("\\A(0|(1|^$|00*11*00*1)(1|00*11*00*1)*0)0*(11*(0|00*(0|^$)|1|^$)|1)$"),
		regexp:     "\\A(0|(1|^$|00*11*00*1)(1|00*11*00*1)*0)0*(11*(0|00*(0|^$)|1|^$)|1)$",
	}
}

func (x FSM) getStatesAndAlphabet() ([]int, []string) {
	states := make(map[int]struct{})
	alphabet := make(map[string]struct{})
	for stateKey, f := range x.delta {
		states[stateKey] = struct{}{}
		for alphabetKey, state := range f {
			states[state] = struct{}{}
			alphabet[alphabetKey] = struct{}{}
		}
	}
	s := []int{}
	a := []string{}
	for key := range states {
		s = append(s, key)
	}
	for key := range alphabet {
		a = append(a, key)
	}
	return s, a
}
func (x FSM) validate(input string) {
	fmt.Println(input)
	if len(x.F) == 0 {
		log.Fatal("Empty set of final states. This is a degenerate case which we're not interested in.")
	}
	states, alphabet := x.getStatesAndAlphabet()
	if len(states) == 0 {
		log.Fatal("empty set given for states")
	}
	for _, c := range input {
		foundIt := false
		char := string(c)
		for _, a := range alphabet {
			if a == char {
				foundIt = true
			}
		}
		if !foundIt {
			log.Fatalf("invalid input: didn't find '%s' from the alphabet: %v\n", char, alphabet)
		}
	}
}

func fsmCoffeeMachine() FSM {
	return FSM{
		name:       "CoffeeMachine",
		delta:      mCoffeeMachine_d(),
		F:          []int{4, 5},
		regexp:     "^(22|1111|211|121|112)[12]*$",
		regexpFunc: regexpMatch("^(22|1111|211|121|112)[12]*$"),
	}
}

func regexpMatch(pattern string, optional ...bool) func(string) bool {
	reverse := false
	if len(optional) == 1 {
		reverse = optional[0]
	}
	if reverse {
		return func(input string) bool {
			match, err := regexp.MatchString(pattern, input)
			if err != nil {
				log.Fatal(err)
			}
			return !match
		}
	}
	return func(input string) bool {
		match, err := regexp.MatchString(pattern, input)
		if err != nil {
			log.Fatal(err)
		}
		return match
	}
}

func mCoffeeMachine_d() map[int]map[string]int {
	m := map[int]map[string]int{
		0: {"1": 1, "2": 2},
		1: {"1": 2, "2": 3},
		2: {"1": 3, "2": 4},
		3: {"1": 4, "2": 5},
		4: {"1": 5, "2": 5},
		5: {"1": 5, "2": 5},
	}
	return m
}

func m0011_d() map[int]map[string]int {
	m := map[int]map[string]int{
		0: {"0": 1, "1": 0},
		1: {"0": 2, "1": 0},
		2: {"0": 2, "1": 3},
		3: {"0": 1, "1": 4},
		4: {"0": 4, "1": 4},
	}
	return m
}

// func (x FSM) machine(state int, language string) bool {
// 	if len(language) == 0 {
// 		for _, finalState := range x.F {
// 			if finalState == state {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	head, tail := string(language[0]), language[1:]
// 	nextState := x.delta[state][head]
// 	return x.machine(nextState, tail)
// }
