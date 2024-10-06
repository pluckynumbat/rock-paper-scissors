package engine

import (
	"errors"
	"math/rand/v2"
)

type Choice int

const (
	None Choice = iota
	Rock
	Paper
	Scissors
)

type Player struct {
	Name   string
	choice Choice
}

var NoPlayer = Player{"No One", None}

func (p Player) String() string {
	return p.Name
}

// implement Stringer
func (c Choice) String() string {
	var result string

	switch c {
	case 0:
		result = "None"

	case 1:
		result = "Rock"

	case 2:
		result = "Paper"

	case 3:
		result = "Scissors"
	}

	return result
}

func (c1 Choice) beats(c2 Choice) bool {

	if c1 == Rock {
		return c2 == Scissors
	} else if c1 == Paper {
		return c2 == Rock
	} else if c1 == Scissors {
		return c2 == Paper
	} else {
		return false
	}
}

func (p *Player) ChooseFixed(c Choice) {
	p.choice = c
}

func (p *Player) ChooseRandom() {
	roll := rand.IntN(3) + 1

	p.choice = Choice(roll)
}

func (p *Player) PrintChoice() string {
	return p.choice.String()
}

func Play(p1, p2 *Player) (*Player, error) {
	if p1.choice == None || p2.choice == None {
		return nil, errors.New("invalid choices!")
	}

	if p1.choice == p2.choice {
		return &NoPlayer, nil
	}

	if p1.choice.beats(p2.choice) {
		return p1, nil
	} else if p2.choice.beats(p1.choice) {
		return p2, nil
	} else {
		return &NoPlayer, nil
	}
}
