/*
  This package implements the logic of the blackjack game.
*/

package blackjack

import (
	"course/internal/console"
	"course/internal/deck"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Player struct {
	Id     int
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
	// Max points
	MaxPoints int
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
		Id:    rand.Int(),
		Name:  DefaultPlayerNames[rand.Intn(len(DefaultPlayerNames))],
		Money: cfg.PlayersStartingMoney,
	}

	players := make([]Player, 0)
	players = append(players, playerUser)

	for i := 0; i < cfg.BotsNumber; i++ {
		players = append(players, Player{
			Id:    rand.Int(),
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
		Id:     rand.Int(),
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
		MaxPoints:                  21,
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

func (bj *Blackjack) getCardPoints(card deck.Card) (int, error) {
	switch card.Value {
	case deck.Ace:
		{
			return AceCostBig, nil
		}

	case deck.King:
		{
			return FaceCost, nil
		}

	case deck.Queen:
		{
			return FaceCost, nil
		}

	case deck.Jack:
		{
			return FaceCost, nil
		}

	default:
		{
			points, err := strconv.Atoi(card.Value)

			if err == nil {
				return points, nil
			} else {
				log.Printf("error when converting the card value to a number: %v", err)
			}

			log.Printf("\nincorrect card value: %v", card.Value)
			return 0, nil
		}
	}
}

func (bj *Blackjack) getPlayerCardsPoints(player *Player) (points int) {
	for i := range player.Cards {
		playerCard := player.Cards[i]
		pts, err := bj.getCardPoints(playerCard)

		if err != nil {
			fmt.Println("error when getting cards points:", err)
		} else {
			points += pts
		}
	}
	return
}

func (bj *Blackjack) printRoundResults() {
	dealerPoints := bj.getPlayerCardsPoints(bj.Dealer)
	dealerLost := dealerPoints > 21

	fmt.Println("\n\n\n--- Round results: ---\n")

	fmt.Printf("%s (%d points)\n\n", bj.Dealer.Name, dealerPoints)

	for i := range bj.Players {
		player := &bj.Players[i]
		playerPoints := bj.getPlayerCardsPoints(player)
		playerName := player.Name

		if player.Id == bj.CurrentPlayer.Id {
			playerName = "You"
		}

		fmt.Printf("%s (%d points): ", playerName, playerPoints)

		if (playerPoints > dealerPoints || dealerLost) && playerPoints <= bj.MaxPoints {
			fmt.Printf("You win!\n")
			player.Money += player.Bet + player.Bet
		} else if playerPoints == dealerPoints && playerPoints <= bj.MaxPoints {
			fmt.Printf("Draw\n")
			player.Money += player.Bet
		} else {
			fmt.Printf("Defeat\n")
		}
	}
}

func (bj *Blackjack) checkRound() {
	allPass := false
	allPlayersPass := false

	for i := range bj.Players {
		player := &bj.Players[i]
		if !player.Pass {
			break
		} else if i == len(bj.Players)-1 {
			allPlayersPass = true
		}
	}

	if bj.Dealer.Pass && allPlayersPass {
		allPass = true
	}
	cnsl := console.NewCmd()
	if allPass {
		bj.printRoundResults()
		lines := "--------------------------"
		fmt.Println("\n\n")
		fmt.Println(lines)
		fmt.Println("  New Round  ")
		fmt.Println(lines)
		// TODO
		//bj.resetRound()
		fmt.Println("Press enter to continue...")
		cnsl.Input()
	}
}
