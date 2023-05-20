package main

import (
	"course/internal/blackjack"
	"fmt"
)

func main() {
	bj, err := blackjack.NewBlackjack(blackjack.Config{})
	if err != nil {
		fmt.Println("error when creating blackjack game:", err)
	}

	err = bj.Run()
	if err != nil {
		fmt.Println("error when running blackjack game:", err)
	}
}
