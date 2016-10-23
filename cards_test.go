package euchre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckDealMust(t *testing.T) {
	assert := assert.New(t)

	cards := Cards{
		heart9,
		heart10,
		heartJ,
	}

	deck := NewDeck(cards)
	assert.Equal(deck.DealMust(3), cards)
	assert.Panics(func() {
		deck.DealMust(1)
	}, "dealing must should panic if there isn't enough")

}

func TestDeckDeal(t *testing.T) {
	assert := assert.New(t)
	deck := NewDeck(Cards{heart9, heart10, heartJ})

	assert.Equal(deck.Deal(2), Cards{heart9, heart10})
	assert.Equal(deck.Deal(2), Cards{heartJ})
	assert.Empty(deck.Deal(2))
}

func TestEuchreDeck(t *testing.T) {
	decks := []*Deck{NewEuchreDeck(), NewShuffledEuchreDeck()}
	for _, d := range decks {
		out := d.DealMust(24)
		assert.Len(t, out, 24)
		// FIXME[matt] someday check the cards are right
		assert.Panics(t, func() {
			d.DealMust(1)
		}, "dealing 25 cards should panic")
	}

}
