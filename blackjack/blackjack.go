package blackjack

import (
	"fmt"
)

type game struct {
	players map[string][]string
	cards map[string]int
}

type BlackjackSim interface {
	Hit(player string, card string)
	Winner() string
}


func New(initialSituation map[string][]string) BlackjackSim {
	cards := map[string]int {
		"AceOne": 1,
		"Ace": 11,
		"Two": 2,
		"Three": 3,
		"Four": 4,
		"Five": 5,
		"Six": 6,
		"Seven": 7,
		"Eight": 8,
		"Nine": 9,
		"Ten": 10,
		"Jack": 10,
		"Queen":10,
		"King":10,
	}

	return &game{
		players:initialSituation,
		cards:cards,
	}
}

func (g *game) Hit(player string, card string) {
	//fmt.Println(g.players[player])
	//fmt.Println(player, card)

	actualPoints := g.calculatePoints(g.players[player])
	if actualPoints + g.cards[card] > 42 {
		var err error = nil
		var position int = -1
		for err == nil {
			position, err = g.acePosition(g.players[player])
			if err == nil {
				g.players[player][position] = "AceOne"
			} else if card == "Ace" {
				card = "AceOne"
			}
		}
	}

	g.players[player] = append(g.players[player], card)
	//fmt.Println(g.players[player])
}

func (g *game) Winner() string {
	totals :=  make(map[string]int, len(g.players))

	for name, cards :=range g.players {
		totals[name] = g.calculatePoints(cards)
	}

	var winners []string

	maxPoints := 0

	for name, points := range totals {
		if points > 42 {
			continue
		}

		if points > maxPoints {
			winners = []string{name}
			maxPoints = points
		} else if points == maxPoints {
			winners = append(winners, name)
		}
	}

	if len(winners) != 1 {
		return "Draw"
	}


	return winners[0]
}

func (g *game) calculatePoints(cards []string) int {
	total := 0
	for _, card := range cards {
		total += g.cards[card]
	}

	return total
}

func (g *game) acePosition(cards []string) (int, error) {
	for k, card := range cards {
		if g.cards[card] == 11 {
			return k, nil
		}
	}

	return -1, fmt.Errorf("no aces found")
}