package camelup

const (
	PlayerStartMoney int = 3
)

type Player struct {
	money int
}

type CamelDice interface {
	Roll() (camelIndex, steps int)
}

type Game struct {
	players []Player
	camelDice *CamelDice
}

func Init(numPlayers int) *Game {
	game := Game{
		players: make([]Player, numPlayers),
	}
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
