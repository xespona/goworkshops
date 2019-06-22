package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultDecksNumber = 1
const MinPlayers = 1
const MaxPlayers = 6
const Croupier = "Croupier"
const Draw = "Draw"
const MaxScore = 42

type (
	PlayableBlackJackSim interface {
		PlayerPoints(player string) int
		CurrentStatus() map[string][]card
		Hit(player string)
		Winner() string
	}

	// game is the struct representation of our bla``sss`sss`ssd.,.,,dezfckjack game. It has a:
	// - deck, representing the type of deck we are using for this game. Who knows, maybe you want to play something fancy?
	// - playerCards, representing the actual cards hold by the different players (not including the Croupier
	// - gameCards, representing ALL the different cards the game has (useful when using multiple decks). The order is important
	// - nextCardToDraw, integer representing at the position of the gameCards slice for the next card to be drawn.
	blackjack struct {
		Deck
		playerCards    map[string][]card
		gameCards      []card
		nextCardToDraw int
	}
)

// Returns the current points of a given player
func (g *blackjack) PlayerPoints(player string) int {
	return g.calculatePoints(g.playerCards[player])
}

// Returns a copy of the current game status
func (g *blackjack) CurrentStatus() map[string][]card {
	var result map[string][]card
	result = make(map[string][]card, len(g.playerCards))
	for name, cards := range g.playerCards {
		result[name] = make([]card, len(cards))
		for i, card := range cards {
			result[name][i] = card
		}
	}
	return result
}

// Deals another card to a given player
func (g *blackjack) Hit(player string) {
	actualPoints := g.calculatePoints(g.playerCards[player])
	card := g.drawCard()
	if actualPoints+g.cardValue(card) > MaxScore {
		var err error = nil
		var position int
		for err == nil {
			position, err = g.acePosition(g.playerCards[player])
			if err == nil {
				g.playerCards[player][position] = "AceOne"
			} else if card == "Ace" {
				card = "AceOne"
			}
		}
	}

	g.playerCards[player] = append(g.playerCards[player], card)
}

// Returns the winner of the game! (or draw if nobody won)
func (g *blackjack) Winner() string {
	totals := make(map[string]int, len(g.playerCards))

	for name, cards := range g.playerCards {
		totals[name] = g.calculatePoints(cards)
	}

	var winners []string

	maxPoints := 0
	for name, points := range totals {
		if points > MaxScore {
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

	return Draw
}

// Creates the game
func New(players []string, deck Deck, requestedDecks int) PlayableBlackJackSim {
	if len(players) < MinPlayers || len(players) > MaxPlayers {
		panic(fmt.Sprintf("invalid number of players: %d", len(players)))
	}

	if requestedDecks <= 0 {
		requestedDecks = defaultDecksNumber
	}

	// Get ALL the cards the game will be played with (on a normal day, this should be 52 cards - 13 * 4)
	deckCards := deck.cards()
	var cards []card
	for i := 0; i < requestedDecks; i++ {
		cards = append(cards, deckCards...)
	}

	g := &blackjack{Deck: deck, gameCards: cards, nextCardToDraw: 0, playerCards: make(map[string][]card)}
	g.shuffleCards()

	// deal to each player
	for _, name := range players {
		for j := 0; j < 2; j++ {
			card := g.drawCard()
			g.playerCards[name] = append(g.playerCards[name], card)
		}
	}

	// deal to croupier (just one)
	card := g.drawCard()
	g.playerCards[Croupier] = append(g.playerCards[Croupier], card)

	return g
}

// Shuffles the cards, in style, like a proper casino croupier
func (g *blackjack) shuffleCards() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.gameCards), func(i, j int) { g.gameCards[i], g.gameCards[j] = g.gameCards[j], g.gameCards[i] })
}

// Calculates the points of a slice of cards
func (g *blackjack) calculatePoints(cards []card) int {
	total := 0
	for _, card := range cards {
		total += g.cardValue(card)
	}

	return total
}

// Finds the first Ace occurrence in a slice of cards and returns its value. If no ace is found an error is returned.
func (g *blackjack) acePosition(cards []card) (int, error) {
	for k, card := range cards {
		if g.cardValue(card) == 11 {
			return k, nil
		}
	}

	return -1, fmt.Errorf("no aces found")
}

// draws the next card from the "deck/s" (aka gameCards). If we run out of cards, an error is returned
func (g *blackjack) drawCard() card {
	if g.nextCardToDraw >= len(g.gameCards) {
		panic("no more cards to deal? :(")
	}

	card := g.gameCards[g.nextCardToDraw]
	g.nextCardToDraw++

	return card
}
