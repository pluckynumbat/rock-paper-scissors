package main

import (
	"net/http"
	"net/http/httptest"
	"os"
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

func TestPlayChoiceHandlers(t *testing.T) {

	gameServer := &GameServer{}

	var tests = []struct {
		name         string
		endpoint     string
		serverChoice engine.Choice
		want         string
	}{
		{"player plays rock, server plays rock", "/play-rock", engine.Rock, "You chose Rock\nServer chose Rock\nNo One Won!\n"},
		{"player plays rock, server plays paper", "/play-rock", engine.Paper, "You chose Rock\nServer chose Paper\nServer Won!\n"},
		{"player plays rock, server plays scissors", "/play-rock", engine.Scissors, "You chose Rock\nServer chose Scissors\nYou Won!\n"},
		{"player plays paper, server plays rock", "/play-paper", engine.Rock, "You chose Paper\nServer chose Rock\nYou Won!\n"},
		{"player plays paper, server plays paper", "/play-paper", engine.Paper, "You chose Paper\nServer chose Paper\nNo One Won!\n"},
		{"player plays paper, server plays scissors", "/play-paper", engine.Scissors, "You chose Paper\nServer chose Scissors\nServer Won!\n"},
		{"player plays scissors, server plays rock", "/play-scissors", engine.Rock, "You chose Scissors\nServer chose Rock\nServer Won!\n"},
		{"player plays scissors, server plays paper", "/play-scissors", engine.Paper, "You chose Scissors\nServer chose Paper\nYou Won!\n"},
		{"player plays scissors, server plays scissors", "/play-scissors", engine.Scissors, "You chose Scissors\nServer chose Scissors\nNo One Won!\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gameServer.fixedChoice = test.serverChoice

			newReq, err := http.NewRequest(http.MethodGet, test.endpoint, nil)
			if err != nil {
				t.Fatalf("new request failed with error: %v", err)
			}

			newResp := httptest.NewRecorder()

			gameServer.ServeHTTP(newResp, newReq)

			want := test.want
			got := newResp.Body.String()

			if got != want {
				t.Errorf("%v handler gave incorrect results, want: %v, got: %v", test.endpoint, want, got)
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
		{"player-scissors-server-rock", playScissorsAgainstServer, "playScissorsAgainstServer", engine.Rock, "You chose Scissors\nServer chose Rock\nServer Won!\n"},
		{"player-scissors-server-paper", playScissorsAgainstServer, "playScissorsAgainstServer", engine.Paper, "You chose Scissors\nServer chose Paper\nYou Won!\n"},
		{"player-scissors-server-scissors", playScissorsAgainstServer, "playScissorsAgainstServer", engine.Scissors, "You chose Scissors\nServer chose Scissors\nNo One Won!\n"},
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

func TestGetPortFromEnv(t *testing.T) {

	var tests = []struct {
		name string
		port string
		want string
	}{
		{"blank", "", defaultPort},
		{"default", defaultPort, defaultPort},
		{"env-8081", "8081", "8081"},
		{"env-8082", "8082", "8082"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			os.Setenv(portEnvironmentVariable, test.port)
			got := getPortFromEnv()
			if got != test.want {
				t.Errorf("got incorrect results, want: %v, got: %v", test.want, got)
			}
		})
	}
}

func TestCreateListenAddress(t *testing.T) {
	var tests = []struct {
		name string
		host string
		port string
		want string
	}{
		{"blank", "", "", ":"},
		{"port-only-default", "", defaultPort, ":8080"},
		{"port-only", "", "8081", ":8081"},
		{"local-host-default-port", "localhost", defaultPort, "localhost:8080"},
		{"local-host-port-specified", "localhost", "8081", "localhost:8081"},
		{"loop-back-default-port", "127.0.0.1", defaultPort, "127.0.0.1:8080"},
		{"loop-back-port-specified", "127.0.0.1", "8081", "127.0.0.1:8081"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := createListenAddress(test.host, test.port)
			if got != test.want {
				t.Errorf("got incorrect results, want: %v, got: %v", test.want, got)
			}
		})
	}
}
