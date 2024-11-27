package main

import (
	"fmt"
	"net/http"
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

func TestCreateServerURLPrefix(t *testing.T) {
	var tests = []struct {
		name       string
		baseURL    string
		portNumber string
		want       string
	}{
		{"blank", "", "", "http://:"},
		{"local-host", "localhost", "8080", "http://localhost:8080"},
		{"loop-back", "127.0.0.1", "8080", "http://127.0.0.1:8080"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := createServerURLPrefix(test.baseURL, test.portNumber)
			if got != test.want {
				t.Errorf("got incorrect results, want: %v, got: %v", test.want, got)
			}
		})
	}
}
