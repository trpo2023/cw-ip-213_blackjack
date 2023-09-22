/*
  This package implements working with the deck
*/

package deck

import (
	"course/pkg/random"
	"strconv"
)

type CardSuit string

// Card Suits
const (
	Spade     CardSuit = "spade"
	Heart     CardSuit = "heart"
	Clover    CardSuit = "clover"
	Diamond   CardSuit = "diamond"
	TrumpCard CardSuit = "trumpCard"
)

type CardValue string

// Card Values
const (
	Ace    CardValue = "ace"
	King   CardValue = "king"
	Queen  CardValue = "queen"
	Jack   CardValue = "jack"
	Number CardValue = "number"
	Joker  CardValue = "joker"
)

type Card struct {
	Suit CardSuit
	// (A, K, Q ...)
	Value CardValue
}

type NewDeckOptions struct {
	// How many cards of each suit should be in the deck
	SuitCardsCount int
	// Jokers count in deck
	JokersNumber int
	// Deck sorting function
	ShuffleFn func(deck []*Card)
	// Whether to shuffle the deck when creating. If the ShuffleFn function is passed, then shuffling will not work by default
	NoShuffle bool
	// Available suits in the deck
	Suits []CardSuit
	// Order of cards of suits
	CardValuesOrder []CardValue
	// Initial value for cards with a numeric value
	CardNumberValueStart int
	// Number of decks
	DecksNumber int
}

// NewDeck
// Creating a new deck.
func NewDeck(options NewDeckOptions) ([]*Card, error) {
	// --- Initializing the settings for creating a new deck
	suitCardsCount := options.SuitCardsCount
	jokersNumber := options.JokersNumber
	shuffleFn := options.ShuffleFn
	suits := options.Suits
	shuffle := !options.NoShuffle
	cardValuesOrder := options.CardValuesOrder
	cardNumberValueStart := options.CardNumberValueStart
	decksNumber := options.DecksNumber

	if suitCardsCount <= 0 {
		suitCardsCount = 13
	}

	if len(suits) == 0 {
		suits = []CardSuit{Spade, Heart, Clover, Diamond}
	}

	if len(cardValuesOrder) == 0 {
		cardValuesOrder = []CardValue{Ace, King, Queen, Jack, Number}
	}

	if cardNumberValueStart == 0 {
		cardNumberValueStart = 2
	}

	if decksNumber == 0 {
		decksNumber = 1
	}
	// -------------------------------------------

	deckSize := (len(suits)*suitCardsCount + jokersNumber) * decksNumber
	deck := make([]*Card, deckSize)

	// Index of the currently created map
	curCardIndex := 0

	for curDeckNumber := 1; curDeckNumber <= decksNumber; curDeckNumber++ {
		for i := range suits {
			cardValuesOrderIndex := 0
			cardNumberValue := cardNumberValueStart

			for j := 0; j < suitCardsCount; j++ {
				suit := suits[i]
				var value CardValue

				cardValueOrder := cardValuesOrder[cardValuesOrderIndex]

				if cardValueOrder == Number {
					value = CardValue(strconv.Itoa(cardNumberValue))
					cardNumberValue++
				} else {
					value = cardValueOrder
					cardValuesOrderIndex++
				}

				card, err := newCard(suit, value)
				if err != nil {
					return nil, err
				}
				deck[curCardIndex] = card

				curCardIndex++
			}
		}

		for i := 0; i < jokersNumber; i++ {
			card, err := newCard(TrumpCard, Joker)
			if err != nil {
				return nil, err
			}
			deck[curCardIndex] = card
			curCardIndex++
		}
	}

	if shuffle && shuffleFn == nil {
		// Sattolo 's Algorithm
		i := len(deck)
		for i > 1 {
			i = i - 1
			j := random.RandInt(0, i)
			deck[j], deck[i] = deck[i], deck[j]
		}
	} else if shuffleFn != nil {
		shuffleFn(deck)
	}

	return deck, nil
}

func newCard(suit CardSuit, value CardValue) (*Card, error) {
	return &Card{
		Suit:  suit,
		Value: value,
	}, nil
}
