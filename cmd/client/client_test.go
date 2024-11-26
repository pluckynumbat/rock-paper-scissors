package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func noOneWins(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "No One Wins")
}

func youWin(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You Win")
}

func serverWins(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Server Wins")
}

func randomResult(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Random Result!")
}


func TestPlayRockRequest(t *testing.T) {
	var tests = []struct {
		name    string
		handler func(w http.ResponseWriter, req *http.Request)
		want    string
	}{
		{"server plays rock", noOneWins, "No One Wins"},
		{"server plays paper", serverWins, "Server Wins"},
		{"server plays scissors", youWin, "You Win"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newServer := httptest.NewServer(http.HandlerFunc(test.handler))

			result, err := sendPlayRockRequest(newServer.URL)
			defer newServer.Close()

			if err != nil {
				t.Fatalf("sendPlayRockRequest failed. Error: %v", err)
			}

			want := test.want
			got := strings.TrimSpace(result)

			if got != want {
				t.Errorf("got incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}
