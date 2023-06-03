package blackjack

import (
	"course/internal/deck"
	"course/pkg/util"
	"fmt"
)

type Player struct {
	Id      string
	Cards   []*deck.Card
	Money   int
	Bet     int
	Name    string
	Bot     bool
	IsSaved bool
	IsLost  bool
}

func newPlayer(username string, money int, bot bool) (*Player, error) {
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	if money <= 0 {
		return nil, fmt.Errorf("user starting money is negative or equal to 0: %d", money)
	}

	return &Player{
		Id:      util.GetRandomString(),
		Cards:   nil,
		Money:   money,
		Bet:     0,
		Name:    username,
		Bot:     bot,
		IsSaved: false,
		IsLost:  false,
	}, nil
}

func (p *Player) getPoints() (int, error) {
	points := 0

	for _, card := range p.Cards {
		cost, err := getCardCost(card)
		if err != nil {
			return -1, err
		}
		points += cost
	}

	return points, nil
}

func (p *Player) resetRound() {
	p.Cards = []*deck.Card{}
	p.IsSaved = false
	p.Bet = 0
}

func (p *Player) checkIsLost() bool {
	if p.Money <= 0 {
		p.IsLost = true
	}

	return p.IsLost
}
