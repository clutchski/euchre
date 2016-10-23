package euchre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckDeal(t *testing.T) {
	assert := assert.New(t)

	cards := Cards{
		Card{Rank: Nine, Suit: Hearts},
		Card{Rank: Ten, Suit: Hearts},
		Card{Rank: Jack, Suit: Hearts},
	}

	deck := NewDeck(cards)

	card, ok := deck.Deal()
	assert.True(ok)
	assert.Equal(card, cards[0])
	assert.Equal(deck.DealMust(), cards[1])
	assert.Equal(deck.DealMust(), cards[2])
	card, ok = deck.Deal()
	assert.False(ok)

	assert.Panics(func() {
		deck.DealMust()
	})
}

func TestEuchreDeck(t *testing.T) {
	decks := []*Deck{NewDeckEuchre(), NewDeckEuchreShuffled()}
	for _, d := range decks {
		for i := 0; i < 24; i++ {
			d.DealMust()
		}
		// FIXME[matt] someday check the cards are right
		assert.Panics(t, func() {
			d.DealMust()
		}, "dealing 25 cards should panic")
	}

}
