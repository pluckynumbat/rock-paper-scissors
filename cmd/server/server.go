package main

import (
	"fmt"
	"net/http"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

func main() {
	fmt.Println("running the rock-paper-server...")

	http.HandleFunc("/random", playRandomGame)

}

func playRandomGame(w http.ResponseWriter, req *http.Request) {
	p1 := &engine.Player{Name: "Player 1"}
	p1.ChooseRandom()
	fmt.Printf("%v chose %v\n", p1.String(), p1.PrintChoice())

	p2 := &engine.Player{Name: "Player 2"}
	p2.ChooseRandom()
	fmt.Printf("%v chose %v\n", p2.String(), p2.PrintChoice())

	win, err := engine.Play(p1, p2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v Won!\n", win.String())
	}
}
