package main

import (
	"testing"
)

func TestValidGetChoiceFromInput(t *testing.T) {

	names := []string{"P1", "P2", "P3", "p1", "p2", "p3"}
	inputs := []string{"R", "P", "S", "r", "p", "s"}
	expected := []string{"Rock", "Paper", "Scissors", "Rock", "Paper", "Scissors"}

	players := createPlayers(names...)

	for i, _ := range inputs {
		choice, err := getChoiceFromInput(inputs[i])
		if err != nil {
			t.Fatalf("Error in getChoiceFromInput for valid input: %v", inputs[i])
		}
		players[i].ChooseFixed(choice)
		want := expected[i]
		have := players[i].PrintChoice()

		if want != have {
			t.Fatalf("Error in FixedChoiceInputLoop: want: %s, have: %s", want, have)
		}
	}
}

func TestInvalidGetChoiceFromInput(t *testing.T) {

	invalid := []string{"R ", " s", "P1", "ro", "Paper", " ", "  "}
	for i, _ := range invalid {
		_, choiceErr := getChoiceFromInput(invalid[i])
		if choiceErr == nil {
			t.Fatalf("Error in getChoiceFromInput for invalid input %v", invalid[i])
		}
	}
}
