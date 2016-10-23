package euchre

import "fmt"

type Phase string

type Hand struct {
	players []Player // in playing order. 0 is the dealer.
	hands   []Cards  // the hands of each player
	kitty   Card     // the top card dealt to the kitty
	deck    *Deck

	called bool
	trump  Suit

	turn int // the index of the next player's turn
}

// NewHand creates a new hand. The first player is the dealer, the third is the
// dealers teammate.
func NewHand(players []Player, deck *Deck) *Hand {
	if len(players) != 4 {
		panic(fmt.Sprintf("players has %d people", len(players)))
	}

	if len(deck.cards) != 24 {
		panic("expected 24 cards")
	}

	h := &Hand{
		players: players,
		deck:    deck,
		turn:    1,
	}

	h.deal()
	return h
}

// Trump returns the current trump suit or false if it's not yet been decided.
func (h *Hand) Trump() (Suit, bool) {
	return h.trump, h.called
}

// Cards returns the cards left in the players hand.
func (h *Hand) Cards(player Player) Cards {
	for i, p := range h.players {
		if p == player {
			return h.hands[i]
		}
	}
	panic(fmt.Sprintf("unknown player %s", player))
}

func (h *Hand) Kitty() Card {
	return h.kitty
}

// deal deals the hand.
func (h *Hand) deal() {
	// not the normal order of things, but that's fine.
	hands := make([]Cards, len(h.players))
	for i := range h.players {
		hands[i] = h.deck.DealMust(5)
	}
	h.hands = hands
	h.kitty = h.deck.DealMust(1)[0]

}
