package blackjack

type card string

const numDecks = 1

const (
	Ace    card = "Ace"
	Two    card = "Two"
	Three  card = "Three"
	Four   card = "Four"
	Five   card = "Five"
	Six    card = "Six"
	Seven  card = "Seven"
	Eight  card = "Eight"
	Nine   card = "Nine"
	Ten    card = "Ten"
	Jack   card = "Jack"
	Queen  card = "Queen"
	King   card = "King"
	AceOne card = "AceOne"
)

var cardValues = map[card]int{
	Ace:    11,
	Two:    2,
	Three:  3,
	Four:   4,
	Five:   5,
	Six:    6,
	Seven:  7,
	Eight:  8,
	Nine:   9,
	Ten:    10,
	Jack:   10,
	Queen:  10,
	King:   10,
	AceOne: 1,
}

type deck struct {
	cards []card
}

func newDeck() []card {
	cards := []card{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	for i := 1; i <= 4; i++ {
		cards = append(cards, cards...)
	}

	return cards
}

func (d *deck) getCards() []card {
	return d.cards
}
