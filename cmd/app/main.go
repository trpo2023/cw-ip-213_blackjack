package main

import (
	"course/internal/blackjack"
	"fmt"
)

func main() {
	bj := blackjack.NewBlackjack(blackjack.Config{})
	err := bj.Run()

	if err != nil {
		fmt.Println("Error when running blackjack game:", err)
	}
}
