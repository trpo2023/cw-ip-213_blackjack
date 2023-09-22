package blackjack

import (
	"course/internal/deck"
	"course/pkg/random"
	"github.com/stretchr/testify/require"
	"testing"
)

func getValidTestCfg() Config {
	return Config{
		PlayersStartingMoney: 1000,
		BotsNumber:           2,
		Username:             "Alex",
	}
}

func getValidTestPlayer() *Player {
	cfg := getValidTestCfg()

	return &Player{
		Id:      random.RandString(10),
		Cards:   []*deck.Card{},
		Money:   cfg.PlayersStartingMoney,
		Bet:     0,
		Name:    cfg.Username,
		Bot:     false,
		IsSaved: false,
		IsLost:  false,
	}
}

func TestNewBlackjack(t *testing.T) {
	cfg := getValidTestCfg()

	testCases := []struct {
		name   string
		config Config
		check  func(b *Blackjack, err error, c Config)
	}{
		{
			name: "Ok",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           cfg.BotsNumber,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.NoError(t, err)
				require.NotNil(t, c)
			},
		},
		{
			name: "Empty Players Starting Money",
			config: Config{
				PlayersStartingMoney: 0,
				BotsNumber:           cfg.BotsNumber,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.EqualError(t, err, ErrInvalidPlayersStartingMoney.Error())
				require.Nil(t, b)
			},
		},
		{
			name: "Negative Players Starting Money",
			config: Config{
				PlayersStartingMoney: -100,
				BotsNumber:           cfg.BotsNumber,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.EqualError(t, err, ErrInvalidPlayersStartingMoney.Error())
				require.Nil(t, b)
			},
		},
		{
			name: "Negative Bots Number",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           -2,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.EqualError(t, err, ErrBotsNumberLessThan.Error())
				require.Nil(t, b)
			},
		},
		{
			name: "Greater Bots Number",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           10,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.EqualError(t, err, ErrBotsNumberGreaterThan.Error())
				require.Nil(t, b)
			},
		},
		{
			name: "Min Bots Number",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           MinBotsNumber,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.NoError(t, err)
				require.NotNil(t, b)
			},
		},
		{
			name: "Max Bots Number",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           MaxBotsNumber,
				Username:             cfg.Username,
			},
			check: func(b *Blackjack, err error, c Config) {
				require.NoError(t, err)
				require.NotNil(t, b)
			},
		},
		{
			name: "Empty Username",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           MaxBotsNumber,
				Username:             "",
			},
			check: func(b *Blackjack, err error, c Config) {
				require.EqualError(t, err, ErrEmptyUsername.Error())
				require.Nil(t, b)
			},
		},
		{
			name:   "Full Invalid Config",
			config: Config{},
			check: func(b *Blackjack, err error, c Config) {
				require.Error(t, err)
				require.Nil(t, b)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			b, err := NewBlackjack(tc.config)
			tc.check(b, err, tc.config)
		})
	}
}

func TestBlackjack_GiveCardsToPlayer(t *testing.T) {
	cfg := getValidTestCfg()

	testCases := []struct {
		name   string
		config Config
		player *Player
		number int
		check  func(card *deck.Card, err error, player *Player, number int)
	}{
		{
			name:   "Ok",
			config: cfg,
			player: getValidTestPlayer(),
			number: 1,
			check: func(card *deck.Card, err error, player *Player, number int) {
				require.NoError(t, err)
				require.NotNil(t, card)
				require.Equal(t, number, len(player.Cards))
			},
		},
		{
			name:   "Invalid Player",
			config: cfg,
			player: nil,
			number: 1,
			check: func(card *deck.Card, err error, player *Player, number int) {
				require.Error(t, err)
				require.Nil(t, card)
				require.Nil(t, player)
			},
		},
		{
			name:   "Big Number",
			config: cfg,
			player: getValidTestPlayer(),
			number: 1000000,
			check: func(card *deck.Card, err error, player *Player, number int) {
				require.Error(t, err)
				require.Nil(t, card)
				require.NotNil(t, player)
				require.Equal(t, 0, len(player.Cards))
			},
		},
		{
			name:   "Negative Number",
			config: cfg,
			player: getValidTestPlayer(),
			number: -10,
			check: func(card *deck.Card, err error, player *Player, number int) {
				require.Error(t, err)
				require.Nil(t, card)
				require.NotNil(t, player)
				require.Equal(t, 0, len(player.Cards))
			},
		},
		{
			name:   "Cards Number Equals Deck Cards Number",
			config: cfg,
			player: getValidTestPlayer(),
			number: 52,
			check: func(card *deck.Card, err error, player *Player, number int) {
				require.NoError(t, err)
				require.NotNil(t, card)
				require.Equal(t, 52, len(player.Cards))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					tc.check(nil, err.(error), tc.player, tc.number)
				}
			}()

			b, err := NewBlackjack(tc.config)
			require.NoError(t, err)

			card, err := b.giveCardToPlayer(tc.player, tc.number)
			tc.check(card, err, tc.player, tc.number)
		})
	}
}

func TestBlackjack_Run(t *testing.T) {
	cfg := getValidTestCfg()

	testCases := []struct {
		name   string
		config Config
		number int
		check  func(card *deck.Card, b *Blackjack, err error, number int)
	}{
		{
			name:   "Ok",
			config: cfg,
			number: 1,
			check: func(card *deck.Card, b *Blackjack, err error, number int) {
				require.NoError(t, err)
				require.NotNil(t, card)
				require.Equal(t, number, len(b.dealer.Cards))
			},
		},
		{
			name:   "Big Number",
			config: cfg,
			number: 1000000,
			check: func(card *deck.Card, b *Blackjack, err error, number int) {
				require.Error(t, err)
				require.Nil(t, card)
				require.Equal(t, 0, len(b.dealer.Cards))
			},
		},
		{
			name:   "Negative Number",
			config: cfg,
			number: -10,
			check: func(card *deck.Card, b *Blackjack, err error, number int) {
				require.Error(t, err)
				require.Nil(t, card)
				require.Equal(t, 0, len(b.dealer.Cards))
			},
		},
		{
			name:   "Cards Number Equals Deck Cards Number",
			config: cfg,
			number: 52,
			check: func(card *deck.Card, b *Blackjack, err error, number int) {
				require.NoError(t, err)
				require.NotNil(t, card)
				require.Equal(t, 52, len(b.dealer.Cards))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					tc.check(nil, nil, err.(error), tc.number)
				}
			}()

			b, err := NewBlackjack(tc.config)
			require.NoError(t, err)

			card, err := b.giveCardToDealer(tc.number)
			tc.check(card, b, err, tc.number)
		})
	}
}

func TestBlackjack_giveCardsToAll(t *testing.T) {
	cfg := getValidTestCfg()

	testCases := []struct {
		name   string
		config Config
		number int
		check  func(b *Blackjack, err error, number int)
	}{
		{
			name: "Ok",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           2,
				Username:             cfg.Username,
			},
			number: 1,
			check: func(b *Blackjack, err error, number int) {
				require.NoError(t, err)
				require.Equal(t, 48, b.getActualDeckCardsCount())
			},
		},
		{
			name: "Max Bots Number",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           MaxBotsNumber,
				Username:             cfg.Username,
			},
			number: 1,
			check: func(b *Blackjack, err error, number int) {
				require.NoError(t, err)
				require.Equal(t, 41, b.getActualDeckCardsCount())
			},
		},
		{
			name: "Min Bots Number",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           MinBotsNumber,
				Username:             cfg.Username,
			},
			number: 1,
			check: func(b *Blackjack, err error, number int) {
				require.NoError(t, err)
				require.Equal(t, 49, b.getActualDeckCardsCount())
			},
		},
		{
			name: "All Cards",
			config: Config{
				PlayersStartingMoney: cfg.PlayersStartingMoney,
				BotsNumber:           2,
				Username:             cfg.Username,
			},
			number: 13,
			check: func(b *Blackjack, err error, number int) {
				require.NoError(t, err)
				require.Equal(t, 0, b.getActualDeckCardsCount())
			},
		},
		{
			name:   "Big Number",
			config: cfg,
			number: 100000,
			check: func(b *Blackjack, err error, number int) {
				require.Error(t, err)
				require.Equal(t, 52, b.getActualDeckCardsCount())
			},
		},
		{
			name:   "Negative Number",
			config: cfg,
			number: -100000,
			check: func(b *Blackjack, err error, number int) {
				require.Error(t, err)
				require.Equal(t, 52, b.getActualDeckCardsCount())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					tc.check(nil, err.(error), tc.number)
				}
			}()

			b, err := NewBlackjack(tc.config)
			require.NoError(t, err)

			err = b.giveCardsToAll(tc.number)
			tc.check(b, err, tc.number)
		})
	}
}
