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
	"os"
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
	bj.gameLoop()

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
			fmt.Printf("Win!\n")
			player.Money += player.Bet + player.Bet
		} else if playerPoints == dealerPoints && playerPoints <= bj.MaxPoints {
			fmt.Printf("Draw\n")
			player.Money += player.Bet
		} else {
			fmt.Printf("Defeat\n")
		}
	}
}

func (bj *Blackjack) resetRound() {
	for i := range bj.Players {
		player := &bj.Players[i]
		player.Cards = []deck.Card{}
		player.Pass = false
		player.Bet = 0
	}

	bj.Dealer.Pass = false
	bj.Dealer.Cards = []deck.Card{}

	bj.Deck = deck.NewDeck(bj.DeckOptions)
	bj.NextDeckCardIndex = 0

	bj.IsStartingCardsDistributed = false
	bj.CurrentTurnIndex = 0
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

	cnsl := console.NewConsole()

	if allPass {
		bj.printRoundResults()
		lines := "--------------------------"
		fmt.Println("\n\n")
		fmt.Println(lines)
		fmt.Println("  New Round  ")
		fmt.Println(lines)
		bj.resetRound()
		fmt.Println("Press enter to continue...")
		cnsl.Input()
	}
}

func (bj *Blackjack) betMakerBot(bot *Player) {
	bet := rand.Intn(bot.Money)

	fmt.Printf("\n\nBot %s makes a bet...\n", bot.Name)
	time.Sleep(time.Second)

	fmt.Printf("An insert of %d coins was made", bet)

	bot.Bet = bet
	bot.Money -= bet
}

func (bj *Blackjack) betMakerPlayer(player *Player) {
	fmt.Printf("\nMake your bet (you have %d c.). %s - Exit.", player.Money, ActionExit)
	userInput := ""
	cnsl := console.NewConsole()

	for userInput == "" {
		fmt.Printf("\n>> ")
		userInput = cnsl.Input()

		if userInput == ActionExit {
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
	for i := range bj.Players {
		player := &bj.Players[i]

		if player.Bot {
			bj.betMakerBot(player)
		} else {
			bj.betMakerPlayer(player)
		}
	}

	fmt.Printf("\n\nBets are made!\n\n")
}

func (bj *Blackjack) increaseNextDeckCardIndex() {
	bj.NextDeckCardIndex++
}

func (bj *Blackjack) giveCardToPlayer(player *Player, cardsNumber int) []deck.Card {
	playerCardsNumber := len(player.Cards)

	for i := playerCardsNumber; i < cardsNumber+playerCardsNumber; i++ {
		player.Cards = append(player.Cards, bj.Deck[bj.NextDeckCardIndex])
		bj.increaseNextDeckCardIndex()
	}

	return player.Cards[len(player.Cards)-cardsNumber:]
}

func (bj *Blackjack) giveCardToDealer(cardsNumber int) []deck.Card {
	dealerCardsNumber := len(bj.Dealer.Cards)

	for i := dealerCardsNumber; i < cardsNumber+dealerCardsNumber; i++ {
		bj.Dealer.Cards = append(bj.Dealer.Cards, bj.Deck[bj.NextDeckCardIndex])
		bj.increaseNextDeckCardIndex()
	}

	return bj.Dealer.Cards[len(bj.Dealer.Cards)-cardsNumber:]
}

func (bj *Blackjack) giveCardsToAll(cardsNumber int) {
	for i := range bj.Players {
		player := &bj.Players[i]
		bj.giveCardToPlayer(player, cardsNumber)
	}
	bj.giveCardToDealer(cardsNumber)
}

func (bj *Blackjack) printStartingCards() {
	fmt.Printf("\nThe following cards were dealt:")

	printCard := func(card *deck.Card) {
		points, err := bj.getCardPoints(*card)
		if err != nil {
			fmt.Println("error getting points:", err)
		} else {
			fmt.Printf("\n%s %s. Gives %d points", card.Suit, card.Value, points)
		}
	}

	printPlayerName := func(playerName string) {
		fmt.Printf("\n\n%s:", playerName)
	}

	printTotalPoints := func(totalPoints int) {
		fmt.Printf("\nTotal: %d", totalPoints)
	}

	for i := range bj.Players {
		player := &bj.Players[i]
		playerName := player.Name

		cardsPoints := 0

		if player.Id == bj.CurrentPlayer.Id {
			playerName = "Your cards"
		}

		printPlayerName(playerName)

		for j := range player.Cards {
			card := &player.Cards[j]
			points, err := bj.getCardPoints(*card)
			if err != nil {
				fmt.Println("error getting points:", err)
			} else {
				cardsPoints += points
			}
			printCard(card)
		}

		printTotalPoints(cardsPoints)
	}

	// --- Дилер

	printPlayerName(bj.Dealer.Name)
	dealerPoints := 0

	for i := range bj.Dealer.Cards {
		card := &bj.Dealer.Cards[i]
		printCard(card)
		points, err := bj.getCardPoints(*card)
		if err != nil {
			fmt.Println("error getting points:", err)
		} else {
			dealerPoints += points
		}
	}

	printTotalPoints(dealerPoints)
}

func (bj *Blackjack) printNewTurn() {
	fmt.Printf("\n\nTurn: %d. Points: %d", bj.CurrentTurnIndex, bj.getPlayerCardsPoints(bj.CurrentPlayer))
}

func (bj *Blackjack) playerPass(player *Player) {
	player.Pass = true
}

func (bj *Blackjack) printPlayerCards(player *Player) {
	totalPoints := 0

	for i := range player.Cards {
		card := &player.Cards[i]
		points, err := bj.getCardPoints(*card)
		if err != nil {
			fmt.Println("error when getting card points: %w", err)
		} else {
			totalPoints += points
			fmt.Printf("%s %s. Give %d pts\n", card.Suit, card.Value, points)
		}
	}

	fmt.Printf("Total: %d\n", totalPoints)
}

func (bj *Blackjack) onUserInput(userInput string) bool {
	if bj.CurrentPlayer.Pass {
		return true
	}

	if userInput == "" {
		return false
	}

	switch userInput {
	case ActionExit:
		{
			fmt.Printf("\nWaiting you again!")
			os.Exit(0)
		}

	case ActionTakeCard:
		{
			receivedCards := bj.giveCardToPlayer(bj.CurrentPlayer, 1)
			receivedCard := receivedCards[0]
			points, err := bj.getCardPoints(receivedCard)
			if err != nil {
				fmt.Println("error when getting card points: %w", err)
			}
			fmt.Printf("\nYou took the card %s %s. Give %d pts\n", receivedCard.Suit, receivedCard.Value, points)
			return true
		}

	case ActionPass:
		{
			fmt.Printf("\nYou passed:")
			bj.playerPass(bj.CurrentPlayer)
			return true
		}

	case ActionViewMyCards:
		{
			fmt.Println("\nYour cards:")
			bj.printPlayerCards(bj.CurrentPlayer)
			return false
		}

	default:
		{
			fmt.Println("Incorrect input")
			return false
		}
	}

	return true
}

func (bj *Blackjack) botTurn(bot *Player) {
	botPoints := bj.getPlayerCardsPoints(bot)
	if botPoints < 15 {
		fmt.Println("\nTake card...")
		time.Sleep(time.Second)
		bj.giveCardToPlayer(bot, 1)
	} else {
		fmt.Println("\nPass")
		bj.playerPass(bot)
	}
}

func (bj *Blackjack) stageBots() {
	// Отсекаем первого игрока, поскольку это пользователь
	bots := bj.Players[1:]

	for i := range bots {
		bot := &bots[i]

		if bot.Pass == true {
			continue
		}

		fmt.Printf("\nBot`s turn: %s...", bot.Name)

		time.Sleep(time.Second)
		bj.botTurn(bot)
	}
}

func (bj *Blackjack) stageDealer() {
	dealer := &bj.Dealer
	dealerPoints := bj.getPlayerCardsPoints(*dealer)

	fmt.Println("\n\nDealer`s turn...")
	time.Sleep(time.Second)

	if dealerPoints <= DealerPointsTakeCardLimit {
		fmt.Println("Dealer takes a card")
		bj.giveCardToDealer(1)
	} else {
		fmt.Printf("Dealer not takes a card")
		bj.playerPass(*dealer)
	}
}

func (bj *Blackjack) gameLoop() {
	cnsl := console.NewConsole()

	for {
		bj.checkRound()

		if !bj.IsStartingCardsDistributed {
			bj.betMakerAll()
			bj.giveCardsToAll(2)
			bj.IsStartingCardsDistributed = true
			bj.printStartingCards()
		}

		bj.printNewTurn()

		userInput := ""

		for !bj.onUserInput(userInput) {
			fmt.Printf("\nMoves:\n%s - Take card. %s - Pass. %s - Your card. %s - exit.", ActionTakeCard, ActionPass, ActionViewMyCards, ActionExit)
			fmt.Printf("\n>> ")
			userInput = cnsl.Input()
		}

		bj.stageBots()
		bj.stageDealer()

		bj.CurrentTurnIndex++
	}
}
