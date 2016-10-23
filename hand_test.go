package euchre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	p := []Player{"a1", "b1", "a2", "b2"}
	hand := NewHand(p, getTestDeck())
	_, ok := hand.Trump()
	assert.False(t, ok)

}

func TestDeal(t *testing.T) {
	// assert we can deal in order
	assert := assert.New(t)
	p := []Player{"a1", "b1", "a2", "b2"}
	hand := NewHand(p, getTestDeck())
	assert.Equal(hand.Cards(p[0]), Cards{club9, club10, clubJ, clubQ, clubK})
	assert.Equal(hand.Cards(p[1]), Cards{clubA, diam9, diam10, diamJ, diamQ})
	assert.Equal(hand.Cards(p[2]), Cards{diamK, diamA, heart9, heart10, heartJ})
	assert.Equal(hand.Cards(p[3]), Cards{heartQ, heartK, heartA, spade9, spade10})
	assert.Equal(hand.Kitty(), spadeJ)
}
