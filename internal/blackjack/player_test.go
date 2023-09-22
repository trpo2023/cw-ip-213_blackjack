package blackjack

import (
	"course/internal/deck"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewPlayer(t *testing.T) {
	cfg := getValidTestCfg()

	testCases := []struct {
		name     string
		username string
		money    int
		bot      bool
		check    func(p *Player, err error, username string, money int, bot bool)
	}{
		{
			name:     "Success Creating User Player",
			username: cfg.Username,
			money:    cfg.PlayersStartingMoney,
			bot:      false,
			check: func(p *Player, err error, username string, money int, bot bool) {
				require.NoError(t, err)
				require.NotNil(t, p)
				require.Equal(t, username, p.Name)
				require.Equal(t, money, p.Money)
				require.Equal(t, bot, p.Bot)
			},
		},
		{
			name:     "Success Creating Bot Player",
			username: cfg.Username,
			money:    cfg.PlayersStartingMoney,
			bot:      true,
			check: func(p *Player, err error, username string, money int, bot bool) {
				require.NoError(t, err)
				require.NotNil(t, p)
				require.Equal(t, username, p.Name)
				require.Equal(t, money, p.Money)
				require.Equal(t, bot, p.Bot)
			},
		},
		{
			name:     "Invalid Username",
			username: "",
			money:    cfg.PlayersStartingMoney,
			bot:      true,
			check: func(p *Player, err error, username string, money int, bot bool) {
				require.Error(t, err)
				require.Nil(t, p)
			},
		},
		{
			name:     "Zero Money",
			username: cfg.Username,
			money:    0,
			bot:      true,
			check: func(p *Player, err error, username string, money int, bot bool) {
				require.Error(t, err)
				require.Nil(t, p)
			},
		},
		{
			name:     "Negative Money",
			username: cfg.Username,
			money:    -1000,
			bot:      true,
			check: func(p *Player, err error, username string, money int, bot bool) {
				require.Error(t, err)
				require.Nil(t, p)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			p, err := newPlayer(tc.username, tc.money, tc.bot)
			tc.check(p, err, tc.username, tc.money, tc.bot)
		})
	}
}

func TestPlayer_getPoints(t *testing.T) {
	cfg := getValidTestCfg()

	testCases := []struct {
		name       string
		username   string
		money      int
		bot        bool
		buildStubs func(d *Player)
		check      func(d *Player, p int, err error)
	}{
		{
			name:       "Zero Points",
			username:   cfg.Username,
			money:      cfg.PlayersStartingMoney,
			bot:        false,
			buildStubs: func(d *Player) {},
			check: func(d *Player, p int, err error) {
				require.NoError(t, err)
				require.Equal(t, 0, p)
			},
		},
		{
			name:     "Some Points",
			username: cfg.Username,
			money:    cfg.PlayersStartingMoney,
			bot:      false,
			buildStubs: func(d *Player) {
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
			check: func(d *Player, p int, err error) {
				require.NoError(t, err)
				require.Equal(t, 20, p)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			d, err := newPlayer(tc.username, tc.money, tc.bot)
			require.NoError(t, err)

			tc.buildStubs(d)

			p, err := d.getPoints()
			tc.check(d, p, err)
		})
	}
}
