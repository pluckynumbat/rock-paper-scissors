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

	serverPlayer := createPlayerWithRandomChoice("Server")

	result := ""
	var err error
	endpoint := req.URL.Path
	switch endpoint {
	case "/random":
		result, err = playRandomAgainstServer(serverPlayer)
	case "/play-rock":
		result, err = playRockAgainstServer(serverPlayer)
	case "/play-paper":
		result, err = playPaperAgainstServer(serverPlayer)
	case "/play-scissors":
		result, err = playScissorsAgainstServer(serverPlayer)
	}

	if err != nil {
		http.Error(w, internalServerErrorMsg, http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, result)
	}
}

func main() {
	fmt.Println("running the rock-paper-server...")
	log.Fatal(http.ListenAndServe(":8080", &GameServer{}))
}

func playRandomGame() (string, error) {
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
		return "", err
	} else {
		fmt.Printf(" Win: %v\n", win.String())
		result += fmt.Sprintf("%v Won!\n", win.String())
	}

	return result, nil
}

func playRandomAgainstServer(serverPlayer *engine.Player) (string, error) {

	result, err := printChoicesAndPlay(createPlayerWithRandomChoice("You"), serverPlayer)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return result, nil
}

func playRockAgainstServer(serverPlayer *engine.Player) (string, error) {

	result, err := printChoicesAndPlay(createPlayerWithFixedChoice("You", engine.Rock), serverPlayer)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return result, nil
}

func playPaperAgainstServer(serverPlayer *engine.Player) (string, error) {
	result, err := printChoicesAndPlay(createPlayerWithFixedChoice("You", engine.Paper), serverPlayer)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return result, nil
}

func playScissorsAgainstServer(serverPlayer *engine.Player) (string, error) {
	result, err := printChoicesAndPlay(createPlayerWithFixedChoice("You", engine.Scissors), serverPlayer)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return result, nil
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

func printChoicesAndPlay(p1, p2 *engine.Player) (string, error) {
	if p1.PrintChoice() == "None" {
		return "", fmt.Errorf("Invalid choice for %v", p1.String())
	}

	if p2.PrintChoice() == "None" {
		return "", fmt.Errorf("Invalid choice for %v", p2.String())
	}

	fmt.Printf(" %v:%v ", p1.String(), p1.PrintChoice())
	result := fmt.Sprintf("%v chose %v\n", p1.String(), p1.PrintChoice())

	fmt.Printf(" %v:%v ", p2.String(), p2.PrintChoice())
	result += fmt.Sprintf("%v chose %v\n", p2.String(), p2.PrintChoice())

	win, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}

	fmt.Printf(" Win: %v\n", win.String())
	result += fmt.Sprintf("%v Won!\n", win.String())
	return result, nil
}
