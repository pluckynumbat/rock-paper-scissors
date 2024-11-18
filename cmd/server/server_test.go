package main

import (
	"testing"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

func TestCreatePlayerWithFixedChoice(t *testing.T) {

	var tests = []struct {
		name       string
		playerName string
		choice     engine.Choice
	}{
		{"rock", "P1", engine.Rock},
		{"paper", "P2", engine.Paper},
		{"scissors", "P3", engine.Scissors},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			want := test.playerName
			player := createPlayerWithFixedChoice(test.playerName, test.choice)
			got := player.String()
			if want != got {
				t.Errorf("Create player with fixed choice has incorrect player name, want: %v, got %v", want, got)
			}

			want = test.choice.String()
			got = player.PrintChoice()
			if want != got {
				t.Errorf("Create player with fixed choice has incorrect choice, want: %v, got %v", want, got)
			}
		})
	}
}
