package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
	case r == 11:
		return "J"
	case r == 12:
		return "Q"
	case r == 13:
		return "K"
	case r == 14:
		return "A"
	default:
		return strconv.Itoa(int(r))
	}

}

// Ranks are the allowed cards in a euchre hand - 9, 10, J, Q, K, A
var Ranks = []Rank{9, 10, 11, 12, 13, 14}

// Card represents one playing card.
type Card struct {
	Rank Rank
	Suit Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%2s %s", c.Rank, c.Suit)
}

// Deck is a set of cards.
type Deck []Card

// NewDeck returns a new Euchre deck.
func NewDeck() Deck {
	d := make([]Card, 0, len(Suits)*len(Ranks))
	for _, s := range Suits {
		for _, r := range Ranks {
			d = append(d, Card{Suit: s, Rank: r})
		}
	}
	return Deck(d)
}

// NewDeckShuffled returns a new shuffled deck.
func NewDeckShuffled() Deck {
	d := NewDeck()
	d.Shuffle()
	return d
}

func (d Deck) Shuffle() {
	for i := 0; i < len(d); i++ {
		j := rand.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}
