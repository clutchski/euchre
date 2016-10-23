package euchre

// makes some unit tests easier
var (
	club9  = Card{Nine, Clubs}
	club10 = Card{Ten, Clubs}
	clubJ  = Card{Jack, Clubs}
	clubQ  = Card{Queen, Clubs}
	clubK  = Card{King, Clubs}
	clubA  = Card{Ace, Clubs}

	diam9  = Card{Nine, Diamonds}
	diam10 = Card{Ten, Diamonds}
	diamJ  = Card{Jack, Diamonds}
	diamQ  = Card{Queen, Diamonds}
	diamK  = Card{King, Diamonds}
	diamA  = Card{Ace, Diamonds}

	heart9  = Card{Nine, Hearts}
	heart10 = Card{Ten, Hearts}
	heartJ  = Card{Jack, Hearts}
	heartQ  = Card{Queen, Hearts}
	heartK  = Card{King, Hearts}
	heartA  = Card{Ace, Hearts}

	spade9  = Card{Nine, Spades}
	spade10 = Card{Ten, Spades}
	spadeJ  = Card{Jack, Spades}
	spadeQ  = Card{Queen, Spades}
	spadeK  = Card{King, Spades}
	spadeA  = Card{Ace, Spades}
)

// getTestDeck returns all the cards in order of suit and rank
func getTestDeck() *Deck {
	return NewEuchreDeck()
}
