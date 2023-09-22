package blackjack

import (
	"course/internal/deck"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_getCardCost(t *testing.T) {
	testCases := []struct {
		name  string
		card  *deck.Card
		check func(cost int, err error)
	}{
		{
			name: "Ok Ace",
			card: &deck.Card{
				Suit:  deck.Clover,
				Value: deck.Ace,
			},
			check: func(cost int, err error) {
				require.NoError(t, err)
				require.Equal(t, int(AceCostBig), cost)
			},
		},
		{
			name: "Ok Face",
			card: &deck.Card{
				Suit:  deck.Heart,
				Value: deck.Jack,
			},
			check: func(cost int, err error) {
				require.NoError(t, err)
				require.Equal(t, int(FaceCost), cost)
			},
		},
		{
			name: "Ok Number",
			card: &deck.Card{
				Suit:  deck.Spade,
				Value: "6",
			},
			check: func(cost int, err error) {
				require.NoError(t, err)
				require.Equal(t, 6, cost)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cost, err := getCardCost(tc.card)
			tc.check(cost, err)
		})
	}
}
