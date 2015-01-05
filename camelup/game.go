package camelup

const (
	PlayerStartMoney int = 3
)

type Player struct {
	money int
}

type Game struct {
	players []Player
}

func Init(numPlayers int) *Game {
	game := Game{}
	game.players = make([]Player, numPlayers)
	for i := range game.players {
		game.players[i].money = PlayerStartMoney
	}
	return &game
}

func (this *Game) Bet(camelIndex int) {
	// TODO implement
}

func (this *Game) Dice() {
	// TODO implement
}