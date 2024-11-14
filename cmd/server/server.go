package main

import (
	"fmt"
	"net/http"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

func main() {
	fmt.Println("running the rock-paper-server...")

	http.HandleFunc("/random", playRandomGame)
	http.ListenAndServe(":8080", nil)
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
