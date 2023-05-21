package blackjack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	testUsername = "TestUser"
)

func newBlackjack(t *testing.T, cfg Config) *Blackjack {
	bj, err := NewBlackjack(Config{
		PlayersStartingMoney: cfg.PlayersStartingMoney,
		BotsNumber:           cfg.BotsNumber,
		Username:             cfg.Username,
	})

	require.NoError(t, err)
	require.NotEmpty(t, bj)

	return bj
}

func TestNewBlackjack(t *testing.T) {
	newBlackjack(t, Config{
		PlayersStartingMoney: 100,
		BotsNumber:           2,
		Username:             testUsername,
	})
}

func TestNewBlackjackErrorPlayerStartingMoney(t *testing.T) {
	bj, err := NewBlackjack(Config{
		PlayersStartingMoney: -100,
		BotsNumber:           1,
		Username:             testUsername,
	})

	require.Error(t, err)
	require.Empty(t, bj)
}

func TestNewBlackjackErrorBotsNumber(t *testing.T) {
	bj, err := NewBlackjack(Config{
		PlayersStartingMoney: 100,
		BotsNumber:           -10,
		Username:             testUsername,
	})

	require.Error(t, err)
	require.Empty(t, bj)
}

func TestNewBlackjackErrorEmptyUsername(t *testing.T) {
	bj, err := NewBlackjack(Config{
		PlayersStartingMoney: 100,
		BotsNumber:           2,
		Username:             "",
	})

	require.Error(t, err)
	require.Empty(t, bj)
}

func TestGiveCardsToAll(t *testing.T) {
	bj := newBlackjack(t, Config{
		PlayersStartingMoney: 100,
		BotsNumber:           3,
		Username:             testUsername,
	})
	err := bj.giveCardsToAll(2)

	require.NoError(t, err)
}

func TestGiveCardsToAllError(t *testing.T) {
	bj := newBlackjack(t, Config{
		PlayersStartingMoney: 100,
		BotsNumber:           3,
		Username:             testUsername,
	})
	err := bj.giveCardsToAll(100)

	require.Error(t, err)
}
