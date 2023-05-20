/*
  This package implements the logic of the blackjack game.
*/

package blackjack

import "fmt"

type Blackjack struct{}

type Config struct{}

func NewBlackjack(cfg Config) *Blackjack {
	return &Blackjack{}
}

func (b *Blackjack) printWelcome() {
	fmt.Println(" ___   ___   ___   ___   ___ \n |A  | |K  | |Q  | |J  | |10 |\n |(`)| |(`)| |(`)| |(`)| |(`)|\n |_\\_| |_\\_| |_\\_| |_\\_| |_\\_|")
	fmt.Println("--- Welcome in the Blackjack Game ---")
	fmt.Println("-------------------------------------")
}

func (b *Blackjack) Run() error {
	b.printWelcome()

	return nil
}
