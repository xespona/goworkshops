package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultDecksNumber = 1

type (
	deck interface {
		cards() []card
		cardValue(card) int
	}

	PlayableBlackJackSim interface {
		CurrentStatus() map[string][]card
		Hit(player string) error
		Winner() string
	}

	// game is the struct representation of our bla``sss`sss`ssd.,.,,dezfckjack game. It has a:
	// - deck, representing the type of deck we are using for this game. Who knows, maybe you want to play something fancy?
	// - playerCards, representing the actual cards hold by the different players (not including the Croupier
	// - croupierCards.
	// - gameCards, representing ALL the different cards the game has (useful when using multiple decks). The order is important
	// - nextCardToDraw, integer representing at the position of the gameCards slice for the next card to be drawn.
	blackjack struct {
		deck
		playerCards    map[string][]card
		croupierCards  []card
		gameCards      []card
		nextCardToDraw int
	}
)

func New(players []string, deck deck, requestedDecks int) PlayableBlackJackSim {
	if len(players) <= 0 || len(players) > 6 {
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

	//fmt.Printf("%v\n", cards)

	g := &blackjack{deck: deck, gameCards: cards, nextCardToDraw: 0, playerCards: make(map[string][]card)}
	g.shuffleCards()

	// deal to each player
	for _, name := range players {
		for j := 0; j < 2; j++ {
			//fmt.Println(j)
			if card, err := g.drawCard(); err == nil {
				//fmt.Printf("card drawn %v\n", string(card))
				g.playerCards[name] = append(g.playerCards[name], card)
			} else {
				//fmt.Println("WTF")
				panic("ERROR WHEN DEALING THE INITIAL CARDS TO PLAYERS...")
			}
		}
	}

	//fmt.Printf("players: %v\n", g.playerCards)

	// deal to croupier (just one)
	if card, err := g.drawCard(); err == nil {
		g.croupierCards = append(g.croupierCards, card)
	} else {
		panic("ERROR WHEN DEALING THE INITIAL CARDS TO CROUPIER...")
	}

	//fmt.Printf("croupier: %v\n", g.croupierCards)

	return g
}

func (g *blackjack) CurrentStatus() map[string][]card {
	var result map[string][]card
	result = make(map[string][]card, len(g.playerCards)+1)
	for name, cards := range g.playerCards {
		result[name] = cards
	}
	result["Croupier"] = g.croupierCards
	return result
}

func (g *blackjack) Hit(player string) error {
	//fmt.Println(g.playerCards[player])
	//fmt.Println(player, card)

	actualPoints := g.calculatePoints(g.playerCards[player])
	if card, err := g.drawCard(); err == nil {
		if actualPoints+g.cardValue(card) > 42 {
			var err error = nil
			var position int = -1
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
		return nil
	} else {
		return fmt.Errorf("%s", err.Error())
	}
}

func (g *blackjack) Winner() string {
	totals := make(map[string]int, len(g.playerCards))

	for name, cards := range g.playerCards {
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
func (g *blackjack) drawCard() (card, error) {
	if g.nextCardToDraw >= len(g.gameCards) {
		return card(""), fmt.Errorf("no more cards to deal :(")
	}

	card := g.gameCards[g.nextCardToDraw]
	g.nextCardToDraw++

	return card, nil
}
