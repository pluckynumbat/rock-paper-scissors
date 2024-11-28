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

func TestSendServerRequest(t *testing.T) {

	newServer := httptest.NewServer(http.HandlerFunc(randomResult))
	res, err := sendServerRequest(newServer.URL, "play-random")
	if err != nil {
		t.Fatalf("sendServerRequest failed, error: %v", err)
	}

	want := "Random Result!"
	got := strings.TrimSpace(res)
	if got != want {
		t.Errorf("sendServerRequest gave incorrect results, want: %v, got %v", want, got)
	}
}
