package deck

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testCases := []struct {
		name   string
		config NewDeckOptions
		check  func(cards []*Card, err error, config NewDeckOptions)
	}{
		{
			name:   "Default Config",
			config: NewDeckOptions{},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 52)
			},
		},
		{
			name: "Custom Suit Count",
			config: NewDeckOptions{
				SuitCardsCount: 12,
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 48)
			},
		},
		{
			name: "Custom Jokers Count",
			config: NewDeckOptions{
				JokersNumber: 2,
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 54)
			},
		},
		{
			name: "Custom Jokers and Suit Count",
			config: NewDeckOptions{
				SuitCardsCount: 12,
				JokersNumber:   2,
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 50)
			},
		},
		{
			name: "No Shuffle",
			config: NewDeckOptions{
				NoShuffle: true,
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 52)
			},
		},
		{
			name: "Custom Suits",
			config: NewDeckOptions{
				Suits: []CardSuit{Spade, Heart},
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 26)
			},
		},
		{
			name: "Custom Card Values Order",
			config: NewDeckOptions{
				CardValuesOrder: []CardValue{Jack, Number, Ace, King, Queen},
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 52)
			},
		},
		{
			name: "Custom Card Number Value Start",
			config: NewDeckOptions{
				CardNumberValueStart: 5,
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 52)
			},
		},
		{
			name: "Custom Decks Number",
			config: NewDeckOptions{
				DecksNumber: 3,
			},
			check: func(cards []*Card, err error, config NewDeckOptions) {
				require.NoError(t, err)
				require.NotNil(t, cards)
				require.Len(t, cards, 156)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			d, err := NewDeck(tc.config)
			tc.check(d, err, tc.config)
		})
	}
}
