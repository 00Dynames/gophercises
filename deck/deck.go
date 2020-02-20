package deck

import (
	"fmt"
)

// Value type
type Value int

// Suit type
type Suit int

// Enum for card value
const (
	_ Value = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Enum for card suit
const (
	Heart Suit = iota
	Diamond
	Club
	Spade
	Joker
)

// Card type
type Card struct {
	Value
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()

	}

	return fmt.Sprintf("%s of %ss", c.Value.String(), c.Suit.String())
}

// NewDeck creates a new standard deck
func NewDeck() *[]Card {
	result := []Card{}

	for j := Heart; j < 4; j++ {
		for i := Ace; i <= King; i++ {
			result = append(result, Card{i, j})
		}
	}

	return &result
}
