/*
  This package implements the logic of the blackjack game.
*/

package blackjack

import (
	"course/internal/deck"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Player struct {
	Cards  []deck.Card
	Money  int
	Bet    int
	Name   string
	Bot    bool
	Dealer bool
	Pass   bool
}

type Blackjack struct {
	// Deck
	Deck []deck.Card
	// Index of the map to be taken next
	NextDeckCardIndex int
	// Deck Settings
	DeckOptions deck.NewDeckOptions
	// Player
	Players []Player
	// Number of bots
	BotsNumber int
	// Dealer
	Dealer *Player
	// Have the starting cards been distributed
	IsStartingCardsDistributed bool
	// Index of the current move
	CurrentTurnIndex int
	// Current player
	CurrentPlayer *Player
}

type Config struct {
	PlayersStartingMoney int
	BotsNumber           int
}

func NewBlackjack(cfg Config) (*Blackjack, error) {
	if cfg.PlayersStartingMoney <= 0 {
		return nil, fmt.Errorf("player starting money is negative or equal to 0: %d", cfg.PlayersStartingMoney)
	}

	if cfg.BotsNumber < 0 {
		return nil, fmt.Errorf("bots number less then 0: %d", cfg.BotsNumber)
	}

	if cfg.BotsNumber > 9 {
		return nil, fmt.Errorf("bots number greater than 9: %d", cfg.BotsNumber)
	}

	rand.Seed(time.Now().UnixNano())

	deckOptions := deck.NewDeckOptions{}

	playerUser := Player{
		Name:  DefaultPlayerNames[rand.Intn(len(DefaultPlayerNames))],
		Money: cfg.PlayersStartingMoney,
	}

	players := make([]Player, 0)
	players = append(players, playerUser)

	for i := 0; i < cfg.BotsNumber; i++ {
		players = append(players, Player{
			Name:  DefaultPlayerNames[rand.Intn(len(DefaultPlayerNames))],
			Money: cfg.PlayersStartingMoney,
			Bot:   true,
		})
	}

	playerCurrentTurnIndex := 0
	// 1 is user
	playersNumber := cfg.BotsNumber + 1

	if playersNumber > 1 {
		playerCurrentTurnIndex = rand.Intn(cfg.BotsNumber)
	}

	currentPlayer := &(players[playerCurrentTurnIndex])

	dealer := &Player{
		Money:  math.MaxInt,
		Name:   "Dealer",
		Dealer: true,
	}

	return &Blackjack{
		Deck:                       deck.NewDeck(deckOptions),
		NextDeckCardIndex:          0,
		DeckOptions:                deckOptions,
		Players:                    players,
		BotsNumber:                 cfg.BotsNumber,
		Dealer:                     dealer,
		IsStartingCardsDistributed: false,
		CurrentTurnIndex:           1,
		CurrentPlayer:              currentPlayer,
	}, nil
}

func (bj *Blackjack) Run() error {
	bj.printWelcome()

	return nil
}

func (bj *Blackjack) printWelcome() {
	fmt.Println(" ___   ___   ___   ___   ___ \n |A  | |K  | |Q  | |J  | |10 |\n |(`)| |(`)| |(`)| |(`)| |(`)|\n |_\\_| |_\\_| |_\\_| |_\\_| |_\\_|")
	fmt.Println("--- Welcome in the Blackjack Game ---")
	fmt.Println("-------------------------------------")
}
