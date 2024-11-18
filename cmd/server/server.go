package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

const internalServerErrorMsg string = "error: internal server error"

type GameServer struct{}

func (gs *GameServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	endpoint := req.URL.Path
	switch endpoint {
	case "/random":
		result, err := playRandomGame()
		if err != nil {
			http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, result)
		}
	case "/play-rock":
		result, err := playRockAgainstServer()
		if err != nil {
			http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, result)
		}
	case "play-paper":
		playPaperAgainstServer(w, req)
	case "/play-scissors":
		playScissorsAgainstServer(w, req)
	}
}

func main() {
	fmt.Println("running the rock-paper-server...")
	log.Fatal(http.ListenAndServe(":8080", &GameServer{}))
}

func playRandomGame(w http.ResponseWriter, req *http.Request) {
	result := ""

	p1 := createPlayerWithRandomChoice("Player 1")
	fmt.Printf(" %v:%v ", p1.String(), p1.PrintChoice())
	result += fmt.Sprintf("%v chose %v\n", p1.String(), p1.PrintChoice())

	p2 := createPlayerWithRandomChoice("Player 2")
	fmt.Printf(" %v:%v ", p2.String(), p2.PrintChoice())
	result += fmt.Sprintf("%v chose %v\n", p2.String(), p2.PrintChoice())

	win, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Printf(" Win: %v\n", win.String())
		result += fmt.Sprintf("%v Won!\n", win.String())
		fmt.Fprintf(w, result)
	}
}

func playRockAgainstServer(w http.ResponseWriter, req *http.Request) {
	result, err := playChoiceAgainstServer(engine.Rock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, result)
	}
}

func playPaperAgainstServer(w http.ResponseWriter, req *http.Request) {

	result, err := playChoiceAgainstServer(engine.Paper)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, result)
	}
}

func playScissorsAgainstServer(w http.ResponseWriter, req *http.Request) {

	result, err := playChoiceAgainstServer(engine.Scissors)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, result)
	}
}

func createPlayerWithRandomChoice(name string) *engine.Player {
	player := &engine.Player{Name: name}
	player.ChooseRandom()
	return player
}

func createPlayerWithFixedChoice(name string, ch engine.Choice) *engine.Player {
	player := &engine.Player{Name: name}
	player.ChooseFixed(ch)
	return player
}

func playChoiceAgainstServer(ch engine.Choice) (string, error) {
	if ch.String() == "None" {
		return "", fmt.Errorf("Invalid choice")
	}

	p1 := createPlayerWithFixedChoice("You", ch)
	fmt.Printf(" Player:%v ", p1.PrintChoice())
	result := fmt.Sprintf("%v chose %v\n", p1.String(), p1.PrintChoice())

	p2 := createPlayerWithRandomChoice("Server")
	fmt.Printf(" Server:%v ", p2.PrintChoice())
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
