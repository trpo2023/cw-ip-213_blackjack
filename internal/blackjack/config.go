package blackjack

const (
	AceCostBig   = 11
	AceCostSmall = 1
	FaceCost     = 10
)

const (
	ActionTakeCard    = "t"
	ActionPass        = "p"
	ActionExit        = "q"
	ActionViewMyCards = "c"
)

const (
	MaxPlayers                = 10
	DealerPointsTakeCardLimit = 16
)

var DefaultPlayerNames = [...]string{"Anton", "Egor", "Karina", "Danil", "Kostya", "Masha", "Roma", "Sasha", "Oleg", "Zhenya", "Nastya", "Lisa", "Maks", "Dima", "Stas", "Anya", "Natasha", "Igor"}
