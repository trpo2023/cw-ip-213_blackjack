package blackjack

import "time"

type CardCost int

const (
	AceCostBig   CardCost = 11
	AceCostSmall CardCost = 1
	FaceCost     CardCost = 10
)

type Action string

const (
	ActionTakeCard    Action = "t"
	ActionPass        Action = "p"
	ActionExit        Action = "q"
	ActionViewMyCards Action = "c"
)

const (
	MaxPlayers                = 10
	DealerPointsTakeCardLimit = 16
	MaxPoints                 = 21
	Delay                     = 1 * time.Second
	LongDelay                 = 2 * time.Second
)

var DefaultPlayerNames = [...]string{"Anton", "Egor", "Karina", "Danil", "Kostya", "Masha", "Roma", "Sasha", "Oleg", "Zhenya", "Nastya", "Lisa", "Maks", "Dima", "Stas", "Anya", "Natasha", "Igor"}
