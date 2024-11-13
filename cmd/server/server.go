package main

import (
	"fmt"

	"github.com/pluckynumbat/rock-paper-scissors/engine"
)

func main() {
	fmt.Println("running the rock-paper-server...")
	playRandomGame()
}

func playRandomGame() {
	p1 := &engine.Player{Name: "Player 1"}
	p1.ChooseRandom()
	fmt.Printf("%v chose %v\n", p1.String(), p1.PrintChoice())

	p2 := &engine.Player{Name: "Player 2"}
	p2.ChooseRandom()
	fmt.Printf("%v chose %v\n", p2.String(), p2.PrintChoice())
}
