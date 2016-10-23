package euchre

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

// Suit represents the suits of cards.
type Suit string

const (
	Hearts   Suit = "H"
	Clubs         = "C"
	Diamonds      = "D"
	Spades        = "S"
)

// Suits are the suits.
var Suits = []Suit{Hearts, Clubs, Diamonds, Spades}

// Rank is the value of the card. 9, 10, J, etc.
type Rank int

func (r Rank) String() string {
	switch {
	case r == Jack:
		return "J"
	case r == Queen:
		return "Q"
	case r == King:
		return "K"
	case r == Ace:
		return "A"
	default:
		return strconv.Itoa(int(r))
	}

}

// The cards used in euchre
const (
	Nine  Rank = 9
	Ten        = 10
	Jack       = 11
	Queen      = 12
	King       = 13
	Ace        = 14
)

// Ranks are the allowed cards in a euchre hand.
var Ranks = []Rank{Nine, Ten, Jack, Queen, King, Ace}

// Card represents one playing card.
type Card struct {
	Rank Rank
	Suit Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%2s %s", c.Rank, c.Suit)
}

// Cards is a set of cards.
type Cards []Card

func (d Cards) String() string {
	var b bytes.Buffer
	for i, c := range d {
		b.WriteString(c.String())
		if i < len(d)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// Deck is a deck of cards.
type Deck struct {
	mu    sync.Mutex
	cards Cards
	idx   int
}

// NewDeck is a new deck with the given cards.
func NewDeck(cards Cards) *Deck {
	return &Deck{
		cards: cards,
		idx:   0,
	}
}

// NewDeckEuchre returns a new Euchre deck.
func NewDeckEuchre() *Deck {
	cards := make(Cards, 0, len(Suits)*len(Ranks))
	for _, s := range Suits {
		for _, r := range Ranks {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}
	return NewDeck(cards)
}

// NewDeckEuchreShuffled returns a new shuffled euchre deck.
func NewDeckEuchreShuffled() *Deck {
	d := NewDeckEuchre()
	d.shuffle()
	return d
}

// Deal a card. If there are no cards left return false.
func (d *Deck) Deal() (c Card, ok bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.idx < len(d.cards) {
		c = d.cards[d.idx]
		ok = true
		d.idx++
	}

	return c, ok
}

// DealMust deals a card and panics if there are none left.
func (d *Deck) DealMust() Card {
	card, ok := d.Deal()
	if ok {
		return card
	}
	panic("dealt past the end of the deck")
}

func (d *Deck) shuffle() {
	d.mu.Lock()
	defer d.mu.Unlock()

	order := rand.Perm(len(d.cards))
	for i, j := range order {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
