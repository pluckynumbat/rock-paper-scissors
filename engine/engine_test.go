package engine

import (
	"testing"
)

func TestPrintChoice(t *testing.T) {
	players := []Player{{"NO1", None}, {"A1", Rock}, {"B1", Paper}, {"C1", Scissors}}
	strings := []string{"None", "Rock", "Paper", "Scissors"}

	for i := range len(players) - 1 {

		have := players[i].PrintChoice()
		want := strings[i]

		if have != want {
			t.Errorf("Error in the Play function, want: %s, have: %s", want, have)
		}
	}
}
func TestPlayer1Wins(t *testing.T) {
	players1 := []Player{{"A1", Rock}, {"B1", Paper}, {"C1", Scissors}}
	players2 := []Player{{"A2", Scissors}, {"B2", Rock}, {"C2", Paper}}

	for i := range len(players1) - 1 {
		winner, err := Play(&players1[i], &players2[i])

		if err != nil {
			t.Fatalf("Error in the inputs for Play: %v, %v", players1[i].choice, players2[i].choice)
		}

		have := winner
		want := &players1[i]

		if have != want {
			t.Errorf("Error in the Play function, want: %s, have: %s", want, have)
		}
	}
}

func TestPlayer2Wins(t *testing.T) {
	players1 := []Player{{"A1", Rock}, {"B1", Paper}, {"C1", Scissors}}
	players2 := []Player{{"A2", Paper}, {"B2", Scissors}, {"C2", Rock}}

	for i := range len(players1) - 1 {
		winner, err := Play(&players1[i], &players2[i])

		if err != nil {
			t.Fatalf("Error in the inputs for Play: %v, %v", players1[i].choice, players2[i].choice)
		}

		have := winner
		want := &players2[i]

		if have != want {
			t.Errorf("Error in the Play function, want: %s, have: %s", want, have)
		}
	}
}

func TestNoPlayerWins(t *testing.T) {
	players1 := []Player{{"A1", Rock}, {"B1", Paper}, {"C1", Scissors}}
	players2 := []Player{{"A2", Rock}, {"B2", Paper}, {"C2", Scissors}}

	for i := range len(players1) - 1 {
		winner, err := Play(&players1[i], &players2[i])

		if err != nil {
			t.Fatalf("Error in the inputs for Play: %v, %v", players1[i].choice, players2[i].choice)
		}

		have := winner
		want := &NoPlayer

		if have != want {
			t.Errorf("Error in the Play function, want: %s, have: %s", want, have)
		}
	}
}

func TestChooseFixed(t *testing.T) {
	players := []Player{NoPlayer, {"A1", Rock}, {"B1", Paper}, {"C1", Scissors}}
	choices := []Choice{None, Rock, Paper, Scissors}

	for i := range len(choices) - 1 {

		player := new(Player)
		player.ChooseFixed(choices[i])
		have := player.choice

		want := players[i].choice

		if want != have {
			t.Fatalf("Error in ChooseFixed: want: %s, have: %s", want, have)
		}
	}
}

func TestChooseRandom(t *testing.T) {

	player1 := Player{Name: "P1"}
	player2 := Player{Name: "P2"}

	for range 1000 {
		player1.ChooseRandom()
		player2.ChooseRandom()

		winner, err := Play(&player1, &player2)
		if err != nil {
			t.Fatalf("Error in the inputs for Play: %v, %v", player1.choice, player2.choice)
		}

		have := winner.Name

		var want string
		if player1.choice.beats(player2.choice) {
			want = player1.Name
		} else if player2.choice.beats(player1.choice) {
			want = player2.Name
		} else {
			want = NoPlayer.Name
		}

		if want != have {
			t.Fatalf("Error in ChooseFixed: want: %s, have: %s", want, have)
		}
	}
}
