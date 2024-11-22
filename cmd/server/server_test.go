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
				t.Errorf("play-rock handler gave incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPlayPaperHandler(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
		want         string
	}{
		{"server plays rock", engine.Rock, "You chose Paper\nServer chose Rock\nYou Won!\n"},
		{"server plays paper", engine.Paper, "You chose Paper\nServer chose Paper\nNo One Won!\n"},
		{"server plays scissors", engine.Scissors, "You chose Paper\nServer chose Scissors\nServer Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice

			newReq, err := http.NewRequest(http.MethodGet, "/play-paper", nil)
			if err != nil {
				t.Fatalf("new request failed with error: %v", err)
			}

			newResp := httptest.NewRecorder()

			gameServer.ServeHTTP(newResp, newReq)

			want := test.want
			got := newResp.Body.String()

			if got != want {
				t.Errorf("play-paper handler gave incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPlayScissorsHandler(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
		want         string
	}{
		{"server plays rock", engine.Rock, "You chose Scissors\nServer chose Rock\nServer Won!\n"},
		{"server plays paper", engine.Paper, "You chose Scissors\nServer chose Paper\nYou Won!\n"},
		{"server plays scissors", engine.Scissors, "You chose Scissors\nServer chose Scissors\nNo One Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice

			newReq, err := http.NewRequest(http.MethodGet, "/play-scissors", nil)
			if err != nil {
				t.Fatalf("new request failed with error: %v", err)
			}

			newResp := httptest.NewRecorder()

			gameServer.ServeHTTP(newResp, newReq)

			want := test.want
			got := newResp.Body.String()

			if got != want {
				t.Errorf("play-scissors handler gave incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPlayAgainstServerFunctions(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		fn           func(*engine.Player) (string, error)
		fnName       string
		serverChoice engine.Choice
		want         string
	}{
		{"player-rock-server-rock", playRockAgainstServer, "playRockAgainstServer", engine.Rock, "You chose Rock\nServer chose Rock\nNo One Won!\n"},
		{"player-rock-server-paper", playRockAgainstServer, "playRockAgainstServer", engine.Paper, "You chose Rock\nServer chose Paper\nServer Won!\n"},
		{"player-rock-server-scissors", playRockAgainstServer, "playRockAgainstServer", engine.Scissors, "You chose Rock\nServer chose Scissors\nYou Won!\n"},
		{"player-paper-server-rock", playPaperAgainstServer, "playPaperAgainstServer", engine.Rock, "You chose Paper\nServer chose Rock\nYou Won!\n"},
		{"player-paper-server-paper", playPaperAgainstServer, "playPaperAgainstServer", engine.Paper, "You chose Paper\nServer chose Paper\nNo One Won!\n"},
		{"player-paper-server-scissors", playPaperAgainstServer, "playPaperAgainstServer", engine.Scissors, "You chose Paper\nServer chose Scissors\nServer Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice
			serverPlayer := gameServer.createServerPlayer()

			want := test.want
			got, err := test.fn(serverPlayer)

			if err != nil {
				t.Errorf("%v failed with error: %v", test.fnName, err)
			}

			if want != got {
				t.Errorf("%v has incorrect results, want: %v, got: %v", test.fnName, want, got)
			}
		})
	}
}

func TestPlayPaperAgainstServerFunction(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
		want         string
	}{
		{"server rock", engine.Rock, "You chose Paper\nServer chose Rock\nYou Won!\n"},
		{"server paper", engine.Paper, "You chose Paper\nServer chose Paper\nNo One Won!\n"},
		{"server scissors", engine.Scissors, "You chose Paper\nServer chose Scissors\nServer Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice
			serverPlayer := gameServer.createServerPlayer()

			want := test.want
			got, err := playPaperAgainstServer(serverPlayer)

			if err != nil {
				t.Errorf("playPaperAgainstServer failed with error: %v", err)
			}

			if want != got {
				t.Errorf("playPaperAgainstServer has incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPlayScissorsAgainstServerFunction(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		serverChoice engine.Choice
		want         string
	}{
		{"server rock", engine.Rock, "You chose Scissors\nServer chose Rock\nServer Won!\n"},
		{"server paper", engine.Paper, "You chose Scissors\nServer chose Paper\nYou Won!\n"},
		{"server scissors", engine.Scissors, "You chose Scissors\nServer chose Scissors\nNo One Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice
			serverPlayer := gameServer.createServerPlayer()

			want := test.want
			got, err := playScissorsAgainstServer(serverPlayer)

			if err != nil {
				t.Errorf("playScissorsAgainstServer failed with error: %v", err)
			}

			if want != got {
				t.Errorf("playScissorsAgainstServer has incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}
