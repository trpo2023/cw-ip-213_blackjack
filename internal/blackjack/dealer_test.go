package blackjack

import (
	"course/internal/deck"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewDealer(t *testing.T) {
	testCases := []struct {
		name  string
		check func(d *Dealer, err error)
	}{
		{
			name: "Ok",
			check: func(d *Dealer, err error) {
				require.NoError(t, err)
				require.NotNil(t, d)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			d, err := newDealer()
			tc.check(d, err)
		})
	}
}

func TestDealer_getPoints(t *testing.T) {
	testCases := []struct {
		name       string
		buildStubs func(d *Dealer)
		check      func(d *Dealer, p int, err error)
	}{
		{
			name:       "Zero Points",
			buildStubs: func(d *Dealer) {},
			check: func(d *Dealer, p int, err error) {
				require.NoError(t, err)
				require.Equal(t, 0, p)
			},
		},
		{
			name: "Some Points",
			buildStubs: func(d *Dealer) {
				d.Cards = []*deck.Card{
					{
						Suit:  deck.Heart,
						Value: deck.Queen,
					},
					{
						Suit:  deck.Spade,
						Value: deck.King,
					},
				}
			},
			check: func(d *Dealer, p int, err error) {
				require.NoError(t, err)
				require.Equal(t, 20, p)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			d, err := newDealer()
			require.NoError(t, err)

			tc.buildStubs(d)

			p, err := d.getPoints()
			tc.check(d, p, err)
		})
	}
}
