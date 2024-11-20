package main

import (
	"net/http"
	"net/http/httptest"
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

func TestCreateServerPlayer(t *testing.T) {
	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
	}{
		{"fixed choice rock", engine.Rock},
		{"fixed choice paper", engine.Paper},
		{"fixed choice scissors", engine.Scissors},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice
			serverPlayer := gameServer.createServerPlayer()

			want := test.serverChoice.String()
			got := serverPlayer.PrintChoice()
			if want != got {
				t.Errorf("Created Server Player has incorrect choice, want: %v, got %v", want, got)
			}
		})
	}

	t.Run("fixed choice none", func(t *testing.T) {

		gameServer.fixedChoice = engine.None
		serverPlayer := gameServer.createServerPlayer()

		choiceString := serverPlayer.PrintChoice()
		options := []engine.Choice{engine.Rock, engine.Paper, engine.Scissors}

		found := false
		for _, val := range options {
			if choiceString == val.String() {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Created Server Player has incorrect choice, wanted: %v or %v or %v,  got: %v", engine.Rock.String(), engine.Paper.String(), engine.Scissors.String(), choiceString)
		}
	})
}

func TestPlayRockHandler(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
		want         string
	}{
		{"server plays rock", engine.Rock, "You chose Rock\nServer chose Rock\nNo One Won!\n"},
		{"server plays paper", engine.Paper, "You chose Rock\nServer chose Paper\nServer Won!\n"},
		{"server plays scissors", engine.Scissors, "You chose Rock\nServer chose Scissors\nYou Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice

			newReq, err := http.NewRequest(http.MethodGet, "/play-rock", nil)
			if err != nil {
				t.Fatalf("new request failed with error: %v", err)
			}

			newResp := httptest.NewRecorder()

			gameServer.ServeHTTP(newResp, newReq)

			want := test.want
			got := newResp.Body.String()

			if got != want {
				t.Errorf("play rock against server rock gave incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPlayRockAgainstServerFunction(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
		want         string
	}{
		{"server rock", engine.Rock, "You chose Rock\nServer chose Rock\nNo One Won!\n"},
		{"server paper", engine.Paper, "You chose Rock\nServer chose Paper\nServer Won!\n"},
		{"server scissors", engine.Scissors, "You chose Rock\nServer chose Scissors\nYou Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice
			serverPlayer := gameServer.createServerPlayer()

			want := test.want
			got, err := playRockAgainstServer(serverPlayer)

			if err != nil {
				t.Errorf("playRockAgainstServer failed with error: %v", err)
			}

			if want != got {
				t.Errorf("playRockAgainstServer has incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}
