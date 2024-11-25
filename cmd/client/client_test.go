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

func TestPlayRockAgainstServerRock(t *testing.T) {
	newServer := httptest.NewServer(http.HandlerFunc(noOneWins))

	result, err := sendPlayRockRequest(newServer.URL)
	defer newServer.Close()

	if err != nil {
		t.Fatalf("http request failed. Error: %v", err)
	}

	want := "No One Wins"
	got := strings.TrimSpace(result)

	if got != want {
		t.Errorf("incorrect results, want: %v, got: %v", want, got)
	}
}

func TestPlayRockAgainstServerPaper(t *testing.T) {
	newServer := httptest.NewServer(http.HandlerFunc(serverWins))

	result, err := sendPlayRockRequest(newServer.URL)
	defer newServer.Close()

	if err != nil {
		t.Fatalf("http request failed. Error: %v", err)
	}

	want := "Server Wins"
	got := strings.TrimSpace(result)

	if got != want {
		t.Errorf("incorrect results, want: %v, got: %v", want, got)
	}
}

func TestPlayRockAgainstServerScissors(t *testing.T) {
	newServer := httptest.NewServer(http.HandlerFunc(youWin))

	result, err := sendPlayRockRequest(newServer.URL)
	defer newServer.Close()

	if err != nil {
		t.Fatalf("http request failed. Error: %v", err)
	}

	want := "You Win"
	got := strings.TrimSpace(result)

	if got != want {
		t.Errorf("incorrect results, want: %v, got: %v", want, got)
	}
}
