package blackjack

import (
	"course/internal/deck"
	"fmt"
	"log"
	"strconv"
)

func getCardCost(card *deck.Card) (int, error) {
	switch card.Value {
	case deck.Ace:
		{
			return int(AceCostBig), nil
		}

	case deck.King:
		{
			return int(FaceCost), nil
		}

	case deck.Queen:
		{
			return int(FaceCost), nil
		}

	case deck.Jack:
		{
			return int(FaceCost), nil
		}

	default:
		{
			points, err := strconv.Atoi(string(card.Value))

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

func printCard(card *deck.Card) error {
	cost, err := getCardCost(card)
	if err != nil {
		return err
	} else {
		fmt.Printf("\n%s %s. Gives %d points", card.Suit, card.Value, cost)
	}

	return nil
}

func printPlayerName(playerName string) {
	fmt.Printf("\n\n%s:", playerName)
}

func printTotalPoints(totalPoints int) {
	fmt.Printf("\nTotal: %d", totalPoints)
}
