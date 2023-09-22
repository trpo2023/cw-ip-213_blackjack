/*
  This package implements the logic of the blackjack game.
*/

package blackjack

import (
	"course/internal/console"
	"course/internal/deck"
	"course/pkg/random"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Blackjack struct {
	// Deck
	deck []*deck.Card
	// Index of the map to be taken next
	nextDeckCardIndex int
	// Deck Settings
	deckOptions deck.NewDeckOptions
	// Player
	players []*Player
	// Number of bots
	botsNumber int
	// Dealer
	dealer *Dealer
	// Have the starting cards been distributed
	isStartingCardsDistributed bool
	// Index of the current move
	currentTurnIndex int
	// Current user player id
	currentUser *Player
	// All the players were saved
	isAllPlayersSaved bool
	// All saved
	isAllSaved bool
	// Console
	console *console.Console
}

type Config struct {
	PlayersStartingMoney int
	BotsNumber           int
	Username             string
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

	// 1 is user
	playersNumber := 1 + cfg.BotsNumber
	players := make([]*Player, 0, playersNumber)

	playerUser, err := newPlayer(cfg.Username, cfg.PlayersStartingMoney, false)
	if err != nil {
		return nil, err
	}

	players = append(players, playerUser)

	for i := 0; i < cfg.BotsNumber; i++ {
		bot, err := newPlayer(DefaultPlayerNames[random.RandInt(0, len(DefaultPlayerNames))], cfg.PlayersStartingMoney, true)
		if err != nil {
			return nil, err
		}
		players = append(players, bot)
	}

	//playerCurrentTurnIndex := 0
	//
	//if playersNumber > 1 {
	//	playerCurrentTurnIndex = util.GetRandomIntn(cfg.BotsNumber)
	//}
	//
	//currentPlayer := players[playerCurrentTurnIndex]

	dealer, err := newDealer()
	if err != nil {
		return nil, err
	}

	deckOptions := deck.NewDeckOptions{}
	bjDeck, err := deck.NewDeck(deckOptions)
	if err != nil {
		return nil, err
	}

	cnsl, err := console.NewConsole()
	if err != nil {
		return nil, err
	}

	return &Blackjack{
		deck:                       bjDeck,
		nextDeckCardIndex:          0,
		deckOptions:                deckOptions,
		players:                    players,
		botsNumber:                 cfg.BotsNumber,
		dealer:                     dealer,
		isStartingCardsDistributed: false,
		currentTurnIndex:           1,
		currentUser:                playerUser,
		isAllPlayersSaved:          false,
		isAllSaved:                 false,
		console:                    cnsl,
	}, nil
}

func (bj *Blackjack) Run() error {
	bj.printWelcome()
	if err := bj.gameLoop(); err != nil {
		return err
	}

	return nil
}

func (bj *Blackjack) printWelcome() {
	fmt.Println(" ___   ___   ___   ___   ___ \n |A  | |K  | |Q  | |J  | |10 |\n |(`)| |(`)| |(`)| |(`)| |(`)|\n |_\\_| |_\\_| |_\\_| |_\\_| |_\\_|")
	fmt.Println("--- Welcome in the Blackjack Game ---")
	fmt.Println("-------------------------------------")
}

func (bj *Blackjack) getParticipantCardsPoints(player *Player) (int, error) {
	points := 0

	for i := range player.Cards {
		playerCard := player.Cards[i]
		pts, err := getCardCost(playerCard)

		if err != nil {
			return -1, err
		} else {
			points += pts
		}
	}
	return points, nil
}

func (bj *Blackjack) printRoundResults() error {
	dealerPoints, err := bj.dealer.getPoints()
	if err != nil {
		return err
	}

	fmt.Println("\n\n\n--- Round results: ---")
	fmt.Printf("\n%s (%d points)", "dealer", dealerPoints)
	fmt.Println("\nDealer cards:")
	err = bj.dealer.printAllCards()
	if err != nil {
		return err
	}
	fmt.Printf("\n\n\n")

	dealerLost := dealerPoints > 21

	for _, player := range bj.players {
		playerPoints, err := player.getPoints()
		if err != nil {
			return err
		}

		playerName := player.Name

		if player.Id == bj.currentUser.Id {
			playerName = "You"
		}

		fmt.Printf("%s (%d points): ", playerName, playerPoints)

		if (playerPoints > dealerPoints || dealerLost) && playerPoints <= MaxPoints {
			fmt.Printf("Win!\n")
			player.Money += player.Bet + player.Bet
		} else if playerPoints == dealerPoints && playerPoints <= MaxPoints {
			fmt.Printf("Draw\n")
			player.Money += player.Bet
		} else {
			fmt.Printf("Defeat\n")
		}
	}

	return nil
}

func (bj *Blackjack) resetRound() error {
	for _, player := range bj.players {
		player.resetRound()
	}

	bj.dealer.resetRound()

	bjDeck, err := deck.NewDeck(bj.deckOptions)
	if err != nil {
		return err
	}
	bj.deck = bjDeck
	bj.nextDeckCardIndex = 0

	bj.isStartingCardsDistributed = false
	bj.currentTurnIndex = 0

	for _, player := range bj.players {
		player.checkIsLost()
	}

	return nil
}

func (bj *Blackjack) checkAllPlayersSaved() {
	if !bj.isAllPlayersSaved {
		for i, player := range bj.players {
			if !player.IsSaved {
				break
			} else if i == len(bj.players)-1 {
				bj.isAllPlayersSaved = true
			}
		}
	}
}

func (bj *Blackjack) checkAllSaved() {
	if bj.dealer.IsSaved && bj.isAllPlayersSaved {
		bj.isAllSaved = true
	}
}

func (bj *Blackjack) checkRound() error {
	bj.checkAllPlayersSaved()
	bj.checkAllSaved()

	if bj.isAllSaved {
		err := bj.printRoundResults()
		if err != nil {
			return err
		}

		err = bj.resetRound()
		if err != nil {
			return err
		}

		fmt.Println("\n\nPress enter to continue...")

		bj.console.Input()

		lines := "--------------------------"
		fmt.Printf("\n\n\n")
		fmt.Println(lines)
		fmt.Println("  New Round  ")
		fmt.Println(lines)
	}

	return nil
}

func (bj *Blackjack) betMakerBot(bot *Player) {
	if bot.IsLost {
		return
	}

	bet := random.RandInt(0, bot.Money)

	fmt.Printf("\n\nBot %s makes a bet...\n", bot.Name)
	time.Sleep(Delay)

	fmt.Printf("An insert of %d coins was made", bet)

	bot.Bet = bet
	bot.Money -= bet
}

func (bj *Blackjack) betMakerPlayer(player *Player) {
	fmt.Printf("\nMake your bet (you have %d c.). %s - Exit.", player.Money, ActionExit)
	userInput := ""

	for userInput == "" {
		fmt.Printf("\n>> ")
		userInput = bj.console.Input()

		if userInput == string(ActionExit) {
			fmt.Printf("\nWe are waiting for you again!")
			os.Exit(0)
		}

		bet, err := strconv.Atoi(userInput)

		if bet > player.Money {
			fmt.Println("You don't have that many coins")
			userInput = ""
		} else if err != nil || bet <= 0 {
			fmt.Println("Incorrect input")
			userInput = ""
		} else {
			player.Bet = bet
			player.Money -= bet
		}
	}
}

func (bj *Blackjack) betMakerAll() {
	for _, player := range bj.players {
		if player.Bot {
			bj.betMakerBot(player)
		} else {
			bj.betMakerPlayer(player)
		}
	}

	fmt.Printf("\n\nBets are made!\n\n")
}

func (bj *Blackjack) increaseNextDeckCardIndex() {
	bj.nextDeckCardIndex++
}

func (bj *Blackjack) giveCardToPlayer(player *Player, cardsNumber int) (*deck.Card, error) {
	playerCardsNumber := len(player.Cards)

	for i := playerCardsNumber; i < cardsNumber+playerCardsNumber; i++ {
		player.Cards = append(player.Cards, bj.deck[bj.nextDeckCardIndex])
		bj.increaseNextDeckCardIndex()
	}

	receivedCard := player.Cards[len(player.Cards)-cardsNumber:]
	return receivedCard[0], nil
}

func (bj *Blackjack) giveCardToDealer(cardsNumber int) (*deck.Card, error) {
	dealerCardsNumber := len(bj.dealer.Cards)

	for i := dealerCardsNumber; i < cardsNumber+dealerCardsNumber; i++ {
		bj.dealer.Cards = append(bj.dealer.Cards, bj.deck[bj.nextDeckCardIndex])
		bj.increaseNextDeckCardIndex()
	}

	receivedCard := bj.dealer.Cards[len(bj.dealer.Cards)-cardsNumber:]
	return receivedCard[0], nil
}

func (bj *Blackjack) giveCardsToAll(cardsNumber int) error {
	if cardsNumber > len(bj.deck) {
		return fmt.Errorf("cards number is larger than deck length")
	}

	for _, player := range bj.players {
		_, err := bj.giveCardToPlayer(player, cardsNumber)
		if err != nil {
			return err
		}
	}

	_, err := bj.giveCardToDealer(cardsNumber)
	if err != nil {
		return err
	}
	return nil
}

func (bj *Blackjack) printStartingCards() error {
	fmt.Printf("\nThe following cards were dealt:")

	for _, player := range bj.players {
		playerName := player.Name

		cardsPoints := 0

		if player.Id == bj.currentUser.Id {
			playerName = "Your cards"
		}

		printPlayerName(playerName)

		for _, card := range player.Cards {
			cost, err := getCardCost(card)
			if err != nil {
				return err
			} else {
				cardsPoints += cost
			}
			err = printCard(card)
			if err != nil {
				return err
			}
		}

		printTotalPoints(cardsPoints)
		time.Sleep(Delay)
	}

	// --- We print only the first card at the dealer

	printPlayerName("dealer")
	openedDealerCard := bj.dealer.Cards[0]
	err := printCard(openedDealerCard)
	if err != nil {
		return err
	}
	cardCost, err := getCardCost(openedDealerCard)
	if err != nil {
		return err
	}
	printTotalPoints(cardCost)

	time.Sleep(Delay)

	return nil
}

func (bj *Blackjack) printNewTurn() error {
	fmt.Println("\n\n-----------------------------")
	points, err := bj.getParticipantCardsPoints(bj.currentUser)
	if err != nil {
		return err
	}
	if !bj.currentUser.IsSaved {
		fmt.Printf("Turn: %d. Points: %d", bj.currentTurnIndex, points)
	}

	return nil
}

func (bj *Blackjack) playerSaved(player *Player) {
	player.IsSaved = true
}

func (bj *Blackjack) dealerSaved() {
	bj.dealer.IsSaved = true
}

func (bj *Blackjack) printPlayerCards(player *Player) error {
	totalPoints := 0

	for _, card := range player.Cards {
		cost, err := getCardCost(card)
		if err != nil {
			return err
		} else {
			totalPoints += cost
			fmt.Printf("%s %s. Give %d pts\n", card.Suit, card.Value, cost)
		}
	}

	fmt.Printf("Total: %d\n", totalPoints)
	return nil
}

func (bj *Blackjack) onUserInput(userInput string) (bool, error) {
	if bj.currentUser.IsSaved {
		return true, nil
	}

	if userInput == "" {
		return false, nil
	}

	switch userInput {
	case string(ActionExit):
		{
			fmt.Printf("\nWaiting you again!")
			os.Exit(0)
		}

	case string(ActionTakeCard):
		{
			receivedCard, err := bj.giveCardToPlayer(bj.currentUser, 1)
			if err != nil {
				return false, err
			}
			points, err := getCardCost(receivedCard)
			if err != nil {
				fmt.Println("error when getting card points: %w", err)
			}
			fmt.Printf("\nYou took the card %s %s. Give %d pts\n", receivedCard.Suit, receivedCard.Value, points)
			return true, nil
		}

	case string(ActionPass):
		{
			fmt.Printf("\nYou saved\n\n")
			bj.playerSaved(bj.currentUser)
			return true, nil
		}

	case string(ActionViewMyCards):
		{
			fmt.Println("\nYour cards:")
			err := bj.printPlayerCards(bj.currentUser)
			return false, err
		}

	default:
		{
			fmt.Println("Incorrect input")
			return false, nil
		}
	}

	return true, nil
}

func (bj *Blackjack) botTurn(bot *Player) error {
	botPoints, err := bj.getParticipantCardsPoints(bot)
	if err != nil {
		return err
	}

	if botPoints < 15 {
		fmt.Println("\nTake card...")
		time.Sleep(Delay)
		card, err := bj.giveCardToPlayer(bot, 1)
		if err != nil {
			return err
		}

		err = printCard(card)
		if err != nil {
			return err
		}
		fmt.Printf("\n\n\n")
	} else {
		fmt.Println("\nSaved")
		bj.playerSaved(bot)
	}

	return nil
}

func (bj *Blackjack) stageBots() error {
	// Отсекаем первого игрока, поскольку это пользователь
	bots := bj.players[1:]

	for _, bot := range bots {
		if bot.IsSaved == true || bot.IsLost {
			continue
		}

		fmt.Printf("\nBot`s turn: %s...", bot.Name)

		time.Sleep(Delay)
		err := bj.botTurn(bot)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bj *Blackjack) stageDealer() error {
	if bj.dealer.IsSaved {
		return fmt.Errorf("dealer is saved")
	}

	if !bj.isAllPlayersSaved {
		return nil
	}

	fmt.Println("\n\nDealer`s turn...")

	for {
		time.Sleep(Delay)

		dealerPoints, err := bj.dealer.getPoints()
		if err != nil {
			return err
		}

		if dealerPoints <= DealerPointsTakeCardLimit {
			fmt.Println("Dealer takes a card")
			card, err := bj.giveCardToDealer(1)
			if err != nil {
				return err
			}

			err = printCard(card)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Dealer not takes a card")
			bj.dealerSaved()

			return nil
		}
	}
}

func (bj *Blackjack) gameLoop() error {
	for {
		err := bj.checkRound()
		if err != nil {
			return err
		}

		if !bj.isStartingCardsDistributed {
			bj.betMakerAll()
			err := bj.giveCardsToAll(2)
			if err != nil {
				return err
			}

			bj.isStartingCardsDistributed = true

			err = bj.printStartingCards()
			if err != nil {
				return err
			}
		}

		err = bj.printNewTurn()
		if err != nil {
			return err
		}

		inputRes := false

		for !inputRes && !bj.currentUser.IsSaved {
			fmt.Printf("\nMoves:\n%s - Take card. %s - Save. %s - Your card. %s - exit.", ActionTakeCard, ActionPass, ActionViewMyCards, ActionExit)
			fmt.Printf("\n>> ")
			userInput := bj.console.Input()
			inputRes, err = bj.onUserInput(userInput)
			if err != nil {
				return err
			}
		}

		err = bj.stageBots()
		if err != nil {
			return err
		}

		err = bj.stageDealer()
		if err != nil {
			return err
		}

		bj.currentTurnIndex++
	}
}
