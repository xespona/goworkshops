package main

import (
	"fmt"
	"github.com/xespona/goworkshops/blackjack"
)

func main() {

	game := blackjack.New(map[string][]string{
		"Crupier": {"Ace"},
		"Uno": {"Ace", "Eight"},
	})

	game.Hit("Uno", "Eight")
	game.Hit("Uno", "King")
	game.Hit("Crupier", "Ace")
	game.Hit("Crupier", "Ace")
	game.Hit("Crupier", "Nine")

	fmt.Println(game.Winner())

}
