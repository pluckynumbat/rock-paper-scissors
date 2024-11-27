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


func TestPlayRandomRequest(t *testing.T) {
	newServer := httptest.NewServer(http.HandlerFunc(randomResult))

	result, err := sendPlayRandomRequest(newServer.URL)
	defer newServer.Close()

	if err != nil {
		t.Fatalf("sendPlayRandomRequest failed. Error: %v", err)
	}

	want := "Random Result!"
	got := strings.TrimSpace(result)

	if got != want {
		t.Errorf("got incorrect results, want: %v, got: %v", want, got)
	}
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

func TestPlayPaperRequest(t *testing.T) {
	var tests = []struct {
		name    string
		handler func(w http.ResponseWriter, req *http.Request)
		want    string
	}{
		{"server plays rock", youWin, "You Win"},
		{"server plays paper", noOneWins, "No One Wins"},
		{"server plays scissors", serverWins, "Server Wins"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newServer := httptest.NewServer(http.HandlerFunc(test.handler))

			result, err := sendPlayPaperRequest(newServer.URL)
			defer newServer.Close()

			if err != nil {
				t.Fatalf("sendPlayPaperRequest failed. Error: %v", err)
			}

			want := test.want
			got := strings.TrimSpace(result)

			if got != want {
				t.Errorf("got incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPlayScissorsRequest(t *testing.T) {
	var tests = []struct {
		name    string
		handler func(w http.ResponseWriter, req *http.Request)
		want    string
	}{
		{"server plays rock", serverWins, "Server Wins"},
		{"server plays paper", youWin, "You Win"},
		{"server plays scissors", noOneWins, "No One Wins"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newServer := httptest.NewServer(http.HandlerFunc(test.handler))

			result, err := sendPlayScissorsRequest(newServer.URL)
			defer newServer.Close()

			if err != nil {
				t.Fatalf("sendPlayScissorsRequest failed. Error: %v", err)
			}

			want := test.want
			got := strings.TrimSpace(result)

			if got != want {
				t.Errorf("got incorrect results, want: %v, got: %v", want, got)
			}
		})
	}
}
