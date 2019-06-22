package blackjack

import (
	"fmt"
)

const Croupier = "Croupier"
const Draw = "Draw"
const MaxScore = 42
const MinPlayers = 1 // Croupier is not counted in here bc is always playing. No croupier no party.
const MaxPlayers = 6

type game struct {
	players map[string][]string
	cards   map[string]int
}

type PlayableBlackJackSim interface {
	PlayerPoints(player string) int
	CurrentStatus() map[string][]string
	Hit(player string)
	Winner() string
}

func New(players []string) PlayableBlackJackSim {
	cards := map[string]int{
		"AceOne": 1,
		"Ace":    11,
		"Two":    2,
		"Three":  3,
		"Four":   4,
		"Five":   5,
		"Six":    6,
		"Seven":  7,
		"Eight":  8,
		"Nine":   9,
		"Ten":    10,
		"Jack":   10,
		"Queen":  10,
		"King":   10,
	}

	// FIXME FIXME FIXME: Deal random cards to the given players and the croupier
	return &game{
		players: map[string][]string{"PLAYER ONE":{"ACE", "JACK"}, "PLAYER TWO":{"ACE", "QUEEN"}, Croupier:{"KING"}},
		cards:   cards,
	}
}

func (g *game) PlayerPoints(player string) int {
	return g.calculatePoints(g.players[player])
}

func (g *game) CurrentStatus() map[string][]string {
	return g.players
}

func (g *game) Hit(player string) {
	card := "ACE" // FIXME FIXME FIXME: get a random card!
	actualPoints := g.calculatePoints(g.players[player])
	if actualPoints+g.cards[card] > 42 {
		var err error = nil
		var position = -1
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
}

func (g *game) Winner() string {
	totals := make(map[string]int, len(g.players))

	for name, cards := range g.players {
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

	if len(winners) == 1 {
		return winners[0]
	}

	return "Draw"
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
