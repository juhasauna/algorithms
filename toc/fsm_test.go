package toc

import (
	"strings"
	"testing"
)

func Test_fsm(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{

		{"fsmTest", fsmTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func fsmTest(t *testing.T) {
	tests := []struct {
		fsm   FSM
		name  string
		input string
		want  bool
	}{
		{fsm_2022_H22_ii(), "odd number of 01", "", false},
		{fsm_2022_H22_ii(), "odd number of 01", "01", true},
		{fsm_2022_H22_ii(), "odd number of 01", "0101", false},
		{fsm_2022_H22_ii(), "odd number of 01", "111101", true},
		{fsm_2022_H22_ii(), "odd number of 01", "111101000000101", true},
		{fsm_2022_H22_ii(), "odd number of 01", "0000", false},
		{fsm_2022_H22_ii(), "odd number of 01", "0101010000", true},
		{fsm_2022_H22_ii(), "odd number of 01", "000001", true},
		// {fsm0011(), "epsilon", "", false},
		// {fsm0011(), "001", "001", false},
		// {fsm0011(), "0011", "0011", true},
		// {fsm0011(), "stuff+0011", "1010101010011", true},
		// {fsm0011(), "stuff+0011+stuff", "101010101001110101010101", true},
		// {fsm0011(), "0011+stuff", "001110101010101", true},

		// {fsm_not_0011(), "epsilon", "", true},
		// {fsm_not_0011(), "001", "001", true},
		// {fsm_not_0011(), "0011", "0011", false},
		// {fsm_not_0011(), "stuff+0011", "1010101010011", false},
		// {fsm_not_0011(), "stuff+0011+stuff", "101010101001110101010101", false},
		// {fsm_not_0011(), "0011+stuff", "001110101010101", false},

		// {fsmCoffeeMachine(), "1111", "1111", true},
		// {fsmCoffeeMachine(), "22", "22", true},
		// {fsmCoffeeMachine(), "211", "211", true},
		// {fsmCoffeeMachine(), "121", "121", true},
		// {fsmCoffeeMachine(), "112", "112", true},
		// {fsmCoffeeMachine(), "222", "222", true},
		// {fsmCoffeeMachine(), "11111", "11111", true},
		// {fsmCoffeeMachine(), "111", "111", false},
		// {fsmCoffeeMachine(), "", "", false},
		// {fsmCoffeeMachine(), "21", "21", false},
		// {fsmCoffeeMachine(), "2", "2", false},
	}
	for _, tt := range tests {
		name := tt.fsm.name + " " + tt.name
		t.Run(name, func(t *testing.T) {
			tt.fsm.validate(tt.input)
			sliceInput := strings.Split(tt.input, ";")
			if len(sliceInput) == 1 {
				sliceInput = strings.Split(tt.input, "")
			}
			got := tt.fsm.machine(0, sliceInput)
			// if got != tt.want {
			// 	t.Errorf("FSM() = %v, want %v", got, tt.want)
			// }
			if got != tt.want || got != tt.fsm.regexpFunc(tt.input) {
				t.Errorf("FSM() = %v, want %v/%v", got, tt.want, tt.fsm.regexpFunc(tt.input))
			}
		})
	}
}
