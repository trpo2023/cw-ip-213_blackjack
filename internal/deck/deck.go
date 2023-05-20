/*
  This package implements working with the deck
*/

package deck

import (
	"math/rand"
	"strconv"
	"time"
)

// Suits
const (
	// Spade
	Spade = "spade"
	// Heart
	Heart = "heart"
	// Clover
	Clover = "clover"
	// Diamond
	Diamond = "diamond"
)

// Card Values
const (
	Ace    = "ace"
	King   = "king"
	Queen  = "queen"
	Jack   = "jack"
	Number = "number"
	Joker  = "joker"
)

type Card struct {
	// Suit
	Suit string
	// Value (A, K, Q ...)
	Value string
}

type NewDeckOptions struct {
	// How many cards of each suit should be in the deck
	SuitCardsCount int
	// Jokers count in deck
	JokersNumber int
	// Deck sorting function
	ShuffleFn func(deck *[]Card)
	// Whether to shuffle the deck when creating. If the ShuffleFn function is passed, then shuffling will not work by default
	NoShuffle bool
	// Available suits in the deck
	Suits []string
	// Order of cards of suits
	CardValuesOrder []string
	// Initial value for cards with a numeric value
	CardNumberValueStart int
	// Number of decks
	DecksNumber int
}

// NewDeck
// Creating a new deck.
func NewDeck(options NewDeckOptions) []Card {
	rand.Seed(time.Now().UnixNano())

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
		suits = []string{Spade, Heart, Clover, Diamond}
	}

	if len(cardValuesOrder) == 0 {
		cardValuesOrder = []string{Ace, King, Queen, Jack, Number}
	}

	if cardNumberValueStart == 0 {
		cardNumberValueStart = 2
	}

	if decksNumber == 0 {
		decksNumber = 1
	}
	// -------------------------------------------

	deckSize := (len(suits)*suitCardsCount + jokersNumber) * decksNumber
	deck := make([]Card, deckSize)

	// Index of the currently created map
	curCardIndex := 0

	for curDeckNumber := 1; curDeckNumber <= decksNumber; curDeckNumber++ {
		for i := range suits {
			cardValuesOrderIndex := 0
			cardNumberValue := cardNumberValueStart

			for j := 0; j < suitCardsCount; j++ {
				suit := suits[i]
				value := ""

				cardValueOrder := cardValuesOrder[cardValuesOrderIndex]

				if cardValueOrder == Number {
					value = strconv.Itoa(cardNumberValue)
					cardNumberValue++
				} else {
					value = cardValueOrder
					cardValuesOrderIndex++
				}

				deck[curCardIndex] = Card{
					Suit:  suit,
					Value: value,
				}

				curCardIndex++
			}
		}

		for i := 0; i < jokersNumber; i++ {
			deck[curCardIndex] = Card{
				Suit:  Joker,
				Value: Joker,
			}
			curCardIndex++
		}
	}

	if shuffle && shuffleFn == nil {
		// Sattolo 's Algorithm
		i := len(deck)
		for i > 1 {
			i = i - 1
			j := rand.Intn(i)
			deck[j], deck[i] = deck[i], deck[j]
		}
	} else if shuffleFn != nil {
		shuffleFn(&deck)
	}

	return deck
}
