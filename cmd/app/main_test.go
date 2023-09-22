package main

import (
	"course/internal/blackjack"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_main(t *testing.T) {
	bj, err := blackjack.NewBlackjack(blackjack.Config{
		PlayersStartingMoney: 100,
		BotsNumber:           3,
		Username:             "Arasaki",
	})

	require.NoError(t, err)
	require.NotNil(t, bj)
}
