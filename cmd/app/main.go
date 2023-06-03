package main

import (
	"course/internal/blackjack"
	"fmt"
)

func main() {
	bj, err := blackjack.NewBlackjack(blackjack.Config{
		PlayersStartingMoney: 100,
		BotsNumber:           3,
		Username:             "Arasaki",
	})
	if err != nil {
		fmt.Println("error when creating blackjack game:", err)
		return
	}

	err = bj.Run()
	if err != nil {
		fmt.Println("error when running blackjack game:", err)
		return
	}
}
