package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

const internalServerErrorMsg string = "error: internal server error"

func main() {
	fmt.Println("running the rock-paper-server...")

	http.HandleFunc("/random", playRandomGame)
	http.HandleFunc("/play-rock", playRockAgainstServer)
	http.HandleFunc("/play-paper", playPaperAgainstServer)
	http.HandleFunc("/play-scissors", playScissorsAgainstServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func playRandomGame(w http.ResponseWriter, req *http.Request) {
	p1 := &engine.Player{Name: "Player 1"}
	p1.ChooseRandom()
	fmt.Printf(" %v:%v ", p1.String(), p1.PrintChoice())
	fmt.Fprintf(w, "%v chose %v\n", p1.String(), p1.PrintChoice())

	p2 := &engine.Player{Name: "Player 2"}
	p2.ChooseRandom()
	fmt.Printf(" %v:%v ", p2.String(), p2.PrintChoice())
	fmt.Fprintf(w, "%v chose %v\n", p2.String(), p2.PrintChoice())

	win, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf(" Win: %v\n", win.String())
		fmt.Fprintf(w, "%v Won!\n", win.String())
	}
}

func playRockAgainstServer(w http.ResponseWriter, req *http.Request) {
	result, err := playAgainstServer(createPlayerFromChoice(engine.Rock))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, result)
	}
}

func playPaperAgainstServer(w http.ResponseWriter, req *http.Request) {

	result, err := playAgainstServer(createPlayerFromChoice(engine.Paper))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, result)
	}
}

func playScissorsAgainstServer(w http.ResponseWriter, req *http.Request) {

	result, err := playAgainstServer(createPlayerFromChoice(engine.Scissors))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, result)
	}
}

func createPlayerFromChoice(ch engine.Choice) *engine.Player {
	player := &engine.Player{Name: "You"}
	player.ChooseFixed(ch)
	return player
}

func playAgainstServer(p1 *engine.Player) (string, error) {
	if p1.PrintChoice() == "None" {
		return "", fmt.Errorf("Invalid choice")
	}
	fmt.Printf(" Player:%v ", p1.PrintChoice())
	result := fmt.Sprintf("%v chose %v\n", p1.String(), p1.PrintChoice())

	p2 := &engine.Player{Name: "Server"}
	p2.ChooseRandom()
	fmt.Printf(" %v:%v ", p2.String(), p2.PrintChoice())
	result += fmt.Sprintf("%v chose %v\n", p2.String(), p2.PrintChoice())

	win, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf(" Win: %v\n", win.String())
		result += fmt.Sprintf("%v Won!\n", win.String())
	}

	return result, nil
}
