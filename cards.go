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
	Clubs    Suit = "C"
	Diamonds      = "D"
	Hearts        = "H"
	Spades        = "S"
)

// Suits are the suits in alphabetical order.
var Suits = []Suit{Clubs, Diamonds, Hearts, Spades}

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

func NewDeck(cards Cards) *Deck {
	return &Deck{
		cards: cards,
		idx:   0,
	}
}

// Deal will return up to the given number of cards.
func (d *Deck) Deal(n int) Cards {
	if n < 0 {
		return nil
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	length := len(d.cards)
	end := d.idx + n
	if end > length {
		end = length
	}

	cards := d.cards[d.idx:end]
	d.idx = end
	return cards
}

func (d *Deck) DealMust(n int) Cards {
	cards := d.Deal(n)
	if len(cards) < n {
		panic("can't deal that many cards")
	}
	return cards
}

func (d *Deck) shuffle() {
	d.mu.Lock()
	defer d.mu.Unlock()

	order := rand.Perm(len(d.cards))
	for i, j := range order {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// NewEuchreDeck returns the cards in a euchre deck order by suit
// and rank.
func NewEuchreDeck() *Deck {
	cards := make(Cards, 0, len(Suits)*len(Ranks))
	for _, s := range Suits {
		for _, r := range Ranks {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}
	return NewDeck(cards)
}

func NewShuffledEuchreDeck() *Deck {
	d := NewEuchreDeck()
	d.shuffle()
	return d
}
