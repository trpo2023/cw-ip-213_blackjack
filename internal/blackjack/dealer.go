package blackjack

import (
	"course/internal/deck"
	"course/pkg/util"
)

type Dealer struct {
	Id      string
	Cards   []*deck.Card
	IsSaved bool
}

func newDealer() (*Dealer, error) {
	return &Dealer{
		Id:      util.GetRandomString(),
		Cards:   nil,
		IsSaved: false,
	}, nil
}

func (d *Dealer) getPoints() (int, error) {
	points := 0

	for _, card := range d.Cards {
		cost, err := getCardCost(card)
		if err != nil {
			return -1, err
		}
		points += cost
	}

	return points, nil
}

func (d *Dealer) resetRound() {
	d.IsSaved = false
	d.Cards = []*deck.Card{}
}

func (d *Dealer) printAllCards() error {
	for _, card := range d.Cards {
		err := printCard(card)
		if err != nil {
			return err
		}
	}

	return nil
}
